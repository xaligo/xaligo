# Amazon Rekognition

API version: 2016-06-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/rekognition/2016-06-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateFaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `UserId` | `string` | yes |
| `FaceIds` | `List<string>` | yes |
| `UserMatchThreshold` | `float` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedFaces` | `List<AssociatedFace>` | no |
| `UnsuccessfulFaceAssociations` | `List<UnsuccessfulFaceAssociation>` | no |
| `UserStatus` | `string` | no |

## CompareFaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceImage` | `Image` | yes |
| `TargetImage` | `Image` | yes |
| `SimilarityThreshold` | `float` | no |
| `QualityFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceImageFace` | `ComparedSourceImageFace` | no |
| `FaceMatches` | `List<CompareFacesMatch>` | no |
| `UnmatchedFaces` | `List<ComparedFace>` | no |
| `SourceImageOrientationCorrection` | `string` | no |
| `TargetImageOrientationCorrection` | `string` | no |

## CopyProjectVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceProjectArn` | `string` | yes |
| `SourceProjectVersionArn` | `string` | yes |
| `DestinationProjectArn` | `string` | yes |
| `VersionName` | `string` | yes |
| `OutputConfig` | `OutputConfig` | yes |
| `Tags` | `Map<string>` | no |
| `KmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectVersionArn` | `string` | no |

## CreateCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatusCode` | `integer` | no |
| `CollectionArn` | `string` | no |
| `FaceModelVersion` | `string` | no |

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetSource` | `DatasetSource` | no |
| `DatasetType` | `string` | yes |
| `ProjectArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | no |

## CreateFaceLivenessSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KmsKeyId` | `string` | no |
| `Settings` | `CreateFaceLivenessSessionRequestSettings` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |

## CreateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectName` | `string` | yes |
| `Feature` | `string` | no |
| `AutoUpdate` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | no |

## CreateProjectVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |
| `VersionName` | `string` | yes |
| `OutputConfig` | `OutputConfig` | yes |
| `TrainingData` | `TrainingData` | no |
| `TestingData` | `TestingData` | no |
| `Tags` | `Map<string>` | no |
| `KmsKeyId` | `string` | no |
| `VersionDescription` | `string` | no |
| `FeatureConfig` | `CustomizationFeatureConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectVersionArn` | `string` | no |

## CreateStreamProcessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Input` | `StreamProcessorInput` | yes |
| `Output` | `StreamProcessorOutput` | yes |
| `Name` | `string` | yes |
| `Settings` | `StreamProcessorSettings` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `NotificationChannel` | `StreamProcessorNotificationChannel` | no |
| `KmsKeyId` | `string` | no |
| `RegionsOfInterest` | `List<RegionOfInterest>` | no |
| `DataSharingPreference` | `StreamProcessorDataSharingPreference` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamProcessorArn` | `string` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `UserId` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatusCode` | `integer` | no |

## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `FaceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletedFaces` | `List<string>` | no |
| `UnsuccessfulFaceDeletions` | `List<UnsuccessfulFaceDeletion>` | no |

## DeleteProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## DeleteProjectPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProjectVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## DeleteStreamProcessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `UserId` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FaceCount` | `long` | no |
| `FaceModelVersion` | `string` | no |
| `CollectionARN` | `string` | no |
| `CreationTimestamp` | `timestamp` | no |
| `UserCount` | `long` | no |

## DescribeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetDescription` | `DatasetDescription` | no |

## DescribeProjectVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |
| `VersionNames` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectVersionDescriptions` | `List<ProjectVersionDescription>` | no |
| `NextToken` | `string` | no |

## DescribeProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ProjectNames` | `List<string>` | no |
| `Features` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectDescriptions` | `List<ProjectDescription>` | no |
| `NextToken` | `string` | no |

## DescribeStreamProcessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `StreamProcessorArn` | `string` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `CreationTimestamp` | `timestamp` | no |
| `LastUpdateTimestamp` | `timestamp` | no |
| `Input` | `StreamProcessorInput` | no |
| `Output` | `StreamProcessorOutput` | no |
| `RoleArn` | `string` | no |
| `Settings` | `StreamProcessorSettings` | no |
| `NotificationChannel` | `StreamProcessorNotificationChannel` | no |
| `KmsKeyId` | `string` | no |
| `RegionsOfInterest` | `List<RegionOfInterest>` | no |
| `DataSharingPreference` | `StreamProcessorDataSharingPreference` | no |

## DetectCustomLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectVersionArn` | `string` | yes |
| `Image` | `Image` | yes |
| `MaxResults` | `integer` | no |
| `MinConfidence` | `float` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomLabels` | `List<CustomLabel>` | no |

## DetectFaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Image` | `Image` | yes |
| `Attributes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FaceDetails` | `List<FaceDetail>` | no |
| `OrientationCorrection` | `string` | no |

## DetectLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Image` | `Image` | yes |
| `MaxLabels` | `integer` | no |
| `MinConfidence` | `float` | no |
| `Features` | `List<string>` | no |
| `Settings` | `DetectLabelsSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Labels` | `List<Label>` | no |
| `OrientationCorrection` | `string` | no |
| `LabelModelVersion` | `string` | no |
| `ImageProperties` | `DetectLabelsImageProperties` | no |

## DetectModerationLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Image` | `Image` | yes |
| `MinConfidence` | `float` | no |
| `HumanLoopConfig` | `HumanLoopConfig` | no |
| `ProjectVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModerationLabels` | `List<ModerationLabel>` | no |
| `ModerationModelVersion` | `string` | no |
| `HumanLoopActivationOutput` | `HumanLoopActivationOutput` | no |
| `ProjectVersion` | `string` | no |
| `ContentTypes` | `List<ContentType>` | no |

## DetectProtectiveEquipment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Image` | `Image` | yes |
| `SummarizationAttributes` | `ProtectiveEquipmentSummarizationAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectiveEquipmentModelVersion` | `string` | no |
| `Persons` | `List<ProtectiveEquipmentPerson>` | no |
| `Summary` | `ProtectiveEquipmentSummary` | no |

## DetectText

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Image` | `Image` | yes |
| `Filters` | `DetectTextFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextDetections` | `List<TextDetection>` | no |
| `TextModelVersion` | `string` | no |

## DisassociateFaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `UserId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `FaceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisassociatedFaces` | `List<DisassociatedFace>` | no |
| `UnsuccessfulFaceDisassociations` | `List<UnsuccessfulFaceDisassociation>` | no |
| `UserStatus` | `string` | no |

## DistributeDatasetEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Datasets` | `List<DistributeDataset>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetCelebrityInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Urls` | `List<string>` | no |
| `Name` | `string` | no |
| `KnownGender` | `KnownGender` | no |

## GetCelebrityRecognition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `VideoMetadata` | `VideoMetadata` | no |
| `NextToken` | `string` | no |
| `Celebrities` | `List<CelebrityRecognition>` | no |
| `JobId` | `string` | no |
| `Video` | `Video` | no |
| `JobTag` | `string` | no |

## GetContentModeration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `AggregateBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `VideoMetadata` | `VideoMetadata` | no |
| `ModerationLabels` | `List<ContentModerationDetection>` | no |
| `NextToken` | `string` | no |
| `ModerationModelVersion` | `string` | no |
| `JobId` | `string` | no |
| `Video` | `Video` | no |
| `JobTag` | `string` | no |
| `GetRequestMetadata` | `GetContentModerationRequestMetadata` | no |

## GetFaceDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `VideoMetadata` | `VideoMetadata` | no |
| `NextToken` | `string` | no |
| `Faces` | `List<FaceDetection>` | no |
| `JobId` | `string` | no |
| `Video` | `Video` | no |
| `JobTag` | `string` | no |

## GetFaceLivenessSessionResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |
| `Status` | `string` | yes |
| `Confidence` | `float` | no |
| `ReferenceImage` | `AuditImage` | no |
| `AuditImages` | `List<AuditImage>` | no |
| `Challenge` | `Challenge` | no |

