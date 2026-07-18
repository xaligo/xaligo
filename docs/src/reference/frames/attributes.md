# Frame and Border Attribute Reference

## Border Attributes

| Target | Attribute | Description |
|---|---|---|
| page `<frame>` | `width`, `height` | Canvas size in layout pixels |
| page `<frame>` | `margin`, `margin-*` | Inset diagram content without shrinking the paper frame |
| identified child `<frame>` | `title`, `version` | Enable the page metadata band and supply visible built-in values; child `version` is a content revision, not the DSL version |
| frame `<metadata>` | `position` | `top` (default) or `bottom` |
| frame `<metadata>` | `align` | `left` (default), `center`, or `right`; applied to each wrapped row independently |
| frame `<metadata>` | `font-family`, `font-size` | Tag typography; defaults are `virgil` and `12`, and height follows `ceil(font-size × 1.2) + 4` |
| frame `<metadata>` | `color`, `key-color`, `background-color`, `key-background-color`, `border-color` | `#RRGGBB` or `transparent`; defaults are `#64748b`, inherited `color`, `transparent`, `#f8fafc`, and `#cbd5e1` respectively |
| frame `<metadata>`, `<entry>` | `width`, `key-width` | Positive manual total/key-cell width; omission means auto, and entry values override metadata defaults |
| frame `<metadata>` | `gap`, `row-gap` | Non-negative horizontal and wrapped-row spacing; defaults are `8` and `4` |
| frame `<metadata>` | `<entry key="..." value="..." />` | Arbitrary non-empty key/value tag, retained in source order |
| frame `<metadata>` `<entry>` | `break-before` | `false` (default) or `true`; starts that entry on a new row when a preceding tag exists |
| generic leaf box | `border="none"` | Hide the visible border |
| generic leaf box | `visible="false"` | Hide component while preserving layout space |
| generic leaf box, rectangle, port | `font-size` | Label font size |
| rectangle, port | `id` | Required unique connection reference |
| port | `side` | `top`, `right`, `bottom`, or `left` |

Metadata tags pack greedily in input order against the usable width, producing
the minimum row count unless `break-before` introduces an earlier boundary.
The usable width is the complete outer frame width: the selected top/bottom
edge and left/right row alignment are not inset by padding, margins, or the
content box. A full-width reservation strip extends from that outer edge to the
final content-box boundary and is at least the band height plus the fixed
8-pixel gap. Normal items, text, connector paths and labels, and page links stay
outside it, even when the frame or a nested container uses
`overflow="visible"`. Metadata cell borders use a fixed `0.75`-pixel stroke.

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

See [Layout: Frame metadata](../../xal/layout.md#frame-metadata) for activation,
defaults, wrapping, output projection, and a complete example.
