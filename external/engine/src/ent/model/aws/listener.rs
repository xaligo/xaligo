pub enum Protocol {
    Http,
    Https,
    Tcp,
    Tls,
    Udp,
    TcpUdp,
    Quic,
    TcpQuic,
}

pub enum MutualTls {
    Off,
    Verify,
    Passthrough,
}

pub struct Listener {
    pub presentation: super::presentation::Presentation,
    pub protocol: Protocol,
    pub port: u16,
    pub mutual_tls: MutualTls,
    pub certificate: String,
    pub trust_store: String,
    pub target_group: String,
    pub backend_tls: Option<bool>,
    pub backend_mtls: Option<bool>,
    pub show_title: Option<bool>,
}
