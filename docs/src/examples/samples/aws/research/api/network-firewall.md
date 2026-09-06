# AWS Network Firewall

API version: 2020-11-12. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/network-firewall/2020-11-12/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptNetworkFirewallTransitGatewayAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `TransitGatewayAttachmentStatus` | `string` | yes |

## AssociateAvailabilityZones

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `AvailabilityZoneMappings` | `List<AvailabilityZoneMapping>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `AvailabilityZoneMappings` | `List<AvailabilityZoneMapping>` | no |
| `UpdateToken` | `string` | no |

## AssociateFirewallPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `FirewallPolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `FirewallPolicyArn` | `string` | no |
| `UpdateToken` | `string` | no |

## AssociateSubnets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `SubnetMappings` | `List<SubnetMapping>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `SubnetMappings` | `List<SubnetMapping>` | no |
| `UpdateToken` | `string` | no |

## AttachRuleGroupsToProxyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurationName` | `string` | no |
| `ProxyConfigurationArn` | `string` | no |
| `RuleGroups` | `List<ProxyRuleGroupAttachment>` | yes |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfiguration` | `ProxyConfiguration` | no |
| `UpdateToken` | `string` | no |

## CreateContainerAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociationName` | `string` | yes |
| `Description` | `string` | no |
| `Type` | `string` | yes |
| `ContainerMonitoringConfigurations` | `List<ContainerMonitoringConfiguration>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociationName` | `string` | no |
| `ContainerAssociationArn` | `string` | no |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `ContainerMonitoringConfigurations` | `List<ContainerMonitoringConfiguration>` | no |
| `Status` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `UpdateToken` | `string` | no |

## CreateFirewall

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallName` | `string` | yes |
| `FirewallPolicyArn` | `string` | yes |
| `VpcId` | `string` | no |
| `SubnetMappings` | `List<SubnetMapping>` | no |
| `DeleteProtection` | `boolean` | no |
| `SubnetChangeProtection` | `boolean` | no |
| `FirewallPolicyChangeProtection` | `boolean` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `EnabledAnalysisTypes` | `List<string>` | no |
| `TransitGatewayId` | `string` | no |
| `AvailabilityZoneMappings` | `List<AvailabilityZoneMapping>` | no |
| `AvailabilityZoneChangeProtection` | `boolean` | no |
| `NatGatewayMappings` | `List<NatGatewayMapping>` | no |
| `ProxySettings` | `ProxySettings` | no |
| `NoSourcePreservation` | `boolean` | no |
| `VpcEndpoint` | `VpcEndpoint` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Firewall` | `Firewall` | no |
| `FirewallStatus` | `FirewallStatus` | no |

## CreateFirewallPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallPolicyName` | `string` | yes |
| `FirewallPolicy` | `FirewallPolicy` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DryRun` | `boolean` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `FirewallPolicyResponse` | `FirewallPolicyResponse` | yes |

## CreateProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyName` | `string` | yes |
| `NatGatewayId` | `string` | yes |
| `ProxyConfigurationName` | `string` | no |
| `ProxyConfigurationArn` | `string` | no |
| `ListenerProperties` | `List<ListenerPropertyRequest>` | no |
| `TlsInterceptProperties` | `TlsInterceptPropertiesRequest` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Proxy` | `Proxy` | no |
| `UpdateToken` | `string` | no |

## CreateProxyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurationName` | `string` | yes |
| `Description` | `string` | no |
| `RuleGroupNames` | `List<string>` | no |
| `RuleGroupArns` | `List<string>` | no |
| `DefaultRulePhaseActions` | `ProxyConfigDefaultRulePhaseActionsRequest` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfiguration` | `ProxyConfiguration` | no |
| `UpdateToken` | `string` | no |

## CreateProxyRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupName` | `string` | yes |
| `Description` | `string` | no |
| `Rules` | `ProxyRulesByRequestPhase` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroup` | `ProxyRuleGroup` | no |
| `UpdateToken` | `string` | no |

## CreateProxyRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupArn` | `string` | no |
| `ProxyRuleGroupName` | `string` | no |
| `Rules` | `CreateProxyRulesByRequestPhase` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroup` | `ProxyRuleGroup` | no |
| `UpdateToken` | `string` | no |

## CreateRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupName` | `string` | yes |
| `RuleGroup` | `RuleGroup` | no |
| `Rules` | `string` | no |
| `Type` | `string` | yes |
| `Description` | `string` | no |
| `Capacity` | `integer` | yes |
| `Tags` | `List<Tag>` | no |
| `DryRun` | `boolean` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `SourceMetadata` | `SourceMetadata` | no |
| `AnalyzeRuleGroup` | `boolean` | no |
| `SummaryConfiguration` | `SummaryConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `RuleGroupResponse` | `RuleGroupResponse` | yes |

## CreateTLSInspectionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TLSInspectionConfigurationName` | `string` | yes |
| `TLSInspectionConfiguration` | `TLSInspectionConfiguration` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `TLSInspectionConfigurationResponse` | `TLSInspectionConfigurationResponse` | yes |

## CreateVpcEndpointAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | yes |
| `VpcId` | `string` | yes |
| `SubnetMapping` | `SubnetMapping` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointAssociation` | `VpcEndpointAssociation` | no |
| `VpcEndpointAssociationStatus` | `VpcEndpointAssociationStatus` | no |

## DeleteContainerAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociationName` | `string` | no |
| `ContainerAssociationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociationName` | `string` | no |
| `ContainerAssociationArn` | `string` | no |
| `Status` | `string` | no |

## DeleteFirewall

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallName` | `string` | no |
| `FirewallArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Firewall` | `Firewall` | no |
| `FirewallStatus` | `FirewallStatus` | no |

## DeleteFirewallPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallPolicyName` | `string` | no |
| `FirewallPolicyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallPolicyResponse` | `FirewallPolicyResponse` | yes |

## DeleteNetworkFirewallTransitGatewayAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `TransitGatewayAttachmentStatus` | `string` | yes |

## DeleteProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | yes |
| `ProxyName` | `string` | no |
| `ProxyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | no |
| `ProxyName` | `string` | no |
| `ProxyArn` | `string` | no |

## DeleteProxyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurationName` | `string` | no |
| `ProxyConfigurationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurationName` | `string` | no |
| `ProxyConfigurationArn` | `string` | no |

## DeleteProxyRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupName` | `string` | no |
| `ProxyRuleGroupArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupName` | `string` | no |
| `ProxyRuleGroupArn` | `string` | no |

## DeleteProxyRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupArn` | `string` | no |
| `ProxyRuleGroupName` | `string` | no |
| `Rules` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroup` | `ProxyRuleGroup` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupName` | `string` | no |
| `RuleGroupArn` | `string` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupResponse` | `RuleGroupResponse` | yes |

## DeleteTLSInspectionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TLSInspectionConfigurationArn` | `string` | no |
| `TLSInspectionConfigurationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TLSInspectionConfigurationResponse` | `TLSInspectionConfigurationResponse` | yes |

## DeleteVpcEndpointAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointAssociationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointAssociation` | `VpcEndpointAssociation` | no |
| `VpcEndpointAssociationStatus` | `VpcEndpointAssociationStatus` | no |

## DescribeContainerAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociationName` | `string` | no |
| `ContainerAssociationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociationName` | `string` | no |
| `ContainerAssociationArn` | `string` | no |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `ContainerMonitoringConfigurations` | `List<ContainerMonitoringConfiguration>` | no |
| `Status` | `string` | no |
| `ResolvedCidrCount` | `integer` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |
| `UpdateToken` | `string` | no |

## DescribeFirewall

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallName` | `string` | no |
| `FirewallArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `Firewall` | `Firewall` | no |
| `FirewallStatus` | `FirewallStatus` | no |

## DescribeFirewallMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallPolicyArn` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `SupportedAvailabilityZones` | `Map<AvailabilityZoneMetadata>` | no |
| `TransitGatewayAttachmentId` | `string` | no |

## DescribeFirewallPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallPolicyName` | `string` | no |
| `FirewallPolicyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `FirewallPolicyResponse` | `FirewallPolicyResponse` | yes |
| `FirewallPolicy` | `FirewallPolicy` | no |

## DescribeFlowOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | yes |
| `AvailabilityZone` | `string` | no |
| `VpcEndpointAssociationArn` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `FlowOperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `VpcEndpointAssociationArn` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `FlowOperationId` | `string` | no |
| `FlowOperationType` | `string` | no |
| `FlowOperationStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `FlowRequestTimestamp` | `timestamp` | no |
| `FlowOperation` | `FlowOperation` | no |

## DescribeLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `LoggingConfiguration` | `LoggingConfiguration` | no |
| `EnableMonitoringDashboard` | `boolean` | no |

## DescribeProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyName` | `string` | no |
| `ProxyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Proxy` | `DescribeProxyResource` | no |
| `UpdateToken` | `string` | no |

## DescribeProxyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurationName` | `string` | no |
| `ProxyConfigurationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfiguration` | `ProxyConfiguration` | no |
| `UpdateToken` | `string` | no |

## DescribeProxyRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleName` | `string` | yes |
| `ProxyRuleGroupName` | `string` | no |
| `ProxyRuleGroupArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRule` | `ProxyRule` | no |
| `UpdateToken` | `string` | no |

## DescribeProxyRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupName` | `string` | no |
| `ProxyRuleGroupArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroup` | `ProxyRuleGroup` | no |
| `UpdateToken` | `string` | no |

## DescribeResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## DescribeRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupName` | `string` | no |
| `RuleGroupArn` | `string` | no |
| `Type` | `string` | no |
| `AnalyzeRuleGroup` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `RuleGroup` | `RuleGroup` | no |
| `RuleGroupResponse` | `RuleGroupResponse` | yes |

## DescribeRuleGroupMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupName` | `string` | no |
| `RuleGroupArn` | `string` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupArn` | `string` | yes |
| `RuleGroupName` | `string` | yes |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `Capacity` | `integer` | no |
| `StatefulRuleOptions` | `StatefulRuleOptions` | no |
| `LastModifiedTime` | `timestamp` | no |
| `VendorName` | `string` | no |
| `ProductId` | `string` | no |
| `ListingName` | `string` | no |

## DescribeRuleGroupSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupName` | `string` | no |
| `RuleGroupArn` | `string` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupName` | `string` | yes |
| `Description` | `string` | no |
| `Summary` | `Summary` | no |

## DescribeTLSInspectionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TLSInspectionConfigurationArn` | `string` | no |
| `TLSInspectionConfigurationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `TLSInspectionConfiguration` | `TLSInspectionConfiguration` | no |
| `TLSInspectionConfigurationResponse` | `TLSInspectionConfigurationResponse` | yes |

## DescribeVpcEndpointAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointAssociationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointAssociation` | `VpcEndpointAssociation` | no |
| `VpcEndpointAssociationStatus` | `VpcEndpointAssociationStatus` | no |

## DetachRuleGroupsFromProxyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurationName` | `string` | no |
| `ProxyConfigurationArn` | `string` | no |
| `RuleGroupNames` | `List<string>` | no |
| `RuleGroupArns` | `List<string>` | no |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfiguration` | `ProxyConfiguration` | no |
| `UpdateToken` | `string` | no |

## DisassociateAvailabilityZones

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `AvailabilityZoneMappings` | `List<AvailabilityZoneMapping>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `AvailabilityZoneMappings` | `List<AvailabilityZoneMapping>` | no |
| `UpdateToken` | `string` | no |

## DisassociateSubnets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `SubnetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `SubnetMappings` | `List<SubnetMapping>` | no |
| `UpdateToken` | `string` | no |

## GetAnalysisReportResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallName` | `string` | no |
| `AnalysisReportId` | `string` | yes |
| `FirewallArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `ReportTime` | `timestamp` | no |
| `AnalysisType` | `string` | no |
| `NextToken` | `string` | no |
| `AnalysisReportResults` | `List<AnalysisTypeReportResult>` | no |

## ListAnalysisReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallName` | `string` | no |
| `FirewallArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisReports` | `List<AnalysisReport>` | no |
| `NextToken` | `string` | no |

## ListContainerAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociations` | `List<ContainerAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListFirewallPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `FirewallPolicies` | `List<FirewallPolicyMetadata>` | no |

## ListFirewalls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `VpcIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Firewalls` | `List<FirewallMetadata>` | no |

## ListFlowOperationResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | yes |
| `FlowOperationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AvailabilityZone` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `VpcEndpointAssociationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `VpcEndpointAssociationArn` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `FlowOperationId` | `string` | no |
| `FlowOperationStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `FlowRequestTimestamp` | `timestamp` | no |
| `Flows` | `List<Flow>` | no |
| `NextToken` | `string` | no |

## ListFlowOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | yes |
| `AvailabilityZone` | `string` | no |
| `VpcEndpointAssociationArn` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `FlowOperationType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowOperations` | `List<FlowOperationMetadata>` | no |
| `NextToken` | `string` | no |

## ListProxies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Proxies` | `List<ProxyMetadata>` | no |
| `NextToken` | `string` | no |

## ListProxyConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurations` | `List<ProxyConfigurationMetadata>` | no |
| `NextToken` | `string` | no |

## ListProxyRuleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroups` | `List<ProxyRuleGroupMetadata>` | no |
| `NextToken` | `string` | no |

## ListRuleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Scope` | `string` | no |
| `ManagedType` | `string` | no |
| `SubscriptionStatus` | `string` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RuleGroups` | `List<RuleGroupMetadata>` | no |

## ListTLSInspectionConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TLSInspectionConfigurations` | `List<TLSInspectionConfigurationMetadata>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

## ListVpcEndpointAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `FirewallArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `VpcEndpointAssociations` | `List<VpcEndpointAssociationMetadata>` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RejectNetworkFirewallTransitGatewayAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `TransitGatewayAttachmentStatus` | `string` | yes |

## StartAnalysisReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallName` | `string` | no |
| `FirewallArn` | `string` | no |
| `AnalysisType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisReportId` | `string` | yes |

## StartFlowCapture

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | yes |
| `AvailabilityZone` | `string` | no |
| `VpcEndpointAssociationArn` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `MinimumFlowAgeInSeconds` | `integer` | no |
| `FlowFilters` | `List<FlowFilter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FlowOperationId` | `string` | no |
| `FlowOperationStatus` | `string` | no |

## StartFlowFlush

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | yes |
| `AvailabilityZone` | `string` | no |
| `VpcEndpointAssociationArn` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `MinimumFlowAgeInSeconds` | `integer` | no |
| `FlowFilters` | `List<FlowFilter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FlowOperationId` | `string` | no |
| `FlowOperationStatus` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAvailabilityZoneChangeProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `AvailabilityZoneChangeProtection` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `AvailabilityZoneChangeProtection` | `boolean` | no |

## UpdateContainerAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociationName` | `string` | no |
| `ContainerAssociationArn` | `string` | no |
| `Description` | `string` | no |
| `Type` | `string` | yes |
| `ContainerMonitoringConfigurations` | `List<ContainerMonitoringConfiguration>` | yes |
| `Tags` | `List<Tag>` | no |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerAssociationName` | `string` | no |
| `ContainerAssociationArn` | `string` | no |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `ContainerMonitoringConfigurations` | `List<ContainerMonitoringConfiguration>` | no |
| `Status` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `UpdateToken` | `string` | no |

## UpdateFirewallAnalysisSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnabledAnalysisTypes` | `List<string>` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `UpdateToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnabledAnalysisTypes` | `List<string>` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `UpdateToken` | `string` | no |

## UpdateFirewallDeleteProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `DeleteProtection` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `DeleteProtection` | `boolean` | no |
| `UpdateToken` | `string` | no |

## UpdateFirewallDescription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `Description` | `string` | no |
| `UpdateToken` | `string` | no |

## UpdateFirewallEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `UpdateToken` | `string` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |

## UpdateFirewallPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `FirewallPolicyArn` | `string` | no |
| `FirewallPolicyName` | `string` | no |
| `FirewallPolicy` | `FirewallPolicy` | yes |
| `Description` | `string` | no |
| `DryRun` | `boolean` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `FirewallPolicyResponse` | `FirewallPolicyResponse` | yes |

## UpdateFirewallPolicyChangeProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `FirewallPolicyChangeProtection` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `FirewallPolicyChangeProtection` | `boolean` | no |

## UpdateLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `LoggingConfiguration` | `LoggingConfiguration` | no |
| `EnableMonitoringDashboard` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `LoggingConfiguration` | `LoggingConfiguration` | no |
| `EnableMonitoringDashboard` | `boolean` | no |

## UpdateProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | yes |
| `ProxyName` | `string` | no |
| `ProxyArn` | `string` | no |
| `ListenerPropertiesToAdd` | `List<ListenerPropertyRequest>` | no |
| `ListenerPropertiesToRemove` | `List<ListenerPropertyRequest>` | no |
| `TlsInterceptProperties` | `TlsInterceptPropertiesRequest` | no |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Proxy` | `Proxy` | no |
| `UpdateToken` | `string` | no |

## UpdateProxyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurationName` | `string` | no |
| `ProxyConfigurationArn` | `string` | no |
| `DefaultRulePhaseActions` | `ProxyConfigDefaultRulePhaseActionsRequest` | yes |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfiguration` | `ProxyConfiguration` | no |
| `UpdateToken` | `string` | no |

## UpdateProxyRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupName` | `string` | no |
| `ProxyRuleGroupArn` | `string` | no |
| `ProxyRuleName` | `string` | yes |
| `Description` | `string` | no |
| `Action` | `string` | no |
| `AddConditions` | `List<ProxyRuleCondition>` | no |
| `RemoveConditions` | `List<ProxyRuleCondition>` | no |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRule` | `ProxyRule` | no |
| `RemovedConditions` | `List<ProxyRuleCondition>` | no |
| `UpdateToken` | `string` | no |

## UpdateProxyRuleGroupPriorities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyConfigurationName` | `string` | no |
| `ProxyConfigurationArn` | `string` | no |
| `RuleGroups` | `List<ProxyRuleGroupPriority>` | yes |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroups` | `List<ProxyRuleGroupPriorityResult>` | no |
| `UpdateToken` | `string` | no |

## UpdateProxyRulePriorities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupName` | `string` | no |
| `ProxyRuleGroupArn` | `string` | no |
| `RuleGroupRequestPhase` | `string` | yes |
| `Rules` | `List<ProxyRulePriority>` | yes |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxyRuleGroupName` | `string` | no |
| `ProxyRuleGroupArn` | `string` | no |
| `RuleGroupRequestPhase` | `string` | no |
| `Rules` | `List<ProxyRulePriority>` | no |
| `UpdateToken` | `string` | no |

## UpdateProxySettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `UpdateToken` | `string` | no |
| `ProxySettings` | `ProxySettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `UpdateToken` | `string` | no |
| `ProxySettings` | `ProxySettings` | no |

## UpdateRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `RuleGroupArn` | `string` | no |
| `RuleGroupName` | `string` | no |
| `RuleGroup` | `RuleGroup` | no |
| `Rules` | `string` | no |
| `Type` | `string` | no |
| `Description` | `string` | no |
| `DryRun` | `boolean` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `SourceMetadata` | `SourceMetadata` | no |
| `AnalyzeRuleGroup` | `boolean` | no |
| `SummaryConfiguration` | `SummaryConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `RuleGroupResponse` | `RuleGroupResponse` | yes |

## UpdateSubnetChangeProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `SubnetChangeProtection` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | no |
| `FirewallArn` | `string` | no |
| `FirewallName` | `string` | no |
| `SubnetChangeProtection` | `boolean` | no |

## UpdateTLSInspectionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TLSInspectionConfigurationArn` | `string` | no |
| `TLSInspectionConfigurationName` | `string` | no |
| `TLSInspectionConfiguration` | `TLSInspectionConfiguration` | yes |
| `Description` | `string` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `UpdateToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateToken` | `string` | yes |
| `TLSInspectionConfigurationResponse` | `TLSInspectionConfigurationResponse` | yes |

