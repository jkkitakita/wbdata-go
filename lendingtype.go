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

func (lt *LendingTypesService) ListLendingTypes(pages PageParams) (*PageSummary, []*LendingType, error) {
	summary := &PageSummary{}
	lendingTypes := []*LendingType{}

	req, err := lt.client.NewRequest("GET", "lendingTypes", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	if err = lt.client.do(req, &[]interface{}{summary, &lendingTypes}); err != nil {
		return nil, nil, err
	}

	return summary, lendingTypes, nil
}

func (lt *LendingTypesService) GetLendingType(lendingTypeID string) (*PageSummary, *LendingType, error) {
	summary := &PageSummary{}
	lendingType := []*LendingType{}

	path := fmt.Sprintf("lendingTypes/%v", lendingTypeID)
	req, err := lt.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = lt.client.do(req, &[]interface{}{summary, &lendingType}); err != nil {
		return nil, nil, err
	}

	return summary, lendingType[0], nil
}
