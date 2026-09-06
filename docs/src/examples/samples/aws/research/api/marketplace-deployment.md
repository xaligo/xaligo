# AWS Marketplace Deployment Service

API version: 2023-01-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/marketplace-deployment/2023-01-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutDeploymentParameter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `catalog` | `string` | yes |
| `clientToken` | `string` | no |
| `deploymentParameter` | `DeploymentParameterInput` | yes |
| `expirationDate` | `timestamp` | no |
| `productId` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementId` | `string` | yes |
| `deploymentParameterId` | `string` | yes |
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | no |

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


