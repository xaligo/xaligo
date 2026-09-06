# Amazon Connect Cases

API version: 2022-10-03. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/connectcases/2022-10-03/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetCaseRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseRules` | `List<CaseRuleIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseRules` | `List<GetCaseRuleResponse>` | yes |
| `errors` | `List<CaseRuleError>` | yes |
| `unprocessedCaseRules` | `List<string>` | no |

## BatchGetField

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `fields` | `List<FieldIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fields` | `List<GetFieldResponse>` | yes |
| `errors` | `List<FieldError>` | yes |

## BatchPutFieldOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `fieldId` | `string` | yes |
| `options` | `List<FieldOption>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<FieldOptionError>` | no |

## CreateCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `templateId` | `string` | yes |
| `fields` | `List<FieldValue>` | yes |
| `clientToken` | `string` | no |
| `performedBy` | `UserUnion` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `caseArn` | `string` | yes |

## CreateCaseRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `rule` | `CaseRuleDetails` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseRuleId` | `string` | yes |
| `caseRuleArn` | `string` | yes |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `domainArn` | `string` | yes |
| `domainStatus` | `string` | yes |

## CreateField

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `description` | `string` | no |
| `attributes` | `FieldAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fieldId` | `string` | yes |
| `fieldArn` | `string` | yes |

## CreateLayout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `content` | `LayoutContent` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `layoutId` | `string` | yes |
| `layoutArn` | `string` | yes |

## CreateRelatedItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseId` | `string` | yes |
| `type` | `string` | yes |
| `content` | `RelatedItemInputContent` | yes |
| `performedBy` | `UserUnion` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relatedItemId` | `string` | yes |
| `relatedItemArn` | `string` | yes |

## CreateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `layoutConfiguration` | `LayoutConfiguration` | no |
| `requiredFields` | `List<RequiredField>` | no |
| `status` | `string` | no |
| `rules` | `List<TemplateRule>` | no |
| `tagPropagationConfigurations` | `List<TagPropagationConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateId` | `string` | yes |
| `templateArn` | `string` | yes |

## DeleteCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCaseRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteField

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `fieldId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLayout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `layoutId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRelatedItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseId` | `string` | yes |
| `relatedItemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `templateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `domainId` | `string` | yes |
| `fields` | `List<FieldIdentifier>` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fields` | `List<FieldValue>` | yes |
| `templateId` | `string` | yes |
| `nextToken` | `string` | no |
| `tags` | `Map<string>` | no |

## GetCaseAuditEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | yes |
| `domainId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `auditEvents` | `List<AuditEvent>` | yes |

## GetCaseEventConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventBridge` | `EventBridgeConfiguration` | yes |

## GetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `domainArn` | `string` | yes |
| `name` | `string` | yes |
| `createdTime` | `timestamp` | yes |
| `domainStatus` | `string` | yes |
| `tags` | `Map<string>` | no |

## GetLayout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `layoutId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `layoutId` | `string` | yes |
| `layoutArn` | `string` | yes |
| `name` | `string` | yes |
| `content` | `LayoutContent` | yes |
| `tags` | `Map<string>` | no |
| `deleted` | `boolean` | no |
| `createdTime` | `timestamp` | no |
| `lastModifiedTime` | `timestamp` | no |

## GetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `templateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateId` | `string` | yes |
| `templateArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `layoutConfiguration` | `LayoutConfiguration` | no |
| `requiredFields` | `List<RequiredField>` | no |
| `tags` | `Map<string>` | no |
| `status` | `string` | yes |
| `deleted` | `boolean` | no |
| `createdTime` | `timestamp` | no |
| `lastModifiedTime` | `timestamp` | no |
| `rules` | `List<TemplateRule>` | no |
| `tagPropagationConfigurations` | `List<TagPropagationConfiguration>` | no |

## ListCaseRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseRules` | `List<CaseRuleSummary>` | yes |
| `nextToken` | `string` | no |

## ListCasesForContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `contactArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cases` | `List<CaseSummary>` | yes |
| `nextToken` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domains` | `List<DomainSummary>` | yes |
| `nextToken` | `string` | no |

## ListFieldOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `fieldId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `values` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `options` | `List<FieldOption>` | yes |
| `nextToken` | `string` | no |

## ListFields

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fields` | `List<FieldSummary>` | yes |
| `nextToken` | `string` | no |

## ListLayouts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `layouts` | `List<LayoutSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `status` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templates` | `List<TemplateSummary>` | yes |
| `nextToken` | `string` | no |

## PutCaseEventConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `eventBridge` | `EventBridgeConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchAllRelatedItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `List<RelatedItemTypeFilter>` | no |
| `sorts` | `List<SearchAllRelatedItemsSort>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `relatedItems` | `List<SearchAllRelatedItemsResponseItem>` | yes |

## SearchCases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `searchTerm` | `string` | no |
| `filter` | `CaseFilter` | no |
| `sorts` | `List<Sort>` | no |
| `fields` | `List<FieldIdentifier>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `cases` | `List<SearchCasesResponseItem>` | yes |
| `totalCount` | `long` | no |

## SearchRelatedItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `List<RelatedItemTypeFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `relatedItems` | `List<SearchRelatedItemsResponseItem>` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseId` | `string` | yes |
| `fields` | `List<FieldValue>` | yes |
| `performedBy` | `UserUnion` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCaseRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseRuleId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `rule` | `CaseRuleDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateField

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `fieldId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `attributes` | `FieldAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLayout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `layoutId` | `string` | yes |
| `name` | `string` | no |
| `content` | `LayoutContent` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRelatedItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `caseId` | `string` | yes |
| `relatedItemId` | `string` | yes |
| `content` | `RelatedItemUpdateContent` | yes |
| `performedBy` | `UserUnion` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relatedItemId` | `string` | yes |
| `relatedItemArn` | `string` | yes |
| `type` | `string` | yes |
| `content` | `RelatedItemContent` | yes |
| `associationTime` | `timestamp` | yes |
| `tags` | `Map<string>` | no |
| `lastUpdatedUser` | `UserUnion` | no |
| `createdBy` | `UserUnion` | no |

## UpdateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `templateId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `layoutConfiguration` | `LayoutConfiguration` | no |
| `requiredFields` | `List<RequiredField>` | no |
| `status` | `string` | no |
| `rules` | `List<TemplateRule>` | no |
| `tagPropagationConfigurations` | `List<TagPropagationConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


