# AWS IoT

API version: 2015-05-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iot/2015-05-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptCertificateTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |
| `setAsActive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddThingToBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | no |
| `billingGroupArn` | `string` | no |
| `thingName` | `string` | no |
| `thingArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddThingToThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | no |
| `thingGroupArn` | `string` | no |
| `thingName` | `string` | no |
| `thingArn` | `string` | no |
| `overrideDynamicGroups` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateSbomWithPackageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `versionName` | `string` | yes |
| `sbom` | `Sbom` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | no |
| `versionName` | `string` | no |
| `sbom` | `Sbom` | no |
| `sbomValidationStatus` | `string` | no |

## AssociateTargetsWithJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targets` | `List<string>` | yes |
| `jobId` | `string` | yes |
| `comment` | `string` | no |
| `namespaceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |
| `jobId` | `string` | no |
| `description` | `string` | no |

## AttachPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `target` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachPrincipalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `principal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | yes |
| `securityProfileTargetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachThingPrincipal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `principal` | `string` | yes |
| `thingPrincipalType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelAuditMitigationActionsTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelAuditTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelCertificateTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelDetectMitigationActionsTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `reasonCode` | `string` | no |
| `comment` | `string` | no |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |
| `jobId` | `string` | no |
| `description` | `string` | no |

## CancelJobExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `thingName` | `string` | yes |
| `force` | `boolean` | no |
| `expectedVersion` | `long` | no |
| `statusDetails` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ClearDefaultAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ConfirmTopicRuleDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `confirmationToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAuditSuppression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkName` | `string` | yes |
| `resourceIdentifier` | `ResourceIdentifier` | yes |
| `expirationDate` | `timestamp` | no |
| `suppressIndefinitely` | `boolean` | no |
| `description` | `string` | no |
| `clientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | yes |
| `authorizerFunctionArn` | `string` | yes |
| `tokenKeyName` | `string` | no |
| `tokenSigningPublicKeys` | `Map<string>` | no |
| `status` | `string` | no |
| `tags` | `List<Tag>` | no |
| `signingDisabled` | `boolean` | no |
| `enableCachingForHttp` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | no |
| `authorizerArn` | `string` | no |

## CreateBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | yes |
| `billingGroupProperties` | `BillingGroupProperties` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | no |
| `billingGroupArn` | `string` | no |
| `billingGroupId` | `string` | no |

## CreateCertificateFromCsr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateSigningRequest` | `string` | yes |
| `setAsActive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateArn` | `string` | no |
| `certificateId` | `string` | no |
| `certificatePem` | `string` | no |

## CreateCertificateProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateProviderName` | `string` | yes |
| `lambdaFunctionArn` | `string` | yes |
| `accountDefaultForOperations` | `List<string>` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateProviderName` | `string` | no |
| `certificateProviderArn` | `string` | no |

## CreateCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandId` | `string` | yes |
| `namespace` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `payload` | `CommandPayload` | no |
| `payloadTemplate` | `string` | no |
| `preprocessor` | `CommandPreprocessor` | no |
| `mandatoryParameters` | `List<CommandParameter>` | no |
| `roleArn` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandId` | `string` | no |
| `commandArn` | `string` | no |

## CreateCustomMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | yes |
| `displayName` | `string` | no |
| `metricType` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `clientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricArn` | `string` | no |

## CreateDimension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `stringValues` | `List<string>` | yes |
| `tags` | `List<Tag>` | no |
| `clientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |

## CreateDomainConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainConfigurationName` | `string` | yes |
| `domainName` | `string` | no |
| `serverCertificateArns` | `List<string>` | no |
| `validationCertificateArn` | `string` | no |
| `authorizerConfig` | `AuthorizerConfig` | no |
| `serviceType` | `string` | no |
| `tags` | `List<Tag>` | no |
| `tlsConfig` | `TlsConfig` | no |
| `serverCertificateConfig` | `ServerCertificateConfig` | no |
| `authenticationType` | `string` | no |
| `applicationProtocol` | `string` | no |
| `clientCertificateConfig` | `ClientCertificateConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainConfigurationName` | `string` | no |
| `domainConfigurationArn` | `string` | no |

## CreateDynamicThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | yes |
| `thingGroupProperties` | `ThingGroupProperties` | no |
| `indexName` | `string` | no |
| `queryString` | `string` | yes |
| `queryVersion` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | no |
| `thingGroupArn` | `string` | no |
| `thingGroupId` | `string` | no |
| `indexName` | `string` | no |
| `queryString` | `string` | no |
| `queryVersion` | `string` | no |

## CreateFleetMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | yes |
| `queryString` | `string` | yes |
| `aggregationType` | `AggregationType` | yes |
| `period` | `integer` | yes |
| `aggregationField` | `string` | yes |
| `description` | `string` | no |
| `queryVersion` | `string` | no |
| `indexName` | `string` | no |
| `unit` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricArn` | `string` | no |

## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `targets` | `List<string>` | yes |
| `documentSource` | `string` | no |
| `document` | `string` | no |
| `description` | `string` | no |
| `presignedUrlConfig` | `PresignedUrlConfig` | no |
| `targetSelection` | `string` | no |
| `jobExecutionsRolloutConfig` | `JobExecutionsRolloutConfig` | no |
| `abortConfig` | `AbortConfig` | no |
| `timeoutConfig` | `TimeoutConfig` | no |
| `tags` | `List<Tag>` | no |
| `namespaceId` | `string` | no |
| `jobTemplateArn` | `string` | no |
| `jobExecutionsRetryConfig` | `JobExecutionsRetryConfig` | no |
| `documentParameters` | `Map<string>` | no |
| `schedulingConfig` | `SchedulingConfig` | no |
| `destinationPackageVersions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |
| `jobId` | `string` | no |
| `description` | `string` | no |

## CreateJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobTemplateId` | `string` | yes |
| `jobArn` | `string` | no |
| `documentSource` | `string` | no |
| `document` | `string` | no |
| `description` | `string` | yes |
| `presignedUrlConfig` | `PresignedUrlConfig` | no |
| `jobExecutionsRolloutConfig` | `JobExecutionsRolloutConfig` | no |
| `abortConfig` | `AbortConfig` | no |
| `timeoutConfig` | `TimeoutConfig` | no |
| `tags` | `List<Tag>` | no |
| `jobExecutionsRetryConfig` | `JobExecutionsRetryConfig` | no |
| `maintenanceWindows` | `List<MaintenanceWindow>` | no |
| `destinationPackageVersions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobTemplateArn` | `string` | no |
| `jobTemplateId` | `string` | no |

## CreateKeysAndCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `setAsActive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateArn` | `string` | no |
| `certificateId` | `string` | no |
| `certificatePem` | `string` | no |
| `keyPair` | `KeyPair` | no |

## CreateMitigationAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionName` | `string` | yes |
| `roleArn` | `string` | yes |
| `actionParams` | `MitigationActionParams` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionArn` | `string` | no |
| `actionId` | `string` | no |

## CreateOTAUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `otaUpdateId` | `string` | yes |
| `description` | `string` | no |
| `targets` | `List<string>` | yes |
| `protocols` | `List<string>` | no |
| `targetSelection` | `string` | no |
| `awsJobExecutionsRolloutConfig` | `AwsJobExecutionsRolloutConfig` | no |
| `awsJobPresignedUrlConfig` | `AwsJobPresignedUrlConfig` | no |
| `awsJobAbortConfig` | `AwsJobAbortConfig` | no |
| `awsJobTimeoutConfig` | `AwsJobTimeoutConfig` | no |
| `files` | `List<OTAUpdateFile>` | yes |
| `roleArn` | `string` | yes |
| `additionalParameters` | `Map<string>` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `otaUpdateId` | `string` | no |
| `awsIotJobId` | `string` | no |
| `otaUpdateArn` | `string` | no |
| `awsIotJobArn` | `string` | no |
| `otaUpdateStatus` | `string` | no |

## CreatePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | no |
| `packageArn` | `string` | no |
| `description` | `string` | no |

## CreatePackageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `versionName` | `string` | yes |
| `description` | `string` | no |
| `attributes` | `Map<string>` | no |
| `artifact` | `PackageVersionArtifact` | no |
| `recipe` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageVersionArn` | `string` | no |
| `packageName` | `string` | no |
| `versionName` | `string` | no |
| `description` | `string` | no |
| `attributes` | `Map<string>` | no |
| `status` | `string` | no |
| `errorReason` | `string` | no |

## CreatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `policyDocument` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | no |
| `policyArn` | `string` | no |
| `policyDocument` | `string` | no |
| `policyVersionId` | `string` | no |

## CreatePolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `policyDocument` | `string` | yes |
| `setAsDefault` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | no |
| `policyDocument` | `string` | no |
| `policyVersionId` | `string` | no |
| `isDefaultVersion` | `boolean` | no |

## CreateProvisioningClaim

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | no |
| `certificatePem` | `string` | no |
| `keyPair` | `KeyPair` | no |
| `expiration` | `timestamp` | no |

## CreateProvisioningTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `description` | `string` | no |
| `templateBody` | `string` | yes |
| `enabled` | `boolean` | no |
| `provisioningRoleArn` | `string` | yes |
| `preProvisioningHook` | `ProvisioningHook` | no |
| `tags` | `List<Tag>` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateArn` | `string` | no |
| `templateName` | `string` | no |
| `defaultVersionId` | `integer` | no |

## CreateProvisioningTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `templateBody` | `string` | yes |
| `setAsDefault` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateArn` | `string` | no |
| `templateName` | `string` | no |
| `versionId` | `integer` | no |
| `isDefaultVersion` | `boolean` | no |

## CreateRoleAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleAlias` | `string` | yes |
| `roleArn` | `string` | yes |
| `credentialDurationSeconds` | `integer` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleAlias` | `string` | no |
| `roleAliasArn` | `string` | no |

## CreateScheduledAudit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frequency` | `string` | yes |
| `dayOfMonth` | `string` | no |
| `dayOfWeek` | `string` | no |
| `targetCheckNames` | `List<string>` | yes |
| `scheduledAuditName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAuditArn` | `string` | no |

## CreateSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | yes |
| `securityProfileDescription` | `string` | no |
| `behaviors` | `List<Behavior>` | no |
| `alertTargets` | `Map<AlertTarget>` | no |
| `additionalMetricsToRetain` | `List<string>` | no |
| `additionalMetricsToRetainV2` | `List<MetricToRetain>` | no |
| `tags` | `List<Tag>` | no |
| `metricsExportConfig` | `MetricsExportConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | no |
| `securityProfileArn` | `string` | no |

## CreateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamId` | `string` | yes |
| `description` | `string` | no |
| `files` | `List<StreamFile>` | yes |
| `roleArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamId` | `string` | no |
| `streamArn` | `string` | no |
| `description` | `string` | no |
| `streamVersion` | `integer` | no |

## CreateThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `thingTypeName` | `string` | no |
| `attributePayload` | `AttributePayload` | no |
| `billingGroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | no |
| `thingArn` | `string` | no |
| `thingId` | `string` | no |

## CreateThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | yes |
| `parentGroupName` | `string` | no |
| `thingGroupProperties` | `ThingGroupProperties` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | no |
| `thingGroupArn` | `string` | no |
| `thingGroupId` | `string` | no |

## CreateThingType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingTypeName` | `string` | yes |
| `thingTypeProperties` | `ThingTypeProperties` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingTypeName` | `string` | no |
| `thingTypeArn` | `string` | no |
| `thingTypeId` | `string` | no |

## CreateTopicRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleName` | `string` | yes |
| `topicRulePayload` | `TopicRulePayload` | yes |
| `tags` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateTopicRuleDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationConfiguration` | `TopicRuleDestinationConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `topicRuleDestination` | `TopicRuleDestination` | no |

## DeleteAccountAuditConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deleteScheduledAudits` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAuditSuppression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkName` | `string` | yes |
| `resourceIdentifier` | `ResourceIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | yes |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCACertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |
| `forceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCertificateProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateProviderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |

## DeleteCommandExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | yes |
| `targetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDimension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomainConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDynamicThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | yes |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFleetMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | yes |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `force` | `boolean` | no |
| `namespaceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteJobExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `thingName` | `string` | yes |
| `executionNumber` | `long` | yes |
| `force` | `boolean` | no |
| `namespaceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobTemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMitigationAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOTAUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `otaUpdateId` | `string` | yes |
| `deleteStream` | `boolean` | no |
| `forceDeleteAWSJob` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePackageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `versionName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `policyVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProvisioningTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProvisioningTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `versionId` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRegistrationCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRoleAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScheduledAudit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAuditName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | yes |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | yes |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteThingType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTopicRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTopicRuleDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteV2LoggingLevel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetType` | `string` | yes |
| `targetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeprecateThingType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingTypeName` | `string` | yes |
| `undoDeprecate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountAuditConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleArn` | `string` | no |
| `auditNotificationTargetConfigurations` | `Map<AuditNotificationTarget>` | no |
| `auditCheckConfigurations` | `Map<AuditCheckConfiguration>` | no |

## DescribeAuditFinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `finding` | `AuditFinding` | no |

## DescribeAuditMitigationActionsTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskStatus` | `string` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `taskStatistics` | `Map<TaskStatisticsForAuditCheck>` | no |
| `target` | `AuditMitigationActionsTaskTarget` | no |
| `auditCheckToActionsMapping` | `Map<List<string>>` | no |
| `actionsDefinition` | `List<MitigationAction>` | no |

## DescribeAuditSuppression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkName` | `string` | yes |
| `resourceIdentifier` | `ResourceIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkName` | `string` | no |
| `resourceIdentifier` | `ResourceIdentifier` | no |
| `expirationDate` | `timestamp` | no |
| `suppressIndefinitely` | `boolean` | no |
| `description` | `string` | no |

## DescribeAuditTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskStatus` | `string` | no |
| `taskType` | `string` | no |
| `taskStartTime` | `timestamp` | no |
| `taskStatistics` | `TaskStatistics` | no |
| `scheduledAuditName` | `string` | no |
| `auditDetails` | `Map<AuditCheckDetails>` | no |

## DescribeAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerDescription` | `AuthorizerDescription` | no |

## DescribeBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | no |
| `billingGroupId` | `string` | no |
| `billingGroupArn` | `string` | no |
| `version` | `long` | no |
| `billingGroupProperties` | `BillingGroupProperties` | no |
| `billingGroupMetadata` | `BillingGroupMetadata` | no |

## DescribeCACertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateDescription` | `CACertificateDescription` | no |
| `registrationConfig` | `RegistrationConfig` | no |

## DescribeCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateDescription` | `CertificateDescription` | no |

## DescribeCertificateProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateProviderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateProviderName` | `string` | no |
| `certificateProviderArn` | `string` | no |
| `lambdaFunctionArn` | `string` | no |
| `accountDefaultForOperations` | `List<string>` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |

## DescribeCustomMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricArn` | `string` | no |
| `metricType` | `string` | no |
| `displayName` | `string` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |

## DescribeDefaultAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerDescription` | `AuthorizerDescription` | no |

## DescribeDetectMitigationActionsTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskSummary` | `DetectMitigationActionsTaskSummary` | no |

## DescribeDimension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |
| `type` | `string` | no |
| `stringValues` | `List<string>` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |

## DescribeDomainConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainConfigurationName` | `string` | no |
| `domainConfigurationArn` | `string` | no |
| `domainName` | `string` | no |
| `serverCertificates` | `List<ServerCertificateSummary>` | no |
| `authorizerConfig` | `AuthorizerConfig` | no |
| `domainConfigurationStatus` | `string` | no |
| `serviceType` | `string` | no |
| `domainType` | `string` | no |
| `lastStatusChangeDate` | `timestamp` | no |
| `tlsConfig` | `TlsConfig` | no |
| `serverCertificateConfig` | `ServerCertificateConfig` | no |
| `authenticationType` | `string` | no |
| `applicationProtocol` | `string` | no |
| `clientCertificateConfig` | `ClientCertificateConfig` | no |

## DescribeEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionType` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `kmsAccessRoleArn` | `string` | no |
| `configurationDetails` | `ConfigurationDetails` | no |
| `lastModifiedDate` | `timestamp` | no |

## DescribeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointAddress` | `string` | no |

## DescribeEventConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventConfigurations` | `Map<Configuration>` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |

## DescribeFleetMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `queryString` | `string` | no |
| `aggregationType` | `AggregationType` | no |
| `period` | `integer` | no |
| `aggregationField` | `string` | no |
| `description` | `string` | no |
| `queryVersion` | `string` | no |
| `indexName` | `string` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |
| `unit` | `string` | no |
| `version` | `long` | no |
| `metricArn` | `string` | no |

## DescribeIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexName` | `string` | no |
| `indexStatus` | `string` | no |
| `schema` | `string` | no |

## DescribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `beforeSubstitution` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentSource` | `string` | no |
| `job` | `Job` | no |

## DescribeJobExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `thingName` | `string` | yes |
| `executionNumber` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `execution` | `JobExecution` | no |

## DescribeJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobTemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobTemplateArn` | `string` | no |
| `jobTemplateId` | `string` | no |
| `description` | `string` | no |
| `documentSource` | `string` | no |
| `document` | `string` | no |
| `createdAt` | `timestamp` | no |
| `presignedUrlConfig` | `PresignedUrlConfig` | no |
| `jobExecutionsRolloutConfig` | `JobExecutionsRolloutConfig` | no |
| `abortConfig` | `AbortConfig` | no |
| `timeoutConfig` | `TimeoutConfig` | no |
| `jobExecutionsRetryConfig` | `JobExecutionsRetryConfig` | no |
| `maintenanceWindows` | `List<MaintenanceWindow>` | no |
| `destinationPackageVersions` | `List<string>` | no |

## DescribeManagedJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `templateVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | no |
| `templateArn` | `string` | no |
| `description` | `string` | no |
| `templateVersion` | `string` | no |
| `environments` | `List<string>` | no |
| `documentParameters` | `List<DocumentParameter>` | no |
| `document` | `string` | no |

## DescribeMitigationAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionName` | `string` | no |
| `actionType` | `string` | no |
| `actionArn` | `string` | no |
| `actionId` | `string` | no |
| `roleArn` | `string` | no |
| `actionParams` | `MitigationActionParams` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |

## DescribeProvisioningTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateArn` | `string` | no |
| `templateName` | `string` | no |
| `description` | `string` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |
| `defaultVersionId` | `integer` | no |
| `templateBody` | `string` | no |
| `enabled` | `boolean` | no |
| `provisioningRoleArn` | `string` | no |
| `preProvisioningHook` | `ProvisioningHook` | no |
| `type` | `string` | no |

## DescribeProvisioningTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `versionId` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versionId` | `integer` | no |
| `creationDate` | `timestamp` | no |
| `templateBody` | `string` | no |
| `isDefaultVersion` | `boolean` | no |

## DescribeRoleAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleAliasDescription` | `RoleAliasDescription` | no |

## DescribeScheduledAudit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAuditName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frequency` | `string` | no |
| `dayOfMonth` | `string` | no |
| `dayOfWeek` | `string` | no |
| `targetCheckNames` | `List<string>` | no |
| `scheduledAuditName` | `string` | no |
| `scheduledAuditArn` | `string` | no |

## DescribeSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | no |
| `securityProfileArn` | `string` | no |
| `securityProfileDescription` | `string` | no |
| `behaviors` | `List<Behavior>` | no |
| `alertTargets` | `Map<AlertTarget>` | no |
| `additionalMetricsToRetain` | `List<string>` | no |
| `additionalMetricsToRetainV2` | `List<MetricToRetain>` | no |
| `version` | `long` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |
| `metricsExportConfig` | `MetricsExportConfig` | no |

## DescribeStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamInfo` | `StreamInfo` | no |

## DescribeThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `defaultClientId` | `string` | no |
| `thingName` | `string` | no |
| `thingId` | `string` | no |
| `thingArn` | `string` | no |
| `thingTypeName` | `string` | no |
| `attributes` | `Map<string>` | no |
| `version` | `long` | no |
| `billingGroupName` | `string` | no |

## DescribeThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | no |
| `thingGroupId` | `string` | no |
| `thingGroupArn` | `string` | no |
| `version` | `long` | no |
| `thingGroupProperties` | `ThingGroupProperties` | no |
| `thingGroupMetadata` | `ThingGroupMetadata` | no |
| `indexName` | `string` | no |
| `queryString` | `string` | no |
| `queryVersion` | `string` | no |
| `status` | `string` | no |

## DescribeThingRegistrationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |
| `templateBody` | `string` | no |
| `inputFileBucket` | `string` | no |
| `inputFileKey` | `string` | no |
| `roleArn` | `string` | no |
| `status` | `string` | no |
| `message` | `string` | no |
| `successCount` | `integer` | no |
| `failureCount` | `integer` | no |
| `percentageProgress` | `integer` | no |

## DescribeThingType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingTypeName` | `string` | no |
| `thingTypeId` | `string` | no |
| `thingTypeArn` | `string` | no |
| `thingTypeProperties` | `ThingTypeProperties` | no |
| `thingTypeMetadata` | `ThingTypeMetadata` | no |

## DetachPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `target` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachPrincipalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `principal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | yes |
| `securityProfileTargetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachThingPrincipal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `principal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableTopicRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSbomFromPackageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `versionName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableTopicRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetBehaviorModelTrainingSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<BehaviorModelTrainingSummary>` | no |
| `nextToken` | `string` | no |

## GetBucketsAggregation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexName` | `string` | no |
| `queryString` | `string` | yes |
| `aggregationField` | `string` | yes |
| `queryVersion` | `string` | no |
| `bucketsAggregationType` | `BucketsAggregationType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `totalCount` | `integer` | no |
| `buckets` | `List<Bucket>` | no |

## GetCardinality

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexName` | `string` | no |
| `queryString` | `string` | yes |
| `aggregationField` | `string` | no |
| `queryVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cardinality` | `integer` | no |

## GetCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandId` | `string` | no |
| `commandArn` | `string` | no |
| `namespace` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `mandatoryParameters` | `List<CommandParameter>` | no |
| `payload` | `CommandPayload` | no |
| `payloadTemplate` | `string` | no |
| `preprocessor` | `CommandPreprocessor` | no |
| `roleArn` | `string` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `deprecated` | `boolean` | no |
| `pendingDeletion` | `boolean` | no |

## GetCommandExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | yes |
| `targetArn` | `string` | yes |
| `includeResult` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | no |
| `commandArn` | `string` | no |
| `targetArn` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `StatusReason` | no |
| `result` | `Map<CommandExecutionResult>` | no |
| `parameters` | `Map<CommandParameterValue>` | no |
| `executionTimeoutSeconds` | `long` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `startedAt` | `timestamp` | no |
| `completedAt` | `timestamp` | no |
| `timeToLive` | `timestamp` | no |

## GetEffectivePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principal` | `string` | no |
| `cognitoIdentityPoolId` | `string` | no |
| `thingName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `effectivePolicies` | `List<EffectivePolicy>` | no |

## GetIndexingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingIndexingConfiguration` | `ThingIndexingConfiguration` | no |
| `thingGroupIndexingConfiguration` | `ThingGroupIndexingConfiguration` | no |

## GetJobDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `beforeSubstitution` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `document` | `string` | no |

## GetLoggingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleArn` | `string` | no |
| `logLevel` | `string` | no |

## GetOTAUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `otaUpdateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `otaUpdateInfo` | `OTAUpdateInfo` | no |

## GetPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | no |
| `packageArn` | `string` | no |
| `description` | `string` | no |
| `defaultVersionName` | `string` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |

## GetPackageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versionUpdateByJobsConfig` | `VersionUpdateByJobsConfig` | no |

## GetPackageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `versionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageVersionArn` | `string` | no |
| `packageName` | `string` | no |
| `versionName` | `string` | no |
| `description` | `string` | no |
| `attributes` | `Map<string>` | no |
| `artifact` | `PackageVersionArtifact` | no |
| `status` | `string` | no |
| `errorReason` | `string` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |
| `sbom` | `Sbom` | no |
| `sbomValidationStatus` | `string` | no |
| `recipe` | `string` | no |

## GetPercentiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexName` | `string` | no |
| `queryString` | `string` | yes |
| `aggregationField` | `string` | no |
| `queryVersion` | `string` | no |
| `percents` | `List<double>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `percentiles` | `List<PercentPair>` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | no |
| `policyArn` | `string` | no |
| `policyDocument` | `string` | no |
| `defaultVersionId` | `string` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |
| `generationId` | `string` | no |

## GetPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `policyVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | no |
| `policyName` | `string` | no |
| `policyDocument` | `string` | no |
| `policyVersionId` | `string` | no |
| `isDefaultVersion` | `boolean` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |
| `generationId` | `string` | no |

## GetRegistrationCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registrationCode` | `string` | no |

## GetStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexName` | `string` | no |
| `queryString` | `string` | yes |
| `aggregationField` | `string` | no |
| `queryVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statistics` | `Statistics` | no |

## GetThingConnectivityData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `includeSocketInformation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | no |
| `connected` | `boolean` | no |
| `timestamp` | `timestamp` | no |
| `disconnectReason` | `string` | no |
| `sourceIp` | `string` | no |
| `sourcePort` | `integer` | no |
| `targetIp` | `string` | no |
| `targetPort` | `integer` | no |
| `vpcEndpointId` | `string` | no |
| `keepAliveDuration` | `integer` | no |
| `cleanSession` | `boolean` | no |
| `sessionExpiry` | `long` | no |
| `clientId` | `string` | no |

## GetTopicRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleArn` | `string` | no |
| `rule` | `TopicRule` | no |

## GetTopicRuleDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `topicRuleDestination` | `TopicRuleDestination` | no |

## GetV2LoggingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `verbose` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleArn` | `string` | no |
| `defaultLogLevel` | `string` | no |
| `disableAllLogs` | `boolean` | no |
| `eventConfigurations` | `List<LogEventConfiguration>` | no |

## ListActiveViolations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | no |
| `securityProfileName` | `string` | no |
| `behaviorCriteriaType` | `string` | no |
| `listSuppressedAlerts` | `boolean` | no |
| `verificationState` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activeViolations` | `List<ActiveViolation>` | no |
| `nextToken` | `string` | no |

## ListAttachedPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `target` | `string` | yes |
| `recursive` | `boolean` | no |
| `marker` | `string` | no |
| `pageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policies` | `List<Policy>` | no |
| `nextMarker` | `string` | no |

## ListAuditFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |
| `checkName` | `string` | no |
| `resourceIdentifier` | `ResourceIdentifier` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `listSuppressedFindings` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<AuditFinding>` | no |
| `nextToken` | `string` | no |

## ListAuditMitigationActionsExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |
| `actionStatus` | `string` | no |
| `findingId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionsExecutions` | `List<AuditMitigationActionExecutionMetadata>` | no |
| `nextToken` | `string` | no |

## ListAuditMitigationActionsTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `auditTaskId` | `string` | no |
| `findingId` | `string` | no |
| `taskStatus` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<AuditMitigationActionsTaskMetadata>` | no |
| `nextToken` | `string` | no |

## ListAuditSuppressions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkName` | `string` | no |
| `resourceIdentifier` | `ResourceIdentifier` | no |
| `ascendingOrder` | `boolean` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suppressions` | `List<AuditSuppression>` | no |
| `nextToken` | `string` | no |

## ListAuditTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `taskType` | `string` | no |
| `taskStatus` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<AuditTaskMetadata>` | no |
| `nextToken` | `string` | no |

## ListAuthorizers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageSize` | `integer` | no |
| `marker` | `string` | no |
| `ascendingOrder` | `boolean` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizers` | `List<AuthorizerSummary>` | no |
| `nextMarker` | `string` | no |

## ListBillingGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `namePrefixFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroups` | `List<GroupNameAndArn>` | no |
| `nextToken` | `string` | no |

## ListCACertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageSize` | `integer` | no |
| `marker` | `string` | no |
| `ascendingOrder` | `boolean` | no |
| `templateName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificates` | `List<CACertificate>` | no |
| `nextMarker` | `string` | no |

## ListCertificateProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateProviders` | `List<CertificateProviderSummary>` | no |
| `nextToken` | `string` | no |

## ListCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageSize` | `integer` | no |
| `marker` | `string` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificates` | `List<Certificate>` | no |
| `nextMarker` | `string` | no |

## ListCertificatesByCA

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caCertificateId` | `string` | yes |
| `pageSize` | `integer` | no |
| `marker` | `string` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificates` | `List<Certificate>` | no |
| `nextMarker` | `string` | no |

## ListCommandExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `namespace` | `string` | no |
| `status` | `string` | no |
| `sortOrder` | `string` | no |
| `startedTimeFilter` | `TimeFilter` | no |
| `completedTimeFilter` | `TimeFilter` | no |
| `targetArn` | `string` | no |
| `commandArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandExecutions` | `List<CommandExecutionSummary>` | no |
| `nextToken` | `string` | no |

## ListCommands

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `namespace` | `string` | no |
| `commandParameterName` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commands` | `List<CommandSummary>` | no |
| `nextToken` | `string` | no |

## ListCustomMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricNames` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListDetectMitigationActionsExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |
| `violationId` | `string` | no |
| `thingName` | `string` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionsExecutions` | `List<DetectMitigationActionExecution>` | no |
| `nextToken` | `string` | no |

## ListDetectMitigationActionsTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<DetectMitigationActionsTaskSummary>` | no |
| `nextToken` | `string` | no |

