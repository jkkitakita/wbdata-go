package main

import (
	"fmt"

	"github.com/jkkitakita/wbdata-go"
)

func main() {
	client := wbdata.NewClient(nil)

	// ListTopics
	summary, topics, err := client.Topics.ListTopics(wbdata.PageParams{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		fmt.Printf("failed to list topics: %+v\n", err)
	}
	fmt.Printf("ListTopics, summary: %+v, topics[0]: %+v\n", summary, topics[0])

	// GetTopic
	_, topic, err := client.Topics.GetTopic("1")
	if err != nil {
		fmt.Printf("failed to get topic: %+v\n", err)
	}
	fmt.Printf("GetTopic, topic: %+v\n", topic)
}
