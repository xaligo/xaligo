# AWS Compute Optimizer

API version: 2019-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/compute-optimizer/2019-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteRecommendationPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |
| `scope` | `Scope` | no |
| `recommendationPreferenceNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeRecommendationExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIds` | `List<string>` | no |
| `filters` | `List<JobFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationExportJobs` | `List<RecommendationExportJob>` | no |
| `nextToken` | `string` | no |

## ExportAutoScalingGroupRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `filters` | `List<Filter>` | no |
| `fieldsToExport` | `List<string>` | no |
| `s3DestinationConfig` | `S3DestinationConfig` | yes |
| `fileFormat` | `string` | no |
| `includeMemberAccounts` | `boolean` | no |
| `recommendationPreferences` | `RecommendationPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `s3Destination` | `S3Destination` | no |

## ExportEBSVolumeRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `filters` | `List<EBSFilter>` | no |
| `fieldsToExport` | `List<string>` | no |
| `s3DestinationConfig` | `S3DestinationConfig` | yes |
| `fileFormat` | `string` | no |
| `includeMemberAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `s3Destination` | `S3Destination` | no |

## ExportEC2InstanceRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `filters` | `List<Filter>` | no |
| `fieldsToExport` | `List<string>` | no |
| `s3DestinationConfig` | `S3DestinationConfig` | yes |
| `fileFormat` | `string` | no |
| `includeMemberAccounts` | `boolean` | no |
| `recommendationPreferences` | `RecommendationPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `s3Destination` | `S3Destination` | no |

## ExportECSServiceRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `filters` | `List<ECSServiceRecommendationFilter>` | no |
| `fieldsToExport` | `List<string>` | no |
| `s3DestinationConfig` | `S3DestinationConfig` | yes |
| `fileFormat` | `string` | no |
| `includeMemberAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `s3Destination` | `S3Destination` | no |

## ExportIdleRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `filters` | `List<IdleRecommendationFilter>` | no |
| `fieldsToExport` | `List<string>` | no |
| `s3DestinationConfig` | `S3DestinationConfig` | yes |
| `fileFormat` | `string` | no |
| `includeMemberAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `s3Destination` | `S3Destination` | no |

## ExportLambdaFunctionRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `filters` | `List<LambdaFunctionRecommendationFilter>` | no |
| `fieldsToExport` | `List<string>` | no |
| `s3DestinationConfig` | `S3DestinationConfig` | yes |
| `fileFormat` | `string` | no |
| `includeMemberAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `s3Destination` | `S3Destination` | no |

## ExportLicenseRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `filters` | `List<LicenseRecommendationFilter>` | no |
| `fieldsToExport` | `List<string>` | no |
| `s3DestinationConfig` | `S3DestinationConfig` | yes |
| `fileFormat` | `string` | no |
| `includeMemberAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `s3Destination` | `S3Destination` | no |

## ExportRDSDatabaseRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `filters` | `List<RDSDBRecommendationFilter>` | no |
| `fieldsToExport` | `List<string>` | no |
| `s3DestinationConfig` | `S3DestinationConfig` | yes |
| `fileFormat` | `string` | no |
| `includeMemberAccounts` | `boolean` | no |
| `recommendationPreferences` | `RecommendationPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `s3Destination` | `S3Destination` | no |

## GetAutoScalingGroupRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `autoScalingGroupArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<Filter>` | no |
| `recommendationPreferences` | `RecommendationPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `autoScalingGroupRecommendations` | `List<AutoScalingGroupRecommendation>` | no |
| `errors` | `List<GetRecommendationError>` | no |

## GetEBSVolumeRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `volumeArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<EBSFilter>` | no |
| `accountIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `volumeRecommendations` | `List<VolumeRecommendation>` | no |
| `errors` | `List<GetRecommendationError>` | no |

