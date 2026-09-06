package entity

// EngineAWSComponentSpec is an explicit, typed native component request.
// Source XML is never passed to the Rust component implementation.
type EngineAWSComponentSpec struct {
	Kind, Domain                         string
	Protocol, MutualTLS                  string
	Port                                 *uint16
	Certificate, TrustStore, TargetGroup string
	BackendTLS, BackendMTLS              *bool
	ShowTitle                            *bool
	DetailLevel, Show, Hide              string
	Type, Name, Value, Aux, Order        string
}
