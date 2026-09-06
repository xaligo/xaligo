# Amazon Aurora DSQL

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/dsql/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deletionProtectionEnabled` | `boolean` | no |
| `kmsEncryptionKey` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |
| `multiRegionProperties` | `MultiRegionProperties` | no |
| `policy` | `string` | no |
| `bypassPolicyLockoutSafetyCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `multiRegionProperties` | `MultiRegionProperties` | no |
| `encryptionDetails` | `EncryptionDetails` | no |
| `deletionProtectionEnabled` | `boolean` | yes |
| `endpoint` | `string` | no |

## CreateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `targetDefinition` | `TargetDefinition` | yes |
| `ordering` | `string` | yes |
| `format` | `string` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `streamIdentifier` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `ordering` | `string` | yes |
| `format` | `string` | yes |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## DeleteClusterPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `expectedPolicyVersion` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyVersion` | `string` | yes |

## DeleteStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `streamIdentifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `streamIdentifier` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## GetCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `deletionProtectionEnabled` | `boolean` | yes |
| `multiRegionProperties` | `MultiRegionProperties` | no |
| `tags` | `Map<string>` | no |
| `encryptionDetails` | `EncryptionDetails` | no |
| `endpoint` | `string` | no |

## GetClusterPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | yes |
| `policyVersion` | `string` | yes |

## GetStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `streamIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `streamIdentifier` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `ordering` | `string` | yes |
| `format` | `string` | yes |
| `targetDefinition` | `TargetDefinition` | no |
| `statusReason` | `StatusReason` | no |
| `tags` | `Map<string>` | no |

## GetVpcEndpointServiceName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |
| `clusterVpcEndpoint` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `clusters` | `List<ClusterSummary>` | yes |

## ListStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `streams` | `List<StreamSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutClusterPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `policy` | `string` | yes |
| `bypassPolicyLockoutSafetyCheck` | `boolean` | no |
| `expectedPolicyVersion` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyVersion` | `string` | yes |

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


## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `deletionProtectionEnabled` | `boolean` | no |
| `kmsEncryptionKey` | `string` | no |
| `clientToken` | `string` | no |
| `multiRegionProperties` | `MultiRegionProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |

