# Inspector Scan

API version: 2023-08-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/inspector-scan/2023-08-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ScanSbom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sbom` | `Sbom` | yes |
| `outputFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sbom` | `Sbom` | no |

