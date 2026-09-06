# AWS Backup Gateway

API version: 2021-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/backup-gateway/2021-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateGatewayToServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |
| `ServerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## CreateGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivationKey` | `string` | yes |
| `GatewayDisplayName` | `string` | yes |
| `GatewayType` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## DeleteGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## DeleteHypervisor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | no |

## DisassociateGatewayFromServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## GetBandwidthRateLimitSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |
| `BandwidthRateLimitIntervals` | `List<BandwidthRateLimitInterval>` | no |

## GetGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Gateway` | `GatewayDetails` | no |

## GetHypervisor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Hypervisor` | `HypervisorDetails` | no |

## GetHypervisorPropertyMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | no |
| `VmwareToAwsTagMappings` | `List<VmwareToAwsTagMapping>` | no |
| `IamRoleArn` | `string` | no |

## GetVirtualMachine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VirtualMachine` | `VirtualMachineDetails` | no |

## ImportHypervisorConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Host` | `string` | yes |
| `Username` | `string` | no |
| `Password` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | no |

## ListGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Gateways` | `List<Gateway>` | no |
| `NextToken` | `string` | no |

## ListHypervisors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Hypervisors` | `List<Hypervisor>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

## ListVirtualMachines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VirtualMachines` | `List<VirtualMachine>` | no |
| `NextToken` | `string` | no |

## PutBandwidthRateLimitSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |
| `BandwidthRateLimitIntervals` | `List<BandwidthRateLimitInterval>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## PutHypervisorPropertyMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | yes |
| `VmwareToAwsTagMappings` | `List<VmwareToAwsTagMapping>` | yes |
| `IamRoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | no |

## PutMaintenanceStartTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |
| `HourOfDay` | `integer` | yes |
| `MinuteOfHour` | `integer` | yes |
| `DayOfWeek` | `integer` | no |
| `DayOfMonth` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## StartVirtualMachinesMetadataSync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | no |

## TestHypervisorConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |
| `Host` | `string` | yes |
| `Username` | `string` | no |
| `Password` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | no |

## UpdateGatewayInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |
| `GatewayDisplayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## UpdateGatewaySoftwareNow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## UpdateHypervisor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | yes |
| `Host` | `string` | no |
| `Username` | `string` | no |
| `Password` | `string` | no |
| `Name` | `string` | no |
| `LogGroupArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HypervisorArn` | `string` | no |

