# `aws-direct-connect-gateway` — AWS Direct Connect Gateway

[SVG preview](sample.svg) · [Editable XAL](sample.xal) · [Catalog](../README.md)

![AWS Direct Connect Gateway](sample.svg)

AWS resource icon. Use a label and explicit annotations to describe its role; scope is selected by the author.

- Kind: `resource`; category: Networking Content Delivery.
- Diagram scope: `logical` (recommendation, not AWS deployment validation).
- Default catalog ID: 1561. Covered catalog IDs: 1561.
- Implementation: V1 and V2; fixed AWS icon with a wrapped label and explicit functional annotations.

## Parameters

`id` is a required, unique connection ID, not a catalog number. `label`/`title`/`name` override the label; an empty label hides it. `size` > 0 defaults to 48 px. `label-width` > 0 defaults to 160 px (default box width, at least icon size + 12 px). Explicit `width`/`height` must contain the icon and label. `visible="false"` hides it. Children and icon overrides are not supported; use a group for containment.

`detail` adds a free-form diagram annotation. `show-details="false"` hides annotation text. Only supplied values are shown; none are sent to AWS. Service/resource annotations appear on separate wrapped lines.

| Parameter | Type | Meaning | Example |
|---|---|---|---|
| `traffic` | text | Traffic/protocol annotation | `HTTPS` |

## Review notes

The catalog provides a baseline for per-component development, not a simulation of the AWS control plane. This component's current functional parameters are the ones listed above. Additional service-specific visual behavior can be developed here without replacing catalog IDs in diagrams. Edit `sample.xal`, then run:

```sh
npm run generate:aws-samples -- --render --tag=aws-direct-connect-gateway
```

<!-- aws-functional-research:start -->
## 機能調査・構成デザイン（2026-09-06）

分類: `service-context`。サービス文脈: [`aws-direct-connect`](../aws-direct-connect/README.md)。

サンプルはアイコンと、設定・内包構造・関連リソース・操作を分離したレビューシートです。設定カードは編集可能な `rectangle`、グループは既存の専用タグで実装しています。カードのフィールド名を新しい XAL 属性として受理するわけではありません。専用タグが受理する属性は上の Parameters 表を参照してください。

実線の通信と、設定の参照・同じサービスに属する型一覧を区別します。スキーマの必須項目は AWS 側の仕様であり、図の必須入力ではありません。記載の構成モデル/API は取り込んだ公式資料の範囲であり、全リージョン・全機能の完全性や稼働可否を保証しません。

**重要:** このアイコンに対応する独立した構成リソースを断定せず、所属サービスの構成モデルを参考表示しています。アイコン名や絵柄から属性・親子関係・通信を推測しません。

### 構成モデル: `AWS::DirectConnect::Connection`

[公式リファレンス](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html)。全 7 プロパティを型付きで列挙します（表示カードには主要項目のみ）。

| Field | Type | Required in AWS schema |
|---|---|---|
| [Bandwidth](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-bandwidth) | `String` | yes |
| [ConnectionName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-connectionname) | `String` | yes |
| [LagId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-lagid) | `String` | no |
| [Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-location) | `String` | yes |
| [ProviderName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-providername) | `String` | no |
| [RequestMACSec](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-requestmacsec) | `Boolean` | no |
| [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-tags) | `List<Tag>` | no |

### 関連する構成リソース（7 型）

同じサービス文脈の型一覧です。すべてがこのアイコンの子リソースという意味ではありません。さらに深い入れ子は [公式スキーマのスナップショット](../research/cloudformation-models.json) に記録しています。

- [AWS::DirectConnect::Connection](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html)
- [AWS::DirectConnect::DirectConnectGateway](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgateway.html)
- [AWS::DirectConnect::DirectConnectGatewayAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgatewayassociation.html)
- [AWS::DirectConnect::Lag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html)
- [AWS::DirectConnect::PrivateVirtualInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-privatevirtualinterface.html)
- [AWS::DirectConnect::PublicVirtualInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html)
- [AWS::DirectConnect::TransitVirtualInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html)

### API の操作・パラメータ

- [AWS Direct Connect: 64 操作の入力・出力一覧](../research/api/directconnect.md)（API version 2012-10-25）

### 出典・調査範囲

- [公式資料 1](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html)
- [公式資料 2](https://github.com/boto/botocore/blob/develop/botocore/data/directconnect/2012-10-25/service-2.json)

CloudFormation 仕様 263.0.0、AWS SDK 431 サービスモデルをオフラインで参照。取得日・元データの SHA-256 は [調査マニフェスト](../research/README.md) を参照。API モデル名・フィールド名は仕様から抽出し、説明本文を転載していません。利用可能性は全サービス一律には確認できないため、提供終了の確認がないものも「現在利用可能」と断定していません。

### 次の部品レビュー

- 本アイコンが独立したリソースか、機能・状態・デバイスの記号かを確認する。
- 詳細カードのうち専用の子タグ・参照属性として実装する範囲を選ぶ。
- 通信、制御、認証、監視の関係を分け、必要な接続点・配置制約を確認する。
- 編集後は `npm run generate:aws-samples -- --render --tag=aws-direct-connect-gateway`。通常の再描画は XAL/README を上書きしない。
<!-- aws-functional-research:end -->
