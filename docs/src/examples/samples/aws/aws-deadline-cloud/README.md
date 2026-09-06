# `aws-deadline-cloud` — AWS Deadline Cloud

[SVG preview](sample.svg) · [Editable XAL](sample.xal) · [Catalog](../README.md)

![AWS Deadline Cloud](sample.svg)

AWS service icon. Use a label and explicit annotations to describe its role; scope is selected by the author.

- Kind: `service`; category: Media Services.
- Diagram scope: `service` (recommendation, not AWS deployment validation).
- Default catalog ID: 815. Covered catalog IDs: 749, 771, 793, 815.
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
npm run generate:aws-samples -- --render --tag=aws-deadline-cloud
```

<!-- aws-functional-research:start -->
## 機能調査・構成デザイン（2026-09-06）

分類: `service-context`。サービス文脈: [`aws-deadline-cloud`](../aws-deadline-cloud/README.md)。

サンプルはアイコンと、設定・内包構造・関連リソース・操作を分離したレビューシートです。設定カードは編集可能な `rectangle`、グループは既存の専用タグで実装しています。カードのフィールド名を新しい XAL 属性として受理するわけではありません。専用タグが受理する属性は上の Parameters 表を参照してください。

実線の通信と、設定の参照・同じサービスに属する型一覧を区別します。スキーマの必須項目は AWS 側の仕様であり、図の必須入力ではありません。記載の構成モデル/API は取り込んだ公式資料の範囲であり、全リージョン・全機能の完全性や稼働可否を保証しません。

**重要:** このアイコンに対応する独立した構成リソースを断定せず、所属サービスの構成モデルを参考表示しています。アイコン名や絵柄から属性・親子関係・通信を推測しません。

### 構成モデル: `AWS::Deadline::Budget`

[公式リファレンス](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html)。全 8 プロパティを型付きで列挙します（表示カードには主要項目のみ）。

| Field | Type | Required in AWS schema |
|---|---|---|
| [Actions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-actions) | `List<BudgetActionToAdd>` | yes |
| [ApproximateDollarLimit](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-approximatedollarlimit) | `Double` | yes |
| [Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-description) | `String` | no |
| [DisplayName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-displayname) | `String` | yes |
| [FarmId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-farmid) | `String` | yes |
| [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-schedule) | `BudgetSchedule` | yes |
| [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-tags) | `List<Tag>` | no |
| [UsageTrackingResource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html#cfn-deadline-budget-usagetrackingresource) | `UsageTrackingResource` | yes |

#### Actions → `AWS::Deadline::Budget.BudgetActionToAdd`

| Field | Type | Required in AWS schema |
|---|---|---|
| [Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-description) | `String` | no |
| [ThresholdPercentage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-thresholdpercentage) | `Double` | yes |
| [Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-type) | `String` | yes |

#### Schedule → `AWS::Deadline::Budget.BudgetSchedule`

| Field | Type | Required in AWS schema |
|---|---|---|
| [Fixed](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetschedule.html#cfn-deadline-budget-budgetschedule-fixed) | `FixedBudgetSchedule` | yes |

#### UsageTrackingResource → `AWS::Deadline::Budget.UsageTrackingResource`

| Field | Type | Required in AWS schema |
|---|---|---|
| [QueueId](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-usagetrackingresource.html#cfn-deadline-budget-usagetrackingresource-queueid) | `String` | yes |

### 関連する構成リソース（14 型）

同じサービス文脈の型一覧です。すべてがこのアイコンの子リソースという意味ではありません。さらに深い入れ子は [公式スキーマのスナップショット](../research/cloudformation-models.json) に記録しています。

- [AWS::Deadline::Budget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html)
- [AWS::Deadline::Farm](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-farm.html)
- [AWS::Deadline::Fleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html)
- [AWS::Deadline::Job](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-job.html)
- [AWS::Deadline::LicenseEndpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-licenseendpoint.html)
- [AWS::Deadline::Limit](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-limit.html)
- [AWS::Deadline::MeteredProduct](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html)
- [AWS::Deadline::Monitor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-monitor.html)
- [AWS::Deadline::Queue](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html)
- [AWS::Deadline::QueueEnvironment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queueenvironment.html)
- [AWS::Deadline::QueueFleetAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html)
- [AWS::Deadline::QueueLimitAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html)
- [AWS::Deadline::StorageProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-storageprofile.html)
- [AWS::Deadline::Worker](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-worker.html)

### API の操作・パラメータ

- [AWSDeadlineCloud: 126 操作の入力・出力一覧](../research/api/deadline.md)（API version 2023-10-12）

### 出典・調査範囲

- [公式資料 1](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-budget.html)
- [公式資料 2](https://github.com/boto/botocore/blob/develop/botocore/data/deadline/2023-10-12/service-2.json)

CloudFormation 仕様 263.0.0、AWS SDK 431 サービスモデルをオフラインで参照。取得日・元データの SHA-256 は [調査マニフェスト](../research/README.md) を参照。API モデル名・フィールド名は仕様から抽出し、説明本文を転載していません。利用可能性は全サービス一律には確認できないため、提供終了の確認がないものも「現在利用可能」と断定していません。

### 次の部品レビュー

- 本アイコンが独立したリソースか、機能・状態・デバイスの記号かを確認する。
- 詳細カードのうち専用の子タグ・参照属性として実装する範囲を選ぶ。
- 通信、制御、認証、監視の関係を分け、必要な接続点・配置制約を確認する。
- 編集後は `npm run generate:aws-samples -- --render --tag=aws-deadline-cloud`。通常の再描画は XAL/README を上書きしない。
<!-- aws-functional-research:end -->
