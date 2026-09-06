# Amazon Voice ID

API version: 2021-09-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/voice-id/2021-09-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateFraudster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `FraudsterId` | `string` | yes |
| `WatchlistId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fraudster` | `Fraudster` | no |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `ServerSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `Domain` | no |

## CreateWatchlist

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `DomainId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Watchlist` | `Watchlist` | no |

## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFraudster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `FraudsterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSpeaker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `SpeakerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWatchlist

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `WatchlistId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `Domain` | no |

## DescribeFraudster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `FraudsterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fraudster` | `Fraudster` | no |

## DescribeFraudsterRegistrationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Job` | `FraudsterRegistrationJob` | no |

## DescribeSpeaker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `SpeakerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Speaker` | `Speaker` | no |

## DescribeSpeakerEnrollmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Job` | `SpeakerEnrollmentJob` | no |

## DescribeWatchlist

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `WatchlistId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Watchlist` | `Watchlist` | no |

## DisassociateFraudster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `FraudsterId` | `string` | yes |
| `WatchlistId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fraudster` | `Fraudster` | no |

## EvaluateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `SessionNameOrId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationResult` | `AuthenticationResult` | no |
| `DomainId` | `string` | no |
| `FraudDetectionResult` | `FraudDetectionResult` | no |
| `SessionId` | `string` | no |
| `SessionName` | `string` | no |
| `StreamingStatus` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainSummaries` | `List<DomainSummary>` | no |
| `NextToken` | `string` | no |

## ListFraudsterRegistrationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `JobStatus` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobSummaries` | `List<FraudsterRegistrationJobSummary>` | no |
| `NextToken` | `string` | no |

## ListFraudsters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `WatchlistId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FraudsterSummaries` | `List<FraudsterSummary>` | no |
| `NextToken` | `string` | no |

## ListSpeakerEnrollmentJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `JobStatus` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobSummaries` | `List<SpeakerEnrollmentJobSummary>` | no |
| `NextToken` | `string` | no |

## ListSpeakers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SpeakerSummaries` | `List<SpeakerSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListWatchlists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WatchlistSummaries` | `List<WatchlistSummary>` | no |

## OptOutSpeaker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `SpeakerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Speaker` | `Speaker` | no |

## StartFraudsterRegistrationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DataAccessRoleArn` | `string` | yes |
| `DomainId` | `string` | yes |
| `InputDataConfig` | `InputDataConfig` | yes |
| `JobName` | `string` | no |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `RegistrationConfig` | `RegistrationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Job` | `FraudsterRegistrationJob` | no |

## StartSpeakerEnrollmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DataAccessRoleArn` | `string` | yes |
| `DomainId` | `string` | yes |
| `EnrollmentConfig` | `EnrollmentConfig` | no |
| `InputDataConfig` | `InputDataConfig` | yes |
| `JobName` | `string` | no |
| `OutputDataConfig` | `OutputDataConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Job` | `SpeakerEnrollmentJob` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DomainId` | `string` | yes |
| `Name` | `string` | yes |
| `ServerSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `Domain` | no |

## UpdateWatchlist

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DomainId` | `string` | yes |
| `Name` | `string` | no |
| `WatchlistId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Watchlist` | `Watchlist` | no |

