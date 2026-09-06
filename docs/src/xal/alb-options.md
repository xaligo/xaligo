# ALB オプション一覧

[ALB 専用タグ](alb.md)の `aws-option` で指定する閉じた設定一覧です。
各設定は `visible="false"` で個別に非表示にできます。`standard` / `summary` で表示する場合は親に `show="options"` を指定してください。
`text` は空でない単一行の参照・注記であり、AWS に対する存在確認や完全な値検証は行いません。
`bool` は `true` / `false`、`int:min:max` は両端を含む整数、`named` は追加の `key` が必要です。
`status` は 200–299 / 400–599。`enum:` は列挙した値だけを受け付けます。
例示値は実環境の推奨設定ではありません。省略された設定は図に追加されません。

## ALB

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `scheme` | Scheme | `enum:internet-facing\|internal` | `internet-facing` |
| `ip-address-type` | IP family | `enum:ipv4\|dualstack\|dualstack-without-public-ipv4` | `dualstack` |
| `subnets` | Subnets | `text` | `subnet-a, subnet-c` |
| `security-groups` | Security groups | `text` | `sg-edge` |
| `ipam-pool` | IPAM pool | `text` | `ipam-pool-example` |
| `placement` | Placement | `enum:regional\|local-zone\|outpost` | `regional` |
| `waf` | WAF | `text` | `web-acl` |
| `shield` | Shield | `enum:standard\|advanced` | `advanced` |
| `capacity-units` | Reserved LCU | `int:0:1000000` | `100` |
| `capacity-reset` | Reset capacity | `bool` | `false` |
| `tags` | Tags | `text` | `env=production` |
| `cloudwatch` | CloudWatch | `text` | `5XX / latency alarm` |
| `cloudtrail` | CloudTrail | `text` | `audit trail` |
| `route53` | DNS alias | `text` | `api.example.test` |
| `global-accelerator` | Global Accelerator | `text` | `accelerator` |
| `cloudfront` | CloudFront | `text` | `distribution` |
| `websockets` | WebSocket | `bool` | `true` |
| `arc-zonal-autoshift` | ARC autoshift | `bool` | `true` |
| `deletion_protection.enabled` | Deletion protection | `bool` | `true` |
| `load_balancing.cross_zone.enabled` | Cross-zone | `enum:true` | `true` |
| `access_logs.s3.enabled` | Access logs | `bool` | `true` |
| `access_logs.s3.bucket` | Access log bucket | `text` | `alb-logs` |
| `access_logs.s3.prefix` | Access log prefix | `text` | `access/` |
| `connection_logs.s3.enabled` | Connection logs | `bool` | `true` |
| `connection_logs.s3.bucket` | Connection log bucket | `text` | `alb-logs` |
| `connection_logs.s3.prefix` | Connection log prefix | `text` | `connection/` |
| `health_check_logs.s3.enabled` | Health logs | `bool` | `true` |
| `health_check_logs.s3.bucket` | Health log bucket | `text` | `alb-logs` |
| `health_check_logs.s3.prefix` | Health log prefix | `text` | `health/` |
| `ipv6.deny_all_igw_traffic` | Block IPv6 IGW | `bool` | `false` |
| `zonal_shift.config.enabled` | Zonal shift | `bool` | `true` |
| `idle_timeout.timeout_seconds` | Idle timeout (s) | `int:1:4000` | `60` |
| `client_keep_alive.seconds` | Keepalive (s) | `int:60:604800` | `3600` |
| `routing.http.desync_mitigation_mode` | Desync protection | `enum:monitor\|defensive\|strictest` | `defensive` |
| `routing.http.drop_invalid_header_fields.enabled` | Drop invalid headers | `bool` | `true` |
| `routing.http.preserve_host_header.enabled` | Preserve host | `bool` | `true` |
| `routing.http.x_amzn_tls_version_and_cipher_suite.enabled` | TLS headers | `bool` | `true` |
| `routing.http.xff_client_port.enabled` | XFF client port | `bool` | `true` |
| `routing.http.xff_header_processing.mode` | XFF | `enum:append\|preserve\|remove` | `append` |
| `routing.http2.enabled` | HTTP/2 | `bool` | `true` |
| `waf.fail_open.enabled` | WAF fail-open | `bool` | `false` |
| `aws-config` | AWS Config | `text` | `compliance recorder` |
| `state` | State | `enum:provisioning\|active\|active_impaired\|failed` | `active` |

