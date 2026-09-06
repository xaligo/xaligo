# Amazon Timestream Write

API version: 2018-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/timestream-write/2018-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateBatchLoadTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DataModelConfiguration` | `DataModelConfiguration` | no |
| `DataSourceConfiguration` | `DataSourceConfiguration` | yes |
| `ReportConfiguration` | `ReportConfiguration` | yes |
| `TargetDatabaseName` | `string` | yes |
| `TargetTableName` | `string` | yes |
| `RecordVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |

## CreateDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Database` | `Database` | no |

## CreateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `RetentionProperties` | `RetentionProperties` | no |
| `Tags` | `List<Tag>` | no |
| `MagneticStoreWriteProperties` | `MagneticStoreWriteProperties` | no |
| `Schema` | `Schema` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Table` | `Table` | no |

## DeleteDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeBatchLoadTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchLoadTaskDescription` | `BatchLoadTaskDescription` | yes |

## DescribeDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Database` | `Database` | no |

## DescribeEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<Endpoint>` | yes |

## DescribeTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Table` | `Table` | no |

## ListBatchLoadTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `TaskStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `BatchLoadTasks` | `List<BatchLoadTask>` | no |

## ListDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Databases` | `List<Database>` | no |
| `NextToken` | `string` | no |

## ListTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tables` | `List<Table>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ResumeBatchLoadTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `KmsKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Database` | `Database` | no |

## UpdateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `RetentionProperties` | `RetentionProperties` | no |
| `MagneticStoreWriteProperties` | `MagneticStoreWriteProperties` | no |
| `Schema` | `Schema` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Table` | `Table` | no |

## WriteRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `CommonAttributes` | `Record` | no |
| `Records` | `List<Record>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordsIngested` | `RecordsIngested` | no |

