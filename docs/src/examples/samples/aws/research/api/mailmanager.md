# MailManager

API version: 2023-10-17. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mailmanager/2023-10-17/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAddonInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `AddonSubscriptionId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonInstanceId` | `string` | yes |

## CreateAddonSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `AddonName` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonSubscriptionId` | `string` | yes |

## CreateAddressList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `AddressListName` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |

## CreateAddressListImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `AddressListId` | `string` | yes |
| `Name` | `string` | yes |
| `ImportDataFormat` | `ImportDataFormat` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `PreSignedUrl` | `string` | yes |

## CreateArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ArchiveName` | `string` | yes |
| `Retention` | `ArchiveRetention` | no |
| `KmsKeyArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |

## CreateIngressPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `IngressPointName` | `string` | yes |
| `Type` | `string` | yes |
| `RuleSetId` | `string` | yes |
| `TrafficPolicyId` | `string` | yes |
| `IngressPointConfiguration` | `IngressPointConfiguration` | no |
| `NetworkConfiguration` | `NetworkConfiguration` | no |
| `TlsPolicy` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IngressPointId` | `string` | yes |

## CreateRelay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `RelayName` | `string` | yes |
| `ServerName` | `string` | yes |
| `ServerPort` | `integer` | yes |
| `Authentication` | `RelayAuthentication` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RelayId` | `string` | yes |

## CreateRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `RuleSetName` | `string` | yes |
| `Rules` | `List<Rule>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetId` | `string` | yes |

## CreateTrafficPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `TrafficPolicyName` | `string` | yes |
| `PolicyStatements` | `List<PolicyStatement>` | yes |
| `DefaultAction` | `string` | yes |
| `MaxMessageSizeBytes` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyId` | `string` | yes |

## DeleteAddonInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAddonSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonSubscriptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAddressList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIngressPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IngressPointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRelay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RelayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrafficPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterMemberFromAddressList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |
| `Address` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAddonInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonSubscriptionId` | `string` | no |
| `AddonName` | `string` | no |
| `AddonInstanceArn` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## GetAddonSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonSubscriptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonName` | `string` | no |
| `AddonSubscriptionArn` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## GetAddressList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |
| `AddressListArn` | `string` | yes |
| `AddressListName` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `LastUpdatedTimestamp` | `timestamp` | yes |

## GetAddressListImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `Name` | `string` | yes |
| `Status` | `string` | yes |
| `PreSignedUrl` | `string` | yes |
| `ImportedItemsCount` | `integer` | no |
| `FailedItemsCount` | `integer` | no |
| `ImportDataFormat` | `ImportDataFormat` | yes |
| `AddressListId` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `StartTimestamp` | `timestamp` | no |
| `CompletedTimestamp` | `timestamp` | no |
| `Error` | `string` | no |

## GetArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |
| `ArchiveName` | `string` | yes |
| `ArchiveArn` | `string` | yes |
| `ArchiveState` | `string` | yes |
| `Retention` | `ArchiveRetention` | yes |
| `CreatedTimestamp` | `timestamp` | no |
| `LastUpdatedTimestamp` | `timestamp` | no |
| `KmsKeyArn` | `string` | no |

## GetArchiveExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | no |
| `Filters` | `ArchiveFilters` | no |
| `FromTimestamp` | `timestamp` | no |
| `ToTimestamp` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `ExportDestinationConfiguration` | `ExportDestinationConfiguration` | no |
| `Status` | `ExportStatus` | no |

## GetArchiveMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchivedMessageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageDownloadLink` | `string` | no |
| `Metadata` | `Metadata` | no |
| `Envelope` | `Envelope` | no |

## GetArchiveMessageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchivedMessageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `MessageBody` | no |

## GetArchiveSearch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | no |
| `Filters` | `ArchiveFilters` | no |
| `FromTimestamp` | `timestamp` | no |
| `ToTimestamp` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `Status` | `SearchStatus` | no |

## GetArchiveSearchResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rows` | `List<Row>` | no |

## GetIngressPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IngressPointId` | `string` | yes |
| `IncludeTrustStoreContents` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IngressPointId` | `string` | yes |
| `IngressPointName` | `string` | yes |
| `IngressPointArn` | `string` | no |
| `Status` | `string` | no |
| `Type` | `string` | no |
| `ARecord` | `string` | no |
| `RuleSetId` | `string` | no |
| `TrafficPolicyId` | `string` | no |
| `IngressPointAuthConfiguration` | `IngressPointAuthConfiguration` | no |
| `NetworkConfiguration` | `NetworkConfiguration` | no |
| `TlsPolicy` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `LastUpdatedTimestamp` | `timestamp` | no |

## GetMemberOfAddressList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |
| `Address` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Address` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |

