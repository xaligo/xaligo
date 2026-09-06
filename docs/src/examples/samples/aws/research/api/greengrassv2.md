# AWS IoT Greengrass V2

API version: 2020-11-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/greengrassv2/2020-11-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateServiceRoleToAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associatedAt` | `string` | no |

## BatchAssociateClientDeviceWithCoreDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entries` | `List<AssociateClientDeviceWithCoreDeviceEntry>` | no |
| `coreDeviceThingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorEntries` | `List<AssociateClientDeviceWithCoreDeviceErrorEntry>` | no |

## BatchDisassociateClientDeviceFromCoreDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entries` | `List<DisassociateClientDeviceFromCoreDeviceEntry>` | no |
| `coreDeviceThingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errorEntries` | `List<DisassociateClientDeviceFromCoreDeviceErrorEntry>` | no |

## CancelDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## CreateComponentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inlineRecipe` | `blob` | no |
| `lambdaFunction` | `LambdaFunctionRecipeSource` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `componentName` | `string` | yes |
| `componentVersion` | `string` | yes |
| `creationTimestamp` | `timestamp` | yes |
| `status` | `CloudComponentStatus` | yes |

## CreateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetArn` | `string` | yes |
| `deploymentName` | `string` | no |
| `components` | `Map<ComponentDeploymentSpecification>` | no |
| `iotJobConfiguration` | `DeploymentIoTJobConfiguration` | no |
| `deploymentPolicies` | `DeploymentPolicies` | no |
| `parentTargetArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `iotJobId` | `string` | no |
| `iotJobArn` | `string` | no |

## DeleteComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCoreDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `coreDeviceThingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `componentName` | `string` | no |
| `componentVersion` | `string` | no |
| `creationTimestamp` | `timestamp` | no |
| `publisher` | `string` | no |
| `description` | `string` | no |
| `status` | `CloudComponentStatus` | no |
| `platforms` | `List<ComponentPlatform>` | no |
| `tags` | `Map<string>` | no |

## DisassociateServiceRoleFromAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `disassociatedAt` | `string` | no |

## GetComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recipeOutputFormat` | `string` | no |
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recipeOutputFormat` | `string` | yes |
| `recipe` | `blob` | yes |
| `tags` | `Map<string>` | no |

## GetComponentVersionArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `artifactName` | `string` | yes |
| `s3EndpointType` | `string` | no |
| `iotEndpointType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `preSignedUrl` | `string` | yes |

## GetConnectivityInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectivityInfo` | `List<ConnectivityInfo>` | no |
| `message` | `string` | no |

## GetCoreDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `coreDeviceThingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `coreDeviceThingName` | `string` | no |
| `coreVersion` | `string` | no |
| `platform` | `string` | no |
| `architecture` | `string` | no |
| `runtime` | `string` | no |
| `status` | `string` | no |
| `lastStatusUpdateTimestamp` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## GetDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetArn` | `string` | no |
| `revisionId` | `string` | no |
| `deploymentId` | `string` | no |
| `deploymentName` | `string` | no |
| `deploymentStatus` | `string` | no |
| `iotJobId` | `string` | no |
| `iotJobArn` | `string` | no |
| `components` | `Map<ComponentDeploymentSpecification>` | no |
| `deploymentPolicies` | `DeploymentPolicies` | no |
| `iotJobConfiguration` | `DeploymentIoTJobConfiguration` | no |
| `creationTimestamp` | `timestamp` | no |
| `isLatestForTarget` | `boolean` | no |
| `parentTargetArn` | `string` | no |
| `tags` | `Map<string>` | no |

## GetServiceRoleForAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associatedAt` | `string` | no |
| `roleArn` | `string` | no |

## ListClientDevicesAssociatedWithCoreDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `coreDeviceThingName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associatedClientDevices` | `List<AssociatedClientDevice>` | no |
| `nextToken` | `string` | no |

## ListComponentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentVersions` | `List<ComponentVersionListItem>` | no |
| `nextToken` | `string` | no |

## ListComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scope` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `components` | `List<Component>` | no |
| `nextToken` | `string` | no |

## ListCoreDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupArn` | `string` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `runtime` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `coreDevices` | `List<CoreDevice>` | no |
| `nextToken` | `string` | no |

## ListDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetArn` | `string` | no |
| `historyFilter` | `string` | no |
| `parentTargetArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployments` | `List<Deployment>` | no |
| `nextToken` | `string` | no |

## ListEffectiveDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `coreDeviceThingName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `effectiveDeployments` | `List<EffectiveDeployment>` | no |
| `nextToken` | `string` | no |

## ListInstalledComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `coreDeviceThingName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `topologyFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `installedComponents` | `List<InstalledComponent>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ResolveComponentCandidates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `platform` | `ComponentPlatform` | no |
| `componentCandidates` | `List<ComponentCandidate>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resolvedComponentVersions` | `List<ResolvedComponentVersion>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

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


## UpdateConnectivityInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `connectivityInfo` | `List<ConnectivityInfo>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | no |
| `message` | `string` | no |

