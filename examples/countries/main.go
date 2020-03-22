package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)
	summary, countries, err := client.Countries.ListCountries(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	fmt.Printf("summary: %+v\n", summary)
	fmt.Printf("countries: %+v\n", countries)
	// e.g. Outputs:
	// summary: {Page:1 Pages:304 PerPage:1 Total:304}
	// countries: [{ID:ABW Name:Aruba CapitalCity:Oranjestad Iso2Code:AW Longitude:-70.0167 Latitude:12.5167 Region:{ID:LCN Iso2Code:ZJ Code:} IncomeLevels:{ID: Iso2Code: Value:} LendingType:{ID:LNX Iso2Code:XX Value:Not classified} AdminRegion:{ID: Iso2Code: Value:}}]
}
