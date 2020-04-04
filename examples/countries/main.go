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
	fmt.Printf("ListCountries, summary: %+v, countries[0]: %+v\n", summary, countries[0])

	// GetCountry
	summary, country, err := client.Countries.GetCountry("jpn")
	if err != nil {
		fmt.Printf("failed to get country: %+v\n", err)
	}
	fmt.Printf("GetCountry, summary: %+v, country: %+v\n", summary, country)
}
