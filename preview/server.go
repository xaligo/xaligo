// Package preview provides the reusable HTTP/SSE live-preview protocol used by
// xaligo serve and future editor integrations.
package preview

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	xaligo "github.com/ryo-arima/xaligo"
)

type Options struct {
	Render       xaligo.RenderOptions
	PollInterval time.Duration
}

type Status struct {
	Version     uint64              `json:"version"`
	Error       string              `json:"error,omitempty"`
	Diagnostics []xaligo.Diagnostic `json:"diagnostics,omitempty"`
}

type Server struct {
	path     string
	opts     Options
	mu       sync.RWMutex
	hash     [sha256.Size]byte
	haveHash bool
	svg      []byte
	status   Status
	nextSub  uint64
	subs     map[uint64]chan uint64
}

func New(path string, opts Options) (*Server, error) {
	if path == "" {
		return nil, fmt.Errorf("preview input path is required")
	}
	if opts.PollInterval <= 0 {
		opts.PollInterval = 500 * time.Millisecond
	}
	opts.Render.Format = xaligo.FormatSVG
	if err := xaligo.ValidateRenderOptions(opts.Render); err != nil {
		return nil, err
	}
	s := &Server{path: path, opts: opts, subs: map[uint64]chan uint64{}}
	if err := s.refresh(true); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/diagram.svg", s.handleSVG)
	mux.HandleFunc("/api/status", s.handleStatus)
	mux.HandleFunc("/events", s.handleEvents)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ok\n"))
	})
	return mux
}

func (s *Server) Run(ctx context.Context, address string) error {
	if address == "" {
		address = "127.0.0.1:8080"
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	httpServer := &http.Server{Addr: address, Handler: s.Handler(), ReadHeaderTimeout: 5 * time.Second}
	watchDone := make(chan struct{})
	go func() {
		defer close(watchDone)
		s.watch(ctx)
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

func (s *Server) Refresh() error { return s.refresh(false) }

func (s *Server) refresh(force bool) error {
	source, err := os.ReadFile(s.path)
	if err != nil {
		return fmt.Errorf("read preview input: %w", err)
	}
	hash := sha256.Sum256(source)
	s.mu.RLock()
	unchanged := s.haveHash && hash == s.hash
	s.mu.RUnlock()
	if unchanged && !force {
		return nil
	}

	svg, renderErr := xaligo.RenderSVG(context.Background(), source, s.opts.Render)
	s.mu.Lock()
	s.hash = hash
	s.haveHash = true
	s.status.Version++
	if renderErr != nil {
		s.status.Error = renderErr.Error()
		diagnostics, diagnoseErr := xaligo.Diagnose(context.Background(), source)
		if diagnoseErr == nil && len(diagnostics) > 0 {
			s.status.Diagnostics = diagnostics
		} else {
			s.status.Diagnostics = []xaligo.Diagnostic{{Severity: xaligo.SeverityError, Message: renderErr.Error()}}
		}
		s.svg = nil
	} else {
		s.status.Error = ""
		s.status.Diagnostics = nil
		s.svg = append(s.svg[:0], svg...)
	}
	version := s.status.Version
	subs := make([]chan uint64, 0, len(s.subs))
	for _, ch := range s.subs {
		subs = append(subs, ch)
	}
	s.mu.Unlock()
	for _, ch := range subs {
		select {
		case ch <- version:
		default:
		}
	}
	return nil
}

func (s *Server) watch(ctx context.Context) {
	ticker := time.NewTicker(s.opts.PollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_ = s.refresh(false)
		}
	}
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(indexHTML))
}

func (s *Server) handleSVG(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	svg := append([]byte(nil), s.svg...)
	errText := s.status.Error
	s.mu.RUnlock()
	if errText != "" {
		http.Error(w, errText, http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-store")
	_, _ = w.Write(svg)
}

func (s *Server) handleStatus(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	status := s.status
	s.mu.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	_ = json.NewEncoder(w).Encode(status)
}

func (s *Server) handleEvents(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	s.mu.Lock()
	s.nextSub++
	id := s.nextSub
	ch := make(chan uint64, 1)
	s.subs[id] = ch
	version := s.status.Version
	s.mu.Unlock()
	defer func() {
		s.mu.Lock()
		delete(s.subs, id)
		s.mu.Unlock()
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
