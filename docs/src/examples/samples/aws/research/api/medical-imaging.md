# AWS Health Imaging

API version: 2023-07-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/medical-imaging/2023-07-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CopyImageSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `sourceImageSetId` | `string` | yes |
| `copyImageSetInformation` | `CopyImageSetInformation` | yes |
| `force` | `boolean` | no |
| `promoteToPrimary` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `sourceImageSetProperties` | `CopySourceImageSetProperties` | yes |
| `destinationImageSetProperties` | `CopyDestinationImageSetProperties` | yes |

## CreateDatastore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreName` | `string` | no |
| `clientToken` | `string` | yes |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |
| `lambdaAuthorizerArn` | `string` | no |
| `losslessStorageFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `datastoreStatus` | `string` | yes |

## DeleteDatastore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `datastoreStatus` | `string` | yes |

## DeleteImageSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |
| `imageSetState` | `string` | yes |
| `imageSetWorkflowStatus` | `string` | yes |

## GetDICOMImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobProperties` | `DICOMImportJobProperties` | yes |

## GetDatastore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreProperties` | `DatastoreProperties` | yes |

## GetImageFrame

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |
| `imageFrameInformation` | `ImageFrameInformation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageFrameBlob` | `blob` | yes |
| `contentType` | `string` | no |

## GetImageSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |
| `versionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |
| `versionId` | `string` | yes |
| `imageSetState` | `string` | yes |
| `imageSetWorkflowStatus` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `deletedAt` | `timestamp` | no |
| `message` | `string` | no |
| `imageSetArn` | `string` | no |
| `overrides` | `Overrides` | no |
| `isPrimary` | `boolean` | no |
| `lastAccessedAt` | `timestamp` | no |
| `storageTier` | `string` | no |

## GetImageSetMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |
| `versionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageSetMetadataBlob` | `blob` | yes |
| `contentType` | `string` | no |
| `contentEncoding` | `string` | no |

## ListDICOMImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `jobStatus` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummaries` | `List<DICOMImportJobSummary>` | yes |
| `nextToken` | `string` | no |

## ListDatastores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreStatus` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreSummaries` | `List<DatastoreSummary>` | no |
| `nextToken` | `string` | no |

## ListImageSetVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageSetPropertiesList` | `List<ImageSetProperties>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## SearchImageSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `searchCriteria` | `SearchCriteria` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageSetsMetadataSummaries` | `List<ImageSetsMetadataSummary>` | yes |
| `sort` | `Sort` | no |
| `nextToken` | `string` | no |

## StartDICOMImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | no |
| `dataAccessRoleArn` | `string` | yes |
| `clientToken` | `string` | yes |
| `datastoreId` | `string` | yes |
| `inputS3Uri` | `string` | yes |
| `outputS3Uri` | `string` | yes |
| `inputOwnerAccountId` | `string` | no |
| `importConfiguration` | `ImportConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `jobId` | `string` | yes |
| `jobStatus` | `string` | yes |
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


## UpdateImageSetMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |
| `latestVersionId` | `string` | yes |
| `force` | `boolean` | no |
| `includeStudyImageSets` | `boolean` | no |
| `updateImageSetMetadataUpdates` | `MetadataUpdates` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datastoreId` | `string` | yes |
| `imageSetId` | `string` | yes |
| `latestVersionId` | `string` | yes |
| `imageSetState` | `string` | yes |
| `imageSetWorkflowStatus` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `message` | `string` | no |

