# Amazon Chime

API version: 2018-05-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/chime/2018-05-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociatePhoneNumberWithUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserId` | `string` | yes |
| `E164PhoneNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateSigninDelegateGroupsWithAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `SigninDelegateGroups` | `List<SigninDelegateGroup>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchCreateRoomMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |
| `MembershipItemList` | `List<MembershipItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<MemberError>` | no |

## BatchDeletePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberErrors` | `List<PhoneNumberError>` | no |

## BatchSuspendUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserIdList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserErrors` | `List<UserError>` | no |

## BatchUnsuspendUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserIdList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserErrors` | `List<UserError>` | no |

## BatchUpdatePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdatePhoneNumberRequestItems` | `List<UpdatePhoneNumberRequestItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberErrors` | `List<PhoneNumberError>` | no |

## BatchUpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UpdateUserRequestItems` | `List<UpdateUserRequestItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserErrors` | `List<UserError>` | no |

## CreateAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Account` | `Account` | no |

## CreateBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `DisplayName` | `string` | yes |
| `Domain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bot` | `Bot` | no |

## CreateMeetingDialOut

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `FromPhoneNumber` | `string` | yes |
| `ToPhoneNumber` | `string` | yes |
| `JoinToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionId` | `string` | no |

## CreatePhoneNumberOrder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductType` | `string` | yes |
| `E164PhoneNumbers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberOrder` | `PhoneNumberOrder` | no |

## CreateRoom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Room` | `Room` | no |

## CreateRoomMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |
| `MemberId` | `string` | yes |
| `Role` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoomMembership` | `RoomMembership` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Username` | `string` | no |
| `Email` | `string` | no |
| `UserType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## DeleteAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRoom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRoomMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |
| `MemberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociatePhoneNumberFromUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSigninDelegateGroupsFromAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `GroupNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Account` | `Account` | no |

## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountSettings` | `AccountSettings` | no |

## GetBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bot` | `Bot` | no |

## GetEventsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventsConfiguration` | `EventsConfiguration` | no |

## GetGlobalSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BusinessCalling` | `BusinessCallingSettings` | no |
| `VoiceConnector` | `VoiceConnectorSettings` | no |

## GetPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `PhoneNumber` | no |

## GetPhoneNumberOrder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberOrderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberOrder` | `PhoneNumberOrder` | no |

## GetPhoneNumberSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallingName` | `string` | no |
| `CallingNameUpdatedTimestamp` | `timestamp` | no |

## GetRetentionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetentionSettings` | `RetentionSettings` | no |
| `InitiateDeletionTimestamp` | `timestamp` | no |

## GetRoom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Room` | `Room` | no |

## GetUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## GetUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserSettings` | `UserSettings` | no |

## InviteUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserEmailList` | `List<string>` | yes |
| `UserType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Invites` | `List<Invite>` | no |

## ListAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `UserEmail` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accounts` | `List<Account>` | no |
| `NextToken` | `string` | no |

## ListBots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bots` | `List<Bot>` | no |
| `NextToken` | `string` | no |

## ListPhoneNumberOrders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberOrders` | `List<PhoneNumberOrder>` | no |
| `NextToken` | `string` | no |

## ListPhoneNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `ProductType` | `string` | no |
| `FilterName` | `string` | no |
| `FilterValue` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumbers` | `List<PhoneNumber>` | no |
| `NextToken` | `string` | no |

## ListRoomMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoomMemberships` | `List<RoomMembership>` | no |
| `NextToken` | `string` | no |

## ListRooms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `MemberId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rooms` | `List<Room>` | no |
| `NextToken` | `string` | no |

## ListSupportedPhoneNumberCountries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberCountries` | `List<PhoneNumberCountry>` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserEmail` | `string` | no |
| `UserType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | no |
| `NextToken` | `string` | no |

## LogoutUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEventsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BotId` | `string` | yes |
| `OutboundEventsHTTPSEndpoint` | `string` | no |
| `LambdaFunctionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventsConfiguration` | `EventsConfiguration` | no |

## PutRetentionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RetentionSettings` | `RetentionSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetentionSettings` | `RetentionSettings` | no |
| `InitiateDeletionTimestamp` | `timestamp` | no |

## RedactConversationMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `ConversationId` | `string` | yes |
| `MessageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RedactRoomMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |
| `MessageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegenerateSecurityToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bot` | `Bot` | no |

## ResetPersonalPIN

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## RestorePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `PhoneNumber` | no |

## SearchAvailablePhoneNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AreaCode` | `string` | no |
| `City` | `string` | no |
| `Country` | `string` | no |
| `State` | `string` | no |
| `TollFreePrefix` | `string` | no |
| `PhoneNumberType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `E164PhoneNumbers` | `List<string>` | no |
| `NextToken` | `string` | no |

## UpdateAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Name` | `string` | no |
| `DefaultLicense` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Account` | `Account` | no |

## UpdateAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `AccountSettings` | `AccountSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BotId` | `string` | yes |
| `Disabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bot` | `Bot` | no |

## UpdateGlobalSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BusinessCalling` | `BusinessCallingSettings` | no |
| `VoiceConnector` | `VoiceConnectorSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |
| `ProductType` | `string` | no |
| `CallingName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `PhoneNumber` | no |

## UpdatePhoneNumberSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Room` | `Room` | no |

## UpdateRoomMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `RoomId` | `string` | yes |
| `MemberId` | `string` | yes |
| `Role` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoomMembership` | `RoomMembership` | no |

## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserId` | `string` | yes |
| `LicenseType` | `string` | no |
| `UserType` | `string` | no |
| `AlexaForBusinessMetadata` | `AlexaForBusinessMetadata` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## UpdateUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `UserId` | `string` | yes |
| `UserSettings` | `UserSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