## GetRelay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RelayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RelayId` | `string` | yes |
| `RelayArn` | `string` | no |
| `RelayName` | `string` | no |
| `ServerName` | `string` | no |
| `ServerPort` | `integer` | no |
| `Authentication` | `RelayAuthentication` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `LastModifiedTimestamp` | `timestamp` | no |

## GetRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetId` | `string` | yes |
| `RuleSetArn` | `string` | yes |
| `RuleSetName` | `string` | yes |
| `CreatedDate` | `timestamp` | yes |
| `LastModificationDate` | `timestamp` | yes |
| `Rules` | `List<Rule>` | yes |

## GetTrafficPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyName` | `string` | yes |
| `TrafficPolicyId` | `string` | yes |
| `TrafficPolicyArn` | `string` | no |
| `PolicyStatements` | `List<PolicyStatement>` | no |
| `MaxMessageSizeBytes` | `integer` | no |
| `DefaultAction` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `LastUpdatedTimestamp` | `timestamp` | no |

## ListAddonInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonInstances` | `List<AddonInstance>` | no |
| `NextToken` | `string` | no |

## ListAddonSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddonSubscriptions` | `List<AddonSubscription>` | no |
| `NextToken` | `string` | no |

## ListAddressListImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportJobs` | `List<ImportJob>` | yes |
| `NextToken` | `string` | no |

## ListAddressLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressLists` | `List<AddressList>` | yes |
| `NextToken` | `string` | no |

## ListArchiveExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Exports` | `List<ExportSummary>` | no |
| `NextToken` | `string` | no |

## ListArchiveSearches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Searches` | `List<SearchSummary>` | no |
| `NextToken` | `string` | no |

## ListArchives

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Archives` | `List<Archive>` | yes |
| `NextToken` | `string` | no |

## ListIngressPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageSize` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IngressPoints` | `List<IngressPoint>` | no |
| `NextToken` | `string` | no |

## ListMembersOfAddressList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |
| `Filter` | `AddressFilter` | no |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Addresses` | `List<SavedAddress>` | yes |
| `NextToken` | `string` | no |

## ListRelays

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageSize` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Relays` | `List<Relay>` | yes |
| `NextToken` | `string` | no |

## ListRuleSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSets` | `List<RuleSet>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |

## ListTrafficPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageSize` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicies` | `List<TrafficPolicy>` | no |
| `NextToken` | `string` | no |

## RegisterMemberToAddressList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressListId` | `string` | yes |
| `Address` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAddressListImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartArchiveExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |
| `Filters` | `ArchiveFilters` | no |
| `FromTimestamp` | `timestamp` | yes |
| `ToTimestamp` | `timestamp` | yes |
| `MaxResults` | `integer` | no |
| `ExportDestinationConfiguration` | `ExportDestinationConfiguration` | yes |
| `IncludeMetadata` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportId` | `string` | no |

## StartArchiveSearch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |
| `Filters` | `ArchiveFilters` | no |
| `FromTimestamp` | `timestamp` | yes |
| `ToTimestamp` | `timestamp` | yes |
| `MaxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchId` | `string` | no |

## StopAddressListImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopArchiveExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopArchiveSearch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveId` | `string` | yes |
| `ArchiveName` | `string` | no |
| `Retention` | `ArchiveRetention` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateIngressPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IngressPointId` | `string` | yes |
| `IngressPointName` | `string` | no |
| `StatusToUpdate` | `string` | no |
| `RuleSetId` | `string` | no |
| `TrafficPolicyId` | `string` | no |
| `IngressPointConfiguration` | `IngressPointConfiguration` | no |
| `TlsPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRelay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RelayId` | `string` | yes |
| `RelayName` | `string` | no |
| `ServerName` | `string` | no |
| `ServerPort` | `integer` | no |
| `Authentication` | `RelayAuthentication` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetId` | `string` | yes |
| `RuleSetName` | `string` | no |
| `Rules` | `List<Rule>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTrafficPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyId` | `string` | yes |
| `TrafficPolicyName` | `string` | no |
| `PolicyStatements` | `List<PolicyStatement>` | no |
| `DefaultAction` | `string` | no |
| `MaxMessageSizeBytes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


