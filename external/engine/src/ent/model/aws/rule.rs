pub struct Rule {
    // None denotes the default rule, always evaluated last.
    pub priority: Option<u16>,
}
