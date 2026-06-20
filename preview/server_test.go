package preview

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	xaligo "github.com/ryo-arima/xaligo"
)

func TestServerRefreshAndDiagnostics(t *testing.T) {
	path := filepath.Join(t.TempDir(), "diagram.xal")
	if err := os.WriteFile(path, []byte(`<frame width="320" height="180"><box title="Preview" /></frame>`), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := New(path, Options{Render: xaligo.RenderOptions{Theme: "dark"}})
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/diagram.svg", nil)
	rec := httptest.NewRecorder()
	server.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `<svg`) || !strings.Contains(rec.Body.String(), `#111827`) {
		t.Fatalf("initial SVG: status=%d body=%s", rec.Code, rec.Body.String())
	}

	if err := os.WriteFile(path, []byte(`<frame>`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := server.Refresh(); err != nil {
		t.Fatal(err)
	}
	rec = httptest.NewRecorder()
	server.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("invalid SVG status = %d, want 422", rec.Code)
	}

	statusReq := httptest.NewRequest(http.MethodGet, "/api/status", nil)
	statusRec := httptest.NewRecorder()
	server.Handler().ServeHTTP(statusRec, statusReq)
	if !strings.Contains(statusRec.Body.String(), `"error"`) || !strings.Contains(statusRec.Body.String(), `"line":1`) {
		t.Fatalf("status has no diagnostic: %s", statusRec.Body.String())
	}
}

func TestServerHealth(t *testing.T) {
	path := filepath.Join(t.TempDir(), "diagram.xal")
	_ = os.WriteFile(path, []byte(`<frame width="10" height="10" />`), 0644)
	server, err := New(path, Options{})
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	server.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if rec.Code != http.StatusOK || rec.Body.String() != "ok\n" {
		t.Fatalf("health response = %d %q", rec.Code, rec.Body.String())
	}
}