## Listener

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `tls-policy` | TLS policy | `text` | `ELBSecurityPolicy-TLS13-1-2-2021-06` |
| `sni-certificates` | SNI certificates | `text` | `api-cert, admin-cert` |
| `mtls-ignore-expiry` | Ignore client expiry | `bool` | `false` |
| `mtls-advertise-ca` | Advertise CA | `bool` | `true` |
| `trust-store-ca-bundle` | CA bundle | `text` | `s3://trust/ca.pem` |
| `trust-store-crls` | CRLs | `text` | `s3://trust/revoked.pem` |
| `routing.http.request.x_amzn_mtls_clientcert_serial_number.header_name` | mTLS Serial header | `text` | `X-Client-Serial` |
| `routing.http.request.x_amzn_mtls_clientcert_issuer.header_name` | mTLS Issuer header | `text` | `X-Client-Issuer` |
| `routing.http.request.x_amzn_mtls_clientcert_subject.header_name` | mTLS Subject header | `text` | `X-Client-Subject` |
| `routing.http.request.x_amzn_mtls_clientcert_validity.header_name` | mTLS Validity header | `text` | `X-Client-Validity` |
| `routing.http.request.x_amzn_mtls_clientcert_leaf.header_name` | mTLS Leaf header | `text` | `X-Client-Leaf` |
| `routing.http.request.x_amzn_mtls_clientcert.header_name` | mTLS Chain header | `text` | `X-Client-Chain` |
| `routing.http.request.x_amzn_tls_version.header_name` | TLS version header | `text` | `X-TLS-version` |
| `routing.http.request.x_amzn_tls_cipher_suite.header_name` | TLS cipher header | `text` | `X-TLS-cipher` |
| `routing.http.response.server.enabled` | Server header | `bool` | `false` |
| `routing.http.response.strict_transport_security.header_value` | HSTS | `text` | `max-age=31536000` |
| `routing.http.response.access_control_allow_origin.header_value` | CORS origin | `text` | `https://app.example.test` |
| `routing.http.response.access_control_allow_methods.header_value` | CORS methods | `text` | `GET, POST` |
| `routing.http.response.access_control_allow_headers.header_value` | CORS headers | `text` | `Authorization` |
| `routing.http.response.access_control_allow_credentials.header_value` | CORS credentials | `text` | `true` |
| `routing.http.response.access_control_expose_headers.header_value` | CORS expose | `text` | `X-Request-Id` |
| `routing.http.response.access_control_max_age.header_value` | CORS max-age | `text` | `600` |
| `routing.http.response.content_security_policy.header_value` | CSP | `text` | `default-src 'self'` |
| `routing.http.response.x_content_type_options.header_value` | Content type options | `text` | `nosniff` |
| `routing.http.response.x_frame_options.header_value` | Frame options | `text` | `DENY` |

## Forward

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `stickiness` | Group stickiness | `bool` | `true` |
| `stickiness-duration` | Group cookie (s) | `int:1:604800` | `3600` |

## Redirect

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `protocol` | Protocol | `enum:HTTP\|HTTPS\|#{protocol}` | `HTTPS` |
| `host` | Host | `text` | `#{host}` |
| `port` | Port | `text` | `443` |
| `path` | Path | `text` | `/#{path}` |
| `query` | Query | `text` | `#{query}` |
| `status-code` | Status | `enum:HTTP_301\|HTTP_302` | `HTTP_301` |

## Fixed response

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `status-code` | Status | `status` | `503` |
| `content-type` | Content type | `enum:text/plain\|text/css\|text/html\|application/javascript\|application/json` | `text/plain` |
| `message-body` | Body | `text` | `Temporarily unavailable` |

## OIDC / Cognito

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `client-id` | Client ID | `text` | `client-id` |
| `on-unauthenticated-request` | Unauthenticated | `enum:deny\|allow\|authenticate` | `authenticate` |
| `scope` | Scopes | `text` | `openid email` |
| `session-cookie` | Session cookie | `text` | `ALBAuthSession` |
| `session-timeout` | Session (s) | `int:1:604800` | `3600` |
| `auth-parameter` | Auth parameter | `named` | `login` |

## OIDC

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `issuer` | Issuer | `text` | `https://id.example.test` |
| `authorization-endpoint` | Authorize | `text` | `https://id.example.test/authorize` |
| `token-endpoint` | Token | `text` | `https://id.example.test/token` |
| `user-info-endpoint` | User info | `text` | `https://id.example.test/userinfo` |
| `client-secret-ref` | Secret reference | `text` | `secret:oidc-client` |
| `use-existing-client-secret` | Existing secret | `bool` | `true` |

## Cognito

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `user-pool` | User pool | `text` | `user-pool-arn` |
| `user-pool-domain` | Pool domain | `text` | `login.example.test` |

## JWT validation

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `jwt-issuer` | JWT issuer | `text` | `https://id.example.test` |
| `jwks-endpoint` | JWKS | `text` | `https://id.example.test/.well-known/jwks.json` |

