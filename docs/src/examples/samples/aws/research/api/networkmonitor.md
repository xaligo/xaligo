# Amazon CloudWatch Network Monitor

API version: 2023-08-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/networkmonitor/2023-08-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `probes` | `List<CreateMonitorProbeInput>` | no |
| `aggregationPeriod` | `long` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorArn` | `string` | yes |
| `monitorName` | `string` | yes |
| `state` | `string` | yes |
| `aggregationPeriod` | `long` | no |
| `tags` | `Map<string>` | no |

## CreateProbe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `probe` | `ProbeInput` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `probeId` | `string` | no |
| `probeArn` | `string` | no |
| `sourceArn` | `string` | yes |
| `destination` | `string` | yes |
| `destinationPort` | `integer` | no |
| `protocol` | `string` | yes |
| `packetSize` | `integer` | no |
| `addressFamily` | `string` | no |
| `vpcId` | `string` | no |
| `state` | `string` | no |
| `createdAt` | `timestamp` | no |
| `modifiedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## DeleteMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProbe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `probeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorArn` | `string` | yes |
| `monitorName` | `string` | yes |
| `state` | `string` | yes |
| `aggregationPeriod` | `long` | yes |
| `tags` | `Map<string>` | no |
| `probes` | `List<Probe>` | no |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | yes |

## GetProbe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `probeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `probeId` | `string` | no |
| `probeArn` | `string` | no |
| `sourceArn` | `string` | yes |
| `destination` | `string` | yes |
| `destinationPort` | `integer` | no |
| `protocol` | `string` | yes |
| `packetSize` | `integer` | no |
| `addressFamily` | `string` | no |
| `vpcId` | `string` | no |
| `state` | `string` | no |
| `createdAt` | `timestamp` | no |
| `modifiedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## ListMonitors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `state` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitors` | `List<MonitorSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

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


## UpdateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `aggregationPeriod` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorArn` | `string` | yes |
| `monitorName` | `string` | yes |
| `state` | `string` | yes |
| `aggregationPeriod` | `long` | no |
| `tags` | `Map<string>` | no |

## UpdateProbe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `probeId` | `string` | yes |
| `state` | `string` | no |
| `destination` | `string` | no |
| `destinationPort` | `integer` | no |
| `protocol` | `string` | no |
| `packetSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `probeId` | `string` | no |
| `probeArn` | `string` | no |
| `sourceArn` | `string` | yes |
| `destination` | `string` | yes |
| `destinationPort` | `integer` | no |
| `protocol` | `string` | yes |
| `packetSize` | `integer` | no |
| `addressFamily` | `string` | no |
| `vpcId` | `string` | no |
| `state` | `string` | no |
| `createdAt` | `timestamp` | no |
| `modifiedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |

