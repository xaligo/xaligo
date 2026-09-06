use super::*;
use crate::cnf::engine_abi::*;

#[test]
fn native_alb_features_preserve_zero_weight_and_reject_foreign_fields() {
    use crate::ent::model::aws::{Component, feature::FeatureKind};
    let mut s: [String; STRING_FIELD_COUNT] = std::array::from_fn(|_| String::new());
    s[STRING_AWS_KIND] = "forward-target".into();
    s[STRING_AWS_VALUE] = "canary".into();
    s[STRING_AWS_ORDER] = "0".into();
    let decoded = decode_aws_component(&s, None, None, None, None)
        .unwrap()
        .unwrap();
    assert!(
        matches!(decoded, Component::Feature(f) if matches!(&f.kind,FeatureKind::ForwardTarget(t) if t.weight==0))
    );
    s[STRING_AWS_TYPE] = "unrelated".into();
    assert!(decode_aws_component(&s, None, None, None, None).is_err());
    s[STRING_AWS_TYPE].clear();
    s[STRING_AWS_ORDER] = "1000".into();
    assert!(decode_aws_component(&s, None, None, None, None).is_err());
    s[STRING_AWS_ORDER] = "1".into();
    s[STRING_AWS_HIDE] = "options".into();
    s[STRING_AWS_SHOW] = "options".into();
    assert!(decode_aws_component(&s, None, None, None, None).is_err());
    s[STRING_AWS_SHOW].clear();
    s[STRING_AWS_DETAIL_LEVEL] = "summary".into();
    assert!(decode_aws_component(&s, None, None, None, None).is_ok());
}

#[test]
fn native_aws_decoder_rejects_untyped_or_invalid_payloads() {
    let mut fields: [String; STRING_FIELD_COUNT] = std::array::from_fn(|_| String::new());
    assert!(
        decode_aws_component(&fields, None, None, None, None)
            .unwrap()
            .is_none()
    );
    assert!(decode_aws_component(&fields, Some(443.0), None, None, None).is_err());
    assert!(decode_aws_component(&fields, None, Some(false), None, None).is_err());
    assert!(decode_aws_component(&fields, None, None, None, Some(false)).is_err());
    fields[STRING_AWS_KIND] = "unknown".into();
    assert!(decode_aws_component(&fields, None, None, None, None).is_err());
    fields[STRING_AWS_KIND] = "listener".into();
    fields[STRING_AWS_PROTOCOL] = "HTTPS".into();
    for value in [
        None,
        Some(0.0),
        Some(65536.0),
        Some(443.5),
        Some(f64::NAN),
        Some(f64::INFINITY),
    ] {
        assert!(decode_aws_component(&fields, value, None, None, None).is_err());
    }
    assert!(
        decode_aws_component(&fields, Some(65535.0), Some(true), Some(false), Some(false)).is_ok()
    );
    fields[STRING_AWS_MUTUAL_TLS] = "on".into();
    assert!(decode_aws_component(&fields, Some(443.0), None, None, None).is_err());
    fields[STRING_AWS_MUTUAL_TLS].clear();
    fields[STRING_AWS_DOMAIN] = "example.test".into();
    assert!(decode_aws_component(&fields, Some(443.0), None, None, None).is_err());
}
