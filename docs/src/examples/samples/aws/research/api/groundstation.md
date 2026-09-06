# AWS Ground Station

API version: 2019-05-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/groundstation/2019-05-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | no |
| `versionId` | `integer` | no |

## CreateConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `configData` | `ConfigTypeData` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | no |
| `configType` | `string` | no |
| `configArn` | `string` | no |

## CreateDataflowEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointDetails` | `List<EndpointDetails>` | yes |
| `tags` | `Map<string>` | no |
| `contactPrePassDurationSeconds` | `integer` | no |
| `contactPostPassDurationSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataflowEndpointGroupId` | `string` | no |

## CreateDataflowEndpointGroupV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoints` | `List<CreateEndpointDetails>` | yes |
| `contactPrePassDurationSeconds` | `integer` | no |
| `contactPostPassDurationSeconds` | `integer` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataflowEndpointGroupId` | `string` | no |

## CreateEphemeris

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `satelliteId` | `string` | no |
| `enabled` | `boolean` | no |
| `priority` | `integer` | no |
| `expirationTime` | `timestamp` | no |
| `name` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `ephemeris` | `EphemerisData` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ephemerisId` | `string` | no |

## CreateMissionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `contactPrePassDurationSeconds` | `integer` | no |
| `contactPostPassDurationSeconds` | `integer` | no |
| `minimumViableContactDurationSeconds` | `integer` | yes |
| `dataflowEdges` | `List<List<string>>` | yes |
| `trackingConfigArn` | `string` | yes |
| `telemetrySinkConfigArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `streamsKmsKey` | `KmsKey` | no |
| `streamsKmsRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `missionProfileId` | `string` | no |

## DeleteConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | yes |
| `configType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | no |
| `configType` | `string` | no |
| `configArn` | `string` | no |

## DeleteDataflowEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataflowEndpointGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataflowEndpointGroupId` | `string` | no |

## DeleteEphemeris

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ephemerisId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ephemerisId` | `string` | no |

## DeleteMissionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `missionProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `missionProfileId` | `string` | no |

## DescribeContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | no |
| `missionProfileArn` | `string` | no |
| `satelliteArn` | `string` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `prePassStartTime` | `timestamp` | no |
| `postPassEndTime` | `timestamp` | no |
| `groundStation` | `string` | no |
| `contactStatus` | `string` | no |
| `errorMessage` | `string` | no |
| `maximumElevation` | `Elevation` | no |
| `tags` | `Map<string>` | no |
| `region` | `string` | no |
| `dataflowList` | `List<DataflowDetail>` | no |
| `visibilityStartTime` | `timestamp` | no |
| `visibilityEndTime` | `timestamp` | no |
| `trackingOverrides` | `TrackingOverrides` | no |
| `ephemeris` | `EphemerisResponseData` | no |
| `version` | `ContactVersion` | no |

## DescribeContactVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | yes |
| `versionId` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | no |
| `missionProfileArn` | `string` | no |
| `satelliteArn` | `string` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `prePassStartTime` | `timestamp` | no |
| `postPassEndTime` | `timestamp` | no |
| `groundStation` | `string` | no |
| `contactStatus` | `string` | no |
| `errorMessage` | `string` | no |
| `maximumElevation` | `Elevation` | no |
| `tags` | `Map<string>` | no |
| `region` | `string` | no |
| `dataflowList` | `List<DataflowDetail>` | no |
| `visibilityStartTime` | `timestamp` | no |
| `visibilityEndTime` | `timestamp` | no |
| `trackingOverrides` | `TrackingOverrides` | no |
| `ephemeris` | `EphemerisResponseData` | no |
| `version` | `ContactVersion` | no |

## DescribeEphemeris

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ephemerisId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ephemerisId` | `string` | no |
| `satelliteId` | `string` | no |
| `status` | `string` | no |
| `priority` | `integer` | no |
| `creationTime` | `timestamp` | no |
| `enabled` | `boolean` | no |
| `name` | `string` | no |
| `tags` | `Map<string>` | no |
| `suppliedData` | `EphemerisTypeDescription` | no |
| `invalidReason` | `string` | no |
| `errorReasons` | `List<EphemerisErrorReason>` | no |

## GetAgentConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | no |
| `taskingDocument` | `string` | no |

## GetAgentTaskResponseUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `taskId` | `string` | yes |
| `presignedLogUrl` | `string` | yes |

## GetConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | yes |
| `configType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | yes |
| `configArn` | `string` | yes |
| `name` | `string` | yes |
| `configType` | `string` | no |
| `configData` | `ConfigTypeData` | yes |
| `tags` | `Map<string>` | no |

## GetDataflowEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataflowEndpointGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataflowEndpointGroupId` | `string` | no |
| `dataflowEndpointGroupArn` | `string` | no |
| `endpointsDetails` | `List<EndpointDetails>` | no |
| `tags` | `Map<string>` | no |
| `contactPrePassDurationSeconds` | `integer` | no |
| `contactPostPassDurationSeconds` | `integer` | no |

## GetMinuteUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `month` | `integer` | yes |
| `year` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `isReservedMinutesCustomer` | `boolean` | no |
| `totalReservedMinuteAllocation` | `integer` | no |
| `upcomingMinutesScheduled` | `integer` | no |
| `totalScheduledMinutes` | `integer` | no |
| `estimatedMinutesRemaining` | `integer` | no |

## GetMissionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `missionProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `missionProfileId` | `string` | no |
| `missionProfileArn` | `string` | no |
| `name` | `string` | no |
| `region` | `string` | no |
| `contactPrePassDurationSeconds` | `integer` | no |
| `contactPostPassDurationSeconds` | `integer` | no |
| `minimumViableContactDurationSeconds` | `integer` | no |
| `dataflowEdges` | `List<List<string>>` | no |
| `trackingConfigArn` | `string` | no |
| `telemetrySinkConfigArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `streamsKmsKey` | `KmsKey` | no |
| `streamsKmsRole` | `string` | no |

## GetSatellite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `satelliteId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `satelliteId` | `string` | no |
| `satelliteArn` | `string` | no |
| `noradSatelliteID` | `integer` | no |
| `groundStations` | `List<string>` | no |
| `currentEphemeris` | `EphemerisMetaData` | no |

## ListAntennas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `groundStationId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `antennaList` | `List<AntennaListItem>` | yes |
| `nextToken` | `string` | no |

## ListConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `configList` | `List<ConfigListItem>` | no |

## ListContactVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `contactVersionsList` | `List<ContactVersion>` | no |

## ListContacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `statusList` | `List<string>` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `groundStation` | `string` | no |
| `satelliteArn` | `string` | no |
| `missionProfileArn` | `string` | no |
| `ephemeris` | `EphemerisFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `contactList` | `List<ContactData>` | no |

## ListDataflowEndpointGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `dataflowEndpointGroupList` | `List<DataflowEndpointListItem>` | no |

## ListEphemerides

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `satelliteId` | `string` | no |
| `ephemerisType` | `string` | no |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `statusList` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `ephemerides` | `List<EphemerisItem>` | no |

## ListGroundStationReservations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `groundStationId` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `reservationTypes` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reservationList` | `List<GroundStationReservationListItem>` | yes |
| `nextToken` | `string` | no |

## ListGroundStations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `satelliteId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `groundStationList` | `List<GroundStationData>` | no |

## ListMissionProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `missionProfileList` | `List<MissionProfileListItem>` | no |

## ListSatellites

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `satellites` | `List<SatelliteListItem>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## RegisterAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `discoveryData` | `DiscoveryData` | yes |
| `agentDetails` | `AgentDetails` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | no |

## ReserveContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `missionProfileArn` | `string` | yes |
| `satelliteArn` | `string` | no |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `groundStation` | `string` | yes |
| `tags` | `Map<string>` | no |
| `trackingOverrides` | `TrackingOverrides` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | no |
| `versionId` | `integer` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAgentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `taskId` | `string` | yes |
| `aggregateStatus` | `AggregateStatus` | yes |
| `componentStatuses` | `List<ComponentStatusData>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |

## UpdateConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | yes |
| `name` | `string` | yes |
| `configType` | `string` | yes |
| `configData` | `ConfigTypeData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | no |
| `configType` | `string` | no |
| `configArn` | `string` | no |

## UpdateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | yes |
| `clientToken` | `string` | no |
| `trackingOverrides` | `TrackingOverrides` | no |
| `satelliteArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactId` | `string` | no |
| `versionId` | `integer` | no |

## UpdateEphemeris

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ephemerisId` | `string` | yes |
| `enabled` | `boolean` | yes |
| `name` | `string` | no |
| `priority` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ephemerisId` | `string` | no |

## UpdateMissionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `missionProfileId` | `string` | yes |
| `name` | `string` | no |
| `contactPrePassDurationSeconds` | `integer` | no |
| `contactPostPassDurationSeconds` | `integer` | no |
| `minimumViableContactDurationSeconds` | `integer` | no |
| `dataflowEdges` | `List<List<string>>` | no |
| `trackingConfigArn` | `string` | no |
| `telemetrySinkConfigArn` | `string` | no |
| `streamsKmsKey` | `KmsKey` | no |
| `streamsKmsRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `missionProfileId` | `string` | no |

