# Amazon S3 on Outposts

API version: 2017-07-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/s3outposts/2017-07-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostId` | `string` | yes |
| `SubnetId` | `string` | yes |
| `SecurityGroupId` | `string` | yes |
| `AccessType` | `string` | no |
| `CustomerOwnedIpv4Pool` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | no |

## DeleteEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointId` | `string` | yes |
| `OutpostId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<Endpoint>` | no |
| `NextToken` | `string` | no |

## ListOutpostsWithS3

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Outposts` | `List<Outpost>` | no |
| `NextToken` | `string` | no |

## ListSharedEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `OutpostId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<Endpoint>` | no |
| `NextToken` | `string` | no |

