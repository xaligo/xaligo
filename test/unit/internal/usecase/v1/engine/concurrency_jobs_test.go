package engine_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

// Parallel scheduling belongs to callers. This regression test runs independent
// synchronous V1 engine jobs concurrently without adding concurrency control
// to the engine itself.
func TestV1EngineSupportsCallerControlledIndependentJobs(t *testing.T) {
	const jobs = 24
	errors := make(chan error, jobs)
	var workers sync.WaitGroup
	workers.Add(jobs)

	for i := 0; i < jobs; i++ {
		i := i
		go func() {
			defer workers.Done()
			width := 240 + i
			source := []byte(fmt.Sprintf(`<frame id="frame-%d" width="%d" height="120"><blank /></frame>`, i, width))
			document, err := v1engine.ParseV1EngineParseDocument(bytes.NewReader(source))
			if err != nil {
				errors <- fmt.Errorf("job %d parse: %w", i, err)
				return
			}
			box, err := v1engine.BuildV1EngineLayoutBuild(document)
			if err != nil {
				errors <- fmt.Errorf("job %d layout: %w", i, err)
				return
			}
			if box.W != float64(width) {
				errors <- fmt.Errorf("job %d got width %.1f, want %d", i, box.W, width)
			}
		}()
	}

	workers.Wait()
	close(errors)
	for err := range errors {
		t.Error(err)
	}
}
