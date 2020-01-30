package wbdata

import (
	"time"
)

type SourcesService service

type Source struct {
	ID                   string
	Name                 string
	Code                 string
	Description          string
	URL                  string
	DataAvailability     string
	MetadataAvailability string
	Concepts             string
	LastUpdated          *time.Time
}
