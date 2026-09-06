# Amazon SageMaker geospatial capabilities

API version: 2020-05-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sagemaker-geospatial/2020-05-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteEarthObservationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVectorEnrichmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExportEarthObservationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ClientToken` | `string` | no |
| `ExecutionRoleArn` | `string` | yes |
| `ExportSourceImages` | `boolean` | no |
| `OutputConfig` | `OutputConfigInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `ExecutionRoleArn` | `string` | yes |
| `ExportSourceImages` | `boolean` | no |
| `ExportStatus` | `string` | yes |
| `OutputConfig` | `OutputConfigInput` | yes |

## ExportVectorEnrichmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ClientToken` | `string` | no |
| `ExecutionRoleArn` | `string` | yes |
| `OutputConfig` | `ExportVectorEnrichmentJobOutputConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `ExecutionRoleArn` | `string` | yes |
| `ExportStatus` | `string` | yes |
| `OutputConfig` | `ExportVectorEnrichmentJobOutputConfig` | yes |

## GetEarthObservationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `DurationInSeconds` | `integer` | yes |
| `ErrorDetails` | `EarthObservationJobErrorDetails` | no |
| `ExecutionRoleArn` | `string` | no |
| `ExportErrorDetails` | `ExportErrorDetails` | no |
| `ExportStatus` | `string` | no |
| `InputConfig` | `InputConfigOutput` | yes |
| `JobConfig` | `JobConfigInput` | yes |
| `KmsKeyId` | `string` | no |
| `Name` | `string` | yes |
| `OutputBands` | `List<OutputBand>` | no |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |

## GetRasterDataCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Description` | `string` | yes |
| `DescriptionPageUrl` | `string` | yes |
| `ImageSourceBands` | `List<string>` | yes |
| `Name` | `string` | yes |
| `SupportedFilters` | `List<Filter>` | yes |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | yes |

## GetTile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ExecutionRoleArn` | `string` | no |
| `ImageAssets` | `List<string>` | yes |
| `ImageMask` | `boolean` | no |
| `OutputDataType` | `string` | no |
| `OutputFormat` | `string` | no |
| `PropertyFilters` | `string` | no |
| `Target` | `string` | yes |
| `TimeRangeFilter` | `string` | no |
| `x` | `integer` | yes |
| `y` | `integer` | yes |
| `z` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BinaryFile` | `blob` | no |

## GetVectorEnrichmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `DurationInSeconds` | `integer` | yes |
| `ErrorDetails` | `VectorEnrichmentJobErrorDetails` | no |
| `ExecutionRoleArn` | `string` | yes |
| `ExportErrorDetails` | `VectorEnrichmentJobExportErrorDetails` | no |
| `ExportStatus` | `string` | no |
| `InputConfig` | `VectorEnrichmentJobInputConfig` | yes |
| `JobConfig` | `VectorEnrichmentJobConfig` | yes |
| `KmsKeyId` | `string` | no |
| `Name` | `string` | yes |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | yes |

## ListEarthObservationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `StatusEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EarthObservationJobSummaries` | `List<ListEarthObservationJobOutputConfig>` | yes |
| `NextToken` | `string` | no |

## ListRasterDataCollections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RasterDataCollectionSummaries` | `List<RasterDataCollectionMetadata>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListVectorEnrichmentJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `StatusEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `VectorEnrichmentJobSummaries` | `List<ListVectorEnrichmentJobOutputConfig>` | yes |

## SearchRasterDataCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `NextToken` | `string` | no |
| `RasterDataCollectionQuery` | `RasterDataCollectionQueryWithBandFilterInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApproximateResultCount` | `integer` | yes |
| `Items` | `List<ItemSource>` | no |
| `NextToken` | `string` | no |

## StartEarthObservationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ExecutionRoleArn` | `string` | yes |
| `InputConfig` | `InputConfigInput` | yes |
| `JobConfig` | `JobConfigInput` | yes |
| `KmsKeyId` | `string` | no |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `DurationInSeconds` | `integer` | yes |
| `ExecutionRoleArn` | `string` | yes |
| `InputConfig` | `InputConfigOutput` | no |
| `JobConfig` | `JobConfigInput` | yes |
| `KmsKeyId` | `string` | no |
| `Name` | `string` | yes |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |

## StartVectorEnrichmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ExecutionRoleArn` | `string` | yes |
| `InputConfig` | `VectorEnrichmentJobInputConfig` | yes |
| `JobConfig` | `VectorEnrichmentJobConfig` | yes |
| `KmsKeyId` | `string` | no |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `DurationInSeconds` | `integer` | yes |
| `ExecutionRoleArn` | `string` | yes |
| `InputConfig` | `VectorEnrichmentJobInputConfig` | yes |
| `JobConfig` | `VectorEnrichmentJobConfig` | yes |
| `KmsKeyId` | `string` | no |
| `Name` | `string` | yes |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | yes |

## StopEarthObservationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopVectorEnrichmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


