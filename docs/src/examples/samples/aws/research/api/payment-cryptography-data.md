# Payment Cryptography Data Plane

API version: 2022-02-03. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/payment-cryptography-data/2022-02-03/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DecryptData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `CipherText` | `string` | yes |
| `DecryptionAttributes` | `EncryptionDecryptionAttributes` | yes |
| `WrappedKey` | `WrappedKey` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |
| `PlainText` | `string` | yes |

## EncryptData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `PlainText` | `string` | yes |
| `EncryptionAttributes` | `EncryptionDecryptionAttributes` | yes |
| `WrappedKey` | `WrappedKey` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | no |
| `CipherText` | `string` | yes |

## GenerateAs2805KekValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `KekValidationType` | `As2805KekValidationType` | yes |
| `RandomKeySendVariantMask` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |
| `RandomKeySend` | `string` | yes |
| `RandomKeyReceive` | `string` | yes |

## GenerateAuthRequestCryptogram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `TransactionData` | `string` | yes |
| `MajorKeyDerivationMode` | `string` | yes |
| `SessionKeyDerivationAttributes` | `SessionKeyDerivation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |
| `AuthRequestCryptogram` | `string` | yes |

## GenerateCardValidationData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `PrimaryAccountNumber` | `string` | yes |
| `GenerationAttributes` | `CardGenerationAttributes` | yes |
| `ValidationDataLength` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |
| `ValidationData` | `string` | yes |

## GenerateMac

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `MessageData` | `string` | yes |
| `GenerationAttributes` | `MacAttributes` | yes |
| `MacLength` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |
| `Mac` | `string` | yes |

## GenerateMacEmvPinChange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NewPinPekIdentifier` | `string` | yes |
| `NewEncryptedPinBlock` | `string` | yes |
| `PinBlockFormat` | `string` | yes |
| `SecureMessagingIntegrityKeyIdentifier` | `string` | yes |
| `SecureMessagingConfidentialityKeyIdentifier` | `string` | yes |
| `MessageData` | `string` | yes |
| `DerivationMethodAttributes` | `DerivationMethodAttributes` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NewPinPekArn` | `string` | yes |
| `SecureMessagingIntegrityKeyArn` | `string` | yes |
| `SecureMessagingConfidentialityKeyArn` | `string` | yes |
| `Mac` | `string` | yes |
| `EncryptedPinBlock` | `string` | yes |
| `NewPinPekKeyCheckValue` | `string` | yes |
| `SecureMessagingIntegrityKeyCheckValue` | `string` | yes |
| `SecureMessagingConfidentialityKeyCheckValue` | `string` | yes |
| `VisaAmexDerivationOutputs` | `VisaAmexDerivationOutputs` | no |

## GeneratePinData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GenerationKeyIdentifier` | `string` | yes |
| `EncryptionKeyIdentifier` | `string` | yes |
| `GenerationAttributes` | `PinGenerationAttributes` | yes |
| `PinDataLength` | `integer` | no |
| `PrimaryAccountNumber` | `string` | no |
| `PinBlockFormat` | `string` | yes |
| `EncryptionWrappedKey` | `WrappedKey` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GenerationKeyArn` | `string` | yes |
| `GenerationKeyCheckValue` | `string` | yes |
| `EncryptionKeyArn` | `string` | yes |
| `EncryptionKeyCheckValue` | `string` | yes |
| `EncryptedPinBlock` | `string` | yes |
| `PinData` | `PinData` | yes |

## ReEncryptData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IncomingKeyIdentifier` | `string` | yes |
| `OutgoingKeyIdentifier` | `string` | yes |
| `CipherText` | `string` | yes |
| `IncomingEncryptionAttributes` | `ReEncryptionAttributes` | yes |
| `OutgoingEncryptionAttributes` | `ReEncryptionAttributes` | yes |
| `IncomingWrappedKey` | `WrappedKey` | no |
| `OutgoingWrappedKey` | `WrappedKey` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |
| `CipherText` | `string` | yes |

## TranslateKeyMaterial

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IncomingKeyMaterial` | `IncomingKeyMaterial` | yes |
| `OutgoingKeyMaterial` | `OutgoingKeyMaterial` | yes |
| `KeyCheckValueAlgorithm` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WrappedKey` | `WrappedWorkingKey` | yes |

## TranslatePinData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IncomingKeyIdentifier` | `string` | yes |
| `OutgoingKeyIdentifier` | `string` | yes |
| `IncomingTranslationAttributes` | `TranslationIsoFormats` | yes |
| `OutgoingTranslationAttributes` | `TranslationIsoFormats` | yes |
| `EncryptedPinBlock` | `string` | yes |
| `IncomingDukptAttributes` | `DukptDerivationAttributes` | no |
| `OutgoingDukptAttributes` | `DukptDerivationAttributes` | no |
| `IncomingWrappedKey` | `WrappedKey` | no |
| `OutgoingWrappedKey` | `WrappedKey` | no |
| `IncomingAs2805Attributes` | `As2805PekDerivationAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PinBlock` | `string` | yes |
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |

## VerifyAuthRequestCryptogram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `TransactionData` | `string` | yes |
| `AuthRequestCryptogram` | `string` | yes |
| `MajorKeyDerivationMode` | `string` | yes |
| `SessionKeyDerivationAttributes` | `SessionKeyDerivation` | yes |
| `AuthResponseAttributes` | `CryptogramAuthResponse` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |
| `AuthResponseValue` | `string` | no |

## VerifyCardValidationData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `PrimaryAccountNumber` | `string` | yes |
| `VerificationAttributes` | `CardVerificationAttributes` | yes |
| `ValidationData` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |

## VerifyMac

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyIdentifier` | `string` | yes |
| `MessageData` | `string` | yes |
| `Mac` | `string` | yes |
| `VerificationAttributes` | `MacAttributes` | yes |
| `MacLength` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyArn` | `string` | yes |
| `KeyCheckValue` | `string` | yes |

## VerifyPinData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerificationKeyIdentifier` | `string` | yes |
| `EncryptionKeyIdentifier` | `string` | yes |
| `VerificationAttributes` | `PinVerificationAttributes` | yes |
| `EncryptedPinBlock` | `string` | yes |
| `PrimaryAccountNumber` | `string` | no |
| `PinBlockFormat` | `string` | yes |
| `PinDataLength` | `integer` | no |
| `DukptAttributes` | `DukptAttributes` | no |
| `EncryptionWrappedKey` | `WrappedKey` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerificationKeyArn` | `string` | yes |
| `VerificationKeyCheckValue` | `string` | yes |
| `EncryptionKeyArn` | `string` | yes |
| `EncryptionKeyCheckValue` | `string` | yes |

