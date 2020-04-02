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
		fmt.Printf("failed to list countries: %+v\n", err)
	}
	fmt.Printf("ListCountries, summary: %+v, countries: %+v\n", summary, countries)

	// GetCountry
	_, country, err := client.Countries.GetCountry("jpn")
	if err != nil {
		fmt.Printf("failed to get country: %+v\n", err)
	}
	fmt.Printf("GetCountry, country: %+v\n", country)
}
