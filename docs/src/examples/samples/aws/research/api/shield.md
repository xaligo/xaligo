# AWS Shield

API version: 2016-06-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/shield/2016-06-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateDRTLogBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LogBucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateDRTRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateHealthCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionId` | `string` | yes |
| `HealthCheckArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateProactiveEngagementDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmergencyContactList` | `List<EmergencyContact>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionId` | `string` | no |

## CreateProtectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionGroupId` | `string` | yes |
| `Aggregation` | `string` | yes |
| `Pattern` | `string` | yes |
| `ResourceType` | `string` | no |
| `Members` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProtectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAttack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttackId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attack` | `AttackDetail` | no |

## DescribeAttackStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimeRange` | `TimeRange` | yes |
| `DataItems` | `List<AttackStatisticsDataItem>` | yes |

## DescribeDRTAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleArn` | `string` | no |
| `LogBucketList` | `List<string>` | no |

## DescribeEmergencyContactSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmergencyContactList` | `List<EmergencyContact>` | no |

## DescribeProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionId` | `string` | no |
| `ResourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Protection` | `Protection` | no |

## DescribeProtectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionGroup` | `ProtectionGroup` | yes |

## DescribeSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subscription` | `Subscription` | no |

## DisableApplicationLayerAutomaticResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableProactiveEngagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateDRTLogBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LogBucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateDRTRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateHealthCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionId` | `string` | yes |
| `HealthCheckArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableApplicationLayerAutomaticResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Action` | `ResponseAction` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableProactiveEngagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetSubscriptionState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionState` | `string` | yes |

## ListAttacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArns` | `List<string>` | no |
| `StartTime` | `TimeRange` | no |
| `EndTime` | `TimeRange` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttackSummaries` | `List<AttackSummary>` | no |
| `NextToken` | `string` | no |

## ListProtectionGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `InclusionFilters` | `InclusionProtectionGroupFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionGroups` | `List<ProtectionGroup>` | yes |
| `NextToken` | `string` | no |

## ListProtections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `InclusionFilters` | `InclusionProtectionFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Protections` | `List<Protection>` | no |
| `NextToken` | `string` | no |

## ListResourcesInProtectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionGroupId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArns` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApplicationLayerAutomaticResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Action` | `ResponseAction` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEmergencyContactSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmergencyContactList` | `List<EmergencyContact>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateProtectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectionGroupId` | `string` | yes |
| `Aggregation` | `string` | yes |
| `Pattern` | `string` | yes |
| `ResourceType` | `string` | no |
| `Members` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoRenew` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


