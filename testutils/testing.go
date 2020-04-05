package testutils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	TestDefaultPage    = 1
	TestDefaultPerPage = 10
)

func UpdateFixture(update *bool) {
	if *update {
		fixtureDir := filepath.Join("testdata", "fixtures")
		os.RemoveAll(fixtureDir)
		if err := os.MkdirAll(fixtureDir, 0755); err != nil {
			panic(fmt.Sprintf("failed to create fixture dirs: %s", err))
		}
	}
}