## Target group

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `protocol-version` | HTTP version | `enum:HTTP1\|HTTP2\|GRPC` | `HTTP1` |
| `vpc` | VPC | `text` | `vpc-production` |
| `target-ip-address-type` | Target IP family | `enum:ipv4\|ipv6` | `ipv4` |
| `health-check-enabled` | Health check | `bool` | `true` |
| `health-check-protocol` | Health protocol | `enum:HTTP\|HTTPS` | `HTTP` |
| `health-check-port` | Health port | `text` | `traffic-port` |
| `health-check-path` | Health path | `text` | `/health` |
| `health-check-interval` | Health interval (s) | `int:5:300` | `30` |
| `health-check-timeout` | Health timeout (s) | `int:2:120` | `5` |
| `healthy-threshold` | Healthy threshold | `int:2:10` | `5` |
| `unhealthy-threshold` | Unhealthy threshold | `int:2:10` | `2` |
| `health-check-matcher` | Success codes | `text` | `200-299` |
| `target-control-port` | Optimizer port | `int:1:65535` | `8081` |
| `deregistration_delay.timeout_seconds` | Drain (s) | `int:0:3600` | `300` |
| `stickiness.enabled` | Target stickiness | `bool` | `true` |
| `stickiness.type` | Cookie type | `enum:lb_cookie\|app_cookie` | `lb_cookie` |
| `stickiness.app_cookie.cookie_name` | App cookie | `text` | `SESSION` |
| `stickiness.app_cookie.duration_seconds` | App cookie (s) | `int:1:604800` | `86400` |
| `stickiness.lb_cookie.duration_seconds` | LB cookie (s) | `int:1:604800` | `86400` |
| `load_balancing.algorithm.type` | Algorithm | `enum:round_robin\|least_outstanding_requests\|weighted_random` | `round_robin` |
| `load_balancing.algorithm.anomaly_mitigation` | Anomaly mitigation | `enum:on\|off` | `off` |
| `slow_start.duration_seconds` | Slow start (s) | `int:0:900` | `30` |
| `target-cross-zone` | Cross-zone override | `enum:true\|false\|use_load_balancer_configuration` | `use_load_balancer_configuration` |
| `lambda.multi_value_headers.enabled` | Lambda multi-headers | `bool` | `false` |
| `target_group_health.dns_failover.minimum_healthy_targets.count` | DNS minimum targets | `text` | `1` |
| `target_group_health.dns_failover.minimum_healthy_targets.percentage` | DNS healthy (%) | `text` | `off` |
| `target_group_health.unhealthy_state_routing.minimum_healthy_targets.count` | Fail-open minimum targets | `text` | `1` |
| `target_group_health.unhealthy_state_routing.minimum_healthy_targets.percentage` | Fail-open healthy (%) | `text` | `off` |
| `waf.http2.traffic_inspection_behavior` | WAF HTTP/2 inspection | `text` | `inspect_after_sufficient_data` |

## Registered target

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `health` | Health | `enum:initial\|healthy\|unhealthy\|unhealthy.draining\|unused\|draining\|unavailable` | `healthy` |
| `health-reason` | Reason | `text` | `Target.ResponseCodeMismatch` |
| `anomaly` | Anomaly | `enum:normal\|anomalous` | `normal` |
| `administrative-override` | Override | `text` | `zonal_shift_active` |
| `optimizer-agent` | Optimizer agent | `text` | `concurrency=100; data=:8080` |

## Target service

| name | 表示 | 型 | 例 |
|---|---|---|---|
| `cluster` | Cluster | `text` | `production` |
| `launch-type` | Launch type | `enum:fargate\|ec2` | `fargate` |
| `network-mode` | Network mode | `enum:awsvpc\|bridge\|host` | `awsvpc` |
| `container-name` | Container | `text` | `orders-api` |
| `container-port` | Container port | `int:1:65535` | `8080` |
| `namespace` | Namespace | `text` | `apps` |
| `kubernetes-service-port` | Service port | `int:1:65535` | `80` |
| `node-port` | NodePort | `int:1:65535` | `30080` |
| `autoscaling-group` | Auto Scaling group | `text` | `web-asg` |

これらは `aws-target-service` の設定です。ECS/EKS等をELBの登録target-typeとして
追加するものではありません。登録方式はtarget groupの `ip` / `instance` / `lambda`、
実体は子の `aws-registered-target` で表します。

## AWS の出典と対応

- [LoadBalancer attributes](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancerAttribute.html): ネットワーク、保護、ログ、HTTP 動作。
- [Listener attributes](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_ListenerAttribute.html): ヘッダー名とレスポンスヘッダー。
- [Listener](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html)・[mTLS](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/mutual-authentication.html): TLS policy、SNI、CA / CRL の参照。
- [Action](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Action.html): forward / redirect / fixed response / OIDC / Cognito / JWT の設定。
- [TargetGroup attributes](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroupAttribute.html)・[Target groups](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html): ターゲット、ヘルスチェック、転送動作。
- [Integrations](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-integrations.html): 外部サービス参照。

XAL の短い名前は図向けの別名です。例えば `target-cross-zone` はターゲットグループの `load_balancing.cross_zone.enabled` に対応します。ALB 側の同名 AWS 属性は `true` 固定です。
`waf` / `shield` / `route53` / `cloudfront` / `aws-config` 等は連携を示す注記で、同名の ALB API パラメータではありません。
`client-secret-ref` は秘密値の代わりに図上の参照名を置くための項目です。
