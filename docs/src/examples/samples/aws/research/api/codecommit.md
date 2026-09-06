# AWS CodeCommit

API version: 2015-04-13. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codecommit/2015-04-13/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateApprovalRuleTemplateWithRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchAssociateApprovalRuleTemplateWithRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |
| `repositoryNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associatedRepositoryNames` | `List<string>` | yes |
| `errors` | `List<BatchAssociateApprovalRuleTemplateWithRepositoriesError>` | yes |

## BatchDescribeMergeConflicts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `mergeOption` | `string` | yes |
| `maxMergeHunks` | `integer` | no |
| `maxConflictFiles` | `integer` | no |
| `filePaths` | `List<string>` | no |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `conflicts` | `List<Conflict>` | yes |
| `nextToken` | `string` | no |
| `errors` | `List<BatchDescribeMergeConflictsError>` | no |
| `destinationCommitId` | `string` | yes |
| `sourceCommitId` | `string` | yes |
| `baseCommitId` | `string` | no |

## BatchDisassociateApprovalRuleTemplateFromRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |
| `repositoryNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `disassociatedRepositoryNames` | `List<string>` | yes |
| `errors` | `List<BatchDisassociateApprovalRuleTemplateFromRepositoriesError>` | yes |

## BatchGetCommits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitIds` | `List<string>` | yes |
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commits` | `List<Commit>` | no |
| `errors` | `List<BatchGetCommitsError>` | no |

## BatchGetRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositories` | `List<RepositoryMetadata>` | no |
| `repositoriesNotFound` | `List<string>` | no |
| `errors` | `List<BatchGetRepositoriesError>` | no |

## CreateApprovalRuleTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |
| `approvalRuleTemplateContent` | `string` | yes |
| `approvalRuleTemplateDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplate` | `ApprovalRuleTemplate` | yes |

## CreateBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `branchName` | `string` | yes |
| `commitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateCommit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `branchName` | `string` | yes |
| `parentCommitId` | `string` | no |
| `authorName` | `string` | no |
| `email` | `string` | no |
| `commitMessage` | `string` | no |
| `keepEmptyFolders` | `boolean` | no |
| `putFiles` | `List<PutFileEntry>` | no |
| `deleteFiles` | `List<DeleteFileEntry>` | no |
| `setFileModes` | `List<SetFileModeEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | no |
| `treeId` | `string` | no |
| `filesAdded` | `List<FileMetadata>` | no |
| `filesUpdated` | `List<FileMetadata>` | no |
| `filesDeleted` | `List<FileMetadata>` | no |

## CreatePullRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `title` | `string` | yes |
| `description` | `string` | no |
| `targets` | `List<Target>` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequest` | `PullRequest` | yes |

## CreatePullRequestApprovalRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `approvalRuleName` | `string` | yes |
| `approvalRuleContent` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRule` | `ApprovalRule` | yes |

## CreateRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `repositoryDescription` | `string` | no |
| `tags` | `Map<string>` | no |
| `kmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryMetadata` | `RepositoryMetadata` | no |

## CreateUnreferencedMergeCommit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `mergeOption` | `string` | yes |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |
| `authorName` | `string` | no |
| `email` | `string` | no |
| `commitMessage` | `string` | no |
| `keepEmptyFolders` | `boolean` | no |
| `conflictResolution` | `ConflictResolution` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | no |
| `treeId` | `string` | no |

## DeleteApprovalRuleTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateId` | `string` | yes |

## DeleteBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `branchName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deletedBranch` | `BranchInfo` | no |

## DeleteCommentContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `comment` | `Comment` | no |

## DeleteFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `branchName` | `string` | yes |
| `filePath` | `string` | yes |
| `parentCommitId` | `string` | yes |
| `keepEmptyFolders` | `boolean` | no |
| `commitMessage` | `string` | no |
| `name` | `string` | no |
| `email` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | yes |
| `blobId` | `string` | yes |
| `treeId` | `string` | yes |
| `filePath` | `string` | yes |

## DeletePullRequestApprovalRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `approvalRuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleId` | `string` | yes |

## DeleteRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryId` | `string` | no |

## DescribeMergeConflicts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `mergeOption` | `string` | yes |
| `maxMergeHunks` | `integer` | no |
| `filePath` | `string` | yes |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `conflictMetadata` | `ConflictMetadata` | yes |
| `mergeHunks` | `List<MergeHunk>` | yes |
| `nextToken` | `string` | no |
| `destinationCommitId` | `string` | yes |
| `sourceCommitId` | `string` | yes |
| `baseCommitId` | `string` | no |

## DescribePullRequestEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `pullRequestEventType` | `string` | no |
| `actorArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestEvents` | `List<PullRequestEvent>` | yes |
| `nextToken` | `string` | no |

## DisassociateApprovalRuleTemplateFromRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EvaluatePullRequestApprovalRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `revisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluation` | `Evaluation` | yes |

## GetApprovalRuleTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplate` | `ApprovalRuleTemplate` | yes |

## GetBlob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `blobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `content` | `blob` | yes |

## GetBlobDifferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `afterBlobId` | `string` | yes |
| `beforeBlobId` | `string` | no |
| `contextLines` | `integer` | no |
| `ignoreWhitespace` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `hunks` | `List<DiffHunk>` | yes |
| `isBinary` | `boolean` | yes |
| `beforeBlobSize` | `long` | no |
| `afterBlobSize` | `long` | yes |
| `NextToken` | `string` | no |

## GetBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | no |
| `branchName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `BranchInfo` | no |

## GetComment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `comment` | `Comment` | no |

## GetCommentReactions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentId` | `string` | yes |
| `reactionUserArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reactionsForComment` | `List<ReactionForComment>` | yes |
| `nextToken` | `string` | no |

## GetCommentsForComparedCommit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `beforeCommitId` | `string` | no |
| `afterCommitId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentsForComparedCommitData` | `List<CommentsForComparedCommit>` | no |
| `nextToken` | `string` | no |

## GetCommentsForPullRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `repositoryName` | `string` | no |
| `beforeCommitId` | `string` | no |
| `afterCommitId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentsForPullRequestData` | `List<CommentsForPullRequest>` | no |
| `nextToken` | `string` | no |

## GetCommit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `commitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commit` | `Commit` | yes |

## GetDifferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `beforeCommitSpecifier` | `string` | no |
| `afterCommitSpecifier` | `string` | yes |
| `beforePath` | `string` | no |
| `afterPath` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `differences` | `List<Difference>` | no |
| `NextToken` | `string` | no |

## GetFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `commitSpecifier` | `string` | no |
| `filePath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | yes |
| `blobId` | `string` | yes |
| `filePath` | `string` | yes |
| `fileMode` | `string` | yes |
| `fileSize` | `long` | yes |
| `fileContent` | `blob` | yes |

## GetFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `commitSpecifier` | `string` | no |
| `folderPath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | yes |
| `folderPath` | `string` | yes |
| `treeId` | `string` | no |
| `subFolders` | `List<Folder>` | no |
| `files` | `List<File>` | no |
| `symbolicLinks` | `List<SymbolicLink>` | no |
| `subModules` | `List<SubModule>` | no |

## GetMergeCommit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceCommitId` | `string` | no |
| `destinationCommitId` | `string` | no |
| `baseCommitId` | `string` | no |
| `mergedCommitId` | `string` | no |

## GetMergeConflicts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `mergeOption` | `string` | yes |
| `conflictDetailLevel` | `string` | no |
| `maxConflictFiles` | `integer` | no |
| `conflictResolutionStrategy` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mergeable` | `boolean` | yes |
| `destinationCommitId` | `string` | yes |
| `sourceCommitId` | `string` | yes |
| `baseCommitId` | `string` | no |
| `conflictMetadataList` | `List<ConflictMetadata>` | yes |
| `nextToken` | `string` | no |

## GetMergeOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mergeOptions` | `List<string>` | yes |
| `sourceCommitId` | `string` | yes |
| `destinationCommitId` | `string` | yes |
| `baseCommitId` | `string` | yes |

## GetPullRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequest` | `PullRequest` | yes |

## GetPullRequestApprovalStates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `revisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvals` | `List<Approval>` | no |

## GetPullRequestOverrideState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `revisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `overridden` | `boolean` | no |
| `overrider` | `string` | no |

## GetRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryMetadata` | `RepositoryMetadata` | no |

## GetRepositoryTriggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationId` | `string` | no |
| `triggers` | `List<RepositoryTrigger>` | no |

## ListApprovalRuleTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateNames` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListAssociatedApprovalRuleTemplatesForRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateNames` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListBranches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branches` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListFileCommitHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `commitSpecifier` | `string` | no |
| `filePath` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `revisionDag` | `List<FileVersion>` | yes |
| `nextToken` | `string` | no |

## ListPullRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `authorArn` | `string` | no |
| `pullRequestStatus` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestIds` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `order` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositories` | `List<RepositoryNameIdPair>` | no |
| `nextToken` | `string` | no |

## ListRepositoriesForApprovalRuleTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryNames` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |
| `nextToken` | `string` | no |

