# Data Automation for Amazon Bedrock

API version: 2023-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bedrock-data-automation/2023-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CopyBlueprintStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprintArn` | `string` | yes |
| `sourceStage` | `string` | yes |
| `targetStage` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprintName` | `string` | yes |
| `type` | `string` | yes |
| `blueprintStage` | `string` | no |
| `schema` | `string` | yes |
| `clientToken` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprint` | `Blueprint` | yes |

## CreateBlueprintVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprintArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprint` | `Blueprint` | yes |

## CreateDataAutomationLibrary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryName` | `string` | yes |
| `libraryDescription` | `string` | no |
| `clientToken` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | no |
| `status` | `string` | no |

## CreateDataAutomationProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |
| `projectDescription` | `string` | no |
| `projectStage` | `string` | no |
| `projectType` | `string` | no |
| `standardOutputConfiguration` | `StandardOutputConfiguration` | yes |
| `customOutputConfiguration` | `CustomOutputConfiguration` | no |
| `overrideConfiguration` | `OverrideConfiguration` | no |
| `dataAutomationLibraryConfiguration` | `DataAutomationLibraryConfiguration` | no |
| `clientToken` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `projectStage` | `string` | no |
| `status` | `string` | no |

## DeleteBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprintArn` | `string` | yes |
| `blueprintVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataAutomationLibrary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | no |
| `status` | `string` | no |

## DeleteDataAutomationProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `status` | `string` | no |

## GetBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprintArn` | `string` | yes |
| `blueprintVersion` | `string` | no |
| `blueprintStage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprint` | `Blueprint` | yes |

## GetBlueprintOptimizationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `errorType` | `string` | no |
| `errorMessage` | `string` | no |
| `outputConfiguration` | `BlueprintOptimizationOutputConfiguration` | no |

## GetDataAutomationLibrary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `library` | `DataAutomationLibrary` | no |

## GetDataAutomationLibraryEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | yes |
| `entityType` | `string` | yes |
| `entityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entity` | `EntityDetails` | no |

## GetDataAutomationLibraryIngestionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | yes |
| `jobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `DataAutomationLibraryIngestionJob` | no |

## GetDataAutomationProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `projectStage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `project` | `DataAutomationProject` | yes |

## InvokeBlueprintOptimizationAsync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprint` | `BlueprintOptimizationObject` | yes |
| `samples` | `List<BlueprintOptimizationSample>` | yes |
| `outputConfiguration` | `BlueprintOptimizationOutputConfiguration` | yes |
| `dataAutomationProfileArn` | `string` | yes |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationArn` | `string` | yes |

## InvokeDataAutomationLibraryIngestionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | yes |
| `clientToken` | `string` | no |
| `inputConfiguration` | `InputConfiguration` | yes |
| `entityType` | `string` | yes |
| `operationType` | `string` | yes |
| `outputConfiguration` | `OutputConfiguration` | yes |
| `notificationConfiguration` | `NotificationConfiguration` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |

## ListBlueprints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprintArn` | `string` | no |
| `resourceOwner` | `string` | no |
| `blueprintStageFilter` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `projectFilter` | `DataAutomationProjectFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprints` | `List<BlueprintSummary>` | yes |
| `nextToken` | `string` | no |

## ListDataAutomationLibraries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `projectFilter` | `DataAutomationProjectFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraries` | `List<DataAutomationLibrarySummary>` | no |
| `nextToken` | `string` | no |

## ListDataAutomationLibraryEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | yes |
| `entityType` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<DataAutomationLibraryEntitySummary>` | no |
| `nextToken` | `string` | no |

## ListDataAutomationLibraryIngestionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<DataAutomationLibraryIngestionJobSummary>` | no |
| `nextToken` | `string` | no |

## ListDataAutomationProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `projectStageFilter` | `string` | no |
| `blueprintFilter` | `BlueprintFilter` | no |
| `resourceOwner` | `string` | no |
| `libraryFilter` | `DataAutomationLibraryFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projects` | `List<DataAutomationProjectSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprintArn` | `string` | yes |
| `schema` | `string` | yes |
| `blueprintStage` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprint` | `Blueprint` | yes |

## UpdateDataAutomationLibrary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | yes |
| `libraryDescription` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryArn` | `string` | no |
| `status` | `string` | no |

## UpdateDataAutomationProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `projectStage` | `string` | no |
| `projectDescription` | `string` | no |
| `standardOutputConfiguration` | `StandardOutputConfiguration` | yes |
| `customOutputConfiguration` | `CustomOutputConfiguration` | no |
| `overrideConfiguration` | `OverrideConfiguration` | no |
| `dataAutomationLibraryConfiguration` | `DataAutomationLibraryConfiguration` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `projectStage` | `string` | no |
| `status` | `string` | no |

