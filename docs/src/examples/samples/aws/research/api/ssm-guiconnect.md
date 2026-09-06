# AWS SSM-GUIConnect

API version: 2021-05-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ssm-guiconnect/2021-05-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteConnectionRecordingPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |

## GetConnectionRecordingPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ConnectionRecordingPreferences` | `ConnectionRecordingPreferences` | no |

## UpdateConnectionRecordingPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ConnectionRecordingPreferences` | `ConnectionRecordingPreferences` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ConnectionRecordingPreferences` | `ConnectionRecordingPreferences` | no |

