# AWS Artifact

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/artifact/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateComplianceInquiry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `inquiryContent` | `InquiryContent` | yes |
| `clientToken` | `string` | no |
| `supportMode` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `complianceInquirySummary` | `InquirySummary` | no |
| `tags` | `Map<string>` | no |

## ExportComplianceInquiry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `complianceInquiryId` | `string` | yes |
| `queryIdentifiers` | `List<integer>` | no |
| `includeCitations` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentPresignedUrl` | `string` | no |
| `tags` | `Map<string>` | no |

## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountSettings` | `AccountSettings` | no |

## GetComplianceInquiryMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `complianceInquiryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `complianceInquiryDetail` | `InquiryDetail` | no |
| `tags` | `Map<string>` | no |

## GetReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |
| `reportVersion` | `long` | no |
| `termToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentPresignedUrl` | `string` | no |

## GetReportMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |
| `reportVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportDetails` | `ReportDetail` | no |

## GetTermForReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |
| `reportVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentPresignedUrl` | `string` | no |
| `termToken` | `string` | no |

## ListComplianceInquiries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `complianceInquiries` | `List<InquirySummary>` | no |
| `nextToken` | `string` | no |

## ListComplianceInquiryQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `complianceInquiryId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queries` | `List<QuerySummary>` | no |
| `nextToken` | `string` | no |

## ListCustomerAgreements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customerAgreements` | `List<CustomerAgreementSummary>` | yes |
| `nextToken` | `string` | no |

## ListReportVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reports` | `List<ReportSummary>` | yes |
| `nextToken` | `string` | no |

## ListReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reports` | `List<ReportSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationSubscriptionStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountSettings` | `AccountSettings` | no |

## PutComplianceInquiryFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `complianceInquiryId` | `string` | yes |
| `queryIdentifier` | `integer` | no |
| `rating` | `string` | yes |
| `responseRevisionId` | `integer` | no |
| `reasonCodes` | `List<string>` | no |
| `comment` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `submittedAt` | `timestamp` | yes |

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


