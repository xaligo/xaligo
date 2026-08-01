// Package icon implements the embedded SQLite SVG registry.
package icon

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xaligo/xaligo/internal/entity"
)

const (
	registrySchemaVersion = 1
	defaultBusyTimeoutMS  = 5_000
	maxSearchResults      = 100
)

var ErrNotFound = errors.New("icon not found")

type RegistryRepository interface {
	Get(context.Context, entity.IconRef) (entity.Icon, error)
	Search(context.Context, string, int) ([]entity.IconSummary, error)
	Put(context.Context, entity.Icon) (entity.Icon, error)
	Delete(context.Context, entity.IconRef) error
	List(context.Context, string, int) ([]entity.IconSummary, error)
	ListNamespaces(context.Context) ([]string, error)
	Close() error
}

type registryRepository struct {
	path    string
	openMu  sync.Mutex
	db      *sql.DB
	closed  bool
	writeMu sync.Mutex
}

func NewRegistryRepository(path string) RegistryRepository {
	return &registryRepository{path: path}
}

func (rcvr *registryRepository) Get(ctx context.Context, ref entity.IconRef) (entity.Icon, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return entity.Icon{}, err
	}
	row := db.QueryRowContext(ctx, `
WITH resolved AS (
    SELECT id AS icon_id, 0 AS priority
      FROM icons
     WHERE namespace = ? AND name = ?
    UNION ALL
    SELECT icon_id, 1 AS priority
      FROM icon_aliases
     WHERE namespace = ? AND alias = ?
)
SELECT i.id, i.namespace, i.name, i.description, i.svg, i.view_box,
       i.width, i.height, i.checksum, i.compression, i.license, i.source,
       i.created_at, i.updated_at
  FROM resolved r
  JOIN icons i ON i.id = r.icon_id
 ORDER BY r.priority, i.namespace, i.name
 LIMIT 1`, ref.Namespace, ref.Name, ref.Namespace, ref.Name)
	icon, err := scanIcon(row)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Icon{}, fmt.Errorf("%w: %s", ErrNotFound, ref.String())
	}
	if err != nil {
		return entity.Icon{}, fmt.Errorf("get icon %s: %w", ref.String(), err)
	}
	icon.Tags, err = loadTags(ctx, db, icon.ID)
	if err != nil {
		return entity.Icon{}, fmt.Errorf("load icon %s tags: %w", ref.String(), err)
	}
	icon.Aliases, err = loadAliases(ctx, db, icon.ID)
	if err != nil {
		return entity.Icon{}, fmt.Errorf("load icon %s aliases: %w", ref.String(), err)
	}
	return icon, nil
}

func (rcvr *registryRepository) Search(ctx context.Context, query string, limit int) ([]entity.IconSummary, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return nil, err
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, errors.New("icon search query must not be empty")
	}
	limit = normalizeLimit(limit)
	rows, err := db.QueryContext(ctx, `
SELECT i.namespace, i.name, i.description, i.view_box,
       i.width, i.height, i.license, i.source
  FROM icon_search s
  JOIN icons i ON i.id = s.rowid
 WHERE icon_search MATCH ?
 ORDER BY bm25(icon_search), i.namespace, i.name
 LIMIT ?`, query, limit)
	if err != nil {
		return nil, fmt.Errorf("search icons: %w", err)
	}
	defer rows.Close()
	return scanIconSummaries(rows)
}

