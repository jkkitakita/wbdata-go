package wbdata_test

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/dnaeon/go-vcr/recorder"

	"github.com/jkkitakita/wbdata-go"
	"github.com/jkkitakita/wbdata-go/testutils"
)

// var update = flag.Bool("update", false, "update fixtures")

func ExampleListCountries() {
	client, save := newExampleClient("ListCountries")
	defer save()
	summary, countries, _ := client.Countries.ListCountries(wbdata.PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	})

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Countries[0] is: %#v\n", countries[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:31, PerPage:10, Total:304}
	// Countries[0] is: &wbdata.Country{ID:"ABW", Name:"Aruba", CapitalCity:"Oranjestad", Iso2Code:"AW", Longitude:"-70.0167", Latitude:"12.5167", Region:wbdata.CountryRegion{ID:"LCN", Iso2Code:"ZJ", Value:"Latin America & Caribbean "}, IncomeLevel:wbdata.IncomeLevel{ID:"HIC", Iso2Code:"XD", Value:"High income"}, LendingType:wbdata.LendingType{ID:"LNX", Iso2Code:"XX", Value:"Not classified"}, AdminRegion:wbdata.CountryRegion{ID:"", Iso2Code:"", Value:""}}
}

func ExampleGetCountry() {
	client, save := newExampleClient("GetCountry")
	defer save()
	summary, country, _ := client.Countries.GetCountry("jpn")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Country is: %#v\n", country)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Country is: &wbdata.Country{ID:"JPN", Name:"Japan", CapitalCity:"Tokyo", Iso2Code:"JP", Longitude:"139.77", Latitude:"35.67", Region:wbdata.CountryRegion{ID:"EAS", Iso2Code:"Z4", Value:"East Asia & Pacific"}, IncomeLevel:wbdata.IncomeLevel{ID:"HIC", Iso2Code:"XD", Value:"High income"}, LendingType:wbdata.LendingType{ID:"LNX", Iso2Code:"XX", Value:"Not classified"}, AdminRegion:wbdata.CountryRegion{ID:"", Iso2Code:"", Value:""}}
}

func ExampleListIncomeLevels() {
	client, save := newExampleClient("ListIncomeLevels")
	defer save()
	summary, incomeLevels, _ := client.IncomeLevels.ListIncomeLevels(wbdata.PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	})

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IncomeLevels[0] is: %#v\n", incomeLevels[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:10, Total:7}
	// IncomeLevels[0] is: &wbdata.IncomeLevel{ID:"HIC", Iso2Code:"XD", Value:"High income"}
}

func ExampleGetIncomeLevel() {
	client, save := newExampleClient("GetIncomeLevel")
	defer save()
	summary, incomeLevel, _ := client.IncomeLevels.GetIncomeLevel("hic")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IncomeLevel is: %#v\n", incomeLevel)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// IncomeLevel is: &wbdata.IncomeLevel{ID:"HIC", Iso2Code:"XD", Value:"High income"}
}

func ExampleListIndicators() {
	client, save := newExampleClient("ListIndicators")
	defer save()
	summary, indicators, _ := client.Indicators.ListIndicators(wbdata.PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	})

	indicators[0].Source = nil
	indicators[0].Topics = nil

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Indicators[0] without Source and Topics is: %#v\n", indicators[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1734, PerPage:10, Total:17338}
	// Indicators[0] without Source and Topics is: &wbdata.Indicator{ID:"1.0.HCount.1.90usd", Name:"Poverty Headcount ($1.90 a day)", Source:(*wbdata.Source)(nil), SourceNote:"The poverty headcount index measures the proportion of the population with daily per capita income (in 2011 PPP) below the poverty line.", SourceOrganization:"LAC Equity Lab tabulations of SEDLAC (CEDLAS and the World Bank).", Topics:[]*wbdata.Topic(nil)}
}

func ExampleGetIndicator() {
	client, save := newExampleClient("GetIndicator")
	defer save()
	summary, indicator, _ := client.Indicators.GetIndicator("1.0.hcount.1.90usd")

	indicator.Source = nil
	indicator.Topics = nil

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Indicator without Source and Topics is: %#v\n", indicator)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Indicator without Source and Topics is: &wbdata.Indicator{ID:"1.0.HCount.1.90usd", Name:"Poverty Headcount ($1.90 a day)", Source:(*wbdata.Source)(nil), SourceNote:"The poverty headcount index measures the proportion of the population with daily per capita income (in 2011 PPP) below the poverty line.", SourceOrganization:"LAC Equity Lab tabulations of SEDLAC (CEDLAS and the World Bank).", Topics:[]*wbdata.Topic(nil)}
}

