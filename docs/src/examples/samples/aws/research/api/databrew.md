# AWS Glue DataBrew

API version: 2017-07-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/databrew/2017-07-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchDeleteRecipeVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RecipeVersions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Errors` | `List<RecipeVersionErrorDetail>` | no |

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Format` | `string` | no |
| `FormatOptions` | `FormatOptions` | no |
| `Input` | `Input` | yes |
| `PathOptions` | `PathOptions` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## CreateProfileJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | yes |
| `EncryptionKeyArn` | `string` | no |
| `EncryptionMode` | `string` | no |
| `Name` | `string` | yes |
| `LogSubscription` | `string` | no |
| `MaxCapacity` | `integer` | no |
| `MaxRetries` | `integer` | no |
| `OutputLocation` | `S3Location` | yes |
| `Configuration` | `ProfileConfiguration` | no |
| `ValidationConfigurations` | `List<ValidationConfiguration>` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `Timeout` | `integer` | no |
| `JobSample` | `JobSample` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## CreateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | yes |
| `Name` | `string` | yes |
| `RecipeName` | `string` | yes |
| `Sample` | `Sample` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## CreateRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Steps` | `List<RecipeStep>` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## CreateRecipeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | no |
| `EncryptionKeyArn` | `string` | no |
| `EncryptionMode` | `string` | no |
| `Name` | `string` | yes |
| `LogSubscription` | `string` | no |
| `MaxCapacity` | `integer` | no |
| `MaxRetries` | `integer` | no |
| `Outputs` | `List<Output>` | no |
| `DataCatalogOutputs` | `List<DataCatalogOutput>` | no |
| `DatabaseOutputs` | `List<DatabaseOutput>` | no |
| `ProjectName` | `string` | no |
| `RecipeReference` | `RecipeReference` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `Timeout` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## CreateRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `TargetArn` | `string` | yes |
| `Rules` | `List<Rule>` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## CreateSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobNames` | `List<string>` | no |
| `CronExpression` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## DeleteJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## DeleteProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## DeleteRecipeVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RecipeVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RecipeVersion` | `string` | yes |

## DeleteRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## DeleteSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## DescribeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedBy` | `string` | no |
| `CreateDate` | `timestamp` | no |
| `Name` | `string` | yes |
| `Format` | `string` | no |
| `FormatOptions` | `FormatOptions` | no |
| `Input` | `Input` | yes |
| `LastModifiedDate` | `timestamp` | no |
| `LastModifiedBy` | `string` | no |
| `Source` | `string` | no |
| `PathOptions` | `PathOptions` | no |
| `Tags` | `Map<string>` | no |
| `ResourceArn` | `string` | no |

## DescribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateDate` | `timestamp` | no |
| `CreatedBy` | `string` | no |
| `DatasetName` | `string` | no |
| `EncryptionKeyArn` | `string` | no |
| `EncryptionMode` | `string` | no |
| `Name` | `string` | yes |
| `Type` | `string` | no |
| `LastModifiedBy` | `string` | no |
| `LastModifiedDate` | `timestamp` | no |
| `LogSubscription` | `string` | no |
| `MaxCapacity` | `integer` | no |
| `MaxRetries` | `integer` | no |
| `Outputs` | `List<Output>` | no |
| `DataCatalogOutputs` | `List<DataCatalogOutput>` | no |
| `DatabaseOutputs` | `List<DatabaseOutput>` | no |
| `ProjectName` | `string` | no |
| `ProfileConfiguration` | `ProfileConfiguration` | no |
| `ValidationConfigurations` | `List<ValidationConfiguration>` | no |
| `RecipeReference` | `RecipeReference` | no |
| `ResourceArn` | `string` | no |
| `RoleArn` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Timeout` | `integer` | no |
| `JobSample` | `JobSample` | no |

## DescribeJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attempt` | `integer` | no |
| `CompletedOn` | `timestamp` | no |
| `DatasetName` | `string` | no |
| `ErrorMessage` | `string` | no |
| `ExecutionTime` | `integer` | no |
| `JobName` | `string` | yes |
| `ProfileConfiguration` | `ProfileConfiguration` | no |
| `ValidationConfigurations` | `List<ValidationConfiguration>` | no |
| `RunId` | `string` | no |
| `State` | `string` | no |
| `LogSubscription` | `string` | no |
| `LogGroupName` | `string` | no |
| `Outputs` | `List<Output>` | no |
| `DataCatalogOutputs` | `List<DataCatalogOutput>` | no |
| `DatabaseOutputs` | `List<DatabaseOutput>` | no |
| `RecipeReference` | `RecipeReference` | no |
| `StartedBy` | `string` | no |
| `StartedOn` | `timestamp` | no |
| `JobSample` | `JobSample` | no |

## DescribeProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateDate` | `timestamp` | no |
| `CreatedBy` | `string` | no |
| `DatasetName` | `string` | no |
| `LastModifiedDate` | `timestamp` | no |
| `LastModifiedBy` | `string` | no |
| `Name` | `string` | yes |
| `RecipeName` | `string` | no |
| `ResourceArn` | `string` | no |
| `Sample` | `Sample` | no |
| `RoleArn` | `string` | no |
| `Tags` | `Map<string>` | no |
| `SessionStatus` | `string` | no |
| `OpenedBy` | `string` | no |
| `OpenDate` | `timestamp` | no |

## DescribeRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RecipeVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedBy` | `string` | no |
| `CreateDate` | `timestamp` | no |
| `LastModifiedBy` | `string` | no |
| `LastModifiedDate` | `timestamp` | no |
| `ProjectName` | `string` | no |
| `PublishedBy` | `string` | no |
| `PublishedDate` | `timestamp` | no |
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Steps` | `List<RecipeStep>` | no |
| `Tags` | `Map<string>` | no |
| `ResourceArn` | `string` | no |
| `RecipeVersion` | `string` | no |

## DescribeRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `TargetArn` | `string` | no |
| `Rules` | `List<Rule>` | no |
| `CreateDate` | `timestamp` | no |
| `CreatedBy` | `string` | no |
| `LastModifiedBy` | `string` | no |
| `LastModifiedDate` | `timestamp` | no |
| `ResourceArn` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribeSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateDate` | `timestamp` | no |
| `CreatedBy` | `string` | no |
| `JobNames` | `List<string>` | no |
| `LastModifiedBy` | `string` | no |
| `LastModifiedDate` | `timestamp` | no |
| `ResourceArn` | `string` | no |
| `CronExpression` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Name` | `string` | yes |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Datasets` | `List<Dataset>` | yes |
| `NextToken` | `string` | no |

## ListJobRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobRuns` | `List<JobRun>` | yes |
| `NextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ProjectName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<Job>` | yes |
| `NextToken` | `string` | no |

## ListProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Projects` | `List<Project>` | yes |
| `NextToken` | `string` | no |

## ListRecipeVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Recipes` | `List<Recipe>` | yes |

## ListRecipes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RecipeVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Recipes` | `List<Recipe>` | yes |
| `NextToken` | `string` | no |

## ListRulesets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rulesets` | `List<RulesetItem>` | yes |
| `NextToken` | `string` | no |

## ListSchedules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Schedules` | `List<Schedule>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## PublishRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## SendProjectSessionAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Preview` | `boolean` | no |
| `Name` | `string` | yes |
| `RecipeStep` | `RecipeStep` | no |
| `StepIndex` | `integer` | no |
| `ClientSessionId` | `string` | no |
| `ViewFrame` | `ViewFrame` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Result` | `string` | no |
| `Name` | `string` | yes |
| `ActionId` | `integer` | no |

## StartJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | yes |

## StartProjectSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `AssumeControl` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ClientSessionId` | `string` | no |

## StopJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | yes |

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


## UpdateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Format` | `string` | no |
| `FormatOptions` | `FormatOptions` | no |
| `Input` | `Input` | yes |
| `PathOptions` | `PathOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## UpdateProfileJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configuration` | `ProfileConfiguration` | no |
| `EncryptionKeyArn` | `string` | no |
| `EncryptionMode` | `string` | no |
| `Name` | `string` | yes |
| `LogSubscription` | `string` | no |
| `MaxCapacity` | `integer` | no |
| `MaxRetries` | `integer` | no |
| `OutputLocation` | `S3Location` | yes |
| `ValidationConfigurations` | `List<ValidationConfiguration>` | no |
| `RoleArn` | `string` | yes |
| `Timeout` | `integer` | no |
| `JobSample` | `JobSample` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## UpdateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sample` | `Sample` | no |
| `RoleArn` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastModifiedDate` | `timestamp` | no |
| `Name` | `string` | yes |

## UpdateRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Steps` | `List<RecipeStep>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## UpdateRecipeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EncryptionKeyArn` | `string` | no |
| `EncryptionMode` | `string` | no |
| `Name` | `string` | yes |
| `LogSubscription` | `string` | no |
| `MaxCapacity` | `integer` | no |
| `MaxRetries` | `integer` | no |
| `Outputs` | `List<Output>` | no |
| `DataCatalogOutputs` | `List<DataCatalogOutput>` | no |
| `DatabaseOutputs` | `List<DatabaseOutput>` | no |
| `RoleArn` | `string` | yes |
| `Timeout` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## UpdateRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Rules` | `List<Rule>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## UpdateSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobNames` | `List<string>` | no |
| `CronExpression` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

