# Frame and Border Attribute Reference

## Border Attributes

| Target | Attribute | Description |
|---|---|---|
| root `<frame>` | `width`, `height` | Canvas size in layout pixels |
| root `<frame>` | `margin`, `margin-*` | Inset diagram content without shrinking the paper frame |
| generic leaf box | `border="none"` | Hide the visible border |
| generic leaf box | `visible="false"` | Hide component while preserving layout space |
| generic leaf box, rectangle, port | `font-size` | Label font size |
| rectangle, port | `id` | Required unique connection reference |
| port | `side` | `top`, `right`, `bottom`, or `left` |

## AWS Group Border Styles

| Tag | Border color | Stroke | Width | Icon |
|---|---|---|---|---|
| `<aws-cloud>` | `#000000` | solid | `2` | `AWS-Cloud-logo_32.svg` |
| `<aws-cloud-alt>` | `#000000` | solid | `2` | `AWS-Cloud_32.svg` |
| `<region>` | `#00A1C9` | dashed | `2` | `Region_32.svg` |
| `<availability-zone>` | `#00A1C9` | dashed | `2` | none |
| `<security-group>` | `#CC0000` | dashed | `2` | none |
| `<auto-scaling-group>` | `#E7601B` | solid | `2` | `Auto-Scaling-group_32.svg` |
| `<vpc>` | `#8C4FFF` | solid | `2` | `Virtual-private-cloud-VPC_32.svg` |
| `<private-subnet>` | `#00A1C9` | solid | `2` | `Private-subnet_32.svg` |
| `<public-subnet>` | `#3F8624` | solid | `2` | `Public-subnet_32.svg` |
| `<server-contents>` | `#7A7C7F` | solid | `2` | `Server-contents_32.svg` |
| `<corporate-data-center>` | `#7A7C7F` | solid | `2` | `Corporate-data-center_32.svg` |
| `<ec2-instance-contents>` | `#E7601B` | solid | `2` | `EC2-instance-contents_32.svg` |
| `<spot-fleet>` | `#E7601B` | solid | `2` | `Spot-Fleet_32.svg` |
| `<aws-account>` | `#E7008A` | solid | `2` | `AWS-Account_32.svg` |
| `<aws-iot-greengrass-deployment>` | `#3F8624` | solid | `2` | `AWS-IoT-Greengrass-Deployment_32.svg` |
| `<aws-iot-greengrass>` | `#3F8624` | solid | `2` | none |
| `<elastic-beanstalk-container>` | `#E7601B` | solid | `2` | none |
| `<aws-step-functions-workflow>` | `#E7008A` | solid | `2` | none |
| `<generic-group>` | `#AAB7B8` | dashed | `1` | configured by positive decimal int32 `icon-id` (`1..2147483647`) |
