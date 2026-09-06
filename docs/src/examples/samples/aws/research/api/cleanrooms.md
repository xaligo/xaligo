# AWS Clean Rooms Service

API version: 2022-02-17. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cleanrooms/2022-02-17/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetCollaborationAnalysisTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `analysisTemplateArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationAnalysisTemplates` | `List<CollaborationAnalysisTemplate>` | yes |
| `errors` | `List<BatchGetCollaborationAnalysisTemplateError>` | yes |

## BatchGetSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `names` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemas` | `List<Schema>` | yes |
| `errors` | `List<BatchGetSchemaError>` | yes |

## BatchGetSchemaAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `schemaAnalysisRuleRequests` | `List<SchemaAnalysisRuleRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRules` | `List<AnalysisRule>` | yes |
| `errors` | `List<BatchGetSchemaAnalysisRuleError>` | yes |

## CreateAnalysisTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `membershipIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `format` | `string` | yes |
| `source` | `AnalysisSource` | yes |
| `tags` | `Map<string>` | no |
| `analysisParameters` | `List<AnalysisParameter>` | no |
| `schema` | `AnalysisSchema` | no |
| `errorMessageConfiguration` | `ErrorMessageConfiguration` | no |
| `syntheticDataParameters` | `SyntheticDataParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisTemplate` | `AnalysisTemplate` | yes |

## CreateCollaboration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `members` | `List<MemberSpecification>` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `creatorMemberAbilities` | `List<string>` | yes |
| `creatorMLMemberAbilities` | `MLMemberAbilities` | no |
| `creatorDisplayName` | `string` | yes |
| `dataEncryptionMetadata` | `DataEncryptionMetadata` | no |
| `queryLogStatus` | `string` | yes |
| `jobLogStatus` | `string` | no |
| `tags` | `Map<string>` | no |
| `creatorPaymentConfiguration` | `PaymentConfiguration` | no |
| `analyticsEngine` | `string` | no |
| `autoApprovedChangeRequestTypes` | `List<string>` | no |
| `allowedResultRegions` | `List<string>` | no |
| `isMetricsEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaboration` | `Collaboration` | yes |

## CreateCollaborationChangeRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `changes` | `List<ChangeInput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationChangeRequest` | `CollaborationChangeRequest` | yes |

## CreateConfiguredAudienceModelAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `configuredAudienceModelArn` | `string` | yes |
| `configuredAudienceModelAssociationName` | `string` | yes |
| `manageResourcePolicies` | `boolean` | yes |
| `tags` | `Map<string>` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelAssociation` | `ConfiguredAudienceModelAssociation` | yes |

## CreateConfiguredTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `tableReference` | `TableReference` | yes |
| `allowedColumns` | `List<string>` | yes |
| `analysisMethod` | `string` | yes |
| `selectedAnalysisMethods` | `List<string>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTable` | `ConfiguredTable` | yes |

## CreateConfiguredTableAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |
| `analysisRulePolicy` | `ConfiguredTableAnalysisRulePolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `ConfiguredTableAnalysisRule` | yes |

## CreateConfiguredTableAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `membershipIdentifier` | `string` | yes |
| `configuredTableIdentifier` | `string` | yes |
| `roleArn` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableAssociation` | `ConfiguredTableAssociation` | yes |

## CreateConfiguredTableAssociationAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `configuredTableAssociationIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |
| `analysisRulePolicy` | `ConfiguredTableAssociationAnalysisRulePolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `ConfiguredTableAssociationAnalysisRule` | yes |

## CreateIdMappingTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `inputReferenceConfig` | `IdMappingTableInputReferenceConfig` | yes |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingTable` | `IdMappingTable` | yes |

## CreateIdNamespaceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `inputReferenceConfig` | `IdNamespaceAssociationInputReferenceConfig` | yes |
| `tags` | `Map<string>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `idMappingConfig` | `IdMappingConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceAssociation` | `IdNamespaceAssociation` | yes |

## CreateIntermediateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `populationAnalysisConfiguration` | `PopulationAnalysisConfiguration` | yes |
| `kmsKeyArn` | `string` | no |
| `retentionInDays` | `integer` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intermediateTable` | `IntermediateTable` | yes |

## CreateIntermediateTableAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `intermediateTableIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |
| `analysisRulePolicy` | `IntermediateTableAnalysisRulePolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `IntermediateTableAnalysisRule` | yes |

## CreateMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `queryLogStatus` | `string` | yes |
| `jobLogStatus` | `string` | no |
| `tags` | `Map<string>` | no |
| `defaultResultConfiguration` | `MembershipProtectedQueryResultConfiguration` | no |
| `defaultJobResultConfiguration` | `MembershipProtectedJobResultConfiguration` | no |
| `paymentConfiguration` | `MembershipPaymentConfiguration` | no |
| `isMetricsEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membership` | `Membership` | yes |

## CreatePrivacyBudgetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `autoRefresh` | `string` | no |
| `privacyBudgetType` | `string` | yes |
| `parameters` | `PrivacyBudgetTemplateParametersInput` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privacyBudgetTemplate` | `PrivacyBudgetTemplate` | yes |

## DeleteAnalysisTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `analysisTemplateIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCollaboration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredAudienceModelAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredTableAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredTableAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredTableAssociationAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `configuredTableAssociationIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIdMappingTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingTableIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIdNamespaceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntermediateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `intermediateTableIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntermediateTableAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `intermediateTableIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePrivacyBudgetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `privacyBudgetTemplateIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisallowIntermediateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `intermediateTableName` | `string` | yes |
| `includeDescendants` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAnalysisLogExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `analysisLogExportIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisLogExport` | `AnalysisLogExport` | yes |

## GetAnalysisTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `analysisTemplateIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisTemplate` | `AnalysisTemplate` | yes |

## GetCollaboration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaboration` | `Collaboration` | yes |

## GetCollaborationAnalysisTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `analysisTemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationAnalysisTemplate` | `CollaborationAnalysisTemplate` | yes |

## GetCollaborationChangeRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `changeRequestIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationChangeRequest` | `CollaborationChangeRequest` | yes |

## GetCollaborationConfiguredAudienceModelAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `configuredAudienceModelAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationConfiguredAudienceModelAssociation` | `CollaborationConfiguredAudienceModelAssociation` | yes |

## GetCollaborationIdNamespaceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `idNamespaceAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdNamespaceAssociation` | `CollaborationIdNamespaceAssociation` | yes |

## GetCollaborationPrivacyBudgetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `privacyBudgetTemplateIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationPrivacyBudgetTemplate` | `CollaborationPrivacyBudgetTemplate` | yes |

## GetConfiguredAudienceModelAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelAssociation` | `ConfiguredAudienceModelAssociation` | yes |

## GetConfiguredTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTable` | `ConfiguredTable` | yes |

## GetConfiguredTableAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `ConfiguredTableAnalysisRule` | yes |

## GetConfiguredTableAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableAssociation` | `ConfiguredTableAssociation` | yes |

## GetConfiguredTableAssociationAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `configuredTableAssociationIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `ConfiguredTableAssociationAnalysisRule` | yes |

## GetIdMappingTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingTableIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingTable` | `IdMappingTable` | yes |

## GetIdNamespaceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceAssociation` | `IdNamespaceAssociation` | yes |

## GetIntermediateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intermediateTableIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intermediateTable` | `IntermediateTable` | yes |

## GetIntermediateTableAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `intermediateTableIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `IntermediateTableAnalysisRule` | yes |

## GetMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membership` | `Membership` | yes |

## GetPrivacyBudgetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `privacyBudgetTemplateIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privacyBudgetTemplate` | `PrivacyBudgetTemplate` | yes |

## GetProtectedJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `protectedJobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protectedJob` | `ProtectedJob` | yes |

## GetProtectedQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `protectedQueryIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protectedQuery` | `ProtectedQuery` | yes |

## GetSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schema` | `Schema` | yes |

## GetSchemaAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `AnalysisRule` | yes |

## ListAnalysisLogExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `analysisIdentifier` | `string` | no |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `analysisLogExports` | `List<AnalysisLogExportSummary>` | yes |

## ListAnalysisTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `analysisTemplateSummaries` | `List<AnalysisTemplateSummary>` | yes |

## ListCollaborationAnalysisTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationAnalysisTemplateSummaries` | `List<CollaborationAnalysisTemplateSummary>` | yes |

## ListCollaborationChangeRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationChangeRequestSummaries` | `List<CollaborationChangeRequestSummary>` | yes |
| `nextToken` | `string` | no |

## ListCollaborationConfiguredAudienceModelAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationConfiguredAudienceModelAssociationSummaries` | `List<CollaborationConfiguredAudienceModelAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListCollaborationIdNamespaceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationIdNamespaceAssociationSummaries` | `List<CollaborationIdNamespaceAssociationSummary>` | yes |

## ListCollaborationPrivacyBudgetTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationPrivacyBudgetTemplateSummaries` | `List<CollaborationPrivacyBudgetTemplateSummary>` | yes |

## ListCollaborationPrivacyBudgets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `privacyBudgetType` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `accessBudgetResourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationPrivacyBudgetSummaries` | `List<CollaborationPrivacyBudgetSummary>` | yes |
| `nextToken` | `string` | no |

## ListCollaborations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `memberStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationList` | `List<CollaborationSummary>` | yes |

## ListConfiguredAudienceModelAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelAssociationSummaries` | `List<ConfiguredAudienceModelAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListConfiguredTableAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableAssociationSummaries` | `List<ConfiguredTableAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListConfiguredTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableSummaries` | `List<ConfiguredTableSummary>` | yes |
| `nextToken` | `string` | no |

## ListIdMappingTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingTableSummaries` | `List<IdMappingTableSummary>` | yes |
| `nextToken` | `string` | no |

## ListIdNamespaceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `idNamespaceAssociationSummaries` | `List<IdNamespaceAssociationSummary>` | yes |

## ListIntermediateTableVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `intermediateTableIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intermediateTableVersionSummaries` | `List<IntermediateTableVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListIntermediateTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intermediateTableSummaries` | `List<IntermediateTableSummary>` | yes |
| `nextToken` | `string` | no |

## ListMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `memberSummaries` | `List<MemberSummary>` | yes |

## ListMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `membershipSummaries` | `List<MembershipSummary>` | yes |

## ListPrivacyBudgetTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `privacyBudgetTemplateSummaries` | `List<PrivacyBudgetTemplateSummary>` | yes |

## ListPrivacyBudgets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `privacyBudgetType` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `accessBudgetResourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privacyBudgetSummaries` | `List<PrivacyBudgetSummary>` | yes |
| `nextToken` | `string` | no |

## ListProtectedJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `protectedJobs` | `List<ProtectedJobSummary>` | yes |

## ListProtectedQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `protectedQueries` | `List<ProtectedQuerySummary>` | yes |

## ListSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `schemaType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaSummaries` | `List<SchemaSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## PopulateIdMappingTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingTableIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `jobType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingJobId` | `string` | yes |

## PopulateIntermediateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intermediateTableIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `parameters` | `Map<string>` | no |
| `computeConfiguration` | `IntermediateTableComputeConfiguration` | no |
| `analysisPayerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisId` | `string` | yes |
| `analysisType` | `string` | yes |
| `versionId` | `string` | yes |

## PreviewPrivacyImpact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `parameters` | `PreviewPrivacyImpactParametersInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privacyImpact` | `PrivacyImpact` | yes |

## StartAnalysisLogExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `analysisId` | `string` | yes |
| `analysisType` | `string` | yes |
| `resultConfiguration` | `AnalysisLogExportResultConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisLogExport` | `AnalysisLogExport` | yes |

## StartProtectedJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `jobParameters` | `ProtectedJobParameters` | yes |
| `resultConfiguration` | `ProtectedJobResultConfigurationInput` | no |
| `computeConfiguration` | `ProtectedJobComputeConfiguration` | no |
| `jobComputePayerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protectedJob` | `ProtectedJob` | yes |

## StartProtectedQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `sqlParameters` | `ProtectedQuerySQLParameters` | yes |
| `resultConfiguration` | `ProtectedQueryResultConfiguration` | no |
| `computeConfiguration` | `ComputeConfiguration` | no |
| `queryComputePayerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protectedQuery` | `ProtectedQuery` | yes |

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


## UpdateAnalysisTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `analysisTemplateIdentifier` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisTemplate` | `AnalysisTemplate` | yes |

## UpdateCollaboration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `analyticsEngine` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaboration` | `Collaboration` | yes |

## UpdateCollaborationChangeRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationIdentifier` | `string` | yes |
| `changeRequestIdentifier` | `string` | yes |
| `action` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collaborationChangeRequest` | `CollaborationChangeRequest` | yes |

## UpdateConfiguredAudienceModelAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `description` | `string` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelAssociation` | `ConfiguredAudienceModelAssociation` | yes |

## UpdateConfiguredTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableIdentifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `tableReference` | `TableReference` | no |
| `allowedColumns` | `List<string>` | no |
| `analysisMethod` | `string` | no |
| `selectedAnalysisMethods` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTable` | `ConfiguredTable` | yes |

## UpdateConfiguredTableAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |
| `analysisRulePolicy` | `ConfiguredTableAnalysisRulePolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `ConfiguredTableAnalysisRule` | yes |

## UpdateConfiguredTableAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredTableAssociation` | `ConfiguredTableAssociation` | yes |

## UpdateConfiguredTableAssociationAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `configuredTableAssociationIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |
| `analysisRulePolicy` | `ConfiguredTableAssociationAnalysisRulePolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `ConfiguredTableAssociationAnalysisRule` | yes |

## UpdateIdMappingTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingTableIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `description` | `string` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idMappingTable` | `IdMappingTable` | yes |

## UpdateIdNamespaceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceAssociationIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `idMappingConfig` | `IdMappingConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceAssociation` | `IdNamespaceAssociation` | yes |

## UpdateIntermediateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intermediateTableIdentifier` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `description` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `columns` | `List<IntermediateTableColumn>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intermediateTable` | `IntermediateTable` | yes |

## UpdateIntermediateTableAnalysisRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `intermediateTableIdentifier` | `string` | yes |
| `analysisRuleType` | `string` | yes |
| `analysisRulePolicy` | `IntermediateTableAnalysisRulePolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analysisRule` | `IntermediateTableAnalysisRule` | yes |

## UpdateMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `queryLogStatus` | `string` | no |
| `jobLogStatus` | `string` | no |
| `defaultResultConfiguration` | `MembershipProtectedQueryResultConfiguration` | no |
| `defaultJobResultConfiguration` | `MembershipProtectedJobResultConfiguration` | no |
| `membershipPaymentConfiguration` | `UpdateMembershipPaymentConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membership` | `Membership` | yes |

## UpdatePrivacyBudgetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `privacyBudgetTemplateIdentifier` | `string` | yes |
| `privacyBudgetType` | `string` | yes |
| `parameters` | `PrivacyBudgetTemplateUpdateParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privacyBudgetTemplate` | `PrivacyBudgetTemplate` | yes |

## UpdateProtectedJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `protectedJobIdentifier` | `string` | yes |
| `targetStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protectedJob` | `ProtectedJob` | yes |

## UpdateProtectedQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `protectedQueryIdentifier` | `string` | yes |
| `targetStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protectedQuery` | `ProtectedQuery` | yes |

