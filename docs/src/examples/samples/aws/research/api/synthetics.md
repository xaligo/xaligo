# Synthetics

API version: 2017-10-11. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/synthetics/2017-10-11/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIdentifier` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateCanary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Code` | `CanaryCodeInput` | yes |
| `ArtifactS3Location` | `string` | yes |
| `ExecutionRoleArn` | `string` | yes |
| `Schedule` | `CanaryScheduleInput` | yes |
| `RunConfig` | `CanaryRunConfigInput` | no |
| `SuccessRetentionPeriodInDays` | `integer` | no |
| `FailureRetentionPeriodInDays` | `integer` | no |
| `RuntimeVersion` | `string` | yes |
| `VpcConfig` | `VpcConfigInput` | no |
| `ResourcesToReplicateTags` | `List<string>` | no |
| `ProvisionedResourceCleanup` | `string` | no |
| `BrowserConfigs` | `List<BrowserConfig>` | no |
| `AddReplicaLocations` | `List<AddReplicaLocationInput>` | no |
| `Tags` | `Map<string>` | no |
| `ArtifactConfig` | `ArtifactConfigInput` | no |
| `KmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Canary` | `Canary` | no |

## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |

## DeleteCanary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DeleteLambda` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeCanaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Names` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Canaries` | `List<Canary>` | no |
| `NextToken` | `string` | no |

## DescribeCanariesLastRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Names` | `List<string>` | no |
| `BrowserType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CanariesLastRun` | `List<CanaryLastRun>` | no |
| `NextToken` | `string` | no |

## DescribeRuntimeVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuntimeVersions` | `List<RuntimeVersion>` | no |
| `NextToken` | `string` | no |

## DisassociateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIdentifier` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetCanary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DryRunId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Canary` | `Canary` | no |

## GetCanaryRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRunId` | `string` | no |
| `RunType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CanaryRuns` | `List<CanaryRun>` | no |
| `NextToken` | `string` | no |

## GetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |

## ListAssociatedGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<GroupSummary>` | no |
| `NextToken` | `string` | no |

## ListGroupResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `GroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<GroupSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## StartCanary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartCanaryDryRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Code` | `CanaryCodeInput` | no |
| `RuntimeVersion` | `string` | no |
| `RunConfig` | `CanaryRunConfigInput` | no |
| `VpcConfig` | `VpcConfigInput` | no |
| `ExecutionRoleArn` | `string` | no |
| `SuccessRetentionPeriodInDays` | `integer` | no |
| `FailureRetentionPeriodInDays` | `integer` | no |
| `VisualReference` | `VisualReferenceInput` | no |
| `ArtifactS3Location` | `string` | no |
| `ArtifactConfig` | `ArtifactConfigInput` | no |
| `ProvisionedResourceCleanup` | `string` | no |
| `BrowserConfigs` | `List<BrowserConfig>` | no |
| `VisualReferences` | `List<VisualReferenceInput>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRunConfig` | `DryRunConfigOutput` | no |

## StopCanary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateCanary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Code` | `CanaryCodeInput` | no |
| `ExecutionRoleArn` | `string` | no |
| `RuntimeVersion` | `string` | no |
| `Schedule` | `CanaryScheduleInput` | no |
| `RunConfig` | `CanaryRunConfigInput` | no |
| `SuccessRetentionPeriodInDays` | `integer` | no |
| `FailureRetentionPeriodInDays` | `integer` | no |
| `VpcConfig` | `VpcConfigInput` | no |
| `VisualReference` | `VisualReferenceInput` | no |
| `ArtifactS3Location` | `string` | no |
| `ArtifactConfig` | `ArtifactConfigInput` | no |
| `ProvisionedResourceCleanup` | `string` | no |
| `DryRunId` | `string` | no |
| `VisualReferences` | `List<VisualReferenceInput>` | no |
| `BrowserConfigs` | `List<BrowserConfig>` | no |
| `AddReplicaLocations` | `List<AddReplicaLocationInput>` | no |
| `RemoveReplicaLocations` | `List<string>` | no |
| `KmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


