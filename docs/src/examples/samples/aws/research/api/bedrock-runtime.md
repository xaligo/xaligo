# Amazon Bedrock Runtime

API version: 2023-09-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bedrock-runtime/2023-09-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ApplyGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailIdentifier` | `string` | yes |
| `guardrailVersion` | `string` | yes |
| `source` | `string` | yes |
| `content` | `List<GuardrailContentBlock>` | yes |
| `outputScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usage` | `GuardrailUsage` | yes |
| `action` | `string` | yes |
| `actionReason` | `string` | no |
| `outputs` | `List<GuardrailOutputContent>` | yes |
| `assessments` | `List<GuardrailAssessment>` | yes |
| `guardrailCoverage` | `GuardrailCoverage` | no |

## Converse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `messages` | `List<Message>` | no |
| `system` | `List<SystemContentBlock>` | no |
| `inferenceConfig` | `InferenceConfiguration` | no |
| `toolConfig` | `ToolConfiguration` | no |
| `guardrailConfig` | `GuardrailConfiguration` | no |
| `additionalModelRequestFields` | `Document` | no |
| `promptVariables` | `Map<PromptVariableValues>` | no |
| `additionalModelResponseFieldPaths` | `List<string>` | no |
| `requestMetadata` | `Map<string>` | no |
| `performanceConfig` | `PerformanceConfiguration` | no |
| `serviceTier` | `ServiceTier` | no |
| `outputConfig` | `OutputConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `output` | `ConverseOutput` | yes |
| `stopReason` | `string` | yes |
| `usage` | `TokenUsage` | yes |
| `metrics` | `ConverseMetrics` | yes |
| `additionalModelResponseFields` | `Document` | no |
| `trace` | `ConverseTrace` | no |
| `performanceConfig` | `PerformanceConfiguration` | no |
| `serviceTier` | `ServiceTier` | no |

## ConverseStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `messages` | `List<Message>` | no |
| `system` | `List<SystemContentBlock>` | no |
| `inferenceConfig` | `InferenceConfiguration` | no |
| `toolConfig` | `ToolConfiguration` | no |
| `guardrailConfig` | `GuardrailStreamConfiguration` | no |
| `additionalModelRequestFields` | `Document` | no |
| `promptVariables` | `Map<PromptVariableValues>` | no |
| `additionalModelResponseFieldPaths` | `List<string>` | no |
| `requestMetadata` | `Map<string>` | no |
| `performanceConfig` | `PerformanceConfiguration` | no |
| `serviceTier` | `ServiceTier` | no |
| `outputConfig` | `OutputConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stream` | `ConverseStreamOutput` | no |

## CountTokens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `input` | `CountTokensInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inputTokens` | `integer` | yes |

## GetAsyncInvoke

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationArn` | `string` | yes |
| `modelArn` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `status` | `string` | yes |
| `failureMessage` | `string` | no |
| `submitTime` | `timestamp` | yes |
| `lastModifiedTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `outputDataConfig` | `AsyncInvokeOutputDataConfig` | yes |

## InvokeGuardrailChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messages` | `List<GuardrailChecksMessage>` | yes |
| `checks` | `GuardrailChecksConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `GuardrailChecksResults` | yes |
| `usage` | `GuardrailChecksUsageResults` | yes |

## InvokeModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `blob` | no |
| `contentType` | `string` | no |
| `accept` | `string` | no |
| `modelId` | `string` | yes |
| `trace` | `string` | no |
| `guardrailIdentifier` | `string` | no |
| `guardrailVersion` | `string` | no |
| `performanceConfigLatency` | `string` | no |
| `serviceTier` | `string` | no |
| `requestMetadata` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `blob` | yes |
| `contentType` | `string` | yes |
| `performanceConfigLatency` | `string` | no |
| `serviceTier` | `string` | no |

## InvokeModelWithBidirectionalStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `body` | `InvokeModelWithBidirectionalStreamInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `InvokeModelWithBidirectionalStreamOutput` | yes |

## InvokeModelWithResponseStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `blob` | no |
| `contentType` | `string` | no |
| `accept` | `string` | no |
| `modelId` | `string` | yes |
| `trace` | `string` | no |
| `guardrailIdentifier` | `string` | no |
| `guardrailVersion` | `string` | no |
| `performanceConfigLatency` | `string` | no |
| `serviceTier` | `string` | no |
| `requestMetadata` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `ResponseStream` | yes |
| `contentType` | `string` | yes |
| `performanceConfigLatency` | `string` | no |
| `serviceTier` | `string` | no |

## ListAsyncInvokes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `submitTimeAfter` | `timestamp` | no |
| `submitTimeBefore` | `timestamp` | no |
| `statusEquals` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `asyncInvokeSummaries` | `List<AsyncInvokeSummary>` | no |

## StartAsyncInvoke

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientRequestToken` | `string` | no |
| `modelId` | `string` | yes |
| `modelInput` | `ModelInputPayload` | yes |
| `outputDataConfig` | `AsyncInvokeOutputDataConfig` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationArn` | `string` | yes |

