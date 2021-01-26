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

	// IndicatorValueWithFootnote represents an indicator value with footnote
	IndicatorValueWithFootnote struct {
		IndicatorValue
		Footnote string `json:"footnote"`
	}
)

// List returns a Response's Summary and Indicator in all countries
func (i *IndicatorValuesService) List(
	indicatorID string,
	filterParams *FilterParams,
	pages *PageParams,
) (*PageSummaryWithSourceID, []*IndicatorValue, error) {
	summary := &PageSummaryWithSourceID{}
	indicatorValues := []*IndicatorValue{}

	path := fmt.Sprintf(
		"countries/all/indicators/%s",
		indicatorID,
	)

	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err := filterParams.addFilterParams(req); err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}

// ListWithFootnote returns a Response's Summary and Indicator with footnote in all countries
func (i *IndicatorValuesService) ListWithFootnote(
	indicatorID string,
	filterParams *FilterParams,
	pages *PageParams,
) (*PageSummaryWithSourceID, []*IndicatorValueWithFootnote, error) {
	summary := &PageSummaryWithSourceID{}
	indicatorValues := []*IndicatorValueWithFootnote{}

	path := fmt.Sprintf(
		"countries/all/indicators/%s",
		indicatorID,
	)

	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err := filterParams.addFilterParams(req); err != nil {
		return nil, nil, err
	}

	addFootNoteParams(req)

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}

// ListByCountryIDs returns a Response's Summary and Indicator By country IDs
func (i *IndicatorValuesService) ListByCountryIDs(
	countryIDs []string,
	indicatorID string,
	filterParams *FilterParams,
	pages *PageParams,
) (*PageSummaryWithSourceID, []*IndicatorValue, error) {
	summary := &PageSummaryWithSourceID{}
	indicatorValues := []*IndicatorValue{}

	path := fmt.Sprintf(
		"countries/%s/indicators/%s",
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

	if err := filterParams.addFilterParams(req); err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}

// ListByCountryIDsWithFootnote returns a Response's Summary and Indicator with footnote By country IDs
func (i *IndicatorValuesService) ListByCountryIDsWithFootnote(
	countryIDs []string,
	indicatorID string,
	filterParams *FilterParams,
	pages *PageParams,
) (*PageSummaryWithSourceID, []*IndicatorValueWithFootnote, error) {
	summary := &PageSummaryWithSourceID{}
	indicatorValues := []*IndicatorValueWithFootnote{}

	path := fmt.Sprintf(
		"countries/%s/indicators/%s",
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

	if err := filterParams.addFilterParams(req); err != nil {
		return nil, nil, err
	}

	addFootNoteParams(req)

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}

// ListBySourceID returns a Response's Summary and Indicator in all countries By source ID
func (i *IndicatorValuesService) ListBySourceID(
	indicatorIDs []string,
	sourceID string,
	filterParams *FilterParams,
	pages *PageParams,
) (*PageSummaryWithLastUpdated, []*IndicatorValue, error) {
	summary := &PageSummaryWithLastUpdated{}
	indicatorValues := []*IndicatorValue{}

	path := fmt.Sprintf(
		"countries/all/indicators/%s?source=%s",
		strings.Join(indicatorIDs, ";"),
		sourceID,
	)

	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err := filterParams.addFilterParams(req); err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}

// ListBySourceIDWithFootnote returns a Response's Summary and Indicator with footnote in all countries By source ID
func (i *IndicatorValuesService) ListBySourceIDWithFootnote(
	indicatorIDs []string,
	sourceID string,
	filterParams *FilterParams,
	pages *PageParams,
) (*PageSummaryWithLastUpdated, []*IndicatorValueWithFootnote, error) {
	summary := &PageSummaryWithLastUpdated{}
	indicatorValues := []*IndicatorValueWithFootnote{}

	path := fmt.Sprintf(
		"countries/all/indicators/%s?source=%s",
		strings.Join(indicatorIDs, ";"),
		sourceID,
	)

	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err := filterParams.addFilterParams(req); err != nil {
		return nil, nil, err
	}

	addFootNoteParams(req)

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}

// ListByCountryIDsAndSourceID returns a Response's Summary and Indicator By country IDs and source ID
func (i *IndicatorValuesService) ListByCountryIDsAndSourceID(
	countryIDs []string,
	indicatorIDs []string,
	sourceID string,
	filterParams *FilterParams,
	pages *PageParams,
) (*PageSummaryWithLastUpdated, []*IndicatorValue, error) {
	summary := &PageSummaryWithLastUpdated{}
	indicatorValues := []*IndicatorValue{}

	path := fmt.Sprintf(
		"countries/%s/indicators/%s?source=%s",
		strings.Join(countryIDs, ";"),
		strings.Join(indicatorIDs, ";"),
		sourceID,
	)

	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err := filterParams.addFilterParams(req); err != nil {
		return nil, nil, err
	}

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}

// ListByCountryIDsAndSourceIDWithFootnote returns a Response's Summary and Indicator with footnote By country IDs and source ID
func (i *IndicatorValuesService) ListByCountryIDsAndSourceIDWithFootnote(
	countryIDs []string,
	indicatorIDs []string,
	sourceID string,
	filterParams *FilterParams,
	pages *PageParams,
) (*PageSummaryWithLastUpdated, []*IndicatorValueWithFootnote, error) {
	summary := &PageSummaryWithLastUpdated{}
	indicatorValues := []*IndicatorValueWithFootnote{}

	path := fmt.Sprintf(
		"countries/%s/indicators/%s?source=%s",
		strings.Join(countryIDs, ";"),
		strings.Join(indicatorIDs, ";"),
		sourceID,
	)

	req, err := i.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err := filterParams.addFilterParams(req); err != nil {
		return nil, nil, err
	}

	addFootNoteParams(req)

	if err = i.client.do(req, &[]interface{}{summary, &indicatorValues}); err != nil {
		return nil, nil, err
	}

	return summary, indicatorValues, nil
}
