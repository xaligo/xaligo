package engine

// Orthogonal arrow routing core (ported from the former TypeScript router).
// Pure, deterministic geometry in layout-pixel space: given anchor rectangles,
// exit/entry sides and obstacle rectangles, it produces a polyline that avoids
// obstacles, prefers empty gutters, minimises crossings and near-parallel
// intrusions, and snaps bends to the grid.

type ptV1EngineRouteTypes struct{ X, Y float64 }

type rectV1EngineRouteTypes struct{ X, Y, W, H float64 }

type sideV1EngineRouteTypes string

const (
	sideTopV1EngineRouteTypes    sideV1EngineRouteTypes = "top"
	sideBottomV1EngineRouteTypes sideV1EngineRouteTypes = "bottom"
	sideLeftV1EngineRouteTypes   sideV1EngineRouteTypes = "left"
	sideRightV1EngineRouteTypes  sideV1EngineRouteTypes = "right"
)

type routeRequestV1EngineRouteTypes struct {
	ID         string
	Kind       string
	Src        rectV1EngineRouteTypes
	Dst        rectV1EngineRouteTypes
	SrcSide    sideV1EngineRouteTypes
	DstSide    sideV1EngineRouteTypes
	SrcAnchor  *ptV1EngineRouteTypes
	DstAnchor  *ptV1EngineRouteTypes
	SrcGap     float64
	DstGap     float64
	SrcLane    float64
	DstLane    float64
	SrcProfile string
	DstProfile string
	Bends      []ptV1EngineRouteTypes
	Grid       float64
	HardAvoid  bool
}

type routedPathV1EngineRouteTypes struct {
	ID     string
	Points []ptV1EngineRouteTypes
}

type segmentV1EngineRouteTypes struct{ A, B ptV1EngineRouteTypes }

type routerOptionsV1EngineRouteTypes struct {
	Grid       float64
	Clearance  float64
	Stub       float64
	LaneGap    float64
	LineMargin float64
	// Reserved paths participate in lane selection and overlap/proximity
	// scoring, but are not solid obstacles. Container borders use this so a
	// connector may cross a frame while avoiding running along its stroke.
	Reserved [][]segmentV1EngineRouteTypes
	// HardObstacles are exclusion zones that the final polyline may never
	// enter. Frame metadata reserve strips use this stricter postcondition.
	HardObstacles []rectV1EngineRouteTypes
	// Bounds is an optional drawing area clamp for local frame routes. Candidate
	// paths outside this rectangle are rejected before scoring.
	Bounds *rectV1EngineRouteTypes
}

func defaultRouterOptionsV1EngineRouteTypes() routerOptionsV1EngineRouteTypes {
	return routerOptionsV1EngineRouteTypes{Grid: 8, Clearance: 12, Stub: 20, LaneGap: 8, LineMargin: 8}
}

const (
	alignTolV1EngineRouteTypes = 4.0
	epsV1EngineRouteTypes      = 1e-6

	wObstacleV1EngineRouteTypes  = 1000.0
	wCrossV1EngineRouteTypes     = 20.0
	wOverlapV1EngineRouteTypes   = 40.0
	wProximityV1EngineRouteTypes = 24.0
	wLenV1EngineRouteTypes       = 5.0
	wBendV1EngineRouteTypes      = 8.0
)
