package wbdata

type CountriesService service

type Country struct {
	Name         string
	CapitalCity  string
	Iso2Code     string
	Longitude    string
	Latitude     string
	Region       Region
	IncomeLevels IncomeLevel
	LendingType  LendingType
	AdminRegion  struct {
		ID    string
		Value string
	}
}
