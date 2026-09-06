# Amazon Comprehend

API version: 2017-11-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/comprehend/2017-11-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchDetectDominantLanguage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultList` | `List<BatchDetectDominantLanguageItemResult>` | yes |
| `ErrorList` | `List<BatchItemError>` | yes |

## BatchDetectEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextList` | `List<string>` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultList` | `List<BatchDetectEntitiesItemResult>` | yes |
| `ErrorList` | `List<BatchItemError>` | yes |

## BatchDetectKeyPhrases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextList` | `List<string>` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultList` | `List<BatchDetectKeyPhrasesItemResult>` | yes |
| `ErrorList` | `List<BatchItemError>` | yes |

## BatchDetectSentiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextList` | `List<string>` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultList` | `List<BatchDetectSentimentItemResult>` | yes |
| `ErrorList` | `List<BatchItemError>` | yes |

## BatchDetectSyntax

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextList` | `List<string>` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultList` | `List<BatchDetectSyntaxItemResult>` | yes |
| `ErrorList` | `List<BatchItemError>` | yes |

## BatchDetectTargetedSentiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextList` | `List<string>` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultList` | `List<BatchDetectTargetedSentimentItemResult>` | yes |
| `ErrorList` | `List<BatchItemError>` | yes |

## ClassifyDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | no |
| `EndpointArn` | `string` | yes |
| `Bytes` | `blob` | no |
| `DocumentReaderConfig` | `DocumentReaderConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Classes` | `List<DocumentClass>` | no |
| `Labels` | `List<DocumentLabel>` | no |
| `DocumentMetadata` | `DocumentMetadata` | no |
| `DocumentType` | `List<DocumentTypeListItem>` | no |
| `Errors` | `List<ErrorsListItem>` | no |
| `Warnings` | `List<WarningsListItem>` | no |

## ContainsPiiEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Labels` | `List<EntityLabel>` | no |

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | yes |
| `DatasetName` | `string` | yes |
| `DatasetType` | `string` | no |
| `Description` | `string` | no |
| `InputDataConfig` | `DatasetInputDataConfig` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | no |

## CreateDocumentClassifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassifierName` | `string` | yes |
| `VersionName` | `string` | no |
| `DataAccessRoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `InputDataConfig` | `DocumentClassifierInputDataConfig` | yes |
| `OutputDataConfig` | `DocumentClassifierOutputDataConfig` | no |
| `ClientRequestToken` | `string` | no |
| `LanguageCode` | `string` | yes |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Mode` | `string` | no |
| `ModelKmsKeyId` | `string` | no |
| `ModelPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassifierArn` | `string` | no |

## CreateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `ModelArn` | `string` | no |
| `DesiredInferenceUnits` | `integer` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DataAccessRoleArn` | `string` | no |
| `FlywheelArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | no |
| `ModelArn` | `string` | no |

## CreateEntityRecognizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecognizerName` | `string` | yes |
| `VersionName` | `string` | no |
| `DataAccessRoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `InputDataConfig` | `EntityRecognizerInputDataConfig` | yes |
| `ClientRequestToken` | `string` | no |
| `LanguageCode` | `string` | yes |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `ModelKmsKeyId` | `string` | no |
| `ModelPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityRecognizerArn` | `string` | no |

## CreateFlywheel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelName` | `string` | yes |
| `ActiveModelArn` | `string` | no |
| `DataAccessRoleArn` | `string` | yes |
| `TaskConfig` | `TaskConfig` | no |
| `ModelType` | `string` | no |
| `DataLakeS3Uri` | `string` | yes |
| `DataSecurityConfig` | `DataSecurityConfig` | no |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | no |
| `ActiveModelArn` | `string` | no |

## DeleteDocumentClassifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassifierArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEntityRecognizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityRecognizerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFlywheel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `PolicyRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetProperties` | `DatasetProperties` | no |

## DescribeDocumentClassificationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassificationJobProperties` | `DocumentClassificationJobProperties` | no |

## DescribeDocumentClassifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassifierArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassifierProperties` | `DocumentClassifierProperties` | no |

## DescribeDominantLanguageDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DominantLanguageDetectionJobProperties` | `DominantLanguageDetectionJobProperties` | no |

## DescribeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointProperties` | `EndpointProperties` | no |

## DescribeEntitiesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntitiesDetectionJobProperties` | `EntitiesDetectionJobProperties` | no |

## DescribeEntityRecognizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityRecognizerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityRecognizerProperties` | `EntityRecognizerProperties` | no |

## DescribeEventsDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventsDetectionJobProperties` | `EventsDetectionJobProperties` | no |

## DescribeFlywheel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelProperties` | `FlywheelProperties` | no |

## DescribeFlywheelIteration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | yes |
| `FlywheelIterationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelIterationProperties` | `FlywheelIterationProperties` | no |

## DescribeKeyPhrasesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyPhrasesDetectionJobProperties` | `KeyPhrasesDetectionJobProperties` | no |

## DescribePiiEntitiesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PiiEntitiesDetectionJobProperties` | `PiiEntitiesDetectionJobProperties` | no |

## DescribeResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePolicy` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `PolicyRevisionId` | `string` | no |

## DescribeSentimentDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SentimentDetectionJobProperties` | `SentimentDetectionJobProperties` | no |

## DescribeTargetedSentimentDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetedSentimentDetectionJobProperties` | `TargetedSentimentDetectionJobProperties` | no |

## DescribeTopicsDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicsDetectionJobProperties` | `TopicsDetectionJobProperties` | no |

## DetectDominantLanguage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Languages` | `List<DominantLanguage>` | no |

## DetectEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | no |
| `LanguageCode` | `string` | no |
| `EndpointArn` | `string` | no |
| `Bytes` | `blob` | no |
| `DocumentReaderConfig` | `DocumentReaderConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<Entity>` | no |
| `DocumentMetadata` | `DocumentMetadata` | no |
| `DocumentType` | `List<DocumentTypeListItem>` | no |
| `Blocks` | `List<Block>` | no |
| `Errors` | `List<ErrorsListItem>` | no |

## DetectKeyPhrases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyPhrases` | `List<KeyPhrase>` | no |

## DetectPiiEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<PiiEntity>` | no |

## DetectSentiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sentiment` | `string` | no |
| `SentimentScore` | `SentimentScore` | no |

## DetectSyntax

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyntaxTokens` | `List<SyntaxToken>` | no |

## DetectTargetedSentiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<TargetedSentimentEntity>` | no |

## DetectToxicContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextSegments` | `List<TextSegment>` | yes |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultList` | `List<ToxicLabels>` | no |

## ImportModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceModelArn` | `string` | yes |
| `ModelName` | `string` | no |
| `VersionName` | `string` | no |
| `ModelKmsKeyId` | `string` | no |
| `DataAccessRoleArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelArn` | `string` | no |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | no |
| `Filter` | `DatasetFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetPropertiesList` | `List<DatasetProperties>` | no |
| `NextToken` | `string` | no |

## ListDocumentClassificationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `DocumentClassificationJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassificationJobPropertiesList` | `List<DocumentClassificationJobProperties>` | no |
| `NextToken` | `string` | no |

## ListDocumentClassifierSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassifierSummariesList` | `List<DocumentClassifierSummary>` | no |
| `NextToken` | `string` | no |

## ListDocumentClassifiers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `DocumentClassifierFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassifierPropertiesList` | `List<DocumentClassifierProperties>` | no |
| `NextToken` | `string` | no |

## ListDominantLanguageDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `DominantLanguageDetectionJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DominantLanguageDetectionJobPropertiesList` | `List<DominantLanguageDetectionJobProperties>` | no |
| `NextToken` | `string` | no |

## ListEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `EndpointFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointPropertiesList` | `List<EndpointProperties>` | no |
| `NextToken` | `string` | no |

## ListEntitiesDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `EntitiesDetectionJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntitiesDetectionJobPropertiesList` | `List<EntitiesDetectionJobProperties>` | no |
| `NextToken` | `string` | no |

## ListEntityRecognizerSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityRecognizerSummariesList` | `List<EntityRecognizerSummary>` | no |
| `NextToken` | `string` | no |

## ListEntityRecognizers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `EntityRecognizerFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityRecognizerPropertiesList` | `List<EntityRecognizerProperties>` | no |
| `NextToken` | `string` | no |

## ListEventsDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `EventsDetectionJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventsDetectionJobPropertiesList` | `List<EventsDetectionJobProperties>` | no |
| `NextToken` | `string` | no |

## ListFlywheelIterationHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | yes |
| `Filter` | `FlywheelIterationFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelIterationPropertiesList` | `List<FlywheelIterationProperties>` | no |
| `NextToken` | `string` | no |

## ListFlywheels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `FlywheelFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelSummaryList` | `List<FlywheelSummary>` | no |
| `NextToken` | `string` | no |

## ListKeyPhrasesDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `KeyPhrasesDetectionJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyPhrasesDetectionJobPropertiesList` | `List<KeyPhrasesDetectionJobProperties>` | no |
| `NextToken` | `string` | no |

## ListPiiEntitiesDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `PiiEntitiesDetectionJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PiiEntitiesDetectionJobPropertiesList` | `List<PiiEntitiesDetectionJobProperties>` | no |
| `NextToken` | `string` | no |

## ListSentimentDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `SentimentDetectionJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SentimentDetectionJobPropertiesList` | `List<SentimentDetectionJobProperties>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

## ListTargetedSentimentDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `TargetedSentimentDetectionJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetedSentimentDetectionJobPropertiesList` | `List<TargetedSentimentDetectionJobProperties>` | no |
| `NextToken` | `string` | no |

## ListTopicsDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `TopicsDetectionJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicsDetectionJobPropertiesList` | `List<TopicsDetectionJobProperties>` | no |
| `NextToken` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourcePolicy` | `string` | yes |
| `PolicyRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyRevisionId` | `string` | no |

## StartDocumentClassificationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |
| `DocumentClassifierArn` | `string` | no |
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Tags` | `List<Tag>` | no |
| `FlywheelArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |
| `DocumentClassifierArn` | `string` | no |

## StartDominantLanguageDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |

## StartEntitiesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `EntityRecognizerArn` | `string` | no |
| `LanguageCode` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Tags` | `List<Tag>` | no |
| `FlywheelArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |
| `EntityRecognizerArn` | `string` | no |

## StartEventsDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `LanguageCode` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `TargetEventTypes` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |

## StartFlywheelIteration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | no |
| `FlywheelIterationId` | `string` | no |

## StartKeyPhrasesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `LanguageCode` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |

## StartPiiEntitiesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `Mode` | `string` | yes |
| `RedactionConfig` | `RedactionConfig` | no |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `LanguageCode` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |

## StartSentimentDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `LanguageCode` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |

## StartTargetedSentimentDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `LanguageCode` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |

## StartTopicsDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `NumberOfTopics` | `integer` | no |
| `ClientRequestToken` | `string` | no |
| `VolumeKmsKeyId` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobArn` | `string` | no |
| `JobStatus` | `string` | no |

## StopDominantLanguageDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## StopEntitiesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## StopEventsDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## StopKeyPhrasesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## StopPiiEntitiesDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## StopSentimentDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## StopTargetedSentimentDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## StopTrainingDocumentClassifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentClassifierArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopTrainingEntityRecognizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityRecognizerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


## UpdateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |
| `DesiredModelArn` | `string` | no |
| `DesiredInferenceUnits` | `integer` | no |
| `DesiredDataAccessRoleArn` | `string` | no |
| `FlywheelArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DesiredModelArn` | `string` | no |

## UpdateFlywheel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelArn` | `string` | yes |
| `ActiveModelArn` | `string` | no |
| `DataAccessRoleArn` | `string` | no |
| `DataSecurityConfig` | `UpdateDataSecurityConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlywheelProperties` | `FlywheelProperties` | no |

