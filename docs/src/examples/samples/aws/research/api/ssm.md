# Amazon Simple Systems Manager (SSM)

API version: 2014-11-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ssm/2014-11-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTagsToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `ResourceId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateOpsItemRelatedItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsItemId` | `string` | yes |
| `AssociationType` | `string` | yes |
| `ResourceType` | `string` | yes |
| `ResourceUri` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |

## CancelCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CommandId` | `string` | yes |
| `InstanceIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelMaintenanceWindowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | no |

## CreateActivation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DefaultInstanceName` | `string` | no |
| `IamRole` | `string` | yes |
| `RegistrationLimit` | `integer` | no |
| `ExpirationDate` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |
| `RegistrationMetadata` | `List<RegistrationMetadataItem>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivationId` | `string` | no |
| `ActivationCode` | `string` | no |

## CreateAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `InstanceId` | `string` | no |
| `Parameters` | `Map<List<string>>` | no |
| `Targets` | `List<Target>` | no |
| `ScheduleExpression` | `string` | no |
| `OutputLocation` | `InstanceAssociationOutputLocation` | no |
| `AssociationName` | `string` | no |
| `AutomationTargetParameterName` | `string` | no |
| `MaxErrors` | `string` | no |
| `MaxConcurrency` | `string` | no |
| `ComplianceSeverity` | `string` | no |
| `SyncCompliance` | `string` | no |
| `ApplyOnlyAtCronInterval` | `boolean` | no |
| `CalendarNames` | `List<string>` | no |
| `TargetLocations` | `List<TargetLocation>` | no |
| `ScheduleOffset` | `integer` | no |
| `Duration` | `integer` | no |
| `TargetMaps` | `List<Map<List<string>>>` | no |
| `Tags` | `List<Tag>` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |
| `AssociationDispatchAssumeRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationDescription` | `AssociationDescription` | no |

## CreateAssociationBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<CreateAssociationBatchRequestEntry>` | yes |
| `AssociationDispatchAssumeRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<AssociationDescription>` | no |
| `Failed` | `List<FailedCreateAssociation>` | no |

## CreateCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `Description` | `string` | no |
| `Configuration` | `CloudConnectorConfiguration` | yes |
| `ConfigConnectorArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorId` | `string` | no |

## CreateDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `string` | yes |
| `Requires` | `List<DocumentRequires>` | no |
| `Attachments` | `List<AttachmentsSource>` | no |
| `Name` | `string` | yes |
| `DisplayName` | `string` | no |
| `VersionName` | `string` | no |
| `DocumentType` | `string` | no |
| `DocumentFormat` | `string` | no |
| `TargetType` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentDescription` | `DocumentDescription` | no |

## CreateMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `StartDate` | `string` | no |
| `EndDate` | `string` | no |
| `Schedule` | `string` | yes |
| `ScheduleTimezone` | `string` | no |
| `ScheduleOffset` | `integer` | no |
| `Duration` | `integer` | yes |
| `Cutoff` | `integer` | yes |
| `AllowUnassociatedTargets` | `boolean` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |

## CreateOpsItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | yes |
| `OpsItemType` | `string` | no |
| `OperationalData` | `Map<OpsItemDataValue>` | no |
| `Notifications` | `List<OpsItemNotification>` | no |
| `Priority` | `integer` | no |
| `RelatedOpsItems` | `List<RelatedOpsItem>` | no |
| `Source` | `string` | yes |
| `Title` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `Category` | `string` | no |
| `Severity` | `string` | no |
| `ActualStartTime` | `timestamp` | no |
| `ActualEndTime` | `timestamp` | no |
| `PlannedStartTime` | `timestamp` | no |
| `PlannedEndTime` | `timestamp` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsItemId` | `string` | no |
| `OpsItemArn` | `string` | no |

## CreateOpsMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Metadata` | `Map<MetadataValue>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsMetadataArn` | `string` | no |

## CreatePatchBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperatingSystem` | `string` | no |
| `Name` | `string` | yes |
| `GlobalFilters` | `PatchFilterGroup` | no |
| `ApprovalRules` | `PatchRuleGroup` | no |
| `ApprovedPatches` | `List<string>` | no |
| `ApprovedPatchesComplianceLevel` | `string` | no |
| `ApprovedPatchesEnableNonSecurity` | `boolean` | no |
| `RejectedPatches` | `List<string>` | no |
| `RejectedPatchesAction` | `string` | no |
| `Description` | `string` | no |
| `Sources` | `List<PatchSource>` | no |
| `AvailableSecurityUpdatesComplianceStatus` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |

