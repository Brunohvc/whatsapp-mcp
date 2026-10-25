package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// A media server that accepts the upload and never answers must make the request fail, not hang forever
// (on 24–25/10/2026 one dead connection held every group send for 13 hours).
func TestMediaHTTPClientDoesNotHangOnSilentServer(t *testing.T) {
	release := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-release
	}))
	defer srv.Close()
	defer close(release)

	done := make(chan error, 1)
	go func() {
		resp, err := newMediaHTTPClient(500 * time.Millisecond).Post(srv.URL, "application/octet-stream", nil)
		if resp != nil {
			resp.Body.Close()
		}
		done <- err
	}()

	select {
	case err := <-done:
		if err == nil {
			t.Fatal("expected a timeout error from a server that never answers")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("media upload hung on a silent server")
	}
}
