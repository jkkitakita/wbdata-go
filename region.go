package wbdata

import "fmt"

type (
	RegionsService service

	Region struct {
		ID       string
		Iso2Code string
		Value    string
	}
)

func (c *RegionsService) ListRegions(pages PageParams) (*PageSummary, []*Region, error) {
	summary := &PageSummary{}
	regions := []*Region{}

	req, err := c.client.NewRequest("GET", "regions", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &regions})
	if err != nil {
		return nil, nil, err
	}

	return summary, regions, nil
}

func (c *RegionsService) GetRegion(code string) (*PageSummary, *Region, error) {
	summary := &PageSummary{}
	region := []*Region{}

	s := fmt.Sprintf("regions/%v", code)
	req, err := c.client.NewRequest("GET", s, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &region})
	if err != nil {
		return nil, nil, err
	}

	return summary, region[0], nil
}