## CreateResourceDataSync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncName` | `string` | yes |
| `S3Destination` | `ResourceDataSyncS3Destination` | no |
| `SyncType` | `string` | no |
| `SyncSource` | `ResourceDataSyncSource` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteActivation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `InstanceId` | `string` | no |
| `AssociationId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorId` | `string` | no |

## DeleteDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `VersionName` | `string` | no |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInventory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | yes |
| `SchemaDeleteOption` | `string` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionId` | `string` | no |
| `TypeName` | `string` | no |
| `DeletionSummary` | `InventoryDeletionSummary` | no |

## DeleteMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |

## DeleteOpsItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsItemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOpsMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsMetadataArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteParameter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletedParameters` | `List<string>` | no |
| `InvalidParameters` | `List<string>` | no |

## DeletePatchBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |

## DeleteResourceDataSync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncName` | `string` | yes |
| `SyncType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `PolicyId` | `string` | yes |
| `PolicyHash` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterManagedInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterPatchBaselineForPatchGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | yes |
| `PatchGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |
| `PatchGroup` | `string` | no |

## DeregisterTargetFromMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `WindowTargetId` | `string` | yes |
| `Safe` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |
| `WindowTargetId` | `string` | no |

## DeregisterTaskFromMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `WindowTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |
| `WindowTaskId` | `string` | no |

## DescribeActivations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<DescribeActivationsFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivationList` | `List<Activation>` | no |
| `NextToken` | `string` | no |

## DescribeAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `InstanceId` | `string` | no |
| `AssociationId` | `string` | no |
| `AssociationVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationDescription` | `AssociationDescription` | no |

## DescribeAssociationExecutionTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |
| `ExecutionId` | `string` | yes |
| `Filters` | `List<AssociationExecutionTargetsFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationExecutionTargets` | `List<AssociationExecutionTarget>` | no |
| `NextToken` | `string` | no |

## DescribeAssociationExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |
| `Filters` | `List<AssociationExecutionFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationExecutions` | `List<AssociationExecution>` | no |
| `NextToken` | `string` | no |

## DescribeAutomationExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<AutomationExecutionFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationExecutionMetadataList` | `List<AutomationExecutionMetadata>` | no |
| `NextToken` | `string` | no |

## DescribeAutomationStepExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationExecutionId` | `string` | yes |
| `Filters` | `List<StepExecutionFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ReverseOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StepExecutions` | `List<StepExecution>` | no |
| `NextToken` | `string` | no |

## DescribeAvailablePatches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<PatchOrchestratorFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Patches` | `List<Patch>` | no |
| `NextToken` | `string` | no |

## DescribeDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `VersionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Document` | `DocumentDescription` | no |

## DescribeDocumentPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `PermissionType` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | no |
| `AccountSharingInfoList` | `List<AccountSharingInfo>` | no |
| `NextToken` | `string` | no |

## DescribeEffectiveInstanceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<InstanceAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeEffectivePatchesForPatchBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EffectivePatches` | `List<EffectivePatch>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceAssociationsStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceAssociationStatusInfos` | `List<InstanceAssociationStatusInfo>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceInformationFilterList` | `List<InstanceInformationFilter>` | no |
| `Filters` | `List<InstanceInformationStringFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceInformationList` | `List<InstanceInformation>` | no |
| `NextToken` | `string` | no |

## DescribeInstancePatchStates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstancePatchStates` | `List<InstancePatchState>` | no |
| `NextToken` | `string` | no |

## DescribeInstancePatchStatesForPatchGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PatchGroup` | `string` | yes |
| `Filters` | `List<InstancePatchStateFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstancePatchStates` | `List<InstancePatchState>` | no |
| `NextToken` | `string` | no |

## DescribeInstancePatches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Filters` | `List<PatchOrchestratorFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Patches` | `List<PatchComplianceData>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstancePropertyFilterList` | `List<InstancePropertyFilter>` | no |
| `FiltersWithOperator` | `List<InstancePropertyStringFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProperties` | `List<InstanceProperty>` | no |
| `NextToken` | `string` | no |

## DescribeInventoryDeletions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InventoryDeletions` | `List<InventoryDeletionStatusItem>` | no |
| `NextToken` | `string` | no |

## DescribeMaintenanceWindowExecutionTaskInvocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | yes |
| `TaskId` | `string` | yes |
| `Filters` | `List<MaintenanceWindowFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionTaskInvocationIdentities` | `List<MaintenanceWindowExecutionTaskInvocationIdentity>` | no |
| `NextToken` | `string` | no |

## DescribeMaintenanceWindowExecutionTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | yes |
| `Filters` | `List<MaintenanceWindowFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionTaskIdentities` | `List<MaintenanceWindowExecutionTaskIdentity>` | no |
| `NextToken` | `string` | no |

## DescribeMaintenanceWindowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `Filters` | `List<MaintenanceWindowFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutions` | `List<MaintenanceWindowExecution>` | no |
| `NextToken` | `string` | no |

## DescribeMaintenanceWindowSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |
| `Targets` | `List<Target>` | no |
| `ResourceType` | `string` | no |
| `Filters` | `List<PatchOrchestratorFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledWindowExecutions` | `List<ScheduledWindowExecution>` | no |
| `NextToken` | `string` | no |

## DescribeMaintenanceWindowTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `Filters` | `List<MaintenanceWindowFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Targets` | `List<MaintenanceWindowTarget>` | no |
| `NextToken` | `string` | no |

## DescribeMaintenanceWindowTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `Filters` | `List<MaintenanceWindowFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tasks` | `List<MaintenanceWindowTask>` | no |
| `NextToken` | `string` | no |

## DescribeMaintenanceWindows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<MaintenanceWindowFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowIdentities` | `List<MaintenanceWindowIdentity>` | no |
| `NextToken` | `string` | no |

## DescribeMaintenanceWindowsForTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Targets` | `List<Target>` | yes |
| `ResourceType` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowIdentities` | `List<MaintenanceWindowIdentityForTarget>` | no |
| `NextToken` | `string` | no |

## DescribeOpsItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsItemFilters` | `List<OpsItemFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `OpsItemSummaries` | `List<OpsItemSummary>` | no |

## DescribeParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ParametersFilter>` | no |
| `ParameterFilters` | `List<ParameterStringFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Shared` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<ParameterMetadata>` | no |
| `NextToken` | `string` | no |

## DescribePatchBaselines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<PatchOrchestratorFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineIdentities` | `List<PatchBaselineIdentity>` | no |
| `NextToken` | `string` | no |

## DescribePatchGroupState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PatchGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `integer` | no |
| `InstancesWithInstalledPatches` | `integer` | no |
| `InstancesWithInstalledOtherPatches` | `integer` | no |
| `InstancesWithInstalledPendingRebootPatches` | `integer` | no |
| `InstancesWithInstalledRejectedPatches` | `integer` | no |
| `InstancesWithMissingPatches` | `integer` | no |
| `InstancesWithFailedPatches` | `integer` | no |
| `InstancesWithNotApplicablePatches` | `integer` | no |
| `InstancesWithUnreportedNotApplicablePatches` | `integer` | no |
| `InstancesWithCriticalNonCompliantPatches` | `integer` | no |
| `InstancesWithSecurityNonCompliantPatches` | `integer` | no |
| `InstancesWithOtherNonCompliantPatches` | `integer` | no |
| `InstancesWithAvailableSecurityUpdates` | `integer` | no |

## DescribePatchGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `Filters` | `List<PatchOrchestratorFilter>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Mappings` | `List<PatchGroupPatchBaselineMapping>` | no |
| `NextToken` | `string` | no |

## DescribePatchProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperatingSystem` | `string` | yes |
| `Property` | `string` | yes |
| `PatchSet` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Properties` | `List<Map<string>>` | no |
| `NextToken` | `string` | no |

## DescribeSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<SessionFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sessions` | `List<Session>` | no |
| `NextToken` | `string` | no |

## DisassociateOpsItemRelatedItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsItemId` | `string` | yes |
| `AssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `AccessRequestStatus` | `string` | no |

## GetAutomationExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationExecution` | `AutomationExecution` | no |

## GetCalendarState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalendarNames` | `List<string>` | yes |
| `AtTime` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |
| `AtTime` | `string` | no |
| `NextTransitionTime` | `string` | no |

## GetCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorArn` | `string` | no |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |
| `Configuration` | `CloudConnectorConfiguration` | no |
| `ConfigConnectorArn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |

## GetCommandInvocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CommandId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `PluginName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CommandId` | `string` | no |
| `InstanceId` | `string` | no |
| `Comment` | `string` | no |
| `DocumentName` | `string` | no |
| `DocumentVersion` | `string` | no |
| `PluginName` | `string` | no |
| `ResponseCode` | `integer` | no |
| `ExecutionStartDateTime` | `string` | no |
| `ExecutionElapsedTime` | `string` | no |
| `ExecutionEndDateTime` | `string` | no |
| `Status` | `string` | no |
| `StatusDetails` | `string` | no |
| `StandardOutputContent` | `string` | no |
| `StandardOutputUrl` | `string` | no |
| `StandardErrorContent` | `string` | no |
| `StandardErrorUrl` | `string` | no |
| `CloudWatchOutputConfig` | `CloudWatchOutputConfig` | no |

## GetConnectionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Target` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Target` | `string` | no |
| `Status` | `string` | no |

## GetDefaultPatchBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperatingSystem` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |
| `OperatingSystem` | `string` | no |

## GetDeployablePatchSnapshotForInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `SnapshotId` | `string` | yes |
| `BaselineOverride` | `BaselineOverride` | no |
| `UseS3DualStackEndpoint` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `SnapshotId` | `string` | no |
| `SnapshotDownloadUrl` | `string` | no |
| `Product` | `string` | no |

## GetDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `VersionName` | `string` | no |
| `DocumentVersion` | `string` | no |
| `DocumentFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `DisplayName` | `string` | no |
| `VersionName` | `string` | no |
| `DocumentVersion` | `string` | no |
| `Status` | `string` | no |
| `StatusInformation` | `string` | no |
| `Content` | `string` | no |
| `DocumentType` | `string` | no |
| `DocumentFormat` | `string` | no |
| `Requires` | `List<DocumentRequires>` | no |
| `AttachmentsContent` | `List<AttachmentContent>` | no |
| `ReviewStatus` | `string` | no |

## GetExecutionPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionPreviewId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionPreviewId` | `string` | no |
| `EndedAt` | `timestamp` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `ExecutionPreview` | `ExecutionPreview` | no |

## GetInventory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<InventoryFilter>` | no |
| `Aggregators` | `List<InventoryAggregator>` | no |
| `ResultAttributes` | `List<ResultAttribute>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<InventoryResultEntity>` | no |
| `NextToken` | `string` | no |

## GetInventorySchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Aggregator` | `boolean` | no |
| `SubType` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Schemas` | `List<InventoryItemSchema>` | no |
| `NextToken` | `string` | no |

## GetMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `StartDate` | `string` | no |
| `EndDate` | `string` | no |
| `Schedule` | `string` | no |
| `ScheduleTimezone` | `string` | no |
| `ScheduleOffset` | `integer` | no |
| `NextExecutionTime` | `string` | no |
| `Duration` | `integer` | no |
| `Cutoff` | `integer` | no |
| `AllowUnassociatedTargets` | `boolean` | no |
| `Enabled` | `boolean` | no |
| `CreatedDate` | `timestamp` | no |
| `ModifiedDate` | `timestamp` | no |

## GetMaintenanceWindowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | no |
| `TaskIds` | `List<string>` | no |
| `Status` | `string` | no |
| `StatusDetails` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |

## GetMaintenanceWindowExecutionTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | yes |
| `TaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | no |
| `TaskExecutionId` | `string` | no |
| `TaskArn` | `string` | no |
| `ServiceRole` | `string` | no |
| `Type` | `string` | no |
| `TaskParameters` | `List<Map<MaintenanceWindowTaskParameterValueExpression>>` | no |
| `Priority` | `integer` | no |
| `MaxConcurrency` | `string` | no |
| `MaxErrors` | `string` | no |
| `Status` | `string` | no |
| `StatusDetails` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |
| `TriggeredAlarms` | `List<AlarmStateInformation>` | no |

## GetMaintenanceWindowExecutionTaskInvocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | yes |
| `TaskId` | `string` | yes |
| `InvocationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowExecutionId` | `string` | no |
| `TaskExecutionId` | `string` | no |
| `InvocationId` | `string` | no |
| `ExecutionId` | `string` | no |
| `TaskType` | `string` | no |
| `Parameters` | `string` | no |
| `Status` | `string` | no |
| `StatusDetails` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `OwnerInformation` | `string` | no |
| `WindowTargetId` | `string` | no |

## GetMaintenanceWindowTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `WindowTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |
| `WindowTaskId` | `string` | no |
| `Targets` | `List<Target>` | no |
| `TaskArn` | `string` | no |
| `ServiceRoleArn` | `string` | no |
| `TaskType` | `string` | no |
| `TaskParameters` | `Map<MaintenanceWindowTaskParameterValueExpression>` | no |
| `TaskInvocationParameters` | `MaintenanceWindowTaskInvocationParameters` | no |
| `Priority` | `integer` | no |
| `MaxConcurrency` | `string` | no |
| `MaxErrors` | `string` | no |
| `LoggingInfo` | `LoggingInfo` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CutoffBehavior` | `string` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |

## GetOpsItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsItemId` | `string` | yes |
| `OpsItemArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsItem` | `OpsItem` | no |

## GetOpsMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsMetadataArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | no |
| `Metadata` | `Map<MetadataValue>` | no |
| `NextToken` | `string` | no |

## GetOpsSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncName` | `string` | no |
| `Filters` | `List<OpsFilter>` | no |
| `Aggregators` | `List<OpsAggregator>` | no |
| `ResultAttributes` | `List<OpsResultAttribute>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<OpsEntity>` | no |
| `NextToken` | `string` | no |

## GetParameter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `WithDecryption` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameter` | `Parameter` | no |

## GetParameterHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `WithDecryption` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<ParameterHistory>` | no |
| `NextToken` | `string` | no |

## GetParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | yes |
| `WithDecryption` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<Parameter>` | no |
| `InvalidParameters` | `List<string>` | no |

## GetParametersByPath

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | yes |
| `Recursive` | `boolean` | no |
| `ParameterFilters` | `List<ParameterStringFilter>` | no |
| `WithDecryption` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<Parameter>` | no |
| `NextToken` | `string` | no |

## GetPatchBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |
| `Name` | `string` | no |
| `OperatingSystem` | `string` | no |
| `GlobalFilters` | `PatchFilterGroup` | no |
| `ApprovalRules` | `PatchRuleGroup` | no |
| `ApprovedPatches` | `List<string>` | no |
| `ApprovedPatchesComplianceLevel` | `string` | no |
| `ApprovedPatchesEnableNonSecurity` | `boolean` | no |
| `RejectedPatches` | `List<string>` | no |
| `RejectedPatchesAction` | `string` | no |
| `PatchGroups` | `List<string>` | no |
| `CreatedDate` | `timestamp` | no |
| `ModifiedDate` | `timestamp` | no |
| `Description` | `string` | no |
| `Sources` | `List<PatchSource>` | no |
| `AvailableSecurityUpdatesComplianceStatus` | `string` | no |

## GetPatchBaselineForPatchGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PatchGroup` | `string` | yes |
| `OperatingSystem` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |
| `PatchGroup` | `string` | no |
| `OperatingSystem` | `string` | no |

## GetResourcePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Policies` | `List<GetResourcePoliciesResponseEntry>` | no |

## GetServiceSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SettingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSetting` | `ServiceSetting` | no |

## LabelParameterVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ParameterVersion` | `long` | no |
| `Labels` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvalidLabels` | `List<string>` | no |
| `ParameterVersion` | `long` | no |

## ListAssociationVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationVersions` | `List<AssociationVersionInfo>` | no |
| `NextToken` | `string` | no |

## ListAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationFilterList` | `List<AssociationFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<Association>` | no |
| `NextToken` | `string` | no |

## ListCloudConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<CloudConnectorFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectors` | `List<CloudConnectorSummary>` | no |
| `NextToken` | `string` | no |

## ListCommandInvocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CommandId` | `string` | no |
| `InstanceId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<CommandFilter>` | no |
| `Details` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CommandInvocations` | `List<CommandInvocation>` | no |
| `NextToken` | `string` | no |

## ListCommands

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CommandId` | `string` | no |
| `InstanceId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<CommandFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Commands` | `List<Command>` | no |
| `NextToken` | `string` | no |

## ListComplianceItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ComplianceStringFilter>` | no |
| `ResourceIds` | `List<string>` | no |
| `ResourceTypes` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComplianceItems` | `List<ComplianceItem>` | no |
| `NextToken` | `string` | no |

## ListComplianceSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ComplianceStringFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComplianceSummaryItems` | `List<ComplianceSummaryItem>` | no |
| `NextToken` | `string` | no |

## ListDocumentMetadataHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `Metadata` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `DocumentVersion` | `string` | no |
| `Author` | `string` | no |
| `Metadata` | `DocumentMetadataResponseInfo` | no |
| `NextToken` | `string` | no |

## ListDocumentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentVersions` | `List<DocumentVersionInfo>` | no |
| `NextToken` | `string` | no |

## ListDocuments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentFilterList` | `List<DocumentFilter>` | no |
| `Filters` | `List<DocumentKeyValuesFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentIdentifiers` | `List<DocumentIdentifier>` | no |
| `NextToken` | `string` | no |

## ListInventoryEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TypeName` | `string` | yes |
| `Filters` | `List<InventoryFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | no |
| `InstanceId` | `string` | no |
| `SchemaVersion` | `string` | no |
| `CaptureTime` | `string` | no |
| `Entries` | `List<Map<string>>` | no |
| `NextToken` | `string` | no |

## ListNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncName` | `string` | no |
| `Filters` | `List<NodeFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Nodes` | `List<Node>` | no |
| `NextToken` | `string` | no |

## ListNodesSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncName` | `string` | no |
| `Filters` | `List<NodeFilter>` | no |
| `Aggregators` | `List<NodeAggregator>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summary` | `List<Map<string>>` | no |
| `NextToken` | `string` | no |

## ListOpsItemEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<OpsItemEventFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Summaries` | `List<OpsItemEventSummary>` | no |

## ListOpsItemRelatedItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsItemId` | `string` | no |
| `Filters` | `List<OpsItemRelatedItemsFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Summaries` | `List<OpsItemRelatedItemSummary>` | no |

## ListOpsMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<OpsMetadataFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsMetadataList` | `List<OpsMetadata>` | no |
| `NextToken` | `string` | no |

## ListResourceComplianceSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ComplianceStringFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceComplianceSummaryItems` | `List<ResourceComplianceSummaryItem>` | no |
| `NextToken` | `string` | no |

## ListResourceDataSync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceDataSyncItems` | `List<ResourceDataSyncItem>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## ModifyDocumentPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `PermissionType` | `string` | yes |
| `AccountIdsToAdd` | `List<string>` | no |
| `AccountIdsToRemove` | `List<string>` | no |
| `SharedDocumentVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutComplianceItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |
| `ComplianceType` | `string` | yes |
| `ExecutionSummary` | `ComplianceExecutionSummary` | yes |
| `Items` | `List<ComplianceItemEntry>` | yes |
| `ItemContentHash` | `string` | no |
| `UploadType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutInventory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Items` | `List<InventoryItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## PutParameter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Value` | `string` | yes |
| `Type` | `string` | no |
| `KeyId` | `string` | no |
| `Overwrite` | `boolean` | no |
| `AllowedPattern` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `Tier` | `string` | no |
| `Policies` | `string` | no |
| `DataType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Version` | `long` | no |
| `Tier` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |
| `PolicyId` | `string` | no |
| `PolicyHash` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | no |
| `PolicyHash` | `string` | no |

## RegisterDefaultPatchBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |

## RegisterPatchBaselineForPatchGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | yes |
| `PatchGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |
| `PatchGroup` | `string` | no |

## RegisterTargetWithMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `ResourceType` | `string` | yes |
| `Targets` | `List<Target>` | yes |
| `OwnerInformation` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowTargetId` | `string` | no |

## RegisterTaskWithMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `Targets` | `List<Target>` | no |
| `TaskArn` | `string` | yes |
| `ServiceRoleArn` | `string` | no |
| `TaskType` | `string` | yes |
| `TaskParameters` | `Map<MaintenanceWindowTaskParameterValueExpression>` | no |
| `TaskInvocationParameters` | `MaintenanceWindowTaskInvocationParameters` | no |
| `Priority` | `integer` | no |
| `MaxConcurrency` | `string` | no |
| `MaxErrors` | `string` | no |
| `LoggingInfo` | `LoggingInfo` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `CutoffBehavior` | `string` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowTaskId` | `string` | no |

## RemoveTagsFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `ResourceId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetServiceSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SettingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSetting` | `ServiceSetting` | no |

## ResumeSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | no |
| `TokenValue` | `string` | no |
| `StreamUrl` | `string` | no |

## SendAutomationSignal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationExecutionId` | `string` | yes |
| `SignalType` | `string` | yes |
| `Payload` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `Targets` | `List<Target>` | no |
| `DocumentName` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `DocumentHash` | `string` | no |
| `DocumentHashType` | `string` | no |
| `TimeoutSeconds` | `integer` | no |
| `Comment` | `string` | no |
| `Parameters` | `Map<List<string>>` | no |
| `OutputS3Region` | `string` | no |
| `OutputS3BucketName` | `string` | no |
| `OutputS3KeyPrefix` | `string` | no |
| `MaxConcurrency` | `string` | no |
| `MaxErrors` | `string` | no |
| `ServiceRoleArn` | `string` | no |
| `NotificationConfig` | `NotificationConfig` | no |
| `CloudWatchOutputConfig` | `CloudWatchOutputConfig` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Command` | `Command` | no |

## StartAccessRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Reason` | `string` | yes |
| `Targets` | `List<Target>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessRequestId` | `string` | no |

