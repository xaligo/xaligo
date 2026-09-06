# AWS Import/Export

API version: 2010-06-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/importexport/2010-06-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `APIVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Success` | `boolean` | no |

## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobType` | `string` | yes |
| `Manifest` | `string` | yes |
| `ManifestAddendum` | `string` | no |
| `ValidateOnly` | `boolean` | yes |
| `APIVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobType` | `string` | no |
| `Signature` | `string` | no |
| `SignatureFileContents` | `string` | no |
| `WarningMessage` | `string` | no |
| `ArtifactList` | `List<Artifact>` | no |

## GetShippingLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIds` | `List<string>` | yes |
| `name` | `string` | no |
| `company` | `string` | no |
| `phoneNumber` | `string` | no |
| `country` | `string` | no |
| `stateOrProvince` | `string` | no |
| `city` | `string` | no |
| `postalCode` | `string` | no |
| `street1` | `string` | no |
| `street2` | `string` | no |
| `street3` | `string` | no |
| `APIVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShippingLabelURL` | `string` | no |
| `Warning` | `string` | no |

## GetStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `APIVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobType` | `string` | no |
| `LocationCode` | `string` | no |
| `LocationMessage` | `string` | no |
| `ProgressCode` | `string` | no |
| `ProgressMessage` | `string` | no |
| `Carrier` | `string` | no |
| `TrackingNumber` | `string` | no |
| `LogBucket` | `string` | no |
| `LogKey` | `string` | no |
| `ErrorCount` | `integer` | no |
| `Signature` | `string` | no |
| `SignatureFileContents` | `string` | no |
| `CurrentManifest` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `ArtifactList` | `List<Artifact>` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxJobs` | `integer` | no |
| `Marker` | `string` | no |
| `APIVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<Job>` | no |
| `IsTruncated` | `boolean` | no |

## UpdateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `Manifest` | `string` | yes |
| `JobType` | `string` | yes |
| `ValidateOnly` | `boolean` | yes |
| `APIVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Success` | `boolean` | no |
| `WarningMessage` | `string` | no |
| `ArtifactList` | `List<Artifact>` | no |