## MergeBranchesByFastForward

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `targetBranch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | no |
| `treeId` | `string` | no |

## MergeBranchesBySquash

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `targetBranch` | `string` | no |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |
| `authorName` | `string` | no |
| `email` | `string` | no |
| `commitMessage` | `string` | no |
| `keepEmptyFolders` | `boolean` | no |
| `conflictResolution` | `ConflictResolution` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | no |
| `treeId` | `string` | no |

## MergeBranchesByThreeWay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `sourceCommitSpecifier` | `string` | yes |
| `destinationCommitSpecifier` | `string` | yes |
| `targetBranch` | `string` | no |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |
| `authorName` | `string` | no |
| `email` | `string` | no |
| `commitMessage` | `string` | no |
| `keepEmptyFolders` | `boolean` | no |
| `conflictResolution` | `ConflictResolution` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | no |
| `treeId` | `string` | no |

## MergePullRequestByFastForward

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `repositoryName` | `string` | yes |
| `sourceCommitId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequest` | `PullRequest` | no |

## MergePullRequestBySquash

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `repositoryName` | `string` | yes |
| `sourceCommitId` | `string` | no |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |
| `commitMessage` | `string` | no |
| `authorName` | `string` | no |
| `email` | `string` | no |
| `keepEmptyFolders` | `boolean` | no |
| `conflictResolution` | `ConflictResolution` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequest` | `PullRequest` | no |

## MergePullRequestByThreeWay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `repositoryName` | `string` | yes |
| `sourceCommitId` | `string` | no |
| `conflictDetailLevel` | `string` | no |
| `conflictResolutionStrategy` | `string` | no |
| `commitMessage` | `string` | no |
| `authorName` | `string` | no |
| `email` | `string` | no |
| `keepEmptyFolders` | `boolean` | no |
| `conflictResolution` | `ConflictResolution` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequest` | `PullRequest` | no |

## OverridePullRequestApprovalRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `revisionId` | `string` | yes |
| `overrideStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PostCommentForComparedCommit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `beforeCommitId` | `string` | no |
| `afterCommitId` | `string` | yes |
| `location` | `Location` | no |
| `content` | `string` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | no |
| `beforeCommitId` | `string` | no |
| `afterCommitId` | `string` | no |
| `beforeBlobId` | `string` | no |
| `afterBlobId` | `string` | no |
| `location` | `Location` | no |
| `comment` | `Comment` | no |

## PostCommentForPullRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `repositoryName` | `string` | yes |
| `beforeCommitId` | `string` | yes |
| `afterCommitId` | `string` | yes |
| `location` | `Location` | no |
| `content` | `string` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | no |
| `pullRequestId` | `string` | no |
| `beforeCommitId` | `string` | no |
| `afterCommitId` | `string` | no |
| `beforeBlobId` | `string` | no |
| `afterBlobId` | `string` | no |
| `location` | `Location` | no |
| `comment` | `Comment` | no |

## PostCommentReply

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inReplyTo` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `content` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `comment` | `Comment` | no |

## PutCommentReaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentId` | `string` | yes |
| `reactionValue` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `branchName` | `string` | yes |
| `fileContent` | `blob` | yes |
| `filePath` | `string` | yes |
| `fileMode` | `string` | no |
| `parentCommitId` | `string` | no |
| `commitMessage` | `string` | no |
| `name` | `string` | no |
| `email` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commitId` | `string` | yes |
| `blobId` | `string` | yes |
| `treeId` | `string` | yes |

## PutRepositoryTriggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `triggers` | `List<RepositoryTrigger>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestRepositoryTriggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `triggers` | `List<RepositoryTrigger>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulExecutions` | `List<string>` | no |
| `failedExecutions` | `List<RepositoryTriggerExecutionFailure>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApprovalRuleTemplateContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |
| `newRuleContent` | `string` | yes |
| `existingRuleContentSha256` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplate` | `ApprovalRuleTemplate` | yes |

## UpdateApprovalRuleTemplateDescription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplateName` | `string` | yes |
| `approvalRuleTemplateDescription` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplate` | `ApprovalRuleTemplate` | yes |

## UpdateApprovalRuleTemplateName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `oldApprovalRuleTemplateName` | `string` | yes |
| `newApprovalRuleTemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRuleTemplate` | `ApprovalRuleTemplate` | yes |

## UpdateComment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commentId` | `string` | yes |
| `content` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `comment` | `Comment` | no |

## UpdateDefaultBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `defaultBranchName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePullRequestApprovalRuleContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `approvalRuleName` | `string` | yes |
| `existingRuleContentSha256` | `string` | no |
| `newRuleContent` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalRule` | `ApprovalRule` | yes |

## UpdatePullRequestApprovalState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `revisionId` | `string` | yes |
| `approvalState` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePullRequestDescription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequest` | `PullRequest` | yes |

## UpdatePullRequestStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `pullRequestStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequest` | `PullRequest` | yes |

## UpdatePullRequestTitle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequestId` | `string` | yes |
| `title` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullRequest` | `PullRequest` | yes |

## UpdateRepositoryDescription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `repositoryDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRepositoryEncryptionKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `kmsKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryId` | `string` | no |
| `kmsKeyId` | `string` | no |
| `originalKmsKeyId` | `string` | no |

## UpdateRepositoryName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `oldName` | `string` | yes |
| `newName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


