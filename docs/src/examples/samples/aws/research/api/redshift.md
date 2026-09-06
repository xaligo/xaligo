# Amazon Redshift

API version: 2012-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/redshift/2012-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptReservedNodeExchange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodeId` | `string` | yes |
| `TargetReservedNodeOfferingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExchangedReservedNode` | `ReservedNode` | no |

## AddPartner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ClusterIdentifier` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `PartnerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | no |
| `PartnerName` | `string` | no |

## AssociateDataShareConsumer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | yes |
| `AssociateEntireAccount` | `boolean` | no |
| `ConsumerArn` | `string` | no |
| `ConsumerRegion` | `string` | no |
| `AllowWrites` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | no |
| `ProducerArn` | `string` | no |
| `AllowPubliclyAccessibleConsumers` | `boolean` | no |
| `DataShareAssociations` | `List<DataShareAssociation>` | no |
| `ManagedBy` | `string` | no |
| `DataShareType` | `string` | no |

## AuthorizeClusterSecurityGroupIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSecurityGroupName` | `string` | yes |
| `CIDRIP` | `string` | no |
| `EC2SecurityGroupName` | `string` | no |
| `EC2SecurityGroupOwnerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSecurityGroup` | `ClusterSecurityGroup` | no |

## AuthorizeDataShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | yes |
| `ConsumerIdentifier` | `string` | yes |
| `AllowWrites` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | no |
| `ProducerArn` | `string` | no |
| `AllowPubliclyAccessibleConsumers` | `boolean` | no |
| `DataShareAssociations` | `List<DataShareAssociation>` | no |
| `ManagedBy` | `string` | no |
| `DataShareType` | `string` | no |

## AuthorizeEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `Account` | `string` | yes |
| `VpcIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grantor` | `string` | no |
| `Grantee` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `AuthorizeTime` | `timestamp` | no |
| `ClusterStatus` | `string` | no |
| `Status` | `string` | no |
| `AllowedAllVPCs` | `boolean` | no |
| `AllowedVPCs` | `List<string>` | no |
| `EndpointCount` | `integer` | no |

## AuthorizeSnapshotAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotIdentifier` | `string` | no |
| `SnapshotArn` | `string` | no |
| `SnapshotClusterIdentifier` | `string` | no |
| `AccountWithRestoreAccess` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## BatchDeleteClusterSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifiers` | `List<DeleteClusterSnapshotMessage>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<string>` | no |
| `Errors` | `List<SnapshotErrorMessage>` | no |

## BatchModifyClusterSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotIdentifierList` | `List<string>` | yes |
| `ManualSnapshotRetentionPeriod` | `integer` | no |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<string>` | no |
| `Errors` | `List<SnapshotErrorMessage>` | no |

## CancelResize

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetNodeType` | `string` | no |
| `TargetNumberOfNodes` | `integer` | no |
| `TargetClusterType` | `string` | no |
| `Status` | `string` | no |
| `ImportTablesCompleted` | `List<string>` | no |
| `ImportTablesInProgress` | `List<string>` | no |
| `ImportTablesNotStarted` | `List<string>` | no |
| `AvgResizeRateInMegaBytesPerSecond` | `double` | no |
| `TotalResizeDataInMegaBytes` | `long` | no |
| `ProgressInMegaBytes` | `long` | no |
| `ElapsedTimeInSeconds` | `long` | no |
| `EstimatedTimeToCompletionInSeconds` | `long` | no |
| `ResizeType` | `string` | no |
| `Message` | `string` | no |
| `TargetEncryptionType` | `string` | no |
| `DataTransferProgressPercent` | `double` | no |

## CopyClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceSnapshotIdentifier` | `string` | yes |
| `SourceSnapshotClusterIdentifier` | `string` | no |
| `TargetSnapshotIdentifier` | `string` | yes |
| `ManualSnapshotRetentionPeriod` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## CreateAuthenticationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileName` | `string` | yes |
| `AuthenticationProfileContent` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileName` | `string` | no |
| `AuthenticationProfileContent` | `string` | no |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBName` | `string` | no |
| `ClusterIdentifier` | `string` | yes |
| `ClusterType` | `string` | no |
| `NodeType` | `string` | yes |
| `MasterUsername` | `string` | yes |
| `MasterUserPassword` | `string` | no |
| `ClusterSecurityGroups` | `List<string>` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `ClusterSubnetGroupName` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `ClusterParameterGroupName` | `string` | no |
| `AutomatedSnapshotRetentionPeriod` | `integer` | no |
| `ManualSnapshotRetentionPeriod` | `integer` | no |
| `Port` | `integer` | no |
| `ClusterVersion` | `string` | no |
| `AllowVersionUpgrade` | `boolean` | no |
| `NumberOfNodes` | `integer` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Encrypted` | `boolean` | no |
| `HsmClientCertificateIdentifier` | `string` | no |
| `HsmConfigurationIdentifier` | `string` | no |
| `ElasticIp` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `EnhancedVpcRouting` | `boolean` | no |
| `AdditionalInfo` | `string` | no |
| `IamRoles` | `List<string>` | no |
| `MaintenanceTrackName` | `string` | no |
| `SnapshotScheduleIdentifier` | `string` | no |
| `AvailabilityZoneRelocation` | `boolean` | no |
| `AquaConfigurationStatus` | `string` | no |
| `DefaultIamRoleArn` | `string` | no |
| `LoadSampleData` | `string` | no |
| `ManageMasterPassword` | `boolean` | no |
| `MasterPasswordSecretKmsKeyId` | `string` | no |
| `IpAddressType` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `RedshiftIdcApplicationArn` | `string` | no |
| `CatalogName` | `string` | no |
| `ExtraComputeForAutomaticOptimization` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## CreateClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `ParameterGroupFamily` | `string` | yes |
| `Description` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterParameterGroup` | `ClusterParameterGroup` | no |

## CreateClusterSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSecurityGroupName` | `string` | yes |
| `Description` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSecurityGroup` | `ClusterSecurityGroup` | no |

## CreateClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotIdentifier` | `string` | yes |
| `ClusterIdentifier` | `string` | yes |
| `ManualSnapshotRetentionPeriod` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## CreateClusterSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSubnetGroupName` | `string` | yes |
| `Description` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSubnetGroup` | `ClusterSubnetGroup` | no |

## CreateCustomDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomDomainName` | `string` | yes |
| `CustomDomainCertificateArn` | `string` | yes |
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomDomainName` | `string` | no |
| `CustomDomainCertificateArn` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `CustomDomainCertExpiryTime` | `string` | no |

## CreateEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `ResourceOwner` | `string` | no |
| `EndpointName` | `string` | yes |
| `SubnetGroupName` | `string` | yes |
| `VpcSecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `ResourceOwner` | `string` | no |
| `SubnetGroupName` | `string` | no |
| `EndpointStatus` | `string` | no |
| `EndpointName` | `string` | no |
| `EndpointCreateTime` | `timestamp` | no |
| `Port` | `integer` | no |
| `Address` | `string` | no |
| `VpcSecurityGroups` | `List<VpcSecurityGroupMembership>` | no |
| `VpcEndpoint` | `VpcEndpoint` | no |

## CreateEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |
| `SnsTopicArn` | `string` | yes |
| `SourceType` | `string` | no |
| `SourceIds` | `List<string>` | no |
| `EventCategories` | `List<string>` | no |
| `Severity` | `string` | no |
| `Enabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## CreateHsmClientCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmClientCertificateIdentifier` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmClientCertificate` | `HsmClientCertificate` | no |

## CreateHsmConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmConfigurationIdentifier` | `string` | yes |
| `Description` | `string` | yes |
| `HsmIpAddress` | `string` | yes |
| `HsmPartitionName` | `string` | yes |
| `HsmPartitionPassword` | `string` | yes |
| `HsmServerPublicCertificate` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmConfiguration` | `HsmConfiguration` | no |

## CreateIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `TargetArn` | `string` | yes |
| `IntegrationName` | `string` | yes |
| `KMSKeyId` | `string` | no |
| `TagList` | `List<Tag>` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationArn` | `string` | no |
| `IntegrationName` | `string` | no |
| `SourceArn` | `string` | no |
| `TargetArn` | `string` | no |
| `Status` | `string` | no |
| `Errors` | `List<IntegrationError>` | no |
| `CreateTime` | `timestamp` | no |
| `Description` | `string` | no |
| `KMSKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |

## CreateQev2IdcApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdcInstanceArn` | `string` | yes |
| `Qev2IdcApplicationName` | `string` | yes |
| `IdcDisplayName` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Qev2IdcApplication` | `Qev2IdcApplication` | no |

## CreateRedshiftIdcApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdcInstanceArn` | `string` | yes |
| `RedshiftIdcApplicationName` | `string` | yes |
| `IdentityNamespace` | `string` | no |
| `IdcDisplayName` | `string` | yes |
| `IamRoleArn` | `string` | yes |
| `AuthorizedTokenIssuerList` | `List<AuthorizedTokenIssuer>` | no |
| `ServiceIntegrations` | `List<ServiceIntegrationsUnion>` | no |
| `ApplicationType` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `SsoTagKeys` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RedshiftIdcApplication` | `RedshiftIdcApplication` | no |

## CreateScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActionName` | `string` | yes |
| `TargetAction` | `ScheduledActionType` | yes |
| `Schedule` | `string` | yes |
| `IamRole` | `string` | yes |
| `ScheduledActionDescription` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Enable` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActionName` | `string` | no |
| `TargetAction` | `ScheduledActionType` | no |
| `Schedule` | `string` | no |
| `IamRole` | `string` | no |
| `ScheduledActionDescription` | `string` | no |
| `State` | `string` | no |
| `NextInvocations` | `List<timestamp>` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |

## CreateSnapshotCopyGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotCopyGrantName` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotCopyGrant` | `SnapshotCopyGrant` | no |

## CreateSnapshotSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduleDefinitions` | `List<string>` | no |
| `ScheduleIdentifier` | `string` | no |
| `ScheduleDescription` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DryRun` | `boolean` | no |
| `NextInvocations` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduleDefinitions` | `List<string>` | no |
| `ScheduleIdentifier` | `string` | no |
| `ScheduleDescription` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `NextInvocations` | `List<timestamp>` | no |
| `AssociatedClusterCount` | `integer` | no |
| `AssociatedClusters` | `List<ClusterAssociatedToSchedule>` | no |

## CreateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateUsageLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `FeatureType` | `string` | yes |
| `LimitType` | `string` | yes |
| `Amount` | `long` | yes |
| `Period` | `string` | no |
| `BreachAction` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageLimitId` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `FeatureType` | `string` | no |
| `LimitType` | `string` | no |
| `Amount` | `long` | no |
| `Period` | `string` | no |
| `BreachAction` | `string` | no |
| `Tags` | `List<Tag>` | no |

## DeauthorizeDataShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | yes |
| `ConsumerIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | no |
| `ProducerArn` | `string` | no |
| `AllowPubliclyAccessibleConsumers` | `boolean` | no |
| `DataShareAssociations` | `List<DataShareAssociation>` | no |
| `ManagedBy` | `string` | no |
| `DataShareType` | `string` | no |

## DeleteAuthenticationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileName` | `string` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `SkipFinalClusterSnapshot` | `boolean` | no |
| `FinalClusterSnapshotIdentifier` | `string` | no |
| `FinalClusterSnapshotRetentionPeriod` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## DeleteClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteClusterSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSecurityGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotIdentifier` | `string` | yes |
| `SnapshotClusterIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## DeleteClusterSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSubnetGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `CustomDomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `ResourceOwner` | `string` | no |
| `SubnetGroupName` | `string` | no |
| `EndpointStatus` | `string` | no |
| `EndpointName` | `string` | no |
| `EndpointCreateTime` | `timestamp` | no |
| `Port` | `integer` | no |
| `Address` | `string` | no |
| `VpcSecurityGroups` | `List<VpcSecurityGroupMembership>` | no |
| `VpcEndpoint` | `VpcEndpoint` | no |

## DeleteEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHsmClientCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmClientCertificateIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHsmConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmConfigurationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationArn` | `string` | no |
| `IntegrationName` | `string` | no |
| `SourceArn` | `string` | no |
| `TargetArn` | `string` | no |
| `Status` | `string` | no |
| `Errors` | `List<IntegrationError>` | no |
| `CreateTime` | `timestamp` | no |
| `Description` | `string` | no |
| `KMSKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |

## DeletePartner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ClusterIdentifier` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `PartnerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | no |
| `PartnerName` | `string` | no |

## DeleteQev2IdcApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Qev2IdcApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRedshiftIdcApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RedshiftIdcApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSnapshotCopyGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotCopyGrantName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSnapshotSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUsageLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageLimitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamespaceIdentifier` | `NamespaceIdentifierUnion` | yes |
| `ConsumerIdentifiers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## DescribeAccountAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttributeNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAttributes` | `List<AccountAttribute>` | no |

## DescribeAuthenticationProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfiles` | `List<AuthenticationProfile>` | no |

## DescribeClusterDbRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ClusterDbRevisions` | `List<ClusterDbRevision>` | no |

## DescribeClusterParameterGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ParameterGroups` | `List<ClusterParameterGroup>` | no |

## DescribeClusterParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `Source` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<Parameter>` | no |
| `Marker` | `string` | no |

## DescribeClusterSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSecurityGroupName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ClusterSecurityGroups` | `List<ClusterSecurityGroup>` | no |

## DescribeClusterSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `SnapshotIdentifier` | `string` | no |
| `SnapshotArn` | `string` | no |
| `SnapshotType` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `OwnerAccount` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |
| `ClusterExists` | `boolean` | no |
| `SortingEntities` | `List<SnapshotSortingEntity>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Snapshots` | `List<Snapshot>` | no |

## DescribeClusterSubnetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSubnetGroupName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ClusterSubnetGroups` | `List<ClusterSubnetGroup>` | no |

## DescribeClusterTracks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaintenanceTrackName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaintenanceTracks` | `List<MaintenanceTrack>` | no |
| `Marker` | `string` | no |

## DescribeClusterVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterVersion` | `string` | no |
| `ClusterParameterGroupFamily` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ClusterVersions` | `List<ClusterVersion>` | no |

## DescribeClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Clusters` | `List<Cluster>` | no |

## DescribeCustomDomainAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomDomainName` | `string` | no |
| `CustomDomainCertificateArn` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Associations` | `List<Association>` | no |

## DescribeDataShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShares` | `List<DataShare>` | no |
| `Marker` | `string` | no |

## DescribeDataSharesForConsumer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsumerArn` | `string` | no |
| `Status` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShares` | `List<DataShare>` | no |
| `Marker` | `string` | no |

## DescribeDataSharesForProducer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProducerArn` | `string` | no |
| `Status` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShares` | `List<DataShare>` | no |
| `Marker` | `string` | no |

## DescribeDefaultClusterParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupFamily` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultClusterParameters` | `DefaultClusterParameters` | no |

## DescribeEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `ResourceOwner` | `string` | no |
| `EndpointName` | `string` | no |
| `VpcId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointAccessList` | `List<EndpointAccess>` | no |
| `Marker` | `string` | no |

## DescribeEndpointAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `Account` | `string` | no |
| `Grantee` | `boolean` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointAuthorizationList` | `List<EndpointAuthorization>` | no |
| `Marker` | `string` | no |

## DescribeEventCategories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventCategoriesMapList` | `List<EventCategoriesMap>` | no |

## DescribeEventSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `EventSubscriptionsList` | `List<EventSubscription>` | no |

## DescribeEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceIdentifier` | `string` | no |
| `SourceType` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Duration` | `integer` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Events` | `List<Event>` | no |

## DescribeHsmClientCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmClientCertificateIdentifier` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `HsmClientCertificates` | `List<HsmClientCertificate>` | no |

## DescribeHsmConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmConfigurationIdentifier` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `HsmConfigurations` | `List<HsmConfiguration>` | no |

## DescribeInboundIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationArn` | `string` | no |
| `TargetArn` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `InboundIntegrations` | `List<InboundIntegration>` | no |

## DescribeIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationArn` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `Filters` | `List<DescribeIntegrationsFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Integrations` | `List<Integration>` | no |

## DescribeLoggingStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingEnabled` | `boolean` | no |
| `BucketName` | `string` | no |
| `S3KeyPrefix` | `string` | no |
| `LastSuccessfulDeliveryTime` | `timestamp` | no |
| `LastFailureTime` | `timestamp` | no |
| `LastFailureMessage` | `string` | no |
| `LogDestinationType` | `string` | no |
| `LogExports` | `List<string>` | no |
| `S3Tables` | `S3TablePublishStatus` | no |

## DescribeNodeConfigurationOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionType` | `string` | yes |
| `ClusterIdentifier` | `string` | no |
| `SnapshotIdentifier` | `string` | no |
| `SnapshotArn` | `string` | no |
| `OwnerAccount` | `string` | no |
| `Filters` | `List<NodeConfigurationOptionsFilter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NodeConfigurationOptionList` | `List<NodeConfigurationOption>` | no |
| `Marker` | `string` | no |

## DescribeOrderableClusterOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterVersion` | `string` | no |
| `NodeType` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrderableClusterOptions` | `List<OrderableClusterOption>` | no |
| `Marker` | `string` | no |

## DescribePartners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ClusterIdentifier` | `string` | yes |
| `DatabaseName` | `string` | no |
| `PartnerName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PartnerIntegrationInfoList` | `List<PartnerIntegrationInfo>` | no |

## DescribeQev2IdcApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Qev2IdcApplicationArn` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Qev2IdcApplications` | `List<Qev2IdcApplication>` | no |
| `Marker` | `string` | no |

## DescribeRedshiftIdcApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RedshiftIdcApplicationArn` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RedshiftIdcApplications` | `List<RedshiftIdcApplication>` | no |
| `Marker` | `string` | no |

## DescribeReservedNodeExchangeStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodeId` | `string` | no |
| `ReservedNodeExchangeRequestId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodeExchangeStatusDetails` | `List<ReservedNodeExchangeStatus>` | no |
| `Marker` | `string` | no |

## DescribeReservedNodeOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodeOfferingId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReservedNodeOfferings` | `List<ReservedNodeOffering>` | no |

## DescribeReservedNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodeId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReservedNodes` | `List<ReservedNode>` | no |

## DescribeResize

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetNodeType` | `string` | no |
| `TargetNumberOfNodes` | `integer` | no |
| `TargetClusterType` | `string` | no |
| `Status` | `string` | no |
| `ImportTablesCompleted` | `List<string>` | no |
| `ImportTablesInProgress` | `List<string>` | no |
| `ImportTablesNotStarted` | `List<string>` | no |
| `AvgResizeRateInMegaBytesPerSecond` | `double` | no |
| `TotalResizeDataInMegaBytes` | `long` | no |
| `ProgressInMegaBytes` | `long` | no |
| `ElapsedTimeInSeconds` | `long` | no |
| `EstimatedTimeToCompletionInSeconds` | `long` | no |
| `ResizeType` | `string` | no |
| `Message` | `string` | no |
| `TargetEncryptionType` | `string` | no |
| `DataTransferProgressPercent` | `double` | no |

## DescribeScheduledActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActionName` | `string` | no |
| `TargetActionType` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Active` | `boolean` | no |
| `Filters` | `List<ScheduledActionFilter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ScheduledActions` | `List<ScheduledAction>` | no |

## DescribeSnapshotCopyGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotCopyGrantName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `SnapshotCopyGrants` | `List<SnapshotCopyGrant>` | no |

## DescribeSnapshotSchedules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `ScheduleIdentifier` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotSchedules` | `List<SnapshotSchedule>` | no |
| `Marker` | `string` | no |

## DescribeStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TotalBackupSizeInMegaBytes` | `double` | no |
| `TotalProvisionedStorageInMegaBytes` | `double` | no |

## DescribeTableRestoreStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `TableRestoreRequestId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableRestoreStatusDetails` | `List<TableRestoreStatus>` | no |
| `Marker` | `string` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | no |
| `ResourceType` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaggedResources` | `List<TaggedResource>` | no |
| `Marker` | `string` | no |

## DescribeUsageLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageLimitId` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `FeatureType` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `TagValues` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageLimits` | `List<UsageLimit>` | no |
| `Marker` | `string` | no |

## DisableLogging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `LogDestinationType` | `string` | no |
| `LogExports` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingEnabled` | `boolean` | no |
| `BucketName` | `string` | no |
| `S3KeyPrefix` | `string` | no |
| `LastSuccessfulDeliveryTime` | `timestamp` | no |
| `LastFailureTime` | `timestamp` | no |
| `LastFailureMessage` | `string` | no |
| `LogDestinationType` | `string` | no |
| `LogExports` | `List<string>` | no |
| `S3Tables` | `S3TablePublishStatus` | no |

