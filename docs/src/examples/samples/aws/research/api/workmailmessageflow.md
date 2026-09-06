# Amazon WorkMail Message Flow

API version: 2019-05-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/workmailmessageflow/2019-05-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetRawMessageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageContent` | `blob` | yes |

## PutRawMessageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageId` | `string` | yes |
| `content` | `RawMessageContent` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


