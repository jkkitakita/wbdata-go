package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)

	// ListCountries
	summary, incomeLevels, err := client.IncomeLevels.ListIncomeLevels(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		// e.g. Outputs:
		// err: [{ID:150 Key:Language with ISO2 code: 'v3' is not yet supported in the API Value:Response requested in an unsupported language.}]
		fmt.Printf("err: %+v\n", err)
	}

	// e.g. Outputs:
	// ListIncomeLevels
	// summary: {Page:1 Pages:7 PerPage:1 Total:7}
	// incomeLevels: [{ID:HIC Iso2Code:XD Value:High income}]
	fmt.Println("ListIncomeLevels")
	fmt.Printf("summary: %+v\n", summary)
	fmt.Printf("incomeLevels: %+v\n", incomeLevels)

	// GetCountry
	_, incomeLevel, err := client.IncomeLevels.GetIncomeLevel("hic")
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	// e.g. Outputs:
	// GetIncomeLevel
	// incomeLevel: {ID:HIC Iso2Code:XD Value:High income}
	fmt.Println("GetIncomeLevel")
	fmt.Printf("incomeLevel: %+v\n", incomeLevel)
}
