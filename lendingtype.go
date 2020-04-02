package wbdata

import "fmt"

type (
	LendingTypesService service

	LendingType struct {
		ID       string
		Iso2Code string
		Value    string
	}
)

func (c *LendingTypesService) ListLendingTypes(pages PageParams) (*PageSummary, []*LendingType, error) {
	summary := &PageSummary{}
	lendingTypes := []*LendingType{}

	req, err := c.client.NewRequest("GET", "lendingTypes", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &lendingTypes})
	if err != nil {
		return nil, nil, err
	}

	return summary, lendingTypes, nil
}

func (c *LendingTypesService) GetLendingType(lendingTypeID string) (*PageSummary, *LendingType, error) {
	summary := &PageSummary{}
	lendingType := []*LendingType{}

	s := fmt.Sprintf("lendingTypes/%v", lendingTypeID)
	req, err := c.client.NewRequest("GET", s, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &lendingType})
	if err != nil {
		return nil, nil, err
	}

	return summary, lendingType[0], nil
}
