# Amazon Verified Permissions

API version: 2021-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/verifiedpermissions/2021-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requests` | `List<BatchGetPolicyInputItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `List<BatchGetPolicyOutputItem>` | yes |
| `errors` | `List<BatchGetPolicyErrorItem>` | yes |

## BatchIsAuthorized

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `entities` | `EntitiesDefinition` | no |
| `requests` | `List<BatchIsAuthorizedInputItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `List<BatchIsAuthorizedOutputItem>` | yes |

## BatchIsAuthorizedWithToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `identityToken` | `string` | no |
| `accessToken` | `string` | no |
| `entities` | `EntitiesDefinition` | no |
| `requests` | `List<BatchIsAuthorizedWithTokenInputItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principal` | `EntityIdentifier` | no |
| `results` | `List<BatchIsAuthorizedWithTokenOutputItem>` | yes |

## CreateIdentitySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `policyStoreId` | `string` | yes |
| `configuration` | `Configuration` | yes |
| `principalEntityType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdDate` | `timestamp` | yes |
| `identitySourceId` | `string` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `policyStoreId` | `string` | yes |

## CreatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `policyStoreId` | `string` | yes |
| `definition` | `PolicyDefinition` | yes |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyId` | `string` | yes |
| `policyType` | `string` | yes |
| `principal` | `EntityIdentifier` | no |
| `resource` | `EntityIdentifier` | no |
| `actions` | `List<ActionIdentifier>` | no |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `effect` | `string` | no |

## CreatePolicyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `validationSettings` | `ValidationSettings` | yes |
| `description` | `string` | no |
| `deletionProtection` | `string` | no |
| `encryptionSettings` | `EncryptionSettings` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `arn` | `string` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |

## CreatePolicyStoreAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aliasName` | `string` | yes |
| `policyStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aliasName` | `string` | yes |
| `policyStoreId` | `string` | yes |
| `aliasArn` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreatePolicyTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `policyStoreId` | `string` | yes |
| `description` | `string` | no |
| `statement` | `string` | yes |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyTemplateId` | `string` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |

## DeleteIdentitySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `identitySourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicyStoreAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aliasName` | `string` | yes |
| `deletionMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicyTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyTemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetIdentitySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `identitySourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdDate` | `timestamp` | yes |
| `details` | `IdentitySourceDetails` | no |
| `identitySourceId` | `string` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `policyStoreId` | `string` | yes |
| `principalEntityType` | `string` | yes |
| `configuration` | `ConfigurationDetail` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyId` | `string` | yes |
| `policyType` | `string` | yes |
| `principal` | `EntityIdentifier` | no |
| `resource` | `EntityIdentifier` | no |
| `actions` | `List<ActionIdentifier>` | no |
| `definition` | `PolicyDefinitionDetail` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `effect` | `string` | no |
| `name` | `string` | no |

## GetPolicyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `tags` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `arn` | `string` | yes |
| `validationSettings` | `ValidationSettings` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `description` | `string` | no |
| `deletionProtection` | `string` | no |
| `encryptionState` | `EncryptionState` | no |
| `cedarVersion` | `string` | no |
| `tags` | `Map<string>` | no |

## GetPolicyStoreAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aliasName` | `string` | yes |
| `policyStoreId` | `string` | yes |
| `aliasArn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `state` | `string` | yes |

## GetPolicyTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyTemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyTemplateId` | `string` | yes |
| `description` | `string` | no |
| `statement` | `string` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `name` | `string` | no |

## GetSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `schema` | `string` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `namespaces` | `List<string>` | no |

## IsAuthorized

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `principal` | `EntityIdentifier` | no |
| `action` | `ActionIdentifier` | no |
| `resource` | `EntityIdentifier` | no |
| `context` | `ContextDefinition` | no |
| `entities` | `EntitiesDefinition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `decision` | `string` | yes |
| `determiningPolicies` | `List<DeterminingPolicyItem>` | yes |
| `errors` | `List<EvaluationErrorItem>` | yes |

## IsAuthorizedWithToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `identityToken` | `string` | no |
| `accessToken` | `string` | no |
| `action` | `ActionIdentifier` | no |
| `resource` | `EntityIdentifier` | no |
| `context` | `ContextDefinition` | no |
| `entities` | `EntitiesDefinition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `decision` | `string` | yes |
| `determiningPolicies` | `List<DeterminingPolicyItem>` | yes |
| `errors` | `List<EvaluationErrorItem>` | yes |
| `principal` | `EntityIdentifier` | no |

## ListIdentitySources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<IdentitySourceFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `identitySources` | `List<IdentitySourceItem>` | yes |

## ListPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `PolicyFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `policies` | `List<PolicyItem>` | yes |

## ListPolicyStoreAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `PolicyStoreAliasFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `policyStoreAliases` | `List<PolicyStoreAliasItem>` | yes |

## ListPolicyStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `policyStores` | `List<PolicyStoreItem>` | yes |

## ListPolicyTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `policyTemplates` | `List<PolicyTemplateItem>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `definition` | `SchemaDefinition` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `namespaces` | `List<string>` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |

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


## UpdateIdentitySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `identitySourceId` | `string` | yes |
| `updateConfiguration` | `UpdateConfiguration` | yes |
| `principalEntityType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdDate` | `timestamp` | yes |
| `identitySourceId` | `string` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `policyStoreId` | `string` | yes |

## UpdatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyId` | `string` | yes |
| `definition` | `UpdatePolicyDefinition` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyId` | `string` | yes |
| `policyType` | `string` | yes |
| `principal` | `EntityIdentifier` | no |
| `resource` | `EntityIdentifier` | no |
| `actions` | `List<ActionIdentifier>` | no |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |
| `effect` | `string` | no |

## UpdatePolicyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `validationSettings` | `ValidationSettings` | yes |
| `deletionProtection` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `arn` | `string` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |

## UpdatePolicyTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyTemplateId` | `string` | yes |
| `description` | `string` | no |
| `statement` | `string` | yes |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyStoreId` | `string` | yes |
| `policyTemplateId` | `string` | yes |
| `createdDate` | `timestamp` | yes |
| `lastUpdatedDate` | `timestamp` | yes |

