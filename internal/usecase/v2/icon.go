package v2

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"unicode"

	"github.com/xaligo/xaligo/internal/entity"
	iconrepository "github.com/xaligo/xaligo/internal/repository/icon"
)

type IconUsecase interface {
	Put(context.Context, entity.IconRegistration) (entity.Icon, error)
	Get(context.Context, string) (entity.Icon, error)
	Search(context.Context, string, int) ([]entity.IconSummary, error)
	Delete(context.Context, string) error
	List(context.Context, string, int) ([]entity.IconSummary, error)
	ListNamespaces(context.Context) ([]string, error)
}

type iconUsecase struct {
	registry iconrepository.RegistryRepository
	engine   EngineUsecase
	defaults []entity.IconRegistration
	seedMu   sync.Mutex
	seeded   bool
}

func NewIconUsecase(registry iconrepository.RegistryRepository, engine EngineUsecase, defaults ...entity.IconRegistration) IconUsecase {
	return &iconUsecase{
		registry: registry,
		engine:   engine,
		defaults: append([]entity.IconRegistration(nil), defaults...),
	}
}

const (
	maxIconDescriptionBytes = 4 * 1024
	maxIconTags             = 64
	maxIconAliases          = 64
	maxIconTagBytes         = 128
	maxIconSearchBytes      = 512
)

func (rcvr *iconUsecase) Put(ctx context.Context, registration entity.IconRegistration) (entity.Icon, error) {
	if err := checkEngineContext(ctx); err != nil {
		return entity.Icon{}, err
	}
	if err := rcvr.ensureDefaults(ctx); err != nil {
		return entity.Icon{}, err
	}
	return rcvr.put(ctx, registration)
}

func (rcvr *iconUsecase) put(ctx context.Context, registration entity.IconRegistration) (entity.Icon, error) {
	if rcvr.registry == nil {
		return entity.Icon{}, errors.New("icon registry repository is required")
	}
	if rcvr.engine == nil {
		return entity.Icon{}, errors.New("V2 Rust engine is required for SVG registration")
	}
	ref, err := parseIconRef(registration.Reference, "")
	if err != nil {
		return entity.Icon{}, err
	}
	if len(registration.Description) > maxIconDescriptionBytes {
		return entity.Icon{}, fmt.Errorf("icon description exceeds %d UTF-8 bytes", maxIconDescriptionBytes)
	}
	tags, err := normalizeIconTags(registration.Tags)
	if err != nil {
		return entity.Icon{}, err
	}
	aliases, err := normalizeIconAliases(registration.Aliases, ref)
	if err != nil {
		return entity.Icon{}, err
	}
	normalized, err := rcvr.engine.NormalizeSVG(ctx, registration.SVG)
	if err != nil {
		return entity.Icon{}, fmt.Errorf("validate icon %s SVG: %w", ref.String(), err)
	}
	width := normalized.Width
	height := normalized.Height
	icon := entity.Icon{
		Ref:         ref,
		Description: strings.TrimSpace(registration.Description),
		SVG:         normalized.Data,
		ViewBox:     normalized.ViewBox,
		Width:       &width,
		Height:      &height,
		Checksum:    sha256.Sum256(normalized.Data),
		Compression: 0,
		License:     strings.TrimSpace(registration.License),
		Source:      strings.TrimSpace(registration.Source),
		Tags:        tags,
		Aliases:     aliases,
	}
	stored, err := rcvr.registry.Put(ctx, icon)
	if err != nil {
		return entity.Icon{}, fmt.Errorf("store icon %s: %w", ref.String(), err)
	}
	return stored, nil
}

