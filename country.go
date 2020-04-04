package wbdata

import (
	"fmt"
)

type (
	CountriesService service

	Country struct {
		ID          string
		Name        string
		CapitalCity string
		Iso2Code    string
		Longitude   string
		Latitude    string
		Region      Region
		IncomeLevel IncomeLevel
		LendingType LendingType
		AdminRegion Region
	}
)

func (c *CountriesService) ListCountries(pages PageParams) (*PageSummary, []*Country, error) {
	summary := &PageSummary{}
	countries := []*Country{}

	req, err := c.client.NewRequest("GET", "countries", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &countries})
	if err != nil {
		return nil, nil, err
	}

	return summary, countries, nil
}

func (c *CountriesService) GetCountry(countryID string) (*PageSummary, *Country, error) {
	summary := &PageSummary{}
	country := []*Country{}

	s := fmt.Sprintf("countries/%v", countryID)
	req, err := c.client.NewRequest("GET", s, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &country})
	if err != nil {
		return nil, nil, err
	}

	return summary, country[0], nil
}
