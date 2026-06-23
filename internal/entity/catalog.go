package entity

// CatalogEntry is a resolved row from the service icon catalog.
type CatalogEntry struct {
	ID       int
	Category string
	Service  string
	SVGFile  string
	RelPath  string
	DataURL  string
}
