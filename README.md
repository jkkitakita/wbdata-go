# World Bank Open Data API in Go

[![GoDoc](https://godoc.org/github.com/jkkitakita/wbdata-go?status.svg)](https://godoc.org/github.com/jkkitakita/wbdata-go)
![Build Status](https://github.com/jkkitakita/wbdata-go/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/jkkitakita/wbdata-go/branch/main/graph/badge.svg)](https://codecov.io/gh/jkkitakita/wbdata-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkkitakita/wbdata-go)](https://goreportcard.com/report/github.com/jkkitakita/wbdata-go)

## Installing

```shell
go get -u github.com/jkkitakita/wbdata-go
```

## Example

```go
package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main () {
	client := wbdata.NewClient(nil)
	summary, countries, _ := client.Countries.List(
		&wbdata.ListCountryParams{
			RegionID: "EAS",
		},
		&wbdata.PageParams{
			Page:    1,
			PerPage: 1,
		},
	)

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Countries[0] is: %#v\n", countries[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:38, PerPage:1, Total:38}
	// Countries[0] is: &wbdata.Country{ID:"ASM", Name:"American Samoa", CapitalCity:"Pago Pago", Iso2Code:"AS", Longitude:"-170.691", Latitude:"-14.2846", Region:wbdata.CountryRegion{ID:"EAS", Iso2Code:"Z4", Value:"East Asia & Pacific"}, IncomeLevel:wbdata.IncomeLevel{ID:"UMC", Iso2Code:"XT", Value:"Upper middle income"}, LendingType:wbdata.LendingType{ID:"LNX", Iso2Code:"XX", Value:"Not classified"}, AdminRegion:wbdata.CountryRegion{ID:"EAP", Iso2Code:"4E", Value:"East Asia & Pacific (excluding high income)"}}
}
```

ref. [example_test.go](./example_test.go)
