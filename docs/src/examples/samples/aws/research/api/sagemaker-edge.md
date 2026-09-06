# Amazon Sagemaker Edge Manager

API version: 2020-09-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sagemaker-edge/2020-09-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceName` | `string` | yes |
| `DeviceFleetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Deployments` | `List<EdgeDeployment>` | no |

## GetDeviceRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceName` | `string` | yes |
| `DeviceFleetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceRegistration` | `string` | no |
| `CacheTTL` | `string` | no |

## SendHeartbeat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentMetrics` | `List<EdgeMetric>` | no |
| `Models` | `List<Model>` | no |
| `AgentVersion` | `string` | yes |
| `DeviceName` | `string` | yes |
| `DeviceFleetName` | `string` | yes |
| `DeploymentResult` | `DeploymentResult` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


