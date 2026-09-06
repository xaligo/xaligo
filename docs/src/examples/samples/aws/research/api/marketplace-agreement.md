# AWS Marketplace Agreement Service

API version: 2020-03-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/marketplace-agreement/2020-03-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptAgreementCancellationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `agreementCancellationRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | no |
| `agreementCancellationRequestId` | `string` | no |
| `status` | `string` | no |
| `reasonCode` | `string` | no |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## AcceptAgreementPaymentRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | yes |
| `agreementId` | `string` | yes |
| `purchaseOrderReference` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | no |
| `agreementId` | `string` | no |
| `status` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `chargeAmount` | `string` | no |
| `currencyCode` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## AcceptAgreementRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementRequestId` | `string` | yes |
| `purchaseOrders` | `List<PurchaseOrder>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | no |

## BatchCreateBillingAdjustmentRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingAdjustmentRequestEntries` | `List<BatchCreateBillingAdjustmentRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BatchCreateBillingAdjustmentItem>` | yes |
| `errors` | `List<BatchCreateBillingAdjustmentError>` | yes |

## CancelAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelAgreementCancellationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `agreementCancellationRequestId` | `string` | yes |
| `cancellationReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementCancellationRequestId` | `string` | no |
| `agreementId` | `string` | no |
| `reasonCode` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## CancelAgreementPaymentRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | yes |
| `agreementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | no |
| `agreementId` | `string` | no |
| `status` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `chargeAmount` | `string` | no |
| `currencyCode` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## CreateAgreementRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `intent` | `string` | yes |
| `requestedTerms` | `List<RequestedTerm>` | yes |
| `sourceAgreementIdentifier` | `string` | no |
| `agreementProposalIdentifier` | `string` | no |
| `taxConfiguration` | `TaxConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementRequestId` | `string` | no |
| `chargeSummary` | `ChargeSummary` | no |

## DescribeAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | no |
| `acceptor` | `Acceptor` | no |
| `proposer` | `Proposer` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `acceptanceTime` | `timestamp` | no |
| `agreementType` | `string` | no |
| `estimatedCharges` | `EstimatedCharges` | no |
| `proposalSummary` | `ProposalSummary` | no |
| `status` | `string` | no |
| `initialAgreementId` | `string` | no |
| `endTimeBehavior` | `EndTimeBehavior` | no |

## GetAgreementCancellationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementCancellationRequestId` | `string` | yes |
| `agreementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementCancellationRequestId` | `string` | no |
| `agreementId` | `string` | no |
| `reasonCode` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## GetAgreementEntitlements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementEntitlements` | `List<AgreementEntitlement>` | no |
| `nextToken` | `string` | no |

## GetAgreementPaymentRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | yes |
| `agreementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | no |
| `agreementId` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `chargeId` | `string` | no |
| `chargeAmount` | `string` | no |
| `currencyCode` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## GetAgreementTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `acceptedTerms` | `List<AcceptedTerm>` | no |
| `nextToken` | `string` | no |

## GetBillingAdjustmentRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `billingAdjustmentRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingAdjustmentRequestId` | `string` | yes |
| `agreementId` | `string` | yes |
| `adjustmentReasonCode` | `string` | yes |
| `description` | `string` | no |
| `originalInvoiceId` | `string` | yes |
| `adjustmentAmount` | `string` | yes |
| `currencyCode` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## ListAgreementCancellationRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `partyType` | `string` | yes |
| `agreementId` | `string` | no |
| `status` | `string` | no |
| `agreementType` | `string` | no |
| `catalog` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<AgreementCancellationRequestSummary>` | no |

## ListAgreementCharges

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | no |
| `agreementId` | `string` | no |
| `agreementType` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Charge>` | no |
| `nextToken` | `string` | no |

## ListAgreementInvoiceLineItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `groupBy` | `string` | yes |
| `invoiceId` | `string` | no |
| `invoiceType` | `string` | no |
| `invoiceBillingPeriod` | `InvoiceBillingPeriod` | no |
| `beforeIssuedTime` | `timestamp` | no |
| `afterIssuedTime` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementInvoiceLineItemGroupSummaries` | `List<AgreementInvoiceLineItemGroupSummary>` | no |
| `nextToken` | `string` | no |

## ListAgreementPaymentRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `partyType` | `string` | yes |
| `agreementType` | `string` | no |
| `catalog` | `string` | no |
| `agreementId` | `string` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<PaymentRequestSummary>` | yes |

## ListBillingAdjustmentRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | no |
| `status` | `string` | no |
| `createdAfter` | `timestamp` | no |
| `createdBefore` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `catalog` | `string` | no |
| `agreementType` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<BillingAdjustmentSummary>` | yes |

## RejectAgreementCancellationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `agreementCancellationRequestId` | `string` | yes |
| `rejectionReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | no |
| `agreementCancellationRequestId` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `reasonCode` | `string` | no |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## RejectAgreementPaymentRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | yes |
| `agreementId` | `string` | yes |
| `rejectionReason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | no |
| `agreementId` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `chargeAmount` | `string` | no |
| `currencyCode` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## SearchAgreements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | no |
| `filters` | `List<Filter>` | no |
| `sort` | `Sort` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementViewSummaries` | `List<AgreementViewSummary>` | no |
| `nextToken` | `string` | no |

## SendAgreementCancellationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `reasonCode` | `string` | yes |
| `clientToken` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | no |
| `agreementCancellationRequestId` | `string` | no |
| `status` | `string` | no |
| `reasonCode` | `string` | no |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## SendAgreementPaymentRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `agreementId` | `string` | yes |
| `termId` | `string` | yes |
| `name` | `string` | yes |
| `chargeAmount` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentRequestId` | `string` | no |
| `agreementId` | `string` | no |
| `status` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `chargeAmount` | `string` | no |
| `currencyCode` | `string` | no |
| `createdAt` | `timestamp` | no |

## UpdatePurchaseOrders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `purchaseOrders` | `List<PurchaseOrder>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


