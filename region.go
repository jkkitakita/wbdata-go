package wbdata

import "fmt"

type (
	// RegionsService ...
	RegionsService service

	// Region is a struct for region
	Region struct {
		ID       string
		Code     string
		Iso2Code string
		Name     string
	}

	// CountryRegion is a struct for region when using the Countries API
	CountryRegion struct {
		ID       string
		Iso2Code string
		Value    string
	}
)

// List returns a Response's Summary and Regions
func (r *RegionsService) List(pages PageParams) (*PageSummary, []*Region, error) {
	summary := &PageSummary{}
	regions := []*Region{}

	req, err := r.client.NewRequest("GET", "regions", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	if err = r.client.do(req, &[]interface{}{summary, &regions}); err != nil {
		return nil, nil, err
	}

	return summary, regions, nil
}

// Get returns a Response's Summary and a Region
func (r *RegionsService) Get(code string) (*PageSummary, *Region, error) {
	summary := &PageSummary{}
	region := []*Region{}

	path := fmt.Sprintf("regions/%v", code)
	req, err := r.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = r.client.do(req, &[]interface{}{summary, &region}); err != nil {
		return nil, nil, err
	}

	return summary, region[0], nil
}
