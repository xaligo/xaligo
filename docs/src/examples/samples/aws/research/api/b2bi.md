# AWS B2B Data Interchange

API version: 2022-06-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/b2bi/2022-06-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `configuration` | `CapabilityConfiguration` | yes |
| `instructionsDocuments` | `List<S3Location>` | no |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityId` | `string` | yes |
| `capabilityArn` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `configuration` | `CapabilityConfiguration` | yes |
| `instructionsDocuments` | `List<S3Location>` | no |
| `createdAt` | `timestamp` | yes |

## CreatePartnership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `name` | `string` | yes |
| `email` | `string` | yes |
| `phone` | `string` | no |
| `capabilities` | `List<string>` | yes |
| `capabilityOptions` | `CapabilityOptions` | no |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `partnershipId` | `string` | yes |
| `partnershipArn` | `string` | yes |
| `name` | `string` | no |
| `email` | `string` | no |
| `phone` | `string` | no |
| `capabilities` | `List<string>` | no |
| `capabilityOptions` | `CapabilityOptions` | no |
| `tradingPartnerId` | `string` | no |
| `createdAt` | `timestamp` | yes |

## CreateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `email` | `string` | no |
| `phone` | `string` | yes |
| `businessName` | `string` | yes |
| `logging` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `profileArn` | `string` | yes |
| `name` | `string` | yes |
| `businessName` | `string` | yes |
| `phone` | `string` | yes |
| `email` | `string` | no |
| `logging` | `string` | no |
| `logGroupName` | `string` | no |
| `createdAt` | `timestamp` | yes |

## CreateStarterMappingTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `outputSampleLocation` | `S3Location` | no |
| `mappingType` | `string` | yes |
| `templateDetails` | `TemplateDetails` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mappingTemplate` | `string` | yes |

## CreateTransformer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |
| `fileFormat` | `string` | no |
| `mappingTemplate` | `string` | no |
| `ediType` | `EdiType` | no |
| `sampleDocument` | `string` | no |
| `inputConversion` | `InputConversion` | no |
| `mapping` | `Mapping` | no |
| `outputConversion` | `OutputConversion` | no |
| `sampleDocuments` | `SampleDocuments` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerId` | `string` | yes |
| `transformerArn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `fileFormat` | `string` | no |
| `mappingTemplate` | `string` | no |
| `ediType` | `EdiType` | no |
| `sampleDocument` | `string` | no |
| `inputConversion` | `InputConversion` | no |
| `mapping` | `Mapping` | no |
| `outputConversion` | `OutputConversion` | no |
| `sampleDocuments` | `SampleDocuments` | no |

## DeleteCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePartnership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `partnershipId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTransformer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GenerateMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inputFileContent` | `string` | yes |
| `outputFileContent` | `string` | yes |
| `mappingType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mappingTemplate` | `string` | yes |
| `mappingAccuracy` | `float` | no |

## GetCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityId` | `string` | yes |
| `capabilityArn` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `configuration` | `CapabilityConfiguration` | yes |
| `instructionsDocuments` | `List<S3Location>` | no |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | no |

## GetPartnership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `partnershipId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `partnershipId` | `string` | yes |
| `partnershipArn` | `string` | yes |
| `name` | `string` | no |
| `email` | `string` | no |
| `phone` | `string` | no |
| `capabilities` | `List<string>` | no |
| `capabilityOptions` | `CapabilityOptions` | no |
| `tradingPartnerId` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | no |

## GetProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `profileArn` | `string` | yes |
| `name` | `string` | yes |
| `email` | `string` | no |
| `phone` | `string` | yes |
| `businessName` | `string` | yes |
| `logging` | `string` | no |
| `logGroupName` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | no |

## GetTransformer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerId` | `string` | yes |
| `transformerArn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | no |
| `fileFormat` | `string` | no |
| `mappingTemplate` | `string` | no |
| `ediType` | `EdiType` | no |
| `sampleDocument` | `string` | no |
| `inputConversion` | `InputConversion` | no |
| `mapping` | `Mapping` | no |
| `outputConversion` | `OutputConversion` | no |
| `sampleDocuments` | `SampleDocuments` | no |

## GetTransformerJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerJobId` | `string` | yes |
| `transformerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `outputFiles` | `List<S3Location>` | no |
| `message` | `string` | no |

## ListCapabilities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilities` | `List<CapabilitySummary>` | yes |
| `nextToken` | `string` | no |

## ListPartnerships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `partnerships` | `List<PartnershipSummary>` | yes |
| `nextToken` | `string` | no |

## ListProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profiles` | `List<ProfileSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListTransformers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformers` | `List<TransformerSummary>` | yes |
| `nextToken` | `string` | no |

## StartTransformerJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inputFile` | `S3Location` | yes |
| `outputLocation` | `S3Location` | yes |
| `transformerId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerJobId` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestConversion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `source` | `ConversionSource` | yes |
| `target` | `ConversionTarget` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `convertedFileContent` | `string` | yes |
| `validationMessages` | `List<string>` | no |

## TestMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inputFileContent` | `string` | yes |
| `mappingTemplate` | `string` | yes |
| `fileFormat` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mappedFileContent` | `string` | yes |

## TestParsing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inputFile` | `S3Location` | yes |
| `fileFormat` | `string` | yes |
| `ediType` | `EdiType` | yes |
| `advancedOptions` | `AdvancedOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `parsedFileContent` | `string` | yes |
| `parsedSplitFileContents` | `List<string>` | no |
| `validationMessages` | `List<string>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityId` | `string` | yes |
| `name` | `string` | no |
| `configuration` | `CapabilityConfiguration` | no |
| `instructionsDocuments` | `List<S3Location>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityId` | `string` | yes |
| `capabilityArn` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `configuration` | `CapabilityConfiguration` | yes |
| `instructionsDocuments` | `List<S3Location>` | no |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | no |

## UpdatePartnership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `partnershipId` | `string` | yes |
| `name` | `string` | no |
| `capabilities` | `List<string>` | no |
| `capabilityOptions` | `CapabilityOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `partnershipId` | `string` | yes |
| `partnershipArn` | `string` | yes |
| `name` | `string` | no |
| `email` | `string` | no |
| `phone` | `string` | no |
| `capabilities` | `List<string>` | no |
| `capabilityOptions` | `CapabilityOptions` | no |
| `tradingPartnerId` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | no |

## UpdateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `name` | `string` | no |
| `email` | `string` | no |
| `phone` | `string` | no |
| `businessName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `profileArn` | `string` | yes |
| `name` | `string` | yes |
| `email` | `string` | no |
| `phone` | `string` | yes |
| `businessName` | `string` | yes |
| `logging` | `string` | no |
| `logGroupName` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | no |

## UpdateTransformer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerId` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | no |
| `fileFormat` | `string` | no |
| `mappingTemplate` | `string` | no |
| `ediType` | `EdiType` | no |
| `sampleDocument` | `string` | no |
| `inputConversion` | `InputConversion` | no |
| `mapping` | `Mapping` | no |
| `outputConversion` | `OutputConversion` | no |
| `sampleDocuments` | `SampleDocuments` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerId` | `string` | yes |
| `transformerArn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | yes |
| `fileFormat` | `string` | no |
| `mappingTemplate` | `string` | no |
| `ediType` | `EdiType` | no |
| `sampleDocument` | `string` | no |
| `inputConversion` | `InputConversion` | no |
| `mapping` | `Mapping` | no |
| `outputConversion` | `OutputConversion` | no |
| `sampleDocuments` | `SampleDocuments` | no |

