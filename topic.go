package wbdata

import "fmt"

type (
	// TopicsService ...
	TopicsService service

	// Topic contains information for an topic field
	Topic struct {
		ID         string
		Value      string
		SourceNote string
	}
)

// ListTopics returns a Response's Summary and Topics
func (t *TopicsService) ListTopics(pages PageParams) (*PageSummary, []*Topic, error) {
	summary := &PageSummary{}
	topics := []*Topic{}

	req, err := t.client.NewRequest("GET", "topics", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	if err = t.client.do(req, &[]interface{}{summary, &topics}); err != nil {
		return nil, nil, err
	}

	return summary, topics, nil
}

// GetTopic returns a Response's Summary and a Topic
func (t *TopicsService) GetTopic(topicID string) (*PageSummary, *Topic, error) {
	summary := &PageSummary{}
	topic := []*Topic{}

	path := fmt.Sprintf("topics/%v", topicID)
	req, err := t.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	if err = t.client.do(req, &[]interface{}{summary, &topic}); err != nil {
		return nil, nil, err
	}

	return summary, topic[0], nil
}
