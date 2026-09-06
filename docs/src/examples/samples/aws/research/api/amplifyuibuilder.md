# AWS Amplify UI Builder

API version: 2021-08-11. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/amplifyuibuilder/2021-08-11/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `clientToken` | `string` | no |
| `componentToCreate` | `CreateComponentData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entity` | `Component` | no |

## CreateForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `clientToken` | `string` | no |
| `formToCreate` | `CreateFormData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entity` | `Form` | no |

## CreateTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `clientToken` | `string` | no |
| `themeToCreate` | `CreateThemeData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entity` | `Theme` | no |

## DeleteComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExchangeCodeForToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `provider` | `string` | yes |
| `request` | `ExchangeCodeForTokenRequestBody` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessToken` | `string` | yes |
| `expiresIn` | `integer` | yes |
| `refreshToken` | `string` | yes |

## ExportComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<Component>` | yes |
| `nextToken` | `string` | no |

## ExportForms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<Form>` | yes |
| `nextToken` | `string` | no |

## ExportThemes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<Theme>` | yes |
| `nextToken` | `string` | no |

## GetCodegenJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `CodegenJob` | no |

## GetComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `component` | `Component` | no |

## GetForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `form` | `Form` | no |

## GetMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `features` | `Map<string>` | yes |

## GetTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `theme` | `Theme` | no |

## ListCodegenJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<CodegenJobSummary>` | yes |
| `nextToken` | `string` | no |

## ListComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<ComponentSummary>` | yes |
| `nextToken` | `string` | no |

## ListForms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<FormSummary>` | yes |
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

## ListThemes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<ThemeSummary>` | yes |
| `nextToken` | `string` | no |

## PutMetadataFlag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `featureName` | `string` | yes |
| `body` | `PutMetadataFlagBody` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RefreshToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `provider` | `string` | yes |
| `refreshTokenBody` | `RefreshTokenRequestBody` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessToken` | `string` | yes |
| `expiresIn` | `integer` | yes |

## StartCodegenJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `clientToken` | `string` | no |
| `codegenJobToCreate` | `StartCodegenJobData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entity` | `CodegenJob` | no |

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


## UpdateComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |
| `clientToken` | `string` | no |
| `updatedComponent` | `UpdateComponentData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entity` | `Component` | no |

## UpdateForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |
| `clientToken` | `string` | no |
| `updatedForm` | `UpdateFormData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entity` | `Form` | no |

## UpdateTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `id` | `string` | yes |
| `clientToken` | `string` | no |
| `updatedTheme` | `UpdateThemeData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entity` | `Theme` | no |

