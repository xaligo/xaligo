# Amazon CloudHSM

API version: 2014-05-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudhsm/2014-05-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTagsToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagList` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## CreateHapg

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Label` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HapgArn` | `string` | no |

## CreateHsm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetId` | `string` | yes |
| `SshKey` | `string` | yes |
| `EniIp` | `string` | no |
| `IamRoleArn` | `string` | yes |
| `ExternalId` | `string` | no |
| `SubscriptionType` | `string` | yes |
| `ClientToken` | `string` | no |
| `SyslogIp` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmArn` | `string` | no |

## CreateLunaClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Label` | `string` | no |
| `Certificate` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientArn` | `string` | no |

## DeleteHapg

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HapgArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## DeleteHsm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## DeleteLunaClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## DescribeHapg

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HapgArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HapgArn` | `string` | no |
| `HapgSerial` | `string` | no |
| `HsmsLastActionFailed` | `List<string>` | no |
| `HsmsPendingDeletion` | `List<string>` | no |
| `HsmsPendingRegistration` | `List<string>` | no |
| `Label` | `string` | no |
| `LastModifiedTimestamp` | `string` | no |
| `PartitionSerialList` | `List<string>` | no |
| `State` | `string` | no |

## DescribeHsm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmArn` | `string` | no |
| `HsmSerialNumber` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmArn` | `string` | no |
| `Status` | `string` | no |
| `StatusDetails` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `EniId` | `string` | no |
| `EniIp` | `string` | no |
| `SubscriptionType` | `string` | no |
| `SubscriptionStartDate` | `string` | no |
| `SubscriptionEndDate` | `string` | no |
| `VpcId` | `string` | no |
| `SubnetId` | `string` | no |
| `IamRoleArn` | `string` | no |
| `SerialNumber` | `string` | no |
| `VendorName` | `string` | no |
| `HsmType` | `string` | no |
| `SoftwareVersion` | `string` | no |
| `SshPublicKey` | `string` | no |
| `SshKeyLastUpdated` | `string` | no |
| `ServerCertUri` | `string` | no |
| `ServerCertLastUpdated` | `string` | no |
| `Partitions` | `List<string>` | no |

## DescribeLunaClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientArn` | `string` | no |
| `CertificateFingerprint` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientArn` | `string` | no |
| `Certificate` | `string` | no |
| `CertificateFingerprint` | `string` | no |
| `LastModifiedTimestamp` | `string` | no |
| `Label` | `string` | no |

## GetConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientArn` | `string` | yes |
| `ClientVersion` | `string` | yes |
| `HapgList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigType` | `string` | no |
| `ConfigFile` | `string` | no |
| `ConfigCred` | `string` | no |

## ListAvailableZones

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AZList` | `List<string>` | no |

## ListHapgs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HapgList` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListHsms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmList` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListLunaClients

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientList` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | yes |

## ModifyHapg

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HapgArn` | `string` | yes |
| `Label` | `string` | no |
| `PartitionSerialList` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HapgArn` | `string` | no |

## ModifyHsm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmArn` | `string` | yes |
| `SubnetId` | `string` | no |
| `EniIp` | `string` | no |
| `IamRoleArn` | `string` | no |
| `ExternalId` | `string` | no |
| `SyslogIp` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmArn` | `string` | no |

## ModifyLunaClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientArn` | `string` | yes |
| `Certificate` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientArn` | `string` | no |

## RemoveTagsFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeyList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

