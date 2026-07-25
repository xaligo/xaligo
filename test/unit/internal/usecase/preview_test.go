package usecase_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
)

func TestPreviewRepositoryHandlers(t *testing.T) {
	path := filepath.Join(t.TempDir(), "preview.xal")
	if err := os.WriteFile(path, []byte(simpleXAL), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := newUsecase().NewPreviewRepository(path, entity.PreviewOptions{Render: entity.RenderOptions{Theme: "light"}})
	if err != nil {
		t.Fatal(err)
	}
	handler := server.Handler()

	index := httptest.NewRecorder()
	handler.ServeHTTP(index, httptest.NewRequest(http.MethodGet, "/", nil))
	if index.Code != http.StatusOK || !strings.Contains(index.Body.String(), "xaligo live preview") {
		t.Fatalf("index status=%d body=%q", index.Code, index.Body.String())
	}

	notFound := httptest.NewRecorder()
	handler.ServeHTTP(notFound, httptest.NewRequest(http.MethodGet, "/missing", nil))
	if notFound.Code != http.StatusNotFound {
		t.Fatalf("missing status = %d", notFound.Code)
	}

	svg := httptest.NewRecorder()
	handler.ServeHTTP(svg, httptest.NewRequest(http.MethodGet, "/diagram.svg", nil))
	if svg.Code != http.StatusOK || !strings.Contains(svg.Body.String(), "<svg") || svg.Header().Get("Content-Type") != "image/svg+xml" {
		t.Fatalf("svg status=%d header=%q body=%q", svg.Code, svg.Header().Get("Content-Type"), svg.Body.String())
	}

	statusResponse := httptest.NewRecorder()
	handler.ServeHTTP(statusResponse, httptest.NewRequest(http.MethodGet, "/api/status", nil))
	if statusResponse.Code != http.StatusOK {
		t.Fatalf("status code = %d", statusResponse.Code)
	}
	var status entity.PreviewStatus
	if err := json.Unmarshal(statusResponse.Body.Bytes(), &status); err != nil {
		t.Fatal(err)
	}
	if status.Version == 0 || status.Error != "" {
		t.Fatalf("status = %#v", status)
	}

	health := httptest.NewRecorder()
	handler.ServeHTTP(health, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if health.Code != http.StatusOK || health.Body.String() != "ok\n" {
		t.Fatalf("health status=%d body=%q", health.Code, health.Body.String())
	}

	eventCtx, cancelEvent := context.WithCancel(context.Background())
	cancelEvent()
	events := httptest.NewRecorder()
	handler.ServeHTTP(events, httptest.NewRequest(http.MethodGet, "/events", nil).WithContext(eventCtx))
	if events.Code != http.StatusOK || !strings.Contains(events.Body.String(), "event: update") {
		t.Fatalf("events status=%d body=%q", events.Code, events.Body.String())
	}
}

func TestPreviewRepositorySVGHandlerReportsRenderError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "broken.xal")
	if err := os.WriteFile(path, []byte(`<frame><item id="bad" /></frame>`), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := newUsecase().NewPreviewRepository(path, entity.PreviewOptions{Render: entity.RenderOptions{Theme: "light"}})
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/diagram.svg", nil))
	if response.Code != http.StatusUnprocessableEntity || !strings.Contains(response.Body.String(), "positive integer") {
		t.Fatalf("response status=%d body=%q", response.Code, response.Body.String())
	}
}

func TestPreviewRepositoryRunStopsWhenContextCanceled(t *testing.T) {
	path := filepath.Join(t.TempDir(), "preview.xal")
	if err := os.WriteFile(path, []byte(simpleXAL), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := newUsecase().NewPreviewRepository(path, entity.PreviewOptions{Render: entity.RenderOptions{Theme: "light"}})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := server.Run(ctx, "127.0.0.1:0"); err != nil {
		t.Fatal(err)
	}
}

const simpleMarkdownXAL = "# Guide\n\nIntro text.\n\n```xal\n" + simpleXAL + "\n```\n\nOutro text.\n"

func TestPreviewRepositoryHandlersMarkdownKind(t *testing.T) {
	path := filepath.Join(t.TempDir(), "guide.md")
	if err := os.WriteFile(path, []byte(simpleMarkdownXAL), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := newUsecase().NewPreviewRepository(path, entity.PreviewOptions{
		Kind:   entity.PreviewKindHTML,
		Render: entity.RenderOptions{Theme: "light"},
	})
	if err != nil {
		t.Fatal(err)
	}
	handler := server.Handler()

	index := httptest.NewRecorder()
	handler.ServeHTTP(index, httptest.NewRequest(http.MethodGet, "/", nil))
	if index.Code != http.StatusOK || !strings.Contains(index.Body.String(), "xaligo live preview") || !strings.Contains(index.Body.String(), "<iframe") {
		t.Fatalf("index status=%d body=%q", index.Code, index.Body.String())
	}

	content := httptest.NewRecorder()
	handler.ServeHTTP(content, httptest.NewRequest(http.MethodGet, "/content.html", nil))
	if content.Code != http.StatusOK || content.Header().Get("Content-Type") != "text/html; charset=utf-8" {
		t.Fatalf("content status=%d header=%q", content.Code, content.Header().Get("Content-Type"))
	}
	body := content.Body.String()
	if !strings.Contains(body, "<svg") || !strings.Contains(body, "<h1>Guide</h1>") || !strings.Contains(body, "Outro text.") {
		t.Fatalf("content body = %q", body)
	}
}

func TestPreviewRepositoryMarkdownHandlerReportsRenderError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "broken.md")
	broken := "```xal\n<frame><item id=\"bad\" /></frame>\n```\n"
	if err := os.WriteFile(path, []byte(broken), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := newUsecase().NewPreviewRepository(path, entity.PreviewOptions{
		Kind:   entity.PreviewKindHTML,
		Render: entity.RenderOptions{Theme: "light"},
	})
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/content.html", nil))
	if response.Code != http.StatusUnprocessableEntity {
		t.Fatalf("response status=%d body=%q", response.Code, response.Body.String())
	}
}
