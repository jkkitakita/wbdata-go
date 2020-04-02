package wbdata

import (
	"fmt"
)

type (
	SourcesService service

	Source struct {
		ID                   string
		LastUpdated          string
		Name                 string
		Code                 string
		Description          string
		URL                  string
		DataAvailability     string
		MetadataAvailability string
		Concepts             string
	}
)

func (c *SourcesService) ListSources(pages PageParams) (*PageSummary, []*Source, error) {
	summary := &PageSummary{}
	sources := []*Source{}

	req, err := c.client.NewRequest("GET", "sources", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &sources})
	if err != nil {
		return nil, nil, err
	}

	return summary, sources, nil
}

func (c *SourcesService) GetSource(sourceID string) (*PageSummary, *Source, error) {
	summary := &PageSummary{}
	source := []*Source{}

	s := fmt.Sprintf("sources/%v", sourceID)
	req, err := c.client.NewRequest("GET", s, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &source})
	if err != nil {
		return nil, nil, err
	}

	return summary, source[0], nil
}
