# Amazon Inspector

API version: 2016-02-16. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/inspector/2016-02-16/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddAttributesToFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingArns` | `List<string>` | yes |
| `attributes` | `List<Attribute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failedItems` | `Map<FailedItemDetails>` | yes |

## CreateAssessmentTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargetName` | `string` | yes |
| `resourceGroupArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargetArn` | `string` | yes |

## CreateAssessmentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargetArn` | `string` | yes |
| `assessmentTemplateName` | `string` | yes |
| `durationInSeconds` | `integer` | yes |
| `rulesPackageArns` | `List<string>` | yes |
| `userAttributesForFindings` | `List<Attribute>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplateArn` | `string` | yes |

## CreateExclusionsPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `previewToken` | `string` | yes |

## CreateResourceGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceGroupTags` | `List<ResourceGroupTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceGroupArn` | `string` | yes |

## DeleteAssessmentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssessmentTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssessmentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAssessmentRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRuns` | `List<AssessmentRun>` | yes |
| `failedItems` | `Map<FailedItemDetails>` | yes |

## DescribeAssessmentTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargetArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargets` | `List<AssessmentTarget>` | yes |
| `failedItems` | `Map<FailedItemDetails>` | yes |

## DescribeAssessmentTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplateArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplates` | `List<AssessmentTemplate>` | yes |
| `failedItems` | `Map<FailedItemDetails>` | yes |

## DescribeCrossAccountAccessRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleArn` | `string` | yes |
| `valid` | `boolean` | yes |
| `registeredAt` | `timestamp` | yes |

## DescribeExclusions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exclusionArns` | `List<string>` | yes |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exclusions` | `Map<Exclusion>` | yes |
| `failedItems` | `Map<FailedItemDetails>` | yes |

## DescribeFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingArns` | `List<string>` | yes |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<Finding>` | yes |
| `failedItems` | `Map<FailedItemDetails>` | yes |

## DescribeResourceGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceGroupArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceGroups` | `List<ResourceGroup>` | yes |
| `failedItems` | `Map<FailedItemDetails>` | yes |

## DescribeRulesPackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rulesPackageArns` | `List<string>` | yes |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rulesPackages` | `List<RulesPackage>` | yes |
| `failedItems` | `Map<FailedItemDetails>` | yes |

## GetAssessmentReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArn` | `string` | yes |
| `reportFileFormat` | `string` | yes |
| `reportType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `url` | `string` | no |

## GetExclusionsPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplateArn` | `string` | yes |
| `previewToken` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `previewStatus` | `string` | yes |
| `exclusionPreviews` | `List<ExclusionPreview>` | no |
| `nextToken` | `string` | no |

## GetTelemetryMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `telemetryMetadata` | `List<TelemetryMetadata>` | yes |

## ListAssessmentRunAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArn` | `string` | yes |
| `filter` | `AgentFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunAgents` | `List<AssessmentRunAgent>` | yes |
| `nextToken` | `string` | no |

## ListAssessmentRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplateArns` | `List<string>` | no |
| `filter` | `AssessmentRunFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArns` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListAssessmentTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `AssessmentTargetFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargetArns` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListAssessmentTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargetArns` | `List<string>` | no |
| `filter` | `AssessmentTemplateFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplateArns` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListEventSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptions` | `List<Subscription>` | yes |
| `nextToken` | `string` | no |

## ListExclusions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exclusionArns` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArns` | `List<string>` | no |
| `filter` | `FindingFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingArns` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListRulesPackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rulesPackageArns` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | yes |

## PreviewAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `previewAgentsArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentPreviews` | `List<AgentPreview>` | yes |
| `nextToken` | `string` | no |

## RegisterCrossAccountAccessRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveAttributesFromFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingArns` | `List<string>` | yes |
| `attributeKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failedItems` | `Map<FailedItemDetails>` | yes |

## SetTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAssessmentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTemplateArn` | `string` | yes |
| `assessmentRunName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArn` | `string` | yes |

## StopAssessmentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentRunArn` | `string` | yes |
| `stopAction` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SubscribeToEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `event` | `string` | yes |
| `topicArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UnsubscribeFromEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `event` | `string` | yes |
| `topicArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAssessmentTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargetArn` | `string` | yes |
| `assessmentTargetName` | `string` | yes |
| `resourceGroupArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


