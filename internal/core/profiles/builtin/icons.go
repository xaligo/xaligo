// Package builtin contains the domain-neutral V2 profile data.
package builtin

import "github.com/xaligo/xaligo/internal/entity"

const iconEnvelopePrefix = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#334155" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">`

type iconDefinition struct {
	name        string
	description string
	body        string
	tags        []string
	aliases     []string
}

var iconDefinitions = []iconDefinition{
	{
		name:        "generic",
		description: "Generic diagram element",
		body:        `<path d="M12 3 21 12 12 21 3 12Z"/><circle cx="12" cy="12" r="2.5"/>`,
		tags:        []string{"generic", "element", "node"},
		aliases:     []string{"node"},
	},
	{
		name:        "user",
		description: "Generic user or actor",
		body:        `<circle cx="12" cy="8" r="3.5"/><path d="M5 21c.7-4.2 3.2-6.5 7-6.5s6.3 2.3 7 6.5"/>`,
		tags:        []string{"user", "person", "actor"},
		aliases:     []string{"person"},
	},
	{
		name:        "service",
		description: "Generic service or component",
		body:        `<path d="m12 3 7.8 4.5v9L12 21l-7.8-4.5v-9Z"/><circle cx="12" cy="12" r="3"/>`,
		tags:        []string{"service", "component", "system"},
		aliases:     []string{"component"},
	},
	{
		name:        "application",
		description: "Generic software application",
		body:        `<rect x="3" y="4" width="18" height="16" rx="2"/><path d="M3 9h18"/><circle cx="6.5" cy="6.5" r=".5" fill="#334155" stroke="none"/><circle cx="9" cy="6.5" r=".5" fill="#334155" stroke="none"/>`,
		tags:        []string{"application", "software", "app"},
		aliases:     []string{"app"},
	},
	{
		name:        "server",
		description: "Generic compute server",
		body:        `<rect x="4" y="3" width="16" height="8" rx="1.5"/><rect x="4" y="13" width="16" height="8" rx="1.5"/><circle cx="7" cy="7" r=".7" fill="#334155" stroke="none"/><circle cx="7" cy="17" r=".7" fill="#334155" stroke="none"/><path d="M10 7h7M10 17h7"/>`,
		tags:        []string{"server", "compute", "host"},
		aliases:     []string{"compute"},
	},
	{
		name:        "database",
		description: "Generic database",
		body:        `<ellipse cx="12" cy="5.5" rx="7.5" ry="3"/><path d="M4.5 5.5v6c0 1.7 3.4 3 7.5 3s7.5-1.3 7.5-3v-6M4.5 11.5v6c0 1.7 3.4 3 7.5 3s7.5-1.3 7.5-3v-6"/>`,
		tags:        []string{"database", "data", "storage"},
		aliases:     []string{"db"},
	},
	{
		name:        "storage",
		description: "Generic data storage",
		body:        `<path d="M4 7h16v12H4Z"/><path d="m4 7 2-3h12l2 3M8 11h8M8 15h5"/>`,
		tags:        []string{"storage", "disk", "archive"},
		aliases:     []string{"disk"},
	},
	{
		name:        "network",
		description: "Generic network topology",
		body:        `<path d="M12 7v4M7 17l5-6 5 6"/><circle cx="12" cy="5" r="2.5"/><circle cx="6" cy="19" r="2.5"/><circle cx="18" cy="19" r="2.5"/>`,
		tags:        []string{"network", "topology", "connection"},
		aliases:     []string{"topology"},
	},
	{
		name:        "cloud",
		description: "Generic cloud environment",
		body:        `<path d="M7 19h10a4 4 0 0 0 .6-8 6 6 0 0 0-11.4-1.8A5 5 0 0 0 7 19Z"/>`,
		tags:        []string{"cloud", "environment", "infrastructure"},
	},
	{
		name:        "queue",
		description: "Generic message queue",
		body:        `<rect x="3" y="5" width="5" height="4" rx="1"/><rect x="3" y="15" width="5" height="4" rx="1"/><path d="M11 7h10M11 17h10M17 4l4 3-4 3M17 14l4 3-4 3"/>`,
		tags:        []string{"queue", "message", "asynchronous"},
		aliases:     []string{"messaging"},
	},
	{
		name:        "gateway",
		description: "Generic gateway or ingress",
		body:        `<path d="M7 4v16M17 4v16M3 8h10M11 5l3 3-3 3M21 16H11M13 13l-3 3 3 3"/>`,
		tags:        []string{"gateway", "ingress", "egress"},
		aliases:     []string{"ingress"},
	},
	{
		name:        "document",
		description: "Generic document or file",
		body:        `<path d="M6 3h8l4 4v14H6Z"/><path d="M14 3v5h4M9 12h6M9 16h6"/>`,
		tags:        []string{"document", "file", "content"},
		aliases:     []string{"file"},
	},
	{
		name:        "terminal",
		description: "Generic command-line terminal",
		body:        `<rect x="3" y="4" width="18" height="16" rx="2"/><path d="m7 9 3 3-3 3M12 16h5"/>`,
		tags:        []string{"terminal", "command", "cli"},
		aliases:     []string{"cli"},
	},
}

// IconRegistrations returns independent registration values for the builtin
// icon catalog. Every SVG still passes through the Rust normalization path
// before persistence.
func IconRegistrations() []entity.IconRegistration {
	registrations := make([]entity.IconRegistration, 0, len(iconDefinitions))
	for _, definition := range iconDefinitions {
		registrations = append(registrations, entity.IconRegistration{
			Reference:   "builtin:" + definition.name,
			SVG:         []byte(iconEnvelopePrefix + definition.body + `</svg>`),
			Description: definition.description,
			Tags:        append([]string(nil), definition.tags...),
			Aliases:     append([]string(nil), definition.aliases...),
			License:     "MIT",
			Source:      "xaligo builtin profile",
		})
	}
	return registrations
}
