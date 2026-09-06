# AWS Resilience Hub

API version: 2020-04-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/resiliencehub/2020-04-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptResourceGroupingRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `entries` | `List<AcceptGroupingRecommendationEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `failedEntries` | `List<FailedGroupingRecommendationEntry>` | yes |

## AddDraftAppVersionResourceMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `resourceMappings` | `List<ResourceMapping>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `resourceMappings` | `List<ResourceMapping>` | yes |

## BatchUpdateRecommendationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `requestEntries` | `List<UpdateRecommendationStatusRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `failedEntries` | `List<BatchUpdateRecommendationStatusFailedEntry>` | yes |
| `successfulEntries` | `List<BatchUpdateRecommendationStatusSuccessfulEntry>` | yes |

## CreateApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentSchedule` | `string` | no |
| `awsApplicationArn` | `string` | no |
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `eventSubscriptions` | `List<EventSubscription>` | no |
| `name` | `string` | yes |
| `permissionModel` | `PermissionModel` | no |
| `policyArn` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `app` | `App` | yes |

## CreateAppVersionAppComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `additionalInfo` | `Map<List<string>>` | no |
| `appArn` | `string` | yes |
| `clientToken` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | yes |
| `type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appComponent` | `AppComponent` | no |
| `appVersion` | `string` | yes |

## CreateAppVersionResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `additionalInfo` | `Map<List<string>>` | no |
| `appArn` | `string` | yes |
| `appComponents` | `List<string>` | yes |
| `awsAccountId` | `string` | no |
| `awsRegion` | `string` | no |
| `clientToken` | `string` | no |
| `logicalResourceId` | `LogicalResourceId` | yes |
| `physicalResourceId` | `string` | yes |
| `resourceName` | `string` | no |
| `resourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `physicalResource` | `PhysicalResource` | no |

## CreateRecommendationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `bucketName` | `string` | no |
| `clientToken` | `string` | no |
| `format` | `string` | no |
| `name` | `string` | yes |
| `recommendationIds` | `List<string>` | no |
| `recommendationTypes` | `List<string>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationTemplate` | `RecommendationTemplate` | no |

## CreateResiliencyPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `dataLocationConstraint` | `string` | no |
| `policy` | `Map<FailurePolicy>` | yes |
| `policyDescription` | `string` | no |
| `policyName` | `string` | yes |
| `tags` | `Map<string>` | no |
| `tier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResiliencyPolicy` | yes |

## DeleteApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `clientToken` | `string` | no |
| `forceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |

## DeleteAppAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `assessmentStatus` | `string` | yes |

## DeleteAppInputSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `clientToken` | `string` | no |
| `eksSourceClusterNamespace` | `EksSourceClusterNamespace` | no |
| `sourceArn` | `string` | no |
| `terraformSource` | `TerraformSource` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | no |
| `appInputSource` | `AppInputSource` | no |

## DeleteAppVersionAppComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `clientToken` | `string` | no |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appComponent` | `AppComponent` | no |
| `appVersion` | `string` | yes |

## DeleteAppVersionResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `awsAccountId` | `string` | no |
| `awsRegion` | `string` | no |
| `clientToken` | `string` | no |
| `logicalResourceId` | `LogicalResourceId` | no |
| `physicalResourceId` | `string` | no |
| `resourceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `physicalResource` | `PhysicalResource` | no |

## DeleteRecommendationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `recommendationTemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationTemplateArn` | `string` | yes |
| `status` | `string` | yes |

## DeleteResiliencyPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `policyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |

## DescribeApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `app` | `App` | yes |

## DescribeAppAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessment` | `AppAssessment` | yes |

## DescribeAppVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `additionalInfo` | `Map<List<string>>` | no |
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |

## DescribeAppVersionAppComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appComponent` | `AppComponent` | no |
| `appVersion` | `string` | yes |

## DescribeAppVersionResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `awsAccountId` | `string` | no |
| `awsRegion` | `string` | no |
| `logicalResourceId` | `LogicalResourceId` | no |
| `physicalResourceId` | `string` | no |
| `resourceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `physicalResource` | `PhysicalResource` | no |

## DescribeAppVersionResourcesResolutionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `resolutionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `errorMessage` | `string` | no |
| `resolutionId` | `string` | yes |
| `status` | `string` | yes |

## DescribeAppVersionTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appTemplateBody` | `string` | yes |
| `appVersion` | `string` | yes |

## DescribeDraftAppVersionResourcesImportStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `errorDetails` | `List<ErrorDetail>` | no |
| `errorMessage` | `string` | no |
| `status` | `string` | yes |
| `statusChangeTime` | `timestamp` | yes |

## DescribeMetricsExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricsExportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorMessage` | `string` | no |
| `exportLocation` | `S3Location` | no |
| `metricsExportId` | `string` | yes |
| `status` | `string` | yes |

## DescribeResiliencyPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResiliencyPolicy` | yes |

## DescribeResourceGroupingRecommendationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `groupingId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorMessage` | `string` | no |
| `groupingId` | `string` | yes |
| `status` | `string` | yes |

## ImportResourcesToDraftAppVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `eksSources` | `List<EksSource>` | no |
| `importStrategy` | `string` | no |
| `sourceArns` | `List<string>` | no |
| `terraformSources` | `List<TerraformSource>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `eksSources` | `List<EksSource>` | no |
| `sourceArns` | `List<string>` | no |
| `status` | `string` | yes |
| `terraformSources` | `List<TerraformSource>` | no |

## ListAlarmRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `alarmRecommendations` | `List<AlarmRecommendation>` | yes |
| `nextToken` | `string` | no |

## ListAppAssessmentComplianceDrifts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `complianceDrifts` | `List<ComplianceDrift>` | yes |
| `nextToken` | `string` | no |

## ListAppAssessmentResourceDrifts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resourceDrifts` | `List<ResourceDrift>` | yes |

