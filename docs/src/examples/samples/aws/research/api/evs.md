# Amazon Elastic VMware Service

API version: 2023-07-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/evs/2023-07-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateEipToVlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `vlanName` | `string` | yes |
| `allocationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vlan` | `Vlan` | no |

## CreateEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `connectorId` | `string` | yes |
| `entitlementType` | `string` | yes |
| `vmIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entitlements` | `List<VmEntitlement>` | no |

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentName` | `string` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `serviceAccessSecurityGroups` | `ServiceAccessSecurityGroups` | no |
| `vpcId` | `string` | yes |
| `serviceAccessSubnetId` | `string` | yes |
| `vcfVersion` | `string` | yes |
| `termsAccepted` | `boolean` | yes |
| `initialVlans` | `InitialVlans` | yes |
| `connectivityInfo` | `ConnectivityInfo` | no |
| `licenseInfo` | `List<LicenseInfo>` | no |
| `hosts` | `List<HostInfoForCreate>` | no |
| `vcfHostnames` | `VcfHostnames` | no |
| `siteId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | no |

## CreateEnvironmentConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `type` | `string` | yes |
| `applianceFqdn` | `string` | yes |
| `secretIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connector` | `Connector` | no |

## CreateEnvironmentHost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `host` | `HostInfoForCreate` | yes |
| `esxVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentSummary` | `EnvironmentSummary` | no |
| `host` | `Host` | no |

## DeleteEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `connectorId` | `string` | yes |
| `entitlementType` | `string` | yes |
| `vmIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entitlements` | `List<VmEntitlement>` | no |

## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | no |

## DeleteEnvironmentConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `connectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connector` | `Connector` | no |
| `environmentSummary` | `EnvironmentSummary` | no |

## DeleteEnvironmentHost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `hostName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentSummary` | `EnvironmentSummary` | no |
| `host` | `Host` | no |

## DisassociateEipFromVlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `vlanName` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vlan` | `Vlan` | no |

## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `List<AccountSetting>` | no |

## GetDepotUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `rotate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `depotUrl` | `string` | yes |
| `token` | `string` | yes |

## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | no |

## GetVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vcfVersions` | `List<VcfVersionInfo>` | yes |
| `instanceTypeEsxVersions` | `List<InstanceTypeEsxVersionsInfo>` | yes |

## ListEnvironmentConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `connectors` | `List<Connector>` | no |

## ListEnvironmentHosts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `environmentHosts` | `List<Host>` | no |

## ListEnvironmentVlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `environmentVlans` | `List<Vlan>` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `state` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `environmentSummaries` | `List<EnvironmentSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListVmEntitlements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `environmentId` | `string` | yes |
| `connectorId` | `string` | yes |
| `entitlementType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `entitlements` | `List<VmEntitlement>` | no |

## PutAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `List<AccountSetting>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `List<AccountSetting>` | no |

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


## UpdateEnvironmentConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `environmentId` | `string` | yes |
| `connectorId` | `string` | yes |
| `applianceFqdn` | `string` | no |
| `secretIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connector` | `Connector` | no |

