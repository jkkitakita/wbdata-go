package wbdata

import "strconv"

type CountriesService service

type Country struct {
	ID           string
	Name         string
	CapitalCity  string
	Iso2Code     string
	Longitude    string
	Latitude     string
	Region       Region
	IncomeLevels IncomeLevel
	LendingType  LendingType
	AdminRegion  struct {
		ID    string
		Value string
	}
}

func (c *CountriesService) ListCountries(pages PageParams) ([]Country, error) {
	country := []Country{}

	req, err := c.client.NewRequest("GET", "countries", nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	if pages.Page != 0 {
		params.Add(`page`, strconv.Itoa(int(pages.Page)))
	}
	if pages.PerPages != 0 {
		params.Add(`per_pages`, strconv.Itoa(int(pages.PerPages)))
	}
	req.URL.RawQuery = params.Encode()

	_, err = c.client.do(req)

	return country, err
}
