#[path = "../../../../external/engine/src/cnf/engine_abi.rs"]
pub mod engine_abi_source;

pub mod cnf {
    pub mod engine {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/cnf/engine.rs"
        ));
    }
    pub use crate::engine_abi_source as engine_abi;
}

pub mod ent {
    pub mod model {
        pub mod aws {
            pub mod action {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/action.rs"
                ));
            }
            pub mod condition {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/condition.rs"
                ));
            }
            pub mod feature {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/feature.rs"
                ));
            }
            pub mod option {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/option.rs"
                ));
            }
            pub mod presentation {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/presentation.rs"
                ));
            }
            pub mod rule {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/rule.rs"
                ));
            }
            pub mod target_service {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/target_service.rs"
                ));
            }
            pub mod target_group {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/target_group.rs"
                ));
            }
            pub mod transform {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/transform.rs"
                ));
            }
            pub mod alb {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/alb.rs"
                ));
            }
            pub mod component {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/component.rs"
                ));
            }
            pub mod listener {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/listener.rs"
                ));
            }
            pub mod nlb {
                include!(concat!(
                    env!("CARGO_MANIFEST_DIR"),
                    "/../../../../external/engine/src/ent/model/aws/nlb.rs"
                ));
            }

            pub use component::Component;
        }
        pub mod document {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/ent/model/document.rs"
            ));
        }
        pub mod svg {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/ent/model/svg.rs"
            ));
        }
    }
    pub mod request {
        pub mod engine {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/ent/request/engine.rs"
            ));
        }
    }
    pub mod response {
        pub mod engine {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/ent/response/engine.rs"
            ));
        }
    }
}

pub mod util {
    pub mod clone {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/clone.rs"
        ));
    }
    pub mod debug {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/debug.rs"
        ));
    }
    pub mod default {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/default.rs"
        ));
    }
    pub mod deserialize {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/deserialize.rs"
        ));
        mod tests {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/util/deserialize_aws_test.rs"
            ));
        }
    }
    pub mod eq {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/eq.rs"
        ));
    }
    pub mod error {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/error.rs"
        ));
    }
    pub mod logger {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/logger.rs"
        ));
        mod tests {
            include!(concat!(env!("CARGO_MANIFEST_DIR"), "/util/logger_test.rs"));
        }
    }
    pub mod mcode {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/mcode.rs"
        ));
        mod tests {
            include!(concat!(env!("CARGO_MANIFEST_DIR"), "/util/mcode_test.rs"));
        }
    }
    pub mod serialize {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/util/serialize.rs"
        ));
    }
}

pub mod usc {
    pub mod aws {
        pub mod target_service {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/target_service.rs"
            ));
        }
        pub mod option_table {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/option_table.rs"
            ));
        }
        pub mod rule_table {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/rule_table.rs"
            ));
        }
        pub mod action {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/action.rs"
            ));
        }
        pub mod condition {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/condition.rs"
            ));
        }
        pub mod detail {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/detail.rs"
            ));
        }
        pub mod option {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/option.rs"
            ));
        }
        pub mod presentation {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/presentation.rs"
            ));
        }
        pub mod rule {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/rule.rs"
            ));
        }
        pub mod target_group {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/target_group.rs"
            ));
        }
        pub mod transform {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/transform.rs"
            ));
        }
        pub mod alb {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/alb.rs"
            ));
        }
        mod composition {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/composition.rs"
            ));
        }
        mod drawing {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/drawing.rs"
            ));
        }
        pub mod listener {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/listener.rs"
            ));
        }
        mod load_balancer {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/load_balancer.rs"
            ));
        }
        pub mod nlb {
            include!(concat!(
                env!("CARGO_MANIFEST_DIR"),
                "/../../../../external/engine/src/usc/aws/nlb.rs"
            ));
        }

        pub(crate) use composition::compose;
    }
    pub mod cancel {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/usc/cancel.rs"
        ));
    }
    pub mod engine {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/usc/engine.rs"
        ));
    }
    pub mod layout {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/usc/layout.rs"
        ));
        include!(concat!(env!("CARGO_MANIFEST_DIR"), "/usc/layout_test.rs"));
    }
    pub mod svg {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/usc/svg.rs"
        ));
        mod tests {
            include!(concat!(env!("CARGO_MANIFEST_DIR"), "/usc/svg_test.rs"));
        }
    }
}

mod base {
    include!(concat!(
        env!("CARGO_MANIFEST_DIR"),
        "/../../../../external/engine/src/base.rs"
    ));
}

pub mod ctl {
    pub mod engine {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/../../../../external/engine/src/ctl/engine.rs"
        ));
        mod tests {
            include!(concat!(env!("CARGO_MANIFEST_DIR"), "/ctl/engine_test.rs"));
        }
    }
}
