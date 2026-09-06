# Amazon Kinesis Video Streams Archived Media

API version: 2017-09-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kinesis-video-archived-media/2017-09-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetClip

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `ClipFragmentSelector` | `ClipFragmentSelector` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | no |
| `Payload` | `blob` | no |

## GetDASHStreamingSessionURL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `PlaybackMode` | `string` | no |
| `DisplayFragmentTimestamp` | `string` | no |
| `DisplayFragmentNumber` | `string` | no |
| `DASHFragmentSelector` | `DASHFragmentSelector` | no |
| `Expires` | `integer` | no |
| `MaxManifestFragmentResults` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DASHStreamingSessionURL` | `string` | no |

## GetHLSStreamingSessionURL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `PlaybackMode` | `string` | no |
| `HLSFragmentSelector` | `HLSFragmentSelector` | no |
| `ContainerFormat` | `string` | no |
| `DiscontinuityMode` | `string` | no |
| `DisplayFragmentTimestamp` | `string` | no |
| `Expires` | `integer` | no |
| `MaxMediaPlaylistFragmentResults` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HLSStreamingSessionURL` | `string` | no |

## GetImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `ImageSelectorType` | `string` | yes |
| `StartTimestamp` | `timestamp` | yes |
| `EndTimestamp` | `timestamp` | yes |
| `SamplingInterval` | `integer` | no |
| `Format` | `string` | yes |
| `FormatConfig` | `Map<string>` | no |
| `WidthPixels` | `integer` | no |
| `HeightPixels` | `integer` | no |
| `MaxResults` | `long` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Images` | `List<Image>` | no |
| `NextToken` | `string` | no |

## GetMediaForFragmentList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `Fragments` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | no |
| `Payload` | `blob` | no |

## ListFragments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `MaxResults` | `long` | no |
| `NextToken` | `string` | no |
| `FragmentSelector` | `FragmentSelector` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fragments` | `List<Fragment>` | no |
| `NextToken` | `string` | no |