func ExampleListLendingTypes() {
	client, save := newExampleClient("ListLendingTypes")
	defer save()
	summary, lendingTypes, _ := client.LendingTypes.ListLendingTypes(wbdata.PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	})

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("LendingTypes[0] is: %#v\n", lendingTypes[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:10, Total:4}
	// LendingTypes[0] is: &wbdata.LendingType{ID:"IBD", Iso2Code:"XF", Value:"IBRD"}
}

func ExampleGetLendingType() {
	client, save := newExampleClient("GetLendingType")
	defer save()
	summary, lendingType, _ := client.LendingTypes.GetLendingType("ibd")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("LendingType is: %#v\n", lendingType)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// LendingType is: &wbdata.LendingType{ID:"IBD", Iso2Code:"XF", Value:"IBRD"}
}

func ExampleListRegions() {
	client, save := newExampleClient("ListRegions")
	defer save()
	summary, regions, _ := client.Regions.ListRegions(wbdata.PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	})

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Regions[0] is: %#v\n", regions[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:5, PerPage:10, Total:48}
	// Regions[0] is: &wbdata.Region{ID:"", Code:"AFR", Iso2Code:"A9", Name:"Africa"}
}

func ExampleGetRegion() {
	client, save := newExampleClient("GetRegion")
	defer save()
	summary, region, _ := client.Regions.GetRegion("xzn")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Region is: %#v\n", region)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Region is: &wbdata.Region{ID:"", Code:"XZN", Iso2Code:"A5", Name:"Sub-Saharan Africa excluding South Africa and Nigeria"}
}

func ExampleListSources() {
	client, save := newExampleClient("ListSources")
	defer save()
	summary, sources, _ := client.Sources.ListSources(wbdata.PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	})

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Sources[0] is: %#v\n", sources[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:6, PerPage:10, Total:59}
	// Sources[0] is: &wbdata.Source{ID:"1", LastUpdated:"2019-10-23", Name:"Doing Business", Code:"DBS", Description:"", URL:"", DataAvailability:"Y", MetadataAvailability:"Y", Concepts:"3"}
}

func ExampleGetSource() {
	client, save := newExampleClient("GetSource")
	defer save()
	summary, source, _ := client.Sources.GetSource("1")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Source is: %#v\n", source)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Source is: &wbdata.Source{ID:"1", LastUpdated:"2019-10-23", Name:"Doing Business", Code:"DBS", Description:"", URL:"", DataAvailability:"Y", MetadataAvailability:"Y", Concepts:"3"}
}

func ExampleListTopics() {
	client, save := newExampleClient("ListTopics")
	defer save()
	summary, topics, _ := client.Topics.ListTopics(wbdata.PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	})

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Topics[0] is: %#v\n", topics[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:3, PerPage:10, Total:21}
	// Topics[0] is: &wbdata.Topic{ID:"1", Value:"Agriculture & Rural Development", SourceNote:"For the 70 percent of the world's poor who live in rural areas, agriculture is the main source of income and employment. But depletion and degradation of land and water pose serious challenges to producing enough food and other agricultural products to sustain livelihoods here and meet the needs of urban populations. Data presented here include measures of agricultural inputs, outputs, and productivity compiled by the UN's Food and Agriculture Organization."}
}

func ExampleGetTopic() {
	client, save := newExampleClient("GetTopic")
	defer save()
	summary, topic, _ := client.Topics.GetTopic("1")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Topic is: %#v\n", topic)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Topic is: &wbdata.Topic{ID:"1", Value:"Agriculture & Rural Development", SourceNote:"For the 70 percent of the world's poor who live in rural areas, agriculture is the main source of income and employment. But depletion and degradation of land and water pose serious challenges to producing enough food and other agricultural products to sustain livelihoods here and meet the needs of urban populations. Data presented here include measures of agricultural inputs, outputs, and productivity compiled by the UN's Food and Agriculture Organization."}
}

func newExampleClient(name string) (*wbdata.Client, func()) {
	cassette := filepath.Join("./testdata/fixtures", name)

	r, err := recorder.New(cassette)
	if err != nil {
		panic(err)
	}
	customHTTPClient := &http.Client{
		Transport: r,
	}

	return wbdata.NewClient(customHTTPClient), func() {
		if err := r.Stop(); err != nil {
			panic(err)
		}
	}
}
