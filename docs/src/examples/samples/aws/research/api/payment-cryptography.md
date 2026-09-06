# Payment Cryptography Control Plane

API version: 2021-09-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/payment-cryptography/2021-09-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddKeyReplicationRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `ReplicationRegions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## AssociateMpaTeam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `string` | yes |
| `MpaTeamArn` | `string` | yes |
| `RequesterComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MpaTeamAssociation` | `MpaTeamAssociation` | yes |

## CreateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasName` | `string` | yes |
| `KeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alias` | `Alias` | yes |

## CreateKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyAttributes` | `KeyAttributes` | yes |
| `KeyCheckValueAlgorithm` | `string` | no |
| `Exportable` | `boolean` | yes |
| `Enabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `DeriveKeyUsage` | `string` | no |
| `ReplicationRegions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## DeleteAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `DeleteKeyInDays` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableDefaultKeyReplicationRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationRegions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnabledReplicationRegions` | `List<string>` | yes |

## DisassociateMpaTeam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `string` | yes |
| `RequesterComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MpaTeamAssociation` | `MpaTeamAssociation` | yes |

## EnableDefaultKeyReplicationRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationRegions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnabledReplicationRegions` | `List<string>` | yes |

## ExportKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyMaterial` | `ExportKeyMaterial` | yes |
| `ExportKeyIdentifier` | `string` | yes |
| `ExportAttributes` | `ExportAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WrappedKey` | `WrappedKey` | no |

## GetAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alias` | `Alias` | yes |

## GetCertificateSigningRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `SigningAlgorithm` | `string` | yes |
| `CertificateSubject` | `CertificateSubjectType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateSigningRequest` | `string` | yes |

## GetDefaultKeyReplicationRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnabledReplicationRegions` | `List<string>` | yes |

## GetKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## GetMpaTeamAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MpaTeamAssociation` | `MpaTeamAssociation` | yes |

## GetParametersForExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyMaterialType` | `string` | yes |
| `SigningKeyAlgorithm` | `string` | yes |
| `ReuseLastGeneratedToken` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SigningKeyCertificate` | `string` | yes |
| `SigningKeyCertificateChain` | `string` | yes |
| `SigningKeyAlgorithm` | `string` | yes |
| `ExportToken` | `string` | yes |
| `ParametersValidUntilTimestamp` | `timestamp` | yes |

## GetParametersForImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyMaterialType` | `string` | yes |
| `WrappingKeyAlgorithm` | `string` | yes |
| `ReuseLastGeneratedToken` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WrappingKeyCertificate` | `string` | yes |
| `WrappingKeyCertificateChain` | `string` | yes |
| `WrappingKeyAlgorithm` | `string` | yes |
| `ImportToken` | `string` | yes |
| `ParametersValidUntilTimestamp` | `timestamp` | yes |

## GetPublicKeyCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyCertificate` | `string` | yes |
| `KeyCertificateChain` | `string` | yes |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

## ImportKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyMaterial` | `ImportKeyMaterial` | yes |
| `KeyCheckValueAlgorithm` | `string` | no |
| `Enabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `ReplicationRegions` | `List<string>` | no |
| `RequesterComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## ListAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Aliases` | `List<Alias>` | yes |
| `NextToken` | `string` | no |

## ListKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyState` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Keys` | `List<KeySummary>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `NextToken` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

## RemoveKeyReplicationRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `ReplicationRegions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## RestoreKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## StartKeyUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## StopKeyUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `Key` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


## UpdateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasName` | `string` | yes |
| `KeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alias` | `Alias` | yes |

