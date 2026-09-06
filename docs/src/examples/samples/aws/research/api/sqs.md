# Amazon Simple Queue Service

API version: 2012-11-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sqs/2012-11-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `Label` | `string` | yes |
| `AWSAccountIds` | `List<string>` | yes |
| `Actions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelMessageMoveTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskHandle` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApproximateNumberOfMessagesMoved` | `long` | no |

## ChangeMessageVisibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `ReceiptHandle` | `string` | yes |
| `VisibilityTimeout` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ChangeMessageVisibilityBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `Entries` | `List<ChangeMessageVisibilityBatchRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<ChangeMessageVisibilityBatchResultEntry>` | yes |
| `Failed` | `List<BatchResultErrorEntry>` | yes |

## CreateQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueName` | `string` | yes |
| `Attributes` | `Map<string>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | no |

## DeleteMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `ReceiptHandle` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMessageBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `Entries` | `List<DeleteMessageBatchRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<DeleteMessageBatchResultEntry>` | yes |
| `Failed` | `List<BatchResultErrorEntry>` | yes |

## DeleteQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetQueueAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `AttributeNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<string>` | no |

## GetQueueUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueName` | `string` | yes |
| `QueueOwnerAWSAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | no |

## ListDeadLetterSourceQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queueUrls` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListMessageMoveTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<ListMessageMoveTasksResultEntry>` | no |

## ListQueueTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueNamePrefix` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrls` | `List<string>` | no |
| `NextToken` | `string` | no |

## PurgeQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReceiveMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `AttributeNames` | `List<string>` | no |
| `MessageSystemAttributeNames` | `List<string>` | no |
| `MessageAttributeNames` | `List<string>` | no |
| `MaxNumberOfMessages` | `integer` | no |
| `VisibilityTimeout` | `integer` | no |
| `WaitTimeSeconds` | `integer` | no |
| `ReceiveRequestAttemptId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Messages` | `List<Message>` | no |

## RemovePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `Label` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `MessageBody` | `string` | yes |
| `DelaySeconds` | `integer` | no |
| `MessageAttributes` | `Map<MessageAttributeValue>` | no |
| `MessageSystemAttributes` | `Map<MessageSystemAttributeValue>` | no |
| `MessageDeduplicationId` | `string` | no |
| `MessageGroupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MD5OfMessageBody` | `string` | no |
| `MD5OfMessageAttributes` | `string` | no |
| `MD5OfMessageSystemAttributes` | `string` | no |
| `MessageId` | `string` | no |
| `SequenceNumber` | `string` | no |

## SendMessageBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `Entries` | `List<SendMessageBatchRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<SendMessageBatchResultEntry>` | yes |
| `Failed` | `List<BatchResultErrorEntry>` | yes |

## SetQueueAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `Attributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartMessageMoveTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `DestinationArn` | `string` | no |
| `MaxNumberOfMessagesPerSecond` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskHandle` | `string` | no |

## TagQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueUrl` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


