# Amazon SageMaker Runtime

API version: 2017-05-13. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sagemaker-runtime/2017-05-13/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## InvokeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `Body` | `blob` | yes |
| `ContentType` | `string` | no |
| `Accept` | `string` | no |
| `CustomAttributes` | `string` | no |
| `TargetModel` | `string` | no |
| `TargetVariant` | `string` | no |
| `TargetContainerHostname` | `string` | no |
| `InferenceId` | `string` | no |
| `EnableExplanations` | `string` | no |
| `InferenceComponentName` | `string` | no |
| `SessionId` | `string` | no |
| `PrefixAwareId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `blob` | yes |
| `ContentType` | `string` | no |
| `InvokedProductionVariant` | `string` | no |
| `CustomAttributes` | `string` | no |
| `NewSessionId` | `string` | no |
| `ClosedSessionId` | `string` | no |

## InvokeEndpointAsync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `ContentType` | `string` | no |
| `Accept` | `string` | no |
| `CustomAttributes` | `string` | no |
| `InferenceId` | `string` | no |
| `InputLocation` | `string` | no |
| `S3OutputPathExtension` | `string` | no |
| `Filename` | `string` | no |
| `RequestTTLSeconds` | `integer` | no |
| `InvocationTimeoutSeconds` | `integer` | no |
| `Body` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceId` | `string` | no |
| `OutputLocation` | `string` | no |
| `FailureLocation` | `string` | no |

## InvokeEndpointWithResponseStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `Body` | `blob` | yes |
| `ContentType` | `string` | no |
| `Accept` | `string` | no |
| `CustomAttributes` | `string` | no |
| `TargetVariant` | `string` | no |
| `TargetContainerHostname` | `string` | no |
| `InferenceId` | `string` | no |
| `InferenceComponentName` | `string` | no |
| `SessionId` | `string` | no |
| `PrefixAwareId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `ResponseStream` | yes |
| `ContentType` | `string` | no |
| `InvokedProductionVariant` | `string` | no |
| `CustomAttributes` | `string` | no |

