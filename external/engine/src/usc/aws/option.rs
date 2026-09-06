use crate::ent::model::aws::option::OptionSetting;
use crate::util::error::LayoutError;

pub struct OptionDefinition {
    pub name: &'static str,
    pub label: &'static str,
    pub owners: &'static str,
    pub constraint: &'static str,
}
// Closed ALB diagram schema; no arbitrary keys or dynamic maps enter layout.
pub static DEFINITIONS: &[OptionDefinition] = &[
    OptionDefinition {
        name: "scheme",
        label: "Scheme",
        owners: "alb",
        constraint: "enum:internet-facing|internal",
    },
    OptionDefinition {
        name: "ip-address-type",
        label: "IP family",
        owners: "alb",
        constraint: "enum:ipv4|dualstack|dualstack-without-public-ipv4",
    },
    OptionDefinition {
        name: "subnets",
        label: "Subnets",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "security-groups",
        label: "Security groups",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "ipam-pool",
        label: "IPAM pool",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "placement",
        label: "Placement",
        owners: "alb",
        constraint: "enum:regional|local-zone|outpost",
    },
    OptionDefinition {
        name: "waf",
        label: "WAF",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "shield",
        label: "Shield",
        owners: "alb",
        constraint: "enum:standard|advanced",
    },
    OptionDefinition {
        name: "capacity-units",
        label: "Reserved LCU",
        owners: "alb",
        constraint: "int:0:1000000",
    },
    OptionDefinition {
        name: "capacity-reset",
        label: "Reset capacity",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "tags",
        label: "Tags",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "cloudwatch",
        label: "CloudWatch",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "cloudtrail",
        label: "CloudTrail",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "route53",
        label: "DNS alias",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "global-accelerator",
        label: "Global Accelerator",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "cloudfront",
        label: "CloudFront",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "websockets",
        label: "WebSocket",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "arc-zonal-autoshift",
        label: "ARC autoshift",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "deletion_protection.enabled",
        label: "Deletion protection",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "load_balancing.cross_zone.enabled",
        label: "Cross-zone",
        owners: "alb",
        constraint: "enum:true",
    },
    OptionDefinition {
        name: "access_logs.s3.enabled",
        label: "Access logs",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "access_logs.s3.bucket",
        label: "Access log bucket",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "access_logs.s3.prefix",
        label: "Access log prefix",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "connection_logs.s3.enabled",
        label: "Connection logs",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "connection_logs.s3.bucket",
        label: "Connection log bucket",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "connection_logs.s3.prefix",
        label: "Connection log prefix",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "health_check_logs.s3.enabled",
        label: "Health logs",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "health_check_logs.s3.bucket",
        label: "Health log bucket",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "health_check_logs.s3.prefix",
        label: "Health log prefix",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "ipv6.deny_all_igw_traffic",
        label: "Block IPv6 IGW",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "zonal_shift.config.enabled",
        label: "Zonal shift",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "idle_timeout.timeout_seconds",
        label: "Idle timeout (s)",
        owners: "alb",
        constraint: "int:1:4000",
    },
    OptionDefinition {
        name: "client_keep_alive.seconds",
        label: "Keepalive (s)",
        owners: "alb",
        constraint: "int:60:604800",
    },
    OptionDefinition {
        name: "routing.http.desync_mitigation_mode",
        label: "Desync protection",
        owners: "alb",
        constraint: "enum:monitor|defensive|strictest",
    },
    OptionDefinition {
        name: "routing.http.drop_invalid_header_fields.enabled",
        label: "Drop invalid headers",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "routing.http.preserve_host_header.enabled",
        label: "Preserve host",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "routing.http.x_amzn_tls_version_and_cipher_suite.enabled",
        label: "TLS headers",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "routing.http.xff_client_port.enabled",
        label: "XFF client port",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "routing.http.xff_header_processing.mode",
        label: "XFF",
        owners: "alb",
        constraint: "enum:append|preserve|remove",
    },
    OptionDefinition {
        name: "routing.http2.enabled",
        label: "HTTP/2",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "waf.fail_open.enabled",
        label: "WAF fail-open",
        owners: "alb",
        constraint: "bool",
    },
    OptionDefinition {
        name: "tls-policy",
        label: "TLS policy",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "sni-certificates",
        label: "SNI certificates",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "mtls-ignore-expiry",
        label: "Ignore client expiry",
        owners: "listener",
        constraint: "bool",
    },
    OptionDefinition {
        name: "mtls-advertise-ca",
        label: "Advertise CA",
        owners: "listener",
        constraint: "bool",
    },
    OptionDefinition {
        name: "trust-store-ca-bundle",
        label: "CA bundle",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "trust-store-crls",
        label: "CRLs",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.request.x_amzn_mtls_clientcert_serial_number.header_name",
        label: "mTLS Serial header",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.request.x_amzn_mtls_clientcert_issuer.header_name",
        label: "mTLS Issuer header",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.request.x_amzn_mtls_clientcert_subject.header_name",
        label: "mTLS Subject header",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.request.x_amzn_mtls_clientcert_validity.header_name",
        label: "mTLS Validity header",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.request.x_amzn_mtls_clientcert_leaf.header_name",
        label: "mTLS Leaf header",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.request.x_amzn_mtls_clientcert.header_name",
        label: "mTLS Chain header",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.request.x_amzn_tls_version.header_name",
        label: "TLS version header",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.request.x_amzn_tls_cipher_suite.header_name",
        label: "TLS cipher header",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.server.enabled",
        label: "Server header",
        owners: "listener",
        constraint: "bool",
    },
    OptionDefinition {
        name: "routing.http.response.strict_transport_security.header_value",
        label: "HSTS",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.access_control_allow_origin.header_value",
        label: "CORS origin",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.access_control_allow_methods.header_value",
        label: "CORS methods",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.access_control_allow_headers.header_value",
        label: "CORS headers",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.access_control_allow_credentials.header_value",
        label: "CORS credentials",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.access_control_expose_headers.header_value",
        label: "CORS expose",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.access_control_max_age.header_value",
        label: "CORS max-age",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.content_security_policy.header_value",
        label: "CSP",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.x_content_type_options.header_value",
        label: "Content type options",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "routing.http.response.x_frame_options.header_value",
        label: "Frame options",
        owners: "listener",
        constraint: "text",
    },
    OptionDefinition {
        name: "stickiness",
        label: "Group stickiness",
        owners: "forward",
        constraint: "bool",
    },
    OptionDefinition {
        name: "stickiness-duration",
        label: "Group cookie (s)",
        owners: "forward",
        constraint: "int:1:604800",
    },
    OptionDefinition {
        name: "protocol",
        label: "Protocol",
        owners: "redirect",
        constraint: "enum:HTTP|HTTPS|#{protocol}",
    },
    OptionDefinition {
        name: "host",
        label: "Host",
        owners: "redirect",
        constraint: "text",
    },
    OptionDefinition {
        name: "port",
        label: "Port",
        owners: "redirect",
        constraint: "text",
    },
    OptionDefinition {
        name: "path",
        label: "Path",
        owners: "redirect",
        constraint: "text",
    },
    OptionDefinition {
        name: "query",
        label: "Query",
        owners: "redirect",
        constraint: "text",
    },
    OptionDefinition {
        name: "status-code",
        label: "Status",
        owners: "redirect",
        constraint: "enum:HTTP_301|HTTP_302",
    },
    OptionDefinition {
        name: "status-code",
        label: "Status",
        owners: "fixed",
        constraint: "status",
    },
    OptionDefinition {
        name: "content-type",
        label: "Content type",
        owners: "fixed",
        constraint: "enum:text/plain|text/css|text/html|application/javascript|application/json",
    },
    OptionDefinition {
        name: "message-body",
        label: "Body",
        owners: "fixed",
        constraint: "text",
    },
    OptionDefinition {
        name: "client-id",
        label: "Client ID",
        owners: "oidc cognito",
        constraint: "text",
    },
    OptionDefinition {
        name: "on-unauthenticated-request",
        label: "Unauthenticated",
        owners: "oidc cognito",
        constraint: "enum:deny|allow|authenticate",
    },
    OptionDefinition {
        name: "scope",
        label: "Scopes",
        owners: "oidc cognito",
        constraint: "text",
    },
    OptionDefinition {
        name: "session-cookie",
        label: "Session cookie",
        owners: "oidc cognito",
        constraint: "text",
    },
    OptionDefinition {
        name: "session-timeout",
        label: "Session (s)",
        owners: "oidc cognito",
        constraint: "int:1:604800",
    },
    OptionDefinition {
        name: "auth-parameter",
        label: "Auth parameter",
        owners: "oidc cognito",
        constraint: "named",
    },
    OptionDefinition {
        name: "issuer",
        label: "Issuer",
        owners: "oidc",
        constraint: "text",
    },
    OptionDefinition {
        name: "authorization-endpoint",
        label: "Authorize",
        owners: "oidc",
        constraint: "text",
    },
    OptionDefinition {
        name: "token-endpoint",
        label: "Token",
        owners: "oidc",
        constraint: "text",
    },
    OptionDefinition {
        name: "user-info-endpoint",
        label: "User info",
        owners: "oidc",
        constraint: "text",
    },
    OptionDefinition {
        name: "client-secret-ref",
        label: "Secret reference",
        owners: "oidc",
        constraint: "text",
    },
    OptionDefinition {
        name: "use-existing-client-secret",
        label: "Existing secret",
        owners: "oidc",
        constraint: "bool",
    },
    OptionDefinition {
        name: "user-pool",
        label: "User pool",
        owners: "cognito",
        constraint: "text",
    },
    OptionDefinition {
        name: "user-pool-domain",
        label: "Pool domain",
        owners: "cognito",
        constraint: "text",
    },
    OptionDefinition {
        name: "jwt-issuer",
        label: "JWT issuer",
        owners: "jwt",
        constraint: "text",
    },
    OptionDefinition {
        name: "jwks-endpoint",
        label: "JWKS",
        owners: "jwt",
        constraint: "text",
    },
    OptionDefinition {
        name: "protocol-version",
        label: "HTTP version",
        owners: "tg",
        constraint: "enum:HTTP1|HTTP2|GRPC",
    },
    OptionDefinition {
        name: "vpc",
        label: "VPC",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "target-ip-address-type",
        label: "Target IP family",
        owners: "tg",
        constraint: "enum:ipv4|ipv6",
    },
    OptionDefinition {
        name: "health-check-enabled",
        label: "Health check",
        owners: "tg",
        constraint: "bool",
    },
    OptionDefinition {
        name: "health-check-protocol",
        label: "Health protocol",
        owners: "tg",
        constraint: "enum:HTTP|HTTPS",
    },
    OptionDefinition {
        name: "health-check-port",
        label: "Health port",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "health-check-path",
        label: "Health path",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "health-check-interval",
        label: "Health interval (s)",
        owners: "tg",
        constraint: "int:5:300",
    },
    OptionDefinition {
        name: "health-check-timeout",
        label: "Health timeout (s)",
        owners: "tg",
        constraint: "int:2:120",
    },
    OptionDefinition {
        name: "healthy-threshold",
        label: "Healthy threshold",
        owners: "tg",
        constraint: "int:2:10",
    },
    OptionDefinition {
        name: "unhealthy-threshold",
        label: "Unhealthy threshold",
        owners: "tg",
        constraint: "int:2:10",
    },
    OptionDefinition {
        name: "health-check-matcher",
        label: "Success codes",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "target-control-port",
        label: "Optimizer port",
        owners: "tg",
        constraint: "int:1:65535",
    },
    OptionDefinition {
        name: "deregistration_delay.timeout_seconds",
        label: "Drain (s)",
        owners: "tg",
        constraint: "int:0:3600",
    },
    OptionDefinition {
        name: "stickiness.enabled",
        label: "Target stickiness",
        owners: "tg",
        constraint: "bool",
    },
    OptionDefinition {
        name: "stickiness.type",
        label: "Cookie type",
        owners: "tg",
        constraint: "enum:lb_cookie|app_cookie",
    },
    OptionDefinition {
        name: "stickiness.app_cookie.cookie_name",
        label: "App cookie",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "stickiness.app_cookie.duration_seconds",
        label: "App cookie (s)",
        owners: "tg",
        constraint: "int:1:604800",
    },
    OptionDefinition {
        name: "stickiness.lb_cookie.duration_seconds",
        label: "LB cookie (s)",
        owners: "tg",
        constraint: "int:1:604800",
    },
    OptionDefinition {
        name: "load_balancing.algorithm.type",
        label: "Algorithm",
        owners: "tg",
        constraint: "enum:round_robin|least_outstanding_requests|weighted_random",
    },
    OptionDefinition {
        name: "load_balancing.algorithm.anomaly_mitigation",
        label: "Anomaly mitigation",
        owners: "tg",
        constraint: "enum:on|off",
    },
    OptionDefinition {
        name: "slow_start.duration_seconds",
        label: "Slow start (s)",
        owners: "tg",
        constraint: "int:0:900",
    },
    OptionDefinition {
        name: "target-cross-zone",
        label: "Cross-zone override",
        owners: "tg",
        constraint: "enum:true|false|use_load_balancer_configuration",
    },
    OptionDefinition {
        name: "lambda.multi_value_headers.enabled",
        label: "Lambda multi-headers",
        owners: "tg",
        constraint: "bool",
    },
    OptionDefinition {
        name: "target_group_health.dns_failover.minimum_healthy_targets.count",
        label: "DNS minimum targets",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "target_group_health.dns_failover.minimum_healthy_targets.percentage",
        label: "DNS healthy (%)",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "target_group_health.unhealthy_state_routing.minimum_healthy_targets.count",
        label: "Fail-open minimum targets",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "target_group_health.unhealthy_state_routing.minimum_healthy_targets.percentage",
        label: "Fail-open healthy (%)",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "health",
        label: "Health",
        owners: "target",
        constraint: "enum:initial|healthy|unhealthy|unhealthy.draining|unused|draining|unavailable",
    },
    OptionDefinition {
        name: "health-reason",
        label: "Reason",
        owners: "target",
        constraint: "text",
    },
    OptionDefinition {
        name: "anomaly",
        label: "Anomaly",
        owners: "target",
        constraint: "enum:normal|anomalous",
    },
    OptionDefinition {
        name: "administrative-override",
        label: "Override",
        owners: "target",
        constraint: "text",
    },
    OptionDefinition {
        name: "aws-config",
        label: "AWS Config",
        owners: "alb",
        constraint: "text",
    },
    OptionDefinition {
        name: "state",
        label: "State",
        owners: "alb",
        constraint: "enum:provisioning|active|active_impaired|failed",
    },
    OptionDefinition {
        name: "waf.http2.traffic_inspection_behavior",
        label: "WAF HTTP/2 inspection",
        owners: "tg",
        constraint: "text",
    },
    OptionDefinition {
        name: "optimizer-agent",
        label: "Optimizer agent",
        owners: "target",
        constraint: "text",
    },
    OptionDefinition {
        name: "cluster",
        label: "Cluster",
        owners: "ecs eks",
        constraint: "text",
    },
    OptionDefinition {
        name: "launch-type",
        label: "Launch type",
        owners: "ecs",
        constraint: "enum:fargate|ec2",
    },
    OptionDefinition {
        name: "network-mode",
        label: "Network mode",
        owners: "ecs",
        constraint: "enum:awsvpc|bridge|host",
    },
    OptionDefinition {
        name: "container-name",
        label: "Container",
        owners: "ecs",
        constraint: "text",
    },
    OptionDefinition {
        name: "container-port",
        label: "Container port",
        owners: "ecs",
        constraint: "int:1:65535",
    },
    OptionDefinition {
        name: "namespace",
        label: "Namespace",
        owners: "eks",
        constraint: "text",
    },
    OptionDefinition {
        name: "kubernetes-service-port",
        label: "Service port",
        owners: "eks",
        constraint: "int:1:65535",
    },
    OptionDefinition {
        name: "node-port",
        label: "NodePort",
        owners: "eks",
        constraint: "int:1:65535",
    },
    OptionDefinition {
        name: "autoscaling-group",
        label: "Auto Scaling group",
        owners: "ec2",
        constraint: "text",
    },
];
pub fn key(name: &str) -> Result<usize, String> {
    DEFINITIONS
        .iter()
        .position(|d| d.name == name)
        .ok_or_else(|| format!("unknown ALB option {name:?}"))
}
pub(super) fn validate(option: &OptionSetting, owner: &str) -> Result<(), LayoutError> {
    let original = DEFINITIONS
        .get(option.key)
        .ok_or_else(|| LayoutError::new("invalid ALB option key"))?;
    let d = DEFINITIONS
        .iter()
        .find(|d| d.name == original.name && d.owners.split_whitespace().any(|o| o == owner))
        .unwrap_or(original);
    if !d.owners.split_whitespace().any(|o| o == owner) {
        return Err(LayoutError::new(format!(
            "ALB option {} is not valid on {owner}",
            d.name
        )));
    }
    let value = option.value.as_str();
    let valid = if d.constraint == "bool" {
        matches!(value, "true" | "false")
    } else if let Some(values) = d.constraint.strip_prefix("enum:") {
        values.split('|').any(|v| v == value)
    } else if let Some(range) = d.constraint.strip_prefix("int:") {
        let (min, max) = range.split_once(':').unwrap();
        value
            .parse::<u32>()
            .is_ok_and(|n| (min.parse::<u32>().unwrap()..=max.parse::<u32>().unwrap()).contains(&n))
    } else if d.constraint == "status" {
        value.len() == 3
            && value
                .parse::<u16>()
                .is_ok_and(|n| (200..=299).contains(&n) || (400..=599).contains(&n))
    } else if d.constraint == "named" {
        !option.name.is_empty() && !value.is_empty()
    } else {
        !value.is_empty()
    };
    if !valid {
        return Err(LayoutError::new(format!(
            "invalid ALB option {} value",
            d.name
        )));
    }
    Ok(())
}
