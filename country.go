package wbdata

import (
	"fmt"
)

type (
	// CountriesService ...
	CountriesService service

	// Country contains information for an country field
	Country struct {
		ID          string
		Name        string
		CapitalCity string
		Iso2Code    string
		Longitude   string
		Latitude    string
		Region      CountryRegion
		IncomeLevel IncomeLevel
		LendingType LendingType
		AdminRegion CountryRegion
	}

	// ListCountryParams contains parameters for List
	ListCountryParams struct {
		RegionID      string
		IncomeLevelID string
		LendingTypeID string
	}
)

// List returns summary and countries with params
func (c *CountriesService) List(params ListCountryParams, pages PageParams) (*PageSummary, []*Country, error) {
	summary := &PageSummary{}
	countries := []*Country{}
	queryParams := map[string]string{
		"region":      params.RegionID,
		"incomelevel": params.IncomeLevelID,
		"lendingtype": params.LendingTypeID,
	}

	req, err := c.client.NewRequest("GET", "countries", queryParams, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err = c.client.do(req, &[]interface{}{summary, &countries}); err != nil {
		return nil, nil, err
	}

	return summary, countries, nil
}

// Get returns summary and a country
func (c *CountriesService) Get(countryID string) (*PageSummary, *Country, error) {
	summary := &PageSummary{}
	country := []*Country{}

	path := fmt.Sprintf("countries/%v", countryID)
	req, err := c.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = c.client.do(req, &[]interface{}{summary, &country}); err != nil {
		return nil, nil, err
	}

	return summary, country[0], nil
}
