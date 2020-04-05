package wbdata

import "fmt"

type (
	// IndicatorsService ...
	IndicatorsService service

	// Indicator contains information for an indicator field
	Indicator struct {
		ID                 string
		Name               string
		Source             *Source
		SourceNote         string
		SourceOrganization string
		Topics             []*Topic
	}
)

// ListIndicators returns a Response's Summary and Indicators
func (i *IndicatorsService) ListIndicators(pages PageParams) (*PageSummary, []*Indicator, error) {
	summary := &PageSummary{}
	indicators := []*Indicator{}

	req, err := i.client.NewRequest("GET", "indicators", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicators}); err != nil {
		return nil, nil, err
	}

	return summary, indicators, nil
}

// GetIndicator returns a Response's Summary and an Indicator
func (i *IndicatorsService) GetIndicator(indicatorID string) (*PageSummary, *Indicator, error) {
	summary := &PageSummary{}
	indicator := []*Indicator{}

	path := fmt.Sprintf("indicators/%v", indicatorID)
	req, err := i.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicator}); err != nil {
		return nil, nil, err
	}

	return summary, indicator[0], nil
}
