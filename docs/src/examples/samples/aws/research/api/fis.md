# AWS Fault Injection Simulator

API version: 2020-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/fis/2020-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateExperimentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `description` | `string` | yes |
| `stopConditions` | `List<CreateExperimentTemplateStopConditionInput>` | yes |
| `targets` | `Map<CreateExperimentTemplateTargetInput>` | no |
| `actions` | `Map<CreateExperimentTemplateActionInput>` | yes |
| `roleArn` | `string` | yes |
| `tags` | `Map<string>` | no |
| `logConfiguration` | `CreateExperimentTemplateLogConfigurationInput` | no |
| `experimentOptions` | `CreateExperimentTemplateExperimentOptionsInput` | no |
| `experimentReportConfiguration` | `CreateExperimentTemplateReportConfigurationInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplate` | `ExperimentTemplate` | no |

## CreateTargetAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `experimentTemplateId` | `string` | yes |
| `accountId` | `string` | yes |
| `roleArn` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetAccountConfiguration` | `TargetAccountConfiguration` | no |

## DeleteExperimentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplate` | `ExperimentTemplate` | no |

## DeleteTargetAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplateId` | `string` | yes |
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetAccountConfiguration` | `TargetAccountConfiguration` | no |

## GetAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `Action` | no |

## GetExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experiment` | `Experiment` | no |

## GetExperimentTargetAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentId` | `string` | yes |
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetAccountConfiguration` | `ExperimentTargetAccountConfiguration` | no |

## GetExperimentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplate` | `ExperimentTemplate` | no |

## GetSafetyLever

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `safetyLever` | `SafetyLever` | no |

## GetTargetAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplateId` | `string` | yes |
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetAccountConfiguration` | `TargetAccountConfiguration` | no |

## GetTargetResourceType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetResourceType` | `TargetResourceType` | no |

## ListActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actions` | `List<ActionSummary>` | no |
| `nextToken` | `string` | no |

## ListExperimentResolvedTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `targetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resolvedTargets` | `List<ResolvedTarget>` | no |
| `nextToken` | `string` | no |

## ListExperimentTargetAccountConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentId` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetAccountConfigurations` | `List<ExperimentTargetAccountConfigurationSummary>` | no |
| `nextToken` | `string` | no |

## ListExperimentTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplates` | `List<ExperimentTemplateSummary>` | no |
| `nextToken` | `string` | no |

## ListExperiments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `experimentTemplateId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experiments` | `List<ExperimentSummary>` | no |
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

## ListTargetAccountConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplateId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetAccountConfigurations` | `List<TargetAccountConfigurationSummary>` | no |
| `nextToken` | `string` | no |

## ListTargetResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetResourceTypes` | `List<TargetResourceTypeSummary>` | no |
| `nextToken` | `string` | no |

## StartExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `experimentTemplateId` | `string` | yes |
| `experimentOptions` | `StartExperimentExperimentOptionsInput` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experiment` | `Experiment` | no |

## StopExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experiment` | `Experiment` | no |

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
| `tagKeys` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateExperimentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `description` | `string` | no |
| `stopConditions` | `List<UpdateExperimentTemplateStopConditionInput>` | no |
| `targets` | `Map<UpdateExperimentTemplateTargetInput>` | no |
| `actions` | `Map<UpdateExperimentTemplateActionInputItem>` | no |
| `roleArn` | `string` | no |
| `logConfiguration` | `UpdateExperimentTemplateLogConfigurationInput` | no |
| `experimentOptions` | `UpdateExperimentTemplateExperimentOptionsInput` | no |
| `experimentReportConfiguration` | `UpdateExperimentTemplateReportConfigurationInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplate` | `ExperimentTemplate` | no |

## UpdateSafetyLeverState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `state` | `UpdateSafetyLeverStateInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `safetyLever` | `SafetyLever` | no |

## UpdateTargetAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `experimentTemplateId` | `string` | yes |
| `accountId` | `string` | yes |
| `roleArn` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetAccountConfiguration` | `TargetAccountConfiguration` | no |

