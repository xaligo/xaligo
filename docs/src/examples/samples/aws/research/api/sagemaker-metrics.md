# Amazon SageMaker Metrics Service

API version: 2022-09-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sagemaker-metrics/2022-09-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricQueries` | `List<MetricQuery>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricQueryResults` | `List<MetricQueryResult>` | no |

## BatchPutMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentName` | `string` | yes |
| `MetricData` | `List<RawMetricData>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchPutMetricsError>` | no |

