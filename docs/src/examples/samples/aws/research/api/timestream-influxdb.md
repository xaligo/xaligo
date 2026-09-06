# Timestream InfluxDB

API version: 2023-01-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/timestream-influxdb/2023-01-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateDbBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `dbResourceId` | `string` | yes |
| `retentionDays` | `integer` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAfter` | `string` | no |
| `dbResourceId` | `string` | no |
| `type` | `string` | no |
| `engineType` | `string` | no |
| `deploymentType` | `string` | no |
| `kmsKeyId` | `string` | no |
| `clusterConfiguration` | `ClusterConfiguration` | no |
| `dbParameterGroupId` | `string` | no |
| `dbInstanceType` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `failoverMode` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `vpcSubnetIds` | `List<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `publiclyAccessible` | `boolean` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |

## CreateDbCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `username` | `string` | no |
| `password` | `string` | no |
| `organization` | `string` | no |
| `bucket` | `string` | no |
| `port` | `integer` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `dbInstanceType` | `string` | yes |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `networkType` | `string` | no |
| `publiclyAccessible` | `boolean` | no |
| `vpcSubnetIds` | `List<string>` | yes |
| `vpcSecurityGroupIds` | `List<string>` | yes |
| `deploymentType` | `string` | no |
| `failoverMode` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `dbBackupConfigurations` | `List<DbBackupConfiguration>` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterId` | `string` | no |
| `dbClusterStatus` | `string` | no |

## CreateDbInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `username` | `string` | no |
| `password` | `string` | yes |
| `organization` | `string` | no |
| `bucket` | `string` | no |
| `dbInstanceType` | `string` | yes |
| `vpcSubnetIds` | `List<string>` | yes |
| `vpcSecurityGroupIds` | `List<string>` | yes |
| `publiclyAccessible` | `boolean` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | yes |
| `dbParameterGroupIdentifier` | `string` | no |
| `deploymentType` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `tags` | `Map<string>` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `dbBackupConfigurations` | `List<DbBackupConfiguration>` | no |
| `kmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `endpoint` | `string` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `dbInstanceType` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `deploymentType` | `string` | no |
| `vpcSubnetIds` | `List<string>` | yes |
| `publiclyAccessible` | `boolean` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `availabilityZone` | `string` | no |
| `secondaryAvailabilityZone` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `dbClusterId` | `string` | no |
| `instanceMode` | `string` | no |
| `instanceModes` | `List<string>` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `lastMaintenanceTime` | `timestamp` | no |
| `nextMaintenanceTime` | `timestamp` | no |
| `dbBackupConfigurations` | `List<DbBackupConfigurationOutput>` | no |
| `kmsKeyId` | `string` | no |

## CreateDbParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `parameters` | `Parameters` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `parameters` | `Parameters` | no |

## DeleteDbBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAfter` | `string` | no |
| `dbResourceId` | `string` | no |
| `type` | `string` | no |
| `engineType` | `string` | no |
| `deploymentType` | `string` | no |
| `kmsKeyId` | `string` | no |
| `clusterConfiguration` | `ClusterConfiguration` | no |
| `dbParameterGroupId` | `string` | no |
| `dbInstanceType` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `failoverMode` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `vpcSubnetIds` | `List<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `publiclyAccessible` | `boolean` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |

## DeleteDbCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterId` | `string` | yes |
| `retainAutomatedBackups` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterStatus` | `string` | no |

## DeleteDbInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `retainAutomatedBackups` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `endpoint` | `string` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `dbInstanceType` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `deploymentType` | `string` | no |
| `vpcSubnetIds` | `List<string>` | yes |
| `publiclyAccessible` | `boolean` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `availabilityZone` | `string` | no |
| `secondaryAvailabilityZone` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `dbClusterId` | `string` | no |
| `instanceMode` | `string` | no |
| `instanceModes` | `List<string>` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `lastMaintenanceTime` | `timestamp` | no |
| `nextMaintenanceTime` | `timestamp` | no |
| `dbBackupConfigurations` | `List<DbBackupConfigurationOutput>` | no |
| `kmsKeyId` | `string` | no |

## GetDbBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAfter` | `string` | no |
| `dbResourceId` | `string` | no |
| `type` | `string` | no |
| `engineType` | `string` | no |
| `deploymentType` | `string` | no |
| `kmsKeyId` | `string` | no |
| `clusterConfiguration` | `ClusterConfiguration` | no |
| `dbParameterGroupId` | `string` | no |
| `dbInstanceType` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `failoverMode` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `vpcSubnetIds` | `List<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `publiclyAccessible` | `boolean` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |

## GetDbCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `endpoint` | `string` | no |
| `readerEndpoint` | `string` | no |
| `port` | `integer` | no |
| `deploymentType` | `string` | no |
| `dbInstanceType` | `string` | no |
| `networkType` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `engineType` | `string` | no |
| `publiclyAccessible` | `boolean` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `effectiveDbParameterGroupIdentifier` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `lastMaintenanceTime` | `timestamp` | no |
| `nextMaintenanceTime` | `timestamp` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `vpcSubnetIds` | `List<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `failoverMode` | `string` | no |
| `clusterConfiguration` | `ClusterConfiguration` | no |
| `dbBackupConfigurations` | `List<DbBackupConfigurationOutput>` | no |
| `kmsKeyId` | `string` | no |

## GetDbInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `endpoint` | `string` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `dbInstanceType` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `deploymentType` | `string` | no |
| `vpcSubnetIds` | `List<string>` | yes |
| `publiclyAccessible` | `boolean` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `availabilityZone` | `string` | no |
| `secondaryAvailabilityZone` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `dbClusterId` | `string` | no |
| `instanceMode` | `string` | no |
| `instanceModes` | `List<string>` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `lastMaintenanceTime` | `timestamp` | no |
| `nextMaintenanceTime` | `timestamp` | no |
| `dbBackupConfigurations` | `List<DbBackupConfigurationOutput>` | no |
| `kmsKeyId` | `string` | no |

## GetDbParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `parameters` | `Parameters` | no |

## ListDbBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbResourceId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DbBackupSummary>` | yes |
| `nextToken` | `string` | no |

## ListDbClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DbClusterSummary>` | yes |
| `nextToken` | `string` | no |

## ListDbInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DbInstanceSummary>` | yes |
| `nextToken` | `string` | no |

## ListDbInstancesForCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DbInstanceForClusterSummary>` | yes |
| `nextToken` | `string` | no |

## ListDbParameterGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DbParameterGroupSummary>` | yes |
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

## RebootDbCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterId` | `string` | yes |
| `instanceIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterStatus` | `string` | no |

## RebootDbInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `endpoint` | `string` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `dbInstanceType` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `deploymentType` | `string` | no |
| `vpcSubnetIds` | `List<string>` | yes |
| `publiclyAccessible` | `boolean` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `availabilityZone` | `string` | no |
| `secondaryAvailabilityZone` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `dbClusterId` | `string` | no |
| `instanceMode` | `string` | no |
| `instanceModes` | `List<string>` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `lastMaintenanceTime` | `timestamp` | no |
| `nextMaintenanceTime` | `timestamp` | no |
| `dbBackupConfigurations` | `List<DbBackupConfigurationOutput>` | no |
| `kmsKeyId` | `string` | no |

## RestoreFromDbBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `dbBackupId` | `string` | yes |
| `restoreToTime` | `timestamp` | no |
| `restoreMode` | `string` | no |
| `vpcSubnetIds` | `List<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `publiclyAccessible` | `boolean` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `tags` | `Map<string>` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `deploymentType` | `string` | no |
| `dbBackupConfigurations` | `List<DbBackupConfiguration>` | no |
| `kmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restoredDbResourceId` | `string` | no |
| `restoreStatus` | `string` | no |
| `resourceType` | `string` | no |
| `engineType` | `string` | no |
| `deploymentType` | `string` | no |

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


## UpdateDbCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterId` | `string` | yes |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `port` | `integer` | no |
| `dbInstanceType` | `string` | no |
| `failoverMode` | `string` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `dbBackupConfigurations` | `List<DbBackupConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterStatus` | `string` | no |

## UpdateDbInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `port` | `integer` | no |
| `dbInstanceType` | `string` | no |
| `deploymentType` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `dbBackupConfigurations` | `List<DbBackupConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `endpoint` | `string` | no |
| `port` | `integer` | no |
| `networkType` | `string` | no |
| `dbInstanceType` | `string` | no |
| `dbStorageType` | `string` | no |
| `allocatedStorage` | `integer` | no |
| `deploymentType` | `string` | no |
| `vpcSubnetIds` | `List<string>` | yes |
| `publiclyAccessible` | `boolean` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |
| `dbParameterGroupIdentifier` | `string` | no |
| `availabilityZone` | `string` | no |
| `secondaryAvailabilityZone` | `string` | no |
| `logDeliveryConfiguration` | `LogDeliveryConfiguration` | no |
| `influxAuthParametersSecretArn` | `string` | no |
| `dbClusterId` | `string` | no |
| `instanceMode` | `string` | no |
| `instanceModes` | `List<string>` | no |
| `maintenanceSchedule` | `MaintenanceSchedule` | no |
| `lastMaintenanceTime` | `timestamp` | no |
| `nextMaintenanceTime` | `timestamp` | no |
| `dbBackupConfigurations` | `List<DbBackupConfigurationOutput>` | no |
| `kmsKeyId` | `string` | no |

