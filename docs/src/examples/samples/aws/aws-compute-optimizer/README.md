# `aws-compute-optimizer` — AWS Compute Optimizer

[SVG preview](sample.svg) · [Editable XAL](sample.xal) · [Catalog](../README.md)

![AWS Compute Optimizer](sample.svg)

AWS service icon. Use a label and explicit annotations to describe its role; scope is selected by the author.

- Kind: `service`; category: Compute.
- Diagram scope: `service` (recommendation, not AWS deployment validation).
- Default catalog ID: 86. Covered catalog IDs: 11, 36, 61, 86, 584, 614, 644, 674.
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
npm run generate:aws-samples -- --render --tag=aws-compute-optimizer
```

<!-- aws-functional-research:start -->
## 機能調査・構成デザイン（2026-09-06）

分類: `service-context`。サービス文脈: [`aws-compute-optimizer`](../aws-compute-optimizer/README.md)。

サンプルはアイコンと、設定・内包構造・関連リソース・操作を分離したレビューシートです。設定カードは編集可能な `rectangle`、グループは既存の専用タグで実装しています。カードのフィールド名を新しい XAL 属性として受理するわけではありません。専用タグが受理する属性は上の Parameters 表を参照してください。

実線の通信と、設定の参照・同じサービスに属する型一覧を区別します。スキーマの必須項目は AWS 側の仕様であり、図の必須入力ではありません。記載の構成モデル/API は取り込んだ公式資料の範囲であり、全リージョン・全機能の完全性や稼働可否を保証しません。

**重要:** このアイコンに対応する独立した構成リソースを断定せず、所属サービスの構成モデルを参考表示しています。アイコン名や絵柄から属性・親子関係・通信を推測しません。

### 構成モデル: `AWS::ComputeOptimizer::AutomationRule`

[公式リファレンス](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html)。全 10 プロパティを型付きで列挙します（表示カードには主要項目のみ）。

| Field | Type | Required in AWS schema |
|---|---|---|
| [Criteria](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-criteria) | `Criteria` | no |
| [Description](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-description) | `String` | no |
| [Name](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-name) | `String` | yes |
| [OrganizationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-organizationconfiguration) | `OrganizationConfiguration` | no |
| [Priority](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-priority) | `String` | no |
| [RecommendedActionTypes](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-recommendedactiontypes) | `List<String>` | yes |
| [RuleType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-ruletype) | `String` | yes |
| [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-schedule) | `Schedule` | yes |
| [Status](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-status) | `String` | yes |
| [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html#cfn-computeoptimizer-automationrule-tags) | `List<Tag>` | no |

#### Schedule → `AWS::ComputeOptimizer::AutomationRule.Schedule`

| Field | Type | Required in AWS schema |
|---|---|---|
| [ExecutionWindowInMinutes](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-schedule.html#cfn-computeoptimizer-automationrule-schedule-executionwindowinminutes) | `Integer` | no |
| [ScheduleExpression](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-schedule.html#cfn-computeoptimizer-automationrule-schedule-scheduleexpression) | `String` | no |
| [ScheduleExpressionTimezone](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-schedule.html#cfn-computeoptimizer-automationrule-schedule-scheduleexpressiontimezone) | `String` | no |

#### Criteria → `AWS::ComputeOptimizer::AutomationRule.Criteria`

| Field | Type | Required in AWS schema |
|---|---|---|
| [EbsVolumeSizeInGib](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-ebsvolumesizeingib) | `List<IntegerCriteriaCondition>` | no |
| [EbsVolumeType](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-ebsvolumetype) | `List<StringCriteriaCondition>` | no |
| [EstimatedMonthlySavings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-estimatedmonthlysavings) | `List<DoubleCriteriaCondition>` | no |
| [LookBackPeriodInDays](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-lookbackperiodindays) | `List<IntegerCriteriaCondition>` | no |
| [Region](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-region) | `List<StringCriteriaCondition>` | no |
| [ResourceArn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-resourcearn) | `List<StringCriteriaCondition>` | no |
| [ResourceTag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-resourcetag) | `List<ResourceTagsCriteriaCondition>` | no |
| [RestartNeeded](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-restartneeded) | `List<StringCriteriaCondition>` | no |

#### OrganizationConfiguration → `AWS::ComputeOptimizer::AutomationRule.OrganizationConfiguration`

| Field | Type | Required in AWS schema |
|---|---|---|
| [AccountIds](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-organizationconfiguration.html#cfn-computeoptimizer-automationrule-organizationconfiguration-accountids) | `List<String>` | no |
| [RuleApplyOrder](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-organizationconfiguration.html#cfn-computeoptimizer-automationrule-organizationconfiguration-ruleapplyorder) | `String` | no |

### 関連する構成リソース（1 型）

同じサービス文脈の型一覧です。すべてがこのアイコンの子リソースという意味ではありません。さらに深い入れ子は [公式スキーマのスナップショット](../research/cloudformation-models.json) に記録しています。

- [AWS::ComputeOptimizer::AutomationRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html)

### API の操作・パラメータ

- [AWS Compute Optimizer: 28 操作の入力・出力一覧](../research/api/compute-optimizer.md)（API version 2019-11-01）

### 出典・調査範囲

- [公式資料 1](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-computeoptimizer-automationrule.html)
- [公式資料 2](https://github.com/boto/botocore/blob/develop/botocore/data/compute-optimizer/2019-11-01/service-2.json)

CloudFormation 仕様 263.0.0、AWS SDK 431 サービスモデルをオフラインで参照。取得日・元データの SHA-256 は [調査マニフェスト](../research/README.md) を参照。API モデル名・フィールド名は仕様から抽出し、説明本文を転載していません。利用可能性は全サービス一律には確認できないため、提供終了の確認がないものも「現在利用可能」と断定していません。

### 次の部品レビュー

- 本アイコンが独立したリソースか、機能・状態・デバイスの記号かを確認する。
- 詳細カードのうち専用の子タグ・参照属性として実装する範囲を選ぶ。
- 通信、制御、認証、監視の関係を分け、必要な接続点・配置制約を確認する。
- 編集後は `npm run generate:aws-samples -- --render --tag=aws-compute-optimizer`。通常の再描画は XAL/README を上書きしない。
<!-- aws-functional-research:end -->
