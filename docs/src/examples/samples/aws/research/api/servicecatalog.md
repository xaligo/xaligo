# AWS Service Catalog

API version: 2015-12-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/servicecatalog/2015-12-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptPortfolioShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `PortfolioShareType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateBudgetWithResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BudgetName` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociatePrincipalWithPortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `PrincipalARN` | `string` | yes |
| `PrincipalType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateProductWithPortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |
| `PortfolioId` | `string` | yes |
| `SourcePortfolioId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateServiceActionWithProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductId` | `string` | yes |
| `ProvisioningArtifactId` | `string` | yes |
| `ServiceActionId` | `string` | yes |
| `AcceptLanguage` | `string` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateTagOptionWithResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagOptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchAssociateServiceActionWithProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionAssociations` | `List<ServiceActionAssociation>` | yes |
| `AcceptLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedServiceActionAssociations` | `List<FailedServiceActionAssociation>` | no |

## BatchDisassociateServiceActionFromProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionAssociations` | `List<ServiceActionAssociation>` | yes |
| `AcceptLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedServiceActionAssociations` | `List<FailedServiceActionAssociation>` | no |

## CopyProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `SourceProductArn` | `string` | yes |
| `TargetProductId` | `string` | no |
| `TargetProductName` | `string` | no |
| `SourceProvisioningArtifactIdentifiers` | `List<Map<string>>` | no |
| `CopyOptions` | `List<string>` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CopyProductToken` | `string` | no |

## CreateConstraint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `ProductId` | `string` | yes |
| `Parameters` | `string` | yes |
| `Type` | `string` | yes |
| `Description` | `string` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConstraintDetail` | `ConstraintDetail` | no |
| `ConstraintParameters` | `string` | no |
| `Status` | `string` | no |

## CreatePortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `DisplayName` | `string` | yes |
| `Description` | `string` | no |
| `ProviderName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioDetail` | `PortfolioDetail` | no |
| `Tags` | `List<Tag>` | no |

## CreatePortfolioShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `AccountId` | `string` | no |
| `OrganizationNode` | `OrganizationNode` | no |
| `ShareTagOptions` | `boolean` | no |
| `SharePrincipals` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioShareToken` | `string` | no |

## CreateProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Name` | `string` | yes |
| `Owner` | `string` | yes |
| `Description` | `string` | no |
| `Distributor` | `string` | no |
| `SupportDescription` | `string` | no |
| `SupportEmail` | `string` | no |
| `SupportUrl` | `string` | no |
| `ProductType` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ProvisioningArtifactParameters` | `ProvisioningArtifactProperties` | no |
| `IdempotencyToken` | `string` | yes |
| `SourceConnection` | `SourceConnection` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductViewDetail` | `ProductViewDetail` | no |
| `ProvisioningArtifactDetail` | `ProvisioningArtifactDetail` | no |
| `Tags` | `List<Tag>` | no |

## CreateProvisionedProductPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PlanName` | `string` | yes |
| `PlanType` | `string` | yes |
| `NotificationArns` | `List<string>` | no |
| `PathId` | `string` | no |
| `ProductId` | `string` | yes |
| `ProvisionedProductName` | `string` | yes |
| `ProvisioningArtifactId` | `string` | yes |
| `ProvisioningParameters` | `List<UpdateProvisioningParameter>` | no |
| `IdempotencyToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlanName` | `string` | no |
| `PlanId` | `string` | no |
| `ProvisionProductId` | `string` | no |
| `ProvisionedProductName` | `string` | no |
| `ProvisioningArtifactId` | `string` | no |

## CreateProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |
| `Parameters` | `ProvisioningArtifactProperties` | yes |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisioningArtifactDetail` | `ProvisioningArtifactDetail` | no |
| `Info` | `Map<string>` | no |
| `Status` | `string` | no |

## CreateServiceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DefinitionType` | `string` | yes |
| `Definition` | `Map<string>` | yes |
| `Description` | `string` | no |
| `AcceptLanguage` | `string` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionDetail` | `ServiceActionDetail` | no |

## CreateTagOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `string` | yes |
| `Value` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagOptionDetail` | `TagOptionDetail` | no |

## DeleteConstraint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePortfolioShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `AccountId` | `string` | no |
| `OrganizationNode` | `OrganizationNode` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioShareToken` | `string` | no |

## DeleteProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProvisionedProductPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PlanId` | `string` | yes |
| `IgnoreErrors` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |
| `ProvisioningArtifactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `AcceptLanguage` | `string` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTagOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeConstraint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConstraintDetail` | `ConstraintDetail` | no |
| `ConstraintParameters` | `string` | no |
| `Status` | `string` | no |

## DescribeCopyProductStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `CopyProductToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CopyProductStatus` | `string` | no |
| `TargetProductId` | `string` | no |
| `StatusDetail` | `string` | no |

## DescribePortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioDetail` | `PortfolioDetail` | no |
| `Tags` | `List<Tag>` | no |
| `TagOptions` | `List<TagOptionDetail>` | no |
| `Budgets` | `List<BudgetDetail>` | no |

## DescribePortfolioShareStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioShareToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioShareToken` | `string` | no |
| `PortfolioId` | `string` | no |
| `OrganizationNodeValue` | `string` | no |
| `Status` | `string` | no |
| `ShareDetails` | `ShareDetails` | no |

## DescribePortfolioShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioId` | `string` | yes |
| `Type` | `string` | yes |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextPageToken` | `string` | no |
| `PortfolioShareDetails` | `List<PortfolioShareDetail>` | no |

## DescribeProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductViewSummary` | `ProductViewSummary` | no |
| `ProvisioningArtifacts` | `List<ProvisioningArtifact>` | no |
| `Budgets` | `List<BudgetDetail>` | no |
| `LaunchPaths` | `List<LaunchPath>` | no |

## DescribeProductAsAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `SourcePortfolioId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductViewDetail` | `ProductViewDetail` | no |
| `ProvisioningArtifactSummaries` | `List<ProvisioningArtifactSummary>` | no |
| `Tags` | `List<Tag>` | no |
| `TagOptions` | `List<TagOptionDetail>` | no |
| `Budgets` | `List<BudgetDetail>` | no |

## DescribeProductView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductViewSummary` | `ProductViewSummary` | no |
| `ProvisioningArtifacts` | `List<ProvisioningArtifact>` | no |

## DescribeProvisionedProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProductDetail` | `ProvisionedProductDetail` | no |
| `CloudWatchDashboards` | `List<CloudWatchDashboard>` | no |

## DescribeProvisionedProductPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PlanId` | `string` | yes |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProductPlanDetails` | `ProvisionedProductPlanDetails` | no |
| `ResourceChanges` | `List<ResourceChange>` | no |
| `NextPageToken` | `string` | no |

## DescribeProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProvisioningArtifactId` | `string` | no |
| `ProductId` | `string` | no |
| `ProvisioningArtifactName` | `string` | no |
| `ProductName` | `string` | no |
| `Verbose` | `boolean` | no |
| `IncludeProvisioningArtifactParameters` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisioningArtifactDetail` | `ProvisioningArtifactDetail` | no |
| `Info` | `Map<string>` | no |
| `Status` | `string` | no |
| `ProvisioningArtifactParameters` | `List<ProvisioningArtifactParameter>` | no |

## DescribeProvisioningParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | no |
| `ProductName` | `string` | no |
| `ProvisioningArtifactId` | `string` | no |
| `ProvisioningArtifactName` | `string` | no |
| `PathId` | `string` | no |
| `PathName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisioningArtifactParameters` | `List<ProvisioningArtifactParameter>` | no |
| `ConstraintSummaries` | `List<ConstraintSummary>` | no |
| `UsageInstructions` | `List<UsageInstruction>` | no |
| `TagOptions` | `List<TagOptionSummary>` | no |
| `ProvisioningArtifactPreferences` | `ProvisioningArtifactPreferences` | no |
| `ProvisioningArtifactOutputs` | `List<ProvisioningArtifactOutput>` | no |
| `ProvisioningArtifactOutputKeys` | `List<ProvisioningArtifactOutput>` | no |

## DescribeRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordDetail` | `RecordDetail` | no |
| `RecordOutputs` | `List<RecordOutput>` | no |
| `NextPageToken` | `string` | no |

## DescribeServiceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `AcceptLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionDetail` | `ServiceActionDetail` | no |

## DescribeServiceActionExecutionParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProductId` | `string` | yes |
| `ServiceActionId` | `string` | yes |
| `AcceptLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionParameters` | `List<ExecutionParameter>` | no |

## DescribeTagOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagOptionDetail` | `TagOptionDetail` | no |

## DisableAWSOrganizationsAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateBudgetFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BudgetName` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociatePrincipalFromPortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `PrincipalARN` | `string` | yes |
| `PrincipalType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateProductFromPortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |
| `PortfolioId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateServiceActionFromProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductId` | `string` | yes |
| `ProvisioningArtifactId` | `string` | yes |
| `ServiceActionId` | `string` | yes |
| `AcceptLanguage` | `string` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateTagOptionFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagOptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableAWSOrganizationsAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExecuteProvisionedProductPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PlanId` | `string` | yes |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordDetail` | `RecordDetail` | no |

## ExecuteProvisionedProductServiceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProductId` | `string` | yes |
| `ServiceActionId` | `string` | yes |
| `ExecuteToken` | `string` | yes |
| `AcceptLanguage` | `string` | no |
| `Parameters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordDetail` | `RecordDetail` | no |

## GetAWSOrganizationsAccessStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessStatus` | `string` | no |

## GetProvisionedProductOutputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProvisionedProductId` | `string` | no |
| `ProvisionedProductName` | `string` | no |
| `OutputKeys` | `List<string>` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Outputs` | `List<RecordOutput>` | no |
| `NextPageToken` | `string` | no |

## ImportAsProvisionedProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |
| `ProvisioningArtifactId` | `string` | yes |
| `ProvisionedProductName` | `string` | yes |
| `PhysicalId` | `string` | yes |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordDetail` | `RecordDetail` | no |

## ListAcceptedPortfolioShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |
| `PortfolioShareType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioDetails` | `List<PortfolioDetail>` | no |
| `NextPageToken` | `string` | no |

## ListBudgetsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ResourceId` | `string` | yes |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Budgets` | `List<BudgetDetail>` | no |
| `NextPageToken` | `string` | no |

## ListConstraintsForPortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `ProductId` | `string` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConstraintDetails` | `List<ConstraintDetail>` | no |
| `NextPageToken` | `string` | no |

## ListLaunchPaths

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchPathSummaries` | `List<LaunchPathSummary>` | no |
| `NextPageToken` | `string` | no |

## ListOrganizationPortfolioAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `OrganizationNodeType` | `string` | yes |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationNodes` | `List<OrganizationNode>` | no |
| `NextPageToken` | `string` | no |

## ListPortfolioAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `OrganizationParentId` | `string` | no |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | no |
| `NextPageToken` | `string` | no |

## ListPortfolios

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioDetails` | `List<PortfolioDetail>` | no |
| `NextPageToken` | `string` | no |

## ListPortfoliosForProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioDetails` | `List<PortfolioDetail>` | no |
| `NextPageToken` | `string` | no |

## ListPrincipalsForPortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Principals` | `List<Principal>` | no |
| `NextPageToken` | `string` | no |

## ListProvisionedProductPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProvisionProductId` | `string` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |
| `AccessLevelFilter` | `AccessLevelFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProductPlans` | `List<ProvisionedProductPlanSummary>` | no |
| `NextPageToken` | `string` | no |

## ListProvisioningArtifacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisioningArtifactDetails` | `List<ProvisioningArtifactDetail>` | no |
| `NextPageToken` | `string` | no |

## ListProvisioningArtifactsForServiceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionId` | `string` | yes |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |
| `AcceptLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisioningArtifactViews` | `List<ProvisioningArtifactView>` | no |
| `NextPageToken` | `string` | no |

## ListRecordHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `AccessLevelFilter` | `AccessLevelFilter` | no |
| `SearchFilter` | `ListRecordHistorySearchFilter` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordDetails` | `List<RecordDetail>` | no |
| `NextPageToken` | `string` | no |

## ListResourcesForTagOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagOptionId` | `string` | yes |
| `ResourceType` | `string` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceDetails` | `List<ResourceDetail>` | no |
| `PageToken` | `string` | no |

## ListServiceActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionSummaries` | `List<ServiceActionSummary>` | no |
| `NextPageToken` | `string` | no |

## ListServiceActionsForProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductId` | `string` | yes |
| `ProvisioningArtifactId` | `string` | yes |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |
| `AcceptLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionSummaries` | `List<ServiceActionSummary>` | no |
| `NextPageToken` | `string` | no |

## ListStackInstancesForProvisionedProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProvisionedProductId` | `string` | yes |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackInstances` | `List<StackInstance>` | no |
| `NextPageToken` | `string` | no |

## ListTagOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `ListTagOptionsFilters` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagOptionDetails` | `List<TagOptionDetail>` | no |
| `PageToken` | `string` | no |

## NotifyProvisionProductEngineWorkflowResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowToken` | `string` | yes |
| `RecordId` | `string` | yes |
| `Status` | `string` | yes |
| `FailureReason` | `string` | no |
| `ResourceIdentifier` | `EngineWorkflowResourceIdentifier` | no |
| `Outputs` | `List<RecordOutput>` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## NotifyTerminateProvisionedProductEngineWorkflowResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowToken` | `string` | yes |
| `RecordId` | `string` | yes |
| `Status` | `string` | yes |
| `FailureReason` | `string` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## NotifyUpdateProvisionedProductEngineWorkflowResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowToken` | `string` | yes |
| `RecordId` | `string` | yes |
| `Status` | `string` | yes |
| `FailureReason` | `string` | no |
| `Outputs` | `List<RecordOutput>` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ProvisionProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | no |
| `ProductName` | `string` | no |
| `ProvisioningArtifactId` | `string` | no |
| `ProvisioningArtifactName` | `string` | no |
| `PathId` | `string` | no |
| `PathName` | `string` | no |
| `ProvisionedProductName` | `string` | yes |
| `ProvisioningParameters` | `List<ProvisioningParameter>` | no |
| `ProvisioningPreferences` | `ProvisioningPreferences` | no |
| `Tags` | `List<Tag>` | no |
| `NotificationArns` | `List<string>` | no |
| `ProvisionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordDetail` | `RecordDetail` | no |

## RejectPortfolioShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `PortfolioShareType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ScanProvisionedProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `AccessLevelFilter` | `AccessLevelFilter` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProducts` | `List<ProvisionedProductDetail>` | no |
| `NextPageToken` | `string` | no |

## SearchProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Filters` | `Map<List<string>>` | no |
| `PageSize` | `integer` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductViewSummaries` | `List<ProductViewSummary>` | no |
| `ProductViewAggregations` | `Map<List<ProductViewAggregationValue>>` | no |
| `NextPageToken` | `string` | no |

## SearchProductsAsAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | no |
| `Filters` | `Map<List<string>>` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `PageToken` | `string` | no |
| `PageSize` | `integer` | no |
| `ProductSource` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductViewDetails` | `List<ProductViewDetail>` | no |
| `NextPageToken` | `string` | no |

## SearchProvisionedProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `AccessLevelFilter` | `AccessLevelFilter` | no |
| `Filters` | `Map<List<string>>` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `PageSize` | `integer` | no |
| `PageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProducts` | `List<ProvisionedProductAttribute>` | no |
| `TotalResultsCount` | `integer` | no |
| `NextPageToken` | `string` | no |

## TerminateProvisionedProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProductName` | `string` | no |
| `ProvisionedProductId` | `string` | no |
| `TerminateToken` | `string` | yes |
| `IgnoreErrors` | `boolean` | no |
| `AcceptLanguage` | `string` | no |
| `RetainPhysicalResources` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordDetail` | `RecordDetail` | no |

## UpdateConstraint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |
| `Description` | `string` | no |
| `Parameters` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConstraintDetail` | `ConstraintDetail` | no |
| `ConstraintParameters` | `string` | no |
| `Status` | `string` | no |

## UpdatePortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `ProviderName` | `string` | no |
| `AddTags` | `List<Tag>` | no |
| `RemoveTags` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioDetail` | `PortfolioDetail` | no |
| `Tags` | `List<Tag>` | no |

## UpdatePortfolioShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `PortfolioId` | `string` | yes |
| `AccountId` | `string` | no |
| `OrganizationNode` | `OrganizationNode` | no |
| `ShareTagOptions` | `boolean` | no |
| `SharePrincipals` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortfolioShareToken` | `string` | no |
| `Status` | `string` | no |

## UpdateProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Owner` | `string` | no |
| `Description` | `string` | no |
| `Distributor` | `string` | no |
| `SupportDescription` | `string` | no |
| `SupportEmail` | `string` | no |
| `SupportUrl` | `string` | no |
| `AddTags` | `List<Tag>` | no |
| `RemoveTags` | `List<string>` | no |
| `SourceConnection` | `SourceConnection` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductViewDetail` | `ProductViewDetail` | no |
| `Tags` | `List<Tag>` | no |

## UpdateProvisionedProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProvisionedProductName` | `string` | no |
| `ProvisionedProductId` | `string` | no |
| `ProductId` | `string` | no |
| `ProductName` | `string` | no |
| `ProvisioningArtifactId` | `string` | no |
| `ProvisioningArtifactName` | `string` | no |
| `PathId` | `string` | no |
| `PathName` | `string` | no |
| `ProvisioningParameters` | `List<UpdateProvisioningParameter>` | no |
| `ProvisioningPreferences` | `UpdateProvisioningPreferences` | no |
| `Tags` | `List<Tag>` | no |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordDetail` | `RecordDetail` | no |

## UpdateProvisionedProductProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProvisionedProductId` | `string` | yes |
| `ProvisionedProductProperties` | `Map<string>` | yes |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedProductId` | `string` | no |
| `ProvisionedProductProperties` | `Map<string>` | no |
| `RecordId` | `string` | no |
| `Status` | `string` | no |

## UpdateProvisioningArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptLanguage` | `string` | no |
| `ProductId` | `string` | yes |
| `ProvisioningArtifactId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Active` | `boolean` | no |
| `Guidance` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisioningArtifactDetail` | `ProvisioningArtifactDetail` | no |
| `Info` | `Map<string>` | no |
| `Status` | `string` | no |

## UpdateServiceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Definition` | `Map<string>` | no |
| `Description` | `string` | no |
| `AcceptLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceActionDetail` | `ServiceActionDetail` | no |

## UpdateTagOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Value` | `string` | no |
| `Active` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagOptionDetail` | `TagOptionDetail` | no |

