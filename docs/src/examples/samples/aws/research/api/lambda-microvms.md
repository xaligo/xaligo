# Lambda MicroVMs

API version: 2025-09-09. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lambda-microvms/2025-09-09/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateMicrovmAuthToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `microvmIdentifier` | `string` | yes |
| `expirationInMinutes` | `integer` | yes |
| `allowedPorts` | `List<PortSpecification>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authToken` | `Map<string>` | yes |

## CreateMicrovmImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baseImageArn` | `string` | yes |
| `baseImageVersion` | `string` | no |
| `buildRoleArn` | `string` | yes |
| `description` | `string` | no |
| `codeArtifact` | `CodeArtifact` | yes |
| `logging` | `Logging` | no |
| `egressNetworkConnectors` | `List<string>` | no |
| `cpuConfigurations` | `List<CpuConfiguration>` | no |
| `resources` | `List<Resources>` | no |
| `additionalOsCapabilities` | `List<string>` | no |
| `hooks` | `Hooks` | no |
| `environmentVariables` | `Map<string>` | no |
| `name` | `string` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageArn` | `string` | yes |
| `name` | `string` | yes |
| `state` | `string` | yes |
| `latestActiveImageVersion` | `string` | no |
| `latestFailedImageVersion` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `baseImageArn` | `string` | yes |
| `baseImageVersion` | `string` | no |
| `buildRoleArn` | `string` | yes |
| `description` | `string` | no |
| `codeArtifact` | `CodeArtifact` | yes |
| `logging` | `Logging` | no |
| `egressNetworkConnectors` | `List<string>` | no |
| `cpuConfigurations` | `List<CpuConfiguration>` | no |
| `resources` | `List<Resources>` | no |
| `additionalOsCapabilities` | `List<string>` | no |
| `hooks` | `Hooks` | no |
| `environmentVariables` | `Map<string>` | no |
| `tags` | `Map<string>` | no |
| `updatedAt` | `timestamp` | no |
| `imageVersion` | `string` | yes |

## CreateMicrovmShellAuthToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `microvmIdentifier` | `string` | yes |
| `expirationInMinutes` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authToken` | `Map<string>` | yes |

## DeleteMicrovmImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIdentifier` | `string` | yes |
| `state` | `string` | yes |

## DeleteMicrovmImageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIdentifier` | `string` | yes |
| `imageVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIdentifier` | `string` | yes |
| `imageVersion` | `string` | yes |
| `state` | `string` | yes |

## GetMicrovm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `microvmIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `microvmId` | `string` | yes |
| `state` | `string` | yes |
| `endpoint` | `string` | yes |
| `imageArn` | `string` | yes |
| `imageVersion` | `string` | yes |
| `executionRoleArn` | `string` | no |
| `idlePolicy` | `IdlePolicy` | no |
| `maximumDurationInSeconds` | `integer` | yes |
| `startedAt` | `timestamp` | yes |
| `terminatedAt` | `timestamp` | no |
| `stateReason` | `string` | no |
| `ingressNetworkConnectors` | `List<string>` | no |
| `egressNetworkConnectors` | `List<string>` | no |

## GetMicrovmImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageArn` | `string` | yes |
| `name` | `string` | yes |
| `state` | `string` | yes |
| `latestActiveImageVersion` | `string` | no |
| `latestFailedImageVersion` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |
| `updatedAt` | `timestamp` | no |

## GetMicrovmImageBuild

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIdentifier` | `string` | yes |
| `imageVersion` | `string` | yes |
| `buildId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageArn` | `string` | yes |
| `imageVersion` | `string` | yes |
| `buildId` | `string` | yes |
| `buildState` | `string` | yes |
| `architecture` | `string` | yes |
| `chipset` | `string` | yes |
| `chipsetGeneration` | `string` | yes |
| `stateReason` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `snapshotBuild` | `SnapshotBuild` | no |

## GetMicrovmImageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIdentifier` | `string` | yes |
| `imageVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baseImageArn` | `string` | yes |
| `baseImageVersion` | `string` | no |
| `buildRoleArn` | `string` | yes |
| `description` | `string` | no |
| `codeArtifact` | `CodeArtifact` | yes |
| `logging` | `Logging` | no |
| `egressNetworkConnectors` | `List<string>` | no |
| `cpuConfigurations` | `List<CpuConfiguration>` | no |
| `resources` | `List<Resources>` | no |
| `additionalOsCapabilities` | `List<string>` | no |
| `hooks` | `Hooks` | no |
| `environmentVariables` | `Map<string>` | no |
| `imageArn` | `string` | yes |
| `imageVersion` | `string` | yes |
| `state` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | no |
| `stateReason` | `string` | no |
| `tags` | `Map<string>` | no |

## ListManagedMicrovmImageVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `imageIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ManagedMicrovmImageVersion>` | yes |

## ListManagedMicrovmImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ManagedMicrovmImageSummary>` | yes |

## ListMicrovmImageBuilds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `imageIdentifier` | `string` | yes |
| `imageVersion` | `string` | yes |
| `architecture` | `string` | no |
| `chipset` | `string` | no |
| `chipsetGeneration` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<MicrovmImageBuildSummary>` | yes |

## ListMicrovmImageVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `imageIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<MicrovmImageVersionSummary>` | yes |

## ListMicrovmImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `nameFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<MicrovmImageSummary>` | yes |

## ListMicrovms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `imageIdentifier` | `string` | no |
| `imageVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<MicrovmItem>` | yes |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ResumeMicrovm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `microvmIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RunMicrovm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingressNetworkConnectors` | `List<string>` | no |
| `egressNetworkConnectors` | `List<string>` | no |
| `imageIdentifier` | `string` | yes |
| `imageVersion` | `string` | no |
| `executionRoleArn` | `string` | no |
| `idlePolicy` | `IdlePolicy` | no |
| `logging` | `Logging` | no |
| `runHookPayload` | `string` | no |
| `maximumDurationInSeconds` | `integer` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `microvmId` | `string` | yes |
| `state` | `string` | yes |
| `endpoint` | `string` | yes |
| `imageArn` | `string` | yes |
| `imageVersion` | `string` | yes |
| `executionRoleArn` | `string` | no |
| `idlePolicy` | `IdlePolicy` | no |
| `maximumDurationInSeconds` | `integer` | yes |
| `startedAt` | `timestamp` | yes |
| `terminatedAt` | `timestamp` | no |
| `stateReason` | `string` | no |
| `ingressNetworkConnectors` | `List<string>` | no |
| `egressNetworkConnectors` | `List<string>` | no |

## SuspendMicrovm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `microvmIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateMicrovm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `microvmIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMicrovmImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baseImageArn` | `string` | yes |
| `baseImageVersion` | `string` | no |
| `buildRoleArn` | `string` | yes |
| `description` | `string` | no |
| `codeArtifact` | `CodeArtifact` | yes |
| `logging` | `Logging` | no |
| `egressNetworkConnectors` | `List<string>` | no |
| `cpuConfigurations` | `List<CpuConfiguration>` | no |
| `resources` | `List<Resources>` | no |
| `additionalOsCapabilities` | `List<string>` | no |
| `hooks` | `Hooks` | no |
| `environmentVariables` | `Map<string>` | no |
| `imageIdentifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageArn` | `string` | yes |
| `name` | `string` | yes |
| `state` | `string` | yes |
| `latestActiveImageVersion` | `string` | no |
| `latestFailedImageVersion` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `baseImageArn` | `string` | yes |
| `baseImageVersion` | `string` | no |
| `buildRoleArn` | `string` | yes |
| `description` | `string` | no |
| `codeArtifact` | `CodeArtifact` | yes |
| `logging` | `Logging` | no |
| `egressNetworkConnectors` | `List<string>` | no |
| `cpuConfigurations` | `List<CpuConfiguration>` | no |
| `resources` | `List<Resources>` | no |
| `additionalOsCapabilities` | `List<string>` | no |
| `hooks` | `Hooks` | no |
| `environmentVariables` | `Map<string>` | no |
| `updatedAt` | `timestamp` | yes |
| `imageVersion` | `string` | yes |

## UpdateMicrovmImageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIdentifier` | `string` | yes |
| `imageVersion` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baseImageArn` | `string` | yes |
| `baseImageVersion` | `string` | no |
| `buildRoleArn` | `string` | yes |
| `description` | `string` | no |
| `codeArtifact` | `CodeArtifact` | yes |
| `logging` | `Logging` | no |
| `egressNetworkConnectors` | `List<string>` | no |
| `cpuConfigurations` | `List<CpuConfiguration>` | no |
| `resources` | `List<Resources>` | no |
| `additionalOsCapabilities` | `List<string>` | no |
| `hooks` | `Hooks` | no |
| `environmentVariables` | `Map<string>` | no |
| `imageArn` | `string` | yes |
| `imageVersion` | `string` | yes |
| `state` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | no |
| `stateReason` | `string` | no |
| `tags` | `Map<string>` | no |