func (rcvr *iconUsecase) Get(ctx context.Context, reference string) (entity.Icon, error) {
	if err := rcvr.ensureDefaults(ctx); err != nil {
		return entity.Icon{}, err
	}
	if rcvr.registry == nil {
		return entity.Icon{}, errors.New("icon registry repository is required")
	}
	ref, err := parseIconRef(reference, "")
	if err != nil {
		return entity.Icon{}, err
	}
	icon, err := rcvr.registry.Get(ctx, ref)
	if err != nil {
		return entity.Icon{}, fmt.Errorf("get icon %s: %w", ref.String(), err)
	}
	return icon, nil
}

func (rcvr *iconUsecase) Search(ctx context.Context, query string, limit int) ([]entity.IconSummary, error) {
	if err := rcvr.ensureDefaults(ctx); err != nil {
		return nil, err
	}
	if rcvr.registry == nil {
		return nil, errors.New("icon registry repository is required")
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, errors.New("icon search query must not be empty")
	}
	if len(query) > maxIconSearchBytes {
		return nil, fmt.Errorf("icon search query exceeds %d UTF-8 bytes", maxIconSearchBytes)
	}
	results, err := rcvr.registry.Search(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("search icons: %w", err)
	}
	return results, nil
}

func (rcvr *iconUsecase) Delete(ctx context.Context, reference string) error {
	if err := rcvr.ensureDefaults(ctx); err != nil {
		return err
	}
	if rcvr.registry == nil {
		return errors.New("icon registry repository is required")
	}
	ref, err := parseIconRef(reference, "")
	if err != nil {
		return err
	}
	if err := rcvr.registry.Delete(ctx, ref); err != nil {
		return fmt.Errorf("delete icon %s: %w", ref.String(), err)
	}
	return nil
}

func (rcvr *iconUsecase) List(ctx context.Context, namespace string, limit int) ([]entity.IconSummary, error) {
	if err := rcvr.ensureDefaults(ctx); err != nil {
		return nil, err
	}
	if rcvr.registry == nil {
		return nil, errors.New("icon registry repository is required")
	}
	var err error
	if strings.TrimSpace(namespace) != "" {
		namespace, err = normalizeIconPart(namespace, false, "namespace")
		if err != nil {
			return nil, err
		}
	}
	results, err := rcvr.registry.List(ctx, namespace, limit)
	if err != nil {
		return nil, fmt.Errorf("list icons: %w", err)
	}
	return results, nil
}

func (rcvr *iconUsecase) ListNamespaces(ctx context.Context) ([]string, error) {
	if err := rcvr.ensureDefaults(ctx); err != nil {
		return nil, err
	}
	if rcvr.registry == nil {
		return nil, errors.New("icon registry repository is required")
	}
	namespaces, err := rcvr.registry.ListNamespaces(ctx)
	if err != nil {
		return nil, fmt.Errorf("list icon namespaces: %w", err)
	}
	return namespaces, nil
}

func (rcvr *iconUsecase) ensureDefaults(ctx context.Context) error {
	if len(rcvr.defaults) == 0 {
		return nil
	}
	if rcvr.registry == nil {
		return errors.New("icon registry repository is required")
	}
	if rcvr.engine == nil {
		return errors.New("V2 Rust engine is required for builtin SVG registration")
	}
	rcvr.seedMu.Lock()
	defer rcvr.seedMu.Unlock()
	if rcvr.seeded {
		return nil
	}
	if err := checkEngineContext(ctx); err != nil {
		return err
	}
	registrations := append([]entity.IconRegistration(nil), rcvr.defaults...)
	sort.SliceStable(registrations, func(left, right int) bool {
		return registrations[left].Reference < registrations[right].Reference
	})
	seen := make(map[entity.IconRef]struct{}, len(registrations))
	for _, registration := range registrations {
		ref, err := parseIconRef(registration.Reference, "")
		if err != nil {
			return fmt.Errorf("invalid default icon %q: %w", registration.Reference, err)
		}
		if _, duplicate := seen[ref]; duplicate {
			return fmt.Errorf("duplicate default icon %s", ref.String())
		}
		seen[ref] = struct{}{}
		if _, err := rcvr.registry.Get(ctx, ref); err == nil {
			continue
		} else if !errors.Is(err, iconrepository.ErrNotFound) {
			return fmt.Errorf("inspect default icon %s: %w", ref.String(), err)
		}
		if _, err := rcvr.put(ctx, registration); err != nil {
			return fmt.Errorf("register default icon %s: %w", ref.String(), err)
		}
	}
	rcvr.seeded = true
	return nil
}

