package builtin_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/core/profiles/builtin"
)

func TestIconRegistrationsAreDomainNeutralAndIndependent(t *testing.T) {
	expected := []string{
		"builtin:application",
		"builtin:cloud",
		"builtin:database",
		"builtin:document",
		"builtin:gateway",
		"builtin:generic",
		"builtin:network",
		"builtin:queue",
		"builtin:server",
		"builtin:service",
		"builtin:storage",
		"builtin:terminal",
		"builtin:user",
	}
	registrations := builtin.IconRegistrations()
	references := make([]string, 0, len(registrations))
	for _, registration := range registrations {
		references = append(references, registration.Reference)
		if registration.License != "MIT" || registration.Source != "xaligo builtin profile" {
			t.Fatalf("attribution for %s = %q, %q", registration.Reference, registration.License, registration.Source)
		}
		if !strings.HasPrefix(string(registration.SVG), `<svg`) {
			t.Fatalf("%s does not contain SVG data", registration.Reference)
		}
	}
	slices.Sort(references)
	if !slices.Equal(references, expected) {
		t.Fatalf("builtin references = %#v", references)
	}

	registrations[0].SVG[0] = 'X'
	registrations[0].Tags[0] = "changed"
	fresh := builtin.IconRegistrations()
	if fresh[0].SVG[0] != '<' || fresh[0].Tags[0] == "changed" {
		t.Fatal("IconRegistrations returned shared mutable data")
	}
}
