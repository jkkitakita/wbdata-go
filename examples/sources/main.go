package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)

	// ListSources
	summary, sources, err := client.Sources.ListSources(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		fmt.Printf("failed to list sources: %+v\n", err)
	}
	fmt.Printf("ListSources, summary: %+v, sources[0]: %+v\n", summary, sources[0])

	// GetSource
	_, source, err := client.Sources.GetSource("1")
	if err != nil {
		fmt.Printf("failed to get source: %+v\n", err)
	}
	fmt.Printf("GetSource, source: %+v\n", source)
}
