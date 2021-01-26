package wbdata_test

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func ExampleCountriesService_List() {
	client := wbdata.NewClient(nil)
	summary, countries, _ := client.Countries.List(
		&wbdata.ListCountryParams{
			RegionID: "EAS",
		},
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		})

	summary.Pages = 30
	summary.Total = 300

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Countries[0] is: %#v\n", countries[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:30, PerPage:10, Total:300}
	// Countries[0] is: &wbdata.Country{ID:"ASM", Name:"American Samoa", CapitalCity:"Pago Pago", Iso2Code:"AS", Longitude:"-170.691", Latitude:"-14.2846", Region:wbdata.CountryRegion{ID:"EAS", Iso2Code:"Z4", Value:"East Asia & Pacific"}, IncomeLevel:wbdata.IncomeLevel{ID:"UMC", Iso2Code:"XT", Value:"Upper middle income"}, LendingType:wbdata.LendingType{ID:"LNX", Iso2Code:"XX", Value:"Not classified"}, AdminRegion:wbdata.CountryRegion{ID:"EAP", Iso2Code:"4E", Value:"East Asia & Pacific (excluding high income)"}}
}

func ExampleCountriesService_Get() {
	client := wbdata.NewClient(nil)
	summary, country, _ := client.Countries.Get("jpn")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Country is: %#v\n", country)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Country is: &wbdata.Country{ID:"JPN", Name:"Japan", CapitalCity:"Tokyo", Iso2Code:"JP", Longitude:"139.77", Latitude:"35.67", Region:wbdata.CountryRegion{ID:"EAS", Iso2Code:"Z4", Value:"East Asia & Pacific"}, IncomeLevel:wbdata.IncomeLevel{ID:"HIC", Iso2Code:"XD", Value:"High income"}, LendingType:wbdata.LendingType{ID:"LNX", Iso2Code:"XX", Value:"Not classified"}, AdminRegion:wbdata.CountryRegion{ID:"", Iso2Code:"", Value:""}}
}

