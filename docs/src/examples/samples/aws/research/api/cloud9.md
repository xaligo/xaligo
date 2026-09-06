# AWS Cloud9

API version: 2017-09-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloud9/2017-09-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateEnvironmentEC2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `clientRequestToken` | `string` | no |
| `instanceType` | `string` | yes |
| `subnetId` | `string` | no |
| `imageId` | `string` | yes |
| `automaticStopTimeMinutes` | `integer` | no |
| `ownerArn` | `string` | no |
| `tags` | `List<Tag>` | no |
| `connectionType` | `string` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | no |

## CreateEnvironmentMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `userArn` | `string` | yes |
| `permissions` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membership` | `EnvironmentMember` | yes |

## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironmentMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `userArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeEnvironmentMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userArn` | `string` | no |
| `environmentId` | `string` | no |
| `permissions` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memberships` | `List<EnvironmentMember>` | no |
| `nextToken` | `string` | no |

## DescribeEnvironmentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `message` | `string` | yes |

## DescribeEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environments` | `List<Environment>` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `environmentIds` | `List<string>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `managedCredentialsAction` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEnvironmentMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `userArn` | `string` | yes |
| `permissions` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membership` | `EnvironmentMember` | no |

