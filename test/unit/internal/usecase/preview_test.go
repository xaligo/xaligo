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
	if index.Code != http.StatusOK || !strings.Contains(index.Body.String(), "xaligo live preview") || !strings.Contains(index.Body.String(), `<iframe sandbox`) {
		t.Fatalf("index status=%d body=%q", index.Code, index.Body.String())
	}
	if !strings.Contains(index.Header().Get("Content-Security-Policy"), "frame-src 'self'") {
		t.Fatalf("index CSP = %q", index.Header().Get("Content-Security-Policy"))
	}

	content := httptest.NewRecorder()
	handler.ServeHTTP(content, httptest.NewRequest(http.MethodGet, "/content.html", nil))
	if content.Code != http.StatusOK || content.Header().Get("Content-Type") != "text/html; charset=utf-8" {
		t.Fatalf("content status=%d header=%q", content.Code, content.Header().Get("Content-Type"))
	}
	body := content.Body.String()
	if !strings.Contains(body, `class="xaligo-diagram"`) || strings.Contains(body, "<svg") || !strings.Contains(body, "<h1>Guide</h1>") || !strings.Contains(body, "Outro text.") {
		t.Fatalf("content body = %q", body)
	}
	if !strings.Contains(body, `.xaligo-diagram img{max-width:100%;height:auto;border:0;box-shadow:none;display:block}`) {
		t.Fatalf("rendered Markdown diagrams must not have a border or shadow: %q", body)
	}
	if !strings.Contains(content.Header().Get("Content-Security-Policy"), "sandbox") {
		t.Fatalf("content CSP = %q", content.Header().Get("Content-Security-Policy"))
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

	statusResponse := httptest.NewRecorder()
	server.Handler().ServeHTTP(statusResponse, httptest.NewRequest(http.MethodGet, "/api/status", nil))
	var status entity.PreviewStatus
	if err := json.Unmarshal(statusResponse.Body.Bytes(), &status); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(status.Error, "render xal code block 1") {
		t.Fatalf("status error = %q", status.Error)
	}
	if len(status.Diagnostics) != 1 || status.Diagnostics[0].Message != status.Error {
		t.Fatalf("Markdown diagnostics must report the block render error, got %#v", status.Diagnostics)
	}
}

func TestPreviewRepositoryMarkdownDisablesRawHTMLAndIsolatesSVGDocuments(t *testing.T) {
	path := filepath.Join(t.TempDir(), "unsafe.md")
	source := "# Safe\n\n<script>globalThis.markdownXSS = true</script>\n\n" +
		`<img src="missing" onerror="globalThis.markdownXSS = true">` + "\n\n" +
		"```xal\n" + simpleXAL + "\n```\n\n" +
		"```xal\n" + simpleXAL + "\n```\n"
	if err := os.WriteFile(path, []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := newUsecase().NewPreviewRepository(path, entity.PreviewOptions{
		Kind:   entity.PreviewKindHTML,
		Render: entity.RenderOptions{Theme: "light"},
	})
	if err != nil {
		t.Fatal(err)
	}

	content := httptest.NewRecorder()
	server.Handler().ServeHTTP(content, httptest.NewRequest(http.MethodGet, "/content.html", nil))
	body := content.Body.String()
	if content.Code != http.StatusOK {
		t.Fatalf("content status=%d body=%q", content.Code, body)
	}
	if strings.Contains(body, "<script") || strings.Contains(body, "onerror") || strings.Contains(body, "markdownXSS") {
		t.Fatalf("raw HTML reached preview output: %q", body)
	}
	if count := strings.Count(body, "data:image/svg+xml;base64,"); count != 2 {
		t.Fatalf("isolated SVG image count = %d, body=%q", count, body)
	}
	if strings.Contains(body, `id="xaligo-slide-clip"`) || strings.Contains(body, `<clipPath`) {
		t.Fatalf("SVG definitions must not share the Markdown document: %q", body)
	}
}

func TestPreviewRepositoryMarkdownServesOnlyImagesInsideSourceDirectory(t *testing.T) {
	root := t.TempDir()
	sourceDir := filepath.Join(root, "docs")
	imageDir := filepath.Join(sourceDir, "images")
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		t.Fatal(err)
	}
	imageData := []byte("\x89PNG\r\n\x1a\nlocal-image")
	imagePath := filepath.Join(imageDir, "pixel.png")
	if err := os.WriteFile(imagePath, imageData, 0644); err != nil {
		t.Fatal(err)
	}
	secretData := []byte("\x89PNG\r\n\x1a\noutside-image")
	secretPath := filepath.Join(root, "secret.png")
	if err := os.WriteFile(secretPath, secretData, 0644); err != nil {
		t.Fatal(err)
	}
	sourcePath := filepath.Join(sourceDir, "guide.md")
	source := "![local](images/pixel.png)\n\n![outside](../secret.png)\n"
	if err := os.WriteFile(sourcePath, []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := newUsecase().NewPreviewRepository(sourcePath, entity.PreviewOptions{
		Kind:   entity.PreviewKindHTML,
		Render: entity.RenderOptions{Theme: "light"},
	})
	if err != nil {
		t.Fatal(err)
	}
	handler := server.Handler()

	content := httptest.NewRecorder()
	handler.ServeHTTP(content, httptest.NewRequest(http.MethodGet, "/content.html", nil))
	body := content.Body.String()
	if !strings.Contains(body, `src="/assets/images/pixel.png"`) {
		t.Fatalf("relative image was not routed through the asset handler: %q", body)
	}
	if strings.Contains(body, "secret.png") {
		t.Fatalf("escaping image path reached preview output: %q", body)
	}

	image := httptest.NewRecorder()
	handler.ServeHTTP(image, httptest.NewRequest(http.MethodGet, "/assets/images/pixel.png", nil))
	if image.Code != http.StatusOK || image.Body.String() != string(imageData) || image.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Fatalf("asset status=%d headers=%v body=%q", image.Code, image.Header(), image.Body.String())
	}

	traversal := httptest.NewRecorder()
	handler.ServeHTTP(traversal, httptest.NewRequest(http.MethodGet, "http://preview.test/assets/%2e%2e/secret.png", nil))
	if traversal.Code == http.StatusOK || traversal.Body.String() == string(secretData) {
		t.Fatalf("path traversal exposed outside file: status=%d body=%q", traversal.Code, traversal.Body.String())
	}

	symlinkPath := filepath.Join(imageDir, "escape.png")
	if err := os.Symlink(secretPath, symlinkPath); err == nil {
		symlink := httptest.NewRecorder()
		handler.ServeHTTP(symlink, httptest.NewRequest(http.MethodGet, "/assets/images/escape.png", nil))
		if symlink.Code == http.StatusOK || symlink.Body.String() == string(secretData) {
			t.Fatalf("symlink exposed outside file: status=%d body=%q", symlink.Code, symlink.Body.String())
		}
	}
}
