# AWS End User Messaging Social

API version: 2024-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/socialmessaging/2024-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateWhatsAppBusinessAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signupCallback` | `WhatsAppSignupCallback` | no |
| `setupFinalization` | `WhatsAppSetupFinalization` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signupCallbackResult` | `WhatsAppSignupCallbackResult` | no |
| `statusCode` | `integer` | no |
| `linkedWhatsAppBusinessAccountId` | `string` | no |

## CreateWhatsAppDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |

## CreateWhatsAppFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowName` | `string` | yes |
| `categories` | `List<string>` | yes |
| `flowJson` | `blob` | no |
| `publish` | `boolean` | no |
| `cloneFlowId` | `string` | no |
| `endpointUri` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowId` | `string` | no |
| `validationErrors` | `List<string>` | no |

## CreateWhatsAppMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateDefinition` | `blob` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metaTemplateId` | `string` | no |
| `templateStatus` | `string` | no |
| `category` | `string` | no |

## CreateWhatsAppMessageTemplateFromLibrary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metaLibraryTemplate` | `MetaLibraryTemplate` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metaTemplateId` | `string` | no |
| `templateStatus` | `string` | no |
| `category` | `string` | no |

## CreateWhatsAppMessageTemplateMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sourceS3File` | `S3File` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metaHeaderHandle` | `string` | no |

## DeleteWhatsAppFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWhatsAppMessageMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mediaId` | `string` | yes |
| `originationPhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `success` | `boolean` | no |

## DeleteWhatsAppMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metaTemplateId` | `string` | no |
| `deleteAllLanguages` | `boolean` | no |
| `id` | `string` | yes |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeprecateWhatsAppFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateWhatsAppBusinessAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetLinkedWhatsAppBusinessAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `account` | `LinkedWhatsAppBusinessAccount` | no |

## GetLinkedWhatsAppBusinessAccountPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `phoneNumber` | `WhatsAppPhoneNumberDetail` | no |
| `linkedWhatsAppBusinessAccountId` | `string` | no |

## GetWhatsAppBusinessPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `originationPhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `businessPublicKey` | `string` | no |
| `businessPublicKeySignatureStatus` | `string` | no |

## GetWhatsAppFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowId` | `string` | yes |
| `flowName` | `string` | yes |
| `flowStatus` | `string` | yes |
| `categories` | `List<string>` | no |
| `validationErrors` | `List<string>` | no |
| `jsonVersion` | `string` | no |
| `dataApiVersion` | `string` | no |
| `endpointUri` | `string` | no |
| `preview` | `MetaFlowPreviewInfo` | no |
| `whatsAppBusinessAccount` | `MetaFlowWhatsAppBusinessAccountInfo` | no |
| `application` | `MetaFlowApplicationInfo` | no |
| `healthStatus` | `MetaFlowHealthStatus` | no |

## GetWhatsAppFlowPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowId` | `string` | yes |
| `invalidate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowId` | `string` | yes |
| `preview` | `MetaFlowPreviewInfo` | yes |

## GetWhatsAppMessageMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mediaId` | `string` | yes |
| `originationPhoneNumberId` | `string` | yes |
| `metadataOnly` | `boolean` | no |
| `destinationS3PresignedUrl` | `S3PresignedUrl` | no |
| `destinationS3File` | `S3File` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mimeType` | `string` | no |
| `fileSize` | `long` | no |

## GetWhatsAppMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metaTemplateId` | `string` | no |
| `id` | `string` | yes |
| `templateName` | `string` | no |
| `templateLanguageCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `template` | `string` | no |

## ListLinkedWhatsAppBusinessAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `linkedAccounts` | `List<LinkedWhatsAppBusinessAccountSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |
| `tags` | `List<Tag>` | no |

## ListWhatsAppFlowAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowAssets` | `List<MetaFlowAsset>` | yes |
| `nextToken` | `string` | no |

## ListWhatsAppFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flows` | `List<MetaFlowSummary>` | yes |
| `nextToken` | `string` | no |

## ListWhatsAppMessageTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templates` | `List<TemplateSummary>` | no |
| `nextToken` | `string` | no |

## ListWhatsAppTemplateLibrary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `id` | `string` | yes |
| `filters` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metaLibraryTemplates` | `List<MetaLibraryTemplateDefinition>` | no |
| `nextToken` | `string` | no |

## PostWhatsAppMessageMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `originationPhoneNumberId` | `string` | yes |
| `sourceS3PresignedUrl` | `S3PresignedUrl` | no |
| `sourceS3File` | `S3File` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mediaId` | `string` | no |

## PublishWhatsAppFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutWhatsAppBusinessAccountEventDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `eventDestinations` | `List<WhatsAppBusinessAccountEventDestination>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutWhatsAppBusinessPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `originationPhoneNumberId` | `string` | yes |
| `businessPublicKey` | `string` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendWhatsAppConversionEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `datasetId` | `string` | yes |
| `eventData` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | yes |

## SendWhatsAppMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `originationPhoneNumberId` | `string` | yes |
| `message` | `blob` | yes |
| `metaApiVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |

## UpdateWhatsAppFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowId` | `string` | yes |
| `flowName` | `string` | no |
| `categories` | `List<string>` | no |
| `endpointUri` | `string` | no |
| `metaAppId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWhatsAppFlowAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `flowId` | `string` | yes |
| `flowJson` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `validationErrors` | `List<string>` | no |

## UpdateWhatsAppMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `metaTemplateId` | `string` | no |
| `templateName` | `string` | no |
| `templateLanguageCode` | `string` | no |
| `parameterFormat` | `string` | no |
| `templateCategory` | `string` | no |
| `templateComponents` | `blob` | no |
| `ctaUrlLinkTrackingOptedOut` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


