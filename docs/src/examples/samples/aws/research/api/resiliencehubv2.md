# AWS Resilience Hub V2

API version: 2026-02-17. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/resiliencehubv2/2026-02-17/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAssertion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `text` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assertion` | `Assertion` | yes |

## CreateInputSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `resourceConfiguration` | `ResourceConfiguration` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `inputSourceId` | `string` | yes |

## CreatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `availabilitySlo` | `AvailabilitySlo` | no |
| `multiAz` | `MultiAzTargets` | no |
| `multiRegion` | `MultiRegionTargets` | no |
| `dataRecovery` | `DataRecoveryTargets` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `Policy` | yes |

## CreateReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `reportType` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportGenerationResult` | `ReportGenerationResult` | yes |

## CreateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `associatedSystems` | `List<AssociatedSystem>` | no |
| `policyArn` | `string` | no |
| `regions` | `List<string>` | yes |
| `permissionModel` | `PermissionModel` | yes |
| `dependencyDiscovery` | `string` | no |
| `reportConfiguration` | `ServiceReportConfiguration` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | yes |

## CreateServiceFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `serviceArn` | `string` | yes |
| `description` | `string` | no |
| `criticality` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceFunction` | `ServiceFunction` | yes |

## CreateServiceFunctionResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `serviceFunctionId` | `string` | yes |
| `resources` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | no |
| `serviceFunctionId` | `string` | no |
| `resources` | `List<string>` | no |

## CreateSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `sharingEnabled` | `boolean` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `system` | `System` | yes |

## CreateTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `testTemplateArn` | `string` | yes |
| `loggingConfiguration` | `LoggingConfiguration` | no |
| `stopConditions` | `List<StopCondition>` | no |
| `roleName` | `string` | no |
| `parameters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `test` | `Test` | yes |

## CreateUserJourney

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `policyArn` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userJourney` | `UserJourney` | yes |

## DeleteAssertion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `assertionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assertionId` | `string` | no |

## DeleteInputSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `inputSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `inputSourceId` | `string` | yes |

## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |

## DeleteService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |

## DeleteServiceFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `serviceFunctionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceFunctionId` | `string` | no |

## DeleteServiceFunctionResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `serviceFunctionId` | `string` | yes |
| `resources` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | no |
| `serviceFunctionId` | `string` | no |
| `resources` | `List<string>` | no |

## DeleteSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |

## DeleteTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | yes |
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | yes |

## DeleteTestSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | yes |
| `serviceArn` | `string` | yes |
| `testSources` | `List<TestSourceInput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserJourney

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |
| `userJourneyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userJourneyId` | `string` | yes |

## GetFailureModeFinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingId` | `string` | yes |
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `finding` | `Finding` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `Policy` | yes |

## GetService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | yes |

## GetSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `system` | `System` | yes |

## GetTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | yes |
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `test` | `Test` | yes |

## GetTestRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRunId` | `string` | yes |
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRun` | `TestRun` | yes |

## GetTestTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testTemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testTemplate` | `TestTemplate` | yes |

## GetUserJourney

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |
| `userJourneyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userJourney` | `UserJourney` | yes |

## ImportApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `v1AppArn` | `string` | yes |
| `policyArn` | `string` | no |
| `kmsKeyId` | `string` | no |
| `skipManuallyAddedResources` | `boolean` | no |
| `associatedSystems` | `List<AssociatedSystem>` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | yes |

## ImportPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `v1PolicyArn` | `string` | yes |
| `kmsKeyId` | `string` | no |
| `availabilitySlo` | `AvailabilitySlo` | no |
| `multiAzDisasterRecoveryApproach` | `string` | no |
| `multiRegionDisasterRecoveryApproach` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `Policy` | yes |

## ListAssertions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `source` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assertions` | `List<Assertion>` | yes |
| `nextToken` | `string` | no |

## ListDependencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | no |
| `queryRangeStartTime` | `timestamp` | no |
| `queryRangeEndTime` | `timestamp` | no |
| `queryRangeGranularity` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dependencySummaries` | `List<DependencySummary>` | yes |
| `nextToken` | `string` | no |

## ListFailureModeAssessments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `assessmentStatuses` | `List<string>` | no |
| `startedAfter` | `timestamp` | no |
| `endedBefore` | `timestamp` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentSummaries` | `List<AssessmentSummary>` | yes |
| `nextToken` | `string` | no |

## ListFailureModeFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `severity` | `string` | no |
| `failureCategory` | `string` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingsSummary` | `List<FindingSummary>` | yes |
| `nextToken` | `string` | no |

## ListInputSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `type` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inputSourceSummaries` | `List<InputSourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policySummaries` | `List<PolicySummary>` | yes |
| `nextToken` | `string` | no |

## ListReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | no |
| `reportType` | `string` | no |
| `testRunId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportGenerationResults` | `List<ReportGenerationResult>` | yes |
| `nextToken` | `string` | no |

## ListResolvedTestRunTargetResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRunId` | `string` | yes |
| `serviceArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resolvedTargetResources` | `List<ResolvedTargetResource>` | yes |
| `nextToken` | `string` | no |

## ListResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `serviceFunctionId` | `string` | no |
| `awsRegion` | `string` | no |
| `resourceTypes` | `List<string>` | no |
| `billable` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceFunctionId` | `string` | no |
| `serviceResources` | `List<ServiceResource>` | no |
| `nextToken` | `string` | no |

## ListServiceEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `eventTypes` | `List<string>` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<ServiceEvent>` | yes |
| `nextToken` | `string` | no |

## ListServiceFunctions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceFunctions` | `List<ServiceFunction>` | yes |
| `nextToken` | `string` | no |

## ListServiceTopologyEdges

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTopologyEdgeSummaries` | `List<ServiceTopologyEdgeSummary>` | no |
| `nextToken` | `string` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | no |
| `userJourneyId` | `string` | no |
| `ouId` | `string` | no |
| `accountId` | `string` | no |
| `assessmentStatus` | `string` | no |
| `policyArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceSummaries` | `List<ServiceSummary>` | yes |
| `nextToken` | `string` | no |

## ListSystemEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |
| `eventTypes` | `List<string>` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<SystemEvent>` | yes |
| `nextToken` | `string` | no |

## ListSystems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ouId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemSummaries` | `List<SystemSummary>` | yes |
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

## ListTestRunEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRunId` | `string` | yes |
| `serviceArn` | `string` | yes |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<TestRunEvent>` | yes |
| `nextToken` | `string` | no |

## ListTestRunSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRunId` | `string` | yes |
| `serviceArn` | `string` | yes |
| `type` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRunSources` | `List<TestRunSourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListTestRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `testId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRuns` | `List<TestRunSummary>` | yes |
| `nextToken` | `string` | no |

## ListTestSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | yes |
| `serviceArn` | `string` | yes |
| `type` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSources` | `List<TestSourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListTestTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testTemplates` | `List<TestTemplateSummary>` | yes |

## ListTests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tests` | `List<TestSummary>` | yes |
| `nextToken` | `string` | no |

## ListUserJourneys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userJourneySummaries` | `List<UserJourneySummary>` | yes |
| `nextToken` | `string` | no |

## PutTestSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | yes |
| `serviceArn` | `string` | yes |
| `testSources` | `List<TestSourceInput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartFailureModeAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | no |
| `serviceArn` | `string` | no |
| `assessmentStatus` | `string` | no |
| `startedAt` | `timestamp` | no |

## StartTestRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | yes |
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRunId` | `string` | yes |
| `status` | `string` | yes |
| `experimentArns` | `List<string>` | yes |

## StopTestRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRunId` | `string` | yes |
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testRunId` | `string` | yes |
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


## UpdateAssertion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `assertionId` | `string` | yes |
| `text` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assertion` | `Assertion` | yes |

## UpdateDependency

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `dependencyId` | `string` | yes |
| `criticality` | `string` | no |
| `comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dependencyId` | `string` | yes |
| `dependencyName` | `string` | yes |
| `location` | `string` | yes |
| `criticality` | `string` | yes |
| `comment` | `string` | no |
| `provider` | `string` | no |
| `updatedAt` | `timestamp` | yes |

## UpdateFailureModeFinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingId` | `string` | yes |
| `status` | `string` | yes |
| `serviceArn` | `string` | yes |
| `comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `finding` | `Finding` | no |

## UpdatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `description` | `string` | no |
| `availabilitySlo` | `AvailabilitySlo` | no |
| `multiAz` | `MultiAzTargets` | no |
| `multiRegion` | `MultiRegionTargets` | no |
| `dataRecovery` | `DataRecoveryTargets` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `Policy` | yes |

## UpdateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `description` | `string` | no |
| `associatedSystems` | `List<AssociatedSystem>` | no |
| `policyArn` | `string` | no |
| `regions` | `List<string>` | no |
| `permissionModel` | `PermissionModel` | no |
| `dependencyDiscovery` | `string` | no |
| `reportConfiguration` | `ServiceReportConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | yes |

## UpdateServiceFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `serviceFunctionId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `criticality` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceFunction` | `ServiceFunction` | yes |

## UpdateSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |
| `description` | `string` | no |
| `sharingEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `system` | `System` | yes |

## UpdateTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | yes |
| `serviceArn` | `string` | yes |
| `loggingConfiguration` | `LoggingConfiguration` | no |
| `stopConditions` | `List<StopCondition>` | no |
| `roleName` | `string` | no |
| `parameters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `test` | `Test` | yes |

## UpdateUserJourney

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemArn` | `string` | yes |
| `userJourneyId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `policyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userJourney` | `UserJourney` | yes |

