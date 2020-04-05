package wbdata

import (
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

const (
	TestDefaultPage    = 1
	TestDefaultPerPage = 10
)

func NewTestClient(t testing.TB, update bool) (*Client, func()) {
	funcName := strings.Split(t.Name(), "_")
	fixtureDir := filepath.Join("testdata", "fixtures")
	cassete := filepath.Join(fixtureDir, funcName[1])

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
