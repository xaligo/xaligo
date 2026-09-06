# Amazon Chime SDK Meetings

API version: 2021-07-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/chime-sdk-meetings/2021-07-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCreateAttendee

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `Attendees` | `List<CreateAttendeeRequestItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attendees` | `List<Attendee>` | no |
| `Errors` | `List<CreateAttendeeError>` | no |

## BatchUpdateAttendeeCapabilitiesExcept

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `ExcludedAttendeeIds` | `List<AttendeeIdItem>` | yes |
| `Capabilities` | `AttendeeCapabilities` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAttendee

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `ExternalUserId` | `string` | yes |
| `Capabilities` | `AttendeeCapabilities` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attendee` | `Attendee` | no |

## CreateMeeting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | yes |
| `MediaRegion` | `string` | yes |
| `MeetingHostId` | `string` | no |
| `ExternalMeetingId` | `string` | yes |
| `NotificationsConfiguration` | `NotificationsConfiguration` | no |
| `MeetingFeatures` | `MeetingFeaturesConfiguration` | no |
| `PrimaryMeetingId` | `string` | no |
| `TenantIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `MediaPlacementNetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Meeting` | `Meeting` | no |

## CreateMeetingWithAttendees

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | yes |
| `MediaRegion` | `string` | yes |
| `MeetingHostId` | `string` | no |
| `ExternalMeetingId` | `string` | yes |
| `MeetingFeatures` | `MeetingFeaturesConfiguration` | no |
| `NotificationsConfiguration` | `NotificationsConfiguration` | no |
| `Attendees` | `List<CreateAttendeeRequestItem>` | yes |
| `PrimaryMeetingId` | `string` | no |
| `TenantIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `MediaPlacementNetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Meeting` | `Meeting` | no |
| `Attendees` | `List<Attendee>` | no |
| `Errors` | `List<CreateAttendeeError>` | no |

## DeleteAttendee

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `AttendeeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMeeting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAttendee

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `AttendeeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attendee` | `Attendee` | no |

## GetMeeting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Meeting` | `Meeting` | no |

## ListAttendees

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attendees` | `List<Attendee>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## StartMeetingTranscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `TranscriptionConfiguration` | `TranscriptionConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopMeetingTranscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAttendeeCapabilities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MeetingId` | `string` | yes |
| `AttendeeId` | `string` | yes |
| `Capabilities` | `AttendeeCapabilities` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attendee` | `Attendee` | no |

