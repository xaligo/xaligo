// Package project implements the durable SQLite project concept index shared
// by RAG, LSP workspace search, and MCP inspection tools.
package project

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xaligo/xaligo/internal/entity"
)

const (
	indexSchemaVersion = 1
	defaultBusyTimeout = 5_000
	maxSearchResults   = 100
)

// IndexRepository persists compact source concepts and their FTS projection.
type IndexRepository interface {
	Put(context.Context, string, entity.ProjectAnalysis) (bool, error)
	Search(context.Context, string, int) ([]entity.ProjectSearchResult, error)
	Symbols(context.Context, string) ([]entity.ProjectSymbol, error)
	Prune(context.Context, string, []string) (int, error)
	Close() error
}

type indexRepository struct {
	path    string
	openMu  sync.Mutex
	db      *sql.DB
	closed  bool
	writeMu sync.Mutex
}

func NewIndexRepository(path string) IndexRepository {
	return &indexRepository{path: path}
}

func (rcvr *indexRepository) Put(ctx context.Context, rootURI string, analysis entity.ProjectAnalysis) (bool, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return false, err
	}
	if strings.TrimSpace(analysis.URI) == "" {
		return false, errors.New("project document URI must not be empty")
	}
	rcvr.writeMu.Lock()
	defer rcvr.writeMu.Unlock()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return false, fmt.Errorf("begin project document index: %w", err)
	}
	defer tx.Rollback()

	var current []byte
	err = tx.QueryRowContext(ctx, `SELECT checksum FROM project_documents WHERE uri = ?`, analysis.URI).Scan(&current)
	if err == nil && bytes.Equal(current, analysis.Checksum[:]) {
		if err := tx.Commit(); err != nil {
			return false, fmt.Errorf("commit unchanged project document: %w", err)
		}
		return false, nil
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("read project document checksum: %w", err)
	}

	now := time.Now().UTC().UnixMilli()
	var documentID int64
	err = tx.QueryRowContext(ctx, `
INSERT INTO project_documents(uri, root_uri, kind, checksum, source, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT(uri) DO UPDATE SET
    root_uri = excluded.root_uri,
    kind = excluded.kind,
    checksum = excluded.checksum,
    source = excluded.source,
    updated_at = excluded.updated_at
RETURNING id`, analysis.URI, rootURI, string(analysis.Kind), analysis.Checksum[:], analysis.Source, now).Scan(&documentID)
	if err != nil {
		return false, fmt.Errorf("upsert project document %s: %w", analysis.URI, err)
	}
	if _, err := tx.ExecContext(ctx, `
DELETE FROM project_search
 WHERE rowid IN (SELECT id FROM project_symbols WHERE document_id = ?)`, documentID); err != nil {
		return false, fmt.Errorf("clear project search rows for %s: %w", analysis.URI, err)
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM project_symbols WHERE document_id = ?`, documentID); err != nil {
		return false, fmt.Errorf("clear project symbols for %s: %w", analysis.URI, err)
	}

	for _, symbol := range analysis.Symbols {
		var symbolID int64
		err := tx.QueryRowContext(ctx, `
INSERT INTO project_symbols(
    document_id, ordinal, parent_ordinal, semantic_id, name, detail, concept,
    source_tag, source_ref, target_ref, line, column_no, offset_no
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id`, documentID, symbol.Ordinal, symbol.ParentOrdinal, symbol.ID,
			symbol.Name, symbol.Detail, string(symbol.Concept), symbol.SourceTag,
			symbol.Source, symbol.Target, symbol.Position.Line, symbol.Position.Column,
			symbol.Position.Offset).Scan(&symbolID)
		if err != nil {
			return false, fmt.Errorf("insert project symbol %s in %s: %w", symbol.ID, analysis.URI, err)
		}
		searchBody := compactSearchBody(symbol)
		if _, err := tx.ExecContext(ctx, `
INSERT INTO project_search(rowid, uri, semantic_id, name, concept, source_tag, body)
VALUES (?, ?, ?, ?, ?, ?, ?)`, symbolID, analysis.URI, symbol.ID, symbol.Name,
			string(symbol.Concept), symbol.SourceTag, searchBody); err != nil {
			return false, fmt.Errorf("index project symbol %s in %s: %w", symbol.ID, analysis.URI, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return false, fmt.Errorf("commit project document %s: %w", analysis.URI, err)
	}
	return true, nil
}

func (rcvr *indexRepository) Search(ctx context.Context, query string, limit int) ([]entity.ProjectSearchResult, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return nil, err
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, errors.New("project search query must not be empty")
	}
	limit = normalizeLimit(limit)
	rows, err := db.QueryContext(ctx, `
SELECT d.uri, s.semantic_id, s.name, s.detail, s.concept, s.source_tag,
       s.line, s.column_no, bm25(project_search)
  FROM project_search
  JOIN project_symbols s ON s.id = project_search.rowid
  JOIN project_documents d ON d.id = s.document_id
 WHERE project_search MATCH ?
 ORDER BY bm25(project_search), d.uri, s.ordinal
 LIMIT ?`, query, limit)
	if err != nil {
		return nil, fmt.Errorf("search project index: %w", err)
	}
	defer rows.Close()
	results := make([]entity.ProjectSearchResult, 0, limit)
	for rows.Next() {
		var result entity.ProjectSearchResult
		if err := rows.Scan(&result.URI, &result.ID, &result.Name, &result.Detail,
			&result.Concept, &result.SourceTag, &result.Line, &result.Column,
			&result.Score); err != nil {
			return nil, fmt.Errorf("scan project search result: %w", err)
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate project search results: %w", err)
	}
	return results, nil
}

func (rcvr *indexRepository) Symbols(ctx context.Context, uri string) ([]entity.ProjectSymbol, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, `
SELECT s.ordinal, s.parent_ordinal, s.semantic_id, s.name, s.detail, s.concept,
       s.source_tag, s.source_ref, s.target_ref, s.offset_no, s.line, s.column_no
  FROM project_symbols s
  JOIN project_documents d ON d.id = s.document_id
 WHERE d.uri = ?
 ORDER BY s.ordinal`, uri)
	if err != nil {
		return nil, fmt.Errorf("list project symbols for %s: %w", uri, err)
	}
	defer rows.Close()
	var symbols []entity.ProjectSymbol
	for rows.Next() {
		var symbol entity.ProjectSymbol
		if err := rows.Scan(&symbol.Ordinal, &symbol.ParentOrdinal, &symbol.ID,
			&symbol.Name, &symbol.Detail, &symbol.Concept, &symbol.SourceTag,
			&symbol.Source, &symbol.Target, &symbol.Position.Offset,
			&symbol.Position.Line, &symbol.Position.Column); err != nil {
			return nil, fmt.Errorf("scan project symbol: %w", err)
		}
		symbols = append(symbols, symbol)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate project symbols: %w", err)
	}
	return symbols, nil
}

func (rcvr *indexRepository) Prune(ctx context.Context, rootURI string, keep []string) (int, error) {
	db, err := rcvr.database(ctx)
	if err != nil {
		return 0, err
	}
	kept := make(map[string]struct{}, len(keep))
	for _, uri := range keep {
		kept[uri] = struct{}{}
	}
	rcvr.writeMu.Lock()
	defer rcvr.writeMu.Unlock()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("begin project prune: %w", err)
	}
	defer tx.Rollback()
	rows, err := tx.QueryContext(ctx, `SELECT id, uri FROM project_documents WHERE root_uri = ? ORDER BY uri`, rootURI)
	if err != nil {
		return 0, fmt.Errorf("list indexed project documents: %w", err)
	}
	type document struct {
		id  int64
		uri string
	}
	var removed []document
	for rows.Next() {
		var candidate document
		if err := rows.Scan(&candidate.id, &candidate.uri); err != nil {
			rows.Close()
			return 0, fmt.Errorf("scan indexed project document: %w", err)
		}
		if _, exists := kept[candidate.uri]; !exists {
			removed = append(removed, candidate)
		}
	}
	if err := rows.Close(); err != nil {
		return 0, fmt.Errorf("close project prune rows: %w", err)
	}
	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("iterate indexed project documents: %w", err)
	}
	for _, document := range removed {
		if _, err := tx.ExecContext(ctx, `DELETE FROM project_search WHERE rowid IN (SELECT id FROM project_symbols WHERE document_id = ?)`, document.id); err != nil {
			return 0, fmt.Errorf("delete project search rows for %s: %w", document.uri, err)
		}
		if _, err := tx.ExecContext(ctx, `DELETE FROM project_documents WHERE id = ?`, document.id); err != nil {
			return 0, fmt.Errorf("delete project document %s: %w", document.uri, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("commit project prune: %w", err)
	}
	return len(removed), nil
}

func (rcvr *indexRepository) Close() error {
	rcvr.openMu.Lock()
	defer rcvr.openMu.Unlock()
	rcvr.closed = true
	if rcvr.db == nil {
		return nil
	}
	err := rcvr.db.Close()
	rcvr.db = nil
	if err != nil {
		return fmt.Errorf("close project index: %w", err)
	}
	return nil
}

func (rcvr *indexRepository) database(ctx context.Context) (*sql.DB, error) {
	rcvr.openMu.Lock()
	defer rcvr.openMu.Unlock()
	if rcvr.closed {
		return nil, errors.New("project index is closed")
	}
	if rcvr.db != nil {
		return rcvr.db, nil
	}
	db, err := openIndex(ctx, rcvr.path)
	if err != nil {
		return nil, err
	}
	rcvr.db = db
	return db, nil
}

func openIndex(ctx context.Context, path string) (*sql.DB, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil, errors.New("project index path must not be empty")
	}
	absolute, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve project index path: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(absolute), 0o750); err != nil {
		return nil, fmt.Errorf("create project index directory: %w", err)
	}
	location := (&url.URL{Scheme: "file", Path: filepath.ToSlash(absolute)}).String()
	parameters := url.Values{}
	parameters.Set("_busy_timeout", fmt.Sprint(defaultBusyTimeout))
	parameters.Set("_foreign_keys", "on")
	parameters.Set("_journal_mode", "WAL")
	parameters.Set("_synchronous", "NORMAL")
	parameters.Set("_txlock", "immediate")
	db, err := sql.Open("sqlite3", location+"?"+parameters.Encode())
	if err != nil {
		return nil, fmt.Errorf("open project index: %w", err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping project index: %w", err)
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
	if err := migrateIndex(ctx, db); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func migrateIndex(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin project index migration: %w", err)
	}
	defer tx.Rollback()
	var version int
	if err := tx.QueryRowContext(ctx, `PRAGMA user_version`).Scan(&version); err != nil {
		return fmt.Errorf("read project index schema version: %w", err)
	}
	if version > indexSchemaVersion {
		return fmt.Errorf("project index schema version %d is newer than supported version %d", version, indexSchemaVersion)
	}
	if version == indexSchemaVersion {
		return tx.Commit()
	}
	statements := []string{
		`CREATE TABLE IF NOT EXISTS project_documents (
            id INTEGER PRIMARY KEY,
            uri TEXT NOT NULL UNIQUE,
            root_uri TEXT NOT NULL,
            kind TEXT NOT NULL,
            checksum BLOB NOT NULL,
            source BLOB NOT NULL,
            updated_at INTEGER NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS project_symbols (
            id INTEGER PRIMARY KEY,
            document_id INTEGER NOT NULL REFERENCES project_documents(id) ON DELETE CASCADE,
            ordinal INTEGER NOT NULL,
            parent_ordinal INTEGER NOT NULL,
            semantic_id TEXT NOT NULL,
            name TEXT NOT NULL,
            detail TEXT NOT NULL,
            concept TEXT NOT NULL,
            source_tag TEXT NOT NULL,
            source_ref TEXT NOT NULL,
            target_ref TEXT NOT NULL,
            line INTEGER NOT NULL,
            column_no INTEGER NOT NULL,
            offset_no INTEGER NOT NULL,
            UNIQUE(document_id, ordinal)
        )`,
		`CREATE INDEX IF NOT EXISTS idx_project_documents_root_uri ON project_documents(root_uri, uri)`,
		`CREATE INDEX IF NOT EXISTS idx_project_symbols_document ON project_symbols(document_id, ordinal)`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS project_search USING fts5(
            uri, semantic_id, name, concept, source_tag, body,
            tokenize = 'unicode61'
        )`,
	}
	for _, statement := range statements {
		if _, err := tx.ExecContext(ctx, statement); err != nil {
			return fmt.Errorf("apply project index schema version 1: %w", err)
		}
	}
	if _, err := tx.ExecContext(ctx, `PRAGMA user_version = 1`); err != nil {
		return fmt.Errorf("set project index schema version: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit project index migration: %w", err)
	}
	return nil
}

func compactSearchBody(symbol entity.ProjectSymbol) string {
	parts := []string{symbol.Name, symbol.Detail, symbol.Source, symbol.Target}
	return strings.Join(parts, " ")
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
