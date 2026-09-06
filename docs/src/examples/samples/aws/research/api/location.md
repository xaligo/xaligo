# Amazon Location Service

API version: 2020-11-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/location/2020-11-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateTrackerConsumer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `ConsumerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchDeleteDevicePositionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `DeviceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchDeleteDevicePositionHistoryError>` | yes |

## BatchDeleteGeofence

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `GeofenceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchDeleteGeofenceError>` | yes |

## BatchEvaluateGeofences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `DevicePositionUpdates` | `List<DevicePositionUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchEvaluateGeofencesError>` | yes |

## BatchGetDevicePosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `DeviceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchGetDevicePositionError>` | yes |
| `DevicePositions` | `List<DevicePosition>` | yes |

## BatchPutGeofence

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `Entries` | `List<BatchPutGeofenceRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successes` | `List<BatchPutGeofenceSuccess>` | yes |
| `Errors` | `List<BatchPutGeofenceError>` | yes |

## BatchUpdateDevicePosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `Updates` | `List<DevicePositionUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchUpdateDevicePositionError>` | yes |

## CalculateRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |
| `DeparturePosition` | `List<double>` | yes |
| `DestinationPosition` | `List<double>` | yes |
| `WaypointPositions` | `List<List<double>>` | no |
| `TravelMode` | `string` | no |
| `DepartureTime` | `timestamp` | no |
| `DepartNow` | `boolean` | no |
| `DistanceUnit` | `string` | no |
| `IncludeLegGeometry` | `boolean` | no |
| `CarModeOptions` | `CalculateRouteCarModeOptions` | no |
| `TruckModeOptions` | `CalculateRouteTruckModeOptions` | no |
| `ArrivalTime` | `timestamp` | no |
| `OptimizeFor` | `string` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Legs` | `List<Leg>` | yes |
| `Summary` | `CalculateRouteSummary` | yes |

## CalculateRouteMatrix

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |
| `DeparturePositions` | `List<List<double>>` | yes |
| `DestinationPositions` | `List<List<double>>` | yes |
| `TravelMode` | `string` | no |
| `DepartureTime` | `timestamp` | no |
| `DepartNow` | `boolean` | no |
| `DistanceUnit` | `string` | no |
| `CarModeOptions` | `CalculateRouteCarModeOptions` | no |
| `TruckModeOptions` | `CalculateRouteTruckModeOptions` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteMatrix` | `List<List<RouteMatrixEntry>>` | yes |
| `SnappedDeparturePositions` | `List<List<double>>` | no |
| `SnappedDestinationPositions` | `List<List<double>>` | no |
| `Summary` | `CalculateRouteMatrixSummary` | yes |

## CancelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobArn` | `string` | yes |
| `JobId` | `string` | yes |
| `Status` | `string` | yes |

## CreateGeofenceCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `PricingPlan` | `string` | no |
| `PricingPlanDataSource` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |
| `KmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `CollectionArn` | `string` | yes |
| `CreateTime` | `timestamp` | yes |

## CreateKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyName` | `string` | yes |
| `Restrictions` | `ApiKeyRestrictions` | yes |
| `Description` | `string` | no |
| `ExpireTime` | `timestamp` | no |
| `NoExpiry` | `boolean` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `string` | yes |
| `KeyArn` | `string` | yes |
| `KeyName` | `string` | yes |
| `CreateTime` | `timestamp` | yes |

## CreateMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `Configuration` | `MapConfiguration` | yes |
| `PricingPlan` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `MapArn` | `string` | yes |
| `CreateTime` | `timestamp` | yes |

## CreatePlaceIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `DataSource` | `string` | yes |
| `PricingPlan` | `string` | no |
| `Description` | `string` | no |
| `DataSourceConfiguration` | `DataSourceConfiguration` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `IndexArn` | `string` | yes |
| `CreateTime` | `timestamp` | yes |

## CreateRouteCalculator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |
| `DataSource` | `string` | yes |
| `PricingPlan` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |
| `CalculatorArn` | `string` | yes |
| `CreateTime` | `timestamp` | yes |

## CreateTracker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `PricingPlan` | `string` | no |
| `KmsKeyId` | `string` | no |
| `PricingPlanDataSource` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |
| `PositionFiltering` | `string` | no |
| `EventBridgeEnabled` | `boolean` | no |
| `KmsKeyEnableGeospatialQueries` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `TrackerArn` | `string` | yes |
| `CreateTime` | `timestamp` | yes |

## DeleteGeofenceCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyName` | `string` | yes |
| `ForceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePlaceIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRouteCalculator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTracker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeGeofenceCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `CollectionArn` | `string` | yes |
| `Description` | `string` | yes |
| `PricingPlan` | `string` | no |
| `PricingPlanDataSource` | `string` | no |
| `KmsKeyId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `CreateTime` | `timestamp` | yes |
| `UpdateTime` | `timestamp` | yes |
| `GeofenceCount` | `integer` | no |

## DescribeKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `string` | yes |
| `KeyArn` | `string` | yes |
| `KeyName` | `string` | yes |
| `Restrictions` | `ApiKeyRestrictions` | yes |
| `CreateTime` | `timestamp` | yes |
| `ExpireTime` | `timestamp` | yes |
| `UpdateTime` | `timestamp` | yes |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribeMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `MapArn` | `string` | yes |
| `PricingPlan` | `string` | no |
| `DataSource` | `string` | yes |
| `Configuration` | `MapConfiguration` | yes |
| `Description` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `CreateTime` | `timestamp` | yes |
| `UpdateTime` | `timestamp` | yes |

## DescribePlaceIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `IndexArn` | `string` | yes |
| `PricingPlan` | `string` | no |
| `Description` | `string` | yes |
| `CreateTime` | `timestamp` | yes |
| `UpdateTime` | `timestamp` | yes |
| `DataSource` | `string` | yes |
| `DataSourceConfiguration` | `DataSourceConfiguration` | yes |
| `Tags` | `Map<string>` | no |

## DescribeRouteCalculator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |
| `CalculatorArn` | `string` | yes |
| `PricingPlan` | `string` | no |
| `Description` | `string` | yes |
| `CreateTime` | `timestamp` | yes |
| `UpdateTime` | `timestamp` | yes |
| `DataSource` | `string` | yes |
| `Tags` | `Map<string>` | no |

## DescribeTracker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `TrackerArn` | `string` | yes |
| `Description` | `string` | yes |
| `PricingPlan` | `string` | no |
| `PricingPlanDataSource` | `string` | no |
| `Tags` | `Map<string>` | no |
| `CreateTime` | `timestamp` | yes |
| `UpdateTime` | `timestamp` | yes |
| `KmsKeyId` | `string` | no |
| `PositionFiltering` | `string` | no |
| `EventBridgeEnabled` | `boolean` | no |
| `KmsKeyEnableGeospatialQueries` | `boolean` | no |

## DisassociateTrackerConsumer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `ConsumerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ForecastGeofenceEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `DeviceState` | `ForecastGeofenceEventsDeviceState` | yes |
| `TimeHorizonMinutes` | `double` | no |
| `DistanceUnit` | `string` | no |
| `SpeedUnit` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastedEvents` | `List<ForecastedEvent>` | yes |
| `NextToken` | `string` | no |
| `DistanceUnit` | `string` | yes |
| `SpeedUnit` | `string` | yes |

## GetDevicePosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `DeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceId` | `string` | no |
| `SampleTime` | `timestamp` | yes |
| `ReceivedTime` | `timestamp` | yes |
| `Position` | `List<double>` | yes |
| `Accuracy` | `PositionalAccuracy` | no |
| `PositionProperties` | `Map<string>` | no |

## GetDevicePositionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `DeviceId` | `string` | yes |
| `NextToken` | `string` | no |
| `StartTimeInclusive` | `timestamp` | no |
| `EndTimeExclusive` | `timestamp` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DevicePositions` | `List<DevicePosition>` | yes |
| `NextToken` | `string` | no |

## GetGeofence

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `GeofenceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeofenceId` | `string` | yes |
| `Geometry` | `GeofenceGeometry` | yes |
| `Status` | `string` | yes |
| `CreateTime` | `timestamp` | yes |
| `UpdateTime` | `timestamp` | yes |
| `GeofenceProperties` | `Map<string>` | no |

## GetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `string` | yes |
| `ActionOptions` | `JobActionOptions` | no |
| `CreatedAt` | `timestamp` | yes |
| `EndedAt` | `timestamp` | no |
| `Error` | `JobError` | no |
| `ExecutionRoleArn` | `string` | yes |
| `InputOptions` | `JobInputOptions` | yes |
| `JobArn` | `string` | yes |
| `JobId` | `string` | yes |
| `Name` | `string` | no |
| `OutputOptions` | `JobOutputOptions` | yes |
| `Status` | `string` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |

## GetMapGlyphs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `FontStack` | `string` | yes |
| `FontUnicodeRange` | `string` | yes |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |

## GetMapSprites

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `FileName` | `string` | yes |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |

## GetMapStyleDescriptor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |

## GetMapTile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `Z` | `string` | yes |
| `X` | `string` | yes |
| `Y` | `string` | yes |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |

## GetPlace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `PlaceId` | `string` | yes |
| `Language` | `string` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Place` | `Place` | yes |

## ListDevicePositions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `FilterGeometry` | `TrackingFilterGeometry` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListDevicePositionsResponseEntry>` | yes |
| `NextToken` | `string` | no |

