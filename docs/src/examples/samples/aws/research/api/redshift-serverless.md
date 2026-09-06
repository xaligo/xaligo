# Redshift Serverless

API version: 2021-04-21. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/redshift-serverless/2021-04-21/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ConvertRecoveryPointToSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPointId` | `string` | yes |
| `retentionPeriod` | `integer` | no |
| `snapshotName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `Snapshot` | no |

## CreateCustomDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainCertificateArn` | `string` | yes |
| `customDomainName` | `string` | yes |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainCertificateArn` | `string` | no |
| `customDomainCertificateExpiryTime` | `timestamp` | no |
| `customDomainName` | `string` | no |
| `workgroupName` | `string` | no |

## CreateEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointName` | `string` | yes |
| `ownerAccount` | `string` | no |
| `subnetIds` | `List<string>` | yes |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `EndpointAccess` | no |

## CreateNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminPasswordSecretKmsKeyId` | `string` | no |
| `adminUserPassword` | `string` | no |
| `adminUsername` | `string` | no |
| `dbName` | `string` | no |
| `defaultIamRoleArn` | `string` | no |
| `iamRoles` | `List<string>` | no |
| `kmsKeyId` | `string` | no |
| `logExports` | `List<string>` | no |
| `manageAdminPassword` | `boolean` | no |
| `namespaceName` | `string` | yes |
| `redshiftIdcApplicationArn` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `Namespace` | no |

## CreateReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacity` | `integer` | yes |
| `clientToken` | `string` | no |
| `offeringId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reservation` | `Reservation` | no |

## CreateScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabled` | `boolean` | no |
| `endTime` | `timestamp` | no |
| `namespaceName` | `string` | yes |
| `roleArn` | `string` | yes |
| `schedule` | `Schedule` | yes |
| `scheduledActionDescription` | `string` | no |
| `scheduledActionName` | `string` | yes |
| `startTime` | `timestamp` | no |
| `targetAction` | `TargetAction` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAction` | `ScheduledActionResponse` | no |

## CreateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaceName` | `string` | yes |
| `retentionPeriod` | `integer` | no |
| `snapshotName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `Snapshot` | no |

## CreateSnapshotCopyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationKmsKeyId` | `string` | no |
| `destinationRegion` | `string` | yes |
| `namespaceName` | `string` | yes |
| `snapshotRetentionPeriod` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotCopyConfiguration` | `SnapshotCopyConfiguration` | yes |

## CreateUsageLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `amount` | `long` | yes |
| `breachAction` | `string` | no |
| `period` | `string` | no |
| `resourceArn` | `string` | yes |
| `usageType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usageLimit` | `UsageLimit` | no |

## CreateWorkgroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baseCapacity` | `integer` | no |
| `configParameters` | `List<ConfigParameter>` | no |
| `enhancedVpcRouting` | `boolean` | no |
| `extraComputeForAutomaticOptimization` | `boolean` | no |
| `ipAddressType` | `string` | no |
| `maxCapacity` | `integer` | no |
| `namespaceName` | `string` | yes |
| `port` | `integer` | no |
| `pricePerformanceTarget` | `PerformanceTarget` | no |
| `publiclyAccessible` | `boolean` | no |
| `securityGroupIds` | `List<string>` | no |
| `subnetIds` | `List<string>` | no |
| `tags` | `List<Tag>` | no |
| `trackName` | `string` | no |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workgroup` | `Workgroup` | no |

## DeleteCustomDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainName` | `string` | yes |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `EndpointAccess` | no |

## DeleteNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `finalSnapshotName` | `string` | no |
| `finalSnapshotRetentionPeriod` | `integer` | no |
| `namespaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `Namespace` | yes |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledActionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAction` | `ScheduledActionResponse` | no |

## DeleteSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `Snapshot` | no |

## DeleteSnapshotCopyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotCopyConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotCopyConfiguration` | `SnapshotCopyConfiguration` | yes |

## DeleteUsageLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usageLimitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usageLimit` | `UsageLimit` | no |

## DeleteWorkgroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workgroup` | `Workgroup` | yes |

## GetCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainName` | `string` | no |
| `dbName` | `string` | no |
| `durationSeconds` | `integer` | no |
| `workgroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbPassword` | `string` | no |
| `dbUser` | `string` | no |
| `expiration` | `timestamp` | no |
| `nextRefreshTime` | `timestamp` | no |

## GetCustomDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainName` | `string` | yes |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainCertificateArn` | `string` | no |
| `customDomainCertificateExpiryTime` | `timestamp` | no |
| `customDomainName` | `string` | no |
| `workgroupName` | `string` | no |

## GetEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `EndpointAccess` | no |

## GetIdentityCenterAuthToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workgroupNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `expirationTime` | `timestamp` | no |
| `token` | `string` | no |

## GetNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `Namespace` | yes |

## GetRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPoint` | `RecoveryPoint` | no |

## GetReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reservationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reservation` | `Reservation` | yes |

## GetReservationOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offeringId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reservationOffering` | `ReservationOffering` | yes |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePolicy` | `ResourcePolicy` | no |

## GetScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledActionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAction` | `ScheduledActionResponse` | no |

## GetSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `snapshotArn` | `string` | no |
| `snapshotName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `Snapshot` | no |

## GetTableRestoreStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableRestoreRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableRestoreStatus` | `TableRestoreStatus` | no |

## GetTrack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `track` | `ServerlessTrack` | no |

## GetUsageLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usageLimitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usageLimit` | `UsageLimit` | no |

## GetWorkgroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workgroup` | `Workgroup` | yes |

## ListCustomDomainAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainCertificateArn` | `string` | no |
| `customDomainName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associations` | `List<Association>` | no |
| `nextToken` | `string` | no |

## ListEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `ownerAccount` | `string` | no |
| `vpcId` | `string` | no |
| `workgroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoints` | `List<EndpointAccess>` | yes |
| `nextToken` | `string` | no |

## ListManagedWorkgroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managedWorkgroups` | `List<ManagedWorkgroupListItem>` | no |
| `nextToken` | `string` | no |

## ListNamespaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaces` | `List<Namespace>` | yes |
| `nextToken` | `string` | no |

## ListRecoveryPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endTime` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `namespaceArn` | `string` | no |
| `namespaceName` | `string` | no |
| `nextToken` | `string` | no |
| `startTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `recoveryPoints` | `List<RecoveryPoint>` | no |

## ListReservationOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `reservationOfferingsList` | `List<ReservationOffering>` | yes |

## ListReservations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `reservationsList` | `List<Reservation>` | yes |

## ListScheduledActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `namespaceName` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `scheduledActions` | `List<ScheduledActionAssociation>` | no |

## ListSnapshotCopyConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `namespaceName` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `snapshotCopyConfigurations` | `List<SnapshotCopyConfiguration>` | yes |

## ListSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endTime` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `namespaceArn` | `string` | no |
| `namespaceName` | `string` | no |
| `nextToken` | `string` | no |
| `ownerAccount` | `string` | no |
| `startTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `snapshots` | `List<Snapshot>` | no |

## ListTableRestoreStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `namespaceName` | `string` | no |
| `nextToken` | `string` | no |
| `workgroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `tableRestoreStatuses` | `List<TableRestoreStatus>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## ListTracks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `tracks` | `List<ServerlessTrack>` | no |

## ListUsageLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `resourceArn` | `string` | no |
| `usageType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `usageLimits` | `List<UsageLimit>` | no |

## ListWorkgroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `ownerAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `workgroups` | `List<Workgroup>` | yes |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePolicy` | `ResourcePolicy` | no |

## RestoreFromRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maintainIntegration` | `boolean` | no |
| `namespaceName` | `string` | yes |
| `recoveryPointId` | `string` | yes |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `Namespace` | no |
| `recoveryPointId` | `string` | no |

## RestoreFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminPasswordSecretKmsKeyId` | `string` | no |
| `maintainIntegration` | `boolean` | no |
| `manageAdminPassword` | `boolean` | no |
| `namespaceName` | `string` | yes |
| `ownerAccount` | `string` | no |
| `snapshotArn` | `string` | no |
| `snapshotName` | `string` | no |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `Namespace` | no |
| `ownerAccount` | `string` | no |
| `snapshotName` | `string` | no |

## RestoreTableFromRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activateCaseSensitiveIdentifier` | `boolean` | no |
| `namespaceName` | `string` | yes |
| `newTableName` | `string` | yes |
| `recoveryPointId` | `string` | yes |
| `sourceDatabaseName` | `string` | yes |
| `sourceSchemaName` | `string` | no |
| `sourceTableName` | `string` | yes |
| `targetDatabaseName` | `string` | no |
| `targetSchemaName` | `string` | no |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableRestoreStatus` | `TableRestoreStatus` | no |

## RestoreTableFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activateCaseSensitiveIdentifier` | `boolean` | no |
| `namespaceName` | `string` | yes |
| `newTableName` | `string` | yes |
| `snapshotName` | `string` | yes |
| `sourceDatabaseName` | `string` | yes |
| `sourceSchemaName` | `string` | no |
| `sourceTableName` | `string` | yes |
| `targetDatabaseName` | `string` | no |
| `targetSchemaName` | `string` | no |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableRestoreStatus` | `TableRestoreStatus` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateCustomDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainCertificateArn` | `string` | yes |
| `customDomainName` | `string` | yes |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customDomainCertificateArn` | `string` | no |
| `customDomainCertificateExpiryTime` | `timestamp` | no |
| `customDomainName` | `string` | no |
| `workgroupName` | `string` | no |

## UpdateEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointName` | `string` | yes |
| `vpcSecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `EndpointAccess` | no |

## UpdateLakehouseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalogName` | `string` | no |
| `dryRun` | `boolean` | no |
| `lakehouseIdcApplicationArn` | `string` | no |
| `lakehouseIdcRegistration` | `string` | no |
| `lakehouseRegistration` | `string` | no |
| `namespaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalogArn` | `string` | no |
| `lakehouseIdcApplicationArn` | `string` | no |
| `lakehouseRegistrationStatus` | `string` | no |
| `namespaceName` | `string` | no |

## UpdateNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminPasswordSecretKmsKeyId` | `string` | no |
| `adminUserPassword` | `string` | no |
| `adminUsername` | `string` | no |
| `defaultIamRoleArn` | `string` | no |
| `iamRoles` | `List<string>` | no |
| `kmsKeyId` | `string` | no |
| `logDestinationType` | `string` | no |
| `logExports` | `List<string>` | no |
| `manageAdminPassword` | `boolean` | no |
| `namespaceName` | `string` | yes |
| `s3TableAction` | `string` | no |
| `s3TableGranularity` | `string` | no |
| `s3TableKmsKeyId` | `string` | no |
| `s3TableNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `Namespace` | yes |

## UpdateScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabled` | `boolean` | no |
| `endTime` | `timestamp` | no |
| `roleArn` | `string` | no |
| `schedule` | `Schedule` | no |
| `scheduledActionDescription` | `string` | no |
| `scheduledActionName` | `string` | yes |
| `startTime` | `timestamp` | no |
| `targetAction` | `TargetAction` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAction` | `ScheduledActionResponse` | no |

## UpdateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `retentionPeriod` | `integer` | no |
| `snapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `Snapshot` | no |

## UpdateSnapshotCopyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotCopyConfigurationId` | `string` | yes |
| `snapshotRetentionPeriod` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotCopyConfiguration` | `SnapshotCopyConfiguration` | yes |

## UpdateUsageLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `amount` | `long` | no |
| `breachAction` | `string` | no |
| `usageLimitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usageLimit` | `UsageLimit` | no |

## UpdateWorkgroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baseCapacity` | `integer` | no |
| `configParameters` | `List<ConfigParameter>` | no |
| `enhancedVpcRouting` | `boolean` | no |
| `extraComputeForAutomaticOptimization` | `boolean` | no |
| `ipAddressType` | `string` | no |
| `maxCapacity` | `integer` | no |
| `port` | `integer` | no |
| `pricePerformanceTarget` | `PerformanceTarget` | no |
| `publiclyAccessible` | `boolean` | no |
| `securityGroupIds` | `List<string>` | no |
| `subnetIds` | `List<string>` | no |
| `trackName` | `string` | no |
| `workgroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workgroup` | `Workgroup` | yes |

