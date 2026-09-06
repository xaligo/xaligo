# AWS Outposts

API version: 2019-12-03. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/outposts/2019-12-03/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelCapacityTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityTaskId` | `string` | yes |
| `OutpostIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelOrder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateOrder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | yes |
| `QuoteIdentifier` | `string` | no |
| `QuoteOptionIdentifier` | `string` | no |
| `LineItems` | `List<LineItemRequest>` | no |
| `PaymentOption` | `string` | yes |
| `PaymentTerm` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Order` | `Order` | no |

## CreateOutpost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `SiteId` | `string` | yes |
| `AvailabilityZone` | `string` | no |
| `AvailabilityZoneId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `SupportedHardwareType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Outpost` | `Outpost` | no |

## CreatePrivateConnectivityConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostId` | `string` | yes |
| `VpcInformationList` | `List<VpcInformation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrivateConnectivityConfig` | `PrivateConnectivityConfig` | no |
| `OutpostId` | `string` | no |

## CreateQuote

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | no |
| `CountryCode` | `string` | yes |
| `RequestedCapacities` | `List<QuoteCapacity>` | yes |
| `RequestedConstraints` | `List<QuoteConstraint>` | no |
| `RequestedPaymentOptions` | `List<string>` | no |
| `RequestedPaymentTerms` | `List<string>` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Quote` | `Quote` | no |

## CreateRenewal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PaymentOption` | `string` | yes |
| `PaymentTerm` | `string` | yes |
| `OutpostIdentifier` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PaymentOption` | `string` | no |
| `PaymentTerm` | `string` | no |
| `OutpostId` | `string` | no |
| `UpfrontPrice` | `float` | no |
| `MonthlyRecurringPrice` | `float` | no |
| `Currency` | `string` | no |

## CreateSite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Notes` | `string` | no |
| `Tags` | `Map<string>` | no |
| `OperatingAddress` | `Address` | no |
| `ShippingAddress` | `Address` | no |
| `RackPhysicalProperties` | `RackPhysicalProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Site` | `Site` | no |

## DeleteOutpost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQuote

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuoteIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetCapacityTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityTaskId` | `string` | yes |
| `OutpostIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityTaskId` | `string` | no |
| `OutpostId` | `string` | no |
| `OrderId` | `string` | no |
| `AssetId` | `string` | no |
| `RequestedInstancePools` | `List<InstanceTypeCapacity>` | no |
| `InstancesToExclude` | `InstancesToExclude` | no |
| `DryRun` | `boolean` | no |
| `CapacityTaskStatus` | `string` | no |
| `Failed` | `CapacityTaskFailure` | no |
| `CreationDate` | `timestamp` | no |
| `CompletionDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |
| `TaskActionOnBlockingInstances` | `string` | no |

## GetCatalogItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogItemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogItem` | `CatalogItem` | no |

## GetConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | no |
| `ConnectionDetails` | `ConnectionDetails` | no |

## GetOrder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Order` | `Order` | no |

## GetOutpost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Outpost` | `Outpost` | no |

## GetOutpostBillingInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `OutpostIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Subscriptions` | `List<Subscription>` | no |
| `ContractEndDate` | `string` | no |
| `PaymentTerm` | `string` | no |
| `PaymentOption` | `string` | no |

## GetOutpostInstanceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypes` | `List<InstanceTypeItem>` | no |
| `NextToken` | `string` | no |
| `OutpostId` | `string` | no |
| `OutpostArn` | `string` | no |

## GetOutpostSupportedInstanceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | yes |
| `OrderId` | `string` | no |
| `AssetId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypes` | `List<InstanceTypeItem>` | no |
| `NextToken` | `string` | no |

## GetPrivateConnectivityConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrivateConnectivityConfig` | `PrivateConnectivityConfig` | no |

## GetQuote

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuoteIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Quote` | `Quote` | no |

## GetRenewalPricing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PricingResult` | `string` | no |
| `PricingOptions` | `List<PricingOption>` | no |

## GetSite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Site` | `Site` | no |

## GetSiteAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteId` | `string` | yes |
| `AddressType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteId` | `string` | no |
| `AddressType` | `string` | no |
| `Address` | `Address` | no |

## ListAssetInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | yes |
| `AssetIdFilter` | `List<string>` | no |
| `InstanceTypeFilter` | `List<string>` | no |
| `AccountIdFilter` | `List<string>` | no |
| `AwsServiceFilter` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetInstances` | `List<AssetInstance>` | no |
| `NextToken` | `string` | no |

## ListAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | yes |
| `HostIdFilter` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `StatusFilter` | `List<string>` | no |
| `AssetTypeFilter` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Assets` | `List<AssetInfo>` | no |
| `NextToken` | `string` | no |

## ListBlockingInstancesForCapacityTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | yes |
| `CapacityTaskId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlockingInstances` | `List<BlockingInstance>` | no |
| `NextToken` | `string` | no |

## ListCapacityTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifierFilter` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `CapacityTaskStatusFilter` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityTasks` | `List<CapacityTaskSummary>` | no |
| `NextToken` | `string` | no |

## ListCatalogItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ItemClassFilter` | `List<string>` | no |
| `SupportedStorageFilter` | `List<string>` | no |
| `EC2FamilyFilter` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogItems` | `List<CatalogItem>` | no |
| `NextToken` | `string` | no |

## ListOrderableInstanceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostGenerationFilter` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypes` | `List<DetailedInstanceTypeItem>` | no |
| `NextToken` | `string` | no |

## ListOrders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifierFilter` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Orders` | `List<OrderSummary>` | no |
| `NextToken` | `string` | no |

## ListOutposts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `LifeCycleStatusFilter` | `List<string>` | no |
| `AvailabilityZoneFilter` | `List<string>` | no |
| `AvailabilityZoneIdFilter` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Outposts` | `List<Outpost>` | no |
| `NextToken` | `string` | no |

## ListQuotes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Quotes` | `List<QuoteSummary>` | no |
| `NextToken` | `string` | no |

## ListSites

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `OperatingAddressCountryCodeFilter` | `List<string>` | no |
| `OperatingAddressStateOrRegionFilter` | `List<string>` | no |
| `OperatingAddressCityFilter` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sites` | `List<Site>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## StartCapacityTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | yes |
| `OrderId` | `string` | no |
| `AssetId` | `string` | no |
| `InstancePools` | `List<InstanceTypeCapacity>` | yes |
| `InstancesToExclude` | `InstancesToExclude` | no |
| `DryRun` | `boolean` | no |
| `TaskActionOnBlockingInstances` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityTaskId` | `string` | no |
| `OutpostId` | `string` | no |
| `OrderId` | `string` | no |
| `AssetId` | `string` | no |
| `RequestedInstancePools` | `List<InstanceTypeCapacity>` | no |
| `InstancesToExclude` | `InstancesToExclude` | no |
| `DryRun` | `boolean` | no |
| `CapacityTaskStatus` | `string` | no |
| `Failed` | `CapacityTaskFailure` | no |
| `CreationDate` | `timestamp` | no |
| `CompletionDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |
| `TaskActionOnBlockingInstances` | `string` | no |

## StartConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceSerialNumber` | `string` | no |
| `AssetId` | `string` | yes |
| `ClientPublicKey` | `string` | yes |
| `NetworkInterfaceDeviceIndex` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | no |
| `UnderlayIpAddress` | `string` | no |

## StartOutpostDecommission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostIdentifier` | `string` | yes |
| `ValidateOnly` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `BlockingResourceTypes` | `List<string>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


## UpdateOutpost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `SupportedHardwareType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Outpost` | `Outpost` | no |

## UpdateQuote

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuoteIdentifier` | `string` | yes |
| `OutpostIdentifier` | `string` | no |
| `CountryCode` | `string` | no |
| `RequestedCapacities` | `List<QuoteCapacity>` | no |
| `RequestedConstraints` | `List<QuoteConstraint>` | no |
| `RequestedPaymentOptions` | `List<string>` | no |
| `RequestedPaymentTerms` | `List<string>` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Quote` | `Quote` | no |

## UpdateSite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Notes` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Site` | `Site` | no |

## UpdateSiteAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteId` | `string` | yes |
| `AddressType` | `string` | yes |
| `Address` | `Address` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressType` | `string` | no |
| `Address` | `Address` | no |

## UpdateSiteRackPhysicalProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteId` | `string` | yes |
| `PowerDrawKva` | `string` | no |
| `PowerPhase` | `string` | no |
| `PowerConnector` | `string` | no |
| `PowerFeedDrop` | `string` | no |
| `UplinkGbps` | `string` | no |
| `UplinkCount` | `string` | no |
| `FiberOpticCableType` | `string` | no |
| `OpticalStandard` | `string` | no |
| `MaximumSupportedWeightLbs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Site` | `Site` | no |