func parseIconRef(value, defaultNamespace string) (entity.IconRef, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return entity.IconRef{}, errors.New("icon reference must not be empty")
	}
	namespace, name, found := strings.Cut(value, ":")
	if !found {
		if defaultNamespace == "" {
			return entity.IconRef{}, fmt.Errorf("icon reference %q must use namespace:name", value)
		}
		namespace = defaultNamespace
		name = value
	}
	if strings.Contains(name, ":") {
		return entity.IconRef{}, fmt.Errorf("icon reference %q contains more than one namespace delimiter", value)
	}
	namespace, err := normalizeIconPart(namespace, false, "namespace")
	if err != nil {
		return entity.IconRef{}, err
	}
	name, err = normalizeIconPart(name, true, "name")
	if err != nil {
		return entity.IconRef{}, err
	}
	return entity.IconRef{Namespace: namespace, Name: name}, nil
}

func normalizeIconPart(value string, allowSlash bool, label string) (string, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	limit := 64
	if allowSlash {
		limit = 128
	}
	if len(value) == 0 || len(value) > limit {
		return "", fmt.Errorf("icon %s must contain 1-%d UTF-8 bytes", label, limit)
	}
	for index, character := range value {
		allowed := character >= 'a' && character <= 'z' || character >= '0' && character <= '9' || character == '-' || character == '_' || character == '.' || allowSlash && character == '/'
		if !allowed {
			return "", fmt.Errorf("icon %s %q contains unsupported character %q", label, value, character)
		}
		if (index == 0 || index == len(value)-1) && !(character >= 'a' && character <= 'z' || character >= '0' && character <= '9') {
			return "", fmt.Errorf("icon %s %q must start and end with an ASCII letter or digit", label, value)
		}
	}
	return value, nil
}

func normalizeIconTags(values []string) ([]string, error) {
	if len(values) > maxIconTags {
		return nil, fmt.Errorf("icon has %d tags; maximum is %d", len(values), maxIconTags)
	}
	unique := make(map[string]struct{}, len(values))
	for _, value := range values {
		value = strings.ToLower(strings.Join(strings.FieldsFunc(value, unicode.IsSpace), " "))
		if value == "" {
			continue
		}
		if len(value) > maxIconTagBytes {
			return nil, fmt.Errorf("icon tag %q exceeds %d UTF-8 bytes", value, maxIconTagBytes)
		}
		unique[value] = struct{}{}
	}
	tags := make([]string, 0, len(unique))
	for value := range unique {
		tags = append(tags, value)
	}
	sort.Strings(tags)
	return tags, nil
}

func normalizeIconAliases(values []string, owner entity.IconRef) ([]entity.IconRef, error) {
	if len(values) > maxIconAliases {
		return nil, fmt.Errorf("icon has %d aliases; maximum is %d", len(values), maxIconAliases)
	}
	unique := make(map[string]entity.IconRef, len(values))
	for _, value := range values {
		alias, err := parseIconRef(value, owner.Namespace)
		if err != nil {
			return nil, fmt.Errorf("invalid icon alias %q: %w", value, err)
		}
		if alias == owner {
			continue
		}
		unique[alias.String()] = alias
	}
	aliases := make([]entity.IconRef, 0, len(unique))
	for _, alias := range unique {
		aliases = append(aliases, alias)
	}
	sort.Slice(aliases, func(left, right int) bool {
		return aliases[left].String() < aliases[right].String()
	})
	return aliases, nil
}
