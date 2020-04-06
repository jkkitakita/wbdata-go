# World Bank Open Data API in Go

[![GoDoc](https://godoc.org/github.com/jkkitakita/wbdata-go?status.svg)](https://godoc.org/github.com/jkkitakita/wbdata-go)
![Build Status](https://github.com/jkkitakita/wbdata-go/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/jkkitakita/wbdata-go/branch/master/graph/badge.svg)](https://codecov.io/gh/jkkitakita/wbdata-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkkitakita/wbdata-go)](https://goreportcard.com/report/github.com/jkkitakita/wbdata-go)

## Installing

### _go get_

    $ go get -u github.com/jkkitakita/wbdata-go

## Example

```go
import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)
	summary, countries, _ := client.Countries.ListCountries(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Countries[0] is: %#v\n", countries[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:304, PerPage:1, Total:304}
	// Countries[0] is: &wbdata.Country{ID:"ABW", Name:"Aruba", CapitalCity:"Oranjestad", Iso2Code:"AW", Longitude:"-70.0167", Latitude:"12.5167", Region:wbdata.Region{ID:"LCN", Code:"", Iso2Code:"ZJ", Value:"Latin America & Caribbean "}, IncomeLevel:wbdata.IncomeLevel{ID:"HIC", Iso2Code:"XD", Value:"High income"}, LendingType:wbdata.LendingType{ID:"LNX", Iso2Code:"XX", Value:"Not classified"}, AdminRegion:wbdata.Region{ID:"", Code:"", Iso2Code:"", Value:""}}
}
```

ref. [example_test.go](./example_test.go)
