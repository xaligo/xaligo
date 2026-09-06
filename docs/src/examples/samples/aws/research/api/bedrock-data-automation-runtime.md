# Runtime for Amazon Bedrock Data Automation

API version: 2024-06-13. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bedrock-data-automation-runtime/2024-06-13/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetDataAutomationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `errorType` | `string` | no |
| `errorMessage` | `string` | no |
| `outputConfiguration` | `OutputConfiguration` | no |
| `jobSubmissionTime` | `timestamp` | no |
| `jobCompletionTime` | `timestamp` | no |
| `jobDurationInSeconds` | `integer` | no |

## InvokeDataAutomation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inputConfiguration` | `SyncInputConfiguration` | yes |
| `dataAutomationConfiguration` | `DataAutomationConfiguration` | no |
| `blueprints` | `List<Blueprint>` | no |
| `dataAutomationProfileArn` | `string` | yes |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `outputConfiguration` | `OutputConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `outputConfiguration` | `OutputConfiguration` | no |
| `semanticModality` | `string` | yes |
| `outputSegments` | `List<OutputSegment>` | no |

## InvokeDataAutomationAsync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `inputConfiguration` | `InputConfiguration` | yes |
| `outputConfiguration` | `OutputConfiguration` | yes |
| `dataAutomationConfiguration` | `DataAutomationConfiguration` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `notificationConfiguration` | `NotificationConfiguration` | no |
| `blueprints` | `List<Blueprint>` | no |
| `dataAutomationProfileArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationArn` | `string` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


