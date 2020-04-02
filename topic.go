package wbdata

import "fmt"

type (
	TopicsService service

	Topic struct {
		ID         string
		Value      string
		SourceNote string
	}
)

func (c *TopicsService) ListTopics(pages PageParams) (*PageSummary, []*Topic, error) {
	summary := &PageSummary{}
	topics := []*Topic{}

	req, err := c.client.NewRequest("GET", "topics", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := pages.pageParams(req); err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &topics})
	if err != nil {
		return nil, nil, err
	}

	return summary, topics, nil
}

func (c *TopicsService) GetTopic(topicID string) (*PageSummary, *Topic, error) {
	summary := &PageSummary{}
	topic := []*Topic{}

	s := fmt.Sprintf("topics/%v", topicID)
	req, err := c.client.NewRequest("GET", s, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = c.client.do(req, &[]interface{}{summary, &topic})
	if err != nil {
		return nil, nil, err
	}

	return summary, topic[0], nil
}
