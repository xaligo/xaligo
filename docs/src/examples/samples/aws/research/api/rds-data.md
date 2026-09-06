# AWS RDS DataService

API version: 2018-08-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/rds-data/2018-08-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchExecuteStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `secretArn` | `string` | yes |
| `sql` | `string` | yes |
| `database` | `string` | no |
| `schema` | `string` | no |
| `parameterSets` | `List<List<SqlParameter>>` | no |
| `transactionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateResults` | `List<UpdateResult>` | no |

## BeginTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `secretArn` | `string` | yes |
| `database` | `string` | no |
| `schema` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transactionId` | `string` | no |

## CommitTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `secretArn` | `string` | yes |
| `transactionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transactionStatus` | `string` | no |

## ExecuteSql

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbClusterOrInstanceArn` | `string` | yes |
| `awsSecretStoreArn` | `string` | yes |
| `sqlStatements` | `string` | yes |
| `database` | `string` | no |
| `schema` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sqlStatementResults` | `List<SqlStatementResult>` | no |

## ExecuteStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `secretArn` | `string` | yes |
| `sql` | `string` | yes |
| `database` | `string` | no |
| `schema` | `string` | no |
| `parameters` | `List<SqlParameter>` | no |
| `transactionId` | `string` | no |
| `includeResultMetadata` | `boolean` | no |
| `continueAfterTimeout` | `boolean` | no |
| `resultSetOptions` | `ResultSetOptions` | no |
| `formatRecordsAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `records` | `List<List<Field>>` | no |
| `columnMetadata` | `List<ColumnMetadata>` | no |
| `numberOfRecordsUpdated` | `long` | no |
| `generatedFields` | `List<Field>` | no |
| `formattedRecords` | `string` | no |

## RollbackTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `secretArn` | `string` | yes |
| `transactionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transactionStatus` | `string` | no |

