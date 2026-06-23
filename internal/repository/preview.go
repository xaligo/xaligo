// Package preview provides the reusable HTTP/SSE live-preview protocol used by
// xaligo serve and future editor integrations.
package repository

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ryo-arima/xaligo/internal/entity"
)

type PreviewServer interface {
	Handler() http.Handler
	Run(context.Context, string) error
	Refresh() error
}

type previewServer struct {
	path     string
	opts     entity.PreviewOptions
	mu       sync.RWMutex
	hash     [sha256.Size]byte
	haveHash bool
	svg      []byte
	status   entity.PreviewStatus
	nextSub  uint64
	subs     map[uint64]chan uint64
	render   func(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	validate func(entity.RenderOptions) error
	diagnose func(context.Context, []byte) ([]entity.Diagnostic, error)
	read     func(string) ([]byte, error)
}

func NewPreviewServer(
	path string,
	opts entity.PreviewOptions,
	render func(context.Context, []byte, entity.RenderOptions) ([]byte, error),
	validate func(entity.RenderOptions) error,
	diagnose func(context.Context, []byte) ([]entity.Diagnostic, error),
	read func(string) ([]byte, error),
) (PreviewServer, error) {
	if path == "" {
		return nil, fmt.Errorf("preview input path is required")
	}
	if opts.PollInterval <= 0 {
		opts.PollInterval = 500 * time.Millisecond
	}
	opts.Render.Format = entity.Format("svg")
	if err := validate(opts.Render); err != nil {
		return nil, err
	}
	s := &previewServer{
		path: path, opts: opts, subs: map[uint64]chan uint64{},
		render: render, validate: validate, diagnose: diagnose, read: read,
	}
	if err := s.refresh(true); err != nil {
		return nil, err
	}
	return s, nil
}

func (rcvr *previewServer) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rcvr.handleIndex)
	mux.HandleFunc("/diagram.svg", rcvr.handleSVG)
	mux.HandleFunc("/api/status", rcvr.handleStatus)
	mux.HandleFunc("/events", rcvr.handleEvents)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ok\n"))
	})
	return mux
}

func (rcvr *previewServer) Run(ctx context.Context, address string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Err() != nil {
		return nil
	}
	if address == "" {
		address = "127.0.0.1:8080"
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	httpServer := &http.Server{Addr: address, Handler: rcvr.Handler(), ReadHeaderTimeout: 5 * time.Second}
	watchDone := make(chan struct{})
	go func() {
		defer close(watchDone)
		rcvr.watch(ctx)
	}()
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = httpServer.Shutdown(shutdownCtx)
	}()
	err := httpServer.ListenAndServe()
	cancel()
	if err == http.ErrServerClosed {
		err = nil
	}
	<-watchDone
	return err
}

func (rcvr *previewServer) Refresh() error { return rcvr.refresh(false) }

func (rcvr *previewServer) refresh(force bool) error {
	source, err := rcvr.read(rcvr.path)
	if err != nil {
		return fmt.Errorf("read preview input: %w", err)
	}
	hash := sha256.Sum256(source)
	rcvr.mu.RLock()
	unchanged := rcvr.haveHash && hash == rcvr.hash
	rcvr.mu.RUnlock()
	if unchanged && !force {
		return nil
	}

	svg, renderErr := rcvr.render(context.Background(), source, rcvr.opts.Render)
	rcvr.mu.Lock()
	rcvr.hash = hash
	rcvr.haveHash = true
	rcvr.status.Version++
	if renderErr != nil {
		rcvr.status.Error = renderErr.Error()
		diagnostics, diagnoseErr := rcvr.diagnose(context.Background(), source)
		if diagnoseErr == nil && len(diagnostics) > 0 {
			rcvr.status.Diagnostics = diagnostics
		} else {
			rcvr.status.Diagnostics = []entity.Diagnostic{{Severity: entity.DiagnosticSeverity("error"), Message: renderErr.Error()}}
		}
		rcvr.svg = nil
	} else {
		rcvr.status.Error = ""
		rcvr.status.Diagnostics = nil
		rcvr.svg = append(rcvr.svg[:0], svg...)
	}
	version := rcvr.status.Version
	subs := make([]chan uint64, 0, len(rcvr.subs))
	for _, ch := range rcvr.subs {
		subs = append(subs, ch)
	}
	rcvr.mu.Unlock()
	for _, ch := range subs {
		select {
		case ch <- version:
		default:
		}
	}
	return nil
}

func (rcvr *previewServer) watch(ctx context.Context) {
	ticker := time.NewTicker(rcvr.opts.PollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_ = rcvr.refresh(false)
		}
	}
}

func (rcvr *previewServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(indexHTML))
}

func (rcvr *previewServer) handleSVG(w http.ResponseWriter, _ *http.Request) {
	rcvr.mu.RLock()
	svg := append([]byte(nil), rcvr.svg...)
	errText := rcvr.status.Error
	rcvr.mu.RUnlock()
	if errText != "" {
		http.Error(w, errText, http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-store")
	_, _ = w.Write(svg)
}

func (rcvr *previewServer) handleStatus(w http.ResponseWriter, _ *http.Request) {
	rcvr.mu.RLock()
	status := rcvr.status
	rcvr.mu.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	_ = json.NewEncoder(w).Encode(status)
}

func (rcvr *previewServer) handleEvents(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	rcvr.mu.Lock()
	rcvr.nextSub++
	id := rcvr.nextSub
	ch := make(chan uint64, 1)
	rcvr.subs[id] = ch
	version := rcvr.status.Version
	rcvr.mu.Unlock()
	defer func() {
		rcvr.mu.Lock()
		delete(rcvr.subs, id)
		rcvr.mu.Unlock()
	}()

	_, _ = fmt.Fprintf(w, "event: update\ndata: %d\n\n", version)
	flusher.Flush()
	for {
		select {
		case <-r.Context().Done():
			return
		case version := <-ch:
			_, _ = fmt.Fprintf(w, "event: update\ndata: %d\n\n", version)
			flusher.Flush()
		}
	}
}

const indexHTML = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>xaligo preview</title><style>
:root{color-scheme:light dark;font-family:system-ui,sans-serif}body{margin:0;background:#111827}header{height:42px;display:flex;align-items:center;padding:0 14px;color:#e5e7eb;background:#0f172a}main{height:calc(100vh - 42px);display:grid;place-items:center;overflow:auto}img{max-width:calc(100% - 32px);max-height:calc(100% - 32px);background:white;box-shadow:0 8px 32px #0008}.error{white-space:pre-wrap;color:#fecaca;background:#7f1d1d;padding:16px;border-radius:6px;max-width:80%}[hidden]{display:none}
</style></head><body><header>xaligo live preview</header><main><img id="diagram" alt="xaligo diagram"><pre id="error" class="error" hidden></pre></main>
<script>
const image=document.querySelector('#diagram'), error=document.querySelector('#error');
async function reload(v){const status=await fetch('/api/status?'+v,{cache:'no-store'}).then(r=>r.json());if(status.error){image.hidden=true;error.hidden=false;const d=status.diagnostics?.[0];error.textContent=d?.line?'Line '+d.line+', column '+d.column+': '+d.message:status.error}else{error.hidden=true;image.hidden=false;image.src='/diagram.svg?v='+status.version}}
new EventSource('/events').addEventListener('update',e=>reload(e.data));
</script></body></html>`