## ListAppAssessments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | no |
| `assessmentName` | `string` | no |
| `assessmentStatus` | `List<string>` | no |
| `complianceStatus` | `string` | no |
| `invoker` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `reverseOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentSummaries` | `List<AppAssessmentSummary>` | yes |
| `nextToken` | `string` | no |

## ListAppComponentCompliances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentCompliances` | `List<AppComponentCompliance>` | yes |
| `nextToken` | `string` | no |

## ListAppComponentRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentRecommendations` | `List<ComponentRecommendation>` | yes |
| `nextToken` | `string` | no |

## ListAppInputSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appInputSources` | `List<AppInputSource>` | yes |
| `nextToken` | `string` | no |

## ListAppVersionAppComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appComponents` | `List<AppComponent>` | no |
| `appVersion` | `string` | yes |
| `nextToken` | `string` | no |

## ListAppVersionResourceMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resourceMappings` | `List<ResourceMapping>` | yes |

## ListAppVersionResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `resolutionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `physicalResources` | `List<PhysicalResource>` | yes |
| `resolutionId` | `string` | yes |

## ListAppVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `endTime` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `startTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appVersions` | `List<AppVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | no |
| `awsApplicationArn` | `string` | no |
| `fromLastAssessmentTime` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `reverseOrder` | `boolean` | no |
| `toLastAssessmentTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appSummaries` | `List<AppSummary>` | yes |
| `nextToken` | `string` | no |

## ListMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `conditions` | `List<Condition>` | no |
| `dataSource` | `string` | no |
| `fields` | `List<Field>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sorts` | `List<Sort>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `rows` | `List<List<string>>` | yes |

## ListRecommendationTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | no |
| `maxResults` | `integer` | no |
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `recommendationTemplateArn` | `string` | no |
| `reverseOrder` | `boolean` | no |
| `status` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `recommendationTemplates` | `List<RecommendationTemplate>` | no |

## ListResiliencyPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `policyName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resiliencyPolicies` | `List<ResiliencyPolicy>` | yes |

## ListResourceGroupingRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `groupingRecommendations` | `List<GroupingRecommendation>` | yes |
| `nextToken` | `string` | no |

## ListSopRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `sopRecommendations` | `List<SopRecommendation>` | yes |

## ListSuggestedResiliencyPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resiliencyPolicies` | `List<ResiliencyPolicy>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTestRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `testRecommendations` | `List<TestRecommendation>` | yes |

## ListUnsupportedAppVersionResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `resolutionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resolutionId` | `string` | yes |
| `unsupportedResources` | `List<UnsupportedResource>` | yes |

## PublishAppVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `versionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | no |
| `identifier` | `long` | no |
| `versionName` | `string` | no |

## PutDraftAppVersionTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appTemplateBody` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | no |
| `appVersion` | `string` | no |

## RejectResourceGroupingRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `entries` | `List<RejectGroupingRecommendationEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `failedEntries` | `List<FailedGroupingRecommendationEntry>` | yes |

## RemoveDraftAppVersionResourceMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appRegistryAppNames` | `List<string>` | no |
| `eksSourceNames` | `List<string>` | no |
| `logicalStackNames` | `List<string>` | no |
| `resourceGroupNames` | `List<string>` | no |
| `resourceNames` | `List<string>` | no |
| `terraformSourceNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | no |
| `appVersion` | `string` | no |

## ResolveAppVersionResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `resolutionId` | `string` | yes |
| `status` | `string` | yes |

## StartAppAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `assessmentName` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessment` | `AppAssessment` | yes |

## StartMetricsExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucketName` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricsExportId` | `string` | yes |
| `status` | `string` | yes |

## StartResourceGroupingRecommendationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `errorMessage` | `string` | no |
| `groupingId` | `string` | yes |
| `status` | `string` | yes |

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


## UpdateApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `assessmentSchedule` | `string` | no |
| `clearResiliencyPolicyArn` | `boolean` | no |
| `description` | `string` | no |
| `eventSubscriptions` | `List<EventSubscription>` | no |
| `permissionModel` | `PermissionModel` | no |
| `policyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `app` | `App` | yes |

## UpdateAppVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `additionalInfo` | `Map<List<string>>` | no |
| `appArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `additionalInfo` | `Map<List<string>>` | no |
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |

## UpdateAppVersionAppComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `additionalInfo` | `Map<List<string>>` | no |
| `appArn` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appComponent` | `AppComponent` | no |
| `appVersion` | `string` | yes |

## UpdateAppVersionResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `additionalInfo` | `Map<List<string>>` | no |
| `appArn` | `string` | yes |
| `appComponents` | `List<string>` | no |
| `awsAccountId` | `string` | no |
| `awsRegion` | `string` | no |
| `excluded` | `boolean` | no |
| `logicalResourceId` | `LogicalResourceId` | no |
| `physicalResourceId` | `string` | no |
| `resourceName` | `string` | no |
| `resourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appArn` | `string` | yes |
| `appVersion` | `string` | yes |
| `physicalResource` | `PhysicalResource` | no |

## UpdateResiliencyPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataLocationConstraint` | `string` | no |
| `policy` | `Map<FailurePolicy>` | no |
| `policyArn` | `string` | yes |
| `policyDescription` | `string` | no |
| `policyName` | `string` | no |
| `tier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResiliencyPolicy` | yes |

