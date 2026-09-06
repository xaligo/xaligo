# Private CA Connector for SCEP

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pca-connector-scep/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateChallenge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Challenge` | `Challenge` | no |

## CreateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `MobileDeviceManagement` | `MobileDeviceManagement` | no |
| `VpcEndpointId` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | no |

## DeleteChallenge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChallengeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetChallengeMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChallengeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChallengeMetadata` | `ChallengeMetadata` | no |

## GetChallengePassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChallengeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Password` | `string` | no |

## GetConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connector` | `Connector` | no |

## ListChallengeMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ConnectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Challenges` | `List<ChallengeMetadataSummary>` | no |
| `NextToken` | `string` | no |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connectors` | `List<ConnectorSummary>` | no |
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


