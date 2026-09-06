# AWS Backup

API version: 2018-11-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/backup/2018-11-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateBackupVaultMpaApprovalTeam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `MpaApprovalTeamArn` | `string` | yes |
| `RequesterComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelLegalHold

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LegalHoldId` | `string` | yes |
| `CancelDescription` | `string` | yes |
| `RetainRecordInDays` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateBackupAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointMetadata` | `Map<string>` | no |
| `AccessPointPolicy` | `string` | no |
| `Name` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointArn` | `string` | yes |
| `Status` | `string` | yes |

## CreateBackupPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlan` | `BackupPlanInput` | yes |
| `BackupPlanTags` | `Map<string>` | no |
| `CreatorRequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | no |
| `BackupPlanArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `VersionId` | `string` | no |
| `AdvancedBackupSettings` | `List<AdvancedBackupSetting>` | no |

## CreateBackupSelection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |
| `BackupSelection` | `BackupSelection` | yes |
| `CreatorRequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SelectionId` | `string` | no |
| `BackupPlanId` | `string` | no |
| `CreationDate` | `timestamp` | no |

## CreateBackupVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `BackupVaultTags` | `Map<string>` | no |
| `EncryptionKeyArn` | `string` | no |
| `CreatorRequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | no |
| `BackupVaultArn` | `string` | no |
| `CreationDate` | `timestamp` | no |

## CreateFramework

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FrameworkName` | `string` | yes |
| `FrameworkDescription` | `string` | no |
| `FrameworkControls` | `List<FrameworkControl>` | yes |
| `IdempotencyToken` | `string` | no |
| `FrameworkTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FrameworkName` | `string` | no |
| `FrameworkArn` | `string` | no |

## CreateLegalHold

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Title` | `string` | yes |
| `Description` | `string` | yes |
| `IdempotencyToken` | `string` | no |
| `RecoveryPointSelection` | `RecoveryPointSelection` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Title` | `string` | no |
| `Status` | `string` | no |
| `Description` | `string` | no |
| `LegalHoldId` | `string` | no |
| `LegalHoldArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `RecoveryPointSelection` | `RecoveryPointSelection` | no |

## CreateLogicallyAirGappedBackupVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `BackupVaultTags` | `Map<string>` | no |
| `CreatorRequestId` | `string` | no |
| `MinRetentionDays` | `long` | yes |
| `MaxRetentionDays` | `long` | yes |
| `EncryptionKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | no |
| `BackupVaultArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `VaultState` | `string` | no |

## CreateReportPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlanName` | `string` | yes |
| `ReportPlanDescription` | `string` | no |
| `ReportDeliveryChannel` | `ReportDeliveryChannel` | yes |
| `ReportSetting` | `ReportSetting` | yes |
| `ReportPlanTags` | `Map<string>` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlanName` | `string` | no |
| `ReportPlanArn` | `string` | no |
| `CreationTime` | `timestamp` | no |

## CreateRestoreAccessBackupVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceBackupVaultArn` | `string` | yes |
| `BackupVaultName` | `string` | no |
| `BackupVaultTags` | `Map<string>` | no |
| `CreatorRequestId` | `string` | no |
| `RequesterComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreAccessBackupVaultArn` | `string` | no |
| `VaultState` | `string` | no |
| `RestoreAccessBackupVaultName` | `string` | no |
| `CreationDate` | `timestamp` | no |

## CreateRestoreTestingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | no |
| `RestoreTestingPlan` | `RestoreTestingPlanForCreate` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | yes |
| `RestoreTestingPlanArn` | `string` | yes |
| `RestoreTestingPlanName` | `string` | yes |

## CreateRestoreTestingSelection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | no |
| `RestoreTestingPlanName` | `string` | yes |
| `RestoreTestingSelection` | `RestoreTestingSelectionForCreate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | yes |
| `RestoreTestingPlanArn` | `string` | yes |
| `RestoreTestingPlanName` | `string` | yes |
| `RestoreTestingSelectionName` | `string` | yes |

## CreateTieringConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TieringConfiguration` | `TieringConfigurationInputForCreate` | yes |
| `TieringConfigurationTags` | `Map<string>` | no |
| `CreatorRequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TieringConfigurationArn` | `string` | no |
| `TieringConfigurationName` | `string` | no |
| `CreationTime` | `timestamp` | no |

## DeleteBackupAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBackupPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | no |
| `BackupPlanArn` | `string` | no |
| `DeletionDate` | `timestamp` | no |
| `VersionId` | `string` | no |

## DeleteBackupSelection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |
| `SelectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBackupVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBackupVaultAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBackupVaultLockConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBackupVaultNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFramework

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FrameworkName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReportPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRestoreTestingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreTestingPlanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRestoreTestingSelection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreTestingPlanName` | `string` | yes |
| `RestoreTestingSelectionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTieringConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TieringConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeBackupAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointArn` | `string` | yes |
| `AccessPointMetadata` | `Map<string>` | no |
| `BackupVaultArn` | `string` | no |
| `BackupVaultName` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `Name` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `ResourceType` | `string` | yes |
| `Status` | `string` | yes |
| `StatusMessage` | `string` | no |

## DescribeBackupJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `BackupJobId` | `string` | no |
| `BackupVaultName` | `string` | no |
| `RecoveryPointLifecycle` | `Lifecycle` | no |
| `BackupVaultArn` | `string` | no |
| `VaultType` | `string` | no |
| `VaultLockState` | `string` | no |
| `RecoveryPointArn` | `string` | no |
| `EncryptionKeyArn` | `string` | no |
| `IsEncrypted` | `boolean` | no |
| `ResourceArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `CompletionDate` | `timestamp` | no |
| `State` | `string` | no |
| `StatusMessage` | `string` | no |
| `PercentDone` | `string` | no |
| `BackupSizeInBytes` | `long` | no |
| `IamRoleArn` | `string` | no |
| `CreatedBy` | `RecoveryPointCreator` | no |
| `ResourceType` | `string` | no |
| `BytesTransferred` | `long` | no |
| `ExpectedCompletionDate` | `timestamp` | no |
| `StartBy` | `timestamp` | no |
| `BackupOptions` | `Map<string>` | no |
| `BackupType` | `string` | no |
| `ParentJobId` | `string` | no |
| `IsParent` | `boolean` | no |
| `NumberOfChildJobs` | `long` | no |
| `ChildJobsInState` | `Map<long>` | no |
| `ResourceName` | `string` | no |
| `InitiationDate` | `timestamp` | no |
| `MessageCategory` | `string` | no |

## DescribeBackupVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `BackupVaultAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | no |
| `BackupVaultArn` | `string` | no |
| `VaultType` | `string` | no |
| `VaultState` | `string` | no |
| `EncryptionKeyArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `CreatorRequestId` | `string` | no |
| `NumberOfRecoveryPoints` | `long` | no |
| `Locked` | `boolean` | no |
| `MinRetentionDays` | `long` | no |
| `MaxRetentionDays` | `long` | no |
| `LockDate` | `timestamp` | no |
| `SourceBackupVaultArn` | `string` | no |
| `MpaApprovalTeamArn` | `string` | no |
| `MpaSessionArn` | `string` | no |
| `LatestMpaApprovalTeamUpdate` | `LatestMpaApprovalTeamUpdate` | no |
| `EncryptionKeyType` | `string` | no |

## DescribeCopyJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CopyJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CopyJob` | `CopyJob` | no |

## DescribeFramework

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FrameworkName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FrameworkName` | `string` | no |
| `FrameworkArn` | `string` | no |
| `FrameworkDescription` | `string` | no |
| `FrameworkControls` | `List<FrameworkControl>` | no |
| `CreationTime` | `timestamp` | no |
| `DeploymentStatus` | `string` | no |
| `FrameworkStatus` | `string` | no |
| `IdempotencyToken` | `string` | no |

## DescribeGlobalSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalSettings` | `Map<string>` | no |
| `LastUpdateTime` | `timestamp` | no |

## DescribeProtectedResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `ResourceType` | `string` | no |
| `LastBackupTime` | `timestamp` | no |
| `ResourceName` | `string` | no |
| `LastBackupVaultArn` | `string` | no |
| `LastRecoveryPointArn` | `string` | no |
| `LatestRestoreExecutionTimeMinutes` | `long` | no |
| `LatestRestoreJobCreationDate` | `timestamp` | no |
| `LatestRestoreRecoveryPointCreationDate` | `timestamp` | no |

## DescribeRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |
| `BackupVaultAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecoveryPointArn` | `string` | no |
| `BackupVaultName` | `string` | no |
| `BackupVaultArn` | `string` | no |
| `SourceBackupVaultArn` | `string` | no |
| `ResourceArn` | `string` | no |
| `ResourceType` | `string` | no |
| `CreatedBy` | `RecoveryPointCreator` | no |
| `IamRoleArn` | `string` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `InitiationDate` | `timestamp` | no |
| `CompletionDate` | `timestamp` | no |
| `BackupSizeInBytes` | `long` | no |
| `CalculatedLifecycle` | `CalculatedLifecycle` | no |
| `Lifecycle` | `Lifecycle` | no |
| `EncryptionKeyArn` | `string` | no |
| `IsEncrypted` | `boolean` | no |
| `StorageClass` | `string` | no |
| `LastRestoreTime` | `timestamp` | no |
| `ParentRecoveryPointArn` | `string` | no |
| `CompositeMemberIdentifier` | `string` | no |
| `IsParent` | `boolean` | no |
| `ResourceName` | `string` | no |
| `VaultType` | `string` | no |
| `IndexStatus` | `string` | no |
| `IndexStatusMessage` | `string` | no |
| `EncryptionKeyType` | `string` | no |
| `ScanResults` | `List<ScanResult>` | no |

## DescribeRegionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTypeOptInPreference` | `Map<boolean>` | no |
| `ResourceTypeManagementPreference` | `Map<boolean>` | no |

## DescribeReportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportJob` | `ReportJob` | no |

## DescribeReportPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlan` | `ReportPlan` | no |

## DescribeRestoreJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `RestoreJobId` | `string` | no |
| `RecoveryPointArn` | `string` | no |
| `SourceResourceArn` | `string` | no |
| `BackupVaultArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `CompletionDate` | `timestamp` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `PercentDone` | `string` | no |
| `BackupSizeInBytes` | `long` | no |
| `IamRoleArn` | `string` | no |
| `ExpectedCompletionTimeMinutes` | `long` | no |
| `CreatedResourceArn` | `string` | no |
| `ResourceType` | `string` | no |
| `RecoveryPointCreationDate` | `timestamp` | no |
| `CreatedBy` | `RestoreJobCreator` | no |
| `ValidationStatus` | `string` | no |
| `ValidationStatusMessage` | `string` | no |
| `DeletionStatus` | `string` | no |
| `DeletionStatusMessage` | `string` | no |
| `IsParent` | `boolean` | no |
| `ParentJobId` | `string` | no |

## DescribeScanJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScanJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BackupVaultArn` | `string` | yes |
| `BackupVaultName` | `string` | yes |
| `CompletionDate` | `timestamp` | no |
| `ContinuousScanEndTime` | `timestamp` | no |
| `ContinuousScanStartTime` | `timestamp` | no |
| `CreatedBy` | `ScanJobCreator` | yes |
| `CreationDate` | `timestamp` | yes |
| `IamRoleArn` | `string` | yes |
| `MalwareScanner` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `ResourceName` | `string` | yes |
| `ResourceType` | `string` | yes |
| `ScanBaseRecoveryPointArn` | `string` | no |
| `ScanId` | `string` | no |
| `ScanJobId` | `string` | yes |
| `ScanMode` | `string` | yes |
| `ScanResult` | `ScanResultInfo` | no |
| `ScannerRoleArn` | `string` | yes |
| `State` | `string` | yes |
| `StatusMessage` | `string` | no |

## DisassociateBackupVaultMpaApprovalTeam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RequesterComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateRecoveryPointFromParent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExportBackupPlanTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanTemplateJson` | `string` | no |

## GetBackupPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |
| `VersionId` | `string` | no |
| `MaxScheduledRunsPreview` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlan` | `BackupPlan` | no |
| `BackupPlanId` | `string` | no |
| `BackupPlanArn` | `string` | no |
| `VersionId` | `string` | no |
| `CreatorRequestId` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `DeletionDate` | `timestamp` | no |
| `LastExecutionDate` | `timestamp` | no |
| `AdvancedBackupSettings` | `List<AdvancedBackupSetting>` | no |
| `ScheduledRunsPreview` | `List<ScheduledPlanExecutionMember>` | no |

## GetBackupPlanFromJSON

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanTemplateJson` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlan` | `BackupPlan` | no |

## GetBackupPlanFromTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanTemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanDocument` | `BackupPlan` | no |

## GetBackupSelection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |
| `SelectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupSelection` | `BackupSelection` | no |
| `SelectionId` | `string` | no |
| `BackupPlanId` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `CreatorRequestId` | `string` | no |

## GetBackupVaultAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | no |
| `BackupVaultArn` | `string` | no |
| `Policy` | `string` | no |

## GetBackupVaultNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | no |
| `BackupVaultArn` | `string` | no |
| `SNSTopicArn` | `string` | no |
| `BackupVaultEvents` | `List<string>` | no |

## GetLegalHold

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LegalHoldId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Title` | `string` | no |
| `Status` | `string` | no |
| `Description` | `string` | no |
| `CancelDescription` | `string` | no |
| `LegalHoldId` | `string` | no |
| `LegalHoldArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `CancellationDate` | `timestamp` | no |
| `RetainRecordUntil` | `timestamp` | no |
| `RecoveryPointSelection` | `RecoveryPointSelection` | no |

## GetPITRMalwareScanResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecoveryPointArn` | `string` | yes |
| `BackupVaultName` | `string` | yes |
| `ScanEndTime` | `timestamp` | yes |
| `MalwareScanner` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScanEndTime` | `timestamp` | yes |
| `ScanResult` | `ScanResultInfo` | yes |
| `LastScanJobTime` | `timestamp` | no |
| `ScanId` | `string` | no |
| `ScanMode` | `string` | no |

## GetRecoveryPointIndexDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecoveryPointArn` | `string` | no |
| `BackupVaultArn` | `string` | no |
| `SourceResourceArn` | `string` | no |
| `IndexCreationDate` | `timestamp` | no |
| `IndexDeletionDate` | `timestamp` | no |
| `IndexCompletionDate` | `timestamp` | no |
| `IndexStatus` | `string` | no |
| `IndexStatusMessage` | `string` | no |
| `TotalItemsIndexed` | `long` | no |

## GetRecoveryPointRestoreMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |
| `BackupVaultAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultArn` | `string` | no |
| `RecoveryPointArn` | `string` | no |
| `RestoreMetadata` | `Map<string>` | no |
| `ResourceType` | `string` | no |

## GetRestoreJobMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreJobId` | `string` | no |
| `Metadata` | `Map<string>` | no |

## GetRestoreTestingInferredMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultAccountId` | `string` | no |
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferredMetadata` | `Map<string>` | yes |

## GetRestoreTestingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreTestingPlanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreTestingPlan` | `RestoreTestingPlanForGet` | yes |

## GetRestoreTestingSelection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreTestingPlanName` | `string` | yes |
| `RestoreTestingSelectionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreTestingSelection` | `RestoreTestingSelectionForGet` | yes |

## GetSupportedResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTypes` | `List<string>` | no |

## GetTieringConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TieringConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TieringConfiguration` | `TieringConfiguration` | no |

## ListBackupAccessPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupAccessPoints` | `List<ListAccessPointsMember>` | yes |
| `NextToken` | `string` | no |

## ListBackupAccessPointsByRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RecoveryPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupAccessPoints` | `List<ListAccessPointsMember>` | yes |
| `NextToken` | `string` | no |

## ListBackupAccessPointsByResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupAccessPoints` | `List<ListAccessPointsMember>` | yes |
| `NextToken` | `string` | no |

## ListBackupJobSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `State` | `string` | no |
| `ResourceType` | `string` | no |
| `MessageCategory` | `string` | no |
| `AggregationPeriod` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupJobSummaries` | `List<BackupJobSummary>` | no |
| `AggregationPeriod` | `string` | no |
| `NextToken` | `string` | no |

## ListBackupJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ByResourceArn` | `string` | no |
| `ByState` | `string` | no |
| `ByBackupVaultName` | `string` | no |
| `ByCreatedBefore` | `timestamp` | no |
| `ByCreatedAfter` | `timestamp` | no |
| `ByResourceType` | `string` | no |
| `ByAccountId` | `string` | no |
| `ByCompleteAfter` | `timestamp` | no |
| `ByCompleteBefore` | `timestamp` | no |
| `ByParentJobId` | `string` | no |
| `ByMessageCategory` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupJobs` | `List<BackupJob>` | no |
| `NextToken` | `string` | no |

## ListBackupPlanTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `BackupPlanTemplatesList` | `List<BackupPlanTemplatesListMember>` | no |

## ListBackupPlanVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `BackupPlanVersionsList` | `List<BackupPlansListMember>` | no |

## ListBackupPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IncludeDeleted` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `BackupPlansList` | `List<BackupPlansListMember>` | no |

## ListBackupSelections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `BackupSelectionsList` | `List<BackupSelectionsListMember>` | no |

## ListBackupVaults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByVaultType` | `string` | no |
| `ByShared` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultList` | `List<BackupVaultListMember>` | no |
| `NextToken` | `string` | no |

## ListCopyJobSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `State` | `string` | no |
| `ResourceType` | `string` | no |
| `MessageCategory` | `string` | no |
| `AggregationPeriod` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CopyJobSummaries` | `List<CopyJobSummary>` | no |
| `AggregationPeriod` | `string` | no |
| `NextToken` | `string` | no |

## ListCopyJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ByResourceArn` | `string` | no |
| `ByState` | `string` | no |
| `ByCreatedBefore` | `timestamp` | no |
| `ByCreatedAfter` | `timestamp` | no |
| `ByResourceType` | `string` | no |
| `ByDestinationVaultArn` | `string` | no |
| `ByAccountId` | `string` | no |
| `ByCompleteBefore` | `timestamp` | no |
| `ByCompleteAfter` | `timestamp` | no |
| `ByParentJobId` | `string` | no |
| `ByMessageCategory` | `string` | no |
| `BySourceRecoveryPointArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CopyJobs` | `List<CopyJob>` | no |
| `NextToken` | `string` | no |

## ListFrameworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Frameworks` | `List<Framework>` | no |
| `NextToken` | `string` | no |

## ListIndexedRecoveryPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SourceResourceArn` | `string` | no |
| `CreatedBefore` | `timestamp` | no |
| `CreatedAfter` | `timestamp` | no |
| `ResourceType` | `string` | no |
| `IndexStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexedRecoveryPoints` | `List<IndexedRecoveryPoint>` | no |
| `NextToken` | `string` | no |

## ListLegalHolds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `LegalHolds` | `List<LegalHold>` | no |

## ListProtectedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<ProtectedResource>` | no |
| `NextToken` | `string` | no |

## ListProtectedResourcesByBackupVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `BackupVaultAccountId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<ProtectedResource>` | no |
| `NextToken` | `string` | no |

## ListRecoveryPointsByBackupVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `BackupVaultAccountId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ByResourceArn` | `string` | no |
| `ByResourceType` | `string` | no |
| `ByBackupPlanId` | `string` | no |
| `ByCreatedBefore` | `timestamp` | no |
| `ByCreatedAfter` | `timestamp` | no |
| `ByParentRecoveryPointArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RecoveryPoints` | `List<RecoveryPointByBackupVault>` | no |

## ListRecoveryPointsByLegalHold

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LegalHoldId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecoveryPoints` | `List<RecoveryPointMember>` | no |
| `NextToken` | `string` | no |

## ListRecoveryPointsByResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ManagedByAWSBackupOnly` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RecoveryPoints` | `List<RecoveryPointByResource>` | no |

## ListReportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByReportPlanName` | `string` | no |
| `ByCreationBefore` | `timestamp` | no |
| `ByCreationAfter` | `timestamp` | no |
| `ByStatus` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportJobs` | `List<ReportJob>` | no |
| `NextToken` | `string` | no |

## ListReportPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlans` | `List<ReportPlan>` | no |
| `NextToken` | `string` | no |