func ExampleIncomeLevelsService_List() {
	client := wbdata.NewClient(nil)
	summary, incomeLevels, _ := client.IncomeLevels.List(
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 1
	summary.Total = 7

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IncomeLevels[0] is: %#v\n", incomeLevels[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:10, Total:7}
	// IncomeLevels[0] is: &wbdata.IncomeLevel{ID:"HIC", Iso2Code:"XD", Value:"High income"}
}

func ExampleIncomeLevelsService_Get() {
	client := wbdata.NewClient(nil)
	summary, incomeLevel, _ := client.IncomeLevels.Get("hic")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IncomeLevel is: %#v\n", incomeLevel)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// IncomeLevel is: &wbdata.IncomeLevel{ID:"HIC", Iso2Code:"XD", Value:"High income"}
}

func ExampleIndicatorsService_List() {
	client := wbdata.NewClient(nil)
	summary, indicators, _ := client.Indicators.List(
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 1735
	summary.Total = 17349
	indicators[0].Source = nil
	indicators[0].Topics = nil

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Indicators[0] without Source and Topics is: %#v\n", indicators[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1735, PerPage:10, Total:17349}
	// Indicators[0] without Source and Topics is: &wbdata.Indicator{ID:"1.0.HCount.1.90usd", Name:"Poverty Headcount ($1.90 a day)", Unit:"", Source:(*wbdata.IDAndValue)(nil), SourceNote:"The poverty headcount index measures the proportion of the population with daily per capita income (in 2011 PPP) below the poverty line.", SourceOrganization:"LAC Equity Lab tabulations of SEDLAC (CEDLAS and the World Bank).", Topics:[]*wbdata.IDAndValue(nil)}
}

func ExampleIndicatorsService_Get() {
	client := wbdata.NewClient(nil)
	summary, indicator, _ := client.Indicators.Get("1.0.hcount.1.90usd")

	indicator.Source = nil
	indicator.Topics = nil

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Indicator without Source and Topics is: %#v\n", indicator)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Indicator without Source and Topics is: &wbdata.Indicator{ID:"1.0.HCount.1.90usd", Name:"Poverty Headcount ($1.90 a day)", Unit:"", Source:(*wbdata.IDAndValue)(nil), SourceNote:"The poverty headcount index measures the proportion of the population with daily per capita income (in 2011 PPP) below the poverty line.", SourceOrganization:"LAC Equity Lab tabulations of SEDLAC (CEDLAS and the World Bank).", Topics:[]*wbdata.IDAndValue(nil)}
}

func ExampleIndicatorsService_ListByTopicID() {
	client := wbdata.NewClient(nil)
	topicID := "1"
	summary, indicators, _ := client.Indicators.ListByTopicID(
		topicID,
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 1735
	summary.Total = 17349
	topics := indicators[0].Topics
	indicators[0].Source = nil
	indicators[0].Topics = nil

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Indicators[0] without Source and Topics is: %#v\n", indicators[0])
	fmt.Printf("Topics is: %+v\n", topics[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1735, PerPage:10, Total:17349}
	// Indicators[0] without Source and Topics is: &wbdata.Indicator{ID:"AG.AGR.TRAC.NO", Name:"Agricultural machinery, tractors", Unit:"", Source:(*wbdata.IDAndValue)(nil), SourceNote:"Agricultural machinery refers to the number of wheel and crawler tractors (excluding garden tractors) in use in agriculture at the end of the calendar year specified or during the first quarter of the following year.", SourceOrganization:"Food and Agriculture Organization, electronic files and web site.", Topics:[]*wbdata.IDAndValue(nil)}
	// Topics is: &{ID:1 Value:Agriculture & Rural Development  }
}

func ExampleIndicatorValuesService_List() {
	client := wbdata.NewClient(nil)
	indicatorID := "NY.GDP.MKTP.CD"
	summary, indicatorValues, _ := client.IndicatorValues.List(
		indicatorID,
		&wbdata.FilterParams{
			FilterParamsType: wbdata.FilterParamsDateRange,
			DateParam: &wbdata.DateParam{
				DateRange: &wbdata.DateRange{
					Start: "2018",
					End:   "2019",
				},
			},
		},
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 53
	summary.Total = 528
	summary.LastUpdated = "2020-12-16"

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IndicatorValues[0]: %#v\n", indicatorValues[0])
	// Output:
	// Summary is: &wbdata.PageSummaryWithSourceID{Page:1, Pages:53, PerPage:10, Total:528, SourceID:"2", LastUpdated:"2020-12-16"}
	// IndicatorValues[0]: &wbdata.IndicatorValue{Indicator:wbdata.IDAndValue{ID:"NY.GDP.MKTP.CD", Value:"GDP (current US$)"}, Country:wbdata.IDAndValue{ID:"1A", Value:"Arab World"}, Countryiso3code:"ARB", Date:"2019", Value:2.81741458466511e+12, Unit:"", ObsStatus:"", Decimal:0}
}

func ExampleIndicatorValuesService_List_second() {
	client := wbdata.NewClient(nil)
	indicatorID := "NY.GDP.MKTP.CD"
	summary, indicatorValues, _ := client.IndicatorValues.List(
		indicatorID,
		&wbdata.FilterParams{
			FilterParamsType: wbdata.FilterParamsMRV,
			RecentParam: &wbdata.RecentParam{
				FrequencyType:    wbdata.FrequencyYearly,
				MostRecentValues: 1,
				IsNotEmpty:       false,
				IsGapFill:        false,
			},
		},
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 53
	summary.Total = 528
	summary.LastUpdated = "2020-12-16"

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IndicatorValues[0]: %#v\n", indicatorValues[0])
	// Output:
	// Summary is: &wbdata.PageSummaryWithSourceID{Page:1, Pages:53, PerPage:10, Total:528, SourceID:"2", LastUpdated:"2020-12-16"}
	// IndicatorValues[0]: &wbdata.IndicatorValue{Indicator:wbdata.IDAndValue{ID:"NY.GDP.MKTP.CD", Value:"GDP (current US$)"}, Country:wbdata.IDAndValue{ID:"1A", Value:"Arab World"}, Countryiso3code:"ARB", Date:"2019", Value:2.81741458466511e+12, Unit:"", ObsStatus:"", Decimal:0}
}

func ExampleIndicatorValuesService_ListByCountryIDs() {
	client := wbdata.NewClient(nil)
	countryIDs := []string{"JPN", "USA"}
	indicatorID := "NY.GDP.MKTP.CD"
	summary, indicatorValues, _ := client.IndicatorValues.ListByCountryIDs(
		countryIDs,
		indicatorID,
		&wbdata.FilterParams{
			FilterParamsType: wbdata.FilterParamsDateRange,
			DateParam: &wbdata.DateParam{
				DateRange: &wbdata.DateRange{
					Start: "2018",
					End:   "2019",
				},
			},
		},
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.LastUpdated = "2020-12-16"

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IndicatorValues[0]: %#v\n", indicatorValues[0])
	// Output:
	// Summary is: &wbdata.PageSummaryWithSourceID{Page:1, Pages:1, PerPage:10, Total:4, SourceID:"2", LastUpdated:"2020-12-16"}
	// IndicatorValues[0]: &wbdata.IndicatorValue{Indicator:wbdata.IDAndValue{ID:"NY.GDP.MKTP.CD", Value:"GDP (current US$)"}, Country:wbdata.IDAndValue{ID:"JP", Value:"Japan"}, Countryiso3code:"JPN", Date:"2019", Value:5.08176954237977e+12, Unit:"", ObsStatus:"", Decimal:0}
}

func ExampleIndicatorValuesService_ListBySourceID() {
	client := wbdata.NewClient(nil)
	indicatorIDs := []string{"NY.GDP.MKTP.CD", "SP.POP.TOTL"}
	sourceID := "2"
	summary, indicatorValues, _ := client.IndicatorValues.ListBySourceID(
		indicatorIDs,
		sourceID,
		&wbdata.FilterParams{
			FilterParamsType: wbdata.FilterParamsDateRange,
			DateParam: &wbdata.DateParam{
				DateRange: &wbdata.DateRange{
					Start: "2018",
					End:   "2019",
				},
			},
		},
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 106
	summary.Total = 1056
	summary.LastUpdated = "2020-12-16"

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IndicatorValues[0]: %#v\n", indicatorValues[0])
	// Output:
	// Summary is: &wbdata.PageSummaryWithLastUpdated{Page:1, Pages:106, PerPage:10, Total:1056, LastUpdated:"2020-12-16"}
	// IndicatorValues[0]: &wbdata.IndicatorValue{Indicator:wbdata.IDAndValue{ID:"NY.GDP.MKTP.CD", Value:"GDP (current US$)"}, Country:wbdata.IDAndValue{ID:"1A", Value:"Arab World"}, Countryiso3code:"ARB", Date:"2019", Value:2.81741458466511e+12, Unit:"", ObsStatus:"", Decimal:0}
}

func ExampleIndicatorValuesService_ListByCountryIDsAndSourceID() {
	client := wbdata.NewClient(nil)
	countryIDs := []string{"JPN", "USA"}
	indicatorIDs := []string{"NY.GDP.MKTP.CD", "SP.POP.TOTL"}
	sourceID := "2"
	summary, indicatorValues, _ := client.IndicatorValues.ListByCountryIDsAndSourceID(
		countryIDs,
		indicatorIDs,
		sourceID,
		&wbdata.FilterParams{
			FilterParamsType: wbdata.FilterParamsDateRange,
			DateParam: &wbdata.DateParam{
				DateRange: &wbdata.DateRange{
					Start: "2018",
					End:   "2019",
				},
			},
		},
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 106
	summary.Total = 1056
	summary.LastUpdated = "2020-12-16"

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("IndicatorValues[0]: %#v\n", indicatorValues[0])
	// Output:
	// Summary is: &wbdata.PageSummaryWithLastUpdated{Page:1, Pages:106, PerPage:10, Total:1056, LastUpdated:"2020-12-16"}
	// IndicatorValues[0]: &wbdata.IndicatorValue{Indicator:wbdata.IDAndValue{ID:"NY.GDP.MKTP.CD", Value:"GDP (current US$)"}, Country:wbdata.IDAndValue{ID:"JP", Value:"Japan"}, Countryiso3code:"JPN", Date:"2019", Value:5.08176954237977e+12, Unit:"", ObsStatus:"", Decimal:0}
}

func ExampleLendingTypesService_List() {
	client := wbdata.NewClient(nil)
	summary, lendingTypes, _ := client.LendingTypes.List(
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 1
	summary.Total = 4

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("LendingTypes[0] is: %#v\n", lendingTypes[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:10, Total:4}
	// LendingTypes[0] is: &wbdata.LendingType{ID:"IBD", Iso2Code:"XF", Value:"IBRD"}
}

func ExampleLendingTypesService_Get() {
	client := wbdata.NewClient(nil)
	summary, lendingType, _ := client.LendingTypes.Get("ibd")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("LendingType is: %#v\n", lendingType)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// LendingType is: &wbdata.LendingType{ID:"IBD", Iso2Code:"XF", Value:"IBRD"}
}

func ExampleRegionsService_List() {
	client := wbdata.NewClient(nil)
	summary, regions, _ := client.Regions.List(
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 5
	summary.Total = 48

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Regions[0] is: %#v\n", regions[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:5, PerPage:10, Total:48}
	// Regions[0] is: &wbdata.Region{ID:"", Code:"AFR", Iso2Code:"A9", Name:"Africa"}
}

func ExampleRegionsService_Get() {
	client := wbdata.NewClient(nil)
	summary, region, _ := client.Regions.Get("xzn")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Region is: %#v\n", region)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Region is: &wbdata.Region{ID:"", Code:"XZN", Iso2Code:"A5", Name:"Sub-Saharan Africa excluding South Africa and Nigeria"}
}

func ExampleSourcesService_List() {
	client := wbdata.NewClient(nil)
	summary, sources, _ := client.Sources.List(
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 6
	summary.Total = 59

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Sources[0] is: %#v\n", sources[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:6, PerPage:10, Total:59}
	// Sources[0] is: &wbdata.Source{ID:"1", LastUpdated:"2019-10-23", Name:"Doing Business", Code:"DBS", Description:"", URL:"", DataAvailability:"Y", MetadataAvailability:"Y", Concepts:"3"}
}

func ExampleSourcesService_Get() {
	client := wbdata.NewClient(nil)
	summary, source, _ := client.Sources.Get("1")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Source is: %#v\n", source)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Source is: &wbdata.Source{ID:"1", LastUpdated:"2019-10-23", Name:"Doing Business", Code:"DBS", Description:"", URL:"", DataAvailability:"Y", MetadataAvailability:"Y", Concepts:"3"}
}

func ExampleTopicsService_List() {
	client := wbdata.NewClient(nil)
	summary, topics, _ := client.Topics.List(
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 3
	summary.Total = 21

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Topics[0] is: %#v\n", topics[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:3, PerPage:10, Total:21}
	// Topics[0] is: &wbdata.Topic{ID:"1", Value:"Agriculture & Rural Development", SourceNote:"For the 70 percent of the world's poor who live in rural areas, agriculture is the main source of income and employment. But depletion and degradation of land and water pose serious challenges to producing enough food and other agricultural products to sustain livelihoods here and meet the needs of urban populations. Data presented here include measures of agricultural inputs, outputs, and productivity compiled by the UN's Food and Agriculture Organization."}
}

func ExampleTopicsService_Get() {
	client := wbdata.NewClient(nil)
	summary, topic, _ := client.Topics.Get("1")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Topic is: %#v\n", topic)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Topic is: &wbdata.Topic{ID:"1", Value:"Agriculture & Rural Development", SourceNote:"For the 70 percent of the world's poor who live in rural areas, agriculture is the main source of income and employment. But depletion and degradation of land and water pose serious challenges to producing enough food and other agricultural products to sustain livelihoods here and meet the needs of urban populations. Data presented here include measures of agricultural inputs, outputs, and productivity compiled by the UN's Food and Agriculture Organization."}
}

func ExampleLanguagesService_List() {
	client := wbdata.NewClient(nil)
	summary, languages, _ := client.Languages.List(
		&wbdata.PageParams{
			Page:    1,
			PerPage: 10,
		},
	)

	summary.Pages = 3
	summary.Total = 23

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Languages[0] is: %#v\n", languages[0])
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:3, PerPage:10, Total:23}
	// Languages[0] is: &wbdata.Language{Code:"en", Name:"English", NativeForm:"English"}
}

func ExampleLanguagesService_Get() {
	client := wbdata.NewClient(nil)
	summary, language, _ := client.Languages.Get("ja")

	fmt.Printf("Summary is: %#v\n", summary)
	fmt.Printf("Language is: %#v\n", language)
	// Output:
	// Summary is: &wbdata.PageSummary{Page:1, Pages:1, PerPage:50, Total:1}
	// Language is: &wbdata.Language{Code:"ja", Name:"Japanese ", NativeForm:"日本語"}
}
