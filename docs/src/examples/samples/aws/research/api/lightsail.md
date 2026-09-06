# Amazon Lightsail

API version: 2016-11-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lightsail/2016-11-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AllocateStaticIp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `staticIpName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## AttachCertificateToDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | yes |
| `certificateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## AttachDisk

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskName` | `string` | yes |
| `instanceName` | `string` | yes |
| `diskPath` | `string` | yes |
| `autoMounting` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## AttachInstancesToLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |
| `instanceNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## AttachLoadBalancerTlsCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |
| `certificateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## AttachStaticIp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `staticIpName` | `string` | yes |
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CloseInstancePublicPorts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portInfo` | `PortInfo` | yes |
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## CopySnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceSnapshotName` | `string` | no |
| `sourceResourceName` | `string` | no |
| `restoreDate` | `string` | no |
| `useLatestRestorableAutoSnapshot` | `boolean` | no |
| `targetSnapshotName` | `string` | yes |
| `sourceRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | yes |
| `bundleId` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `enableObjectVersioning` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucket` | `Bucket` | no |
| `operations` | `List<Operation>` | no |

## CreateBucketAccessKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessKey` | `AccessKey` | no |
| `operations` | `List<Operation>` | no |

## CreateCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateName` | `string` | yes |
| `domainName` | `string` | yes |
| `subjectAlternativeNames` | `List<string>` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificate` | `CertificateSummary` | no |
| `operations` | `List<Operation>` | no |

## CreateCloudFormationStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instances` | `List<InstanceEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateContactMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protocol` | `string` | yes |
| `contactEndpoint` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateContainerService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |
| `power` | `string` | yes |
| `scale` | `integer` | yes |
| `tags` | `List<Tag>` | no |
| `publicDomainNames` | `Map<List<string>>` | no |
| `deployment` | `ContainerServiceDeploymentRequest` | no |
| `privateRegistryAccess` | `PrivateRegistryAccessRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerService` | `ContainerService` | no |

## CreateContainerServiceDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |
| `containers` | `Map<Container>` | no |
| `publicEndpoint` | `EndpointRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerService` | `ContainerService` | no |

## CreateContainerServiceRegistryLogin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryLogin` | `ContainerServiceRegistryLogin` | no |

## CreateDisk

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskName` | `string` | yes |
| `availabilityZone` | `string` | yes |
| `sizeInGb` | `integer` | yes |
| `tags` | `List<Tag>` | no |
| `addOns` | `List<AddOnRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateDiskFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskName` | `string` | yes |
| `diskSnapshotName` | `string` | no |
| `availabilityZone` | `string` | yes |
| `sizeInGb` | `integer` | yes |
| `tags` | `List<Tag>` | no |
| `addOns` | `List<AddOnRequest>` | no |
| `sourceDiskName` | `string` | no |
| `restoreDate` | `string` | no |
| `useLatestRestorableAutoSnapshot` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateDiskSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskName` | `string` | no |
| `diskSnapshotName` | `string` | yes |
| `instanceName` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | yes |
| `origin` | `InputOrigin` | yes |
| `defaultCacheBehavior` | `CacheBehavior` | yes |
| `cacheBehaviorSettings` | `CacheSettings` | no |
| `cacheBehaviors` | `List<CacheBehaviorPerPath>` | no |
| `bundleId` | `string` | yes |
| `ipAddressType` | `string` | no |
| `tags` | `List<Tag>` | no |
| `certificateName` | `string` | no |
| `viewerMinimumTlsProtocolVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distribution` | `LightsailDistribution` | no |
| `operation` | `Operation` | no |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## CreateDomainEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainEntry` | `DomainEntry` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## CreateGUISessionAccessDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | no |
| `status` | `string` | no |
| `percentageComplete` | `integer` | no |
| `failureReason` | `string` | no |
| `sessions` | `List<Session>` | no |

## CreateInstanceSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceSnapshotName` | `string` | yes |
| `instanceName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceNames` | `List<string>` | yes |
| `availabilityZone` | `string` | yes |
| `customImageName` | `string` | no |
| `blueprintId` | `string` | yes |
| `bundleId` | `string` | yes |
| `userData` | `string` | no |
| `keyPairName` | `string` | no |
| `tags` | `List<Tag>` | no |
| `addOns` | `List<AddOnRequest>` | no |
| `ipAddressType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateInstancesFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceNames` | `List<string>` | yes |
| `attachedDiskMapping` | `Map<List<DiskMap>>` | no |
| `availabilityZone` | `string` | yes |
| `instanceSnapshotName` | `string` | no |
| `bundleId` | `string` | yes |
| `userData` | `string` | no |
| `keyPairName` | `string` | no |
| `tags` | `List<Tag>` | no |
| `addOns` | `List<AddOnRequest>` | no |
| `ipAddressType` | `string` | no |
| `sourceInstanceName` | `string` | no |
| `restoreDate` | `string` | no |
| `useLatestRestorableAutoSnapshot` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPairName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPair` | `KeyPair` | no |
| `publicKeyBase64` | `string` | no |
| `privateKeyBase64` | `string` | no |
| `operation` | `Operation` | no |

## CreateLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |
| `instancePort` | `integer` | yes |
| `healthCheckPath` | `string` | no |
| `certificateName` | `string` | no |
| `certificateDomainName` | `string` | no |
| `certificateAlternativeNames` | `List<string>` | no |
| `tags` | `List<Tag>` | no |
| `ipAddressType` | `string` | no |
| `tlsPolicyName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateLoadBalancerTlsCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |
| `certificateName` | `string` | yes |
| `certificateDomainName` | `string` | yes |
| `certificateAlternativeNames` | `List<string>` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateRelationalDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `availabilityZone` | `string` | no |
| `relationalDatabaseBlueprintId` | `string` | yes |
| `relationalDatabaseBundleId` | `string` | yes |
| `masterDatabaseName` | `string` | yes |
| `masterUsername` | `string` | yes |
| `masterUserPassword` | `string` | no |
| `preferredBackupWindow` | `string` | no |
| `preferredMaintenanceWindow` | `string` | no |
| `publiclyAccessible` | `boolean` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateRelationalDatabaseFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `availabilityZone` | `string` | no |
| `publiclyAccessible` | `boolean` | no |
| `relationalDatabaseSnapshotName` | `string` | no |
| `relationalDatabaseBundleId` | `string` | no |
| `sourceRelationalDatabaseName` | `string` | no |
| `restoreTime` | `timestamp` | no |
| `useLatestRestorableTime` | `boolean` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## CreateRelationalDatabaseSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `relationalDatabaseSnapshotName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteAlarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alarmName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteAutoSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |
| `date` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | yes |
| `forceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteBucketAccessKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | yes |
| `accessKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteContactMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protocol` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteContainerImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |
| `image` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContainerService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDisk

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskName` | `string` | yes |
| `forceDeleteAddOns` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteDiskSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## DeleteDomainEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainEntry` | `DomainEntry` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## DeleteInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |
| `forceDeleteAddOns` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteInstanceSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPairName` | `string` | yes |
| `expectedFingerprint` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## DeleteKnownHostKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteLoadBalancerTlsCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |
| `certificateName` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteRelationalDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `skipFinalSnapshot` | `boolean` | no |
| `finalRelationalDatabaseSnapshotName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DeleteRelationalDatabaseSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DetachCertificateFromDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## DetachDisk

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DetachInstancesFromLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |
| `instanceNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DetachStaticIp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `staticIpName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DisableAddOn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addOnType` | `string` | yes |
| `resourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## DownloadDefaultKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `publicKeyBase64` | `string` | no |
| `privateKeyBase64` | `string` | no |
| `createdAt` | `timestamp` | no |

## EnableAddOn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |
| `addOnRequest` | `AddOnRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## ExportSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## GetActiveNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activeNames` | `List<string>` | no |
| `nextPageToken` | `string` | no |

## GetAlarms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alarmName` | `string` | no |
| `pageToken` | `string` | no |
| `monitoredResourceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alarms` | `List<Alarm>` | no |
| `nextPageToken` | `string` | no |

## GetAutoSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | no |
| `resourceType` | `string` | no |
| `autoSnapshots` | `List<AutoSnapshotDetails>` | no |

## GetBlueprints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeInactive` | `boolean` | no |
| `pageToken` | `string` | no |
| `appCategory` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprints` | `List<Blueprint>` | no |
| `nextPageToken` | `string` | no |

## GetBucketAccessKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessKeys` | `List<AccessKey>` | no |

## GetBucketBundles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeInactive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundles` | `List<BucketBundle>` | no |

## GetBucketMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | yes |
| `metricName` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `period` | `integer` | yes |
| `statistics` | `List<string>` | yes |
| `unit` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricData` | `List<MetricDatapoint>` | no |

## GetBuckets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | no |
| `pageToken` | `string` | no |
| `includeConnectedResources` | `boolean` | no |
| `includeCors` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `buckets` | `List<Bucket>` | no |
| `nextPageToken` | `string` | no |
| `accountLevelBpaSync` | `AccountLevelBpaSync` | no |

## GetBundles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeInactive` | `boolean` | no |
| `pageToken` | `string` | no |
| `appCategory` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundles` | `List<Bundle>` | no |
| `nextPageToken` | `string` | no |

## GetCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateStatuses` | `List<string>` | no |
| `includeCertificateDetails` | `boolean` | no |
| `certificateName` | `string` | no |
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificates` | `List<CertificateSummary>` | no |
| `nextPageToken` | `string` | no |

## GetCloudFormationStackRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudFormationStackRecords` | `List<CloudFormationStackRecord>` | no |
| `nextPageToken` | `string` | no |

## GetContactMethods

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protocols` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactMethods` | `List<ContactMethod>` | no |

## GetContainerAPIMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metadata` | `List<Map<string>>` | no |

## GetContainerImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerImages` | `List<ContainerImage>` | no |

## GetContainerLog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |
| `containerName` | `string` | yes |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `filterPattern` | `string` | no |
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logEvents` | `List<ContainerServiceLogEvent>` | no |
| `nextPageToken` | `string` | no |

## GetContainerServiceDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployments` | `List<ContainerServiceDeployment>` | no |

## GetContainerServiceMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |
| `metricName` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `period` | `integer` | yes |
| `statistics` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricData` | `List<MetricDatapoint>` | no |

## GetContainerServicePowers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `powers` | `List<ContainerServicePower>` | no |

## GetContainerServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerServices` | `List<ContainerService>` | no |

## GetCostEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcesBudgetEstimate` | `List<ResourceBudgetEstimate>` | no |

## GetDisk

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `disk` | `Disk` | no |

## GetDiskSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskSnapshot` | `DiskSnapshot` | no |

## GetDiskSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `diskSnapshots` | `List<DiskSnapshot>` | no |
| `nextPageToken` | `string` | no |

## GetDisks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `disks` | `List<Disk>` | no |
| `nextPageToken` | `string` | no |

## GetDistributionBundles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundles` | `List<DistributionBundle>` | no |

## GetDistributionLatestCacheReset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `createTime` | `timestamp` | no |

## GetDistributionMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | yes |
| `metricName` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `period` | `integer` | yes |
| `unit` | `string` | yes |
| `statistics` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricData` | `List<MetricDatapoint>` | no |

## GetDistributions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | no |
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributions` | `List<LightsailDistribution>` | no |
| `nextPageToken` | `string` | no |

## GetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `Domain` | no |

## GetDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domains` | `List<Domain>` | no |
| `nextPageToken` | `string` | no |

## GetExportSnapshotRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportSnapshotRecords` | `List<ExportSnapshotRecord>` | no |
| `nextPageToken` | `string` | no |

## GetInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instance` | `Instance` | no |

## GetInstanceAccessDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |
| `protocol` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessDetails` | `InstanceAccessDetails` | no |

## GetInstanceMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |
| `metricName` | `string` | yes |
| `period` | `integer` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `unit` | `string` | yes |
| `statistics` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricData` | `List<MetricDatapoint>` | no |

## GetInstancePortStates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portStates` | `List<InstancePortState>` | no |

## GetInstanceSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceSnapshot` | `InstanceSnapshot` | no |

## GetInstanceSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceSnapshots` | `List<InstanceSnapshot>` | no |
| `nextPageToken` | `string` | no |

## GetInstanceState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `state` | `InstanceState` | no |

## GetInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instances` | `List<Instance>` | no |
| `nextPageToken` | `string` | no |

## GetKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPairName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPair` | `KeyPair` | no |

## GetKeyPairs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |
| `includeDefaultKeyPair` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPairs` | `List<KeyPair>` | no |
| `nextPageToken` | `string` | no |

## GetLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancer` | `LoadBalancer` | no |

## GetLoadBalancerMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |
| `metricName` | `string` | yes |
| `period` | `integer` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `unit` | `string` | yes |
| `statistics` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricData` | `List<MetricDatapoint>` | no |

## GetLoadBalancerTlsCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tlsCertificates` | `List<LoadBalancerTlsCertificate>` | no |

## GetLoadBalancerTlsPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tlsPolicies` | `List<LoadBalancerTlsPolicy>` | no |
| `nextPageToken` | `string` | no |

## GetLoadBalancers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancers` | `List<LoadBalancer>` | no |
| `nextPageToken` | `string` | no |

## GetOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## GetOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |
| `nextPageToken` | `string` | no |

## GetOperationsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |
| `nextPageCount` | `string` | no |
| `nextPageToken` | `string` | no |

## GetProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileType` | `string` | yes |
| `partner` | `PartnerInfo` | no |

## GetRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeAvailabilityZones` | `boolean` | no |
| `includeRelationalDatabaseAvailabilityZones` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `regions` | `List<Region>` | no |

## GetRelationalDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabase` | `RelationalDatabase` | no |

## GetRelationalDatabaseBlueprints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `blueprints` | `List<RelationalDatabaseBlueprint>` | no |
| `nextPageToken` | `string` | no |

## GetRelationalDatabaseBundles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |
| `includeInactive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundles` | `List<RelationalDatabaseBundle>` | no |
| `nextPageToken` | `string` | no |

## GetRelationalDatabaseEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `durationInMinutes` | `integer` | no |
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseEvents` | `List<RelationalDatabaseEvent>` | no |
| `nextPageToken` | `string` | no |

## GetRelationalDatabaseLogEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `logStreamName` | `string` | yes |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `startFromHead` | `boolean` | no |
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceLogEvents` | `List<LogEvent>` | no |
| `nextBackwardToken` | `string` | no |
| `nextForwardToken` | `string` | no |

## GetRelationalDatabaseLogStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logStreams` | `List<string>` | no |

## GetRelationalDatabaseMasterUserPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `passwordVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `masterUserPassword` | `string` | no |
| `createdAt` | `timestamp` | no |

## GetRelationalDatabaseMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `metricName` | `string` | yes |
| `period` | `integer` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `unit` | `string` | yes |
| `statistics` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricData` | `List<MetricDatapoint>` | no |

## GetRelationalDatabaseParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `parameters` | `List<RelationalDatabaseParameter>` | no |
| `nextPageToken` | `string` | no |

## GetRelationalDatabaseSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseSnapshot` | `RelationalDatabaseSnapshot` | no |

## GetRelationalDatabaseSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseSnapshots` | `List<RelationalDatabaseSnapshot>` | no |
| `nextPageToken` | `string` | no |

## GetRelationalDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabases` | `List<RelationalDatabase>` | no |
| `nextPageToken` | `string` | no |

## GetSetupHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `setupHistory` | `List<SetupHistory>` | no |
| `nextPageToken` | `string` | no |

## GetStaticIp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `staticIpName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `staticIp` | `StaticIp` | no |

## GetStaticIps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `staticIps` | `List<StaticIp>` | no |
| `nextPageToken` | `string` | no |

## ImportKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPairName` | `string` | yes |
| `publicKeyBase64` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## IsVpcPeered

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `isPeered` | `boolean` | no |

