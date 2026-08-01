package entity

import "time"

// IconRef is the stable namespaced identity of an SVG icon.
type IconRef struct {
	Namespace string
	Name      string
}

func (rcvr IconRef) String() string {
	return rcvr.Namespace + ":" + rcvr.Name
}

// Icon is one canonical SVG registry record.
type Icon struct {
	ID          int64
	Ref         IconRef
	Description string
	SVG         []byte
	ViewBox     string
	Width       *float64
	Height      *float64
	Checksum    [32]byte
	Compression int
	License     string
	Source      string
	Tags        []string
	Aliases     []IconRef
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// IconSummary is the searchable metadata projection without SVG bytes.
type IconSummary struct {
	Ref         IconRef
	Description string
	ViewBox     string
	Width       *float64
	Height      *float64
	License     string
	Source      string
}

// IconRegistration is the controller-to-use-case input for an add or update.
type IconRegistration struct {
	Reference   string
	SVG         []byte
	Description string
	Tags        []string
	Aliases     []string
	License     string
	Source      string
}