## DisableSnapshotCopy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## DisassociateDataShareConsumer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | yes |
| `DisassociateEntireAccount` | `boolean` | no |
| `ConsumerArn` | `string` | no |
| `ConsumerRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | no |
| `ProducerArn` | `string` | no |
| `AllowPubliclyAccessibleConsumers` | `boolean` | no |
| `DataShareAssociations` | `List<DataShareAssociation>` | no |
| `ManagedBy` | `string` | no |
| `DataShareType` | `string` | no |

## EnableLogging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `BucketName` | `string` | no |
| `S3KeyPrefix` | `string` | no |
| `LogDestinationType` | `string` | no |
| `LogExports` | `List<string>` | no |
| `S3TableKmsKeyId` | `string` | no |
| `S3TableGranularity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingEnabled` | `boolean` | no |
| `BucketName` | `string` | no |
| `S3KeyPrefix` | `string` | no |
| `LastSuccessfulDeliveryTime` | `timestamp` | no |
| `LastFailureTime` | `timestamp` | no |
| `LastFailureMessage` | `string` | no |
| `LogDestinationType` | `string` | no |
| `LogExports` | `List<string>` | no |
| `S3Tables` | `S3TablePublishStatus` | no |

## EnableSnapshotCopy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `DestinationRegion` | `string` | yes |
| `RetentionPeriod` | `integer` | no |
| `SnapshotCopyGrantName` | `string` | no |
| `ManualSnapshotRetentionPeriod` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## FailoverPrimaryCompute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## GetClusterCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DbUser` | `string` | yes |
| `DbName` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `DurationSeconds` | `integer` | no |
| `AutoCreate` | `boolean` | no |
| `DbGroups` | `List<string>` | no |
| `CustomDomainName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DbUser` | `string` | no |
| `DbPassword` | `string` | no |
| `Expiration` | `timestamp` | no |

## GetClusterCredentialsWithIAM

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DbName` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `DurationSeconds` | `integer` | no |
| `CustomDomainName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DbUser` | `string` | no |
| `DbPassword` | `string` | no |
| `Expiration` | `timestamp` | no |
| `NextRefreshTime` | `timestamp` | no |

## GetIdentityCenterAuthToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Token` | `string` | no |
| `ExpirationTime` | `timestamp` | no |

## GetReservedNodeExchangeConfigurationOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionType` | `string` | yes |
| `ClusterIdentifier` | `string` | no |
| `SnapshotIdentifier` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReservedNodeConfigurationOptionList` | `List<ReservedNodeConfigurationOption>` | no |

## GetReservedNodeExchangeOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodeId` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReservedNodeOfferings` | `List<ReservedNodeOffering>` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePolicy` | `ResourcePolicy` | no |

## ListRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `NamespaceArn` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Recommendations` | `List<Recommendation>` | no |
| `Marker` | `string` | no |

## ModifyAquaConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `AquaConfigurationStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AquaConfiguration` | `AquaConfiguration` | no |

## ModifyAuthenticationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileName` | `string` | yes |
| `AuthenticationProfileContent` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileName` | `string` | no |
| `AuthenticationProfileContent` | `string` | no |

## ModifyCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `ClusterType` | `string` | no |
| `NodeType` | `string` | no |
| `NumberOfNodes` | `integer` | no |
| `ClusterSecurityGroups` | `List<string>` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `MasterUserPassword` | `string` | no |
| `ClusterParameterGroupName` | `string` | no |
| `AutomatedSnapshotRetentionPeriod` | `integer` | no |
| `ManualSnapshotRetentionPeriod` | `integer` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `ClusterVersion` | `string` | no |
| `AllowVersionUpgrade` | `boolean` | no |
| `HsmClientCertificateIdentifier` | `string` | no |
| `HsmConfigurationIdentifier` | `string` | no |
| `NewClusterIdentifier` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `ElasticIp` | `string` | no |
| `EnhancedVpcRouting` | `boolean` | no |
| `MaintenanceTrackName` | `string` | no |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `AvailabilityZoneRelocation` | `boolean` | no |
| `AvailabilityZone` | `string` | no |
| `Port` | `integer` | no |
| `ManageMasterPassword` | `boolean` | no |
| `MasterPasswordSecretKmsKeyId` | `string` | no |
| `IpAddressType` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `ExtraComputeForAutomaticOptimization` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## ModifyClusterDbRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `RevisionTarget` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## ModifyClusterIamRoles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `AddIamRoles` | `List<string>` | no |
| `RemoveIamRoles` | `List<string>` | no |
| `DefaultIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## ModifyClusterMaintenance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `DeferMaintenance` | `boolean` | no |
| `DeferMaintenanceIdentifier` | `string` | no |
| `DeferMaintenanceStartTime` | `timestamp` | no |
| `DeferMaintenanceEndTime` | `timestamp` | no |
| `DeferMaintenanceDuration` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## ModifyClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `Parameters` | `List<Parameter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | no |
| `ParameterGroupStatus` | `string` | no |

## ModifyClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotIdentifier` | `string` | yes |
| `ManualSnapshotRetentionPeriod` | `integer` | no |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## ModifyClusterSnapshotSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `ScheduleIdentifier` | `string` | no |
| `DisassociateSchedule` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyClusterSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSubnetGroupName` | `string` | yes |
| `Description` | `string` | no |
| `SubnetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSubnetGroup` | `ClusterSubnetGroup` | no |

## ModifyCustomDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomDomainName` | `string` | yes |
| `CustomDomainCertificateArn` | `string` | yes |
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomDomainName` | `string` | no |
| `CustomDomainCertificateArn` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `CustomDomainCertExpiryTime` | `string` | no |

## ModifyEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `VpcSecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `ResourceOwner` | `string` | no |
| `SubnetGroupName` | `string` | no |
| `EndpointStatus` | `string` | no |
| `EndpointName` | `string` | no |
| `EndpointCreateTime` | `timestamp` | no |
| `Port` | `integer` | no |
| `Address` | `string` | no |
| `VpcSecurityGroups` | `List<VpcSecurityGroupMembership>` | no |
| `VpcEndpoint` | `VpcEndpoint` | no |

## ModifyEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |
| `SnsTopicArn` | `string` | no |
| `SourceType` | `string` | no |
| `SourceIds` | `List<string>` | no |
| `EventCategories` | `List<string>` | no |
| `Severity` | `string` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## ModifyIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationArn` | `string` | yes |
| `Description` | `string` | no |
| `IntegrationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationArn` | `string` | no |
| `IntegrationName` | `string` | no |
| `SourceArn` | `string` | no |
| `TargetArn` | `string` | no |
| `Status` | `string` | no |
| `Errors` | `List<IntegrationError>` | no |
| `CreateTime` | `timestamp` | no |
| `Description` | `string` | no |
| `KMSKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |

## ModifyLakehouseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `LakehouseRegistration` | `string` | no |
| `CatalogName` | `string` | no |
| `LakehouseIdcRegistration` | `string` | no |
| `LakehouseIdcApplicationArn` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `LakehouseIdcApplicationArn` | `string` | no |
| `LakehouseRegistrationStatus` | `string` | no |
| `CatalogArn` | `string` | no |

## ModifyQev2IdcApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Qev2IdcApplicationArn` | `string` | yes |
| `IdcDisplayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Qev2IdcApplication` | `Qev2IdcApplication` | no |

## ModifyRedshiftIdcApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RedshiftIdcApplicationArn` | `string` | yes |
| `IdentityNamespace` | `string` | no |
| `IamRoleArn` | `string` | no |
| `IdcDisplayName` | `string` | no |
| `AuthorizedTokenIssuerList` | `List<AuthorizedTokenIssuer>` | no |
| `ServiceIntegrations` | `List<ServiceIntegrationsUnion>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RedshiftIdcApplication` | `RedshiftIdcApplication` | no |

## ModifyScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActionName` | `string` | yes |
| `TargetAction` | `ScheduledActionType` | no |
| `Schedule` | `string` | no |
| `IamRole` | `string` | no |
| `ScheduledActionDescription` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Enable` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActionName` | `string` | no |
| `TargetAction` | `ScheduledActionType` | no |
| `Schedule` | `string` | no |
| `IamRole` | `string` | no |
| `ScheduledActionDescription` | `string` | no |
| `State` | `string` | no |
| `NextInvocations` | `List<timestamp>` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |

## ModifySnapshotCopyRetentionPeriod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `RetentionPeriod` | `integer` | yes |
| `Manual` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## ModifySnapshotSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduleIdentifier` | `string` | yes |
| `ScheduleDefinitions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduleDefinitions` | `List<string>` | no |
| `ScheduleIdentifier` | `string` | no |
| `ScheduleDescription` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `NextInvocations` | `List<timestamp>` | no |
| `AssociatedClusterCount` | `integer` | no |
| `AssociatedClusters` | `List<ClusterAssociatedToSchedule>` | no |

## ModifyUsageLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageLimitId` | `string` | yes |
| `Amount` | `long` | no |
| `BreachAction` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageLimitId` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `FeatureType` | `string` | no |
| `LimitType` | `string` | no |
| `Amount` | `long` | no |
| `Period` | `string` | no |
| `BreachAction` | `string` | no |
| `Tags` | `List<Tag>` | no |

## PauseCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## PurchaseReservedNodeOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodeOfferingId` | `string` | yes |
| `NodeCount` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNode` | `ReservedNode` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePolicy` | `ResourcePolicy` | no |

## RebootCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## RegisterNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamespaceIdentifier` | `NamespaceIdentifierUnion` | yes |
| `ConsumerIdentifiers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## RejectDataShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataShareArn` | `string` | no |
| `ProducerArn` | `string` | no |
| `AllowPubliclyAccessibleConsumers` | `boolean` | no |
| `DataShareAssociations` | `List<DataShareAssociation>` | no |
| `ManagedBy` | `string` | no |
| `DataShareType` | `string` | no |

## ResetClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `ResetAllParameters` | `boolean` | no |
| `Parameters` | `List<Parameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | no |
| `ParameterGroupStatus` | `string` | no |

## ResizeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `ClusterType` | `string` | no |
| `NodeType` | `string` | no |
| `NumberOfNodes` | `integer` | no |
| `Classic` | `boolean` | no |
| `ReservedNodeId` | `string` | no |
| `TargetReservedNodeOfferingId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## RestoreFromClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `SnapshotIdentifier` | `string` | no |
| `SnapshotArn` | `string` | no |
| `SnapshotClusterIdentifier` | `string` | no |
| `Port` | `integer` | no |
| `AvailabilityZone` | `string` | no |
| `AllowVersionUpgrade` | `boolean` | no |
| `ClusterSubnetGroupName` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `OwnerAccount` | `string` | no |
| `HsmClientCertificateIdentifier` | `string` | no |
| `HsmConfigurationIdentifier` | `string` | no |
| `ElasticIp` | `string` | no |
| `ClusterParameterGroupName` | `string` | no |
| `ClusterSecurityGroups` | `List<string>` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `AutomatedSnapshotRetentionPeriod` | `integer` | no |
| `ManualSnapshotRetentionPeriod` | `integer` | no |
| `KmsKeyId` | `string` | no |
| `NodeType` | `string` | no |
| `EnhancedVpcRouting` | `boolean` | no |
| `AdditionalInfo` | `string` | no |
| `IamRoles` | `List<string>` | no |
| `MaintenanceTrackName` | `string` | no |
| `SnapshotScheduleIdentifier` | `string` | no |
| `NumberOfNodes` | `integer` | no |
| `AvailabilityZoneRelocation` | `boolean` | no |
| `AquaConfigurationStatus` | `string` | no |
| `DefaultIamRoleArn` | `string` | no |
| `ReservedNodeId` | `string` | no |
| `TargetReservedNodeOfferingId` | `string` | no |
| `Encrypted` | `boolean` | no |
| `ManageMasterPassword` | `boolean` | no |
| `MasterPasswordSecretKmsKeyId` | `string` | no |
| `IpAddressType` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `CatalogName` | `string` | no |
| `RedshiftIdcApplicationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## RestoreTableFromClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |
| `SnapshotIdentifier` | `string` | yes |
| `SourceDatabaseName` | `string` | yes |
| `SourceSchemaName` | `string` | no |
| `SourceTableName` | `string` | yes |
| `TargetDatabaseName` | `string` | no |
| `TargetSchemaName` | `string` | no |
| `NewTableName` | `string` | yes |
| `EnableCaseSensitiveIdentifier` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableRestoreStatus` | `TableRestoreStatus` | no |

## ResumeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## RevokeClusterSecurityGroupIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSecurityGroupName` | `string` | yes |
| `CIDRIP` | `string` | no |
| `EC2SecurityGroupName` | `string` | no |
| `EC2SecurityGroupOwnerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSecurityGroup` | `ClusterSecurityGroup` | no |

## RevokeEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `Account` | `string` | no |
| `VpcIds` | `List<string>` | no |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grantor` | `string` | no |
| `Grantee` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `AuthorizeTime` | `timestamp` | no |
| `ClusterStatus` | `string` | no |
| `Status` | `string` | no |
| `AllowedAllVPCs` | `boolean` | no |
| `AllowedVPCs` | `List<string>` | no |
| `EndpointCount` | `integer` | no |

## RevokeSnapshotAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotIdentifier` | `string` | no |
| `SnapshotArn` | `string` | no |
| `SnapshotClusterIdentifier` | `string` | no |
| `AccountWithRestoreAccess` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## RotateEncryptionKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## UpdatePartnerStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ClusterIdentifier` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `PartnerName` | `string` | yes |
| `Status` | `string` | yes |
| `StatusMessage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | no |
| `PartnerName` | `string` | no |

