package wbdata

type IndicatorsService service

type Indicator struct {
	ID                 string
	Name               string
	Source             *Source
	SourceNote         string
	SourceOrganization string
	Topicss            []*Topic
}
