# Amazon Import/Export Snowball

API version: 2016-06-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/snowball/2016-06-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Address` | `Address` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressId` | `string` | no |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobType` | `string` | yes |
| `Resources` | `JobResource` | no |
| `OnDeviceServiceConfiguration` | `OnDeviceServiceConfiguration` | no |
| `Description` | `string` | no |
| `AddressId` | `string` | yes |
| `KmsKeyARN` | `string` | no |
| `RoleARN` | `string` | no |
| `SnowballType` | `string` | yes |
| `ShippingOption` | `string` | yes |
| `Notification` | `Notification` | no |
| `ForwardingAddressId` | `string` | no |
| `TaxDocuments` | `TaxDocuments` | no |
| `RemoteManagement` | `string` | no |
| `InitialClusterSize` | `integer` | no |
| `ForceCreateJobs` | `boolean` | no |
| `LongTermPricingIds` | `List<string>` | no |
| `SnowballCapacityPreference` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | no |
| `JobListEntries` | `List<JobListEntry>` | no |

## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobType` | `string` | no |
| `Resources` | `JobResource` | no |
| `OnDeviceServiceConfiguration` | `OnDeviceServiceConfiguration` | no |
| `Description` | `string` | no |
| `AddressId` | `string` | no |
| `KmsKeyARN` | `string` | no |
| `RoleARN` | `string` | no |
| `SnowballCapacityPreference` | `string` | no |
| `ShippingOption` | `string` | no |
| `Notification` | `Notification` | no |
| `ClusterId` | `string` | no |
| `SnowballType` | `string` | no |
| `ForwardingAddressId` | `string` | no |
| `TaxDocuments` | `TaxDocuments` | no |
| `DeviceConfiguration` | `DeviceConfiguration` | no |
| `RemoteManagement` | `string` | no |
| `LongTermPricingId` | `string` | no |
| `ImpactLevel` | `string` | no |
| `PickupDetails` | `PickupDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## CreateLongTermPricing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LongTermPricingType` | `string` | yes |
| `IsLongTermPricingAutoRenew` | `boolean` | no |
| `SnowballType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LongTermPricingId` | `string` | no |

## CreateReturnShippingLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `ShippingOption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## DescribeAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Address` | `Address` | no |

## DescribeAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Addresses` | `List<Address>` | no |
| `NextToken` | `string` | no |

## DescribeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterMetadata` | `ClusterMetadata` | no |

## DescribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobMetadata` | `JobMetadata` | no |
| `SubJobMetadata` | `List<JobMetadata>` | no |

## DescribeReturnShippingLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `ExpirationDate` | `timestamp` | no |
| `ReturnShippingLabelURI` | `string` | no |

## GetJobManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManifestURI` | `string` | no |

## GetJobUnlockCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnlockCode` | `string` | no |

## GetSnowballUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnowballLimit` | `integer` | no |
| `SnowballsInUse` | `integer` | no |

## GetSoftwareUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdatesURI` | `string` | no |

## ListClusterJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobListEntries` | `List<JobListEntry>` | no |
| `NextToken` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterListEntries` | `List<ClusterListEntry>` | no |
| `NextToken` | `string` | no |

## ListCompatibleImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompatibleImages` | `List<CompatibleImage>` | no |
| `NextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobListEntries` | `List<JobListEntry>` | no |
| `NextToken` | `string` | no |

## ListLongTermPricing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LongTermPricingEntries` | `List<LongTermPricingListEntry>` | no |
| `NextToken` | `string` | no |

## ListPickupLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Addresses` | `List<Address>` | no |
| `NextToken` | `string` | no |

## ListServiceVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceName` | `string` | yes |
| `DependentServices` | `List<DependentService>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceVersions` | `List<ServiceVersion>` | yes |
| `ServiceName` | `string` | yes |
| `DependentServices` | `List<DependentService>` | no |
| `NextToken` | `string` | no |

## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `RoleARN` | `string` | no |
| `Description` | `string` | no |
| `Resources` | `JobResource` | no |
| `OnDeviceServiceConfiguration` | `OnDeviceServiceConfiguration` | no |
| `AddressId` | `string` | no |
| `ShippingOption` | `string` | no |
| `Notification` | `Notification` | no |
| `ForwardingAddressId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `RoleARN` | `string` | no |
| `Notification` | `Notification` | no |
| `Resources` | `JobResource` | no |
| `OnDeviceServiceConfiguration` | `OnDeviceServiceConfiguration` | no |
| `AddressId` | `string` | no |
| `ShippingOption` | `string` | no |
| `Description` | `string` | no |
| `SnowballCapacityPreference` | `string` | no |
| `ForwardingAddressId` | `string` | no |
| `PickupDetails` | `PickupDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateJobShipmentState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `ShipmentState` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLongTermPricing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LongTermPricingId` | `string` | yes |
| `ReplacementJob` | `string` | no |
| `IsLongTermPricingAutoRenew` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


