# Agent Registry

API version: 2025-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/agent-registry/2025-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetDiscoverableRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entries` | `List<RegistryRecordsEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryRecords` | `List<RegistryRecordSummary>` | yes |
| `errors` | `List<BatchGetDiscoverableRegistryRecordError>` | yes |

## ListDiscoverableRegistryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `List<RegistryRecordFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryRecords` | `List<DiscoverableRegistryRecordSummary>` | yes |
| `nextToken` | `string` | no |

## SearchDiscoverableRegistryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchQuery` | `string` | yes |
| `registryIds` | `List<string>` | yes |
| `maxResults` | `integer` | no |
| `filters` | `MetadataFilterExpression` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryRecords` | `List<RegistryRecordSummary>` | yes |

