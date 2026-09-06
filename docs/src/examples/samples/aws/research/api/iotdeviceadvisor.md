# AWS IoT Core Device Advisor

API version: 2020-09-18. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iotdeviceadvisor/2020-09-18/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateSuiteDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionConfiguration` | `SuiteDefinitionConfiguration` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | no |
| `suiteDefinitionArn` | `string` | no |
| `suiteDefinitionName` | `string` | no |
| `createdAt` | `timestamp` | no |

## DeleteSuiteDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingArn` | `string` | no |
| `certificateArn` | `string` | no |
| `deviceRoleArn` | `string` | no |
| `authenticationMethod` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `string` | no |

## GetSuiteDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | yes |
| `suiteDefinitionVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | no |
| `suiteDefinitionArn` | `string` | no |
| `suiteDefinitionVersion` | `string` | no |
| `latestVersion` | `string` | no |
| `suiteDefinitionConfiguration` | `SuiteDefinitionConfiguration` | no |
| `createdAt` | `timestamp` | no |
| `lastModifiedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## GetSuiteRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | yes |
| `suiteRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | no |
| `suiteDefinitionVersion` | `string` | no |
| `suiteRunId` | `string` | no |
| `suiteRunArn` | `string` | no |
| `suiteRunConfiguration` | `SuiteRunConfiguration` | no |
| `testResult` | `TestResult` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `status` | `string` | no |
| `errorReason` | `string` | no |
| `tags` | `Map<string>` | no |

## GetSuiteRunReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | yes |
| `suiteRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `qualificationReportDownloadUrl` | `string` | no |

## ListSuiteDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionInformationList` | `List<SuiteDefinitionInformation>` | no |
| `nextToken` | `string` | no |

## ListSuiteRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | no |
| `suiteDefinitionVersion` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteRunsList` | `List<SuiteRunInformation>` | no |
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

## StartSuiteRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | yes |
| `suiteDefinitionVersion` | `string` | no |
| `suiteRunConfiguration` | `SuiteRunConfiguration` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteRunId` | `string` | no |
| `suiteRunArn` | `string` | no |
| `createdAt` | `timestamp` | no |
| `endpoint` | `string` | no |

## StopSuiteRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | yes |
| `suiteRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateSuiteDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | yes |
| `suiteDefinitionConfiguration` | `SuiteDefinitionConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suiteDefinitionId` | `string` | no |
| `suiteDefinitionArn` | `string` | no |
| `suiteDefinitionName` | `string` | no |
| `suiteDefinitionVersion` | `string` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |

