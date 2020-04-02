package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)

	// ListRegions
	summary, regions, err := client.Regions.ListRegions(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		fmt.Printf("failed to list regions: %+v\n", err)
	}
	fmt.Printf("ListRegions, summary: %+v, regions[0]: %+v\n", summary, regions[0])

	// GetRegion
	_, region, err := client.Regions.GetRegion("XZN")
	if err != nil {
		fmt.Printf("failed to get region: %+v\n", err)
	}
	fmt.Printf("GetRegion, region: %+v\n", region)
}
