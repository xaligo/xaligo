# Route53 Recovery Cluster

API version: 2019-12-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/route53-recovery-cluster/2019-12-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetRoutingControlState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControlArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControlArn` | `string` | yes |
| `RoutingControlState` | `string` | yes |
| `RoutingControlName` | `string` | no |

## ListRoutingControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlPanelArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControls` | `List<RoutingControl>` | yes |
| `NextToken` | `string` | no |

## UpdateRoutingControlState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingControlArn` | `string` | yes |
| `RoutingControlState` | `string` | yes |
| `SafetyRulesToOverride` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoutingControlStates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateRoutingControlStateEntries` | `List<UpdateRoutingControlStateEntry>` | yes |
| `SafetyRulesToOverride` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


