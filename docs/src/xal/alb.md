# ALB — ルールと詳細度

V2 の ALB 専用部品は、リスナー、条件、書き換え、認証、転送先、運用設定を
小さなカードとして組み立てます。同じ設定から詳細図と全体構成図を作れます。
アイコンとドメインは左上、TLS / mTLS は縦並びで ON が緑、OFF が灰色です。
透かしはありません。表示中の内容だけでサイズを計算します。

## 最小例

既存の V2 `<frame>` 内に配置します。子タグの `id` はすべて必須・一意です。

```xml
<aws-elastic-load-balancing-application-load-balancer
    id="alb" domain="api.example.test" detail-level="standard">
  <aws-listener id="https" protocol="HTTPS" port="443" certificate="server-cert">
    <aws-listener-rule id="api-route" priority="10">
      <aws-rule-condition id="host" field="host-header">
        <aws-rule-match id="api-host" value="api.example.test" />
      </aws-rule-condition>
      <aws-rule-condition id="path" field="path-pattern">
        <aws-rule-match id="api-path" value="/v1/*" />
      </aws-rule-condition>
      <aws-rule-action id="to-api" type="forward" target-group="api" />
    </aws-listener-rule>
    <aws-listener-rule id="fallback" priority="default">
      <aws-rule-action id="to-web" type="forward" target-group="web" />
    </aws-listener-rule>
  </aws-listener>
</aws-elastic-load-balancing-application-load-balancer>
```

## 詳細度と非表示

| 設定 | 表示 |
|---|---|
| `detail-level="summary"` | アイコン、ドメイン、リスナーポート、TLS / mTLS。ルール・参照・設定・ターゲットは折りたたむ |
| `detail-level="standard"` | ルール、条件、アクション、書き換え、ターゲットも表示。`aws-option` は非表示 |
| `detail-level="detailed"` | 指定された全項目を表示 |
| `hide="conditions,transforms"` | 指定した種類をまとめて非表示 |
| `show="options"` | プリセットで非表示の種類を表示 |
| `visible="false"` | その子カードと子孫を非表示にして、配置領域も除く |
| `show-title="false"` | リスナーの既定見出し `Listener` だけを非表示 |

`detail-level` / `show` / `hide` は ALB と全専用子タグに指定できます。
未指定は親を継承し、子の明示プリセット → `show` → `hide` の順で適用します。
ただし、親が子の種類自体を隠した場合、子の `show` だけでは再表示しません。
例えば summary のリスナーに `show="rules"`、そのルールに
`detail-level="standard"` を指定すると、そのルールだけを詳しく表示できます。
同じ種類を `show` と `hide` の両方に指定するのはエラーです。
詳細指定も新しい子タグもない従来の ALB/NLB は既存のコンパクト描画を保ちます。

指定できる種類:
`title,protocol,tls,mtls,certificate,trust-store,target-group,icon,domain,rules,conditions,actions,transforms,options,targets,priority,values,weights,name,type`。
`options` は設定表、`values` は値、`targets` はターゲット群・登録ターゲット、
`target-group` は転送先の参照文字列です。個別オプションはその `aws-option` の
`visible="false"` で隠します。描画を隠しても入力検証は省略しません。

幅・高さの省略時は自動調整します。サブカードの種別・名前・プロトコル・ポート・参照先は
可能な範囲で同じヘッダー行にまとめ、同じ段のカードは各カードの実幅で詰めて配置します。
長いヘッダーだけを折り返し、非表示項目の領域は確保しません。明示した親サイズは内容を収める必要があります。
リスナールールは Listener ヘッダーの直下に `Priority / Conditions / Action` の表として配置します。
同一条件の候補値は `OR`、異なる条件は `AND` でセル内に要約し、既定ルールは `Otherwise` と表示します。
子の座標・幅・高さ・レイアウト weight は指定できません。
非表示カード宛ての明示的な `<connection>` は最も近い表示中の親へ接続し、
両端が同じ親に集約された線は描きません。元の ID と設定はソースに残ります。
非表示は秘密情報の削除ではありません。秘密鍵・client secret 自体は記載しないでください。

## 部品と AWS の対応

| 子タグ | 親 | パラメータ |
|---|---|---|
| `aws-listener-rule` | listener | `priority="1..50000\|default"` |
| `aws-rule-condition` | rule | `field`、HTTP ヘッダーの場合 `name` |
| `aws-rule-match` | condition / jwt-claim | `value` または `regex`、query-string の任意の `key` |
| `aws-rule-action` | rule | `type`、`order="1..50000"`（既定 1）、forward の `target-group` |
| `aws-forward-target` | forward action | `target-group`、`weight="0..999"`（既定 1、比率であり百分率ではない） |
| `aws-jwt-claim` | jwt-validation action | `name`、`format="single-string\|string-array\|space-separated-values"`、値は子の match |
| `aws-rule-transform` | rule | `type="host-header-rewrite\|url-rewrite"` |
| `aws-rule-rewrite` | transform | `regex`、`replace` |
| `aws-target-group` | ALB | `name`、`target-type="instance\|ip\|lambda"`、`protocol="HTTP\|HTTPS"`、`port`（Lambda は両方省略） |
| `aws-target-service` | target-group | `service="ecs\|eks\|ec2\|lambda\|ip"`、`name`、任意の外部 `ref` |
| `aws-registered-target` | target-group | `name`（IP / instance ID / Lambda 参照）、任意の `port`・`zone` |
| `aws-option` | ALB / listener / action / target-group / registered-target | 閉じた定義の `name`・`value`。`auth-parameter` のみ追加の `key` |

