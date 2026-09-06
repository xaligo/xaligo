# `aws-marketplace-light` — AWS Marketplace Light

[SVG preview](sample.svg) · [Editable XAL](sample.xal) · [Catalog](../README.md)

![AWS Marketplace Light](sample.svg)

AWS service icon. Use a label and explicit annotations to describe its role; scope is selected by the author.

- Kind: `service`; category: General Icons.
- Diagram scope: `service` (recommendation, not AWS deployment validation).
- Default catalog ID: 540. Covered catalog IDs: 534, 536, 538, 540.
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
npm run generate:aws-samples -- --render --tag=aws-marketplace-light
```

<!-- aws-functional-research:start -->
## 機能調査・構成デザイン（2026-09-06）

分類: `api-context`。サービス文脈: [`aws-marketplace-light`](../aws-marketplace-light/README.md)。

サンプルはアイコンと、設定・内包構造・関連リソース・操作を分離したレビューシートです。設定カードは編集可能な `rectangle`、グループは既存の専用タグで実装しています。カードのフィールド名を新しい XAL 属性として受理するわけではありません。専用タグが受理する属性は上の Parameters 表を参照してください。

実線の通信と、設定の参照・同じサービスに属する型一覧を区別します。スキーマの必須項目は AWS 側の仕様であり、図の必須入力ではありません。記載の構成モデル/API は取り込んだ公式資料の範囲であり、全リージョン・全機能の完全性や稼働可否を保証しません。

### API の操作・パラメータ

- [AWS Marketplace Catalog Service: 15 操作の入力・出力一覧](../research/api/marketplace-catalog.md)（API version 2018-09-17）
- [AWS Marketplace Agreement Service: 25 操作の入力・出力一覧](../research/api/marketplace-agreement.md)（API version 2020-03-01）
- [AWSMarketplace Metering: 4 操作の入力・出力一覧](../research/api/meteringmarketplace.md)（API version 2016-01-14）

### 出典・調査範囲

- [公式資料 1](https://github.com/boto/botocore/blob/develop/botocore/data/marketplace-catalog/2018-09-17/service-2.json)
- [公式資料 2](https://github.com/boto/botocore/blob/develop/botocore/data/marketplace-agreement/2020-03-01/service-2.json)
- [公式資料 3](https://github.com/boto/botocore/blob/develop/botocore/data/meteringmarketplace/2016-01-14/service-2.json)

CloudFormation 仕様 263.0.0、AWS SDK 431 サービスモデルをオフラインで参照。取得日・元データの SHA-256 は [調査マニフェスト](../research/README.md) を参照。API モデル名・フィールド名は仕様から抽出し、説明本文を転載していません。利用可能性は全サービス一律には確認できないため、提供終了の確認がないものも「現在利用可能」と断定していません。

### 次の部品レビュー

- 本アイコンが独立したリソースか、機能・状態・デバイスの記号かを確認する。
- 詳細カードのうち専用の子タグ・参照属性として実装する範囲を選ぶ。
- 通信、制御、認証、監視の関係を分け、必要な接続点・配置制約を確認する。
- 編集後は `npm run generate:aws-samples -- --render --tag=aws-marketplace-light`。通常の再描画は XAL/README を上書きしない。
<!-- aws-functional-research:end -->
