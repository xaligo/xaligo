# Amazon Location Service Routes V2

API version: 2020-11-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/geo-routes/2020-11-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CalculateIsolines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Allow` | `IsolineAllowOptions` | no |
| `ArrivalTime` | `string` | no |
| `Avoid` | `IsolineAvoidanceOptions` | no |
| `DepartNow` | `boolean` | no |
| `DepartureTime` | `string` | no |
| `Destination` | `List<double>` | no |
| `DestinationOptions` | `IsolineDestinationOptions` | no |
| `IsolineGeometryFormat` | `string` | no |
| `IsolineGranularity` | `IsolineGranularityOptions` | no |
| `Key` | `string` | no |
| `OptimizeIsolineFor` | `string` | no |
| `OptimizeRoutingFor` | `string` | no |
| `Origin` | `List<double>` | no |
| `OriginOptions` | `IsolineOriginOptions` | no |
| `Thresholds` | `IsolineThresholds` | yes |
| `Traffic` | `IsolineTrafficOptions` | no |
| `TravelMode` | `string` | no |
| `TravelModeOptions` | `IsolineTravelModeOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArrivalTime` | `string` | no |
| `DepartureTime` | `string` | no |
| `IsolineGeometryFormat` | `string` | yes |
| `Isolines` | `List<Isoline>` | yes |
| `PricingBucket` | `string` | yes |
| `SnappedDestination` | `List<double>` | no |
| `SnappedOrigin` | `List<double>` | no |

## CalculateRouteMatrix

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Allow` | `RouteMatrixAllowOptions` | no |
| `Avoid` | `RouteMatrixAvoidanceOptions` | no |
| `DepartNow` | `boolean` | no |
| `DepartureTime` | `string` | no |
| `Destinations` | `List<RouteMatrixDestination>` | yes |
| `Exclude` | `RouteMatrixExclusionOptions` | no |
| `Key` | `string` | no |
| `OptimizeRoutingFor` | `string` | no |
| `Origins` | `List<RouteMatrixOrigin>` | yes |
| `RoutingBoundary` | `RouteMatrixBoundary` | no |
| `Traffic` | `RouteMatrixTrafficOptions` | no |
| `TravelMode` | `string` | no |
| `TravelModeOptions` | `RouteMatrixTravelModeOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ErrorCount` | `integer` | yes |
| `PricingBucket` | `string` | yes |
| `RouteMatrix` | `List<List<RouteMatrixEntry>>` | yes |
| `RoutingBoundary` | `RouteMatrixBoundary` | yes |

## CalculateRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Allow` | `RouteAllowOptions` | no |
| `ArrivalTime` | `string` | no |
| `Avoid` | `RouteAvoidanceOptions` | no |
| `DepartNow` | `boolean` | no |
| `DepartureTime` | `string` | no |
| `Destination` | `List<double>` | yes |
| `DestinationOptions` | `RouteDestinationOptions` | no |
| `Driver` | `RouteDriverOptions` | no |
| `Exclude` | `RouteExclusionOptions` | no |
| `InstructionsMeasurementSystem` | `string` | no |
| `Key` | `string` | no |
| `Languages` | `List<string>` | no |
| `LegAdditionalFeatures` | `List<string>` | no |
| `LegGeometryFormat` | `string` | no |
| `MaxAlternatives` | `integer` | no |
| `OptimizeRoutingFor` | `string` | no |
| `Origin` | `List<double>` | yes |
| `OriginOptions` | `RouteOriginOptions` | no |
| `SpanAdditionalFeatures` | `List<string>` | no |
| `Tolls` | `RouteTollOptions` | no |
| `Traffic` | `RouteTrafficOptions` | no |
| `TravelMode` | `string` | no |
| `TravelModeOptions` | `RouteTravelModeOptions` | no |
| `TravelStepType` | `string` | no |
| `Waypoints` | `List<RouteWaypoint>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LegGeometryFormat` | `string` | yes |
| `Notices` | `List<RouteResponseNotice>` | yes |
| `PricingBucket` | `string` | yes |
| `Routes` | `List<Route>` | yes |

## OptimizeWaypoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Avoid` | `WaypointOptimizationAvoidanceOptions` | no |
| `Clustering` | `WaypointOptimizationClusteringOptions` | no |
| `DepartureTime` | `string` | no |
| `Destination` | `List<double>` | no |
| `DestinationOptions` | `WaypointOptimizationDestinationOptions` | no |
| `Driver` | `WaypointOptimizationDriverOptions` | no |
| `Exclude` | `WaypointOptimizationExclusionOptions` | no |
| `Key` | `string` | no |
| `OptimizeSequencingFor` | `string` | no |
| `Origin` | `List<double>` | yes |
| `OriginOptions` | `WaypointOptimizationOriginOptions` | no |
| `Traffic` | `WaypointOptimizationTrafficOptions` | no |
| `TravelMode` | `string` | no |
| `TravelModeOptions` | `WaypointOptimizationTravelModeOptions` | no |
| `Waypoints` | `List<WaypointOptimizationWaypoint>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connections` | `List<WaypointOptimizationConnection>` | yes |
| `Distance` | `long` | yes |
| `Duration` | `long` | yes |
| `ImpedingWaypoints` | `List<WaypointOptimizationImpedingWaypoint>` | yes |
| `OptimizedWaypoints` | `List<WaypointOptimizationOptimizedWaypoint>` | yes |
| `PricingBucket` | `string` | yes |
| `TimeBreakdown` | `WaypointOptimizationTimeBreakdown` | yes |

## SnapToRoads

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `string` | no |
| `SnappedGeometryFormat` | `string` | no |
| `SnapRadius` | `long` | no |
| `TracePoints` | `List<RoadSnapTracePoint>` | yes |
| `TravelMode` | `string` | no |
| `TravelModeOptions` | `RoadSnapTravelModeOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Notices` | `List<RoadSnapNotice>` | yes |
| `PricingBucket` | `string` | yes |
| `SnappedGeometry` | `RoadSnapSnappedGeometry` | no |
| `SnappedGeometryFormat` | `string` | yes |
| `SnappedTracePoints` | `List<RoadSnapSnappedTracePoint>` | yes |

