# AWS IoT SiteWise

API version: 2019-12-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iotsitewise/2019-12-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `hierarchyId` | `string` | yes |
| `childAssetId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateTimeSeriesToAssetProperty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alias` | `string` | yes |
| `assetId` | `string` | yes |
| `propertyId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchAssociateDataSegmentsToDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `associateDataSegmentEntries` | `List<AssociateDataSegmentEntry>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `datasetVersion` | `string` | yes |
| `failedAssociations` | `List<FailedDataSegmentAssociation>` | yes |

## BatchAssociateProjectAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `assetIds` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<AssetErrorDetails>` | no |

## BatchDeleteDatasetDataSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `deleteDataSegmentEntries` | `List<DeleteDataSegmentEntry>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `datasetVersion` | `string` | yes |
| `errors` | `List<FailedDataSegmentDeletion>` | yes |

## BatchDisassociateDataSegmentsFromDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `disassociateDataSegmentEntries` | `List<DisassociateDataSegmentEntry>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `datasetVersion` | `string` | yes |
| `failedDisassociations` | `List<FailedDataSegmentDisassociation>` | yes |

## BatchDisassociateProjectAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `assetIds` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<AssetErrorDetails>` | no |

## BatchGetAssetPropertyAggregates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entries` | `List<BatchGetAssetPropertyAggregatesEntry>` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorEntries` | `List<BatchGetAssetPropertyAggregatesErrorEntry>` | yes |
| `successEntries` | `List<BatchGetAssetPropertyAggregatesSuccessEntry>` | yes |
| `skippedEntries` | `List<BatchGetAssetPropertyAggregatesSkippedEntry>` | yes |
| `nextToken` | `string` | no |

## BatchGetAssetPropertyValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entries` | `List<BatchGetAssetPropertyValueEntry>` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorEntries` | `List<BatchGetAssetPropertyValueErrorEntry>` | yes |
| `successEntries` | `List<BatchGetAssetPropertyValueSuccessEntry>` | yes |
| `skippedEntries` | `List<BatchGetAssetPropertyValueSkippedEntry>` | yes |
| `nextToken` | `string` | no |

## BatchGetAssetPropertyValueHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entries` | `List<BatchGetAssetPropertyValueHistoryEntry>` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorEntries` | `List<BatchGetAssetPropertyValueHistoryErrorEntry>` | yes |
| `successEntries` | `List<BatchGetAssetPropertyValueHistorySuccessEntry>` | yes |
| `skippedEntries` | `List<BatchGetAssetPropertyValueHistorySkippedEntry>` | yes |
| `nextToken` | `string` | no |

## BatchPutAssetPropertyValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enablePartialEntryProcessing` | `boolean` | no |
| `entries` | `List<PutAssetPropertyValueEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorEntries` | `List<BatchPutAssetPropertyErrorEntry>` | yes |

## CancelEnrichmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `status` | `string` | yes |

## CancelPipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `pipelineName` | `string` | yes |
| `pipelineExecutionId` | `string` | yes |
| `reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `state` | `string` | yes |

## CancelQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |
| `status` | `string` | yes |

## CreateAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyIdentity` | `Identity` | yes |
| `accessPolicyResource` | `Resource` | yes |
| `accessPolicyPermission` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyId` | `string` | yes |
| `accessPolicyArn` | `string` | yes |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `idcInstanceArn` | `string` | yes |
| `workspaceName` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `dnsSubdomain` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |

## CreateAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetName` | `string` | yes |
| `assetModelId` | `string` | yes |
| `assetId` | `string` | no |
| `assetExternalId` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `assetDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `assetArn` | `string` | yes |
| `assetStatus` | `AssetStatus` | yes |

## CreateAssetModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelName` | `string` | yes |
| `assetModelType` | `string` | no |
| `assetModelId` | `string` | no |
| `assetModelExternalId` | `string` | no |
| `assetModelDescription` | `string` | no |
| `assetModelProperties` | `List<AssetModelPropertyDefinition>` | no |
| `assetModelHierarchies` | `List<AssetModelHierarchyDefinition>` | no |
| `assetModelCompositeModels` | `List<AssetModelCompositeModelDefinition>` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `assetModelArn` | `string` | yes |
| `assetModelStatus` | `AssetModelStatus` | yes |

## CreateAssetModelCompositeModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `assetModelCompositeModelExternalId` | `string` | no |
| `parentAssetModelCompositeModelId` | `string` | no |
| `assetModelCompositeModelId` | `string` | no |
| `assetModelCompositeModelDescription` | `string` | no |
| `assetModelCompositeModelName` | `string` | yes |
| `assetModelCompositeModelType` | `string` | yes |
| `clientToken` | `string` | no |
| `composedAssetModelId` | `string` | no |
| `assetModelCompositeModelProperties` | `List<AssetModelPropertyDefinition>` | no |
| `ifMatch` | `string` | no |
| `ifNoneMatch` | `string` | no |
| `matchForVersionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelCompositeModelId` | `string` | yes |
| `assetModelCompositeModelPath` | `List<AssetModelCompositeModelPathSegment>` | yes |
| `assetModelStatus` | `AssetModelStatus` | yes |
| `assetModelId` | `string` | no |

## CreateBulkImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `jobRoleArn` | `string` | yes |
| `files` | `List<File>` | yes |
| `errorReportLocation` | `ErrorReportLocation` | yes |
| `jobConfiguration` | `JobConfiguration` | no |
| `adaptiveIngestion` | `boolean` | no |
| `deleteFilesAfterImport` | `boolean` | no |
| `datasetId` | `string` | no |
| `workspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `jobName` | `string` | yes |
| `jobStatus` | `string` | yes |

## CreateComputationModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelName` | `string` | yes |
| `computationModelDescription` | `string` | no |
| `computationModelConfiguration` | `ComputationModelConfiguration` | yes |
| `computationModelDataBinding` | `Map<ComputationModelDataBindingValue>` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelId` | `string` | yes |
| `computationModelArn` | `string` | yes |
| `computationModelStatus` | `ComputationModelStatus` | yes |

## CreateDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `dashboardName` | `string` | yes |
| `dashboardDescription` | `string` | no |
| `dashboardDefinition` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dashboardId` | `string` | yes |
| `dashboardArn` | `string` | yes |

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | no |
| `datasetName` | `string` | yes |
| `datasetDescription` | `string` | no |
| `datasetType` | `string` | no |
| `datasetConfig` | `DatasetConfig` | no |
| `workspaceName` | `string` | no |
| `metadata` | `Map<string>` | no |
| `datasetSource` | `DatasetSource` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `datasetArn` | `string` | yes |
| `datasetStatus` | `DatasetStatus` | yes |

## CreateDatasetExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `clientToken` | `string` | no |
| `destinationS3Uri` | `string` | yes |
| `input` | `ProcessingInput` | yes |
| `errorReportLocation` | `ExportErrorReportLocation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `workspaceName` | `string` | yes |

## CreateEnrichmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `jobConfiguration` | `EnrichmentJobConfiguration` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayName` | `string` | yes |
| `gatewayPlatform` | `GatewayPlatform` | yes |
| `gatewayVersion` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `gatewayArn` | `string` | yes |

## CreatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `pipelineName` | `string` | yes |
| `description` | `string` | no |
| `environmentVariables` | `Map<string>` | no |
| `computations` | `List<ComputeNode>` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `pipelineArn` | `string` | yes |
| `version` | `string` | yes |
| `status` | `ResourceStatus` | yes |

## CreatePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalName` | `string` | yes |
| `portalDescription` | `string` | no |
| `portalContactEmail` | `string` | yes |
| `clientToken` | `string` | no |
| `portalLogoImageFile` | `ImageFile` | no |
| `roleArn` | `string` | yes |
| `tags` | `Map<string>` | no |
| `portalAuthMode` | `string` | no |
| `notificationSenderEmail` | `string` | no |
| `alarms` | `Alarms` | no |
| `portalType` | `string` | no |
| `portalTypeConfiguration` | `Map<PortalTypeEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `portalArn` | `string` | yes |
| `portalStartUrl` | `string` | yes |
| `portalStatus` | `PortalStatus` | yes |
| `ssoApplicationId` | `string` | yes |

## CreateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `projectName` | `string` | yes |
| `projectDescription` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `projectArn` | `string` | yes |

## CreateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `taskName` | `string` | yes |
| `description` | `string` | no |
| `taskConfiguration` | `TaskConfiguration` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskName` | `string` | yes |
| `taskArn` | `string` | yes |
| `version` | `string` | yes |
| `status` | `ResourceStatus` | yes |

## CreateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `workspaceDescription` | `string` | no |
| `encryptionConfiguration` | `WorkspaceEncryptionConfiguration` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `workspaceArn` | `string` | yes |
| `workspaceStatus` | `WorkspaceStatus` | yes |

## DeleteAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | no |
| `assetStatus` | `AssetStatus` | yes |

## DeleteAssetModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `clientToken` | `string` | no |
| `ifMatch` | `string` | no |
| `ifNoneMatch` | `string` | no |
| `matchForVersionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | no |
| `assetModelStatus` | `AssetModelStatus` | yes |

## DeleteAssetModelCompositeModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `assetModelCompositeModelId` | `string` | yes |
| `clientToken` | `string` | no |
| `ifMatch` | `string` | no |
| `ifNoneMatch` | `string` | no |
| `matchForVersionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelStatus` | `AssetModelStatus` | yes |
| `assetModelId` | `string` | no |

## DeleteAssetModelInterfaceRelationship

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `interfaceAssetModelId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `interfaceAssetModelId` | `string` | yes |
| `assetModelArn` | `string` | yes |
| `assetModelStatus` | `AssetModelStatus` | yes |

## DeleteComputationModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelStatus` | `ComputationModelStatus` | yes |

## DeleteDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dashboardId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `workspaceName` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetStatus` | `DatasetStatus` | yes |

## DeleteGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `pipelineName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `ResourceStatus` | yes |

## DeletePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalStatus` | `PortalStatus` | yes |

## DeleteProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `taskName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `ResourceStatus` | yes |

## DeleteTimeSeries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alias` | `string` | no |
| `assetId` | `string` | no |
| `propertyId` | `string` | no |
| `clientToken` | `string` | no |
| `workspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceStatus` | `WorkspaceStatus` | yes |

## DescribeAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyId` | `string` | yes |
| `accessPolicyArn` | `string` | yes |
| `accessPolicyIdentity` | `Identity` | yes |
| `accessPolicyResource` | `Resource` | yes |
| `accessPolicyPermission` | `string` | yes |
| `accessPolicyCreationDate` | `timestamp` | yes |
| `accessPolicyLastUpdateDate` | `timestamp` | yes |

## DescribeAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionId` | `string` | yes |
| `targetResource` | `TargetResource` | yes |
| `actionDefinitionId` | `string` | yes |
| `actionPayload` | `ActionPayload` | yes |
| `executionTime` | `timestamp` | yes |
| `resolveTo` | `ResolveTo` | no |

## DescribeApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `dnsSubdomain` | `string` | yes |
| `description` | `string` | no |
| `id` | `string` | yes |
| `idcApplicationArn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |
| `workspaceName` | `string` | yes |

## DescribeAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `excludeProperties` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `assetExternalId` | `string` | no |
| `assetArn` | `string` | yes |
| `assetName` | `string` | yes |
| `assetModelId` | `string` | yes |
| `assetProperties` | `List<AssetProperty>` | yes |
| `assetHierarchies` | `List<AssetHierarchy>` | yes |
| `assetCompositeModels` | `List<AssetCompositeModel>` | no |
| `assetCreationDate` | `timestamp` | yes |
| `assetLastUpdateDate` | `timestamp` | yes |
| `assetStatus` | `AssetStatus` | yes |
| `assetDescription` | `string` | no |
| `assetCompositeModelSummaries` | `List<AssetCompositeModelSummary>` | no |

## DescribeAssetCompositeModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `assetCompositeModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `assetCompositeModelId` | `string` | yes |
| `assetCompositeModelExternalId` | `string` | no |
| `assetCompositeModelPath` | `List<AssetCompositeModelPathSegment>` | yes |
| `assetCompositeModelName` | `string` | yes |
| `assetCompositeModelDescription` | `string` | yes |
| `assetCompositeModelType` | `string` | yes |
| `assetCompositeModelProperties` | `List<AssetProperty>` | yes |
| `assetCompositeModelSummaries` | `List<AssetCompositeModelSummary>` | yes |
| `actionDefinitions` | `List<ActionDefinition>` | no |

## DescribeAssetModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `excludeProperties` | `boolean` | no |
| `assetModelVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `assetModelExternalId` | `string` | no |
| `assetModelArn` | `string` | yes |
| `assetModelName` | `string` | yes |
| `assetModelType` | `string` | no |
| `assetModelDescription` | `string` | yes |
| `assetModelProperties` | `List<AssetModelProperty>` | yes |
| `assetModelHierarchies` | `List<AssetModelHierarchy>` | yes |
| `assetModelCompositeModels` | `List<AssetModelCompositeModel>` | no |
| `assetModelCompositeModelSummaries` | `List<AssetModelCompositeModelSummary>` | no |
| `assetModelCreationDate` | `timestamp` | yes |
| `assetModelLastUpdateDate` | `timestamp` | yes |
| `assetModelStatus` | `AssetModelStatus` | yes |
| `assetModelVersion` | `string` | no |
| `interfaceDetails` | `List<InterfaceRelationship>` | no |
| `eTag` | `string` | no |

## DescribeAssetModelCompositeModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `assetModelCompositeModelId` | `string` | yes |
| `assetModelVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `assetModelCompositeModelId` | `string` | yes |
| `assetModelCompositeModelExternalId` | `string` | no |
| `assetModelCompositeModelPath` | `List<AssetModelCompositeModelPathSegment>` | yes |
| `assetModelCompositeModelName` | `string` | yes |
| `assetModelCompositeModelDescription` | `string` | yes |
| `assetModelCompositeModelType` | `string` | yes |
| `assetModelCompositeModelProperties` | `List<AssetModelProperty>` | yes |
| `compositionDetails` | `CompositionDetails` | no |
| `assetModelCompositeModelSummaries` | `List<AssetModelCompositeModelSummary>` | yes |
| `actionDefinitions` | `List<ActionDefinition>` | no |

