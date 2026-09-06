# Service Quotas

API version: 2019-06-24. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/service-quotas/2019-06-24/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateServiceQuotaTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateSupportCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceQuotaIncreaseRequestFromTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `QuotaCode` | `string` | yes |
| `AwsRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateServiceQuotaTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAWSDefaultServiceQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `QuotaCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Quota` | `ServiceQuota` | no |

## GetAssociationForServiceQuotaTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceQuotaTemplateAssociationStatus` | `string` | no |

## GetAutoManagementConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptInLevel` | `string` | no |
| `OptInType` | `string` | no |
| `NotificationArn` | `string` | no |
| `OptInStatus` | `string` | no |
| `ExclusionList` | `Map<List<QuotaInfo>>` | no |

## GetQuotaUtilizationReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | no |
| `Status` | `string` | no |
| `GeneratedAt` | `timestamp` | no |
| `TotalCount` | `integer` | no |
| `Quotas` | `List<QuotaUtilizationInfo>` | no |
| `NextToken` | `string` | no |
| `ErrorCode` | `string` | no |
| `ErrorMessage` | `string` | no |

## GetRequestedServiceQuotaChange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestedQuota` | `RequestedServiceQuotaChange` | no |

## GetServiceQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `QuotaCode` | `string` | yes |
| `ContextId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Quota` | `ServiceQuota` | no |

## GetServiceQuotaIncreaseRequestFromTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `QuotaCode` | `string` | yes |
| `AwsRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceQuotaIncreaseRequestInTemplate` | `ServiceQuotaIncreaseRequestInTemplate` | no |

## ListAWSDefaultServiceQuotas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Quotas` | `List<ServiceQuota>` | no |

## ListRequestedServiceQuotaChangeHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | no |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `QuotaRequestedAtLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RequestedQuotas` | `List<RequestedServiceQuotaChange>` | no |

## ListRequestedServiceQuotaChangeHistoryByQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `QuotaCode` | `string` | yes |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `QuotaRequestedAtLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RequestedQuotas` | `List<RequestedServiceQuotaChange>` | no |

## ListServiceQuotaIncreaseRequestsInTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | no |
| `AwsRegion` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceQuotaIncreaseRequestInTemplateList` | `List<ServiceQuotaIncreaseRequestInTemplate>` | no |
| `NextToken` | `string` | no |

## ListServiceQuotas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `QuotaCode` | `string` | no |
| `QuotaAppliedAtLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Quotas` | `List<ServiceQuota>` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Services` | `List<ServiceInfo>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutServiceQuotaIncreaseRequestIntoTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuotaCode` | `string` | yes |
| `ServiceCode` | `string` | yes |
| `AwsRegion` | `string` | yes |
| `DesiredValue` | `double` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceQuotaIncreaseRequestInTemplate` | `ServiceQuotaIncreaseRequestInTemplate` | no |

## RequestServiceQuotaIncrease

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `QuotaCode` | `string` | yes |
| `DesiredValue` | `double` | yes |
| `ContextId` | `string` | no |
| `SupportCaseAllowed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestedQuota` | `RequestedServiceQuotaChange` | no |

## StartAutoManagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptInLevel` | `string` | yes |
| `OptInType` | `string` | yes |
| `NotificationArn` | `string` | no |
| `ExclusionList` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartQuotaUtilizationReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |

## StopAutoManagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


## UpdateAutoManagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptInType` | `string` | no |
| `NotificationArn` | `string` | no |
| `ExclusionList` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