## GetEC2InstanceRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<Filter>` | no |
| `accountIds` | `List<string>` | no |
| `recommendationPreferences` | `RecommendationPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `instanceRecommendations` | `List<InstanceRecommendation>` | no |
| `errors` | `List<GetRecommendationError>` | no |

## GetEC2RecommendationProjectedMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceArn` | `string` | yes |
| `stat` | `string` | yes |
| `period` | `integer` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `recommendationPreferences` | `RecommendationPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendedOptionProjectedMetrics` | `List<RecommendedOptionProjectedMetric>` | no |

## GetECSServiceRecommendationProjectedMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `stat` | `string` | yes |
| `period` | `integer` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendedOptionProjectedMetrics` | `List<ECSServiceRecommendedOptionProjectedMetric>` | no |

## GetECSServiceRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<ECSServiceRecommendationFilter>` | no |
| `accountIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `ecsServiceRecommendations` | `List<ECSServiceRecommendation>` | no |
| `errors` | `List<GetRecommendationError>` | no |

## GetEffectiveRecommendationPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enhancedInfrastructureMetrics` | `string` | no |
| `externalMetricsPreference` | `ExternalMetricsPreference` | no |
| `lookBackPeriod` | `string` | no |
| `utilizationPreferences` | `List<UtilizationPreference>` | no |
| `preferredResources` | `List<EffectivePreferredResource>` | no |

## GetEnrollmentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `memberAccountsEnrolled` | `boolean` | no |
| `lastUpdatedTimestamp` | `timestamp` | no |
| `numberOfMemberAccountsOptedIn` | `integer` | no |

## GetEnrollmentStatusesForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<EnrollmentFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountEnrollmentStatuses` | `List<AccountEnrollmentStatus>` | no |
| `nextToken` | `string` | no |

## GetIdleRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<IdleRecommendationFilter>` | no |
| `accountIds` | `List<string>` | no |
| `orderBy` | `OrderBy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `idleRecommendations` | `List<IdleRecommendation>` | no |
| `errors` | `List<IdleRecommendationError>` | no |

## GetLambdaFunctionRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `functionArns` | `List<string>` | no |
| `accountIds` | `List<string>` | no |
| `filters` | `List<LambdaFunctionRecommendationFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `lambdaFunctionRecommendations` | `List<LambdaFunctionRecommendation>` | no |

## GetLicenseRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<LicenseRecommendationFilter>` | no |
| `accountIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `licenseRecommendations` | `List<LicenseRecommendation>` | no |
| `errors` | `List<GetRecommendationError>` | no |

## GetRDSDatabaseRecommendationProjectedMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `stat` | `string` | yes |
| `period` | `integer` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `recommendationPreferences` | `RecommendationPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendedOptionProjectedMetrics` | `List<RDSDatabaseRecommendedOptionProjectedMetric>` | no |

## GetRDSDatabaseRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<RDSDBRecommendationFilter>` | no |
| `accountIds` | `List<string>` | no |
| `recommendationPreferences` | `RecommendationPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `rdsDBRecommendations` | `List<RDSDBRecommendation>` | no |
| `errors` | `List<GetRecommendationError>` | no |

## GetRecommendationPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |
| `scope` | `Scope` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `recommendationPreferencesDetails` | `List<RecommendationPreferencesDetail>` | no |

## GetRecommendationSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `recommendationSummaries` | `List<RecommendationSummary>` | no |

## PutRecommendationPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |
| `scope` | `Scope` | no |
| `enhancedInfrastructureMetrics` | `string` | no |
| `inferredWorkloadTypes` | `string` | no |
| `externalMetricsPreference` | `ExternalMetricsPreference` | no |
| `lookBackPeriod` | `string` | no |
| `utilizationPreferences` | `List<UtilizationPreference>` | no |
| `preferredResources` | `List<PreferredResource>` | no |
| `savingsEstimationMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEnrollmentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `includeMemberAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `statusReason` | `string` | no |