## ListDimensions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dimensionNames` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListDomainConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marker` | `string` | no |
| `pageSize` | `integer` | no |
| `serviceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainConfigurations` | `List<DomainConfigurationSummary>` | no |
| `nextMarker` | `string` | no |

## ListFleetMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetMetrics` | `List<FleetMetricNameAndArn>` | no |
| `nextToken` | `string` | no |

## ListIndices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexNames` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListJobExecutionsForJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionSummaries` | `List<JobExecutionSummaryForJob>` | no |
| `nextToken` | `string` | no |

## ListJobExecutionsForThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `status` | `string` | no |
| `namespaceId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `jobId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionSummaries` | `List<JobExecutionSummaryForThing>` | no |
| `nextToken` | `string` | no |

## ListJobTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobTemplates` | `List<JobTemplateSummary>` | no |
| `nextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `targetSelection` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `thingGroupName` | `string` | no |
| `thingGroupId` | `string` | no |
| `namespaceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<JobSummary>` | no |
| `nextToken` | `string` | no |

## ListManagedJobTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managedJobTemplates` | `List<ManagedJobTemplateSummary>` | no |
| `nextToken` | `string` | no |

## ListMetricValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `metricName` | `string` | yes |
| `dimensionName` | `string` | no |
| `dimensionValueOperator` | `string` | no |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricDatumList` | `List<MetricDatum>` | no |
| `nextToken` | `string` | no |

## ListMitigationActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionType` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionIdentifiers` | `List<MitigationActionIdentifier>` | no |
| `nextToken` | `string` | no |

## ListOTAUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `otaUpdateStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `otaUpdates` | `List<OTAUpdateSummary>` | no |
| `nextToken` | `string` | no |

## ListOutgoingCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageSize` | `integer` | no |
| `marker` | `string` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `outgoingCertificates` | `List<OutgoingCertificate>` | no |
| `nextMarker` | `string` | no |

## ListPackageVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageVersionSummaries` | `List<PackageVersionSummary>` | no |
| `nextToken` | `string` | no |

## ListPackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageSummaries` | `List<PackageSummary>` | no |
| `nextToken` | `string` | no |

## ListPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marker` | `string` | no |
| `pageSize` | `integer` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policies` | `List<Policy>` | no |
| `nextMarker` | `string` | no |

## ListPolicyPrincipals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `marker` | `string` | no |
| `pageSize` | `integer` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principals` | `List<string>` | no |
| `nextMarker` | `string` | no |

## ListPolicyVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyVersions` | `List<PolicyVersion>` | no |

## ListPrincipalPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principal` | `string` | yes |
| `marker` | `string` | no |
| `pageSize` | `integer` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policies` | `List<Policy>` | no |
| `nextMarker` | `string` | no |

## ListPrincipalThings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `principal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `things` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListPrincipalThingsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `principal` | `string` | yes |
| `thingPrincipalType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principalThingObjects` | `List<PrincipalThingObject>` | no |
| `nextToken` | `string` | no |

## ListProvisioningTemplateVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versions` | `List<ProvisioningTemplateVersionSummary>` | no |
| `nextToken` | `string` | no |

## ListProvisioningTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templates` | `List<ProvisioningTemplateSummary>` | no |
| `nextToken` | `string` | no |

## ListRelatedResourcesForAuditFinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relatedResources` | `List<RelatedResource>` | no |
| `nextToken` | `string` | no |

## ListRoleAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pageSize` | `integer` | no |
| `marker` | `string` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleAliases` | `List<string>` | no |
| `nextMarker` | `string` | no |

## ListSbomValidationResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `versionName` | `string` | yes |
| `validationResult` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `validationResultSummaries` | `List<SbomValidationResultSummary>` | no |
| `nextToken` | `string` | no |

## ListScheduledAudits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAudits` | `List<ScheduledAuditMetadata>` | no |
| `nextToken` | `string` | no |

## ListSecurityProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `dimensionName` | `string` | no |
| `metricName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileIdentifiers` | `List<SecurityProfileIdentifier>` | no |
| `nextToken` | `string` | no |

## ListSecurityProfilesForTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `recursive` | `boolean` | no |
| `securityProfileTargetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileTargetMappings` | `List<SecurityProfileTargetMapping>` | no |
| `nextToken` | `string` | no |

## ListStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `ascendingOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streams` | `List<StreamSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |
| `nextToken` | `string` | no |

## ListTargetsForPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `marker` | `string` | no |
| `pageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targets` | `List<string>` | no |
| `nextMarker` | `string` | no |

## ListTargetsForSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileTargets` | `List<SecurityProfileTarget>` | no |
| `nextToken` | `string` | no |

## ListThingGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `parentGroup` | `string` | no |
| `namePrefixFilter` | `string` | no |
| `recursive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroups` | `List<GroupNameAndArn>` | no |
| `nextToken` | `string` | no |

## ListThingGroupsForThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroups` | `List<GroupNameAndArn>` | no |
| `nextToken` | `string` | no |

## ListThingPrincipals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `thingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principals` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListThingPrincipalsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `thingName` | `string` | yes |
| `thingPrincipalType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingPrincipalObjects` | `List<ThingPrincipalObject>` | no |
| `nextToken` | `string` | no |

## ListThingRegistrationTaskReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |
| `reportType` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceLinks` | `List<string>` | no |
| `reportType` | `string` | no |
| `nextToken` | `string` | no |

## ListThingRegistrationTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskIds` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListThingTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `thingTypeName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingTypes` | `List<ThingTypeDefinition>` | no |
| `nextToken` | `string` | no |

## ListThings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `attributeName` | `string` | no |
| `attributeValue` | `string` | no |
| `thingTypeName` | `string` | no |
| `usePrefixAttributeValue` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `things` | `List<ThingAttribute>` | no |
| `nextToken` | `string` | no |

## ListThingsInBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `things` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListThingsInThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | yes |
| `recursive` | `boolean` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `things` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListTopicRuleDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationSummaries` | `List<TopicRuleDestinationSummary>` | no |
| `nextToken` | `string` | no |

## ListTopicRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `topic` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `ruleDisabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rules` | `List<TopicRuleListItem>` | no |
| `nextToken` | `string` | no |

## ListV2LoggingLevels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logTargetConfigurations` | `List<LogTargetConfiguration>` | no |
| `nextToken` | `string` | no |

## ListViolationEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `thingName` | `string` | no |
| `securityProfileName` | `string` | no |
| `behaviorCriteriaType` | `string` | no |
| `listSuppressedAlerts` | `boolean` | no |
| `verificationState` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `violationEvents` | `List<ViolationEvent>` | no |
| `nextToken` | `string` | no |

## PutVerificationStateOnViolation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `violationId` | `string` | yes |
| `verificationState` | `string` | yes |
| `verificationStateDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterCACertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caCertificate` | `string` | yes |
| `verificationCertificate` | `string` | no |
| `setAsActive` | `boolean` | no |
| `allowAutoRegistration` | `boolean` | no |
| `registrationConfig` | `RegistrationConfig` | no |
| `tags` | `List<Tag>` | no |
| `certificateMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateArn` | `string` | no |
| `certificateId` | `string` | no |

## RegisterCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificatePem` | `string` | yes |
| `caCertificatePem` | `string` | no |
| `setAsActive` | `boolean` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateArn` | `string` | no |
| `certificateId` | `string` | no |

## RegisterCertificateWithoutCA

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificatePem` | `string` | yes |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateArn` | `string` | no |
| `certificateId` | `string` | no |

## RegisterThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateBody` | `string` | yes |
| `parameters` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificatePem` | `string` | no |
| `resourceArns` | `Map<string>` | no |

## RejectCertificateTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |
| `rejectReason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveThingFromBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | no |
| `billingGroupArn` | `string` | no |
| `thingName` | `string` | no |
| `thingArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveThingFromThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | no |
| `thingGroupArn` | `string` | no |
| `thingName` | `string` | no |
| `thingArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReplaceTopicRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleName` | `string` | yes |
| `topicRulePayload` | `TopicRulePayload` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexName` | `string` | no |
| `queryString` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `queryVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `things` | `List<ThingDocument>` | no |
| `thingGroups` | `List<ThingGroupDocument>` | no |

## SetDefaultAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | no |
| `authorizerArn` | `string` | no |

## SetDefaultPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `policyVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetLoggingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loggingOptionsPayload` | `LoggingOptionsPayload` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetV2LoggingLevel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logTarget` | `LogTarget` | yes |
| `logLevel` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetV2LoggingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleArn` | `string` | no |
| `defaultLogLevel` | `string` | no |
| `disableAllLogs` | `boolean` | no |
| `eventConfigurations` | `List<LogEventConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAuditMitigationActionsTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |
| `target` | `AuditMitigationActionsTaskTarget` | yes |
| `auditCheckToActionsMapping` | `Map<List<string>>` | yes |
| `clientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |

## StartDetectMitigationActionsTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |
| `target` | `DetectMitigationActionsTaskTarget` | yes |
| `actions` | `List<string>` | yes |
| `violationEventOccurrenceRange` | `ViolationEventOccurrenceRange` | no |
| `includeOnlyActiveViolations` | `boolean` | no |
| `includeSuppressedAlerts` | `boolean` | no |
| `clientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |

## StartOnDemandAuditTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetCheckNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |

## StartThingRegistrationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateBody` | `string` | yes |
| `inputFileBucket` | `string` | yes |
| `inputFileKey` | `string` | yes |
| `roleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |

## StopThingRegistrationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principal` | `string` | no |
| `cognitoIdentityPoolId` | `string` | no |
| `authInfos` | `List<AuthInfo>` | yes |
| `clientId` | `string` | no |
| `policyNamesToAdd` | `List<string>` | no |
| `policyNamesToSkip` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authResults` | `List<AuthResult>` | no |

## TestInvokeAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | yes |
| `token` | `string` | no |
| `tokenSignature` | `string` | no |
| `httpContext` | `HttpContext` | no |
| `mqttContext` | `MqttContext` | no |
| `tlsContext` | `TlsContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `isAuthenticated` | `boolean` | no |
| `principalId` | `string` | no |
| `policyDocuments` | `List<string>` | no |
| `refreshAfterInSeconds` | `integer` | no |
| `disconnectAfterInSeconds` | `integer` | no |

## TransferCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |
| `targetAwsAccount` | `string` | yes |
| `transferMessage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transferredCertificateArn` | `string` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccountAuditConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleArn` | `string` | no |
| `auditNotificationTargetConfigurations` | `Map<AuditNotificationTarget>` | no |
| `auditCheckConfigurations` | `Map<AuditCheckConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAuditSuppression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checkName` | `string` | yes |
| `resourceIdentifier` | `ResourceIdentifier` | yes |
| `expirationDate` | `timestamp` | no |
| `suppressIndefinitely` | `boolean` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | yes |
| `authorizerFunctionArn` | `string` | no |
| `tokenKeyName` | `string` | no |
| `tokenSigningPublicKeys` | `Map<string>` | no |
| `status` | `string` | no |
| `enableCachingForHttp` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizerName` | `string` | no |
| `authorizerArn` | `string` | no |

## UpdateBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingGroupName` | `string` | yes |
| `billingGroupProperties` | `BillingGroupProperties` | yes |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `long` | no |

## UpdateCACertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |
| `newStatus` | `string` | no |
| `newAutoRegistrationStatus` | `string` | no |
| `registrationConfig` | `RegistrationConfig` | no |
| `removeAutoRegistration` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateId` | `string` | yes |
| `newStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCertificateProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateProviderName` | `string` | yes |
| `lambdaFunctionArn` | `string` | no |
| `accountDefaultForOperations` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateProviderName` | `string` | no |
| `certificateProviderArn` | `string` | no |

## UpdateCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandId` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `deprecated` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandId` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `deprecated` | `boolean` | no |
| `lastUpdatedAt` | `timestamp` | no |

## UpdateCustomMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | yes |
| `displayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | no |
| `metricArn` | `string` | no |
| `metricType` | `string` | no |
| `displayName` | `string` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |

## UpdateDimension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `stringValues` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |
| `type` | `string` | no |
| `stringValues` | `List<string>` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |

## UpdateDomainConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainConfigurationName` | `string` | yes |
| `authorizerConfig` | `AuthorizerConfig` | no |
| `domainConfigurationStatus` | `string` | no |
| `removeAuthorizerConfig` | `boolean` | no |
| `tlsConfig` | `TlsConfig` | no |
| `serverCertificateConfig` | `ServerCertificateConfig` | no |
| `authenticationType` | `string` | no |
| `applicationProtocol` | `string` | no |
| `clientCertificateConfig` | `ClientCertificateConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainConfigurationName` | `string` | no |
| `domainConfigurationArn` | `string` | no |

## UpdateDynamicThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | yes |
| `thingGroupProperties` | `ThingGroupProperties` | yes |
| `expectedVersion` | `long` | no |
| `indexName` | `string` | no |
| `queryString` | `string` | no |
| `queryVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `long` | no |

## UpdateEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionType` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `kmsAccessRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEventConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventConfigurations` | `Map<Configuration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFleetMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricName` | `string` | yes |
| `queryString` | `string` | no |
| `aggregationType` | `AggregationType` | no |
| `period` | `integer` | no |
| `aggregationField` | `string` | no |
| `description` | `string` | no |
| `queryVersion` | `string` | no |
| `indexName` | `string` | yes |
| `unit` | `string` | no |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateIndexingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingIndexingConfiguration` | `ThingIndexingConfiguration` | no |
| `thingGroupIndexingConfiguration` | `ThingGroupIndexingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `description` | `string` | no |
| `presignedUrlConfig` | `PresignedUrlConfig` | no |
| `jobExecutionsRolloutConfig` | `JobExecutionsRolloutConfig` | no |
| `abortConfig` | `AbortConfig` | no |
| `timeoutConfig` | `TimeoutConfig` | no |
| `namespaceId` | `string` | no |
| `jobExecutionsRetryConfig` | `JobExecutionsRetryConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMitigationAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionName` | `string` | yes |
| `roleArn` | `string` | no |
| `actionParams` | `MitigationActionParams` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionArn` | `string` | no |
| `actionId` | `string` | no |

## UpdatePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `description` | `string` | no |
| `defaultVersionName` | `string` | no |
| `unsetDefaultVersion` | `boolean` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePackageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versionUpdateByJobsConfig` | `VersionUpdateByJobsConfig` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePackageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageName` | `string` | yes |
| `versionName` | `string` | yes |
| `description` | `string` | no |
| `attributes` | `Map<string>` | no |
| `artifact` | `PackageVersionArtifact` | no |
| `action` | `string` | no |
| `recipe` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateProvisioningTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `description` | `string` | no |
| `enabled` | `boolean` | no |
| `defaultVersionId` | `integer` | no |
| `provisioningRoleArn` | `string` | no |
| `preProvisioningHook` | `ProvisioningHook` | no |
| `removePreProvisioningHook` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoleAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleAlias` | `string` | yes |
| `roleArn` | `string` | no |
| `credentialDurationSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleAlias` | `string` | no |
| `roleAliasArn` | `string` | no |

## UpdateScheduledAudit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frequency` | `string` | no |
| `dayOfMonth` | `string` | no |
| `dayOfWeek` | `string` | no |
| `targetCheckNames` | `List<string>` | no |
| `scheduledAuditName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledAuditArn` | `string` | no |

## UpdateSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | yes |
| `securityProfileDescription` | `string` | no |
| `behaviors` | `List<Behavior>` | no |
| `alertTargets` | `Map<AlertTarget>` | no |
| `additionalMetricsToRetain` | `List<string>` | no |
| `additionalMetricsToRetainV2` | `List<MetricToRetain>` | no |
| `deleteBehaviors` | `boolean` | no |
| `deleteAlertTargets` | `boolean` | no |
| `deleteAdditionalMetricsToRetain` | `boolean` | no |
| `expectedVersion` | `long` | no |
| `metricsExportConfig` | `MetricsExportConfig` | no |
| `deleteMetricsExportConfig` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityProfileName` | `string` | no |
| `securityProfileArn` | `string` | no |
| `securityProfileDescription` | `string` | no |
| `behaviors` | `List<Behavior>` | no |
| `alertTargets` | `Map<AlertTarget>` | no |
| `additionalMetricsToRetain` | `List<string>` | no |
| `additionalMetricsToRetainV2` | `List<MetricToRetain>` | no |
| `version` | `long` | no |
| `creationDate` | `timestamp` | no |
| `lastModifiedDate` | `timestamp` | no |
| `metricsExportConfig` | `MetricsExportConfig` | no |

## UpdateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamId` | `string` | yes |
| `description` | `string` | no |
| `files` | `List<StreamFile>` | no |
| `roleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamId` | `string` | no |
| `streamArn` | `string` | no |
| `description` | `string` | no |
| `streamVersion` | `integer` | no |

## UpdateThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `thingTypeName` | `string` | no |
| `attributePayload` | `AttributePayload` | no |
| `expectedVersion` | `long` | no |
| `removeThingType` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateThingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingGroupName` | `string` | yes |
| `thingGroupProperties` | `ThingGroupProperties` | yes |
| `expectedVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `long` | no |

## UpdateThingGroupsForThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | no |
| `thingGroupsToAdd` | `List<string>` | no |
| `thingGroupsToRemove` | `List<string>` | no |
| `overrideDynamicGroups` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateThingType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingTypeName` | `string` | yes |
| `thingTypeProperties` | `ThingTypeProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTopicRuleDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ValidateSecurityProfileBehaviors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `behaviors` | `List<Behavior>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `valid` | `boolean` | no |
| `validationErrors` | `List<ValidationError>` | no |

