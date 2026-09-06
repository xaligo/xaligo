# Amazon EKS Auth

API version: 2023-11-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/eks-auth/2023-11-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssumeRoleForPodIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `token` | `string` | yes |
| `eksNodeName` | `string` | no |
| `instanceId` | `string` | no |
| `zone` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subject` | `Subject` | yes |
| `audience` | `string` | yes |
| `podIdentityAssociation` | `PodIdentityAssociation` | yes |
| `assumedRoleUser` | `AssumedRoleUser` | yes |
| `credentials` | `Credentials` | yes |

