# Amazon Prometheus Service

API version: 2020-08-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/amp/2020-08-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAlertManagerDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `data` | `blob` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `AlertManagerDefinitionStatus` | yes |

## CreateAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `alias` | `string` | yes |
| `evaluationIntervalInSeconds` | `integer` | no |
| `missingDataAction` | `AnomalyDetectorMissingDataAction` | no |
| `configuration` | `AnomalyDetectorConfiguration` | yes |
| `labels` | `Map<string>` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectorId` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `AnomalyDetectorStatus` | yes |
| `tags` | `Map<string>` | no |

## CreateLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `logGroupArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `LoggingConfigurationStatus` | yes |

## CreateQueryLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `destinations` | `List<LoggingDestination>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `QueryLoggingConfigurationStatus` | yes |

## CreateRuleGroupsNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `name` | `string` | yes |
| `data` | `blob` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `RuleGroupsNamespaceStatus` | yes |
| `tags` | `Map<string>` | no |

## CreateScraper

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alias` | `string` | no |
| `scrapeConfiguration` | `ScrapeConfiguration` | yes |
| `source` | `Source` | yes |
| `destination` | `Destination` | yes |
| `roleConfiguration` | `RoleConfiguration` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `exporters` | `List<ExporterConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `ScraperStatus` | yes |
| `tags` | `Map<string>` | no |

## CreateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alias` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `WorkspaceStatus` | yes |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

## DeleteAlertManagerDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `anomalyDetectorId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueryLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `clientToken` | `string` | no |
| `revisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRuleGroupsNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `name` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScraper

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |
| `status` | `ScraperStatus` | yes |

## DeleteScraperLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAlertManagerDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alertManagerDefinition` | `AlertManagerDefinitionDescription` | yes |

## DescribeAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `anomalyDetectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetector` | `AnomalyDetectorDescription` | yes |

## DescribeLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loggingConfiguration` | `LoggingConfigurationMetadata` | yes |

## DescribeQueryLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryLoggingConfiguration` | `QueryLoggingConfigurationMetadata` | yes |

## DescribeResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyDocument` | `string` | yes |
| `policyStatus` | `string` | yes |
| `revisionId` | `string` | yes |

## DescribeRuleGroupsNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleGroupsNamespace` | `RuleGroupsNamespaceDescription` | yes |

## DescribeScraper

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraper` | `ScraperDescription` | yes |

## DescribeScraperLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `ScraperLoggingConfigurationStatus` | yes |
| `scraperId` | `string` | yes |
| `loggingDestination` | `ScraperLoggingDestination` | yes |
| `scraperComponents` | `List<ScraperComponent>` | yes |
| `modifiedAt` | `timestamp` | yes |

## DescribeWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspace` | `WorkspaceDescription` | yes |

## DescribeWorkspaceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceConfiguration` | `WorkspaceConfigurationDescription` | yes |

## GetDefaultScraperConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `blob` | yes |

## ListAnomalyDetectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `alias` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectors` | `List<AnomalyDetectorSummary>` | yes |
| `nextToken` | `string` | no |

## ListRuleGroupsNamespaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleGroupsNamespaces` | `List<RuleGroupsNamespaceSummary>` | yes |
| `nextToken` | `string` | no |

## ListScrapers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `Map<List<string>>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scrapers` | `List<ScraperSummary>` | yes |
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

## ListWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `alias` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaces` | `List<WorkspaceSummary>` | yes |
| `nextToken` | `string` | no |

## PutAlertManagerDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `data` | `blob` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `AlertManagerDefinitionStatus` | yes |

## PutAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `anomalyDetectorId` | `string` | yes |
| `evaluationIntervalInSeconds` | `integer` | no |
| `missingDataAction` | `AnomalyDetectorMissingDataAction` | no |
| `configuration` | `AnomalyDetectorConfiguration` | yes |
| `labels` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectorId` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `AnomalyDetectorStatus` | yes |
| `tags` | `Map<string>` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `policyDocument` | `string` | yes |
| `clientToken` | `string` | no |
| `revisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStatus` | `string` | yes |
| `revisionId` | `string` | yes |

## PutRuleGroupsNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `name` | `string` | yes |
| `data` | `blob` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `RuleGroupsNamespaceStatus` | yes |
| `tags` | `Map<string>` | no |

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


## UpdateLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `logGroupArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `LoggingConfigurationStatus` | yes |

## UpdateQueryLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `destinations` | `List<LoggingDestination>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `QueryLoggingConfigurationStatus` | yes |

## UpdateScraper

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |
| `alias` | `string` | no |
| `scrapeConfiguration` | `ScrapeConfiguration` | no |
| `destination` | `Destination` | no |
| `roleConfiguration` | `RoleConfiguration` | no |
| `clientToken` | `string` | no |
| `exporters` | `List<ExporterConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `ScraperStatus` | yes |
| `tags` | `Map<string>` | no |

## UpdateScraperLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scraperId` | `string` | yes |
| `loggingDestination` | `ScraperLoggingDestination` | yes |
| `scraperComponents` | `List<ScraperComponent>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `ScraperLoggingConfigurationStatus` | yes |

## UpdateWorkspaceAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `alias` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkspaceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `clientToken` | `string` | no |
| `limitsPerLabelSet` | `List<LimitsPerLabelSet>` | no |
| `retentionPeriodInDays` | `integer` | no |
| `outOfOrderTimeWindowInSeconds` | `integer` | no |
| `ruleQueryOffsetInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `WorkspaceConfigurationStatus` | yes |

