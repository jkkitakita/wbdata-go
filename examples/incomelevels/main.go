package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)

	// ListIncomeLevels
	summary, incomeLevels, err := client.IncomeLevels.ListIncomeLevels(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		fmt.Printf("failed to list incomelevels: %+v\n", err)
	}
	fmt.Printf("ListIncomeLevels, summary: %+v, incomeLevels[0]: %+v\n", summary, incomeLevels[0])

	// GetIncomeLevel
	_, incomeLevel, err := client.IncomeLevels.GetIncomeLevel("hic")
	if err != nil {
		fmt.Printf("failed to get incomeLevel: %+v\n", err)
	}
	fmt.Printf("GetIncomeLevel, incomeLevel: %+v\n", incomeLevel)
}
