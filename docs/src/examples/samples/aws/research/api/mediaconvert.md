# AWS Elemental MediaConvert

API version: 2017-08-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mediaconvert/2017-08-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccelerationSettings` | `AccelerationSettings` | no |
| `BillingTagsSource` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `HopDestinations` | `List<HopDestination>` | no |
| `JobEngineVersion` | `string` | no |
| `JobTemplate` | `string` | no |
| `Priority` | `integer` | no |
| `Queue` | `string` | no |
| `Role` | `string` | yes |
| `Settings` | `JobSettings` | yes |
| `SimulateReservedQueue` | `string` | no |
| `StatusUpdateInterval` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UserMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Job` | `Job` | no |

## CreateJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccelerationSettings` | `AccelerationSettings` | no |
| `Category` | `string` | no |
| `Description` | `string` | no |
| `HopDestinations` | `List<HopDestination>` | no |
| `Name` | `string` | yes |
| `Priority` | `integer` | no |
| `Queue` | `string` | no |
| `Settings` | `JobTemplateSettings` | yes |
| `StatusUpdateInterval` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobTemplate` | `JobTemplate` | no |

## CreatePreset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Category` | `string` | no |
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Settings` | `PresetSettings` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Preset` | `Preset` | no |

## CreateQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConcurrentJobs` | `integer` | no |
| `Description` | `string` | no |
| `MaximumConcurrentFeeds` | `integer` | no |
| `Name` | `string` | yes |
| `PricingPlan` | `string` | no |
| `ReservationPlanSettings` | `ReservationPlanSettings` | no |
| `Status` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Queue` | `Queue` | no |

## CreateResourceShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `SupportCaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePreset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `Mode` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<Endpoint>` | no |
| `NextToken` | `string` | no |

## DisassociateCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Job` | `Job` | no |

## GetJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobTemplate` | `JobTemplate` | no |

## GetJobsQueryResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<Job>` | no |
| `NextToken` | `string` | no |
| `Status` | `string` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |

## GetPreset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Preset` | `Preset` | no |

## GetQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Queue` | `Queue` | no |

## ListJobTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Category` | `string` | no |
| `ListBy` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Order` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobTemplates` | `List<JobTemplate>` | no |
| `NextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Order` | `string` | no |
| `Queue` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<Job>` | no |
| `NextToken` | `string` | no |

## ListPresets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Category` | `string` | no |
| `ListBy` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Order` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Presets` | `List<Preset>` | no |

## ListQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListBy` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Order` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Queues` | `List<Queue>` | no |
| `TotalConcurrentJobs` | `integer` | no |
| `UnallocatedConcurrentJobs` | `integer` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTags` | `ResourceTags` | no |

## ListVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<JobEngineVersion>` | no |

## Probe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputFiles` | `List<ProbeInputFile>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProbeResults` | `List<ProbeResult>` | no |

## PutPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `Policy` | no |

## SearchJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputFile` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Order` | `string` | no |
| `Queue` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<Job>` | no |
| `NextToken` | `string` | no |

## StartJobsQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterList` | `List<JobsQueryFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Order` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `TagKeys` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccelerationSettings` | `AccelerationSettings` | no |
| `Category` | `string` | no |
| `Description` | `string` | no |
| `HopDestinations` | `List<HopDestination>` | no |
| `Name` | `string` | yes |
| `Priority` | `integer` | no |
| `Queue` | `string` | no |
| `Settings` | `JobTemplateSettings` | no |
| `StatusUpdateInterval` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobTemplate` | `JobTemplate` | no |

## UpdatePreset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Category` | `string` | no |
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Settings` | `PresetSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Preset` | `Preset` | no |

## UpdateQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConcurrentJobs` | `integer` | no |
| `Description` | `string` | no |
| `MaximumConcurrentFeeds` | `integer` | no |
| `Name` | `string` | yes |
| `ReservationPlanSettings` | `ReservationPlanSettings` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Queue` | `Queue` | no |

