# FinSpace User Environment Management service

API version: 2021-03-12. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/finspace/2021-03-12/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `federationMode` | `string` | no |
| `federationParameters` | `FederationParameters` | no |
| `superuserParameters` | `SuperuserParameters` | no |
| `dataBundles` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | no |
| `environmentArn` | `string` | no |
| `environmentUrl` | `string` | no |

## CreateKxChangeset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `changeRequests` | `List<ChangeRequest>` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `changesetId` | `string` | no |
| `databaseName` | `string` | no |
| `environmentId` | `string` | no |
| `changeRequests` | `List<ChangeRequest>` | no |
| `createdTimestamp` | `timestamp` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `status` | `string` | no |
| `errorInfo` | `ErrorInfo` | no |

## CreateKxCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `clusterName` | `string` | yes |
| `clusterType` | `string` | yes |
| `tickerplantLogConfiguration` | `TickerplantLogConfiguration` | no |
| `databases` | `List<KxDatabaseConfiguration>` | no |
| `cacheStorageConfigurations` | `List<KxCacheStorageConfiguration>` | no |
| `autoScalingConfiguration` | `AutoScalingConfiguration` | no |
| `clusterDescription` | `string` | no |
| `capacityConfiguration` | `CapacityConfiguration` | no |
| `releaseLabel` | `string` | yes |
| `vpcConfiguration` | `VpcConfiguration` | yes |
| `initializationScript` | `string` | no |
| `commandLineArguments` | `List<KxCommandLineArgument>` | no |
| `code` | `CodeConfiguration` | no |
| `executionRole` | `string` | no |
| `savedownStorageConfiguration` | `KxSavedownStorageConfiguration` | no |
| `azMode` | `string` | yes |
| `availabilityZoneId` | `string` | no |
| `tags` | `Map<string>` | no |
| `scalingGroupConfiguration` | `KxScalingGroupConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `clusterName` | `string` | no |
| `clusterType` | `string` | no |
| `tickerplantLogConfiguration` | `TickerplantLogConfiguration` | no |
| `volumes` | `List<Volume>` | no |
| `databases` | `List<KxDatabaseConfiguration>` | no |
| `cacheStorageConfigurations` | `List<KxCacheStorageConfiguration>` | no |
| `autoScalingConfiguration` | `AutoScalingConfiguration` | no |
| `clusterDescription` | `string` | no |
| `capacityConfiguration` | `CapacityConfiguration` | no |
| `releaseLabel` | `string` | no |
| `vpcConfiguration` | `VpcConfiguration` | no |
| `initializationScript` | `string` | no |
| `commandLineArguments` | `List<KxCommandLineArgument>` | no |
| `code` | `CodeConfiguration` | no |
| `executionRole` | `string` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `savedownStorageConfiguration` | `KxSavedownStorageConfiguration` | no |
| `azMode` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `scalingGroupConfiguration` | `KxScalingGroupConfiguration` | no |

## CreateKxDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `databaseName` | `string` | no |
| `databaseArn` | `string` | no |
| `environmentId` | `string` | no |
| `description` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `lastModifiedTimestamp` | `timestamp` | no |

## CreateKxDataview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `dataviewName` | `string` | yes |
| `azMode` | `string` | yes |
| `availabilityZoneId` | `string` | no |
| `changesetId` | `string` | no |
| `segmentConfigurations` | `List<KxDataviewSegmentConfiguration>` | no |
| `autoUpdate` | `boolean` | no |
| `readWrite` | `boolean` | no |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataviewName` | `string` | no |
| `databaseName` | `string` | no |
| `environmentId` | `string` | no |
| `azMode` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `changesetId` | `string` | no |
| `segmentConfigurations` | `List<KxDataviewSegmentConfiguration>` | no |
| `description` | `string` | no |
| `autoUpdate` | `boolean` | no |
| `readWrite` | `boolean` | no |
| `createdTimestamp` | `timestamp` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `status` | `string` | no |

## CreateKxEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `kmsKeyId` | `string` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `status` | `string` | no |
| `environmentId` | `string` | no |
| `description` | `string` | no |
| `environmentArn` | `string` | no |
| `kmsKeyId` | `string` | no |
| `creationTimestamp` | `timestamp` | no |

## CreateKxScalingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `environmentId` | `string` | yes |
| `scalingGroupName` | `string` | yes |
| `hostType` | `string` | yes |
| `availabilityZoneId` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | no |
| `scalingGroupName` | `string` | no |
| `hostType` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `status` | `string` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `createdTimestamp` | `timestamp` | no |

## CreateKxUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `userName` | `string` | yes |
| `iamRole` | `string` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userName` | `string` | no |
| `userArn` | `string` | no |
| `environmentId` | `string` | no |
| `iamRole` | `string` | no |

## CreateKxVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `volumeType` | `string` | yes |
| `volumeName` | `string` | yes |
| `description` | `string` | no |
| `nas1Configuration` | `KxNAS1Configuration` | no |
| `azMode` | `string` | yes |
| `availabilityZoneIds` | `List<string>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | no |
| `volumeName` | `string` | no |
| `volumeType` | `string` | no |
| `volumeArn` | `string` | no |
| `nas1Configuration` | `KxNAS1Configuration` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `azMode` | `string` | no |
| `description` | `string` | no |
| `availabilityZoneIds` | `List<string>` | no |
| `createdTimestamp` | `timestamp` | no |

## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKxCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `clusterName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKxClusterNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `clusterName` | `string` | yes |
| `nodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKxDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKxDataview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `dataviewName` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKxEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKxScalingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `scalingGroupName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKxUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userName` | `string` | yes |
| `environmentId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKxVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `volumeName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | no |

## GetKxChangeset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `changesetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `changesetId` | `string` | no |
| `databaseName` | `string` | no |
| `environmentId` | `string` | no |
| `changeRequests` | `List<ChangeRequest>` | no |
| `createdTimestamp` | `timestamp` | no |
| `activeFromTimestamp` | `timestamp` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `status` | `string` | no |
| `errorInfo` | `ErrorInfo` | no |

## GetKxCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `clusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `clusterName` | `string` | no |
| `clusterType` | `string` | no |
| `tickerplantLogConfiguration` | `TickerplantLogConfiguration` | no |
| `volumes` | `List<Volume>` | no |
| `databases` | `List<KxDatabaseConfiguration>` | no |
| `cacheStorageConfigurations` | `List<KxCacheStorageConfiguration>` | no |
| `autoScalingConfiguration` | `AutoScalingConfiguration` | no |
| `clusterDescription` | `string` | no |
| `capacityConfiguration` | `CapacityConfiguration` | no |
| `releaseLabel` | `string` | no |
| `vpcConfiguration` | `VpcConfiguration` | no |
| `initializationScript` | `string` | no |
| `commandLineArguments` | `List<KxCommandLineArgument>` | no |
| `code` | `CodeConfiguration` | no |
| `executionRole` | `string` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `savedownStorageConfiguration` | `KxSavedownStorageConfiguration` | no |
| `azMode` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `scalingGroupConfiguration` | `KxScalingGroupConfiguration` | no |

## GetKxConnectionString

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userArn` | `string` | yes |
| `environmentId` | `string` | yes |
| `clusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signedConnectionString` | `string` | no |

## GetKxDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `databaseName` | `string` | no |
| `databaseArn` | `string` | no |
| `environmentId` | `string` | no |
| `description` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `lastCompletedChangesetId` | `string` | no |
| `numBytes` | `long` | no |
| `numChangesets` | `integer` | no |
| `numFiles` | `integer` | no |

## GetKxDataview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `dataviewName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `databaseName` | `string` | no |
| `dataviewName` | `string` | no |
| `azMode` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `changesetId` | `string` | no |
| `segmentConfigurations` | `List<KxDataviewSegmentConfiguration>` | no |
| `activeVersions` | `List<KxDataviewActiveVersion>` | no |
| `description` | `string` | no |
| `autoUpdate` | `boolean` | no |
| `readWrite` | `boolean` | no |
| `environmentId` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## GetKxEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `environmentId` | `string` | no |
| `awsAccountId` | `string` | no |
| `status` | `string` | no |
| `tgwStatus` | `string` | no |
| `dnsStatus` | `string` | no |
| `errorMessage` | `string` | no |
| `description` | `string` | no |
| `environmentArn` | `string` | no |
| `kmsKeyId` | `string` | no |
| `dedicatedServiceAccountId` | `string` | no |
| `transitGatewayConfiguration` | `TransitGatewayConfiguration` | no |
| `customDNSConfiguration` | `List<CustomDNSServer>` | no |
| `creationTimestamp` | `timestamp` | no |
| `updateTimestamp` | `timestamp` | no |
| `availabilityZoneIds` | `List<string>` | no |
| `certificateAuthorityArn` | `string` | no |

## GetKxScalingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `scalingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scalingGroupName` | `string` | no |
| `scalingGroupArn` | `string` | no |
| `hostType` | `string` | no |
| `clusters` | `List<string>` | no |
| `availabilityZoneId` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `createdTimestamp` | `timestamp` | no |

## GetKxUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userName` | `string` | yes |
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userName` | `string` | no |
| `userArn` | `string` | no |
| `environmentId` | `string` | no |
| `iamRole` | `string` | no |

## GetKxVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `volumeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | no |
| `volumeName` | `string` | no |
| `volumeType` | `string` | no |
| `volumeArn` | `string` | no |
| `nas1Configuration` | `KxNAS1Configuration` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `description` | `string` | no |
| `azMode` | `string` | no |
| `availabilityZoneIds` | `List<string>` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `attachedClusters` | `List<KxAttachedCluster>` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environments` | `List<Environment>` | no |
| `nextToken` | `string` | no |

## ListKxChangesets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kxChangesets` | `List<KxChangesetListEntry>` | no |
| `nextToken` | `string` | no |

## ListKxClusterNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `clusterName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodes` | `List<KxNode>` | no |
| `nextToken` | `string` | no |

## ListKxClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `clusterType` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kxClusterSummaries` | `List<KxCluster>` | no |
| `nextToken` | `string` | no |

## ListKxDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kxDatabases` | `List<KxDatabaseListEntry>` | no |
| `nextToken` | `string` | no |

## ListKxDataviews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kxDataviews` | `List<KxDataviewListEntry>` | no |
| `nextToken` | `string` | no |

## ListKxEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environments` | `List<KxEnvironment>` | no |
| `nextToken` | `string` | no |

## ListKxScalingGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scalingGroups` | `List<KxScalingGroup>` | no |
| `nextToken` | `string` | no |

## ListKxUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `users` | `List<KxUser>` | no |
| `nextToken` | `string` | no |

## ListKxVolumes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `volumeType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kxVolumeSummaries` | `List<KxVolume>` | no |
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


## UpdateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `federationMode` | `string` | no |
| `federationParameters` | `FederationParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | no |

## UpdateKxClusterCodeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `clusterName` | `string` | yes |
| `clientToken` | `string` | no |
| `code` | `CodeConfiguration` | yes |
| `initializationScript` | `string` | no |
| `commandLineArguments` | `List<KxCommandLineArgument>` | no |
| `deploymentConfiguration` | `KxClusterCodeDeploymentConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateKxClusterDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `clusterName` | `string` | yes |
| `clientToken` | `string` | no |
| `databases` | `List<KxDatabaseConfiguration>` | yes |
| `deploymentConfiguration` | `KxDeploymentConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateKxDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `databaseName` | `string` | no |
| `environmentId` | `string` | no |
| `description` | `string` | no |
| `lastModifiedTimestamp` | `timestamp` | no |

## UpdateKxDataview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `databaseName` | `string` | yes |
| `dataviewName` | `string` | yes |
| `description` | `string` | no |
| `changesetId` | `string` | no |
| `segmentConfigurations` | `List<KxDataviewSegmentConfiguration>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | no |
| `databaseName` | `string` | no |
| `dataviewName` | `string` | no |
| `azMode` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `changesetId` | `string` | no |
| `segmentConfigurations` | `List<KxDataviewSegmentConfiguration>` | no |
| `activeVersions` | `List<KxDataviewActiveVersion>` | no |
| `status` | `string` | no |
| `autoUpdate` | `boolean` | no |
| `readWrite` | `boolean` | no |
| `description` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `lastModifiedTimestamp` | `timestamp` | no |

## UpdateKxEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `environmentId` | `string` | no |
| `awsAccountId` | `string` | no |
| `status` | `string` | no |
| `tgwStatus` | `string` | no |
| `dnsStatus` | `string` | no |
| `errorMessage` | `string` | no |
| `description` | `string` | no |
| `environmentArn` | `string` | no |
| `kmsKeyId` | `string` | no |
| `dedicatedServiceAccountId` | `string` | no |
| `transitGatewayConfiguration` | `TransitGatewayConfiguration` | no |
| `customDNSConfiguration` | `List<CustomDNSServer>` | no |
| `creationTimestamp` | `timestamp` | no |
| `updateTimestamp` | `timestamp` | no |
| `availabilityZoneIds` | `List<string>` | no |

## UpdateKxEnvironmentNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `transitGatewayConfiguration` | `TransitGatewayConfiguration` | no |
| `customDNSConfiguration` | `List<CustomDNSServer>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `environmentId` | `string` | no |
| `awsAccountId` | `string` | no |
| `status` | `string` | no |
| `tgwStatus` | `string` | no |
| `dnsStatus` | `string` | no |
| `errorMessage` | `string` | no |
| `description` | `string` | no |
| `environmentArn` | `string` | no |
| `kmsKeyId` | `string` | no |
| `dedicatedServiceAccountId` | `string` | no |
| `transitGatewayConfiguration` | `TransitGatewayConfiguration` | no |
| `customDNSConfiguration` | `List<CustomDNSServer>` | no |
| `creationTimestamp` | `timestamp` | no |
| `updateTimestamp` | `timestamp` | no |
| `availabilityZoneIds` | `List<string>` | no |

## UpdateKxUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `userName` | `string` | yes |
| `iamRole` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userName` | `string` | no |
| `userArn` | `string` | no |
| `environmentId` | `string` | no |
| `iamRole` | `string` | no |

## UpdateKxVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `volumeName` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `nas1Configuration` | `KxNAS1Configuration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | no |
| `volumeName` | `string` | no |
| `volumeType` | `string` | no |
| `volumeArn` | `string` | no |
| `nas1Configuration` | `KxNAS1Configuration` | no |
| `status` | `string` | no |
| `description` | `string` | no |
| `statusReason` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `azMode` | `string` | no |
| `availabilityZoneIds` | `List<string>` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `attachedClusters` | `List<KxAttachedCluster>` | no |

