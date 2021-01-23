package wbdata

import (
	"fmt"
)

type (
	// LanguagesService ...
	LanguagesService service

	// Language contains information for an language field
	Language struct {
		Code       string
		Name       string
		NativeForm string
	}
)

// List returns summary and languages
func (c *LanguagesService) List(pages *PageParams) (*PageSummary, []*Language, error) {
	summary := &PageSummary{}
	languages := []*Language{}

	req, err := c.client.NewRequest("GET", "languages", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.addPageParams(req); err != nil {
		return nil, nil, err
	}

	if err = c.client.do(req, &[]interface{}{summary, &languages}); err != nil {
		return nil, nil, err
	}

	return summary, languages, nil
}

// Get returns summary and a language
func (c *LanguagesService) Get(languageCode string) (*PageSummary, *Language, error) {
	summary := &PageSummary{}
	language := []*Language{}

	path := fmt.Sprintf("languages/%v", languageCode)
	req, err := c.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = c.client.do(req, &[]interface{}{summary, &language}); err != nil {
		return nil, nil, err
	}

	return summary, language[0], nil
}