func (rcvr *registryRepository) Put(ctx context.Context, icon entity.Icon) (entity.Icon, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return entity.Icon{}, err
	}
	rcvr.writeMu.Lock()
	defer rcvr.writeMu.Unlock()
	sortIconMetadata(&icon)

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return entity.Icon{}, fmt.Errorf("begin icon put: %w", err)
	}
	defer tx.Rollback()

	now := time.Now().UTC().UnixMilli()
	var createdAt, updatedAt int64
	err = tx.QueryRowContext(ctx, `
INSERT INTO icons (
    namespace, name, description, svg, view_box, width, height, checksum,
    compression, license, source, created_at, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(namespace, name) DO UPDATE SET
    description = excluded.description,
    svg = excluded.svg,
    view_box = excluded.view_box,
    width = excluded.width,
    height = excluded.height,
    checksum = excluded.checksum,
    compression = excluded.compression,
    license = excluded.license,
    source = excluded.source,
    updated_at = excluded.updated_at
RETURNING id, created_at, updated_at`,
		icon.Ref.Namespace, icon.Ref.Name, icon.Description, icon.SVG, icon.ViewBox,
		nullableFloat(icon.Width), nullableFloat(icon.Height), icon.Checksum[:],
		icon.Compression, icon.License, icon.Source, now, now,
	).Scan(&icon.ID, &createdAt, &updatedAt)
	if err != nil {
		return entity.Icon{}, fmt.Errorf("upsert icon %s: %w", icon.Ref.String(), err)
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM icon_tags WHERE icon_id = ?`, icon.ID); err != nil {
		return entity.Icon{}, fmt.Errorf("replace icon %s tags: %w", icon.Ref.String(), err)
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM icon_aliases WHERE icon_id = ?`, icon.ID); err != nil {
		return entity.Icon{}, fmt.Errorf("replace icon %s aliases: %w", icon.Ref.String(), err)
	}
	for _, tag := range icon.Tags {
		if _, err := tx.ExecContext(ctx, `INSERT INTO icon_tags(icon_id, tag) VALUES (?, ?)`, icon.ID, tag); err != nil {
			return entity.Icon{}, fmt.Errorf("insert icon %s tag %q: %w", icon.Ref.String(), tag, err)
		}
	}
	for _, alias := range icon.Aliases {
		if _, err := tx.ExecContext(ctx, `INSERT INTO icon_aliases(namespace, alias, icon_id) VALUES (?, ?, ?)`, alias.Namespace, alias.Name, icon.ID); err != nil {
			return entity.Icon{}, fmt.Errorf("insert icon %s alias %s: %w", icon.Ref.String(), alias.String(), err)
		}
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM icon_search WHERE rowid = ?`, icon.ID); err != nil {
		return entity.Icon{}, fmt.Errorf("replace icon %s search row: %w", icon.Ref.String(), err)
	}
	aliases := make([]string, 0, len(icon.Aliases)*2)
	for _, alias := range icon.Aliases {
		aliases = append(aliases, alias.Name, alias.String())
	}
	if _, err := tx.ExecContext(ctx, `
INSERT INTO icon_search(rowid, namespace, name, description, tags, aliases)
VALUES (?, ?, ?, ?, ?, ?)`, icon.ID, icon.Ref.Namespace, icon.Ref.Name,
		icon.Description, strings.Join(icon.Tags, " "), strings.Join(aliases, " ")); err != nil {
		return entity.Icon{}, fmt.Errorf("index icon %s: %w", icon.Ref.String(), err)
	}
	if err := tx.Commit(); err != nil {
		return entity.Icon{}, fmt.Errorf("commit icon %s: %w", icon.Ref.String(), err)
	}
	icon.CreatedAt = time.UnixMilli(createdAt).UTC()
	icon.UpdatedAt = time.UnixMilli(updatedAt).UTC()
	return icon, nil
}

func (rcvr *registryRepository) Delete(ctx context.Context, ref entity.IconRef) error {
	db, err := rcvr.database(ctx)
	if err != nil {
		return err
	}
	rcvr.writeMu.Lock()
	defer rcvr.writeMu.Unlock()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin icon delete: %w", err)
	}
	defer tx.Rollback()
	var id int64
	if err := tx.QueryRowContext(ctx, `SELECT id FROM icons WHERE namespace = ? AND name = ?`, ref.Namespace, ref.Name).Scan(&id); errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: %s", ErrNotFound, ref.String())
	} else if err != nil {
		return fmt.Errorf("resolve icon %s delete: %w", ref.String(), err)
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM icon_search WHERE rowid = ?`, id); err != nil {
		return fmt.Errorf("delete icon %s search row: %w", ref.String(), err)
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM icons WHERE id = ?`, id); err != nil {
		return fmt.Errorf("delete icon %s: %w", ref.String(), err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit icon %s delete: %w", ref.String(), err)
	}
	return nil
}

func (rcvr *registryRepository) List(ctx context.Context, namespace string, limit int) ([]entity.IconSummary, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return nil, err
	}
	limit = normalizeLimit(limit)
	query := `
SELECT namespace, name, description, view_box, width, height, license, source
  FROM icons`
	arguments := []any{}
	if namespace != "" {
		query += ` WHERE namespace = ?`
		arguments = append(arguments, namespace)
	}
	query += ` ORDER BY namespace, name LIMIT ?`
	arguments = append(arguments, limit)
	rows, err := db.QueryContext(ctx, query, arguments...)
	if err != nil {
		return nil, fmt.Errorf("list icons: %w", err)
	}
	defer rows.Close()
	return scanIconSummaries(rows)
}

func (rcvr *registryRepository) ListNamespaces(ctx context.Context) ([]string, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, `SELECT DISTINCT namespace FROM icons ORDER BY namespace`)
	if err != nil {
		return nil, fmt.Errorf("list icon namespaces: %w", err)
	}
	defer rows.Close()
	var namespaces []string
	for rows.Next() {
		var namespace string
		if err := rows.Scan(&namespace); err != nil {
			return nil, fmt.Errorf("scan icon namespace: %w", err)
		}
		namespaces = append(namespaces, namespace)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate icon namespaces: %w", err)
	}
	return namespaces, nil
}

func (rcvr *registryRepository) Close() error {
	rcvr.openMu.Lock()
	defer rcvr.openMu.Unlock()
	rcvr.closed = true
	if rcvr.db == nil {
		return nil
	}
	err := rcvr.db.Close()
	rcvr.db = nil
	if err != nil {
		return fmt.Errorf("close icon registry: %w", err)
	}
	return nil
}

func (rcvr *registryRepository) database(ctx context.Context) (*sql.DB, error) {
	rcvr.openMu.Lock()
	defer rcvr.openMu.Unlock()
	if rcvr.closed {
		return nil, errors.New("icon registry is closed")
	}
	if rcvr.db != nil {
		return rcvr.db, nil
	}
	db, err := openRegistry(ctx, rcvr.path)
	if err != nil {
		return nil, err
	}
	rcvr.db = db
	return db, nil
}

func openRegistry(ctx context.Context, path string) (*sql.DB, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil, errors.New("icon registry path must not be empty")
	}
	absolute, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve icon registry path: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(absolute), 0o750); err != nil {
		return nil, fmt.Errorf("create icon registry directory: %w", err)
	}
	location := (&url.URL{Scheme: "file", Path: filepath.ToSlash(absolute)}).String()
	parameters := url.Values{}
	parameters.Set("_busy_timeout", fmt.Sprint(defaultBusyTimeoutMS))
	parameters.Set("_foreign_keys", "on")
	parameters.Set("_journal_mode", "WAL")
	parameters.Set("_synchronous", "NORMAL")
	parameters.Set("_txlock", "immediate")
	db, err := sql.Open("sqlite3", location+"?"+parameters.Encode())
	if err != nil {
		return nil, fmt.Errorf("open icon registry: %w", err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping icon registry: %w", err)
	}
	var fts5 int
	if err := db.QueryRowContext(ctx, `SELECT sqlite_compileoption_used('ENABLE_FTS5')`).Scan(&fts5); err != nil {
		db.Close()
		return nil, fmt.Errorf("check SQLite FTS5 support: %w", err)
	}
	if fts5 != 1 {
		db.Close()
		return nil, errors.New("SQLite FTS5 is unavailable; build with the sqlite_fts5 tag")
	}
	if err := migrateRegistry(ctx, db); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func migrateRegistry(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin icon registry migration: %w", err)
	}
	defer tx.Rollback()
	var version int
	if err := tx.QueryRowContext(ctx, `PRAGMA user_version`).Scan(&version); err != nil {
		return fmt.Errorf("read icon registry schema version: %w", err)
	}
	if version > registrySchemaVersion {
		return fmt.Errorf("icon registry schema version %d is newer than supported version %d", version, registrySchemaVersion)
	}
	if version == registrySchemaVersion {
		return tx.Commit()
	}
	statements := []string{
		`CREATE TABLE IF NOT EXISTS icons (
            id INTEGER PRIMARY KEY,
            namespace TEXT NOT NULL,
            name TEXT NOT NULL,
            description TEXT NOT NULL DEFAULT '',
            svg BLOB NOT NULL,
            view_box TEXT NOT NULL,
            width REAL,
            height REAL,
            checksum BLOB NOT NULL,
            compression INTEGER NOT NULL DEFAULT 0,
            license TEXT NOT NULL DEFAULT '',
            source TEXT NOT NULL DEFAULT '',
            created_at INTEGER NOT NULL,
            updated_at INTEGER NOT NULL,
            UNIQUE(namespace, name)
        )`,
		`CREATE TABLE IF NOT EXISTS icon_tags (
            icon_id INTEGER NOT NULL REFERENCES icons(id) ON DELETE CASCADE,
            tag TEXT NOT NULL,
            PRIMARY KEY (icon_id, tag)
        )`,
		`CREATE TABLE IF NOT EXISTS icon_aliases (
            namespace TEXT NOT NULL,
            alias TEXT NOT NULL,
            icon_id INTEGER NOT NULL REFERENCES icons(id) ON DELETE CASCADE,
            PRIMARY KEY (namespace, alias)
        )`,
		`CREATE INDEX IF NOT EXISTS idx_icons_namespace_name ON icons(namespace, name)`,
		`CREATE INDEX IF NOT EXISTS idx_icon_aliases_icon_id ON icon_aliases(icon_id)`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS icon_search USING fts5(
            namespace, name, description, tags, aliases, tokenize = 'unicode61'
        )`,
	}
	for _, statement := range statements {
		if _, err := tx.ExecContext(ctx, statement); err != nil {
			return fmt.Errorf("apply icon registry schema version 1: %w", err)
		}
	}
	if _, err := tx.ExecContext(ctx, `PRAGMA user_version = 1`); err != nil {
		return fmt.Errorf("set icon registry schema version: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit icon registry migration: %w", err)
	}
	return nil
}

