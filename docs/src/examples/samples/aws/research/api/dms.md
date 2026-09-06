# AWS Database Migration Service

API version: 2016-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/dms/2016-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTagsToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ApplyPendingMaintenanceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceArn` | `string` | yes |
| `ApplyAction` | `string` | yes |
| `OptInType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePendingMaintenanceActions` | `ResourcePendingMaintenanceActions` | no |

## BatchStartRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Data` | `List<StartRecommendationsRequestEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ErrorEntries` | `List<BatchStartRecommendationsErrorEntry>` | no |

## CancelMetadataModelConversion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `RequestIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Request` | `SchemaConversionRequest` | no |

## CancelMetadataModelCreation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `RequestIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Request` | `SchemaConversionRequest` | no |

## CancelReplicationTaskAssessmentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskAssessmentRunArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskAssessmentRun` | `ReplicationTaskAssessmentRun` | no |

## CreateDataMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigrationName` | `string` | no |
| `MigrationProjectIdentifier` | `string` | yes |
| `DataMigrationType` | `string` | yes |
| `ServiceAccessRoleArn` | `string` | yes |
| `EnableCloudwatchLogs` | `boolean` | no |
| `SourceDataSettings` | `List<SourceDataSetting>` | no |
| `TargetDataSettings` | `List<TargetDataSetting>` | no |
| `NumberOfJobs` | `integer` | no |
| `Tags` | `List<Tag>` | no |
| `SelectionRules` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigration` | `DataMigration` | no |

## CreateDataProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataProviderName` | `string` | no |
| `Description` | `string` | no |
| `Engine` | `string` | yes |
| `Virtual` | `boolean` | no |
| `Settings` | `DataProviderSettings` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataProvider` | `DataProvider` | no |

## CreateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointIdentifier` | `string` | yes |
| `EndpointType` | `string` | yes |
| `EngineName` | `string` | yes |
| `Username` | `string` | no |
| `Password` | `string` | no |
| `ServerName` | `string` | no |
| `Port` | `integer` | no |
| `DatabaseName` | `string` | no |
| `ExtraConnectionAttributes` | `string` | no |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `CertificateArn` | `string` | no |
| `SslMode` | `string` | no |
| `ServiceAccessRoleArn` | `string` | no |
| `ExternalTableDefinition` | `string` | no |
| `DynamoDbSettings` | `DynamoDbSettings` | no |
| `S3Settings` | `S3Settings` | no |
| `DmsTransferSettings` | `DmsTransferSettings` | no |
| `MongoDbSettings` | `MongoDbSettings` | no |
| `KinesisSettings` | `KinesisSettings` | no |
| `KafkaSettings` | `KafkaSettings` | no |
| `ElasticsearchSettings` | `ElasticsearchSettings` | no |
| `NeptuneSettings` | `NeptuneSettings` | no |
| `RedshiftSettings` | `RedshiftSettings` | no |
| `PostgreSQLSettings` | `PostgreSQLSettings` | no |
| `MySQLSettings` | `MySQLSettings` | no |
| `OracleSettings` | `OracleSettings` | no |
| `SybaseSettings` | `SybaseSettings` | no |
| `MicrosoftSQLServerSettings` | `MicrosoftSQLServerSettings` | no |
| `IBMDb2Settings` | `IBMDb2Settings` | no |
| `ResourceIdentifier` | `string` | no |
| `DocDbSettings` | `DocDbSettings` | no |
| `RedisSettings` | `RedisSettings` | no |
| `GcpMySQLSettings` | `GcpMySQLSettings` | no |
| `TimestreamSettings` | `TimestreamSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoint` | `Endpoint` | no |

## CreateEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |
| `SnsTopicArn` | `string` | yes |
| `SourceType` | `string` | no |
| `EventCategories` | `List<string>` | no |
| `SourceIds` | `List<string>` | no |
| `Enabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## CreateFleetAdvisorCollector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectorName` | `string` | yes |
| `Description` | `string` | no |
| `ServiceAccessRoleArn` | `string` | yes |
| `S3BucketName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectorReferencedId` | `string` | no |
| `CollectorName` | `string` | no |
| `Description` | `string` | no |
| `ServiceAccessRoleArn` | `string` | no |
| `S3BucketName` | `string` | no |

## CreateInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZone` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `NetworkType` | `string` | no |
| `InstanceProfileName` | `string` | no |
| `Description` | `string` | no |
| `SubnetGroupIdentifier` | `string` | no |
| `VpcSecurityGroups` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfile` | `InstanceProfile` | no |

## CreateMigrationProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectName` | `string` | no |
| `SourceDataProviderDescriptors` | `List<DataProviderDescriptorDefinition>` | yes |
| `TargetDataProviderDescriptors` | `List<DataProviderDescriptorDefinition>` | yes |
| `InstanceProfileIdentifier` | `string` | yes |
| `TransformationRules` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `SchemaConversionApplicationAttributes` | `SCApplicationAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProject` | `MigrationProject` | no |

## CreateReplicationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigIdentifier` | `string` | yes |
| `SourceEndpointArn` | `string` | yes |
| `TargetEndpointArn` | `string` | yes |
| `ComputeConfig` | `ComputeConfig` | yes |
| `ReplicationType` | `string` | yes |
| `TableMappings` | `string` | yes |
| `ReplicationSettings` | `string` | no |
| `SupplementalSettings` | `string` | no |
| `ResourceIdentifier` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfig` | `ReplicationConfig` | no |

## CreateReplicationInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceIdentifier` | `string` | yes |
| `AllocatedStorage` | `integer` | no |
| `ReplicationInstanceClass` | `string` | yes |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `AvailabilityZone` | `string` | no |
| `ReplicationSubnetGroupIdentifier` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `EngineVersion` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `DnsNameServers` | `string` | no |
| `ResourceIdentifier` | `string` | no |
| `NetworkType` | `string` | no |
| `KerberosAuthenticationSettings` | `KerberosAuthenticationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstance` | `ReplicationInstance` | no |

## CreateReplicationSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationSubnetGroupIdentifier` | `string` | yes |
| `ReplicationSubnetGroupDescription` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationSubnetGroup` | `ReplicationSubnetGroup` | no |

## CreateReplicationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskIdentifier` | `string` | yes |
| `SourceEndpointArn` | `string` | yes |
| `TargetEndpointArn` | `string` | yes |
| `ReplicationInstanceArn` | `string` | yes |
| `MigrationType` | `string` | yes |
| `TableMappings` | `string` | yes |
| `ReplicationTaskSettings` | `string` | no |
| `CdcStartTime` | `timestamp` | no |
| `CdcStartPosition` | `string` | no |
| `CdcStopPosition` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `TaskData` | `string` | no |
| `ResourceIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTask` | `ReplicationTask` | no |

## DeleteCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `Certificate` | no |

## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |
| `ReplicationInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | no |

## DeleteDataMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigrationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigration` | `DataMigration` | no |

## DeleteDataProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataProviderIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataProvider` | `DataProvider` | no |

## DeleteEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoint` | `Endpoint` | no |

## DeleteEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## DeleteFleetAdvisorCollector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectorReferencedId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFleetAdvisorDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseIds` | `List<string>` | no |

## DeleteInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfile` | `InstanceProfile` | no |

## DeleteMigrationProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProject` | `MigrationProject` | no |

## DeleteReplicationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfig` | `ReplicationConfig` | no |

## DeleteReplicationInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstance` | `ReplicationInstance` | no |

## DeleteReplicationSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationSubnetGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReplicationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTask` | `ReplicationTask` | no |

## DeleteReplicationTaskAssessmentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskAssessmentRunArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskAssessmentRun` | `ReplicationTaskAssessmentRun` | no |

## DescribeAccountAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountQuotas` | `List<AccountQuota>` | no |
| `UniqueAccountIdentifier` | `string` | no |

## DescribeApplicableIndividualAssessments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | no |
| `ReplicationInstanceArn` | `string` | no |
| `ReplicationConfigArn` | `string` | no |
| `SourceEngineName` | `string` | no |
| `TargetEngineName` | `string` | no |
| `MigrationType` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndividualAssessmentNames` | `List<string>` | no |
| `Marker` | `string` | no |

## DescribeCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Certificates` | `List<Certificate>` | no |

## DescribeConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Connections` | `List<Connection>` | no |

## DescribeConversionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | no |
| `ConversionConfiguration` | `string` | no |

## DescribeDataMigrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `WithoutSettings` | `boolean` | no |
| `WithoutStatistics` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigrations` | `List<DataMigration>` | no |
| `Marker` | `string` | no |

## DescribeDataProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DataProviders` | `List<DataProvider>` | no |

## DescribeEndpointSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineName` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `EndpointSettings` | `List<EndpointSetting>` | no |

## DescribeEndpointTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `SupportedEndpointTypes` | `List<SupportedEndpointType>` | no |

## DescribeEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Endpoints` | `List<Endpoint>` | no |

## DescribeEngineVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineVersions` | `List<EngineVersion>` | no |
| `Marker` | `string` | no |

## DescribeEventCategories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceType` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventCategoryGroupList` | `List<EventCategoryGroup>` | no |

## DescribeEventSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

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
| `EventCategories` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Events` | `List<Event>` | no |

## DescribeExtensionPackAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Requests` | `List<SchemaConversionRequest>` | no |

## DescribeFleetAdvisorCollectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Collectors` | `List<CollectorResponse>` | no |
| `NextToken` | `string` | no |

## DescribeFleetAdvisorDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Databases` | `List<DatabaseResponse>` | no |
| `NextToken` | `string` | no |

## DescribeFleetAdvisorLsaAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Analysis` | `List<FleetAdvisorLsaAnalysisResponse>` | no |
| `NextToken` | `string` | no |

## DescribeFleetAdvisorSchemaObjectSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetAdvisorSchemaObjects` | `List<FleetAdvisorSchemaObjectResponse>` | no |
| `NextToken` | `string` | no |

## DescribeFleetAdvisorSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetAdvisorSchemas` | `List<SchemaResponse>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `InstanceProfiles` | `List<InstanceProfile>` | no |

## DescribeMetadataModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SelectionRules` | `string` | yes |
| `MigrationProjectIdentifier` | `string` | yes |
| `Origin` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetadataModelName` | `string` | no |
| `MetadataModelType` | `string` | no |
| `TargetMetadataModels` | `List<MetadataModelReference>` | no |
| `Definition` | `string` | no |

## DescribeMetadataModelAssessments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Requests` | `List<SchemaConversionRequest>` | no |

## DescribeMetadataModelChildren

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SelectionRules` | `string` | yes |
| `MigrationProjectIdentifier` | `string` | yes |
| `Origin` | `string` | yes |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MetadataModelChildren` | `List<MetadataModelReference>` | no |

## DescribeMetadataModelConversions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Requests` | `List<SchemaConversionRequest>` | no |

## DescribeMetadataModelCreations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |
| `MigrationProjectIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Requests` | `List<SchemaConversionRequest>` | no |

## DescribeMetadataModelExportsAsScript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Requests` | `List<SchemaConversionRequest>` | no |

## DescribeMetadataModelExportsToTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Requests` | `List<SchemaConversionRequest>` | no |

## DescribeMetadataModelImports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Requests` | `List<SchemaConversionRequest>` | no |

## DescribeMigrationProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MigrationProjects` | `List<MigrationProject>` | no |

## DescribeOrderableReplicationInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrderableReplicationInstances` | `List<OrderableReplicationInstance>` | no |
| `Marker` | `string` | no |

## DescribePendingMaintenanceActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceArn` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PendingMaintenanceActions` | `List<ResourcePendingMaintenanceActions>` | no |
| `Marker` | `string` | no |

## DescribeRecommendationLimitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Limitations` | `List<Limitation>` | no |

## DescribeRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Recommendations` | `List<Recommendation>` | no |

## DescribeRefreshSchemasStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RefreshSchemasStatus` | `RefreshSchemasStatus` | no |

## DescribeReplicationConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReplicationConfigs` | `List<ReplicationConfig>` | no |

## DescribeReplicationInstanceTaskLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceArn` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceArn` | `string` | no |
| `ReplicationInstanceTaskLogs` | `List<ReplicationInstanceTaskLog>` | no |
| `Marker` | `string` | no |

## DescribeReplicationInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReplicationInstances` | `List<ReplicationInstance>` | no |

## DescribeReplicationSubnetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReplicationSubnetGroups` | `List<ReplicationSubnetGroup>` | no |

## DescribeReplicationTableStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigArn` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigArn` | `string` | no |
| `Marker` | `string` | no |
| `ReplicationTableStatistics` | `List<TableStatistics>` | no |

## DescribeReplicationTaskAssessmentResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `BucketName` | `string` | no |
| `ReplicationTaskAssessmentResults` | `List<ReplicationTaskAssessmentResult>` | no |

## DescribeReplicationTaskAssessmentRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReplicationTaskAssessmentRuns` | `List<ReplicationTaskAssessmentRun>` | no |

## DescribeReplicationTaskIndividualAssessments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReplicationTaskIndividualAssessments` | `List<ReplicationTaskIndividualAssessment>` | no |

## DescribeReplicationTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `WithoutSettings` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReplicationTasks` | `List<ReplicationTask>` | no |

## DescribeReplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Replications` | `List<Replication>` | no |

## DescribeSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Schemas` | `List<string>` | no |

## DescribeTableStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | no |
| `TableStatistics` | `List<TableStatistics>` | no |
| `Marker` | `string` | no |

## ExportMetadataModelAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `SelectionRules` | `string` | yes |
| `FileName` | `string` | no |
| `AssessmentReportTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PdfReport` | `ExportMetadataModelAssessmentResultEntry` | no |
| `CsvReport` | `ExportMetadataModelAssessmentResultEntry` | no |

## GetTargetSelectionRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `SelectionRules` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetSelectionRules` | `string` | no |

## ImportCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateIdentifier` | `string` | yes |
| `CertificatePem` | `string` | no |
| `CertificateWallet` | `blob` | no |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `Certificate` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `ResourceArnList` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## ModifyConversionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `ConversionConfiguration` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | no |

## ModifyDataMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigrationIdentifier` | `string` | yes |
| `DataMigrationName` | `string` | no |
| `EnableCloudwatchLogs` | `boolean` | no |
| `ServiceAccessRoleArn` | `string` | no |
| `DataMigrationType` | `string` | no |
| `SourceDataSettings` | `List<SourceDataSetting>` | no |
| `TargetDataSettings` | `List<TargetDataSetting>` | no |
| `NumberOfJobs` | `integer` | no |
| `SelectionRules` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigration` | `DataMigration` | no |

## ModifyDataProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataProviderIdentifier` | `string` | yes |
| `DataProviderName` | `string` | no |
| `Description` | `string` | no |
| `Engine` | `string` | no |
| `Virtual` | `boolean` | no |
| `ExactSettings` | `boolean` | no |
| `Settings` | `DataProviderSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataProvider` | `DataProvider` | no |

## ModifyEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |
| `EndpointIdentifier` | `string` | no |
| `EndpointType` | `string` | no |
| `EngineName` | `string` | no |
| `Username` | `string` | no |
| `Password` | `string` | no |
| `ServerName` | `string` | no |
| `Port` | `integer` | no |
| `DatabaseName` | `string` | no |
| `ExtraConnectionAttributes` | `string` | no |
| `CertificateArn` | `string` | no |
| `SslMode` | `string` | no |
| `ServiceAccessRoleArn` | `string` | no |
| `ExternalTableDefinition` | `string` | no |
| `DynamoDbSettings` | `DynamoDbSettings` | no |
| `S3Settings` | `S3Settings` | no |
| `DmsTransferSettings` | `DmsTransferSettings` | no |
| `MongoDbSettings` | `MongoDbSettings` | no |
| `KinesisSettings` | `KinesisSettings` | no |
| `KafkaSettings` | `KafkaSettings` | no |
| `ElasticsearchSettings` | `ElasticsearchSettings` | no |
| `NeptuneSettings` | `NeptuneSettings` | no |
| `RedshiftSettings` | `RedshiftSettings` | no |
| `PostgreSQLSettings` | `PostgreSQLSettings` | no |
| `MySQLSettings` | `MySQLSettings` | no |
| `OracleSettings` | `OracleSettings` | no |
| `SybaseSettings` | `SybaseSettings` | no |
| `MicrosoftSQLServerSettings` | `MicrosoftSQLServerSettings` | no |
| `IBMDb2Settings` | `IBMDb2Settings` | no |
| `DocDbSettings` | `DocDbSettings` | no |
| `RedisSettings` | `RedisSettings` | no |
| `ExactSettings` | `boolean` | no |
| `GcpMySQLSettings` | `GcpMySQLSettings` | no |
| `TimestreamSettings` | `TimestreamSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoint` | `Endpoint` | no |

## ModifyEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |
| `SnsTopicArn` | `string` | no |
| `SourceType` | `string` | no |
| `EventCategories` | `List<string>` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## ModifyInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfileIdentifier` | `string` | yes |
| `AvailabilityZone` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `NetworkType` | `string` | no |
| `InstanceProfileName` | `string` | no |
| `Description` | `string` | no |
| `SubnetGroupIdentifier` | `string` | no |
| `VpcSecurityGroups` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceProfile` | `InstanceProfile` | no |

## ModifyMigrationProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `MigrationProjectName` | `string` | no |
| `SourceDataProviderDescriptors` | `List<DataProviderDescriptorDefinition>` | no |
| `TargetDataProviderDescriptors` | `List<DataProviderDescriptorDefinition>` | no |
| `InstanceProfileIdentifier` | `string` | no |
| `TransformationRules` | `string` | no |
| `Description` | `string` | no |
| `SchemaConversionApplicationAttributes` | `SCApplicationAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProject` | `MigrationProject` | no |

## ModifyReplicationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigArn` | `string` | yes |
| `ReplicationConfigIdentifier` | `string` | no |
| `ReplicationType` | `string` | no |
| `TableMappings` | `string` | no |
| `ReplicationSettings` | `string` | no |
| `SupplementalSettings` | `string` | no |
| `ComputeConfig` | `ComputeConfig` | no |
| `SourceEndpointArn` | `string` | no |
| `TargetEndpointArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfig` | `ReplicationConfig` | no |

## ModifyReplicationInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceArn` | `string` | yes |
| `AllocatedStorage` | `integer` | no |
| `ApplyImmediately` | `boolean` | no |
| `ReplicationInstanceClass` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `EngineVersion` | `string` | no |
| `AllowMajorVersionUpgrade` | `boolean` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `ReplicationInstanceIdentifier` | `string` | no |
| `NetworkType` | `string` | no |
| `KerberosAuthenticationSettings` | `KerberosAuthenticationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstance` | `ReplicationInstance` | no |

## ModifyReplicationSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationSubnetGroupIdentifier` | `string` | yes |
| `ReplicationSubnetGroupDescription` | `string` | no |
| `SubnetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationSubnetGroup` | `ReplicationSubnetGroup` | no |

## ModifyReplicationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |
| `ReplicationTaskIdentifier` | `string` | no |
| `MigrationType` | `string` | no |
| `TableMappings` | `string` | no |
| `ReplicationTaskSettings` | `string` | no |
| `CdcStartTime` | `timestamp` | no |
| `CdcStartPosition` | `string` | no |
| `CdcStopPosition` | `string` | no |
| `TaskData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTask` | `ReplicationTask` | no |

## MoveReplicationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |
| `TargetReplicationInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTask` | `ReplicationTask` | no |

## RebootReplicationInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceArn` | `string` | yes |
| `ForceFailover` | `boolean` | no |
| `ForcePlannedFailover` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstance` | `ReplicationInstance` | no |

## RefreshSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |
| `ReplicationInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RefreshSchemasStatus` | `RefreshSchemasStatus` | no |

## ReloadReplicationTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigArn` | `string` | yes |
| `TablesToReload` | `List<TableToReload>` | yes |
| `ReloadOption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigArn` | `string` | no |

## ReloadTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |
| `TablesToReload` | `List<TableToReload>` | yes |
| `ReloadOption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | no |

## RemoveTagsFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RunFleetAdvisorLsaAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LsaAnalysisId` | `string` | no |
| `Status` | `string` | no |

## StartDataMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigrationIdentifier` | `string` | yes |
| `StartType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigration` | `DataMigration` | no |

## StartExtensionPackAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestIdentifier` | `string` | no |

## StartMetadataModelAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `SelectionRules` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestIdentifier` | `string` | no |

## StartMetadataModelConversion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `SelectionRules` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestIdentifier` | `string` | no |

## StartMetadataModelCreation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `SelectionRules` | `string` | yes |
| `MetadataModelName` | `string` | yes |
| `Properties` | `MetadataModelProperties` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestIdentifier` | `string` | no |

## StartMetadataModelExportAsScript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `SelectionRules` | `string` | yes |
| `Origin` | `string` | yes |
| `FileName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestIdentifier` | `string` | no |

## StartMetadataModelExportToTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `SelectionRules` | `string` | yes |
| `OverwriteExtensionPack` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestIdentifier` | `string` | no |

## StartMetadataModelImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationProjectIdentifier` | `string` | yes |
| `SelectionRules` | `string` | yes |
| `Origin` | `string` | yes |
| `Refresh` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestIdentifier` | `string` | no |

## StartRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseId` | `string` | yes |
| `Settings` | `RecommendationSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigArn` | `string` | yes |
| `StartReplicationType` | `string` | yes |
| `PremigrationAssessmentSettings` | `string` | no |
| `CdcStartTime` | `timestamp` | no |
| `CdcStartPosition` | `string` | no |
| `CdcStopPosition` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Replication` | `Replication` | no |

## StartReplicationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |
| `StartReplicationTaskType` | `string` | yes |
| `CdcStartTime` | `timestamp` | no |
| `CdcStartPosition` | `string` | no |
| `CdcStopPosition` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTask` | `ReplicationTask` | no |

## StartReplicationTaskAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTask` | `ReplicationTask` | no |

## StartReplicationTaskAssessmentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |
| `ServiceAccessRoleArn` | `string` | yes |
| `ResultLocationBucket` | `string` | yes |
| `ResultLocationFolder` | `string` | no |
| `ResultEncryptionMode` | `string` | no |
| `ResultKmsKeyArn` | `string` | no |
| `AssessmentRunName` | `string` | yes |
| `IncludeOnly` | `List<string>` | no |
| `Exclude` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskAssessmentRun` | `ReplicationTaskAssessmentRun` | no |

## StopDataMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigrationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataMigration` | `DataMigration` | no |

## StopReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationConfigArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Replication` | `Replication` | no |

## StopReplicationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTaskArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationTask` | `ReplicationTask` | no |

## TestConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationInstanceArn` | `string` | yes |
| `EndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | no |

## UpdateSubscriptionsToEventBridge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForceMove` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Result` | `string` | no |