## StartAssociationsOnce

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAutomationExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentName` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `Parameters` | `Map<List<string>>` | no |
| `ClientToken` | `string` | no |
| `Mode` | `string` | no |
| `TargetParameterName` | `string` | no |
| `Targets` | `List<Target>` | no |
| `TargetMaps` | `List<Map<List<string>>>` | no |
| `MaxConcurrency` | `string` | no |
| `MaxErrors` | `string` | no |
| `TargetLocations` | `List<TargetLocation>` | no |
| `Tags` | `List<Tag>` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |
| `TargetLocationsURL` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationExecutionId` | `string` | no |

## StartChangeRequestExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledTime` | `timestamp` | no |
| `DocumentName` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `Parameters` | `Map<List<string>>` | no |
| `ChangeRequestName` | `string` | no |
| `ClientToken` | `string` | no |
| `AutoApprove` | `boolean` | no |
| `Runbooks` | `List<Runbook>` | yes |
| `Tags` | `List<Tag>` | no |
| `ScheduledEndTime` | `timestamp` | no |
| `ChangeDetails` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationExecutionId` | `string` | no |

## StartExecutionPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentName` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `ExecutionInputs` | `ExecutionInputs` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionPreviewId` | `string` | no |

## StartSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Target` | `string` | yes |
| `DocumentName` | `string` | no |
| `Reason` | `string` | no |
| `Parameters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | no |
| `TokenValue` | `string` | no |
| `StreamUrl` | `string` | no |

## StopAutomationExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationExecutionId` | `string` | yes |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | no |

## UnlabelParameterVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ParameterVersion` | `long` | yes |
| `Labels` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RemovedLabels` | `List<string>` | no |
| `InvalidLabels` | `List<string>` | no |

## UpdateAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |
| `Parameters` | `Map<List<string>>` | no |
| `DocumentVersion` | `string` | no |
| `ScheduleExpression` | `string` | no |
| `OutputLocation` | `InstanceAssociationOutputLocation` | no |
| `Name` | `string` | no |
| `Targets` | `List<Target>` | no |
| `AssociationName` | `string` | no |
| `AssociationVersion` | `string` | no |
| `AutomationTargetParameterName` | `string` | no |
| `MaxErrors` | `string` | no |
| `MaxConcurrency` | `string` | no |
| `ComplianceSeverity` | `string` | no |
| `SyncCompliance` | `string` | no |
| `ApplyOnlyAtCronInterval` | `boolean` | no |
| `CalendarNames` | `List<string>` | no |
| `TargetLocations` | `List<TargetLocation>` | no |
| `ScheduleOffset` | `integer` | no |
| `Duration` | `integer` | no |
| `TargetMaps` | `List<Map<List<string>>>` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |
| `AssociationDispatchAssumeRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationDescription` | `AssociationDescription` | no |

## UpdateAssociationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `InstanceId` | `string` | yes |
| `AssociationStatus` | `AssociationStatus` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationDescription` | `AssociationDescription` | no |

## UpdateCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorId` | `string` | yes |
| `DisplayName` | `string` | no |
| `Configuration` | `CloudConnectorConfiguration` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorId` | `string` | no |

## UpdateDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `string` | yes |
| `Attachments` | `List<AttachmentsSource>` | no |
| `Name` | `string` | yes |
| `DisplayName` | `string` | no |
| `VersionName` | `string` | no |
| `DocumentVersion` | `string` | no |
| `DocumentFormat` | `string` | no |
| `TargetType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentDescription` | `DocumentDescription` | no |

## UpdateDocumentDefaultVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DocumentVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `DocumentDefaultVersionDescription` | no |

## UpdateDocumentMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DocumentVersion` | `string` | no |
| `DocumentReviews` | `DocumentReviews` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `StartDate` | `string` | no |
| `EndDate` | `string` | no |
| `Schedule` | `string` | no |
| `ScheduleTimezone` | `string` | no |
| `ScheduleOffset` | `integer` | no |
| `Duration` | `integer` | no |
| `Cutoff` | `integer` | no |
| `AllowUnassociatedTargets` | `boolean` | no |
| `Enabled` | `boolean` | no |
| `Replace` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `StartDate` | `string` | no |
| `EndDate` | `string` | no |
| `Schedule` | `string` | no |
| `ScheduleTimezone` | `string` | no |
| `ScheduleOffset` | `integer` | no |
| `Duration` | `integer` | no |
| `Cutoff` | `integer` | no |
| `AllowUnassociatedTargets` | `boolean` | no |
| `Enabled` | `boolean` | no |

## UpdateMaintenanceWindowTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `WindowTargetId` | `string` | yes |
| `Targets` | `List<Target>` | no |
| `OwnerInformation` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Replace` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |
| `WindowTargetId` | `string` | no |
| `Targets` | `List<Target>` | no |
| `OwnerInformation` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |

## UpdateMaintenanceWindowTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | yes |
| `WindowTaskId` | `string` | yes |
| `Targets` | `List<Target>` | no |
| `TaskArn` | `string` | no |
| `ServiceRoleArn` | `string` | no |
| `TaskParameters` | `Map<MaintenanceWindowTaskParameterValueExpression>` | no |
| `TaskInvocationParameters` | `MaintenanceWindowTaskInvocationParameters` | no |
| `Priority` | `integer` | no |
| `MaxConcurrency` | `string` | no |
| `MaxErrors` | `string` | no |
| `LoggingInfo` | `LoggingInfo` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Replace` | `boolean` | no |
| `CutoffBehavior` | `string` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WindowId` | `string` | no |
| `WindowTaskId` | `string` | no |
| `Targets` | `List<Target>` | no |
| `TaskArn` | `string` | no |
| `ServiceRoleArn` | `string` | no |
| `TaskParameters` | `Map<MaintenanceWindowTaskParameterValueExpression>` | no |
| `TaskInvocationParameters` | `MaintenanceWindowTaskInvocationParameters` | no |
| `Priority` | `integer` | no |
| `MaxConcurrency` | `string` | no |
| `MaxErrors` | `string` | no |
| `LoggingInfo` | `LoggingInfo` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CutoffBehavior` | `string` | no |
| `AlarmConfiguration` | `AlarmConfiguration` | no |

## UpdateManagedInstanceRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IamRole` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOpsItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `OperationalData` | `Map<OpsItemDataValue>` | no |
| `OperationalDataToDelete` | `List<string>` | no |
| `Notifications` | `List<OpsItemNotification>` | no |
| `Priority` | `integer` | no |
| `RelatedOpsItems` | `List<RelatedOpsItem>` | no |
| `Status` | `string` | no |
| `OpsItemId` | `string` | yes |
| `Title` | `string` | no |
| `Category` | `string` | no |
| `Severity` | `string` | no |
| `ActualStartTime` | `timestamp` | no |
| `ActualEndTime` | `timestamp` | no |
| `PlannedStartTime` | `timestamp` | no |
| `PlannedEndTime` | `timestamp` | no |
| `OpsItemArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOpsMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsMetadataArn` | `string` | yes |
| `MetadataToUpdate` | `Map<MetadataValue>` | no |
| `KeysToDelete` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpsMetadataArn` | `string` | no |

## UpdatePatchBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | yes |
| `Name` | `string` | no |
| `GlobalFilters` | `PatchFilterGroup` | no |
| `ApprovalRules` | `PatchRuleGroup` | no |
| `ApprovedPatches` | `List<string>` | no |
| `ApprovedPatchesComplianceLevel` | `string` | no |
| `ApprovedPatchesEnableNonSecurity` | `boolean` | no |
| `RejectedPatches` | `List<string>` | no |
| `RejectedPatchesAction` | `string` | no |
| `Description` | `string` | no |
| `Sources` | `List<PatchSource>` | no |
| `AvailableSecurityUpdatesComplianceStatus` | `string` | no |
| `Replace` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaselineId` | `string` | no |
| `Name` | `string` | no |
| `OperatingSystem` | `string` | no |
| `GlobalFilters` | `PatchFilterGroup` | no |
| `ApprovalRules` | `PatchRuleGroup` | no |
| `ApprovedPatches` | `List<string>` | no |
| `ApprovedPatchesComplianceLevel` | `string` | no |
| `ApprovedPatchesEnableNonSecurity` | `boolean` | no |
| `RejectedPatches` | `List<string>` | no |
| `RejectedPatchesAction` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `ModifiedDate` | `timestamp` | no |
| `Description` | `string` | no |
| `Sources` | `List<PatchSource>` | no |
| `AvailableSecurityUpdatesComplianceStatus` | `string` | no |

## UpdateResourceDataSync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncName` | `string` | yes |
| `SyncType` | `string` | yes |
| `SyncSource` | `ResourceDataSyncSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateServiceSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SettingId` | `string` | yes |
| `SettingValue` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ValidateCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ValidationFindings` | `List<ValidationFinding>` | no |
| `NextToken` | `string` | no |

