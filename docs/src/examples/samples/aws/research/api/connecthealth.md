# Connect Health

API version: 2025-01-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/connecthealth/2025-01-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `subscriptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `SubscriptionDescription` | no |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `webAppSetupConfiguration` | `CreateWebAppConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `encryptionContext` | `EncryptionContext` | no |
| `status` | `string` | yes |
| `webAppUrl` | `string` | no |
| `webAppConfiguration` | `WebAppConfiguration` | no |
| `createdAt` | `timestamp` | yes |

## CreateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `subscriptionId` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `activatedAt` | `timestamp` | no |
| `deactivatedAt` | `timestamp` | no |

## DeactivateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `subscriptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `SubscriptionDescription` | no |

## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |

## GetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `encryptionContext` | `EncryptionContext` | no |
| `status` | `string` | yes |
| `webAppUrl` | `string` | no |
| `webAppConfiguration` | `WebAppConfiguration` | no |
| `createdAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## GetMedicalScribeListeningSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `domainId` | `string` | yes |
| `subscriptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `medicalScribeListeningSessionDetails` | `MedicalScribeListeningSessionDetails` | no |

## GetPatientInsightsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `jobArn` | `string` | yes |
| `jobStatus` | `string` | yes |
| `creationTime` | `timestamp` | no |
| `updatedTime` | `timestamp` | no |
| `insightsOutput` | `InsightsOutput` | no |
| `statusDetails` | `string` | no |
| `patientContext` | `PatientInsightsPatientContext` | yes |
| `insightsContext` | `InsightsContext` | yes |
| `encounterContext` | `PatientInsightsEncounterContext` | yes |
| `userContext` | `UserContext` | yes |
| `inputDataConfig` | `InputDataConfig` | yes |
| `outputDataConfig` | `OutputDataConfig` | yes |

## GetSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `subscriptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `SubscriptionDescription` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domains` | `List<DomainSummary>` | yes |
| `nextToken` | `string` | no |

## ListSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptions` | `List<SubscriptionDescription>` | yes |
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

## StartMedicalScribeListeningSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `domainId` | `string` | yes |
| `subscriptionId` | `string` | yes |
| `languageCode` | `string` | yes |
| `mediaSampleRateHertz` | `integer` | yes |
| `mediaEncoding` | `string` | yes |
| `inputStream` | `MedicalScribeInputStream` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | no |
| `domainId` | `string` | no |
| `subscriptionId` | `string` | no |
| `requestId` | `string` | no |
| `languageCode` | `string` | no |
| `mediaSampleRateHertz` | `integer` | no |
| `mediaEncoding` | `string` | no |
| `responseStream` | `MedicalScribeOutputStream` | no |

## StartPatientInsightsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `patientContext` | `PatientInsightsPatientContext` | yes |
| `insightsContext` | `InsightsContext` | yes |
| `encounterContext` | `PatientInsightsEncounterContext` | yes |
| `userContext` | `UserContext` | yes |
| `inputDataConfig` | `InputDataConfig` | yes |
| `outputDataConfig` | `OutputDataConfig` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |
| `jobId` | `string` | yes |
| `creationTime` | `timestamp` | no |

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


