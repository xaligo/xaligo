# `aws-application-migration-service` — AWS Application Migration Service

[SVG preview](sample.svg) · [Editable XAL](sample.xal) · [Catalog](../README.md)

![AWS Application Migration Service](sample.svg)

AWS service icon. Use a label and explicit annotations to describe its role; scope is selected by the author.

- Kind: `service`; category: Migration Modernization.
- Diagram scope: `service` (recommendation, not AWS deployment validation).
- Default catalog ID: 525. Covered catalog IDs: 498, 507, 516, 525.
- Implementation: V1 and V2; fixed AWS icon with a wrapped label and explicit functional annotations.

## Parameters

`id` is a required, unique connection ID, not a catalog number. `label`/`title`/`name` override the label; an empty label hides it. `size` > 0 defaults to 48 px. `label-width` > 0 defaults to 160 px (default box width, at least icon size + 12 px). Explicit `width`/`height` must contain the icon and label. `visible="false"` hides it. Children and icon overrides are not supported; use a group for containment.

`detail` adds a free-form diagram annotation. `show-details="false"` hides annotation text. Only supplied values are shown; none are sent to AWS. Service/resource annotations appear on separate wrapped lines.

| Parameter | Type | Meaning | Example |
|---|---|---|---|
| `role` | text | Architectural role annotation | `Application component` |

## Review notes

The catalog provides a baseline for per-component development, not a simulation of the AWS control plane. This component's current functional parameters are the ones listed above. Additional service-specific visual behavior can be developed here without replacing catalog IDs in diagrams. Edit `sample.xal`, then run:

```sh
npm run generate:aws-samples -- --render --tag=aws-application-migration-service
```

<!-- aws-functional-research:start -->
## 機能調査・構成デザイン（2026-09-06）

分類: `service-context`。サービス文脈: [`aws-application-migration-service`](../aws-application-migration-service/README.md)。

サンプルはアイコンと、設定・内包構造・関連リソース・操作を分離したレビューシートです。設定カードは編集可能な `rectangle`、グループは既存の専用タグで実装しています。カードのフィールド名を新しい XAL 属性として受理するわけではありません。専用タグが受理する属性は上の Parameters 表を参照してください。

実線の通信と、設定の参照・同じサービスに属する型一覧を区別します。スキーマの必須項目は AWS 側の仕様であり、図の必須入力ではありません。記載の構成モデル/API は取り込んだ公式資料の範囲であり、全リージョン・全機能の完全性や稼働可否を保証しません。

**重要:** このアイコンに対応する独立した構成リソースを断定せず、所属サービスの構成モデルを参考表示しています。アイコン名や絵柄から属性・親子関係・通信を推測しません。

### 構成モデル: `AWS::MGN::NetworkMigrationDefinition`

[公式リファレンス](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html)。全 8 プロパティを型付きで列挙します（表示カードには主要項目のみ）。

| Field | Type | Required in AWS schema |
|---|---|---|
| [Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-description) | `String` | no |
| [Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-name) | `String` | yes |
| [ScopeTags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-scopetags) | `Map<String>` | no |
| [SourceConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-sourceconfigurations) | `List<SourceConfiguration>` | yes |
| [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-tags) | `List<Tag>` | no |
| [TargetDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-targetdeployment) | `String` | no |
| [TargetNetwork](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-targetnetwork) | `TargetNetwork` | yes |
| [TargetS3Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html#cfn-mgn-networkmigrationdefinition-targets3configuration) | `TargetS3Configuration` | yes |

#### SourceConfigurations → `AWS::MGN::NetworkMigrationDefinition.SourceConfiguration`

| Field | Type | Required in AWS schema |
|---|---|---|
| [SourceEnvironment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sourceconfiguration.html#cfn-mgn-networkmigrationdefinition-sourceconfiguration-sourceenvironment) | `String` | yes |
| [SourceS3Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sourceconfiguration.html#cfn-mgn-networkmigrationdefinition-sourceconfiguration-sources3configuration) | `SourceS3Configuration` | yes |

#### TargetNetwork → `AWS::MGN::NetworkMigrationDefinition.TargetNetwork`

| Field | Type | Required in AWS schema |
|---|---|---|
| [InboundCidr](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html#cfn-mgn-networkmigrationdefinition-targetnetwork-inboundcidr) | `String` | no |
| [InspectionCidr](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html#cfn-mgn-networkmigrationdefinition-targetnetwork-inspectioncidr) | `String` | no |
| [OutboundCidr](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html#cfn-mgn-networkmigrationdefinition-targetnetwork-outboundcidr) | `String` | no |
| [Topology](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targetnetwork.html#cfn-mgn-networkmigrationdefinition-targetnetwork-topology) | `String` | yes |

#### TargetS3Configuration → `AWS::MGN::NetworkMigrationDefinition.TargetS3Configuration`

| Field | Type | Required in AWS schema |
|---|---|---|
| [S3Bucket](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targets3configuration.html#cfn-mgn-networkmigrationdefinition-targets3configuration-s3bucket) | `String` | yes |
| [S3BucketOwner](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-targets3configuration.html#cfn-mgn-networkmigrationdefinition-targets3configuration-s3bucketowner) | `String` | yes |

### 関連する構成リソース（1 型）

同じサービス文脈の型一覧です。すべてがこのアイコンの子リソースという意味ではありません。さらに深い入れ子は [公式スキーマのスナップショット](../research/cloudformation-models.json) に記録しています。

- [AWS::MGN::NetworkMigrationDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html)

### API の操作・パラメータ

- [Application Migration Service: 95 操作の入力・出力一覧](../research/api/mgn.md)（API version 2020-02-26）

### 出典・調査範囲

- [公式資料 1](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-networkmigrationdefinition.html)
- [公式資料 2](https://github.com/boto/botocore/blob/develop/botocore/data/mgn/2020-02-26/service-2.json)

CloudFormation 仕様 263.0.0、AWS SDK 431 サービスモデルをオフラインで参照。取得日・元データの SHA-256 は [調査マニフェスト](../research/README.md) を参照。API モデル名・フィールド名は仕様から抽出し、説明本文を転載していません。利用可能性は全サービス一律には確認できないため、提供終了の確認がないものも「現在利用可能」と断定していません。

### 次の部品レビュー

- 本アイコンが独立したリソースか、機能・状態・デバイスの記号かを確認する。
- 詳細カードのうち専用の子タグ・参照属性として実装する範囲を選ぶ。
- 通信、制御、認証、監視の関係を分け、必要な接続点・配置制約を確認する。
- 編集後は `npm run generate:aws-samples -- --render --tag=aws-application-migration-service`。通常の再描画は XAL/README を上書きしない。
<!-- aws-functional-research:end -->
