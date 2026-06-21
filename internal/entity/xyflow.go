package entity

type XYFlowDocument struct {
	Nodes          []XYFlowNode   `json:"nodes"`
	Edges          []XYFlowEdge   `json:"edges"`
	XYFlowViewport XYFlowViewport `json:"viewport"`
	Width          float64        `json:"width"`
	Height         float64        `json:"height"`
	Background     string         `json:"background"`
}

type XYFlowViewport struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Zoom float64 `json:"zoom"`
}

type XYFlowPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type XYFlowNode struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	XYFlowPosition XYFlowPosition `json:"position"`
	Width          float64        `json:"width"`
	Height         float64        `json:"height"`
	ParentID       string         `json:"parentId,omitempty"`
	Extent         string         `json:"extent,omitempty"`
	Data           map[string]any `json:"data"`
	Style          map[string]any `json:"style,omitempty"`
}

type XYFlowEdge struct {
	ID           string         `json:"id"`
	Source       string         `json:"source"`
	Target       string         `json:"target"`
	SourceHandle string         `json:"sourceHandle,omitempty"`
	TargetHandle string         `json:"targetHandle,omitempty"`
	Type         string         `json:"type"`
	ZIndex       int            `json:"zIndex,omitempty"`
	Data         map[string]any `json:"data"`
	Style        map[string]any `json:"style,omitempty"`
	MarkerStart  *XYFlowMarker  `json:"markerStart,omitempty"`
	MarkerEnd    *XYFlowMarker  `json:"markerEnd,omitempty"`
}

type XYFlowMarker struct {
	Type  string `json:"type"`
	Color string `json:"color,omitempty"`
}
