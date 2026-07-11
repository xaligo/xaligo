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
		{"invalid item dx", `<frame><item id="1" dx="east" /></frame>`, `dx="east"> must be a number`},
		{"spacer item", `<frame><item /></frame>`, ""},
		{"missing src", `<frame><connection dst="2" /></frame>`, "src attribute"},
		{"missing dst", `<frame><connection src="1" /></frame>`, "dst attribute"},
		{"missing group id", `<frame><generic-group title="g" /></frame>`, `<generic-group> requires a non-empty id attribute`},
		{"missing rectangle id", `<frame><rectangle title="r" /></frame>`, `<rectangle> requires a non-empty id attribute`},
		{"missing port id", `<frame><rectangle id="r"><port title="p" /></rectangle></frame>`, `<port> requires a non-empty id attribute`},
		{"duplicate frame id", `<frame><generic-group id="dup" /><rectangle id="dup" /></frame>`, `duplicate frame reference id "dup"`},
		{"unknown src", `<frame><item id="2" /><connection src="one" dst="2" /></frame>`, `src="one"> does not match any connection endpoint id/name/ref`},
		{"unknown dst", `<frame><item id="1" /><connection src="1" dst="two" /></frame>`, `dst="two"> does not match any connection endpoint id/name/ref`},
		{"missing endpoint item", `<frame><item id="1" /><connection src="1" dst="2" /></frame>`, `dst="2"> does not match any connection endpoint id/name/ref`},
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

func TestValidateRejectsInvalidLayoutNumbers(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"NaN width", `<frame width="NaN" />`, "must be a finite number"},
		{"empty width", `<frame width="" />`, "width must not be empty"},
		{"infinite height", `<frame height="+Inf" />`, "must be a finite number"},
		{"zero width", `<frame width="0" />`, "must be greater than zero"},
		{"negative gap", `<frame gap="-1"><blank /></frame>`, "must be zero or greater"},
		{"negative margin", `<frame><blank margin-left="-1" /></frame>`, "must be zero or greater"},
		{"zero row ratio", `<frame><blank row="0" /></frame>`, "must be greater than zero"},
		{"zero col ratio", `<frame layout="horizontal"><blank col="0" /></frame>`, "must be greater than zero"},
		{"zero content width", `<frame content-width="0"><blank /></frame>`, "must be greater than zero"},
		{"span above grid", `<frame><row><col span="13" /></row></frame>`, "exceeds the 12-column grid"},
		{"span total above grid", `<frame><row><col span="7" /><col span="6" /></row></frame>`, "exceeds the 12-column grid"},
		{"ratio total overflow", `<frame><blank row="1e308" /><blank row="1e308" /></frame>`, "ratio total must be finite"},
		{"negative spacing class", `<frame class="ma--1"><blank /></frame>`, "non-negative integer"},
		{"non-finite item offset", `<frame><item dx="NaN" /></frame>`, "must be a finite number"},
		{"non-finite connection scale", `<frame><item id="1" /><item id="2" /><connection src="1" dst="2" scale="NaN" /></frame>`, "must be a finite number"},
		{"zero connection grid", `<frame><item id="1" /><item id="2" /><connection src="1" dst="2" grid="0" /></frame>`, "must be greater than zero"},
		{"negative stroke width", `<frame><item id="1" /><item id="2" /><connection src="1" dst="2" stroke-width="-1" /></frame>`, "must be greater than zero"},
		{"non-finite bend", `<frame><item id="1" /><item id="2" /><connection src="1" dst="2"><bend x="NaN" y="10" /></connection></frame>`, "must be a finite number"},
		{"incomplete bend", `<frame><item id="1" /><item id="2" /><connection src="1" dst="2"><bend x="10" /></connection></frame>`, "require both x and y"},
		{"malformed inline bends", `<frame><item id="1" /><item id="2" /><connection src="1" dst="2" bends="10,20 nope" /></frame>`, "must be an x,y pair"},
		{"unknown overflow", `<frame overflow="hidden"><blank /></frame>`, "must be error or visible"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := usecase.Validate(context.Background(), []byte(tc.input))
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("Validate() error = %v, want %q", err, tc.want)
			}
		})
	}
}