## ListGeofenceCollections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListGeofenceCollectionsResponseEntry>` | yes |
| `NextToken` | `string` | no |

## ListGeofences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListGeofenceResponseEntry>` | yes |
| `NextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `JobsFilter` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListJobsResponseEntry>` | yes |
| `NextToken` | `string` | no |

## ListKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filter` | `ApiKeyFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListKeysResponseEntry>` | yes |
| `NextToken` | `string` | no |

## ListMaps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListMapsResponseEntry>` | yes |
| `NextToken` | `string` | no |

## ListPlaceIndexes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListPlaceIndexesResponseEntry>` | yes |
| `NextToken` | `string` | no |

## ListRouteCalculators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListRouteCalculatorsResponseEntry>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListTrackerConsumers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsumerArns` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListTrackers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<ListTrackersResponseEntry>` | yes |
| `NextToken` | `string` | no |

## PutGeofence

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `GeofenceId` | `string` | yes |
| `Geometry` | `GeofenceGeometry` | yes |
| `GeofenceProperties` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeofenceId` | `string` | yes |
| `CreateTime` | `timestamp` | yes |
| `UpdateTime` | `timestamp` | yes |

## SearchPlaceIndexForPosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `Position` | `List<double>` | yes |
| `MaxResults` | `integer` | no |
| `Language` | `string` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summary` | `SearchPlaceIndexForPositionSummary` | yes |
| `Results` | `List<SearchForPositionResult>` | yes |

## SearchPlaceIndexForSuggestions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `Text` | `string` | yes |
| `BiasPosition` | `List<double>` | no |
| `FilterBBox` | `List<double>` | no |
| `FilterCountries` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `Language` | `string` | no |
| `FilterCategories` | `List<string>` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summary` | `SearchPlaceIndexForSuggestionsSummary` | yes |
| `Results` | `List<SearchForSuggestionsResult>` | yes |

## SearchPlaceIndexForText

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `Text` | `string` | yes |
| `BiasPosition` | `List<double>` | no |
| `FilterBBox` | `List<double>` | no |
| `FilterCountries` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `Language` | `string` | no |
| `FilterCategories` | `List<string>` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summary` | `SearchPlaceIndexForTextSummary` | yes |
| `Results` | `List<SearchForTextResult>` | yes |

## StartJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Action` | `string` | yes |
| `ActionOptions` | `JobActionOptions` | no |
| `ExecutionRoleArn` | `string` | yes |
| `InputOptions` | `JobInputOptions` | yes |
| `Name` | `string` | no |
| `OutputOptions` | `JobOutputOptions` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAt` | `timestamp` | yes |
| `JobArn` | `string` | yes |
| `JobId` | `string` | yes |
| `Status` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGeofenceCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `PricingPlan` | `string` | no |
| `PricingPlanDataSource` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionName` | `string` | yes |
| `CollectionArn` | `string` | yes |
| `UpdateTime` | `timestamp` | yes |

## UpdateKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyName` | `string` | yes |
| `Description` | `string` | no |
| `ExpireTime` | `timestamp` | no |
| `NoExpiry` | `boolean` | no |
| `ForceUpdate` | `boolean` | no |
| `Restrictions` | `ApiKeyRestrictions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyName` | `string` | yes |
| `UpdateTime` | `timestamp` | yes |

## UpdateMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `PricingPlan` | `string` | no |
| `Description` | `string` | no |
| `ConfigurationUpdate` | `MapConfigurationUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MapName` | `string` | yes |
| `MapArn` | `string` | yes |
| `UpdateTime` | `timestamp` | yes |

## UpdatePlaceIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `PricingPlan` | `string` | no |
| `Description` | `string` | no |
| `DataSourceConfiguration` | `DataSourceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexName` | `string` | yes |
| `IndexArn` | `string` | yes |
| `UpdateTime` | `timestamp` | yes |

## UpdateRouteCalculator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |
| `PricingPlan` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatorName` | `string` | yes |
| `CalculatorArn` | `string` | yes |
| `UpdateTime` | `timestamp` | yes |

## UpdateTracker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `PricingPlan` | `string` | no |
| `PricingPlanDataSource` | `string` | no |
| `Description` | `string` | no |
| `PositionFiltering` | `string` | no |
| `EventBridgeEnabled` | `boolean` | no |
| `KmsKeyEnableGeospatialQueries` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `TrackerArn` | `string` | yes |
| `UpdateTime` | `timestamp` | yes |

## VerifyDevicePosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackerName` | `string` | yes |
| `DeviceState` | `DeviceState` | yes |
| `DistanceUnit` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferredState` | `InferredState` | yes |
| `DeviceId` | `string` | yes |
| `SampleTime` | `timestamp` | yes |
| `ReceivedTime` | `timestamp` | yes |
| `DistanceUnit` | `string` | yes |

