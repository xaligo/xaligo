# Amazon Mechanical Turk

API version: 2017-01-17. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mturk/2017-01-17/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptQualificationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationRequestId` | `string` | yes |
| `IntegerValue` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ApproveAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentId` | `string` | yes |
| `RequesterFeedback` | `string` | no |
| `OverrideRejection` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateQualificationWithWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationTypeId` | `string` | yes |
| `WorkerId` | `string` | yes |
| `IntegerValue` | `integer` | no |
| `SendNotification` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAdditionalAssignmentsForHIT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | yes |
| `NumberOfAdditionalAssignments` | `integer` | yes |
| `UniqueRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateHIT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxAssignments` | `integer` | no |
| `AutoApprovalDelayInSeconds` | `long` | no |
| `LifetimeInSeconds` | `long` | yes |
| `AssignmentDurationInSeconds` | `long` | yes |
| `Reward` | `string` | yes |
| `Title` | `string` | yes |
| `Keywords` | `string` | no |
| `Description` | `string` | yes |
| `Question` | `string` | no |
| `RequesterAnnotation` | `string` | no |
| `QualificationRequirements` | `List<QualificationRequirement>` | no |
| `UniqueRequestToken` | `string` | no |
| `AssignmentReviewPolicy` | `ReviewPolicy` | no |
| `HITReviewPolicy` | `ReviewPolicy` | no |
| `HITLayoutId` | `string` | no |
| `HITLayoutParameters` | `List<HITLayoutParameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HIT` | `HIT` | no |

## CreateHITType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoApprovalDelayInSeconds` | `long` | no |
| `AssignmentDurationInSeconds` | `long` | yes |
| `Reward` | `string` | yes |
| `Title` | `string` | yes |
| `Keywords` | `string` | no |
| `Description` | `string` | yes |
| `QualificationRequirements` | `List<QualificationRequirement>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITTypeId` | `string` | no |

## CreateHITWithHITType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITTypeId` | `string` | yes |
| `MaxAssignments` | `integer` | no |
| `LifetimeInSeconds` | `long` | yes |
| `Question` | `string` | no |
| `RequesterAnnotation` | `string` | no |
| `UniqueRequestToken` | `string` | no |
| `AssignmentReviewPolicy` | `ReviewPolicy` | no |
| `HITReviewPolicy` | `ReviewPolicy` | no |
| `HITLayoutId` | `string` | no |
| `HITLayoutParameters` | `List<HITLayoutParameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HIT` | `HIT` | no |

## CreateQualificationType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Keywords` | `string` | no |
| `Description` | `string` | yes |
| `QualificationTypeStatus` | `string` | yes |
| `RetryDelayInSeconds` | `long` | no |
| `Test` | `string` | no |
| `AnswerKey` | `string` | no |
| `TestDurationInSeconds` | `long` | no |
| `AutoGranted` | `boolean` | no |
| `AutoGrantedValue` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationType` | `QualificationType` | no |

## CreateWorkerBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkerId` | `string` | yes |
| `Reason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHIT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQualificationType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationTypeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkerBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkerId` | `string` | yes |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateQualificationFromWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkerId` | `string` | yes |
| `QualificationTypeId` | `string` | yes |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountBalance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailableBalance` | `string` | no |
| `OnHoldBalance` | `string` | no |

## GetAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Assignment` | `Assignment` | no |
| `HIT` | `HIT` | no |

## GetFileUploadURL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentId` | `string` | yes |
| `QuestionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileUploadURL` | `string` | no |

## GetHIT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HIT` | `HIT` | no |

## GetQualificationScore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationTypeId` | `string` | yes |
| `WorkerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Qualification` | `Qualification` | no |

## GetQualificationType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationTypeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationType` | `QualificationType` | no |

## ListAssignmentsForHIT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AssignmentStatuses` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NumResults` | `integer` | no |
| `Assignments` | `List<Assignment>` | no |

## ListBonusPayments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | no |
| `AssignmentId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NumResults` | `integer` | no |
| `NextToken` | `string` | no |
| `BonusPayments` | `List<BonusPayment>` | no |

## ListHITs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NumResults` | `integer` | no |
| `HITs` | `List<HIT>` | no |

## ListHITsForQualificationType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationTypeId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NumResults` | `integer` | no |
| `HITs` | `List<HIT>` | no |

## ListQualificationRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationTypeId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NumResults` | `integer` | no |
| `NextToken` | `string` | no |
| `QualificationRequests` | `List<QualificationRequest>` | no |

## ListQualificationTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Query` | `string` | no |
| `MustBeRequestable` | `boolean` | yes |
| `MustBeOwnedByCaller` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NumResults` | `integer` | no |
| `NextToken` | `string` | no |
| `QualificationTypes` | `List<QualificationType>` | no |

## ListReviewPolicyResultsForHIT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | yes |
| `PolicyLevels` | `List<string>` | no |
| `RetrieveActions` | `boolean` | no |
| `RetrieveResults` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | no |
| `AssignmentReviewPolicy` | `ReviewPolicy` | no |
| `HITReviewPolicy` | `ReviewPolicy` | no |
| `AssignmentReviewReport` | `ReviewReport` | no |
| `HITReviewReport` | `ReviewReport` | no |
| `NextToken` | `string` | no |

## ListReviewableHITs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITTypeId` | `string` | no |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NumResults` | `integer` | no |
| `HITs` | `List<HIT>` | no |

## ListWorkerBlocks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NumResults` | `integer` | no |
| `WorkerBlocks` | `List<WorkerBlock>` | no |

## ListWorkersWithQualificationType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationTypeId` | `string` | yes |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NumResults` | `integer` | no |
| `Qualifications` | `List<Qualification>` | no |

## NotifyWorkers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subject` | `string` | yes |
| `MessageText` | `string` | yes |
| `WorkerIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyWorkersFailureStatuses` | `List<NotifyWorkersFailureStatus>` | no |

## RejectAssignment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignmentId` | `string` | yes |
| `RequesterFeedback` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RejectQualificationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationRequestId` | `string` | yes |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendBonus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkerId` | `string` | yes |
| `BonusAmount` | `string` | yes |
| `AssignmentId` | `string` | yes |
| `Reason` | `string` | yes |
| `UniqueRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendTestEventNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Notification` | `NotificationSpecification` | yes |
| `TestEventType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateExpirationForHIT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | yes |
| `ExpireAt` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateHITReviewStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | yes |
| `Revert` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateHITTypeOfHIT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITId` | `string` | yes |
| `HITTypeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateNotificationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HITTypeId` | `string` | yes |
| `Notification` | `NotificationSpecification` | no |
| `Active` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQualificationType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationTypeId` | `string` | yes |
| `Description` | `string` | no |
| `QualificationTypeStatus` | `string` | no |
| `Test` | `string` | no |
| `AnswerKey` | `string` | no |
| `TestDurationInSeconds` | `long` | no |
| `RetryDelayInSeconds` | `long` | no |
| `AutoGranted` | `boolean` | no |
| `AutoGrantedValue` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QualificationType` | `QualificationType` | no |

