# AWS Systems Manager Incident Manager

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ssm-incidents/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetIncidentFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingIds` | `List<string>` | yes |
| `incidentRecordArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchGetIncidentFindingsError>` | yes |
| `findings` | `List<Finding>` | yes |

## CreateReplicationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `regions` | `Map<RegionMapInputValue>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## CreateResponsePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actions` | `List<Action>` | no |
| `chatChannel` | `ChatChannel` | no |
| `clientToken` | `string` | no |
| `displayName` | `string` | no |
| `engagements` | `List<string>` | no |
| `incidentTemplate` | `IncidentTemplate` | yes |
| `integrations` | `List<Integration>` | no |
| `name` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## CreateTimelineEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `eventData` | `string` | yes |
| `eventReferences` | `List<EventReference>` | no |
| `eventTime` | `timestamp` | yes |
| `eventType` | `string` | yes |
| `incidentRecordArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `incidentRecordArn` | `string` | yes |

## DeleteIncidentRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReplicationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyId` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResponsePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTimelineEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `incidentRecordArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetIncidentRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `incidentRecord` | `IncidentRecord` | yes |

## GetReplicationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationSet` | `ReplicationSet` | yes |

## GetResourcePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resourcePolicies` | `List<ResourcePolicy>` | yes |

## GetResponsePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actions` | `List<Action>` | no |
| `arn` | `string` | yes |
| `chatChannel` | `ChatChannel` | no |
| `displayName` | `string` | no |
| `engagements` | `List<string>` | no |
| `incidentTemplate` | `IncidentTemplate` | yes |
| `integrations` | `List<Integration>` | no |
| `name` | `string` | yes |

## GetTimelineEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `incidentRecordArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `event` | `TimelineEvent` | yes |

## ListIncidentFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `incidentRecordArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<FindingSummary>` | yes |
| `nextToken` | `string` | no |

## ListIncidentRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `incidentRecordSummaries` | `List<IncidentRecordSummary>` | yes |
| `nextToken` | `string` | no |

## ListRelatedItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `incidentRecordArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `relatedItems` | `List<RelatedItem>` | yes |

## ListReplicationSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `replicationSetArns` | `List<string>` | yes |

## ListResponsePlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `responsePlanSummaries` | `List<ResponsePlanSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## ListTimelineEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<Filter>` | no |
| `incidentRecordArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventSummaries` | `List<EventSummary>` | yes |
| `nextToken` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyId` | `string` | yes |

## StartIncident

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `impact` | `integer` | no |
| `relatedItems` | `List<RelatedItem>` | no |
| `responsePlanArn` | `string` | yes |
| `title` | `string` | no |
| `triggerDetails` | `TriggerDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `incidentRecordArn` | `string` | yes |

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


## UpdateDeletionProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `deletionProtected` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateIncidentRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `chatChannel` | `ChatChannel` | no |
| `clientToken` | `string` | no |
| `impact` | `integer` | no |
| `notificationTargets` | `List<NotificationTargetItem>` | no |
| `status` | `string` | no |
| `summary` | `string` | no |
| `title` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRelatedItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `incidentRecordArn` | `string` | yes |
| `relatedItemsUpdate` | `RelatedItemsUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateReplicationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actions` | `List<UpdateReplicationSetAction>` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResponsePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actions` | `List<Action>` | no |
| `arn` | `string` | yes |
| `chatChannel` | `ChatChannel` | no |
| `clientToken` | `string` | no |
| `displayName` | `string` | no |
| `engagements` | `List<string>` | no |
| `incidentTemplateDedupeString` | `string` | no |
| `incidentTemplateImpact` | `integer` | no |
| `incidentTemplateNotificationTargets` | `List<NotificationTargetItem>` | no |
| `incidentTemplateSummary` | `string` | no |
| `incidentTemplateTags` | `Map<string>` | no |
| `incidentTemplateTitle` | `string` | no |
| `integrations` | `List<Integration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTimelineEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `eventData` | `string` | no |
| `eventId` | `string` | yes |
| `eventReferences` | `List<EventReference>` | no |
| `eventTime` | `timestamp` | no |
| `eventType` | `string` | no |
| `incidentRecordArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


