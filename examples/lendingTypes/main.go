package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)

	// ListLendingTypes
	summary, lendingTypes, err := client.LendingTypes.ListLendingTypes(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		fmt.Printf("failed to list incomelevels: %+v\n", err)
	}
	fmt.Printf("ListLendingTypes, summary: %+v, lendingTypes[0]: %+v\n", summary, lendingTypes[0])

	// GetLendingType
	_, lendingType, err := client.LendingTypes.GetLendingType("IBD")
	if err != nil {
		fmt.Printf("failed to get lendingType: %+v\n", err)
	}
	fmt.Printf("GetLendingType, lendingType: %+v\n", lendingType)
}
