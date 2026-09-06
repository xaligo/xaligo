# Route 53 Profiles

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/route53profiles/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ProfileId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileAssociation` | `ProfileAssociation` | no |

## AssociateResourceToProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ProfileId` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `ResourceProperties` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileResourceAssociation` | `ProfileResourceAssociation` | no |

## CreateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | yes |
| `Name` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Profile` | `Profile` | no |

## DeleteProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Profile` | `Profile` | no |

## DisassociateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileAssociation` | `ProfileAssociation` | no |

## DisassociateResourceFromProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileResourceAssociation` | `ProfileResourceAssociation` | no |

## GetProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Profile` | `Profile` | no |

## GetProfileAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileAssociation` | `ProfileAssociation` | no |

## GetProfileResourceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileResourceAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileResourceAssociation` | `ProfileResourceAssociation` | no |

## ListProfileAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ProfileId` | `string` | no |
| `ResourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ProfileAssociations` | `List<ProfileAssociation>` | no |

## ListProfileResourceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ProfileId` | `string` | yes |
| `ResourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ProfileResourceAssociations` | `List<ProfileResourceAssociation>` | no |

## ListProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ProfileSummaries` | `List<ProfileSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | yes |

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


## UpdateProfileResourceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `ProfileResourceAssociationId` | `string` | yes |
| `ResourceProperties` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileResourceAssociation` | `ProfileResourceAssociation` | no |

