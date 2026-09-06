# AWS Control Catalog

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/controlcatalog/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Aliases` | `List<string>` | no |
| `Name` | `string` | yes |
| `Description` | `string` | yes |
| `Behavior` | `string` | yes |
| `Severity` | `string` | no |
| `RegionConfiguration` | `RegionConfiguration` | yes |
| `Implementation` | `ImplementationDetails` | no |
| `ParameterRequirementSummary` | `string` | no |
| `Parameters` | `List<ControlParameter>` | no |
| `CreateTime` | `timestamp` | no |
| `GovernedResources` | `List<string>` | no |
| `GovernedProviders` | `List<string>` | no |

## ListCommonControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `CommonControlFilter` | `CommonControlFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CommonControls` | `List<CommonControlSummary>` | yes |
| `NextToken` | `string` | no |

## ListControlMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `ControlMappingFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlMappings` | `List<ControlMapping>` | yes |
| `NextToken` | `string` | no |

## ListControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `ControlFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Controls` | `List<ControlSummary>` | yes |
| `NextToken` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domains` | `List<DomainSummary>` | yes |
| `NextToken` | `string` | no |

## ListObjectives

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ObjectiveFilter` | `ObjectiveFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Objectives` | `List<ObjectiveSummary>` | yes |
| `NextToken` | `string` | no |

