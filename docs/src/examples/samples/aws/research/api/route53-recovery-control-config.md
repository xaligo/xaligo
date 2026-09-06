# AWS Route53 Recovery Control Config

API version: 2020-11-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/route53-recovery-control-config/2020-11-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ClusterName` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## CreateControlPanel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ClusterArn` | `string` | yes |
| `ControlPanelName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanel` | `ControlPanel` | no |

## CreateRoutingControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ClusterArn` | `string` | yes |
| `ControlPanelArn` | `string` | no |
| `RoutingControlName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControl` | `RoutingControl` | no |

## CreateSafetyRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssertionRule` | `NewAssertionRule` | no |
| `ClientToken` | `string` | no |
| `GatingRule` | `NewGatingRule` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssertionRule` | `AssertionRule` | no |
| `GatingRule` | `GatingRule` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteControlPanel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRoutingControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControlArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSafetyRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SafetyRuleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## DescribeControlPanel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanel` | `ControlPanel` | no |

## DescribeRoutingControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControlArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControl` | `RoutingControl` | no |

## DescribeSafetyRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SafetyRuleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssertionRule` | `AssertionRule` | no |
| `GatingRule` | `GatingRule` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## ListAssociatedRoute53HealthChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RoutingControlArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckIds` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Clusters` | `List<Cluster>` | no |
| `NextToken` | `string` | no |

## ListControlPanels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanels` | `List<ControlPanel>` | no |
| `NextToken` | `string` | no |

## ListRoutingControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanelArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RoutingControls` | `List<RoutingControl>` | no |

## ListSafetyRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanelArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SafetyRules` | `List<Rule>` | no |

## ListTagsForResource

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


## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `NetworkType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## UpdateControlPanel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanelArn` | `string` | yes |
| `ControlPanelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanel` | `ControlPanel` | no |

## UpdateRoutingControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControlArn` | `string` | yes |
| `RoutingControlName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControl` | `RoutingControl` | no |

## UpdateSafetyRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssertionRuleUpdate` | `AssertionRuleUpdate` | no |
| `GatingRuleUpdate` | `GatingRuleUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssertionRule` | `AssertionRule` | no |
| `GatingRule` | `GatingRule` | no |

