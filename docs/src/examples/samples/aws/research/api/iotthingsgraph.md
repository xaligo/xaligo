# AWS IoT Things Graph

API version: 2018-09-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iotthingsgraph/2018-09-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateEntityToThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `entityId` | `string` | yes |
| `namespaceVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateFlowTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `definition` | `DefinitionDocument` | yes |
| `compatibleNamespaceVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summary` | `FlowTemplateSummary` | no |

## CreateSystemInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |
| `definition` | `DefinitionDocument` | yes |
| `target` | `string` | yes |
| `greengrassGroupName` | `string` | no |
| `s3BucketName` | `string` | no |
| `metricsConfiguration` | `MetricsConfiguration` | no |
| `flowActionsRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summary` | `SystemInstanceSummary` | no |

## CreateSystemTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `definition` | `DefinitionDocument` | yes |
| `compatibleNamespaceVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summary` | `SystemTemplateSummary` | no |

## DeleteFlowTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaceArn` | `string` | no |
| `namespaceName` | `string` | no |

## DeleteSystemInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSystemTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeploySystemInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summary` | `SystemInstanceSummary` | yes |
| `greengrassDeploymentId` | `string` | no |

## DeprecateFlowTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeprecateSystemTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaceArn` | `string` | no |
| `namespaceName` | `string` | no |
| `trackingNamespaceName` | `string` | no |
| `trackingNamespaceVersion` | `long` | no |
| `namespaceVersion` | `long` | no |

## DissociateEntityFromThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `entityType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | yes |
| `namespaceVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `descriptions` | `List<EntityDescription>` | no |

## GetFlowTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `revisionNumber` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `FlowTemplateDescription` | no |

## GetFlowTemplateRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<FlowTemplateSummary>` | no |
| `nextToken` | `string` | no |

## GetNamespaceDeletionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaceArn` | `string` | no |
| `namespaceName` | `string` | no |
| `status` | `string` | no |
| `errorCode` | `string` | no |
| `errorMessage` | `string` | no |

## GetSystemInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `SystemInstanceDescription` | no |

## GetSystemTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `revisionNumber` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `SystemTemplateDescription` | no |

## GetSystemTemplateRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<SystemTemplateSummary>` | no |
| `nextToken` | `string` | no |

## GetUploadStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadId` | `string` | yes |
| `uploadStatus` | `string` | yes |
| `namespaceArn` | `string` | no |
| `namespaceName` | `string` | no |
| `namespaceVersion` | `long` | no |
| `failureReason` | `List<string>` | no |
| `createdDate` | `timestamp` | yes |

## ListFlowExecutionMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowExecutionId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messages` | `List<FlowExecutionMessage>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `resourceArn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |
| `nextToken` | `string` | no |

## SearchEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entityTypes` | `List<string>` | yes |
| `filters` | `List<EntityFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `namespaceVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `descriptions` | `List<EntityDescription>` | no |
| `nextToken` | `string` | no |

## SearchFlowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `systemInstanceId` | `string` | yes |
| `flowExecutionId` | `string` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<FlowExecutionSummary>` | no |
| `nextToken` | `string` | no |

## SearchFlowTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<FlowTemplateFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<FlowTemplateSummary>` | no |
| `nextToken` | `string` | no |

## SearchSystemInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<SystemInstanceFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<SystemInstanceSummary>` | no |
| `nextToken` | `string` | no |

## SearchSystemTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<SystemTemplateFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<SystemTemplateSummary>` | no |
| `nextToken` | `string` | no |

## SearchThings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entityId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `namespaceVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `things` | `List<Thing>` | no |
| `nextToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UndeploySystemInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summary` | `SystemInstanceSummary` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFlowTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `definition` | `DefinitionDocument` | yes |
| `compatibleNamespaceVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summary` | `FlowTemplateSummary` | no |

## UpdateSystemTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `definition` | `DefinitionDocument` | yes |
| `compatibleNamespaceVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summary` | `SystemTemplateSummary` | no |

## UploadEntityDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `document` | `DefinitionDocument` | no |
| `syncWithPublicNamespace` | `boolean` | no |
| `deprecateExistingEntities` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadId` | `string` | yes |

