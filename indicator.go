package wbdata

import "fmt"

type (
	IndicatorsService service

	Indicator struct {
		ID                 string
		Name               string
		Source             *Source
		SourceNote         string
		SourceOrganization string
		Topics             []*Topic
	}
)

func (c *IndicatorsService) ListIndicators(pages PageParams) (*PageSummary, []*Indicator, error) {
	summary := &PageSummary{}
	indicators := []*Indicator{}

	req, err := c.client.NewRequest("GET", "indicators", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &indicators})
	if err != nil {
		return nil, nil, err
	}

	return summary, indicators, nil
}

func (c *IndicatorsService) GetIndicator(indicatorID string) (*PageSummary, *Indicator, error) {
	summary := &PageSummary{}
	indicator := []*Indicator{}

	s := fmt.Sprintf("indicators/%v", indicatorID)
	req, err := c.client.NewRequest("GET", s, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &indicator})
	if err != nil {
		return nil, nil, err
	}

	return summary, indicator[0], nil
}
