package wbdata

import (
	"fmt"
	"strconv"
)

type IncomeLevelsService service

type IncomeLevel struct {
	ID       string
	Iso2Code string
	Value    string
}

func (i *IncomeLevelsService) ListIncomeLevels(pages PageParams) (PageSummary, []IncomeLevel, error) {
	summary := PageSummary{}
	incomeLevels := []IncomeLevel{}

	req, err := i.client.NewRequest("GET", "incomeLevels", nil)
	if err != nil {
		return PageSummary{}, nil, err
	}

	params := req.URL.Query()
	if pages.Page != 0 {
		params.Add(`page`, strconv.Itoa(int(pages.Page)))
	}
	if pages.PerPage != 0 {
		params.Add(`per_page`, strconv.Itoa(int(pages.PerPage)))
	}
	req.URL.RawQuery = params.Encode()

	_, err = i.client.do(req, &[]interface{}{&summary, &incomeLevels})
	if err != nil {
		return PageSummary{}, nil, err
	}

	return summary, incomeLevels, err
}

func (i *IncomeLevelsService) GetIncomeLevel(incomeLevelID string) (PageSummary, IncomeLevel, error) {
	summary := PageSummary{}
	incomeLevels := []IncomeLevel{}

	s := fmt.Sprintf("incomeLevels/%v", incomeLevelID)
	req, err := i.client.NewRequest("GET", s, nil)
	if err != nil {
		return PageSummary{}, IncomeLevel{}, err
	}

	_, err = i.client.do(req, &[]interface{}{&summary, &incomeLevels})
	if err != nil {
		return PageSummary{}, IncomeLevel{}, err
	}

	return summary, incomeLevels[0], nil
}
