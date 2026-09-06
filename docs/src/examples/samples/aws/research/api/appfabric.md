# AppFabric

API version: 2023-05-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/appfabric/2023-05-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetUserAccessTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `taskIdList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessResultsList` | `List<UserAccessResultItem>` | no |

## ConnectAppAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `appAuthorizationIdentifier` | `string` | yes |
| `authRequest` | `AuthRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appAuthorizationSummary` | `AppAuthorizationSummary` | yes |

## CreateAppAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `app` | `string` | yes |
| `credential` | `Credential` | yes |
| `tenant` | `Tenant` | yes |
| `authType` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appAuthorization` | `AppAuthorization` | yes |

## CreateAppBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `customerManagedKeyIdentifier` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundle` | `AppBundle` | yes |

## CreateIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `app` | `string` | yes |
| `tenantId` | `string` | yes |
| `ingestionType` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestion` | `Ingestion` | yes |

## CreateIngestionDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `ingestionIdentifier` | `string` | yes |
| `processingConfiguration` | `ProcessingConfiguration` | yes |
| `destinationConfiguration` | `DestinationConfiguration` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionDestination` | `IngestionDestination` | yes |

## DeleteAppAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `appAuthorizationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAppBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `ingestionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIngestionDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `ingestionIdentifier` | `string` | yes |
| `ingestionDestinationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAppAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `appAuthorizationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appAuthorization` | `AppAuthorization` | yes |

## GetAppBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundle` | `AppBundle` | yes |

## GetIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `ingestionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestion` | `Ingestion` | yes |

## GetIngestionDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `ingestionIdentifier` | `string` | yes |
| `ingestionDestinationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionDestination` | `IngestionDestination` | yes |

## ListAppAuthorizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appAuthorizationSummaryList` | `List<AppAuthorizationSummary>` | yes |
| `nextToken` | `string` | no |

## ListAppBundles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleSummaryList` | `List<AppBundleSummary>` | yes |
| `nextToken` | `string` | no |

## ListIngestionDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `ingestionIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionDestinations` | `List<IngestionDestinationSummary>` | yes |
| `nextToken` | `string` | no |

## ListIngestions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestions` | `List<IngestionSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## StartIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionIdentifier` | `string` | yes |
| `appBundleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartUserAccessTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `email` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessTasksList` | `List<UserAccessTaskItem>` | no |

## StopIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionIdentifier` | `string` | yes |
| `appBundleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateAppAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `appAuthorizationIdentifier` | `string` | yes |
| `credential` | `Credential` | no |
| `tenant` | `Tenant` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appAuthorization` | `AppAuthorization` | yes |

## UpdateIngestionDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appBundleIdentifier` | `string` | yes |
| `ingestionIdentifier` | `string` | yes |
| `ingestionDestinationIdentifier` | `string` | yes |
| `destinationConfiguration` | `DestinationConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionDestination` | `IngestionDestination` | yes |

