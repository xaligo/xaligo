# `aws-outposts-servers` — AWS Outposts servers

[SVG preview](sample.svg) · [Editable XAL](sample.xal) · [Catalog](../README.md)

![AWS Outposts servers](sample.svg)

AWS service icon. Use a label and explicit annotations to describe its role; scope is selected by the author.

- Kind: `service`; category: Compute.
- Diagram scope: `service` (recommendation, not AWS deployment validation).
- Default catalog ID: 93. Covered catalog IDs: 18, 43, 68, 93.
- Implementation: V1 and V2; fixed AWS icon with a wrapped label and explicit functional annotations.

## Parameters

`id` is a required, unique connection ID, not a catalog number. `label`/`title`/`name` override the label; an empty label hides it. `size` > 0 defaults to 48 px. `label-width` > 0 defaults to 160 px (default box width, at least icon size + 12 px). Explicit `width`/`height` must contain the icon and label. `visible="false"` hides it. Children and icon overrides are not supported; use a group for containment.

`detail` adds a free-form diagram annotation. `show-details="false"` hides annotation text. Only supplied values are shown; none are sent to AWS. Service/resource annotations appear on separate wrapped lines.

| Parameter | Type | Meaning | Example |
|---|---|---|---|
| `workload` | text | Workload or process annotation | `Web application` |

## Review notes

The catalog provides a baseline for per-component development, not a simulation of the AWS control plane. This component's current functional parameters are the ones listed above. Additional service-specific visual behavior can be developed here without replacing catalog IDs in diagrams. Edit `sample.xal`, then run:

```sh
npm run generate:aws-samples -- --render --tag=aws-outposts-servers
```

<!-- aws-functional-research:start -->
## 機能調査・構成デザイン（2026-09-06）

分類: `service-context`。サービス文脈: [`aws-outposts-servers`](../aws-outposts-servers/README.md)。

サンプルはアイコンと、設定・内包構造・関連リソース・操作を分離したレビューシートです。設定カードは編集可能な `rectangle`、グループは既存の専用タグで実装しています。カードのフィールド名を新しい XAL 属性として受理するわけではありません。専用タグが受理する属性は上の Parameters 表を参照してください。

実線の通信と、設定の参照・同じサービスに属する型一覧を区別します。スキーマの必須項目は AWS 側の仕様であり、図の必須入力ではありません。記載の構成モデル/API は取り込んだ公式資料の範囲であり、全リージョン・全機能の完全性や稼働可否を保証しません。

**重要:** このアイコンに対応する独立した構成リソースを断定せず、所属サービスの構成モデルを参考表示しています。アイコン名や絵柄から属性・親子関係・通信を推測しません。

### 構成モデル: `AWS::Outposts::Site`

[公式リファレンス](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html)。全 7 プロパティを型付きで列挙します（表示カードには主要項目のみ）。

| Field | Type | Required in AWS schema |
|---|---|---|
| [Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-description) | `String` | no |
| [Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-name) | `String` | yes |
| [Notes](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-notes) | `String` | no |
| [OperatingAddress](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-operatingaddress) | `Address` | no |
| [RackPhysicalProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-rackphysicalproperties) | `RackPhysicalProperties` | no |
| [ShippingAddress](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-shippingaddress) | `Address` | no |
| [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html#cfn-outposts-site-tags) | `List<Tag>` | no |

#### OperatingAddress → `AWS::Outposts::Site.Address`

| Field | Type | Required in AWS schema |
|---|---|---|
| [AddressLine1](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline1) | `String` | yes |
| [AddressLine2](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline2) | `String` | no |
| [AddressLine3](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline3) | `String` | no |
| [City](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-city) | `String` | yes |
| [ContactName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-contactname) | `String` | yes |
| [ContactPhoneNumber](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-contactphonenumber) | `String` | yes |
| [CountryCode](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-countrycode) | `String` | yes |
| [DistrictOrCounty](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-districtorcounty) | `String` | no |
| [Municipality](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-municipality) | `String` | no |
| [PostalCode](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-postalcode) | `String` | yes |
| [StateOrRegion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-stateorregion) | `String` | yes |

#### RackPhysicalProperties → `AWS::Outposts::Site.RackPhysicalProperties`

| Field | Type | Required in AWS schema |
|---|---|---|
| [FiberOpticCableType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-fiberopticcabletype) | `String` | no |
| [MaximumSupportedWeightLbs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-maximumsupportedweightlbs) | `String` | no |
| [OpticalStandard](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-opticalstandard) | `String` | no |
| [PowerConnector](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-powerconnector) | `String` | no |
| [PowerDrawKva](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-powerdrawkva) | `String` | no |
| [PowerFeedDrop](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-powerfeeddrop) | `String` | no |
| [PowerPhase](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-powerphase) | `String` | no |
| [UplinkCount](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-uplinkcount) | `String` | no |
| [UplinkGbps](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-rackphysicalproperties.html#cfn-outposts-site-rackphysicalproperties-uplinkgbps) | `String` | no |

#### ShippingAddress → `AWS::Outposts::Site.Address`

| Field | Type | Required in AWS schema |
|---|---|---|
| [AddressLine1](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline1) | `String` | yes |
| [AddressLine2](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline2) | `String` | no |
| [AddressLine3](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline3) | `String` | no |
| [City](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-city) | `String` | yes |
| [ContactName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-contactname) | `String` | yes |
| [ContactPhoneNumber](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-contactphonenumber) | `String` | yes |
| [CountryCode](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-countrycode) | `String` | yes |
| [DistrictOrCounty](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-districtorcounty) | `String` | no |
| [Municipality](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-municipality) | `String` | no |
| [PostalCode](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-postalcode) | `String` | yes |
| [StateOrRegion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-stateorregion) | `String` | yes |

### 関連する構成リソース（1 型）

同じサービス文脈の型一覧です。すべてがこのアイコンの子リソースという意味ではありません。さらに深い入れ子は [公式スキーマのスナップショット](../research/cloudformation-models.json) に記録しています。

- [AWS::Outposts::Site](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html)

### API の操作・パラメータ

- [AWS Outposts: 45 操作の入力・出力一覧](../research/api/outposts.md)（API version 2019-12-03）

### 出典・調査範囲

- [公式資料 1](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html)
- [公式資料 2](https://github.com/boto/botocore/blob/develop/botocore/data/outposts/2019-12-03/service-2.json)

CloudFormation 仕様 263.0.0、AWS SDK 431 サービスモデルをオフラインで参照。取得日・元データの SHA-256 は [調査マニフェスト](../research/README.md) を参照。API モデル名・フィールド名は仕様から抽出し、説明本文を転載していません。利用可能性は全サービス一律には確認できないため、提供終了の確認がないものも「現在利用可能」と断定していません。

### 次の部品レビュー

- 本アイコンが独立したリソースか、機能・状態・デバイスの記号かを確認する。
- 詳細カードのうち専用の子タグ・参照属性として実装する範囲を選ぶ。
- 通信、制御、認証、監視の関係を分け、必要な接続点・配置制約を確認する。
- 編集後は `npm run generate:aws-samples -- --render --tag=aws-outposts-servers`。通常の再描画は XAL/README を上書きしない。
<!-- aws-functional-research:end -->
