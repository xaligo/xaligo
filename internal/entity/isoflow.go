package entity

type IsoflowDocument struct {
	Version     string             `json:"version"`
	Title       string             `json:"title"`
	Description string             `json:"description,omitempty"`
	Items       []IsoflowModelItem `json:"items"`
	Views       []IsoflowView      `json:"views"`
	Icons       []IsoflowIcon      `json:"icons"`
	Colors      []IsoflowColor     `json:"colors"`
	FitToView   bool               `json:"fitToView,omitempty"`
}

type IsoflowModelItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsoflowIcon string `json:"icon,omitempty"`
}

type IsoflowView struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Items       []IsoflowViewItem  `json:"items"`
	Rectangles  []IsoflowRectangle `json:"rectangles,omitempty"`
	Connectors  []IsoflowConnector `json:"connectors,omitempty"`
}

type IsoflowViewItem struct {
	ID          string        `json:"id"`
	Tile        IsoflowCoords `json:"tile"`
	LabelHeight float64       `json:"labelHeight"`
}

type IsoflowRectangle struct {
	ID           string        `json:"id"`
	IsoflowColor string        `json:"color,omitempty"`
	From         IsoflowCoords `json:"from"`
	To           IsoflowCoords `json:"to"`
}

type IsoflowConnector struct {
	ID           string                   `json:"id"`
	IsoflowColor string                   `json:"color,omitempty"`
	Width        float64                  `json:"width,omitempty"`
	Style        string                   `json:"style,omitempty"`
	Anchors      []IsoflowConnectorAnchor `json:"anchors"`
}

type IsoflowConnectorAnchor struct {
	ID  string           `json:"id"`
	Ref IsoflowAnchorRef `json:"ref"`
}

type IsoflowAnchorRef struct {
	Item string `json:"item,omitempty"`
}

type IsoflowIcon struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Collection  string `json:"collection,omitempty"`
	IsIsometric bool   `json:"isIsometric,omitempty"`
}

type IsoflowColor struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type IsoflowCoords struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type IsoflowIconManifest struct {
	Icons map[string]IsoflowIconManifestEntry `json:"icons"`
}

type IsoflowIconManifestEntry struct {
	DataURL string `json:"dataURL"`
}