type rowScanner interface {
	Scan(...any) error
}

func scanIcon(scanner rowScanner) (entity.Icon, error) {
	var icon entity.Icon
	var width, height sql.NullFloat64
	var checksum []byte
	var createdAt, updatedAt int64
	err := scanner.Scan(
		&icon.ID, &icon.Ref.Namespace, &icon.Ref.Name, &icon.Description, &icon.SVG,
		&icon.ViewBox, &width, &height, &checksum, &icon.Compression, &icon.License,
		&icon.Source, &createdAt, &updatedAt,
	)
	if err != nil {
		return entity.Icon{}, err
	}
	if len(checksum) != len(icon.Checksum) {
		return entity.Icon{}, fmt.Errorf("icon checksum has %d bytes", len(checksum))
	}
	copy(icon.Checksum[:], checksum)
	icon.Width = floatPointer(width)
	icon.Height = floatPointer(height)
	icon.CreatedAt = time.UnixMilli(createdAt).UTC()
	icon.UpdatedAt = time.UnixMilli(updatedAt).UTC()
	return icon, nil
}

func loadTags(ctx context.Context, db *sql.DB, iconID int64) ([]string, error) {
	rows, err := db.QueryContext(ctx, `SELECT tag FROM icon_tags WHERE icon_id = ? ORDER BY tag`, iconID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, rows.Err()
}

func loadAliases(ctx context.Context, db *sql.DB, iconID int64) ([]entity.IconRef, error) {
	rows, err := db.QueryContext(ctx, `SELECT namespace, alias FROM icon_aliases WHERE icon_id = ? ORDER BY namespace, alias`, iconID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var aliases []entity.IconRef
	for rows.Next() {
		var alias entity.IconRef
		if err := rows.Scan(&alias.Namespace, &alias.Name); err != nil {
			return nil, err
		}
		aliases = append(aliases, alias)
	}
	return aliases, rows.Err()
}

func scanIconSummaries(rows *sql.Rows) ([]entity.IconSummary, error) {
	var summaries []entity.IconSummary
	for rows.Next() {
		var summary entity.IconSummary
		var width, height sql.NullFloat64
		if err := rows.Scan(
			&summary.Ref.Namespace, &summary.Ref.Name, &summary.Description,
			&summary.ViewBox, &width, &height, &summary.License, &summary.Source,
		); err != nil {
			return nil, fmt.Errorf("scan icon summary: %w", err)
		}
		summary.Width = floatPointer(width)
		summary.Height = floatPointer(height)
		summaries = append(summaries, summary)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate icon summaries: %w", err)
	}
	return summaries, nil
}

func normalizeLimit(limit int) int {
	if limit <= 0 {
		return 30
	}
	if limit > maxSearchResults {
		return maxSearchResults
	}
	return limit
}

func nullableFloat(value *float64) any {
	if value == nil {
		return nil
	}
	return *value
}

func floatPointer(value sql.NullFloat64) *float64 {
	if !value.Valid {
		return nil
	}
	result := value.Float64
	return &result
}

func sortIconMetadata(icon *entity.Icon) {
	sort.Strings(icon.Tags)
	sort.Slice(icon.Aliases, func(left, right int) bool {
		return icon.Aliases[left].String() < icon.Aliases[right].String()
	})
}
