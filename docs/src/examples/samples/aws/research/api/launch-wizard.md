# AWS Launch Wizard

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/launch-wizard/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |
| `deploymentPatternName` | `string` | yes |
| `name` | `string` | yes |
| `specifications` | `Map<string>` | yes |
| `dryRun` | `boolean` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |

## DeleteDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `statusReason` | `string` | no |

## GetDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployment` | `DeploymentData` | no |

## GetDeploymentPatternVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |
| `deploymentPatternName` | `string` | yes |
| `deploymentPatternVersionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentPatternVersion` | `DeploymentPatternVersionDataSummary` | no |

## GetWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workload` | `WorkloadData` | no |

## GetWorkloadDeploymentPattern

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |
| `deploymentPatternName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadDeploymentPattern` | `WorkloadDeploymentPatternData` | no |

## ListDeploymentEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentEvents` | `List<DeploymentEventDataSummary>` | no |
| `nextToken` | `string` | no |

## ListDeploymentPatternVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |
| `deploymentPatternName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `List<DeploymentPatternVersionFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentPatternVersions` | `List<DeploymentPatternVersionDataSummary>` | no |
| `nextToken` | `string` | no |

## ListDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<DeploymentFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployments` | `List<DeploymentDataSummary>` | no |
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

## ListWorkloadDeploymentPatterns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadDeploymentPatterns` | `List<WorkloadDeploymentPatternDataSummary>` | no |
| `nextToken` | `string` | no |

## ListWorkloads

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloads` | `List<WorkloadDataSummary>` | no |
| `nextToken` | `string` | no |

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


## UpdateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `specifications` | `Map<string>` | yes |
| `workloadVersionName` | `string` | no |
| `deploymentPatternVersionName` | `string` | no |
| `dryRun` | `boolean` | no |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployment` | `DeploymentDataSummary` | no |

