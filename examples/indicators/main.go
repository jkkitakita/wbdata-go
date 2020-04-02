package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)

	// ListIndicators
	summary, indicators, err := client.Indicators.ListIndicators(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		fmt.Printf("failed to list incomelevels: %+v\n", err)
	}
	fmt.Printf("ListIndicators, summary: %+v, indicators[0]: %+v\n", summary, indicators[0])

	// GetIndicator
	_, indicator, err := client.Indicators.GetIndicator("1.0.HCount.1.90usd")
	if err != nil {
		fmt.Printf("failed to get indicator: %+v\n", err)
	}
	fmt.Printf("GetIndicator, indicator: %+v\n", indicator)
}
