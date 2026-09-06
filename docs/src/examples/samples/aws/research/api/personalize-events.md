# Amazon Personalize Events

API version: 2018-03-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/personalize-events/2018-03-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## PutActionInteractions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trackingId` | `string` | yes |
| `actionInteractions` | `List<ActionInteraction>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `actions` | `List<Action>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trackingId` | `string` | yes |
| `userId` | `string` | no |
| `sessionId` | `string` | yes |
| `eventList` | `List<Event>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `items` | `List<Item>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `users` | `List<User>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


