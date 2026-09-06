# AWS App Mesh

API version: 2019-01-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/appmesh/2019-01-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `gatewayRouteName` | `string` | yes |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `GatewayRouteSpec` | yes |
| `tags` | `List<TagRef>` | no |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayRoute` | `GatewayRouteData` | yes |

## CreateMesh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `spec` | `MeshSpec` | no |
| `tags` | `List<TagRef>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mesh` | `MeshData` | yes |

## CreateRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `routeName` | `string` | yes |
| `spec` | `RouteSpec` | yes |
| `tags` | `List<TagRef>` | no |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `route` | `RouteData` | yes |

## CreateVirtualGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `VirtualGatewaySpec` | yes |
| `tags` | `List<TagRef>` | no |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualGateway` | `VirtualGatewayData` | yes |

## CreateVirtualNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `VirtualNodeSpec` | yes |
| `tags` | `List<TagRef>` | no |
| `virtualNodeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualNode` | `VirtualNodeData` | yes |

## CreateVirtualRouter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `VirtualRouterSpec` | yes |
| `tags` | `List<TagRef>` | no |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualRouter` | `VirtualRouterData` | yes |

## CreateVirtualService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `VirtualServiceSpec` | yes |
| `tags` | `List<TagRef>` | no |
| `virtualServiceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualService` | `VirtualServiceData` | yes |

## DeleteGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayRouteName` | `string` | yes |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayRoute` | `GatewayRouteData` | yes |

## DeleteMesh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mesh` | `MeshData` | yes |

## DeleteRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `routeName` | `string` | yes |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `route` | `RouteData` | yes |

## DeleteVirtualGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualGateway` | `VirtualGatewayData` | yes |

## DeleteVirtualNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualNodeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualNode` | `VirtualNodeData` | yes |

## DeleteVirtualRouter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualRouter` | `VirtualRouterData` | yes |

## DeleteVirtualService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualServiceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualService` | `VirtualServiceData` | yes |

## DescribeGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayRouteName` | `string` | yes |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayRoute` | `GatewayRouteData` | yes |

## DescribeMesh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mesh` | `MeshData` | yes |

## DescribeRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `routeName` | `string` | yes |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `route` | `RouteData` | yes |

## DescribeVirtualGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualGateway` | `VirtualGatewayData` | yes |

## DescribeVirtualNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualNodeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualNode` | `VirtualNodeData` | yes |

## DescribeVirtualRouter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualRouter` | `VirtualRouterData` | yes |

## DescribeVirtualService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `virtualServiceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualService` | `VirtualServiceData` | yes |

## ListGatewayRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `nextToken` | `string` | no |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayRoutes` | `List<GatewayRouteRef>` | yes |
| `nextToken` | `string` | no |

## ListMeshes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meshes` | `List<MeshRef>` | yes |
| `nextToken` | `string` | no |

## ListRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `nextToken` | `string` | no |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `routes` | `List<RouteRef>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `nextToken` | `string` | no |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `tags` | `List<TagRef>` | yes |

## ListVirtualGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `virtualGateways` | `List<VirtualGatewayRef>` | yes |

## ListVirtualNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `virtualNodes` | `List<VirtualNodeRef>` | yes |

## ListVirtualRouters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `virtualRouters` | `List<VirtualRouterRef>` | yes |

## ListVirtualServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `virtualServices` | `List<VirtualServiceRef>` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<TagRef>` | yes |

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


## UpdateGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `gatewayRouteName` | `string` | yes |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `GatewayRouteSpec` | yes |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayRoute` | `GatewayRouteData` | yes |

## UpdateMesh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `spec` | `MeshSpec` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mesh` | `MeshData` | yes |

## UpdateRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `routeName` | `string` | yes |
| `spec` | `RouteSpec` | yes |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `route` | `RouteData` | yes |

## UpdateVirtualGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `VirtualGatewaySpec` | yes |
| `virtualGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualGateway` | `VirtualGatewayData` | yes |

## UpdateVirtualNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `VirtualNodeSpec` | yes |
| `virtualNodeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualNode` | `VirtualNodeData` | yes |

## UpdateVirtualRouter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `VirtualRouterSpec` | yes |
| `virtualRouterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualRouter` | `VirtualRouterData` | yes |

## UpdateVirtualService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `meshName` | `string` | yes |
| `meshOwner` | `string` | no |
| `spec` | `VirtualServiceSpec` | yes |
| `virtualServiceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualService` | `VirtualServiceData` | yes |

