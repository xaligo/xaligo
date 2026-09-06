# AWS ARC - Zonal Shift

API version: 2022-10-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/arc-zonal-shift/2022-10-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelPracticeRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `zonalShiftId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `zonalShiftId` | `string` | yes |
| `resourceIdentifier` | `string` | yes |
| `awayFrom` | `string` | yes |
| `expiryTime` | `timestamp` | yes |
| `startTime` | `timestamp` | yes |
| `status` | `string` | yes |
| `comment` | `string` | yes |

## CancelZonalShift

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `zonalShiftId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `zonalShiftId` | `string` | yes |
| `resourceIdentifier` | `string` | yes |
| `awayFrom` | `string` | yes |
| `expiryTime` | `timestamp` | yes |
| `startTime` | `timestamp` | yes |
| `status` | `string` | yes |
| `comment` | `string` | yes |

## CreatePracticeRunConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |
| `blockedWindows` | `List<string>` | no |
| `blockedDates` | `List<string>` | no |
| `blockingAlarms` | `List<ControlCondition>` | no |
| `allowedWindows` | `List<string>` | no |
| `outcomeAlarms` | `List<ControlCondition>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `zonalAutoshiftStatus` | `string` | yes |
| `practiceRunConfiguration` | `PracticeRunConfiguration` | yes |

## DeletePracticeRunConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `zonalAutoshiftStatus` | `string` | yes |

## GetAutoshiftObserverNotificationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## GetManagedResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `name` | `string` | no |
| `appliedWeights` | `Map<float>` | yes |
| `zonalShifts` | `List<ZonalShiftInResource>` | yes |
| `autoshifts` | `List<AutoshiftInResource>` | no |
| `practiceRunConfiguration` | `PracticeRunConfiguration` | no |
| `zonalAutoshiftStatus` | `string` | no |

## ListAutoshifts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AutoshiftSummary>` | no |
| `nextToken` | `string` | no |

## ListManagedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ManagedResourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListZonalShifts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `resourceIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ZonalShiftSummary>` | no |
| `nextToken` | `string` | no |

## StartPracticeRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |
| `awayFrom` | `string` | yes |
| `comment` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `zonalShiftId` | `string` | yes |
| `resourceIdentifier` | `string` | yes |
| `awayFrom` | `string` | yes |
| `expiryTime` | `timestamp` | yes |
| `startTime` | `timestamp` | yes |
| `status` | `string` | yes |
| `comment` | `string` | yes |

## StartZonalShift

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |
| `awayFrom` | `string` | yes |
| `expiresIn` | `string` | yes |
| `comment` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `zonalShiftId` | `string` | yes |
| `resourceIdentifier` | `string` | yes |
| `awayFrom` | `string` | yes |
| `expiryTime` | `timestamp` | yes |
| `startTime` | `timestamp` | yes |
| `status` | `string` | yes |
| `comment` | `string` | yes |

## UpdateAutoshiftObserverNotificationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## UpdatePracticeRunConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |
| `blockedWindows` | `List<string>` | no |
| `blockedDates` | `List<string>` | no |
| `blockingAlarms` | `List<ControlCondition>` | no |
| `allowedWindows` | `List<string>` | no |
| `outcomeAlarms` | `List<ControlCondition>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `zonalAutoshiftStatus` | `string` | yes |
| `practiceRunConfiguration` | `PracticeRunConfiguration` | yes |

## UpdateZonalAutoshiftConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |
| `zonalAutoshiftStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |
| `zonalAutoshiftStatus` | `string` | yes |

## UpdateZonalShift

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `zonalShiftId` | `string` | yes |
| `comment` | `string` | no |
| `expiresIn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `zonalShiftId` | `string` | yes |
| `resourceIdentifier` | `string` | yes |
| `awayFrom` | `string` | yes |
| `expiryTime` | `timestamp` | yes |
| `startTime` | `timestamp` | yes |
| `status` | `string` | yes |
| `comment` | `string` | yes |

