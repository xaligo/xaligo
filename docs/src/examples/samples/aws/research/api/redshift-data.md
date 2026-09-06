# Redshift Data API Service

API version: 2019-12-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/redshift-data/2019-12-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchExecuteStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sqls` | `List<string>` | yes |
| `ClusterIdentifier` | `string` | no |
| `SecretArn` | `string` | no |
| `DbUser` | `string` | no |
| `Database` | `string` | no |
| `WithEvent` | `boolean` | no |
| `StatementName` | `string` | no |
| `Parameters` | `List<SqlParameter>` | no |
| `WorkgroupName` | `string` | no |
| `ClientToken` | `string` | no |
| `ResultFormat` | `string` | no |
| `SessionKeepAliveSeconds` | `integer` | no |
| `SessionId` | `string` | no |
| `ExecutionMode` | `string` | no |
| `WaitTimeSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `ClusterIdentifier` | `string` | no |
| `DbUser` | `string` | no |
| `DbGroups` | `List<string>` | no |
| `Database` | `string` | no |
| `SecretArn` | `string` | no |
| `WorkgroupName` | `string` | no |
| `SessionId` | `string` | no |
| `Status` | `string` | no |
| `RedshiftPid` | `long` | no |
| `HasResultSet` | `boolean` | no |

## CancelStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `boolean` | no |

## DescribeStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `WaitTimeSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `SecretArn` | `string` | no |
| `DbUser` | `string` | no |
| `Database` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `Duration` | `long` | no |
| `Error` | `string` | no |
| `Status` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `RedshiftPid` | `long` | no |
| `HasResultSet` | `boolean` | no |
| `QueryString` | `string` | no |
| `ResultRows` | `long` | no |
| `ResultSize` | `long` | no |
| `RedshiftQueryId` | `long` | no |
| `QueryParameters` | `List<SqlParameter>` | no |
| `SubStatements` | `List<SubStatementData>` | no |
| `WorkgroupName` | `string` | no |
| `ResultFormat` | `string` | no |
| `SessionId` | `string` | no |
| `ExecutionMode` | `string` | no |

## DescribeTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `SecretArn` | `string` | no |
| `DbUser` | `string` | no |
| `Database` | `string` | yes |
| `ConnectedDatabase` | `string` | no |
| `Schema` | `string` | no |
| `Table` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `WorkgroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `ColumnList` | `List<ColumnMetadata>` | no |
| `NextToken` | `string` | no |

## ExecuteStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sql` | `string` | yes |
| `ClusterIdentifier` | `string` | no |
| `SecretArn` | `string` | no |
| `DbUser` | `string` | no |
| `Database` | `string` | no |
| `WithEvent` | `boolean` | no |
| `StatementName` | `string` | no |
| `Parameters` | `List<SqlParameter>` | no |
| `WorkgroupName` | `string` | no |
| `ClientToken` | `string` | no |
| `ResultFormat` | `string` | no |
| `SessionKeepAliveSeconds` | `integer` | no |
| `SessionId` | `string` | no |
| `WaitTimeSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `ClusterIdentifier` | `string` | no |
| `DbUser` | `string` | no |
| `DbGroups` | `List<string>` | no |
| `Database` | `string` | no |
| `SecretArn` | `string` | no |
| `WorkgroupName` | `string` | no |
| `SessionId` | `string` | no |
| `Status` | `string` | no |
| `RedshiftPid` | `long` | no |
| `HasResultSet` | `boolean` | no |

## GetStatementResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `NextToken` | `string` | no |
| `WaitTimeSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<List<Field>>` | yes |
| `ColumnMetadata` | `List<ColumnMetadata>` | no |
| `TotalNumRows` | `long` | no |
| `NextToken` | `string` | no |

## GetStatementResultV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `NextToken` | `string` | no |
| `WaitTimeSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<QueryRecords>` | yes |
| `ColumnMetadata` | `List<ColumnMetadata>` | no |
| `TotalNumRows` | `long` | no |
| `ResultFormat` | `string` | no |
| `NextToken` | `string` | no |

## ListDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `Database` | `string` | yes |
| `SecretArn` | `string` | no |
| `DbUser` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `WorkgroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Databases` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `SecretArn` | `string` | no |
| `DbUser` | `string` | no |
| `Database` | `string` | yes |
| `ConnectedDatabase` | `string` | no |
| `SchemaPattern` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `WorkgroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Schemas` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SessionId` | `string` | no |
| `Status` | `string` | no |
| `RoleLevel` | `boolean` | no |
| `ClusterIdentifier` | `string` | no |
| `WorkgroupName` | `string` | no |
| `Database` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sessions` | `List<SessionData>` | yes |
| `NextToken` | `string` | no |

## ListStatements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StatementName` | `string` | no |
| `Status` | `string` | no |
| `RoleLevel` | `boolean` | no |
| `Database` | `string` | no |
| `ClusterIdentifier` | `string` | no |
| `WorkgroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statements` | `List<StatementData>` | yes |
| `NextToken` | `string` | no |

## ListTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterIdentifier` | `string` | no |
| `SecretArn` | `string` | no |
| `DbUser` | `string` | no |
| `Database` | `string` | yes |
| `ConnectedDatabase` | `string` | no |
| `SchemaPattern` | `string` | no |
| `TablePattern` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `WorkgroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tables` | `List<TableMember>` | no |
| `NextToken` | `string` | no |

