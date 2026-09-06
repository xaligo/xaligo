# AmplifyBackend

API version: 2020-08-11. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/amplifybackend/2020-08-11/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CloneBackend

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `TargetEnvironmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## CreateBackend

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `AppName` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `ResourceConfig` | no |
| `ResourceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## CreateBackendAPI

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `BackendAPIResourceConfig` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## CreateBackendAuth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `CreateBackendAuthResourceConfig` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## CreateBackendConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendManagerAppId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `JobId` | `string` | no |
| `Status` | `string` | no |

## CreateBackendStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `CreateBackendStorageResourceConfig` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `JobId` | `string` | no |
| `Status` | `string` | no |

## CreateToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `ChallengeCode` | `string` | no |
| `SessionId` | `string` | no |
| `Ttl` | `string` | no |

## DeleteBackend

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## DeleteBackendAPI

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `BackendAPIResourceConfig` | no |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## DeleteBackendAuth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## DeleteBackendStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceName` | `string` | yes |
| `ServiceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `JobId` | `string` | no |
| `Status` | `string` | no |

## DeleteToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IsSuccess` | `boolean` | no |

## GenerateBackendAPIModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## GetBackend

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmplifyFeatureFlags` | `string` | no |
| `AmplifyMetaConfig` | `string` | no |
| `AppId` | `string` | no |
| `AppName` | `string` | no |
| `BackendEnvironmentList` | `List<string>` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |

## GetBackendAPI

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `BackendAPIResourceConfig` | no |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `ResourceConfig` | `BackendAPIResourceConfig` | no |
| `ResourceName` | `string` | no |

## GetBackendAPIModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Models` | `string` | no |
| `Status` | `string` | no |
| `ModelIntrospectionSchema` | `string` | no |

## GetBackendAuth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `ResourceConfig` | `CreateBackendAuthResourceConfig` | no |
| `ResourceName` | `string` | no |

## GetBackendJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `CreateTime` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |
| `UpdateTime` | `string` | no |

## GetBackendStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `ResourceConfig` | `GetBackendStorageResourceConfig` | no |
| `ResourceName` | `string` | no |

## GetToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `ChallengeCode` | `string` | no |
| `SessionId` | `string` | no |
| `Ttl` | `string` | no |

## ImportBackendAuth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `IdentityPoolId` | `string` | no |
| `NativeClientId` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `WebClientId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## ImportBackendStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `BucketName` | `string` | no |
| `ServiceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `JobId` | `string` | no |
| `Status` | `string` | no |

## ListBackendJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `JobId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<BackendJobRespObj>` | no |
| `NextToken` | `string` | no |

## ListS3Buckets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Buckets` | `List<S3BucketInfo>` | no |
| `NextToken` | `string` | no |

## RemoveAllBackends

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `CleanAmplifyApp` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## RemoveBackendConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Error` | `string` | no |

## UpdateBackendAPI

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `BackendAPIResourceConfig` | no |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## UpdateBackendAuth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `UpdateBackendAuthResourceConfig` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |

## UpdateBackendConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `LoginAuthConfig` | `LoginAuthConfigReqObj` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendManagerAppId` | `string` | no |
| `Error` | `string` | no |
| `LoginAuthConfig` | `LoginAuthConfigReqObj` | no |

## UpdateBackendJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `JobId` | `string` | yes |
| `Operation` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `CreateTime` | `string` | no |
| `Error` | `string` | no |
| `JobId` | `string` | no |
| `Operation` | `string` | no |
| `Status` | `string` | no |
| `UpdateTime` | `string` | no |

## UpdateBackendStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | yes |
| `BackendEnvironmentName` | `string` | yes |
| `ResourceConfig` | `UpdateBackendStorageResourceConfig` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppId` | `string` | no |
| `BackendEnvironmentName` | `string` | no |
| `JobId` | `string` | no |
| `Status` | `string` | no |

