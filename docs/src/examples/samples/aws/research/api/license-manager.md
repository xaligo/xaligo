# AWS License Manager

API version: 2018-08-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/license-manager/2018-08-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | no |
| `Status` | `string` | no |
| `Version` | `string` | no |

## CheckInLicense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConsumptionToken` | `string` | yes |
| `Beneficiary` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CheckoutBorrowLicense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | yes |
| `Entitlements` | `List<EntitlementData>` | yes |
| `DigitalSignatureMethod` | `string` | yes |
| `NodeId` | `string` | no |
| `CheckoutMetadata` | `List<Metadata>` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | no |
| `LicenseConsumptionToken` | `string` | no |
| `EntitlementsAllowed` | `List<EntitlementData>` | no |
| `NodeId` | `string` | no |
| `SignedToken` | `string` | no |
| `IssuedAt` | `string` | no |
| `Expiration` | `string` | no |
| `CheckoutMetadata` | `List<Metadata>` | no |

## CheckoutLicense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductSKU` | `string` | yes |
| `CheckoutType` | `string` | yes |
| `KeyFingerprint` | `string` | yes |
| `Entitlements` | `List<EntitlementData>` | yes |
| `ClientToken` | `string` | yes |
| `Beneficiary` | `string` | no |
| `NodeId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CheckoutType` | `string` | no |
| `LicenseConsumptionToken` | `string` | no |
| `EntitlementsAllowed` | `List<EntitlementData>` | no |
| `SignedToken` | `string` | no |
| `NodeId` | `string` | no |
| `IssuedAt` | `string` | no |
| `Expiration` | `string` | no |
| `LicenseArn` | `string` | no |

## CreateGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | yes |
| `GrantName` | `string` | yes |
| `LicenseArn` | `string` | yes |
| `Principals` | `List<string>` | yes |
| `HomeRegion` | `string` | yes |
| `AllowedOperations` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | no |
| `Status` | `string` | no |
| `Version` | `string` | no |

## CreateGrantVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | yes |
| `GrantArn` | `string` | yes |
| `GrantName` | `string` | no |
| `AllowedOperations` | `List<string>` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `SourceVersion` | `string` | no |
| `Options` | `Options` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | no |
| `Status` | `string` | no |
| `Version` | `string` | no |

## CreateLicense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseName` | `string` | yes |
| `ProductName` | `string` | yes |
| `ProductSKU` | `string` | yes |
| `Issuer` | `Issuer` | yes |
| `HomeRegion` | `string` | yes |
| `Validity` | `DatetimeRange` | yes |
| `Entitlements` | `List<Entitlement>` | yes |
| `Beneficiary` | `string` | yes |
| `ConsumptionConfiguration` | `ConsumptionConfiguration` | yes |
| `LicenseMetadata` | `List<Metadata>` | no |
| `ClientToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | no |
| `Status` | `string` | no |
| `Version` | `string` | no |

## CreateLicenseAssetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `LicenseAssetGroupConfigurations` | `List<LicenseAssetGroupConfiguration>` | yes |
| `AssociatedLicenseAssetRulesetARNs` | `List<string>` | yes |
| `Properties` | `List<LicenseAssetGroupProperty>` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetGroupArn` | `string` | yes |
| `Status` | `string` | yes |

## CreateLicenseAssetRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Rules` | `List<LicenseAssetRule>` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetRulesetArn` | `string` | yes |

## CreateLicenseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `LicenseCountingType` | `string` | yes |
| `LicenseCount` | `long` | no |
| `LicenseCountHardLimit` | `boolean` | no |
| `LicenseRules` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `DisassociateWhenNotFound` | `boolean` | no |
| `ProductInformationList` | `List<ProductInformation>` | no |
| `LicenseExpiry` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArn` | `string` | no |

## CreateLicenseConversionTaskForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `SourceLicenseContext` | `LicenseConversionContext` | yes |
| `DestinationLicenseContext` | `LicenseConversionContext` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConversionTaskId` | `string` | no |

## CreateLicenseManagerReportGenerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportGeneratorName` | `string` | yes |
| `Type` | `List<string>` | yes |
| `ReportContext` | `ReportContext` | yes |
| `ReportFrequency` | `ReportFrequency` | yes |
| `ClientToken` | `string` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseManagerReportGeneratorArn` | `string` | no |

## CreateLicenseVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | yes |
| `LicenseName` | `string` | yes |
| `ProductName` | `string` | yes |
| `Issuer` | `Issuer` | yes |
| `HomeRegion` | `string` | yes |
| `Validity` | `DatetimeRange` | yes |
| `LicenseMetadata` | `List<Metadata>` | no |
| `Entitlements` | `List<Entitlement>` | yes |
| `ConsumptionConfiguration` | `ConsumptionConfiguration` | yes |
| `Status` | `string` | yes |
| `ClientToken` | `string` | yes |
| `SourceVersion` | `string` | no |
| `ResetUsage` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | no |
| `Version` | `string` | no |
| `Status` | `string` | no |

## CreateToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | yes |
| `RoleArns` | `List<string>` | no |
| `ExpirationInDays` | `integer` | no |
| `TokenProperties` | `List<string>` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TokenId` | `string` | no |
| `TokenType` | `string` | no |
| `Token` | `string` | no |

## DeleteGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | yes |
| `StatusReason` | `string` | no |
| `Version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | no |
| `Status` | `string` | no |
| `Version` | `string` | no |

## DeleteLicense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | yes |
| `SourceVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `DeletionDate` | `string` | no |

## DeleteLicenseAssetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## DeleteLicenseAssetRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetRulesetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLicenseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLicenseManagerReportGenerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseManagerReportGeneratorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TokenId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExtendLicenseConsumption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConsumptionToken` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConsumptionToken` | `string` | no |
| `Expiration` | `string` | no |

## GetAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Token` | `string` | yes |
| `TokenProperties` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | no |

## GetGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grant` | `Grant` | no |

## GetLicense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `License` | `License` | no |

## GetLicenseAssetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetGroup` | `LicenseAssetGroup` | yes |

## GetLicenseAssetRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetRulesetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetRuleset` | `LicenseAssetRuleset` | yes |

## GetLicenseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationId` | `string` | no |
| `LicenseConfigurationArn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LicenseCountingType` | `string` | no |
| `LicenseRules` | `List<string>` | no |
| `LicenseCount` | `long` | no |
| `LicenseCountHardLimit` | `boolean` | no |
| `ConsumedLicenses` | `long` | no |
| `Status` | `string` | no |
| `OwnerAccountId` | `string` | no |
| `ConsumedLicenseSummaryList` | `List<ConsumedLicenseSummary>` | no |
| `ManagedResourceSummaryList` | `List<ManagedResourceSummary>` | no |
| `Tags` | `List<Tag>` | no |
| `ProductInformationList` | `List<ProductInformation>` | no |
| `AutomatedDiscoveryInformation` | `AutomatedDiscoveryInformation` | no |
| `DisassociateWhenNotFound` | `boolean` | no |
| `LicenseExpiry` | `long` | no |

## GetLicenseConversionTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConversionTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConversionTaskId` | `string` | no |
| `ResourceArn` | `string` | no |
| `SourceLicenseContext` | `LicenseConversionContext` | no |
| `DestinationLicenseContext` | `LicenseConversionContext` | no |
| `StatusMessage` | `string` | no |
| `Status` | `string` | no |
| `StartTime` | `timestamp` | no |
| `LicenseConversionTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |

## GetLicenseManagerReportGenerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseManagerReportGeneratorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportGenerator` | `ReportGenerator` | no |

## GetLicenseUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseUsage` | `LicenseUsage` | no |

## GetServiceSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3BucketArn` | `string` | no |
| `SnsTopicArn` | `string` | no |
| `OrganizationConfiguration` | `OrganizationConfiguration` | no |
| `EnableCrossAccountsDiscovery` | `boolean` | no |
| `LicenseManagerResourceShareArn` | `string` | no |
| `CrossRegionDiscoveryHomeRegion` | `string` | no |
| `CrossRegionDiscoverySourceRegions` | `List<string>` | no |
| `ServiceStatus` | `ServiceStatus` | no |

## ListAssetsForLicenseAssetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetGroupArn` | `string` | yes |
| `AssetType` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Assets` | `List<Asset>` | no |
| `NextToken` | `string` | no |

## ListAssociationsForLicenseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationAssociations` | `List<LicenseConfigurationAssociation>` | no |
| `NextToken` | `string` | no |

## ListDistributedGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArns` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grants` | `List<Grant>` | no |
| `NextToken` | `string` | no |

## ListFailuresForLicenseConfigurationOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseOperationFailureList` | `List<LicenseOperationFailure>` | no |
| `NextToken` | `string` | no |

## ListLicenseAssetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetGroups` | `List<LicenseAssetGroup>` | no |
| `NextToken` | `string` | no |

## ListLicenseAssetRulesets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `ShowAWSManagedLicenseAssetRulesets` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetRulesets` | `List<LicenseAssetRuleset>` | no |
| `NextToken` | `string` | no |

## ListLicenseConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArns` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurations` | `List<LicenseConfiguration>` | no |
| `NextToken` | `string` | no |

## ListLicenseConfigurationsForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArns` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurations` | `List<LicenseConfiguration>` | no |
| `NextToken` | `string` | no |

## ListLicenseConversionTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConversionTasks` | `List<LicenseConversionTask>` | no |
| `NextToken` | `string` | no |

## ListLicenseManagerReportGenerators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportGenerators` | `List<ReportGenerator>` | no |
| `NextToken` | `string` | no |

## ListLicenseSpecificationsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseSpecifications` | `List<LicenseSpecification>` | no |
| `NextToken` | `string` | no |

## ListLicenseVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Licenses` | `List<License>` | no |
| `NextToken` | `string` | no |

## ListLicenses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArns` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Licenses` | `List<License>` | no |
| `NextToken` | `string` | no |

## ListReceivedGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArns` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grants` | `List<Grant>` | no |
| `NextToken` | `string` | no |

## ListReceivedGrantsForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArn` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grants` | `List<Grant>` | no |
| `NextToken` | `string` | no |

## ListReceivedLicenses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseArns` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Licenses` | `List<GrantedLicense>` | no |
| `NextToken` | `string` | no |

## ListReceivedLicensesForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Licenses` | `List<GrantedLicense>` | no |
| `NextToken` | `string` | no |

## ListResourceInventory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<InventoryFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceInventoryList` | `List<ResourceInventory>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListTokens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TokenIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tokens` | `List<TokenData>` | no |
| `NextToken` | `string` | no |

## ListUsageForLicenseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationUsageList` | `List<LicenseConfigurationUsage>` | no |
| `NextToken` | `string` | no |

## RejectGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantArn` | `string` | no |
| `Status` | `string` | no |
| `Version` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLicenseAssetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LicenseAssetGroupConfigurations` | `List<LicenseAssetGroupConfiguration>` | no |
| `AssociatedLicenseAssetRulesetARNs` | `List<string>` | yes |
| `Properties` | `List<LicenseAssetGroupProperty>` | no |
| `LicenseAssetGroupArn` | `string` | yes |
| `Status` | `string` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetGroupArn` | `string` | yes |
| `Status` | `string` | yes |

## UpdateLicenseAssetRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Rules` | `List<LicenseAssetRule>` | yes |
| `LicenseAssetRulesetArn` | `string` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseAssetRulesetArn` | `string` | yes |

## UpdateLicenseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseConfigurationArn` | `string` | yes |
| `LicenseConfigurationStatus` | `string` | no |
| `LicenseRules` | `List<string>` | no |
| `LicenseCount` | `long` | no |
| `LicenseCountHardLimit` | `boolean` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ProductInformationList` | `List<ProductInformation>` | no |
| `DisassociateWhenNotFound` | `boolean` | no |
| `LicenseExpiry` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLicenseManagerReportGenerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseManagerReportGeneratorArn` | `string` | yes |
| `ReportGeneratorName` | `string` | yes |
| `Type` | `List<string>` | yes |
| `ReportContext` | `ReportContext` | yes |
| `ReportFrequency` | `ReportFrequency` | yes |
| `ClientToken` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLicenseSpecificationsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `AddLicenseSpecifications` | `List<LicenseSpecification>` | no |
| `RemoveLicenseSpecifications` | `List<LicenseSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateServiceSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3BucketArn` | `string` | no |
| `SnsTopicArn` | `string` | no |
| `OrganizationConfiguration` | `OrganizationConfiguration` | no |
| `EnableCrossAccountsDiscovery` | `boolean` | no |
| `EnabledDiscoverySourceRegions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


