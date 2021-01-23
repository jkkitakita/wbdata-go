package wbdata

import (
	"fmt"
)

type (
	// IndicatorsService ...
	IndicatorsService service

	// Indicator contains information for an indicator field
	Indicator struct {
		ID                 string
		Name               string
		Unit               string
		Source             *IDAndValue
		SourceNote         string
		SourceOrganization string
		Topics             []*IDAndValue
	}

	// IDAndValue represents ID and Value
	IDAndValue struct {
		ID    string
		Value string
	}
)

// List returns a Response's Summary and Indicators
func (i *IndicatorsService) List(pages *PageParams) (*PageSummary, []*Indicator, error) {
	summary := &PageSummary{}
	indicators := []*Indicator{}

	req, err := i.client.NewRequest("GET", "indicators", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicators}); err != nil {
		return nil, nil, err
	}

	return summary, indicators, nil
}

// Get returns a Response's Summary and an Indicator
func (i *IndicatorsService) Get(indicatorID string) (*PageSummary, *Indicator, error) {
	summary := &PageSummary{}
	indicator := []*Indicator{}

	path := fmt.Sprintf("indicators/%v", indicatorID)
	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicator}); err != nil {
		return nil, nil, err
	}

	return summary, indicator[0], nil
}

// ListByTopicID returns a Response's Summary and Indicators By topic id
func (i *IndicatorsService) ListByTopicID(topicID string, pages *PageParams) (*PageSummary, []*Indicator, error) {
	summary := &PageSummary{}
	indicators := []*Indicator{}

	path := fmt.Sprintf("topics/%v/indicators", topicID)
	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicators}); err != nil {
		return nil, nil, err
	}

	return summary, indicators, nil
}
