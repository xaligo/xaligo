# Amazon WorkSpaces Thin Client

API version: 2023-08-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/workspaces-thin-client/2023-08-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `desktopArn` | `string` | yes |
| `desktopEndpoint` | `string` | no |
| `softwareSetUpdateSchedule` | `string` | no |
| `maintenanceWindow` | `MaintenanceWindow` | no |
| `softwareSetUpdateMode` | `string` | no |
| `desiredSoftwareSetId` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `deviceCreationTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `EnvironmentSummary` | no |

## DeleteDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `targetDeviceStatus` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `device` | `Device` | no |

## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | no |

## GetSoftwareSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `softwareSet` | `SoftwareSet` | no |

## ListDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devices` | `List<DeviceSummary>` | no |
| `nextToken` | `string` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environments` | `List<EnvironmentSummary>` | no |
| `nextToken` | `string` | no |

## ListSoftwareSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `softwareSets` | `List<SoftwareSetSummary>` | no |
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


## UpdateDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `desiredSoftwareSetId` | `string` | no |
| `softwareSetUpdateSchedule` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `device` | `DeviceSummary` | no |

## UpdateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `desktopArn` | `string` | no |
| `desktopEndpoint` | `string` | no |
| `softwareSetUpdateSchedule` | `string` | no |
| `maintenanceWindow` | `MaintenanceWindow` | no |
| `softwareSetUpdateMode` | `string` | no |
| `desiredSoftwareSetId` | `string` | no |
| `deviceCreationTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `EnvironmentSummary` | no |

## UpdateSoftwareSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `validationStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


