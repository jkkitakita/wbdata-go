package main

import (
	"fmt"

	wbdata "github.com/jkkitakita/wbdata-go"
)

func main() {
	fmt.Println("hello")

	client := wbdata.NewClient(nil)
	countries, err := client.Countries.ListCountries()

	fmt.Prinln(countries)
}