### 条件・優先順位

6 種類の `field` は `host-header`、`path-pattern`、`http-header`、
`http-request-method`、`query-string`、`source-ip` です。
異なる条件を AND、同じ条件内の match を OR として表示します。
ホスト・パス・ヘッダーは値比較または正規表現を選べます（混在不可）。
メソッドは文字列、送信元は IPv4 / IPv6 CIDR、クエリーは key/value または value を指定します。
1 条件 1–3 比較、1 ルール最大 5 比較・5 ワイルドカードを検証します。
条件の種類別重複制約、優先順位の重複も検証します。
数値の小さいルールから並び、default は最後です。default に条件・書き換えは置けません。
[AWS 条件](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/rule-condition-types.html)、
[AWS ルール](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-rules.html)。

### アクション・書き換え

`forward` / `redirect` / `fixed-response` のいずれかを最後に 1 個配置します。
その前に `authenticate-oidc` / `authenticate-cognito` / `jwt-validation` の
いずれかを追加できます。認証には HTTPS リスナーが必要です。
forward は単一の参照、または最大 5 個の重み付き転送先を指定します。
redirect の HTTP 301 / 302 と URL 部品、固定応答のステータス・Content-Type・本文、
認証の設定は `aws-option` で指定します。
[AWS アクション](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/rule-action-types.html)。

JWT は issuer / JWKS と追加 claim をカード化します。3 種類の claim 形式に対応し、
標準 claim と追加 claim を区別します。トークンや認証を実行する機能ではありません。
[AWS JWT](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-verify-jwt.html)。

書き換えはホスト名と URL をそれぞれ最大 1 個、各カード内に regex / replace を置きます。
条件で選択された後、転送前に適用する意味です。書き換え後の値で再ルーティングしません。
URL 書き換えはパス・クエリーの変更であり、プロトコル・ポートの変更ではありません。
[AWS transforms](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/rule-transforms.html)。

## 設定の範囲と限界

[オプション一覧](alb-options.md)に全実装キー、配置先、値の型を記載しています。
細かい `aws-option` はカードを増やさず、配置先ごとに `Setting / Value` の2列表へまとめます。
長い値はセル内で折り返し、非表示行は領域を残しません。ネットワーク、TLS / mTLS、ヘッダー変更、認証、スティッキーセッション、ヘルスチェック、
アルゴリズム、ATW、Target Optimizer、ログ、保護、容量、外部連携を表現します。
証明書、CA、CRL、サブネット、セキュリティグループ、連携先は参照・設定カードです。
これらを独立した AWS リソースとして作成・操作する機能ではありません。
ALB 内のターゲットグループは設定上のまとまりで、ネットワーク上の包含ではありません。
`aws-target-service` は ECS service、EKS/Kubernetes service、EC2 fleet、Lambda、または
一般のIPサービスとの論理的な関連を表します。その子の `aws-registered-target` が実際に
登録されるタスクIP、EC2 instance IDと動的host port、Pod IP、worker nodeとNodePort、
Lambda関数を表します。ECS `awsvpc` / Fargate は `ip`、ECS `bridge` / `host` は
`instance`、EKSはPod向け `ip` またはnode向け `instance` を検証します。
[ECSとALB](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/alb.html)、
[EKS load balancing](https://docs.aws.amazon.com/eks/latest/best-practices/load-balancing.html)。
通信線は自動生成せず、`connection` で明示します。

2026-09-06 に AWS 公式資料と照合した**図の表現モデル**です。
全条件・アクション・書き換え種類と一覧の設定を扱いますが、AWS API 全操作や将来追加される
機能まで網羅するデプロイ検証器ではありません。名前解決、AWS アカウント・リージョンの
提供可否、証明書検証、正規表現の実行、実トラフィックの評価はしません。
部分図のため、省略された既定ルールや外部ターゲット参照を許可します。
値の型・範囲、親子関係、TLS/認証、ローカルで宣言された HTTP2/GRPC ターゲットとの
互換性など、実装された組み合わせを検証します。ALB 専用の子タグは NLB と V1 ではエラーです。

## サンプル

[ALB サンプル一覧](../examples/samples/aws/aws-elastic-load-balancing-application-load-balancer/README.md#native-rule-and-option-examples)
には、同じ構成の詳細・標準・概要・選択表示と、機能ごとの `.xal` / `.svg` ペアがあります。