## DescribeAssetModelInterfaceRelationship

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `interfaceAssetModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `interfaceAssetModelId` | `string` | yes |
| `propertyMappings` | `List<PropertyMapping>` | yes |
| `hierarchyMappings` | `List<HierarchyMapping>` | yes |

## DescribeAssetProperty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `propertyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `assetExternalId` | `string` | no |
| `assetName` | `string` | yes |
| `assetModelId` | `string` | yes |
| `assetProperty` | `Property` | no |
| `compositeModel` | `CompositeModelProperty` | no |

## DescribeBulkImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `workspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `jobName` | `string` | yes |
| `jobStatus` | `string` | yes |
| `jobRoleArn` | `string` | yes |
| `files` | `List<File>` | yes |
| `errorReportLocation` | `ErrorReportLocation` | yes |
| `jobConfiguration` | `JobConfiguration` | no |
| `jobCreationDate` | `timestamp` | yes |
| `jobLastUpdateDate` | `timestamp` | yes |
| `adaptiveIngestion` | `boolean` | no |
| `deleteFilesAfterImport` | `boolean` | no |
| `datasetId` | `string` | no |
| `workspaceName` | `string` | no |

## DescribeComputationModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelId` | `string` | yes |
| `computationModelVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelId` | `string` | yes |
| `computationModelArn` | `string` | yes |
| `computationModelName` | `string` | yes |
| `computationModelDescription` | `string` | no |
| `computationModelConfiguration` | `ComputationModelConfiguration` | yes |
| `computationModelDataBinding` | `Map<ComputationModelDataBindingValue>` | yes |
| `computationModelCreationDate` | `timestamp` | yes |
| `computationModelLastUpdateDate` | `timestamp` | yes |
| `computationModelStatus` | `ComputationModelStatus` | yes |
| `computationModelVersion` | `string` | yes |
| `actionDefinitions` | `List<ActionDefinition>` | yes |

## DescribeComputationModelExecutionSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelId` | `string` | yes |
| `resolveToResourceType` | `string` | no |
| `resolveToResourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelId` | `string` | yes |
| `resolveTo` | `ResolveTo` | no |
| `computationModelExecutionSummary` | `Map<string>` | yes |

## DescribeDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dashboardId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dashboardId` | `string` | yes |
| `dashboardArn` | `string` | yes |
| `dashboardName` | `string` | yes |
| `projectId` | `string` | yes |
| `dashboardDescription` | `string` | no |
| `dashboardDefinition` | `string` | yes |
| `dashboardCreationDate` | `timestamp` | yes |
| `dashboardLastUpdateDate` | `timestamp` | yes |

## DescribeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `workspaceName` | `string` | no |
| `datasetVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `datasetArn` | `string` | yes |
| `datasetName` | `string` | yes |
| `datasetDescription` | `string` | yes |
| `datasetType` | `string` | no |
| `datasetConfig` | `DatasetConfig` | no |
| `workspaceName` | `string` | no |
| `metadata` | `Map<string>` | no |
| `datasetSource` | `DatasetSource` | yes |
| `datasetStatus` | `DatasetStatus` | yes |
| `datasetCreationDate` | `timestamp` | yes |
| `datasetLastUpdateDate` | `timestamp` | yes |
| `datasetVersion` | `string` | no |
| `enrichmentStatus` | `DatasetEnrichment` | no |

## DescribeDatasetExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `status` | `string` | yes |
| `startedAt` | `timestamp` | yes |
| `completedAt` | `timestamp` | no |
| `destinationS3Uri` | `string` | yes |
| `errorReportLocation` | `ExportErrorReportLocation` | yes |
| `input` | `ProcessingInput` | yes |

## DescribeDefaultEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionType` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `configurationStatus` | `ConfigurationStatus` | yes |

## DescribeEnrichmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `status` | `string` | yes |
| `workspaceName` | `string` | yes |
| `jobType` | `string` | yes |
| `jobConfiguration` | `EnrichmentJobConfiguration` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | no |
| `completedAt` | `timestamp` | no |
| `cancelledAt` | `timestamp` | no |
| `failureMessage` | `string` | no |

## DescribeExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | yes |
| `actionType` | `string` | no |
| `targetResource` | `TargetResource` | yes |
| `targetResourceVersion` | `string` | yes |
| `resolveTo` | `ResolveTo` | no |
| `executionStartTime` | `timestamp` | yes |
| `executionEndTime` | `timestamp` | no |
| `executionStatus` | `ExecutionStatus` | yes |
| `executionResult` | `Map<string>` | no |
| `executionDetails` | `Map<string>` | no |
| `executionEntityVersion` | `string` | no |

## DescribeGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `gatewayName` | `string` | yes |
| `gatewayArn` | `string` | yes |
| `gatewayPlatform` | `GatewayPlatform` | no |
| `gatewayVersion` | `string` | no |
| `gatewayCapabilitySummaries` | `List<GatewayCapabilitySummary>` | yes |
| `creationDate` | `timestamp` | yes |
| `lastUpdateDate` | `timestamp` | yes |

## DescribeGatewayCapabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `capabilityNamespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `capabilityNamespace` | `string` | yes |
| `capabilityConfiguration` | `string` | yes |
| `capabilitySyncStatus` | `string` | yes |

## DescribeLoggingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loggingOptions` | `LoggingOptions` | yes |

## DescribePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `pipelineName` | `string` | yes |
| `pipelineVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `workspaceName` | `string` | yes |
| `description` | `string` | no |
| `pipelineArn` | `string` | yes |
| `version` | `string` | yes |
| `environmentVariables` | `Map<string>` | no |
| `computations` | `List<ComputeNode>` | yes |
| `status` | `ResourceStatus` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## DescribePipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `pipelineName` | `string` | yes |
| `pipelineExecutionId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecutionId` | `string` | yes |
| `pipelineName` | `string` | yes |
| `workspaceName` | `string` | yes |
| `pipelineVersion` | `string` | yes |
| `status` | `PipelineExecutionStatus` | yes |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `requestEnvironmentVariables` | `ExecutionEnvironmentVariables` | yes |
| `requestMountOverrides` | `MountOverrides` | no |
| `executionPriority` | `integer` | no |
| `computeNodeExecutionDetails` | `List<ComputeNodeExecutionDetails>` | yes |
| `nextToken` | `string` | no |

## DescribePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `portalArn` | `string` | yes |
| `portalName` | `string` | yes |
| `portalDescription` | `string` | no |
| `portalClientId` | `string` | yes |
| `portalStartUrl` | `string` | yes |
| `portalContactEmail` | `string` | yes |
| `portalStatus` | `PortalStatus` | yes |
| `portalCreationDate` | `timestamp` | yes |
| `portalLastUpdateDate` | `timestamp` | yes |
| `portalLogoImageLocation` | `ImageLocation` | no |
| `roleArn` | `string` | no |
| `portalAuthMode` | `string` | no |
| `notificationSenderEmail` | `string` | no |
| `alarms` | `Alarms` | no |
| `portalType` | `string` | no |
| `portalTypeConfiguration` | `Map<PortalTypeEntry>` | no |

## DescribeProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `projectArn` | `string` | yes |
| `projectName` | `string` | yes |
| `portalId` | `string` | yes |
| `projectDescription` | `string` | no |
| `projectCreationDate` | `timestamp` | yes |
| `projectLastUpdateDate` | `timestamp` | yes |

## DescribeQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |
| `status` | `string` | yes |
| `submittedAt` | `timestamp` | yes |
| `completedAt` | `timestamp` | no |
| `statistics` | `QueryStatistics` | no |
| `errorMessage` | `string` | no |

## DescribeSearch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `searchId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `status` | `string` | yes |
| `queryStatement` | `string` | yes |
| `searchType` | `string` | yes |
| `statusReason` | `string` | no |
| `startedAt` | `timestamp` | no |
| `groupId` | `string` | no |

## DescribeStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageType` | `string` | yes |
| `multiLayerStorage` | `MultiLayerStorage` | no |
| `disassociatedDataStorage` | `string` | no |
| `retentionPeriod` | `RetentionPeriod` | no |
| `configurationStatus` | `ConfigurationStatus` | yes |
| `lastUpdateDate` | `timestamp` | no |
| `warmTier` | `string` | no |
| `warmTierRetentionPeriod` | `WarmTierRetentionPeriod` | no |
| `disallowIngestNullNaN` | `boolean` | no |

## DescribeTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `taskName` | `string` | yes |
| `taskVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `taskName` | `string` | yes |
| `description` | `string` | no |
| `taskArn` | `string` | yes |
| `version` | `string` | yes |
| `taskConfiguration` | `TaskConfiguration` | yes |
| `status` | `ResourceStatus` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## DescribeTimeSeries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alias` | `string` | no |
| `assetId` | `string` | no |
| `propertyId` | `string` | no |
| `workspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | no |
| `propertyId` | `string` | no |
| `alias` | `string` | no |
| `timeSeriesId` | `string` | yes |
| `dataType` | `string` | yes |
| `dataTypeSpec` | `string` | no |
| `timeSeriesCreationDate` | `timestamp` | yes |
| `timeSeriesLastUpdateDate` | `timestamp` | yes |
| `timeSeriesArn` | `string` | yes |
| `workspaceName` | `string` | no |

## DescribeWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceArn` | `string` | yes |
| `workspaceName` | `string` | yes |
| `workspaceDescription` | `string` | no |
| `workspaceStatus` | `WorkspaceStatus` | yes |
| `encryptionConfiguration` | `WorkspaceEncryptionConfigurationInfo` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## DisassociateAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `hierarchyId` | `string` | yes |
| `childAssetId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateTimeSeriesFromAssetProperty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alias` | `string` | yes |
| `assetId` | `string` | yes |
| `propertyId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExecuteAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetResource` | `TargetResource` | yes |
| `actionDefinitionId` | `string` | yes |
| `actionPayload` | `ActionPayload` | yes |
| `clientToken` | `string` | no |
| `resolveTo` | `ResolveTo` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionId` | `string` | yes |

## ExecuteQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryStatement` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `columns` | `List<ColumnInfo>` | no |
| `rows` | `List<Row>` | no |
| `nextToken` | `string` | no |

