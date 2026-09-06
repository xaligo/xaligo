# AWS AppConfig Data

API version: 2021-11-11. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/appconfigdata/2021-11-11/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetLatestConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextPollConfigurationToken` | `string` | no |
| `NextPollIntervalInSeconds` | `integer` | no |
| `ContentType` | `string` | no |
| `Configuration` | `blob` | no |
| `VersionLabel` | `string` | no |

## StartConfigurationSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `ConfigurationProfileIdentifier` | `string` | yes |
| `RequiredMinimumPollIntervalInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InitialConfigurationToken` | `string` | no |

