# Amazon Managed Blockchain

API version: 2018-09-24. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/managedblockchain/2018-09-24/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAccessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | yes |
| `AccessorType` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessorId` | `string` | no |
| `BillingToken` | `string` | no |
| `NetworkType` | `string` | no |

## CreateMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | yes |
| `InvitationId` | `string` | yes |
| `NetworkId` | `string` | yes |
| `MemberConfiguration` | `MemberConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberId` | `string` | no |

## CreateNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Framework` | `string` | yes |
| `FrameworkVersion` | `string` | yes |
| `FrameworkConfiguration` | `NetworkFrameworkConfiguration` | no |
| `VotingPolicy` | `VotingPolicy` | yes |
| `MemberConfiguration` | `MemberConfiguration` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | no |
| `MemberId` | `string` | no |

## CreateNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | yes |
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | no |
| `NodeConfiguration` | `NodeConfiguration` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NodeId` | `string` | no |

## CreateProposal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | yes |
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | yes |
| `Actions` | `ProposalActions` | yes |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProposalId` | `string` | no |

## DeleteAccessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | no |
| `NodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accessor` | `Accessor` | no |

## GetMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Member` | `Member` | no |

## GetNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Network` | `Network` | no |

## GetNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | no |
| `NodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Node` | `Node` | no |

## GetProposal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `ProposalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Proposal` | `Proposal` | no |

## ListAccessors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accessors` | `List<AccessorSummary>` | no |
| `NextToken` | `string` | no |

## ListInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Invitations` | `List<Invitation>` | no |
| `NextToken` | `string` | no |

## ListMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `IsOwned` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Members` | `List<MemberSummary>` | no |
| `NextToken` | `string` | no |

## ListNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Framework` | `string` | no |
| `Status` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Networks` | `List<NetworkSummary>` | no |
| `NextToken` | `string` | no |

## ListNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | no |
| `Status` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Nodes` | `List<NodeSummary>` | no |
| `NextToken` | `string` | no |

## ListProposalVotes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `ProposalId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProposalVotes` | `List<VoteSummary>` | no |
| `NextToken` | `string` | no |

## ListProposals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Proposals` | `List<ProposalSummary>` | no |
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

## RejectInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvitationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | yes |
| `LogPublishingConfiguration` | `MemberLogPublishingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `MemberId` | `string` | no |
| `NodeId` | `string` | yes |
| `LogPublishingConfiguration` | `NodeLogPublishingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## VoteOnProposal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |
| `ProposalId` | `string` | yes |
| `VoterMemberId` | `string` | yes |
| `Vote` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


