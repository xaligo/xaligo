# AWS Route53 Recovery Readiness

API version: 2019-12-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/route53-recovery-readiness/2019-12-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCell

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CellName` | `string` | yes |
| `Cells` | `List<string>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CellArn` | `string` | no |
| `CellName` | `string` | no |
| `Cells` | `List<string>` | no |
| `ParentReadinessScopes` | `List<string>` | no |
| `Tags` | `Map<string>` | no |

## CreateCrossAccountAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountAuthorization` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountAuthorization` | `string` | no |

## CreateReadinessCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReadinessCheckName` | `string` | yes |
| `ResourceSetName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReadinessCheckArn` | `string` | no |
| `ReadinessCheckName` | `string` | no |
| `ResourceSet` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateRecoveryGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cells` | `List<string>` | no |
| `RecoveryGroupName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cells` | `List<string>` | no |
| `RecoveryGroupArn` | `string` | no |
| `RecoveryGroupName` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateResourceSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetName` | `string` | yes |
| `ResourceSetType` | `string` | yes |
| `Resources` | `List<Resource>` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetArn` | `string` | no |
| `ResourceSetName` | `string` | no |
| `ResourceSetType` | `string` | no |
| `Resources` | `List<Resource>` | no |
| `Tags` | `Map<string>` | no |

## DeleteCell

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CellName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCrossAccountAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountAuthorization` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReadinessCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReadinessCheckName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRecoveryGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecoveryGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourceSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetArchitectureRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RecoveryGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastAuditTimestamp` | `timestamp` | no |
| `NextToken` | `string` | no |
| `Recommendations` | `List<Recommendation>` | no |

## GetCell

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CellName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CellArn` | `string` | no |
| `CellName` | `string` | no |
| `Cells` | `List<string>` | no |
| `ParentReadinessScopes` | `List<string>` | no |
| `Tags` | `Map<string>` | no |

## GetCellReadinessSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CellName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Readiness` | `string` | no |
| `ReadinessChecks` | `List<ReadinessCheckSummary>` | no |

## GetReadinessCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReadinessCheckName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReadinessCheckArn` | `string` | no |
| `ReadinessCheckName` | `string` | no |
| `ResourceSet` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetReadinessCheckResourceStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ReadinessCheckName` | `string` | yes |
| `ResourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Readiness` | `string` | no |
| `Rules` | `List<RuleResult>` | no |

## GetReadinessCheckStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ReadinessCheckName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Messages` | `List<Message>` | no |
| `NextToken` | `string` | no |
| `Readiness` | `string` | no |
| `Resources` | `List<ResourceResult>` | no |

## GetRecoveryGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecoveryGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cells` | `List<string>` | no |
| `RecoveryGroupArn` | `string` | no |
| `RecoveryGroupName` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetRecoveryGroupReadinessSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RecoveryGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Readiness` | `string` | no |
| `ReadinessChecks` | `List<ReadinessCheckSummary>` | no |

## GetResourceSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetArn` | `string` | no |
| `ResourceSetName` | `string` | no |
| `ResourceSetType` | `string` | no |
| `Resources` | `List<Resource>` | no |
| `Tags` | `Map<string>` | no |

## ListCells

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cells` | `List<CellOutput>` | no |
| `NextToken` | `string` | no |

## ListCrossAccountAuthorizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountAuthorizations` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListReadinessChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReadinessChecks` | `List<ReadinessCheckOutput>` | no |

## ListRecoveryGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RecoveryGroups` | `List<RecoveryGroupOutput>` | no |

## ListResourceSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ResourceSets` | `List<ResourceSetOutput>` | no |

## ListRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ResourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Rules` | `List<ListRulesOutput>` | no |

## ListTagsForResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

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
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCell

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CellName` | `string` | yes |
| `Cells` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CellArn` | `string` | no |
| `CellName` | `string` | no |
| `Cells` | `List<string>` | no |
| `ParentReadinessScopes` | `List<string>` | no |
| `Tags` | `Map<string>` | no |

## UpdateReadinessCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReadinessCheckName` | `string` | yes |
| `ResourceSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReadinessCheckArn` | `string` | no |
| `ReadinessCheckName` | `string` | no |
| `ResourceSet` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateRecoveryGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cells` | `List<string>` | yes |
| `RecoveryGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cells` | `List<string>` | no |
| `RecoveryGroupArn` | `string` | no |
| `RecoveryGroupName` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateResourceSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetName` | `string` | yes |
| `ResourceSetType` | `string` | yes |
| `Resources` | `List<Resource>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceSetArn` | `string` | no |
| `ResourceSetName` | `string` | no |
| `ResourceSetType` | `string` | no |
| `Resources` | `List<Resource>` | no |
| `Tags` | `Map<string>` | no |

