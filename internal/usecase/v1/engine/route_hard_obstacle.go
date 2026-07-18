package engine

import (
	"math"
	"sort"
)

// visibilityGridFallback builds a deterministic orthogonal path through the
// free coordinate grid around every hard obstacle. It is the final drawable
// fallback after the normal scored router and the metadata-strip fast path.
func visibilityGridFallbackV1EngineRouteHardObstacle(req routeRequestV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) ([]ptV1EngineRouteTypes, bool) {
	source, sourceStub, destination, destinationStub := routeEndpointStubsV1EngineRouteHardObstacle(req, opt)
	if pointBlockedV1EngineRouteHardObstacle(source, obstacles) || pointBlockedV1EngineRouteHardObstacle(destination, obstacles) {
		return nil, false
	}

	useSourceStub := !pointBlockedV1EngineRouteHardObstacle(sourceStub, obstacles) && !segmentBlockedV1EngineRouteHardObstacle(source, sourceStub, obstacles)
	useDestinationStub := !pointBlockedV1EngineRouteHardObstacle(destinationStub, obstacles) && !segmentBlockedV1EngineRouteHardObstacle(destinationStub, destination, obstacles)
	start := source
	if useSourceStub {
		start = sourceStub
	}
	end := destination
	if useDestinationStub {
		end = destinationStub
	}

	const clearance = 1.0
	xCoordinates := []float64{start.X, end.X}
	yCoordinates := []float64{start.Y, end.Y}
	for _, obstacle := range obstacles {
		xCoordinates = append(xCoordinates, obstacle.X-clearance, obstacle.X+obstacle.W+clearance)
		yCoordinates = append(yCoordinates, obstacle.Y-clearance, obstacle.Y+obstacle.H+clearance)
	}
	xCoordinates = sortedUniqueCoordinatesV1EngineRouteHardObstacle(xCoordinates)
	yCoordinates = sortedUniqueCoordinatesV1EngineRouteHardObstacle(yCoordinates)

	type gridNode struct {
		point ptV1EngineRouteTypes
		x     int
		y     int
	}
	nodeIDs := make([][]int, len(yCoordinates))
	nodes := make([]gridNode, 0, len(xCoordinates)*len(yCoordinates))
	for yIndex, y := range yCoordinates {
		nodeIDs[yIndex] = make([]int, len(xCoordinates))
		for xIndex := range xCoordinates {
			nodeIDs[yIndex][xIndex] = -1
		}
		for xIndex, x := range xCoordinates {
			point := ptV1EngineRouteTypes{X: x, Y: y}
			if pointBlockedV1EngineRouteHardObstacle(point, obstacles) {
				continue
			}
			nodeIDs[yIndex][xIndex] = len(nodes)
			nodes = append(nodes, gridNode{point: point, x: xIndex, y: yIndex})
		}
	}
	if len(nodes) == 0 {
		return nil, false
	}

	startX := coordinateIndexV1EngineRouteHardObstacle(xCoordinates, start.X)
	startY := coordinateIndexV1EngineRouteHardObstacle(yCoordinates, start.Y)
	endX := coordinateIndexV1EngineRouteHardObstacle(xCoordinates, end.X)
	endY := coordinateIndexV1EngineRouteHardObstacle(yCoordinates, end.Y)
	if startX < 0 || startY < 0 || endX < 0 || endY < 0 {
		return nil, false
	}
	startID := nodeIDs[startY][startX]
	endID := nodeIDs[endY][endX]
	if startID < 0 || endID < 0 {
		return nil, false
	}

	adjacent := make([][]int, len(nodes))
	connect := func(left, right int) {
		if left < 0 || right < 0 || segmentBlockedV1EngineRouteHardObstacle(nodes[left].point, nodes[right].point, obstacles) {
			return
		}
		adjacent[left] = append(adjacent[left], right)
		adjacent[right] = append(adjacent[right], left)
	}
	for yIndex := range yCoordinates {
		previous := -1
		for xIndex := range xCoordinates {
			current := nodeIDs[yIndex][xIndex]
			if current < 0 {
				continue
			}
			connect(previous, current)
			previous = current
		}
	}
	for xIndex := range xCoordinates {
		previous := -1
		for yIndex := range yCoordinates {
			current := nodeIDs[yIndex][xIndex]
			if current < 0 {
				continue
			}
			connect(previous, current)
			previous = current
		}
	}

	previous := make([]int, len(nodes))
	for index := range previous {
		previous[index] = -1
	}
	previous[startID] = startID
	queue := []int{startID}
	for len(queue) > 0 && previous[endID] < 0 {
		current := queue[0]
		queue = queue[1:]
		for _, next := range adjacent[current] {
			if previous[next] >= 0 {
				continue
			}
			previous[next] = current
			queue = append(queue, next)
		}
	}
	if previous[endID] < 0 {
		return nil, false
	}

	reversed := []ptV1EngineRouteTypes{}
	for current := endID; ; current = previous[current] {
		reversed = append(reversed, nodes[current].point)
		if current == startID {
			break
		}
	}
	gridPath := make([]ptV1EngineRouteTypes, len(reversed))
	for index := range reversed {
		gridPath[index] = reversed[len(reversed)-1-index]
	}
	points := make([]ptV1EngineRouteTypes, 0, len(gridPath)+2)
	if useSourceStub {
		points = append(points, source)
	}
	points = append(points, gridPath...)
	if useDestinationStub {
		points = append(points, destination)
	}
	points = enforceOrthogonalPolylineV1EngineRoutePath(points)
	if len(points) < 2 || obstacleHitCountV1EngineRouteCandidate(points, obstacles) > 0 {
		return nil, false
	}
	return points, true
}

