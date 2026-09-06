# AWS IoT TwinMaker

API version: 2021-11-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iottwinmaker/2021-11-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchPutPropertyValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `entries` | `List<PropertyValueEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorEntries` | `List<BatchPutPropertyErrorEntry>` | yes |

## CancelMetadataTransferJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metadataTransferJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metadataTransferJobId` | `string` | yes |
| `arn` | `string` | yes |
| `updateDateTime` | `timestamp` | yes |
| `status` | `MetadataTransferJobStatus` | yes |
| `progress` | `MetadataTransferJobProgress` | no |

## CreateComponentType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `isSingleton` | `boolean` | no |
| `componentTypeId` | `string` | yes |
| `description` | `string` | no |
| `propertyDefinitions` | `Map<PropertyDefinitionRequest>` | no |
| `extendsFrom` | `List<string>` | no |
| `functions` | `Map<FunctionRequest>` | no |
| `tags` | `Map<string>` | no |
| `propertyGroups` | `Map<PropertyGroupRequest>` | no |
| `componentTypeName` | `string` | no |
| `compositeComponentTypes` | `Map<CompositeComponentTypeRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `creationDateTime` | `timestamp` | yes |
| `state` | `string` | yes |

## CreateEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `entityId` | `string` | no |
| `entityName` | `string` | yes |
| `description` | `string` | no |
| `components` | `Map<ComponentRequest>` | no |
| `compositeComponents` | `Map<CompositeComponentRequest>` | no |
| `parentEntityId` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entityId` | `string` | yes |
| `arn` | `string` | yes |
| `creationDateTime` | `timestamp` | yes |
| `state` | `string` | yes |

## CreateMetadataTransferJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metadataTransferJobId` | `string` | no |
| `description` | `string` | no |
| `sources` | `List<SourceConfiguration>` | yes |
| `destination` | `DestinationConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metadataTransferJobId` | `string` | yes |
| `arn` | `string` | yes |
| `creationDateTime` | `timestamp` | yes |
| `status` | `MetadataTransferJobStatus` | yes |

## CreateScene

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `sceneId` | `string` | yes |
| `contentLocation` | `string` | yes |
| `description` | `string` | no |
| `capabilities` | `List<string>` | no |
| `tags` | `Map<string>` | no |
| `sceneMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `creationDateTime` | `timestamp` | yes |

## CreateSyncJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `syncSource` | `string` | yes |
| `syncRole` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `creationDateTime` | `timestamp` | yes |
| `state` | `string` | yes |

## CreateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `description` | `string` | no |
| `s3Location` | `string` | no |
| `role` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `creationDateTime` | `timestamp` | yes |

## DeleteComponentType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `componentTypeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `state` | `string` | yes |

## DeleteEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `entityId` | `string` | yes |
| `isRecursive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `state` | `string` | yes |

## DeleteScene

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `sceneId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSyncJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `syncSource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `state` | `string` | yes |

## DeleteWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## ExecuteQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `queryStatement` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `columnDescriptions` | `List<ColumnDescription>` | no |
| `rows` | `List<Row>` | no |
| `nextToken` | `string` | no |

## GetComponentType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `componentTypeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `isSingleton` | `boolean` | no |
| `componentTypeId` | `string` | yes |
| `description` | `string` | no |
| `propertyDefinitions` | `Map<PropertyDefinitionResponse>` | no |
| `extendsFrom` | `List<string>` | no |
| `functions` | `Map<FunctionResponse>` | no |
| `creationDateTime` | `timestamp` | yes |
| `updateDateTime` | `timestamp` | yes |
| `arn` | `string` | yes |
| `isAbstract` | `boolean` | no |
| `isSchemaInitialized` | `boolean` | no |
| `status` | `Status` | no |
| `propertyGroups` | `Map<PropertyGroupResponse>` | no |
| `syncSource` | `string` | no |
| `componentTypeName` | `string` | no |
| `compositeComponentTypes` | `Map<CompositeComponentTypeResponse>` | no |

## GetEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `entityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entityId` | `string` | yes |
| `entityName` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `Status` | yes |
| `workspaceId` | `string` | yes |
| `description` | `string` | no |
| `components` | `Map<ComponentResponse>` | no |
| `parentEntityId` | `string` | yes |
| `hasChildEntities` | `boolean` | yes |
| `creationDateTime` | `timestamp` | yes |
| `updateDateTime` | `timestamp` | yes |
| `syncSource` | `string` | no |
| `areAllComponentsReturned` | `boolean` | no |

## GetMetadataTransferJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metadataTransferJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metadataTransferJobId` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `sources` | `List<SourceConfiguration>` | yes |
| `destination` | `DestinationConfiguration` | yes |
| `metadataTransferJobRole` | `string` | yes |
| `reportUrl` | `string` | no |
| `creationDateTime` | `timestamp` | yes |
| `updateDateTime` | `timestamp` | yes |
| `status` | `MetadataTransferJobStatus` | yes |
| `progress` | `MetadataTransferJobProgress` | no |

## GetPricingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `currentPricingPlan` | `PricingPlan` | yes |
| `pendingPricingPlan` | `PricingPlan` | no |

## GetPropertyValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentName` | `string` | no |
| `componentPath` | `string` | no |
| `componentTypeId` | `string` | no |
| `entityId` | `string` | no |
| `selectedProperties` | `List<string>` | yes |
| `workspaceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `propertyGroupName` | `string` | no |
| `tabularConditions` | `TabularConditions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `propertyValues` | `Map<PropertyLatestValue>` | no |
| `nextToken` | `string` | no |
| `tabularPropertyValues` | `List<List<Map<DataValue>>>` | no |

## GetPropertyValueHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `entityId` | `string` | no |
| `componentName` | `string` | no |
| `componentPath` | `string` | no |
| `componentTypeId` | `string` | no |
| `selectedProperties` | `List<string>` | yes |
| `propertyFilters` | `List<PropertyFilter>` | no |
| `startDateTime` | `timestamp` | no |
| `endDateTime` | `timestamp` | no |
| `interpolation` | `InterpolationParameters` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `orderByTime` | `string` | no |
| `startTime` | `string` | no |
| `endTime` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `propertyValues` | `List<PropertyValueHistory>` | yes |
| `nextToken` | `string` | no |

## GetScene

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `sceneId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `sceneId` | `string` | yes |
| `contentLocation` | `string` | yes |
| `arn` | `string` | yes |
| `creationDateTime` | `timestamp` | yes |
| `updateDateTime` | `timestamp` | yes |
| `description` | `string` | no |
| `capabilities` | `List<string>` | no |
| `sceneMetadata` | `Map<string>` | no |
| `generatedSceneMetadata` | `Map<string>` | no |
| `error` | `SceneError` | no |

## GetSyncJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `syncSource` | `string` | yes |
| `workspaceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `workspaceId` | `string` | yes |
| `syncSource` | `string` | yes |
| `syncRole` | `string` | yes |
| `status` | `SyncJobStatus` | yes |
| `creationDateTime` | `timestamp` | yes |
| `updateDateTime` | `timestamp` | yes |

## GetWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `linkedServices` | `List<string>` | no |
| `s3Location` | `string` | no |
| `role` | `string` | no |
| `creationDateTime` | `timestamp` | yes |
| `updateDateTime` | `timestamp` | yes |

## ListComponentTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `filters` | `List<ListComponentTypesFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `componentTypeSummaries` | `List<ComponentTypeSummary>` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

## ListComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `entityId` | `string` | yes |
| `componentPath` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentSummaries` | `List<ComponentSummary>` | yes |
| `nextToken` | `string` | no |

## ListEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `filters` | `List<ListEntitiesFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entitySummaries` | `List<EntitySummary>` | no |
| `nextToken` | `string` | no |

## ListMetadataTransferJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceType` | `string` | yes |
| `destinationType` | `string` | yes |
| `filters` | `List<ListMetadataTransferJobsFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metadataTransferJobSummaries` | `List<MetadataTransferJobSummary>` | yes |
| `nextToken` | `string` | no |

## ListProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `componentName` | `string` | no |
| `componentPath` | `string` | no |
| `entityId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `propertySummaries` | `List<PropertySummary>` | yes |
| `nextToken` | `string` | no |

## ListScenes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sceneSummaries` | `List<SceneSummary>` | no |
| `nextToken` | `string` | no |

## ListSyncJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `syncJobSummaries` | `List<SyncJobSummary>` | no |
| `nextToken` | `string` | no |

## ListSyncResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `syncSource` | `string` | yes |
| `filters` | `List<SyncResourceFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `syncResources` | `List<SyncResourceSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |
| `nextToken` | `string` | no |

## ListWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceSummaries` | `List<WorkspaceSummary>` | no |
| `nextToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateComponentType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `isSingleton` | `boolean` | no |
| `componentTypeId` | `string` | yes |
| `description` | `string` | no |
| `propertyDefinitions` | `Map<PropertyDefinitionRequest>` | no |
| `extendsFrom` | `List<string>` | no |
| `functions` | `Map<FunctionRequest>` | no |
| `propertyGroups` | `Map<PropertyGroupRequest>` | no |
| `componentTypeName` | `string` | no |
| `compositeComponentTypes` | `Map<CompositeComponentTypeRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `arn` | `string` | yes |
| `componentTypeId` | `string` | yes |
| `state` | `string` | yes |

## UpdateEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `entityId` | `string` | yes |
| `entityName` | `string` | no |
| `description` | `string` | no |
| `componentUpdates` | `Map<ComponentUpdateRequest>` | no |
| `compositeComponentUpdates` | `Map<CompositeComponentUpdateRequest>` | no |
| `parentEntityUpdate` | `ParentEntityUpdateRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateDateTime` | `timestamp` | yes |
| `state` | `string` | yes |

## UpdatePricingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pricingMode` | `string` | yes |
| `bundleNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `currentPricingPlan` | `PricingPlan` | yes |
| `pendingPricingPlan` | `PricingPlan` | no |

## UpdateScene

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `sceneId` | `string` | yes |
| `contentLocation` | `string` | no |
| `description` | `string` | no |
| `capabilities` | `List<string>` | no |
| `sceneMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateDateTime` | `timestamp` | yes |

## UpdateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `description` | `string` | no |
| `role` | `string` | no |
| `s3Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateDateTime` | `timestamp` | yes |

