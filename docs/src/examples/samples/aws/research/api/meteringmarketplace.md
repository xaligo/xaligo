# AWSMarketplace Metering

API version: 2016-01-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/meteringmarketplace/2016-01-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchMeterUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageRecords` | `List<UsageRecord>` | yes |
| `ProductCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<UsageRecordResult>` | no |
| `UnprocessedRecords` | `List<UsageRecord>` | no |

## MeterUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductCode` | `string` | yes |
| `Timestamp` | `timestamp` | yes |
| `UsageDimension` | `string` | yes |
| `UsageQuantity` | `integer` | no |
| `DryRun` | `boolean` | no |
| `UsageAllocations` | `List<UsageAllocation>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeteringRecordId` | `string` | no |

## RegisterUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductCode` | `string` | yes |
| `PublicKeyVersion` | `integer` | yes |
| `Nonce` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKeyRotationTimestamp` | `timestamp` | no |
| `Signature` | `string` | no |

## ResolveCustomer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerIdentifier` | `string` | no |
| `ProductCode` | `string` | no |
| `CustomerAWSAccountId` | `string` | no |
| `LicenseArn` | `string` | no |

