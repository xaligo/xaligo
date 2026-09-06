# AWS Control Tower

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/controltower/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateLandingZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | yes |
| `remediationTypes` | `List<string>` | no |
| `tags` | `Map<string>` | no |
| `manifest` | `Manifest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `operationIdentifier` | `string` | yes |

## DeleteLandingZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `landingZoneIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

## DisableBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledBaselineIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

## DisableControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlIdentifier` | `string` | no |
| `targetIdentifier` | `string` | no |
| `enabledControlIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

## EnableBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baselineVersion` | `string` | yes |
| `parameters` | `List<EnabledBaselineParameter>` | no |
| `baselineIdentifier` | `string` | yes |
| `targetIdentifier` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |
| `arn` | `string` | yes |

## EnableControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlIdentifier` | `string` | yes |
| `targetIdentifier` | `string` | yes |
| `tags` | `Map<string>` | no |
| `parameters` | `List<EnabledControlParameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |
| `arn` | `string` | no |

## GetBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baselineIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |

## GetBaselineOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baselineOperation` | `BaselineOperation` | yes |

## GetControlOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlOperation` | `ControlOperation` | yes |

## GetEnabledBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledBaselineIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledBaselineDetails` | `EnabledBaselineDetails` | no |

## GetEnabledControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledControlIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledControlDetails` | `EnabledControlDetails` | yes |

## GetLandingZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `landingZoneIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `landingZone` | `LandingZoneDetail` | yes |

## GetLandingZoneOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationDetails` | `LandingZoneOperationDetail` | yes |

## ListBaselines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baselines` | `List<BaselineSummary>` | yes |
| `nextToken` | `string` | no |

## ListControlOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `ControlOperationFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `controlOperations` | `List<ControlOperationSummary>` | yes |
| `nextToken` | `string` | no |

## ListEnabledBaselines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `EnabledBaselineFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `includeChildren` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledBaselines` | `List<EnabledBaselineSummary>` | yes |
| `nextToken` | `string` | no |

## ListEnabledControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetIdentifier` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `EnabledControlFilter` | no |
| `includeChildren` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledControls` | `List<EnabledControlSummary>` | yes |
| `nextToken` | `string` | no |

## ListLandingZoneOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `LandingZoneOperationFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `landingZoneOperations` | `List<LandingZoneOperationSummary>` | yes |
| `nextToken` | `string` | no |

## ListLandingZones

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `landingZones` | `List<LandingZoneSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## ResetEnabledBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledBaselineIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

## ResetEnabledControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enabledControlIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

## ResetLandingZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `landingZoneIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

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


## UpdateEnabledBaseline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baselineVersion` | `string` | yes |
| `parameters` | `List<EnabledBaselineParameter>` | no |
| `enabledBaselineIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

## UpdateEnabledControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `parameters` | `List<EnabledControlParameter>` | yes |
| `enabledControlIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

## UpdateLandingZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | yes |
| `remediationTypes` | `List<string>` | no |
| `landingZoneIdentifier` | `string` | yes |
| `manifest` | `Manifest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operationIdentifier` | `string` | yes |

