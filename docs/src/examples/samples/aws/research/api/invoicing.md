# AWS Invoicing

API version: 2024-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/invoicing/2024-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetInvoiceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Profiles` | `List<InvoiceProfile>` | no |

## CreateInvoiceUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `InvoiceReceiver` | `string` | yes |
| `Description` | `string` | no |
| `TaxInheritanceDisabled` | `boolean` | no |
| `Rule` | `InvoiceUnitRule` | yes |
| `ResourceTags` | `List<ResourceTag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceUnitArn` | `string` | no |

## CreateProcurementPortalPreference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalName` | `string` | yes |
| `BuyerDomain` | `string` | yes |
| `BuyerIdentifier` | `string` | yes |
| `SupplierDomain` | `string` | yes |
| `SupplierIdentifier` | `string` | yes |
| `Selector` | `ProcurementPortalPreferenceSelector` | no |
| `ProcurementPortalSharedSecret` | `string` | no |
| `ProcurementPortalInstanceEndpoint` | `string` | no |
| `TestEnvPreference` | `TestEnvPreferenceInput` | no |
| `EinvoiceDeliveryEnabled` | `boolean` | yes |
| `EinvoiceDeliveryPreference` | `EinvoiceDeliveryPreference` | no |
| `PurchaseOrderRetrievalEnabled` | `boolean` | yes |
| `Contacts` | `List<Contact>` | yes |
| `ResourceTags` | `List<ResourceTag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |

## DeleteInvoiceUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceUnitArn` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceUnitArn` | `string` | no |

## DeleteProcurementPortalPreference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |

## GetInvoicePDF

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoicePDF` | `InvoicePDF` | no |

## GetInvoiceUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceUnitArn` | `string` | yes |
| `AsOf` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceUnitArn` | `string` | no |
| `InvoiceReceiver` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `TaxInheritanceDisabled` | `boolean` | no |
| `Rule` | `InvoiceUnitRule` | no |
| `LastModified` | `timestamp` | no |

## GetProcurementPortalPreference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreference` | `ProcurementPortalPreference` | yes |

## ListInvoiceSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Selector` | `InvoiceSummariesSelector` | yes |
| `Filter` | `InvoiceSummariesFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceSummaries` | `List<InvoiceSummary>` | yes |
| `NextToken` | `string` | no |

## ListInvoiceUnits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `Filters` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AsOf` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceUnits` | `List<InvoiceUnit>` | no |
| `NextToken` | `string` | no |

## ListProcurementPortalPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferences` | `List<ProcurementPortalPreferenceSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTags` | `List<ResourceTag>` | no |

## PutProcurementPortalPreference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |
| `Selector` | `ProcurementPortalPreferenceSelector` | no |
| `ProcurementPortalSharedSecret` | `string` | no |
| `ProcurementPortalInstanceEndpoint` | `string` | no |
| `TestEnvPreference` | `TestEnvPreferenceInput` | no |
| `EinvoiceDeliveryEnabled` | `boolean` | yes |
| `EinvoiceDeliveryPreference` | `EinvoiceDeliveryPreference` | no |
| `PurchaseOrderRetrievalEnabled` | `boolean` | yes |
| `Contacts` | `List<Contact>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |

## SendProcurementPortalValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourceTags` | `List<ResourceTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourceTagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateInvoiceUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceUnitArn` | `string` | yes |
| `Description` | `string` | no |
| `TaxInheritanceDisabled` | `boolean` | no |
| `Rule` | `InvoiceUnitRule` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvoiceUnitArn` | `string` | no |

## UpdateProcurementPortalPreferenceStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |
| `EinvoiceDeliveryPreferenceStatus` | `string` | no |
| `EinvoiceDeliveryPreferenceStatusReason` | `string` | no |
| `PurchaseOrderRetrievalPreferenceStatus` | `string` | no |
| `PurchaseOrderRetrievalPreferenceStatusReason` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |

## VerifyProcurementPortalValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |
| `Code` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcurementPortalPreferenceArn` | `string` | yes |

