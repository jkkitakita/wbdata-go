package wbdata

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func NewTestClient(t testing.TB, update bool) (*Client, func()) {
	cassete := filepath.Join("./fixtures/", t.Name())
	if update {
		if err := os.Remove(cassete + ".yaml"); err != nil {
			t.Fatal(err)
		}
	}

	r, err := recorder.New(cassete)
	if err != nil {
		t.Fatal(err)
	}
	customHTTPClient := &http.Client{
		Transport: r,
	}

	return NewClient(customHTTPClient), func() {
		if err := r.Stop(); err != nil {
			t.Errorf("failed to update fixtures: %s", err)
		}
	}
}