## OpenInstancePublicPorts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portInfo` | `PortInfo` | yes |
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## PeerVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## PutAlarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alarmName` | `string` | yes |
| `metricName` | `string` | yes |
| `monitoredResourceName` | `string` | yes |
| `comparisonOperator` | `string` | yes |
| `threshold` | `double` | yes |
| `evaluationPeriods` | `integer` | yes |
| `datapointsToAlarm` | `integer` | no |
| `treatMissingData` | `string` | no |
| `contactProtocols` | `List<string>` | no |
| `notificationTriggers` | `List<string>` | no |
| `notificationEnabled` | `boolean` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## PutInstancePublicPorts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portInfos` | `List<PortInfo>` | yes |
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## RebootInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## RebootRelationalDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## RegisterContainerImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |
| `label` | `string` | yes |
| `digest` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerImage` | `ContainerImage` | no |

## ReleaseStaticIp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `staticIpName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## ResetDistributionCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `createTime` | `timestamp` | no |
| `operation` | `Operation` | no |

## SendContactMethodVerification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protocol` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## SetIpAddressType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |
| `resourceName` | `string` | yes |
| `ipAddressType` | `string` | yes |
| `acceptBundleUpdate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## SetResourceAccessForBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |
| `bucketName` | `string` | yes |
| `access` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## SetupInstanceHttps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |
| `emailAddress` | `string` | yes |
| `domainNames` | `List<string>` | yes |
| `certificateProvider` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## StartGUISession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## StartInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## StartRelationalDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## StopGUISession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## StopInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## StopRelationalDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `relationalDatabaseSnapshotName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |
| `resourceArn` | `string` | no |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## TestAlarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alarmName` | `string` | yes |
| `state` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## UnpeerVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceName` | `string` | yes |
| `resourceArn` | `string` | no |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## UpdateBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | yes |
| `accessRules` | `AccessRules` | no |
| `versioning` | `string` | no |
| `readonlyAccessAccounts` | `List<string>` | no |
| `accessLogConfig` | `BucketAccessLogConfig` | no |
| `cors` | `BucketCorsConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucket` | `Bucket` | no |
| `operations` | `List<Operation>` | no |

## UpdateBucketBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | yes |
| `bundleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## UpdateContainerService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |
| `power` | `string` | no |
| `scale` | `integer` | no |
| `isDisabled` | `boolean` | no |
| `publicDomainNames` | `Map<List<string>>` | no |
| `privateRegistryAccess` | `PrivateRegistryAccessRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerService` | `ContainerService` | no |

## UpdateDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | yes |
| `origin` | `InputOrigin` | no |
| `defaultCacheBehavior` | `CacheBehavior` | no |
| `cacheBehaviorSettings` | `CacheSettings` | no |
| `cacheBehaviors` | `List<CacheBehaviorPerPath>` | no |
| `isEnabled` | `boolean` | no |
| `viewerMinimumTlsProtocolVersion` | `string` | no |
| `certificateName` | `string` | no |
| `useDefaultCertificate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## UpdateDistributionBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionName` | `string` | no |
| `bundleId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## UpdateDomainEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainEntry` | `DomainEntry` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## UpdateInstanceMetadataOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |
| `httpTokens` | `string` | no |
| `httpEndpoint` | `string` | no |
| `httpPutResponseHopLimit` | `integer` | no |
| `httpProtocolIpv6` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operation` | `Operation` | no |

## UpdateLoadBalancerAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadBalancerName` | `string` | yes |
| `attributeName` | `string` | yes |
| `attributeValue` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## UpdateRelationalDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `masterUserPassword` | `string` | no |
| `rotateMasterUserPassword` | `boolean` | no |
| `preferredBackupWindow` | `string` | no |
| `preferredMaintenanceWindow` | `string` | no |
| `enableBackupRetention` | `boolean` | no |
| `disableBackupRetention` | `boolean` | no |
| `publiclyAccessible` | `boolean` | no |
| `applyImmediately` | `boolean` | no |
| `caCertificateIdentifier` | `string` | no |
| `relationalDatabaseBlueprintId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

## UpdateRelationalDatabaseParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationalDatabaseName` | `string` | yes |
| `parameters` | `List<RelationalDatabaseParameter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operations` | `List<Operation>` | no |

