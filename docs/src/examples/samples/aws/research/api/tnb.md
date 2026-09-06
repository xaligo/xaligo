# AWS Telco Network Builder

API version: 2008-10-21. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/tnb/2008-10-21/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelSolNetworkOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsLcmOpOccId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateSolFunctionPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `onboardingState` | `string` | yes |
| `operationalState` | `string` | yes |
| `tags` | `Map<string>` | no |
| `usageState` | `string` | yes |

## CreateSolNetworkInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsDescription` | `string` | no |
| `nsName` | `string` | yes |
| `nsdInfoId` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `nsInstanceName` | `string` | yes |
| `nsdInfoId` | `string` | yes |
| `tags` | `Map<string>` | no |

## CreateSolNetworkPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `nsdOnboardingState` | `string` | yes |
| `nsdOperationalState` | `string` | yes |
| `nsdUsageState` | `string` | yes |
| `tags` | `Map<string>` | no |

## DeleteSolFunctionPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vnfPkgId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSolNetworkInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSolNetworkPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsdInfoId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetSolFunctionInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vnfInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `instantiatedVnfInfo` | `GetSolVnfInfo` | no |
| `instantiationState` | `string` | yes |
| `metadata` | `GetSolFunctionInstanceMetadata` | yes |
| `nsInstanceId` | `string` | yes |
| `tags` | `Map<string>` | no |
| `vnfPkgId` | `string` | yes |
| `vnfProductName` | `string` | no |
| `vnfProvider` | `string` | no |
| `vnfdId` | `string` | yes |
| `vnfdVersion` | `string` | no |

## GetSolFunctionPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vnfPkgId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `metadata` | `GetSolFunctionPackageMetadata` | no |
| `onboardingState` | `string` | yes |
| `operationalState` | `string` | yes |
| `tags` | `Map<string>` | no |
| `usageState` | `string` | yes |
| `vnfProductName` | `string` | no |
| `vnfProvider` | `string` | no |
| `vnfdId` | `string` | no |
| `vnfdVersion` | `string` | no |

## GetSolFunctionPackageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accept` | `string` | yes |
| `vnfPkgId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `packageContent` | `blob` | no |

## GetSolFunctionPackageDescriptor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accept` | `string` | yes |
| `vnfPkgId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `vnfd` | `blob` | no |

## GetSolNetworkInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `lcmOpInfo` | `LcmOperationInfo` | no |
| `metadata` | `GetSolNetworkInstanceMetadata` | yes |
| `nsInstanceDescription` | `string` | yes |
| `nsInstanceName` | `string` | yes |
| `nsState` | `string` | no |
| `nsdId` | `string` | yes |
| `nsdInfoId` | `string` | yes |
| `tags` | `Map<string>` | no |

## GetSolNetworkOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsLcmOpOccId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `error` | `ProblemDetails` | no |
| `id` | `string` | no |
| `lcmOperationType` | `string` | no |
| `metadata` | `GetSolNetworkOperationMetadata` | no |
| `nsInstanceId` | `string` | no |
| `operationState` | `string` | no |
| `tags` | `Map<string>` | no |
| `tasks` | `List<GetSolNetworkOperationTaskDetails>` | no |
| `updateType` | `string` | no |

## GetSolNetworkPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsdInfoId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `metadata` | `GetSolNetworkPackageMetadata` | yes |
| `nsdId` | `string` | yes |
| `nsdName` | `string` | yes |
| `nsdOnboardingState` | `string` | yes |
| `nsdOperationalState` | `string` | yes |
| `nsdUsageState` | `string` | yes |
| `nsdVersion` | `string` | yes |
| `tags` | `Map<string>` | no |
| `vnfPkgIds` | `List<string>` | yes |

## GetSolNetworkPackageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accept` | `string` | yes |
| `nsdInfoId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `nsdContent` | `blob` | no |

## GetSolNetworkPackageDescriptor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsdInfoId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `nsd` | `blob` | no |

## InstantiateSolNetworkInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `additionalParamsForNs` | `Document` | no |
| `dryRun` | `boolean` | no |
| `nsInstanceId` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsLcmOpOccId` | `string` | yes |
| `tags` | `Map<string>` | no |

## ListSolFunctionInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `functionInstances` | `List<ListSolFunctionInstanceInfo>` | no |
| `nextToken` | `string` | no |

## ListSolFunctionPackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `functionPackages` | `List<ListSolFunctionPackageInfo>` | yes |
| `nextToken` | `string` | no |

## ListSolNetworkInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkInstances` | `List<ListSolNetworkInstanceInfo>` | no |
| `nextToken` | `string` | no |

## ListSolNetworkOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `nsInstanceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkOperations` | `List<ListSolNetworkOperationsInfo>` | no |
| `nextToken` | `string` | no |

## ListSolNetworkPackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkPackages` | `List<ListSolNetworkPackageInfo>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## PutSolFunctionPackageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `file` | `blob` | yes |
| `vnfPkgId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `metadata` | `PutSolFunctionPackageContentMetadata` | yes |
| `vnfProductName` | `string` | yes |
| `vnfProvider` | `string` | yes |
| `vnfdId` | `string` | yes |
| `vnfdVersion` | `string` | yes |

## PutSolNetworkPackageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `file` | `blob` | yes |
| `nsdInfoId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `metadata` | `PutSolNetworkPackageContentMetadata` | yes |
| `nsdId` | `string` | yes |
| `nsdName` | `string` | yes |
| `nsdVersion` | `string` | yes |
| `vnfPkgIds` | `List<string>` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateSolNetworkInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsInstanceId` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsLcmOpOccId` | `string` | no |
| `tags` | `Map<string>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSolFunctionPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationalState` | `string` | yes |
| `vnfPkgId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationalState` | `string` | yes |

## UpdateSolNetworkInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modifyVnfInfoData` | `UpdateSolNetworkModify` | no |
| `nsInstanceId` | `string` | yes |
| `tags` | `Map<string>` | no |
| `updateNs` | `UpdateSolNetworkServiceData` | no |
| `updateType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsLcmOpOccId` | `string` | no |
| `tags` | `Map<string>` | no |

## UpdateSolNetworkPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsdInfoId` | `string` | yes |
| `nsdOperationalState` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nsdOperationalState` | `string` | yes |

## ValidateSolFunctionPackageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `file` | `blob` | yes |
| `vnfPkgId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `metadata` | `ValidateSolFunctionPackageContentMetadata` | yes |
| `vnfProductName` | `string` | yes |
| `vnfProvider` | `string` | yes |
| `vnfdId` | `string` | yes |
| `vnfdVersion` | `string` | yes |

## ValidateSolNetworkPackageContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `file` | `blob` | yes |
| `nsdInfoId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `metadata` | `ValidateSolNetworkPackageContentMetadata` | yes |
| `nsdId` | `string` | yes |
| `nsdName` | `string` | yes |
| `nsdVersion` | `string` | yes |
| `vnfPkgIds` | `List<string>` | yes |

