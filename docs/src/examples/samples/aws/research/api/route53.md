# Amazon Route 53

API version: 2013-04-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/route53/2013-04-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivateKeySigningKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## AssociateVPCWithHostedZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `VPC` | `VPC` | yes |
| `Comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## ChangeCidrCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `CollectionVersion` | `long` | no |
| `Changes` | `List<CidrCollectionChange>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

## ChangeResourceRecordSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `ChangeBatch` | `ChangeBatch` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## ChangeTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `ResourceId` | `string` | yes |
| `AddTags` | `List<Tag>` | no |
| `RemoveTagKeys` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateCidrCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CallerReference` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Collection` | `CidrCollection` | no |
| `Location` | `string` | no |

## CreateHealthCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallerReference` | `string` | yes |
| `HealthCheckConfig` | `HealthCheckConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheck` | `HealthCheck` | yes |
| `Location` | `string` | yes |

## CreateHostedZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `VPC` | `VPC` | no |
| `CallerReference` | `string` | yes |
| `HostedZoneConfig` | `HostedZoneConfig` | no |
| `DelegationSetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZone` | `HostedZone` | yes |
| `ChangeInfo` | `ChangeInfo` | yes |
| `DelegationSet` | `DelegationSet` | yes |
| `VPC` | `VPC` | no |
| `Location` | `string` | yes |

## CreateKeySigningKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallerReference` | `string` | yes |
| `HostedZoneId` | `string` | yes |
| `KeyManagementServiceArn` | `string` | yes |
| `Name` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |
| `KeySigningKey` | `KeySigningKey` | yes |
| `Location` | `string` | yes |

## CreateQueryLoggingConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `CloudWatchLogsLogGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryLoggingConfig` | `QueryLoggingConfig` | yes |
| `Location` | `string` | yes |

## CreateReusableDelegationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallerReference` | `string` | yes |
| `HostedZoneId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationSet` | `DelegationSet` | yes |
| `Location` | `string` | yes |

## CreateTrafficPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Document` | `string` | yes |
| `Comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicy` | `TrafficPolicy` | yes |
| `Location` | `string` | yes |

## CreateTrafficPolicyInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `Name` | `string` | yes |
| `TTL` | `long` | yes |
| `TrafficPolicyId` | `string` | yes |
| `TrafficPolicyVersion` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyInstance` | `TrafficPolicyInstance` | yes |
| `Location` | `string` | yes |

## CreateTrafficPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Document` | `string` | yes |
| `Comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicy` | `TrafficPolicy` | yes |
| `Location` | `string` | yes |

## CreateVPCAssociationAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `VPC` | `VPC` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `VPC` | `VPC` | yes |

## DeactivateKeySigningKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## DeleteCidrCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHealthCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHostedZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## DeleteKeySigningKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## DeleteQueryLoggingConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReusableDelegationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrafficPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Version` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrafficPolicyInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVPCAssociationAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `VPC` | `VPC` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableHostedZoneDNSSEC

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## DisassociateVPCFromHostedZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `VPC` | `VPC` | yes |
| `Comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## EnableHostedZoneDNSSEC

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## GetAccountLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `AccountLimit` | yes |
| `Count` | `long` | yes |

## GetChange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeInfo` | `ChangeInfo` | yes |

## GetCheckerIpRanges

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CheckerIpRanges` | `List<string>` | yes |

## GetDNSSEC

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `DNSSECStatus` | yes |
| `KeySigningKeys` | `List<KeySigningKey>` | yes |

## GetGeoLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinentCode` | `string` | no |
| `CountryCode` | `string` | no |
| `SubdivisionCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoLocationDetails` | `GeoLocationDetails` | yes |

## GetHealthCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheck` | `HealthCheck` | yes |

## GetHealthCheckCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckCount` | `long` | yes |

## GetHealthCheckLastFailureReason

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckObservations` | `List<HealthCheckObservation>` | yes |

## GetHealthCheckStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckObservations` | `List<HealthCheckObservation>` | yes |

## GetHostedZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZone` | `HostedZone` | yes |
| `DelegationSet` | `DelegationSet` | no |
| `VPCs` | `List<VPC>` | no |

## GetHostedZoneCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneCount` | `long` | yes |

## GetHostedZoneLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `HostedZoneId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `HostedZoneLimit` | yes |
| `Count` | `long` | yes |

## GetQueryLoggingConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryLoggingConfig` | `QueryLoggingConfig` | yes |

## GetReusableDelegationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationSet` | `DelegationSet` | yes |

## GetReusableDelegationSetLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `DelegationSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `ReusableDelegationSetLimit` | yes |
| `Count` | `long` | yes |

## GetTrafficPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Version` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicy` | `TrafficPolicy` | yes |

## GetTrafficPolicyInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyInstance` | `TrafficPolicyInstance` | yes |

## GetTrafficPolicyInstanceCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyInstanceCount` | `integer` | yes |

## ListCidrBlocks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `LocationName` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CidrBlocks` | `List<CidrBlockSummary>` | no |

## ListCidrCollections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CidrCollections` | `List<CollectionSummary>` | no |

## ListCidrLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CidrLocations` | `List<LocationSummary>` | no |

## ListGeoLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartContinentCode` | `string` | no |
| `StartCountryCode` | `string` | no |
| `StartSubdivisionCode` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoLocationDetailsList` | `List<GeoLocationDetails>` | yes |
| `IsTruncated` | `boolean` | yes |
| `NextContinentCode` | `string` | no |
| `NextCountryCode` | `string` | no |
| `NextSubdivisionCode` | `string` | no |
| `MaxItems` | `string` | yes |

## ListHealthChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthChecks` | `List<HealthCheck>` | yes |
| `Marker` | `string` | yes |
| `IsTruncated` | `boolean` | yes |
| `NextMarker` | `string` | no |
| `MaxItems` | `string` | yes |

## ListHostedZones

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `DelegationSetId` | `string` | no |
| `HostedZoneType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZones` | `List<HostedZone>` | yes |
| `Marker` | `string` | yes |
| `IsTruncated` | `boolean` | yes |
| `NextMarker` | `string` | no |
| `MaxItems` | `string` | yes |

## ListHostedZonesByName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DNSName` | `string` | no |
| `HostedZoneId` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZones` | `List<HostedZone>` | yes |
| `DNSName` | `string` | no |
| `HostedZoneId` | `string` | no |
| `IsTruncated` | `boolean` | yes |
| `NextDNSName` | `string` | no |
| `NextHostedZoneId` | `string` | no |
| `MaxItems` | `string` | yes |

## ListHostedZonesByVPC

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VPCId` | `string` | yes |
| `VPCRegion` | `string` | yes |
| `MaxItems` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneSummaries` | `List<HostedZoneSummary>` | yes |
| `MaxItems` | `string` | yes |
| `NextToken` | `string` | no |

## ListQueryLoggingConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryLoggingConfigs` | `List<QueryLoggingConfig>` | yes |
| `NextToken` | `string` | no |

## ListResourceRecordSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `StartRecordName` | `string` | no |
| `StartRecordType` | `string` | no |
| `StartRecordIdentifier` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceRecordSets` | `List<ResourceRecordSet>` | yes |
| `IsTruncated` | `boolean` | yes |
| `NextRecordName` | `string` | no |
| `NextRecordType` | `string` | no |
| `NextRecordIdentifier` | `string` | no |
| `MaxItems` | `string` | yes |

## ListReusableDelegationSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegationSets` | `List<DelegationSet>` | yes |
| `Marker` | `string` | yes |
| `IsTruncated` | `boolean` | yes |
| `NextMarker` | `string` | no |
| `MaxItems` | `string` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTagSet` | `ResourceTagSet` | yes |

## ListTagsForResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `ResourceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTagSets` | `List<ResourceTagSet>` | yes |

## ListTrafficPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyIdMarker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicySummaries` | `List<TrafficPolicySummary>` | yes |
| `IsTruncated` | `boolean` | yes |
| `TrafficPolicyIdMarker` | `string` | yes |
| `MaxItems` | `string` | yes |

## ListTrafficPolicyInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneIdMarker` | `string` | no |
| `TrafficPolicyInstanceNameMarker` | `string` | no |
| `TrafficPolicyInstanceTypeMarker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyInstances` | `List<TrafficPolicyInstance>` | yes |
| `HostedZoneIdMarker` | `string` | no |
| `TrafficPolicyInstanceNameMarker` | `string` | no |
| `TrafficPolicyInstanceTypeMarker` | `string` | no |
| `IsTruncated` | `boolean` | yes |
| `MaxItems` | `string` | yes |

## ListTrafficPolicyInstancesByHostedZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `TrafficPolicyInstanceNameMarker` | `string` | no |
| `TrafficPolicyInstanceTypeMarker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyInstances` | `List<TrafficPolicyInstance>` | yes |
| `TrafficPolicyInstanceNameMarker` | `string` | no |
| `TrafficPolicyInstanceTypeMarker` | `string` | no |
| `IsTruncated` | `boolean` | yes |
| `MaxItems` | `string` | yes |

## ListTrafficPolicyInstancesByPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyId` | `string` | yes |
| `TrafficPolicyVersion` | `integer` | yes |
| `HostedZoneIdMarker` | `string` | no |
| `TrafficPolicyInstanceNameMarker` | `string` | no |
| `TrafficPolicyInstanceTypeMarker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyInstances` | `List<TrafficPolicyInstance>` | yes |
| `HostedZoneIdMarker` | `string` | no |
| `TrafficPolicyInstanceNameMarker` | `string` | no |
| `TrafficPolicyInstanceTypeMarker` | `string` | no |
| `IsTruncated` | `boolean` | yes |
| `MaxItems` | `string` | yes |

## ListTrafficPolicyVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `TrafficPolicyVersionMarker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicies` | `List<TrafficPolicy>` | yes |
| `IsTruncated` | `boolean` | yes |
| `TrafficPolicyVersionMarker` | `string` | yes |
| `MaxItems` | `string` | yes |

## ListVPCAssociationAuthorizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `NextToken` | `string` | no |
| `VPCs` | `List<VPC>` | yes |

## TestDNSAnswer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `RecordName` | `string` | yes |
| `RecordType` | `string` | yes |
| `ResolverIP` | `string` | no |
| `EDNS0ClientSubnetIP` | `string` | no |
| `EDNS0ClientSubnetMask` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Nameserver` | `string` | yes |
| `RecordName` | `string` | yes |
| `RecordType` | `string` | yes |
| `RecordData` | `List<string>` | yes |
| `ResponseCode` | `string` | yes |
| `Protocol` | `string` | yes |

## UpdateHealthCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckId` | `string` | yes |
| `HealthCheckVersion` | `long` | no |
| `IPAddress` | `string` | no |
| `Port` | `integer` | no |
| `ResourcePath` | `string` | no |
| `FullyQualifiedDomainName` | `string` | no |
| `SearchString` | `string` | no |
| `FailureThreshold` | `integer` | no |
| `Inverted` | `boolean` | no |
| `Disabled` | `boolean` | no |
| `HealthThreshold` | `integer` | no |
| `ChildHealthChecks` | `List<string>` | no |
| `EnableSNI` | `boolean` | no |
| `Regions` | `List<string>` | no |
| `AlarmIdentifier` | `AlarmIdentifier` | no |
| `InsufficientDataHealthStatus` | `string` | no |
| `ResetElements` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheck` | `HealthCheck` | yes |

## UpdateHostedZoneComment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZone` | `HostedZone` | yes |

## UpdateHostedZoneFeatures

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostedZoneId` | `string` | yes |
| `EnableAcceleratedRecovery` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTrafficPolicyComment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Version` | `integer` | yes |
| `Comment` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicy` | `TrafficPolicy` | yes |

## UpdateTrafficPolicyInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `TTL` | `long` | yes |
| `TrafficPolicyId` | `string` | yes |
| `TrafficPolicyVersion` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficPolicyInstance` | `TrafficPolicyInstance` | yes |

