# AWS Key Management Service

API version: 2014-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kms/2014-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelKeyDeletion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |

## ConnectCustomKeyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomKeyStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasName` | `string` | yes |
| `TargetKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateCustomKeyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomKeyStoreName` | `string` | yes |
| `CloudHsmClusterId` | `string` | no |
| `TrustAnchorCertificate` | `string` | no |
| `KeyStorePassword` | `string` | no |
| `CustomKeyStoreType` | `string` | no |
| `XksProxyUriEndpoint` | `string` | no |
| `XksProxyUriPath` | `string` | no |
| `XksProxyVpcEndpointServiceName` | `string` | no |
| `XksProxyVpcEndpointServiceOwner` | `string` | no |
| `XksProxyAuthenticationCredential` | `XksProxyAuthenticationCredentialType` | no |
| `XksProxyConnectivity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomKeyStoreId` | `string` | no |

## CreateGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `GranteePrincipal` | `string` | no |
| `RetiringPrincipal` | `string` | no |
| `Operations` | `List<string>` | yes |
| `Constraints` | `GrantConstraints` | no |
| `GrantTokens` | `List<string>` | no |
| `Name` | `string` | no |
| `DryRun` | `boolean` | no |
| `GranteeServicePrincipal` | `string` | no |
| `RetiringServicePrincipal` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantToken` | `string` | no |
| `GrantId` | `string` | no |

## CreateKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `Description` | `string` | no |
| `KeyUsage` | `string` | no |
| `CustomerMasterKeySpec` | `string` | no |
| `KeySpec` | `string` | no |
| `Origin` | `string` | no |
| `CustomKeyStoreId` | `string` | no |
| `BypassPolicyLockoutSafetyCheck` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `MultiRegion` | `boolean` | no |
| `XksKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyMetadata` | `KeyMetadata` | no |

## Decrypt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CiphertextBlob` | `blob` | no |
| `EncryptionContext` | `Map<string>` | no |
| `GrantTokens` | `List<string>` | no |
| `KeyId` | `string` | no |
| `EncryptionAlgorithm` | `string` | no |
| `Recipient` | `RecipientInfo` | no |
| `DryRun` | `boolean` | no |
| `DryRunModifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `Plaintext` | `blob` | no |
| `EncryptionAlgorithm` | `string` | no |
| `CiphertextForRecipient` | `blob` | no |
| `KeyMaterialId` | `string` | no |

## DeleteAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomKeyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomKeyStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteImportedKeyMaterial

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `KeyMaterialId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `KeyMaterialId` | `string` | no |

## DeriveSharedSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `KeyAgreementAlgorithm` | `string` | yes |
| `PublicKey` | `blob` | yes |
| `GrantTokens` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Recipient` | `RecipientInfo` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `SharedSecret` | `blob` | no |
| `CiphertextForRecipient` | `blob` | no |
| `KeyAgreementAlgorithm` | `string` | no |
| `KeyOrigin` | `string` | no |

## DescribeCustomKeyStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomKeyStoreId` | `string` | no |
| `CustomKeyStoreName` | `string` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomKeyStores` | `List<CustomKeyStoresListEntry>` | no |
| `NextMarker` | `string` | no |
| `Truncated` | `boolean` | no |

## DescribeKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `GrantTokens` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyMetadata` | `KeyMetadata` | no |

## DisableKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableKeyRotation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisconnectCustomKeyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomKeyStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableKeyRotation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `RotationPeriodInDays` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## Encrypt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `Plaintext` | `blob` | yes |
| `EncryptionContext` | `Map<string>` | no |
| `GrantTokens` | `List<string>` | no |
| `EncryptionAlgorithm` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CiphertextBlob` | `blob` | no |
| `KeyId` | `string` | no |
| `EncryptionAlgorithm` | `string` | no |

## GenerateDataKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `EncryptionContext` | `Map<string>` | no |
| `NumberOfBytes` | `integer` | no |
| `KeySpec` | `string` | no |
| `GrantTokens` | `List<string>` | no |
| `Recipient` | `RecipientInfo` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CiphertextBlob` | `blob` | no |
| `Plaintext` | `blob` | no |
| `KeyId` | `string` | no |
| `CiphertextForRecipient` | `blob` | no |
| `KeyMaterialId` | `string` | no |

## GenerateDataKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EncryptionContext` | `Map<string>` | no |
| `KeyId` | `string` | yes |
| `KeyPairSpec` | `string` | yes |
| `GrantTokens` | `List<string>` | no |
| `Recipient` | `RecipientInfo` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrivateKeyCiphertextBlob` | `blob` | no |
| `PrivateKeyPlaintext` | `blob` | no |
| `PublicKey` | `blob` | no |
| `KeyId` | `string` | no |
| `KeyPairSpec` | `string` | no |
| `CiphertextForRecipient` | `blob` | no |
| `KeyMaterialId` | `string` | no |

## GenerateDataKeyPairWithoutPlaintext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EncryptionContext` | `Map<string>` | no |
| `KeyId` | `string` | yes |
| `KeyPairSpec` | `string` | yes |
| `GrantTokens` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrivateKeyCiphertextBlob` | `blob` | no |
| `PublicKey` | `blob` | no |
| `KeyId` | `string` | no |
| `KeyPairSpec` | `string` | no |
| `KeyMaterialId` | `string` | no |

## GenerateDataKeyWithoutPlaintext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `EncryptionContext` | `Map<string>` | no |
| `KeySpec` | `string` | no |
| `NumberOfBytes` | `integer` | no |
| `GrantTokens` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CiphertextBlob` | `blob` | no |
| `KeyId` | `string` | no |
| `KeyMaterialId` | `string` | no |

## GenerateMac

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `blob` | yes |
| `KeyId` | `string` | yes |
| `MacAlgorithm` | `string` | yes |
| `GrantTokens` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Mac` | `blob` | no |
| `MacAlgorithm` | `string` | no |
| `KeyId` | `string` | no |

## GenerateRandom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NumberOfBytes` | `integer` | no |
| `CustomKeyStoreId` | `string` | no |
| `Recipient` | `RecipientInfo` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Plaintext` | `blob` | no |
| `CiphertextForRecipient` | `blob` | no |

## GetKeyLastUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `KeyLastUsage` | `KeyLastUsageData` | no |
| `TrackingStartDate` | `timestamp` | no |
| `KeyCreationDate` | `timestamp` | no |

## GetKeyPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `PolicyName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `PolicyName` | `string` | no |

## GetKeyRotationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyRotationEnabled` | `boolean` | no |
| `KeyId` | `string` | no |
| `RotationPeriodInDays` | `integer` | no |
| `NextRotationDate` | `timestamp` | no |
| `OnDemandRotationStartDate` | `timestamp` | no |

## GetParametersForImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `WrappingAlgorithm` | `string` | yes |
| `WrappingKeySpec` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `ImportToken` | `blob` | no |
| `PublicKey` | `blob` | no |
| `ParametersValidTo` | `timestamp` | no |

## GetPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `GrantTokens` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `PublicKey` | `blob` | no |
| `CustomerMasterKeySpec` | `string` | no |
| `KeySpec` | `string` | no |
| `KeyUsage` | `string` | no |
| `EncryptionAlgorithms` | `List<string>` | no |
| `SigningAlgorithms` | `List<string>` | no |
| `KeyAgreementAlgorithms` | `List<string>` | no |

## ImportKeyMaterial

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `ImportToken` | `blob` | yes |
| `EncryptedKeyMaterial` | `blob` | yes |
| `ValidTo` | `timestamp` | no |
| `ExpirationModel` | `string` | no |
| `ImportType` | `string` | no |
| `KeyMaterialDescription` | `string` | no |
| `KeyMaterialId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `KeyMaterialId` | `string` | no |

## ListAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Aliases` | `List<AliasListEntry>` | no |
| `NextMarker` | `string` | no |
| `Truncated` | `boolean` | no |

## ListGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `Marker` | `string` | no |
| `KeyId` | `string` | yes |
| `GrantId` | `string` | no |
| `GranteePrincipal` | `string` | no |
| `GranteeServicePrincipal` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grants` | `List<GrantListEntry>` | no |
| `NextMarker` | `string` | no |
| `Truncated` | `boolean` | no |

## ListKeyPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyNames` | `List<string>` | no |
| `NextMarker` | `string` | no |
| `Truncated` | `boolean` | no |

