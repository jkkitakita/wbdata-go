package wbdata

type (
	IndicatorsService service

	Indicator struct {
		ID                 string
		Name               string
		Source             *Source
		SourceNote         string
		SourceOrganization string
		Topicss            []*Topic
	}
)
