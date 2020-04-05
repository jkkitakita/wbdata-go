package wbdata

import (
	"fmt"
)

type (
	// SourcesService ...
	SourcesService service

	// Source contains information for a source field
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

// ListSources returns a Response's Summary and Sources
func (s *SourcesService) ListSources(pages PageParams) (*PageSummary, []*Source, error) {
	summary := &PageSummary{}
	sources := []*Source{}

	req, err := s.client.NewRequest("GET", "sources", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	if err = s.client.do(req, &[]interface{}{summary, &sources}); err != nil {
		return nil, nil, err
	}

	return summary, sources, nil
}

// GetSource returns a Response's Summary and a Source
func (s *SourcesService) GetSource(sourceID string) (*PageSummary, *Source, error) {
	summary := &PageSummary{}
	source := []*Source{}

	path := fmt.Sprintf("sources/%v", sourceID)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = s.client.do(req, &[]interface{}{summary, &source}); err != nil {
		return nil, nil, err
	}

	return summary, source[0], nil
}