## ListKeyRotations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `IncludeKeyMaterial` | `string` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rotations` | `List<RotationsListEntry>` | no |
| `NextMarker` | `string` | no |
| `Truncated` | `boolean` | no |

## ListKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Keys` | `List<KeyListEntry>` | no |
| `NextMarker` | `string` | no |
| `Truncated` | `boolean` | no |

## ListResourceTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextMarker` | `string` | no |
| `Truncated` | `boolean` | no |

## ListRetirableGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `Marker` | `string` | no |
| `RetiringPrincipal` | `string` | no |
| `RetiringServicePrincipal` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Grants` | `List<GrantListEntry>` | no |
| `NextMarker` | `string` | no |
| `Truncated` | `boolean` | no |

## PutKeyPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `PolicyName` | `string` | no |
| `Policy` | `string` | yes |
| `BypassPolicyLockoutSafetyCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReEncrypt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CiphertextBlob` | `blob` | no |
| `SourceEncryptionContext` | `Map<string>` | no |
| `SourceKeyId` | `string` | no |
| `DestinationKeyId` | `string` | yes |
| `DestinationEncryptionContext` | `Map<string>` | no |
| `SourceEncryptionAlgorithm` | `string` | no |
| `DestinationEncryptionAlgorithm` | `string` | no |
| `GrantTokens` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `DryRunModifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CiphertextBlob` | `blob` | no |
| `SourceKeyId` | `string` | no |
| `KeyId` | `string` | no |
| `SourceEncryptionAlgorithm` | `string` | no |
| `DestinationEncryptionAlgorithm` | `string` | no |
| `SourceKeyMaterialId` | `string` | no |
| `DestinationKeyMaterialId` | `string` | no |

## ReplicateKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `ReplicaRegion` | `string` | yes |
| `Policy` | `string` | no |
| `BypassPolicyLockoutSafetyCheck` | `boolean` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicaKeyMetadata` | `KeyMetadata` | no |
| `ReplicaPolicy` | `string` | no |
| `ReplicaTags` | `List<Tag>` | no |

## RetireGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrantToken` | `string` | no |
| `KeyId` | `string` | no |
| `GrantId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `GrantId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RotateKeyOnDemand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |

## ScheduleKeyDeletion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `PendingWindowInDays` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `DeletionDate` | `timestamp` | no |
| `KeyState` | `string` | no |
| `PendingWindowInDays` | `integer` | no |

## Sign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `Message` | `blob` | yes |
| `MessageType` | `string` | no |
| `GrantTokens` | `List<string>` | no |
| `SigningAlgorithm` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `Signature` | `blob` | no |
| `SigningAlgorithm` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasName` | `string` | yes |
| `TargetKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCustomKeyStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomKeyStoreId` | `string` | yes |
| `NewCustomKeyStoreName` | `string` | no |
| `KeyStorePassword` | `string` | no |
| `CloudHsmClusterId` | `string` | no |
| `XksProxyUriEndpoint` | `string` | no |
| `XksProxyUriPath` | `string` | no |
| `XksProxyVpcEndpointServiceName` | `string` | no |
| `XksProxyVpcEndpointServiceOwner` | `string` | no |
| `XksProxyAuthenticationCredential` | `XksProxyAuthenticationCredentialType` | no |
| `XksProxyConnectivity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateKeyDescription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `Description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePrimaryRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `PrimaryRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## Verify

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | yes |
| `Message` | `blob` | yes |
| `MessageType` | `string` | no |
| `Signature` | `blob` | yes |
| `SigningAlgorithm` | `string` | yes |
| `GrantTokens` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `SignatureValid` | `boolean` | no |
| `SigningAlgorithm` | `string` | no |

## VerifyMac

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `blob` | yes |
| `KeyId` | `string` | yes |
| `MacAlgorithm` | `string` | yes |
| `Mac` | `blob` | yes |
| `GrantTokens` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `MacValid` | `boolean` | no |
| `MacAlgorithm` | `string` | no |