## ListRestoreAccessBackupVaults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RestoreAccessBackupVaults` | `List<RestoreAccessBackupVaultListMember>` | no |

## ListRestoreJobSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `State` | `string` | no |
| `ResourceType` | `string` | no |
| `AggregationPeriod` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreJobSummaries` | `List<RestoreJobSummary>` | no |
| `AggregationPeriod` | `string` | no |
| `NextToken` | `string` | no |

## ListRestoreJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ByAccountId` | `string` | no |
| `ByResourceType` | `string` | no |
| `ByCreatedBefore` | `timestamp` | no |
| `ByCreatedAfter` | `timestamp` | no |
| `ByStatus` | `string` | no |
| `ByCompleteBefore` | `timestamp` | no |
| `ByCompleteAfter` | `timestamp` | no |
| `ByRestoreTestingPlanArn` | `string` | no |
| `ByParentJobId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreJobs` | `List<RestoreJobsListMember>` | no |
| `NextToken` | `string` | no |

## ListRestoreJobsByProtectedResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ByStatus` | `string` | no |
| `ByRecoveryPointCreationDateAfter` | `timestamp` | no |
| `ByRecoveryPointCreationDateBefore` | `timestamp` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreJobs` | `List<RestoreJobsListMember>` | no |
| `NextToken` | `string` | no |

## ListRestoreTestingPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RestoreTestingPlans` | `List<RestoreTestingPlanForList>` | yes |

## ListRestoreTestingSelections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RestoreTestingPlanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RestoreTestingSelections` | `List<RestoreTestingSelectionForList>` | yes |

## ListScanJobSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `ResourceType` | `string` | no |
| `MalwareScanner` | `string` | no |
| `ScanResultStatus` | `string` | no |
| `State` | `string` | no |
| `AggregationPeriod` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScanJobSummaries` | `List<ScanJobSummary>` | no |
| `AggregationPeriod` | `string` | no |
| `NextToken` | `string` | no |

## ListScanJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByAccountId` | `string` | no |
| `ByBackupVaultName` | `string` | no |
| `ByCompleteAfter` | `timestamp` | no |
| `ByCompleteBefore` | `timestamp` | no |
| `ByMalwareScanner` | `string` | no |
| `ByRecoveryPointArn` | `string` | no |
| `ByResourceArn` | `string` | no |
| `ByResourceType` | `string` | no |
| `ByScanResultStatus` | `string` | no |
| `ByState` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ScanJobs` | `List<ScanJob>` | yes |

## ListTags

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
| `Tags` | `Map<string>` | no |

## ListTieringConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TieringConfigurations` | `List<TieringConfigurationsListMember>` | no |
| `NextToken` | `string` | no |

## PutBackupVaultAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `Policy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutBackupVaultLockConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `MinRetentionDays` | `long` | no |
| `MaxRetentionDays` | `long` | no |
| `ChangeableForDays` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutBackupVaultNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `SNSTopicArn` | `string` | yes |
| `BackupVaultEvents` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRestoreValidationResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreJobId` | `string` | yes |
| `ValidationStatus` | `string` | yes |
| `ValidationStatusMessage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeRestoreAccessBackupVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RestoreAccessBackupVaultArn` | `string` | yes |
| `RequesterComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartBackupJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `LogicallyAirGappedBackupVaultArn` | `string` | no |
| `ResourceArn` | `string` | yes |
| `IamRoleArn` | `string` | yes |
| `IdempotencyToken` | `string` | no |
| `StartWindowMinutes` | `long` | no |
| `CompleteWindowMinutes` | `long` | no |
| `Lifecycle` | `Lifecycle` | no |
| `RecoveryPointTags` | `Map<string>` | no |
| `BackupOptions` | `Map<string>` | no |
| `Index` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupJobId` | `string` | no |
| `RecoveryPointArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `IsParent` | `boolean` | no |

## StartCopyJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecoveryPointArn` | `string` | yes |
| `SourceBackupVaultName` | `string` | yes |
| `DestinationBackupVaultArn` | `string` | yes |
| `IamRoleArn` | `string` | yes |
| `IdempotencyToken` | `string` | no |
| `Lifecycle` | `Lifecycle` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CopyJobId` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `IsParent` | `boolean` | no |

## StartReportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlanName` | `string` | yes |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportJobId` | `string` | no |

## StartRestoreJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecoveryPointArn` | `string` | yes |
| `Metadata` | `Map<string>` | yes |
| `IamRoleArn` | `string` | no |
| `IdempotencyToken` | `string` | no |
| `ResourceType` | `string` | no |
| `CopySourceTagsToRestoredResource` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreJobId` | `string` | no |

## StartScanJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `ContinuousScanEndTime` | `timestamp` | no |
| `IamRoleArn` | `string` | yes |
| `IdempotencyToken` | `string` | no |
| `MalwareScanner` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |
| `ScanBaseRecoveryPointArn` | `string` | no |
| `ScanMode` | `string` | yes |
| `ScannerRoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationDate` | `timestamp` | yes |
| `ScanJobId` | `string` | yes |

## StopBackupJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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
| `TagKeyList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBackupPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | yes |
| `BackupPlan` | `BackupPlanInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPlanId` | `string` | no |
| `BackupPlanArn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `VersionId` | `string` | no |
| `AdvancedBackupSettings` | `List<AdvancedBackupSetting>` | no |
| `ScanSettings` | `List<ScanSetting>` | no |

## UpdateFramework

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FrameworkName` | `string` | yes |
| `FrameworkDescription` | `string` | no |
| `FrameworkControls` | `List<FrameworkControl>` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FrameworkName` | `string` | no |
| `FrameworkArn` | `string` | no |
| `CreationTime` | `timestamp` | no |

## UpdateGlobalSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalSettings` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRecoveryPointIndexSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |
| `IamRoleArn` | `string` | no |
| `Index` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | no |
| `RecoveryPointArn` | `string` | no |
| `IndexStatus` | `string` | no |
| `Index` | `string` | no |

## UpdateRecoveryPointLifecycle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultName` | `string` | yes |
| `RecoveryPointArn` | `string` | yes |
| `Lifecycle` | `Lifecycle` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupVaultArn` | `string` | no |
| `RecoveryPointArn` | `string` | no |
| `Lifecycle` | `Lifecycle` | no |
| `CalculatedLifecycle` | `CalculatedLifecycle` | no |

## UpdateRegionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTypeOptInPreference` | `Map<boolean>` | no |
| `ResourceTypeManagementPreference` | `Map<boolean>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateReportPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlanName` | `string` | yes |
| `ReportPlanDescription` | `string` | no |
| `ReportDeliveryChannel` | `ReportDeliveryChannel` | no |
| `ReportSetting` | `ReportSetting` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportPlanName` | `string` | no |
| `ReportPlanArn` | `string` | no |
| `CreationTime` | `timestamp` | no |

## UpdateRestoreTestingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreTestingPlan` | `RestoreTestingPlanForUpdate` | yes |
| `RestoreTestingPlanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | yes |
| `RestoreTestingPlanArn` | `string` | yes |
| `RestoreTestingPlanName` | `string` | yes |
| `UpdateTime` | `timestamp` | yes |

## UpdateRestoreTestingSelection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestoreTestingPlanName` | `string` | yes |
| `RestoreTestingSelection` | `RestoreTestingSelectionForUpdate` | yes |
| `RestoreTestingSelectionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | yes |
| `RestoreTestingPlanArn` | `string` | yes |
| `RestoreTestingPlanName` | `string` | yes |
| `RestoreTestingSelectionName` | `string` | yes |
| `UpdateTime` | `timestamp` | yes |

## UpdateTieringConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TieringConfigurationName` | `string` | yes |
| `TieringConfiguration` | `TieringConfigurationInputForUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TieringConfigurationArn` | `string` | no |
| `TieringConfigurationName` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastUpdatedTime` | `timestamp` | no |