func TestValidateAndRenderShareGeometryInvariant(t *testing.T) {
	input := []byte(`<frame width="145" height="100"><rectangle id="wide" width="500" height="50" /></frame>`)
	validationErr := usecase.Validate(context.Background(), input)
	if validationErr == nil || !strings.Contains(validationErr.Error(), "overflows parent <frame> content box") {
		t.Fatalf("Validate() error = %v", validationErr)
	}
	var diagnosticsErr *entity.DiagnosticsError
	if !errors.As(validationErr, &diagnosticsErr) || len(diagnosticsErr.Diagnostics) != 1 || diagnosticsErr.Diagnostics[0].Line == 0 {
		t.Fatalf("Validate() error lacks positioned diagnostic: %#v", validationErr)
	}

	_, renderErr := newUsecase().RenderSVG(context.Background(), input, entity.RenderOptions{})
	if renderErr == nil || !strings.Contains(renderErr.Error(), "overflows parent <frame> content box") {
		t.Fatalf("RenderSVG() error = %v", renderErr)
	}
}

func TestValidateAndRenderBothRejectZeroRowBeforeSerialization(t *testing.T) {
	input := []byte(`<frame width="200" height="100"><blank row="0" /><blank row="0" /></frame>`)
	validationErr := usecase.Validate(context.Background(), input)
	if validationErr == nil || !strings.Contains(validationErr.Error(), "row=\"0\" must be greater than zero") {
		t.Fatalf("Validate() error = %v", validationErr)
	}
	_, renderErr := newUsecase().RenderSVG(context.Background(), input, entity.RenderOptions{})
	if renderErr == nil || !strings.Contains(renderErr.Error(), "row=\"0\" must be greater than zero") {
		t.Fatalf("RenderSVG() error = %v", renderErr)
	}
}

func TestV1RouteHeadlessValidationIsSharedByPublicRenderFormats(t *testing.T) {
	input := []byte(`<frame width="320" height="180" layout="horizontal">
  <rectangle id="source" title="Source" width="120" height="80" />
  <rectangle id="target" title="Target" width="120" height="80" />
  <connections kind="route" end-arrowhead="triangle">
    <connection src="source" dst="target" />
  </connections>
</frame>`)
	want := `kind="route"> must be headless`

	_, parseErr := usecase.Parse(strings.NewReader(string(input)))
	var positioned *entity.ParseError
	if !errors.As(parseErr, &positioned) || positioned.Position.Line != 5 || !strings.Contains(parseErr.Error(), want) {
		t.Fatalf("Parse() error = %#v, want positioned route-headless error on line 5", parseErr)
	}

	if err := usecase.Validate(context.Background(), input); err == nil || !strings.Contains(err.Error(), want) {
		t.Fatalf("Validate() error = %v, want %q", err, want)
	}

	renderUsecase := newUsecaseWithPPTX(&fakePPTXExporter{})
	renders := []struct {
		name   string
		format entity.Format
		call   func(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	}{
		{name: "Excalidraw", format: usecase.FormatExcalidraw, call: renderUsecase.RenderExcalidraw},
		{name: "SVG", format: usecase.FormatSVG, call: renderUsecase.RenderSVG},
		{name: "PPTX", format: usecase.FormatPPTX, call: renderUsecase.RenderPPTX},
		{name: "XYFlow", format: usecase.FormatXYFlow, call: renderUsecase.RenderXYFlow},
		{name: "Isoflow", format: usecase.FormatIsoflow, call: renderUsecase.RenderIsoflow},
	}
	for _, renderer := range renders {
		t.Run(renderer.name, func(t *testing.T) {
			_, err := renderer.call(context.Background(), input, entity.RenderOptions{Format: renderer.format, Theme: "light"})
			if err == nil || !strings.Contains(err.Error(), want) {
				t.Fatalf("render error = %v, want %q", err, want)
			}
		})
	}
}

func errorsAs(err error, target interface{}) bool {
	return errors.As(err, target)
}
