# CloudWatch Observability Admin Service

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/observabilityadmin/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCentralizationRuleForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | yes |
| `Rule` | `CentralizationRule` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |

## CreateS3TableIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Encryption` | `Encryption` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## CreateTelemetryPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Configuration` | `TelemetryPipelineConfiguration` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## CreateTelemetryRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | yes |
| `Rule` | `TelemetryRule` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |

## CreateTelemetryRuleForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | yes |
| `Rule` | `TelemetryRule` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |

## DeleteCentralizationRuleForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteS3TableIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTelemetryPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTelemetryRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTelemetryRuleForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetCentralizationRuleForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | no |
| `RuleArn` | `string` | no |
| `CreatorAccountId` | `string` | no |
| `CreatedTimeStamp` | `long` | no |
| `CreatedRegion` | `string` | no |
| `LastUpdateTimeStamp` | `long` | no |
| `RuleHealth` | `string` | no |
| `FailureReason` | `string` | no |
| `TagPropagationStatus` | `string` | no |
| `TagPropagationFailureReason` | `string` | no |
| `CentralizationRule` | `CentralizationRule` | no |

## GetS3TableIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `RoleArn` | `string` | no |
| `Status` | `string` | no |
| `Encryption` | `Encryption` | no |
| `DestinationTableBucketArn` | `string` | no |
| `CreatedTimeStamp` | `long` | no |

## GetTelemetryEnrichmentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `AwsResourceExplorerManagedViewArn` | `string` | no |

## GetTelemetryEvaluationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `FailureReason` | `string` | no |
| `HomeRegion` | `string` | no |
| `RegionStatuses` | `List<RegionStatus>` | no |

## GetTelemetryEvaluationStatusForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `FailureReason` | `string` | no |
| `HomeRegion` | `string` | no |
| `RegionStatuses` | `List<RegionStatus>` | no |

## GetTelemetryPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Pipeline` | `TelemetryPipeline` | no |

## GetTelemetryRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | no |
| `RuleArn` | `string` | no |
| `CreatedTimeStamp` | `long` | no |
| `LastUpdateTimeStamp` | `long` | no |
| `TelemetryRule` | `TelemetryRule` | no |
| `HomeRegion` | `string` | no |
| `IsReplicated` | `boolean` | no |
| `RegionStatuses` | `List<RegionStatus>` | no |

## GetTelemetryRuleForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | no |
| `RuleArn` | `string` | no |
| `CreatedTimeStamp` | `long` | no |
| `LastUpdateTimeStamp` | `long` | no |
| `TelemetryRule` | `TelemetryRule` | no |
| `HomeRegion` | `string` | no |
| `IsReplicated` | `boolean` | no |
| `RegionStatuses` | `List<RegionStatus>` | no |

## ListCentralizationRulesForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleNamePrefix` | `string` | no |
| `AllRegions` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CentralizationRuleSummaries` | `List<CentralizationRuleSummary>` | no |
| `NextToken` | `string` | no |

## ListResourceTelemetry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifierPrefix` | `string` | no |
| `ResourceTypes` | `List<string>` | no |
| `TelemetryConfigurationState` | `Map<string>` | no |
| `ResourceTags` | `Map<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TelemetryConfigurations` | `List<TelemetryConfiguration>` | no |
| `NextToken` | `string` | no |

## ListResourceTelemetryForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIdentifiers` | `List<string>` | no |
| `ResourceIdentifierPrefix` | `string` | no |
| `ResourceTypes` | `List<string>` | no |
| `TelemetryConfigurationState` | `Map<string>` | no |
| `ResourceTags` | `Map<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TelemetryConfigurations` | `List<TelemetryConfiguration>` | no |
| `NextToken` | `string` | no |

## ListS3TableIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationSummaries` | `List<IntegrationSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | yes |

## ListTelemetryPipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineSummaries` | `List<TelemetryPipelineSummary>` | no |
| `NextToken` | `string` | no |

## ListTelemetryRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleNamePrefix` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TelemetryRuleSummaries` | `List<TelemetryRuleSummary>` | no |
| `NextToken` | `string` | no |

## ListTelemetryRulesForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleNamePrefix` | `string` | no |
| `SourceAccountIds` | `List<string>` | no |
| `SourceOrganizationUnitIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TelemetryRuleSummaries` | `List<TelemetryRuleSummary>` | no |
| `NextToken` | `string` | no |

## StartTelemetryEnrichment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `AwsResourceExplorerManagedViewArn` | `string` | no |

## StartTelemetryEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Regions` | `List<string>` | no |
| `AllRegions` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartTelemetryEvaluationForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Regions` | `List<string>` | no |
| `AllRegions` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopTelemetryEnrichment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## StopTelemetryEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopTelemetryEvaluationForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestTelemetryPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<Record>` | yes |
| `Configuration` | `TelemetryPipelineConfiguration` | yes |
| `SignalType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<PipelineOutput>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCentralizationRuleForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |
| `Rule` | `CentralizationRule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |

## UpdateTelemetryPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineIdentifier` | `string` | yes |
| `Configuration` | `TelemetryPipelineConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTelemetryRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |
| `Rule` | `TelemetryRule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |

## UpdateTelemetryRuleForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleIdentifier` | `string` | yes |
| `Rule` | `TelemetryRule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |

## ValidateTelemetryPipelineConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configuration` | `TelemetryPipelineConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<ValidationError>` | no |

