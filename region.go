package wbdata

import "fmt"

type (
	RegionsService service

	Region struct {
		ID       string
		Code     string
		Iso2Code string
		Value    string
	}
)

func (r *RegionsService) ListRegions(pages PageParams) (*PageSummary, []*Region, error) {
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

func (r *RegionsService) GetRegion(code string) (*PageSummary, *Region, error) {
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
