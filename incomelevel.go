package wbdata

import (
	"fmt"
)

type (
	// IncomeLevelsService ...
	IncomeLevelsService service

	// IncomeLevel contains information for an incomelevel field
	IncomeLevel struct {
		ID       string
		Iso2Code string
		Value    string
	}
)

// ListIncomeLevels returns a Response's Summary and IncomeLevels
func (il *IncomeLevelsService) ListIncomeLevels(pages PageParams) (*PageSummary, []*IncomeLevel, error) {
	summary := &PageSummary{}
	incomeLevels := []*IncomeLevel{}

	req, err := il.client.NewRequest("GET", "incomeLevels", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	if err = il.client.do(req, &[]interface{}{summary, &incomeLevels}); err != nil {
		return nil, nil, err
	}

	return summary, incomeLevels, err
}

// GetIncomeLevel returns a Response's Summary and an IncomeLevel
func (il *IncomeLevelsService) GetIncomeLevel(incomeLevelID string) (*PageSummary, *IncomeLevel, error) {
	summary := &PageSummary{}
	incomeLevels := []*IncomeLevel{}

	path := fmt.Sprintf("incomeLevels/%v", incomeLevelID)
	req, err := il.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = il.client.do(req, &[]interface{}{summary, &incomeLevels}); err != nil {
		return nil, nil, err
	}

	return summary, incomeLevels[0], nil
}
