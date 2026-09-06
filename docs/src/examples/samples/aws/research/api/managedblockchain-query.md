# Amazon Managed Blockchain Query

API version: 2023-05-04. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/managedblockchain-query/2023-05-04/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetTokenBalance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `getTokenBalanceInputs` | `List<BatchGetTokenBalanceInputItem>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenBalances` | `List<BatchGetTokenBalanceOutputItem>` | yes |
| `errors` | `List<BatchGetTokenBalanceErrorItem>` | yes |

## GetAssetContract

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contractIdentifier` | `ContractIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contractIdentifier` | `ContractIdentifier` | yes |
| `tokenStandard` | `string` | yes |
| `deployerAddress` | `string` | yes |
| `metadata` | `ContractMetadata` | no |

## GetTokenBalance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenIdentifier` | `TokenIdentifier` | yes |
| `ownerIdentifier` | `OwnerIdentifier` | yes |
| `atBlockchainInstant` | `BlockchainInstant` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerIdentifier` | `OwnerIdentifier` | no |
| `tokenIdentifier` | `TokenIdentifier` | no |
| `balance` | `string` | yes |
| `atBlockchainInstant` | `BlockchainInstant` | yes |
| `lastUpdatedTime` | `BlockchainInstant` | no |

## GetTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transactionHash` | `string` | no |
| `transactionId` | `string` | no |
| `network` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transaction` | `Transaction` | yes |

## ListAssetContracts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contractFilter` | `ContractFilter` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contracts` | `List<AssetContract>` | yes |
| `nextToken` | `string` | no |

## ListFilteredTransactionEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `network` | `string` | yes |
| `addressIdentifierFilter` | `AddressIdentifierFilter` | yes |
| `timeFilter` | `TimeFilter` | no |
| `voutFilter` | `VoutFilter` | no |
| `confirmationStatusFilter` | `ConfirmationStatusFilter` | no |
| `sort` | `ListFilteredTransactionEventsSort` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<TransactionEvent>` | yes |
| `nextToken` | `string` | no |

## ListTokenBalances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerFilter` | `OwnerFilter` | no |
| `tokenFilter` | `TokenFilter` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenBalances` | `List<TokenBalance>` | yes |
| `nextToken` | `string` | no |

## ListTransactionEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transactionHash` | `string` | no |
| `transactionId` | `string` | no |
| `network` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<TransactionEvent>` | yes |
| `nextToken` | `string` | no |

## ListTransactions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `address` | `string` | yes |
| `network` | `string` | yes |
| `fromBlockchainInstant` | `BlockchainInstant` | no |
| `toBlockchainInstant` | `BlockchainInstant` | no |
| `sort` | `ListTransactionsSort` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `confirmationStatusFilter` | `ConfirmationStatusFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transactions` | `List<TransactionOutputItem>` | yes |
| `nextToken` | `string` | no |

