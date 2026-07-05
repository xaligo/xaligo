package usecase_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestValidateReturnsDiagnosticsError(t *testing.T) {
	err := usecase.Validate(context.Background(), []byte(`<frame><item id="abc" /></frame>`))
	if err == nil {
		t.Fatal("Validate invalid document error = nil")
	}
	var diagnosticsErr *entity.DiagnosticsError
	if !errorsAs(err, &diagnosticsErr) || len(diagnosticsErr.Diagnostics) == 0 {
		t.Fatalf("error = %T %v, want DiagnosticsError", err, err)
	}
	if !strings.Contains(err.Error(), "positive integer") {
		t.Fatalf("error = %v, want positive integer", err)
	}
}

func TestValidateReportsItemAndConnectionBranches(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"comma item", `<frame><item id="1,2" /></frame>`, "single ID"},
		{"spacer item", `<frame><item /></frame>`, ""},
		{"missing src", `<frame><connection dst="2" /></frame>`, "src attribute"},
		{"missing dst", `<frame><connection src="1" /></frame>`, "dst attribute"},
		{"missing group id", `<frame><generic-group title="g" /></frame>`, `<generic-group> requires a non-empty id attribute`},
		{"missing rectangle id", `<frame><rectangle title="r" /></frame>`, `<rectangle> requires a non-empty id attribute`},
		{"missing port id", `<frame><rectangle id="r"><port title="p" /></rectangle></frame>`, `<port> requires a non-empty id attribute`},
		{"duplicate frame id", `<frame><generic-group id="dup" /><rectangle id="dup" /></frame>`, `duplicate frame reference id "dup"`},
		{"unknown src", `<frame><item id="2" /><connection src="one" dst="2" /></frame>`, `src="one"> does not match any <item> or group id/name/ref`},
		{"unknown dst", `<frame><item id="1" /><connection src="1" dst="two" /></frame>`, `dst="two"> does not match any <item> or group id/name/ref`},
		{"missing endpoint item", `<frame><item id="1" /><connection src="1" dst="2" /></frame>`, `dst="2"> does not match any <item> or group id/name/ref`},
		{"ambiguous endpoint item", `<frame><item id="1" /><item id="1" /><item id="2" /><connection src="1" dst="2" /></frame>`, `ambiguous because endpoint id="1" appears 2 times`},
		{"nested connection", `<frame><container><connection src="1" dst="2" /></container><item id="1" /><item id="2" /></frame>`, "direct child"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := usecase.Validate(context.Background(), []byte(tc.input))
			if tc.want == "" {
				if err != nil {
					t.Fatalf("Validate() err = %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("Validate() err = %v, want %q", err, tc.want)
			}
		})
	}
}

func TestDiagnoseReportsParsePositionAndContext(t *testing.T) {
	diagnostics, err := usecase.Diagnose(context.Background(), []byte("<frame>"))
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != 1 || diagnostics[0].Severity != usecase.SeverityError || diagnostics[0].Line == 0 || diagnostics[0].Message == "" {
		t.Fatalf("diagnostics = %#v", diagnostics)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := usecase.Diagnose(ctx, []byte(simpleXAL)); err == nil {
		t.Fatal("Diagnose canceled context error = nil")
	}
}

func errorsAs(err error, target interface{}) bool {
	return errors.As(err, target)
}
