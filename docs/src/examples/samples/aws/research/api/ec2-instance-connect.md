# AWS EC2 Instance Connect

API version: 2018-04-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ec2-instance-connect/2018-04-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## SendSSHPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `InstanceOSUser` | `string` | yes |
| `SSHPublicKey` | `string` | yes |
| `AvailabilityZone` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Success` | `boolean` | no |

## SendSerialConsoleSSHPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `SerialPort` | `integer` | no |
| `SSHPublicKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestId` | `string` | no |
| `Success` | `boolean` | no |