## GetAssetPropertyAggregates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | no |
| `propertyId` | `string` | no |
| `propertyAlias` | `string` | no |
| `aggregateTypes` | `List<string>` | yes |
| `resolution` | `string` | yes |
| `qualities` | `List<string>` | no |
| `startDate` | `timestamp` | yes |
| `endDate` | `timestamp` | yes |
| `timeOrdering` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aggregatedValues` | `List<AggregatedValue>` | yes |
| `nextToken` | `string` | no |

## GetAssetPropertyValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | no |
| `propertyId` | `string` | no |
| `propertyAlias` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `propertyValue` | `AssetPropertyValue` | no |

## GetAssetPropertyValueHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | no |
| `propertyId` | `string` | no |
| `propertyAlias` | `string` | no |
| `startDate` | `timestamp` | no |
| `endDate` | `timestamp` | no |
| `qualities` | `List<string>` | no |
| `timeOrdering` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetPropertyValueHistory` | `List<AssetPropertyValue>` | yes |
| `nextToken` | `string` | no |

## GetCaptureData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `startTime` | `TimeInNanos` | yes |
| `endTime` | `TimeInNanos` | yes |
| `timeSeriesId` | `string` | no |
| `propertyAlias` | `string` | no |
| `formatSettings` | `FormatSettings` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `data` | `blob` | yes |
| `startTime` | `TimeInNanos` | yes |
| `endTime` | `TimeInNanos` | yes |
| `dataType` | `string` | yes |
| `nextToken` | `string` | no |

## GetInterpolatedAssetPropertyValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | no |
| `propertyId` | `string` | no |
| `propertyAlias` | `string` | no |
| `startTimeInSeconds` | `long` | yes |
| `startTimeOffsetInNanos` | `integer` | no |
| `endTimeInSeconds` | `long` | yes |
| `endTimeOffsetInNanos` | `integer` | no |
| `quality` | `string` | yes |
| `intervalInSeconds` | `long` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `type` | `string` | yes |
| `intervalWindowInSeconds` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interpolatedAssetPropertyValues` | `List<InterpolatedAssetPropertyValue>` | yes |
| `nextToken` | `string` | no |

## GetQueryResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `queryId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `columnInfo` | `List<ColumnInformation>` | no |
| `rows` | `List<List<string>>` | no |
| `nextToken` | `string` | no |

## GetSearchResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchResults` | `List<SearchResult>` | yes |
| `nextToken` | `string` | no |

## InvokeAssistant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `conversationId` | `string` | no |
| `message` | `string` | yes |
| `enableTrace` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `ResponseStream` | yes |
| `conversationId` | `string` | yes |

## ListAccessPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityType` | `string` | no |
| `identityId` | `string` | no |
| `resourceType` | `string` | no |
| `resourceId` | `string` | no |
| `iamArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicySummaries` | `List<AccessPolicySummary>` | yes |
| `nextToken` | `string` | no |

## ListActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetResourceType` | `string` | yes |
| `targetResourceId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `resolveToResourceType` | `string` | no |
| `resolveToResourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionSummaries` | `List<ActionSummary>` | yes |
| `nextToken` | `string` | yes |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `applications` | `List<ApplicationSummary>` | yes |

## ListAssetModelCompositeModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `assetModelVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelCompositeModelSummaries` | `List<AssetModelCompositeModelSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssetModelProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `string` | no |
| `assetModelVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelPropertySummaries` | `List<AssetModelPropertySummary>` | yes |
| `nextToken` | `string` | no |

## ListAssetModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelTypes` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `assetModelVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelSummaries` | `List<AssetModelSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssetProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetPropertySummaries` | `List<AssetPropertySummary>` | yes |
| `nextToken` | `string` | no |

## ListAssetRelationships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `traversalType` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetRelationshipSummaries` | `List<AssetRelationshipSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `assetModelId` | `string` | no |
| `filter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetSummaries` | `List<AssetSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssociatedAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `hierarchyId` | `string` | no |
| `traversalDirection` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetSummaries` | `List<AssociatedAssetsSummary>` | yes |
| `nextToken` | `string` | no |

## ListBulkImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `string` | no |
| `workspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummaries` | `List<JobSummary>` | yes |
| `nextToken` | `string` | no |

## ListCompositionRelationships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `compositionRelationshipSummaries` | `List<CompositionRelationshipSummary>` | yes |
| `nextToken` | `string` | no |

## ListComputationModelDataBindingUsages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataBindingValueFilter` | `DataBindingValueFilter` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataBindingUsageSummaries` | `List<ComputationModelDataBindingUsageSummary>` | yes |
| `nextToken` | `string` | no |

## ListComputationModelResolveToResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelResolveToResourceSummaries` | `List<ComputationModelResolveToResourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListComputationModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelSummaries` | `List<ComputationModelSummary>` | yes |
| `nextToken` | `string` | no |

## ListDashboards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dashboardSummaries` | `List<DashboardSummary>` | yes |
| `nextToken` | `string` | no |

## ListDatasetDataSegmentRelationships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSegmentRelationshipSummaries` | `List<DataSegmentRelationshipSummary>` | yes |
| `nextToken` | `string` | no |

## ListDatasetDataSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `datasetVersion` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSegments` | `List<DataSegmentSummary>` | yes |
| `nextToken` | `string` | no |

## ListDatasetExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `filter` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<ExportJobSummary>` | yes |
| `nextToken` | `string` | no |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceType` | `string` | yes |
| `workspaceName` | `string` | no |
| `datasetType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetSummaries` | `List<DatasetSummary>` | yes |
| `nextToken` | `string` | no |
| `workspaceName` | `string` | no |

## ListEnrichmentJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `datasetId` | `string` | no |
| `propertyAlias` | `string` | no |
| `timeSeriesId` | `string` | no |
| `status` | `string` | no |
| `jobType` | `string` | no |
| `startDate` | `timestamp` | no |
| `endDate` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<EnrichmentJobSummary>` | yes |
| `nextToken` | `string` | no |

## ListExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetResourceType` | `string` | yes |
| `targetResourceId` | `string` | yes |
| `resolveToResourceType` | `string` | no |
| `resolveToResourceId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `actionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionSummaries` | `List<ExecutionSummary>` | yes |
| `nextToken` | `string` | no |

## ListGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewaySummaries` | `List<GatewaySummary>` | yes |
| `nextToken` | `string` | no |

## ListInterfaceRelationships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interfaceAssetModelId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interfaceRelationshipSummaries` | `List<InterfaceRelationshipSummary>` | yes |
| `nextToken` | `string` | no |

## ListPipelineExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `pipelineName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `state` | `string` | no |
| `startTimeAfter` | `timestamp` | no |
| `startTimeBefore` | `timestamp` | no |
| `endTimeAfter` | `timestamp` | no |
| `endTimeBefore` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecutionSummaries` | `List<PipelineExecutionSummary>` | yes |
| `nextToken` | `string` | no |

## ListPipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineSummaries` | `List<PipelineSummary>` | yes |
| `nextToken` | `string` | no |

## ListPortals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalSummaries` | `List<PortalSummary>` | no |
| `nextToken` | `string` | no |

## ListProjectAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetIds` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectSummaries` | `List<ProjectSummary>` | yes |
| `nextToken` | `string` | no |

## ListQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `filter` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queries` | `List<QuerySummary>` | yes |
| `nextToken` | `string` | no |

## ListSearches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `listSearchesFilters` | `ListSearchesFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchSummaries` | `List<SearchSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskSummaries` | `List<TaskSummary>` | yes |
| `nextToken` | `string` | no |

## ListTimeSeries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `assetId` | `string` | no |
| `aliasPrefix` | `string` | no |
| `timeSeriesType` | `string` | no |
| `workspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimeSeriesSummaries` | `List<TimeSeriesSummary>` | yes |
| `nextToken` | `string` | no |
| `workspaceName` | `string` | no |

## ListWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceSummaries` | `List<WorkspaceSummary>` | yes |
| `nextToken` | `string` | no |

## PutAssetModelInterfaceRelationship

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `interfaceAssetModelId` | `string` | yes |
| `propertyMappingConfiguration` | `PropertyMappingConfiguration` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `interfaceAssetModelId` | `string` | yes |
| `assetModelArn` | `string` | yes |
| `assetModelStatus` | `AssetModelStatus` | yes |

## PutDefaultEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionType` | `string` | yes |
| `kmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionType` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `configurationStatus` | `ConfigurationStatus` | yes |

## PutLoggingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loggingOptions` | `LoggingOptions` | yes |
| `workspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageType` | `string` | yes |
| `multiLayerStorage` | `MultiLayerStorage` | no |
| `disassociatedDataStorage` | `string` | no |
| `retentionPeriod` | `RetentionPeriod` | no |
| `warmTier` | `string` | no |
| `warmTierRetentionPeriod` | `WarmTierRetentionPeriod` | no |
| `disallowIngestNullNaN` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageType` | `string` | yes |
| `multiLayerStorage` | `MultiLayerStorage` | no |
| `disassociatedDataStorage` | `string` | no |
| `retentionPeriod` | `RetentionPeriod` | no |
| `configurationStatus` | `ConfigurationStatus` | yes |
| `warmTier` | `string` | no |
| `warmTierRetentionPeriod` | `WarmTierRetentionPeriod` | no |
| `disallowIngestNullNaN` | `boolean` | no |

## StartPipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `pipelineName` | `string` | yes |
| `executionEnvironmentVariableOverrides` | `ExecutionEnvironmentVariables` | no |
| `executionMountOverrides` | `MountOverrides` | no |
| `executionPriority` | `integer` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecutionId` | `string` | yes |

## StartQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `workspaceName` | `string` | yes |
| `queryStatement` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |
| `status` | `string` | yes |

## StartSearch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `queryStatement` | `string` | yes |
| `clientToken` | `string` | no |
| `searchType` | `string` | no |
| `searchFilters` | `SearchFilters` | no |
| `groupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchId` | `string` | yes |
| `workspaceName` | `string` | yes |
| `status` | `string` | yes |
| `groupId` | `string` | no |

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


## UpdateAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyId` | `string` | yes |
| `accessPolicyIdentity` | `Identity` | yes |
| `accessPolicyResource` | `Resource` | yes |
| `accessPolicyPermission` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `assetExternalId` | `string` | no |
| `assetName` | `string` | yes |
| `clientToken` | `string` | no |
| `assetDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | no |
| `assetStatus` | `AssetStatus` | yes |

## UpdateAssetModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `assetModelExternalId` | `string` | no |
| `assetModelName` | `string` | yes |
| `assetModelDescription` | `string` | no |
| `assetModelProperties` | `List<AssetModelProperty>` | no |
| `assetModelHierarchies` | `List<AssetModelHierarchy>` | no |
| `assetModelCompositeModels` | `List<AssetModelCompositeModel>` | no |
| `clientToken` | `string` | no |
| `ifMatch` | `string` | no |
| `ifNoneMatch` | `string` | no |
| `matchForVersionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | no |
| `assetModelStatus` | `AssetModelStatus` | yes |

## UpdateAssetModelCompositeModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelId` | `string` | yes |
| `assetModelCompositeModelId` | `string` | yes |
| `assetModelCompositeModelExternalId` | `string` | no |
| `assetModelCompositeModelDescription` | `string` | no |
| `assetModelCompositeModelName` | `string` | yes |
| `clientToken` | `string` | no |
| `assetModelCompositeModelProperties` | `List<AssetModelProperty>` | no |
| `ifMatch` | `string` | no |
| `ifNoneMatch` | `string` | no |
| `matchForVersionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetModelCompositeModelPath` | `List<AssetModelCompositeModelPathSegment>` | yes |
| `assetModelStatus` | `AssetModelStatus` | yes |
| `assetModelId` | `string` | no |

## UpdateAssetProperty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assetId` | `string` | yes |
| `propertyId` | `string` | yes |
| `propertyAlias` | `string` | no |
| `propertyNotificationState` | `string` | no |
| `clientToken` | `string` | no |
| `propertyUnit` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateComputationModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelId` | `string` | yes |
| `computationModelName` | `string` | yes |
| `computationModelDescription` | `string` | no |
| `computationModelConfiguration` | `ComputationModelConfiguration` | yes |
| `computationModelDataBinding` | `Map<ComputationModelDataBindingValue>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computationModelStatus` | `ComputationModelStatus` | yes |

## UpdateDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dashboardId` | `string` | yes |
| `dashboardName` | `string` | yes |
| `dashboardDescription` | `string` | no |
| `dashboardDefinition` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `workspaceName` | `string` | no |
| `datasetName` | `string` | yes |
| `datasetDescription` | `string` | no |
| `datasetConfig` | `DatasetConfig` | no |
| `metadata` | `Map<string>` | no |
| `datasetSource` | `DatasetSource` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | no |
| `datasetArn` | `string` | no |
| `datasetStatus` | `DatasetStatus` | no |

## UpdateGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `gatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGatewayCapabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `capabilityNamespace` | `string` | yes |
| `capabilityConfiguration` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityNamespace` | `string` | yes |
| `capabilitySyncStatus` | `string` | yes |

## UpdatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `pipelineName` | `string` | yes |
| `description` | `string` | no |
| `environmentVariables` | `Map<string>` | no |
| `computations` | `List<ComputeNode>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | yes |
| `status` | `ResourceStatus` | yes |

## UpdatePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `portalName` | `string` | yes |
| `portalDescription` | `string` | no |
| `portalContactEmail` | `string` | yes |
| `portalLogoImage` | `Image` | no |
| `roleArn` | `string` | yes |
| `clientToken` | `string` | no |
| `notificationSenderEmail` | `string` | no |
| `alarms` | `Alarms` | no |
| `portalType` | `string` | no |
| `portalTypeConfiguration` | `Map<PortalTypeEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalStatus` | `PortalStatus` | yes |

## UpdateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `projectName` | `string` | yes |
| `projectDescription` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `taskName` | `string` | yes |
| `description` | `string` | no |
| `taskConfiguration` | `TaskConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | yes |
| `status` | `ResourceStatus` | yes |

## UpdateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceName` | `string` | yes |
| `workspaceDescription` | `string` | no |
| `encryptionConfiguration` | `WorkspaceEncryptionConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceStatus` | `WorkspaceStatus` | yes |