func routeEndpointStubsV1EngineRouteHardObstacle(req routeRequestV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) (ptV1EngineRouteTypes, ptV1EngineRouteTypes, ptV1EngineRouteTypes, ptV1EngineRouteTypes) {
	source := edgeMidpointV1EngineRouteGeometry(req.Src, req.SrcSide)
	if req.SrcAnchor != nil {
		source = *req.SrcAnchor
	}
	destination := edgeMidpointV1EngineRouteGeometry(req.Dst, req.DstSide)
	if req.DstAnchor != nil {
		destination = *req.DstAnchor
	}
	if req.SrcGap > 0 {
		source = extendV1EngineRouteGeometry(source, req.SrcSide, req.SrcGap)
	}
	if req.DstGap > 0 {
		destination = extendV1EngineRouteGeometry(destination, req.DstSide, req.DstGap)
	}
	sourceStub := extendV1EngineRouteGeometry(source, req.SrcSide, math.Max(opt.LaneGap, opt.Stub+req.SrcLane*opt.LaneGap))
	destinationStub := extendV1EngineRouteGeometry(destination, req.DstSide, math.Max(opt.LaneGap, opt.Stub+req.DstLane*opt.LaneGap))
	return source, sourceStub, destination, destinationStub
}

func sortedUniqueCoordinatesV1EngineRouteHardObstacle(values []float64) []float64 {
	sort.Float64s(values)
	out := make([]float64, 0, len(values))
	for _, value := range values {
		if len(out) == 0 || math.Abs(value-out[len(out)-1]) > epsV1EngineRouteTypes {
			out = append(out, value)
		}
	}
	return out
}

func coordinateIndexV1EngineRouteHardObstacle(values []float64, wanted float64) int {
	for index, value := range values {
		if math.Abs(value-wanted) <= epsV1EngineRouteTypes {
			return index
		}
	}
	return -1
}

func pointBlockedV1EngineRouteHardObstacle(point ptV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes) bool {
	for _, obstacle := range obstacles {
		if point.X >= obstacle.X && point.X <= obstacle.X+obstacle.W && point.Y >= obstacle.Y && point.Y <= obstacle.Y+obstacle.H {
			return true
		}
	}
	return false
}

func segmentBlockedV1EngineRouteHardObstacle(start, end ptV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes) bool {
	return obstacleHitCountV1EngineRouteCandidate([]ptV1EngineRouteTypes{start, end}, obstacles) > 0
}
