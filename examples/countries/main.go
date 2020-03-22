package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)

	// ListCountries
	summary, countries, err := client.Countries.ListCountries(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		// e.g. Outputs:
		// err: [{ID:150 Key:Language with ISO2 code: 'v3' is not yet supported in the API Value:Response requested in an unsupported language.}]
		fmt.Printf("err: %+v\n", err)
	}

	// e.g. Outputs:
	// ListCountries
	// summary: {Page:1 Pages:304 PerPage:1 Total:304}
	// countries: [{ID:ABW Name:Aruba CapitalCity:Oranjestad Iso2Code:AW Longitude:-70.0167 Latitude:12.5167 Region:{ID:LCN Iso2Code:ZJ Code:} IncomeLevels:{ID: Iso2Code: Value:} LendingType:{ID:LNX Iso2Code:XX Value:Not classified} AdminRegion:{ID: Iso2Code: Value:}}]
	fmt.Println("ListCountries")
	fmt.Printf("summary: %+v\n", summary)
	fmt.Printf("countries: %+v\n", countries)

	// GetCountry
	_, country, err := client.Countries.GetCountry("jpn")
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}

	// e.g. Outputs:
	// GetCountry
	// country: {ID:JPN Name:Japan CapitalCity:Tokyo Iso2Code:JP Longitude:139.77 Latitude:35.67 Region:{ID:EAS Iso2Code:Z4 Code:} IncomeLevels:{ID: Iso2Code: Value:} LendingType:{ID:LNX Iso2Code:XX Value:Not classified} AdminRegion:{ID: Iso2Code: Value:}}
	fmt.Println("GetCountry")
	fmt.Printf("country: %+v\n", country)
}
