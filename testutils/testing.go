package testutils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// TestDefaultPage is the default number of pages for testing
	TestDefaultPage = 1
	// TestDefaultPerPage is the default number of pages per page for testing
	TestDefaultPerPage = 2
	// JaLocalLanguage is local language for Japan
	JaLocalLanguage = "ja"
)

// UpdateFixture removes fixtures when `update` is true.
func UpdateFixture(update *bool) {
	if *update {
		fixtureDir := filepath.Join("testdata", "fixtures")
		os.RemoveAll(fixtureDir)
		if err := os.MkdirAll(fixtureDir, 0755); err != nil {
			panic(fmt.Sprintf("failed to create fixture dirs: %s", err))
		}
	}
}
