package wbdata

const (
	// OutputFormatJSON is output format for json
	OutputFormatJSON OutputFormat = "json"
	// OutputFormatJSONP is output format for jsonP
	OutputFormatJSONP = "jsonP"
	// OutputFormatJSONStat is output format for json-stat
	OutputFormatJSONStat = "jsonstat"
	// OutputFormatXML is output format for xml
	OutputFormatXML = "xml"
)

// OutputFormat is output format
type OutputFormat string

func (o OutputFormat) String() string {
	return string(o)
}
