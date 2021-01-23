package wbdata

import (
	"fmt"
	"strings"
)

type (
	// IndicatorValuesService ...
	IndicatorValuesService service

	// IndicatorValue represents an indicator value
	IndicatorValue struct {
		Indicator       IDAndValue `json:"indicator"`
		Country         IDAndValue `json:"country"`
		Countryiso3code string     `json:"countryiso3code"`
		Date            string     `json:"date"`
		Value           float64    `json:"value"`
		Unit            string     `json:"unit"`
		ObsStatus       string     `json:"obs_status"`
		Decimal         int32      `json:"decimal"`
	}
)

// ListByCountryIDs returns a Response's Summary and Indicator By country IDs
func (i *IndicatorValuesService) ListByCountryIDs(
	countryIDs []string,
	indicatorID string,
	datePatams *DateParams,
	pages *PageParams,
) (*PageSummaryWithSource, []*IndicatorValue, error) {
	summary := &PageSummaryWithSource{}
	indicatorValues := []*IndicatorValue{}

	path := fmt.Sprintf(
		"countries/%v/indicators/%v",
		strings.Join(countryIDs, ";"),
		indicatorID,
	)

	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err := datePatams.addDateParams(req); err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}
