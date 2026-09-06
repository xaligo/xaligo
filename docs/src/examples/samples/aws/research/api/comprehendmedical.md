# AWS Comprehend Medical

API version: 2018-10-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/comprehendmedical/2018-10-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DescribeEntitiesDetectionV2Job

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobProperties` | `ComprehendMedicalAsyncJobProperties` | no |

## DescribeICD10CMInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobProperties` | `ComprehendMedicalAsyncJobProperties` | no |

## DescribePHIDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobProperties` | `ComprehendMedicalAsyncJobProperties` | no |

## DescribeRxNormInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobProperties` | `ComprehendMedicalAsyncJobProperties` | no |

## DescribeSNOMEDCTInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobProperties` | `ComprehendMedicalAsyncJobProperties` | no |

## DetectEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<Entity>` | yes |
| `UnmappedAttributes` | `List<UnmappedAttribute>` | no |
| `PaginationToken` | `string` | no |
| `ModelVersion` | `string` | yes |

## DetectEntitiesV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<Entity>` | yes |
| `UnmappedAttributes` | `List<UnmappedAttribute>` | no |
| `PaginationToken` | `string` | no |
| `ModelVersion` | `string` | yes |

## DetectPHI

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<Entity>` | yes |
| `PaginationToken` | `string` | no |
| `ModelVersion` | `string` | yes |

## InferICD10CM

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<ICD10CMEntity>` | yes |
| `PaginationToken` | `string` | no |
| `ModelVersion` | `string` | no |

## InferRxNorm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<RxNormEntity>` | yes |
| `PaginationToken` | `string` | no |
| `ModelVersion` | `string` | no |

## InferSNOMEDCT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<SNOMEDCTEntity>` | yes |
| `PaginationToken` | `string` | no |
| `ModelVersion` | `string` | no |
| `SNOMEDCTDetails` | `SNOMEDCTDetails` | no |
| `Characters` | `Characters` | no |

## ListEntitiesDetectionV2Jobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `ComprehendMedicalAsyncJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobPropertiesList` | `List<ComprehendMedicalAsyncJobProperties>` | no |
| `NextToken` | `string` | no |

## ListICD10CMInferenceJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `ComprehendMedicalAsyncJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobPropertiesList` | `List<ComprehendMedicalAsyncJobProperties>` | no |
| `NextToken` | `string` | no |

## ListPHIDetectionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `ComprehendMedicalAsyncJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobPropertiesList` | `List<ComprehendMedicalAsyncJobProperties>` | no |
| `NextToken` | `string` | no |

## ListRxNormInferenceJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `ComprehendMedicalAsyncJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobPropertiesList` | `List<ComprehendMedicalAsyncJobProperties>` | no |
| `NextToken` | `string` | no |

## ListSNOMEDCTInferenceJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `ComprehendMedicalAsyncJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComprehendMedicalAsyncJobPropertiesList` | `List<ComprehendMedicalAsyncJobProperties>` | no |
| `NextToken` | `string` | no |

## StartEntitiesDetectionV2Job

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `KMSKey` | `string` | no |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartICD10CMInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `KMSKey` | `string` | no |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartPHIDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `KMSKey` | `string` | no |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartRxNormInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `KMSKey` | `string` | no |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartSNOMEDCTInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `JobName` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `KMSKey` | `string` | no |
| `LanguageCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StopEntitiesDetectionV2Job

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StopICD10CMInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StopPHIDetectionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StopRxNormInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StopSNOMEDCTInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

