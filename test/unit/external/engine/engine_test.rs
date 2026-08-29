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