## GetFaceSearch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `NextToken` | `string` | no |
| `VideoMetadata` | `VideoMetadata` | no |
| `Persons` | `List<PersonMatch>` | no |
| `JobId` | `string` | no |
| `Video` | `Video` | no |
| `JobTag` | `string` | no |

## GetLabelDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `AggregateBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `VideoMetadata` | `VideoMetadata` | no |
| `NextToken` | `string` | no |
| `Labels` | `List<LabelDetection>` | no |
| `LabelModelVersion` | `string` | no |
| `JobId` | `string` | no |
| `Video` | `Video` | no |
| `JobTag` | `string` | no |
| `GetRequestMetadata` | `GetLabelDetectionRequestMetadata` | no |

## GetMediaAnalysisJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `JobName` | `string` | no |
| `OperationsConfig` | `MediaAnalysisOperationsConfig` | yes |
| `Status` | `string` | yes |
| `FailureDetails` | `MediaAnalysisJobFailureDetails` | no |
| `CreationTimestamp` | `timestamp` | yes |
| `CompletionTimestamp` | `timestamp` | no |
| `Input` | `MediaAnalysisInput` | yes |
| `OutputConfig` | `MediaAnalysisOutputConfig` | yes |
| `KmsKeyId` | `string` | no |
| `Results` | `MediaAnalysisResults` | no |
| `ManifestSummary` | `MediaAnalysisManifestSummary` | no |

## GetPersonTracking

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `VideoMetadata` | `VideoMetadata` | no |
| `NextToken` | `string` | no |
| `Persons` | `List<PersonDetection>` | no |
| `JobId` | `string` | no |
| `Video` | `Video` | no |
| `JobTag` | `string` | no |

## GetSegmentDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `VideoMetadata` | `List<VideoMetadata>` | no |
| `AudioMetadata` | `List<AudioMetadata>` | no |
| `NextToken` | `string` | no |
| `Segments` | `List<SegmentDetection>` | no |
| `SelectedSegmentTypes` | `List<SegmentTypeInfo>` | no |
| `JobId` | `string` | no |
| `Video` | `Video` | no |
| `JobTag` | `string` | no |

## GetTextDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `VideoMetadata` | `VideoMetadata` | no |
| `TextDetections` | `List<TextDetectionResult>` | no |
| `NextToken` | `string` | no |
| `TextModelVersion` | `string` | no |
| `JobId` | `string` | no |
| `Video` | `Video` | no |
| `JobTag` | `string` | no |

## IndexFaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `Image` | `Image` | yes |
| `ExternalImageId` | `string` | no |
| `DetectionAttributes` | `List<string>` | no |
| `MaxFaces` | `integer` | no |
| `QualityFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FaceRecords` | `List<FaceRecord>` | no |
| `OrientationCorrection` | `string` | no |
| `FaceModelVersion` | `string` | no |
| `UnindexedFaces` | `List<UnindexedFace>` | no |

## ListCollections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `FaceModelVersions` | `List<string>` | no |

## ListDatasetEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | yes |
| `ContainsLabels` | `List<string>` | no |
| `Labeled` | `boolean` | no |
| `SourceRefContains` | `string` | no |
| `HasErrors` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetEntries` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListDatasetLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetLabelDescriptions` | `List<DatasetLabelDescription>` | no |
| `NextToken` | `string` | no |

## ListFaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `UserId` | `string` | no |
| `FaceIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Faces` | `List<Face>` | no |
| `NextToken` | `string` | no |
| `FaceModelVersion` | `string` | no |

## ListMediaAnalysisJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MediaAnalysisJobs` | `List<MediaAnalysisJobDescription>` | yes |

## ListProjectPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectPolicies` | `List<ProjectPolicy>` | no |
| `NextToken` | `string` | no |

## ListStreamProcessors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `StreamProcessors` | `List<StreamProcessor>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | no |
| `NextToken` | `string` | no |

## PutProjectPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyRevisionId` | `string` | no |
| `PolicyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyRevisionId` | `string` | no |

## RecognizeCelebrities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Image` | `Image` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CelebrityFaces` | `List<Celebrity>` | no |
| `UnrecognizedFaces` | `List<ComparedFace>` | no |
| `OrientationCorrection` | `string` | no |

## SearchFaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `FaceId` | `string` | yes |
| `MaxFaces` | `integer` | no |
| `FaceMatchThreshold` | `float` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchedFaceId` | `string` | no |
| `FaceMatches` | `List<FaceMatch>` | no |
| `FaceModelVersion` | `string` | no |

## SearchFacesByImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `Image` | `Image` | yes |
| `MaxFaces` | `integer` | no |
| `FaceMatchThreshold` | `float` | no |
| `QualityFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchedFaceBoundingBox` | `BoundingBox` | no |
| `SearchedFaceConfidence` | `float` | no |
| `FaceMatches` | `List<FaceMatch>` | no |
| `FaceModelVersion` | `string` | no |

## SearchUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `UserId` | `string` | no |
| `FaceId` | `string` | no |
| `UserMatchThreshold` | `float` | no |
| `MaxUsers` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserMatches` | `List<UserMatch>` | no |
| `FaceModelVersion` | `string` | no |
| `SearchedFace` | `SearchedFace` | no |
| `SearchedUser` | `SearchedUser` | no |

## SearchUsersByImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CollectionId` | `string` | yes |
| `Image` | `Image` | yes |
| `UserMatchThreshold` | `float` | no |
| `MaxUsers` | `integer` | no |
| `QualityFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserMatches` | `List<UserMatch>` | no |
| `FaceModelVersion` | `string` | no |
| `SearchedFace` | `SearchedFaceDetails` | no |
| `UnsearchedFaces` | `List<UnsearchedFace>` | no |

## StartCelebrityRecognition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Video` | `Video` | yes |
| `ClientRequestToken` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `JobTag` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartContentModeration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Video` | `Video` | yes |
| `MinConfidence` | `float` | no |
| `ClientRequestToken` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `JobTag` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartFaceDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Video` | `Video` | yes |
| `ClientRequestToken` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `FaceAttributes` | `string` | no |
| `JobTag` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartFaceSearch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Video` | `Video` | yes |
| `ClientRequestToken` | `string` | no |
| `FaceMatchThreshold` | `float` | no |
| `CollectionId` | `string` | yes |
| `NotificationChannel` | `NotificationChannel` | no |
| `JobTag` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartLabelDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Video` | `Video` | yes |
| `ClientRequestToken` | `string` | no |
| `MinConfidence` | `float` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `JobTag` | `string` | no |
| `Features` | `List<string>` | no |
| `Settings` | `LabelDetectionSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartMediaAnalysisJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `JobName` | `string` | no |
| `OperationsConfig` | `MediaAnalysisOperationsConfig` | yes |
| `Input` | `MediaAnalysisInput` | yes |
| `OutputConfig` | `MediaAnalysisOutputConfig` | yes |
| `KmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

## StartPersonTracking

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Video` | `Video` | yes |
| `ClientRequestToken` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `JobTag` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartProjectVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectVersionArn` | `string` | yes |
| `MinInferenceUnits` | `integer` | yes |
| `MaxInferenceUnits` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## StartSegmentDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Video` | `Video` | yes |
| `ClientRequestToken` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `JobTag` | `string` | no |
| `Filters` | `StartSegmentDetectionFilters` | no |
| `SegmentTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartStreamProcessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `StartSelector` | `StreamProcessingStartSelector` | no |
| `StopSelector` | `StreamProcessingStopSelector` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | no |

## StartTextDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Video` | `Video` | yes |
| `ClientRequestToken` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `JobTag` | `string` | no |
| `Filters` | `StartTextDetectionFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StopProjectVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## StopStreamProcessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

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


## UpdateDatasetEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | yes |
| `Changes` | `DatasetChanges` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStreamProcessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SettingsForUpdate` | `StreamProcessorSettingsForUpdate` | no |
| `RegionsOfInterestForUpdate` | `List<RegionOfInterest>` | no |
| `DataSharingPreferenceForUpdate` | `StreamProcessorDataSharingPreference` | no |
| `ParametersToDelete` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


