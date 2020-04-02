package wbdata

import (
	"fmt"
)

type (
	IncomeLevelsService service

	IncomeLevel struct {
		ID       string
		Iso2Code string
		Value    string
	}
)

func (i *IncomeLevelsService) ListIncomeLevels(pages PageParams) (*PageSummary, *[]IncomeLevel, error) {
	summary := PageSummary{}
	incomeLevels := []IncomeLevel{}

	req, err := i.client.NewRequest("GET", "incomeLevels", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	_, err = i.client.do(req, &[]interface{}{&summary, &incomeLevels})
	if err != nil {
		return nil, nil, err
	}

	return &summary, &incomeLevels, err
}

func (i *IncomeLevelsService) GetIncomeLevel(incomeLevelID string) (*PageSummary, *IncomeLevel, error) {
	summary := PageSummary{}
	incomeLevels := []IncomeLevel{}

	s := fmt.Sprintf("incomeLevels/%v", incomeLevelID)
	req, err := i.client.NewRequest("GET", s, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = i.client.do(req, &[]interface{}{&summary, &incomeLevels})
	if err != nil {
		return nil, nil, err
	}

	return &summary, &incomeLevels[0], nil
}
