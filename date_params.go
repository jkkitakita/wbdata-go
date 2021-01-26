package wbdata

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	// FilterParamsUnknown is type of filter params for unknown
	FilterParamsUnknown FilterParamsType = iota
	// FilterParamsDate is type of filter params for date
	FilterParamsDate
	// FilterParamsDateRange is type of filter params for date range
	FilterParamsDateRange
	// FilterParamsYearToDate is type of filter params for year-to-date
	FilterParamsYearToDate
	// FilterParamsMRV is type of filter params for most recent values
	FilterParamsMRV

	// FrequencyUnknown is frequency type of unknown
	FrequencyUnknown FrequencyType = iota
	// FrequencyMonthly is frequency type of monthly (M)
	FrequencyMonthly
	// FrequencyQuarterly is frequency type of quarterly (Q)
	FrequencyQuarterly
	// FrequencyYearly is frequency type of yearly (Y)
	FrequencyYearly
)

type (
	// FilterParamsType is type of filter params
	FilterParamsType uint
	// FrequencyType is type of frequency
	FrequencyType uint

	// FilterParams is struct for filter params
	FilterParams struct {
		FilterParamsType FilterParamsType
		DateParam        *DateParam
		RecentParam      *RecentParam
	}

	// DateParam is struct for date params
	DateParam struct {
		Date      string
		DateRange *DateRange
	}

	// RecentParam is struct for recent params
	RecentParam struct {
		FrequencyType    FrequencyType
		MostRecentValues uint
		IsNotEmpty       bool
		IsGapFill        bool
	}

	// DateRange is a struct for API's query params about date
	DateRange struct {
		Start string
		End   string
	}
)

func (dp *FilterParams) addFilterParams(req *http.Request) error {
	if dp == nil {
		return nil
	}

	if err := dp.validate(); err != nil {
		return err
	}

	switch dp.FilterParamsType {
	case FilterParamsDate, FilterParamsDateRange, FilterParamsYearToDate:
		dp.addDateParams(req)
	case FilterParamsMRV:
		dp.addRecentParam(req)
	}

	return nil
}

func (dp *FilterParams) addDateParams(req *http.Request) {
	params := req.URL.Query()
	params.Set(`date`, dp.buildDateParams())
	req.URL.RawQuery = params.Encode()
}

func (dp *FilterParams) addRecentParam(req *http.Request) {
	params := req.URL.Query()

	if dp.RecentParam.IsNotEmpty {
		params.Set(`mrnev`, fmt.Sprint(dp.RecentParam.MostRecentValues))
	} else {
		params.Set(`mrv`, fmt.Sprint(dp.RecentParam.MostRecentValues))
	}

	if dp.RecentParam.IsGapFill {
		params.Set(`gapfill`, `Y`)
	}

	switch dp.RecentParam.FrequencyType {
	case FrequencyMonthly:
		params.Set(`frequency`, `M`)
	case FrequencyQuarterly:
		params.Set(`frequency`, `Q`)
	case FrequencyYearly:
		params.Set(`frequency`, `Y`)
	}

	req.URL.RawQuery = params.Encode()
}

func (dp *FilterParams) validate() error {
	switch dp.FilterParamsType {
	case FilterParamsDate:
		if err := dp.validateFilterParamsDate(); err != nil {
			return err
		}
	case FilterParamsDateRange:
		if err := dp.validateFilterParamsDateRange(); err != nil {
			return err
		}
	case FilterParamsYearToDate:
		if err := dp.validateFilterParamsYearToDate(); err != nil {
			return err
		}
	case FilterParamsMRV:
		if err := dp.validateFilterParamsMRV(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("filter params type is unknown. filter params: %v", dp)
	}

	return nil
}

func (dp *FilterParams) validateFilterParamsDate() error {
	if dp.RecentParam != nil {
		return fmt.Errorf("recent param should not be specified with FilterParamsDate. filter params: %v", dp)
	}
	if dp.DateParam.DateRange != nil {
		return fmt.Errorf("date range should not be specified with FilterParamsDate. filter params: %v", dp)
	}
	if dp.DateParam.Date == "" {
		return fmt.Errorf("date is required. filter params: %v", dp)
	}
	_, err := parseDate(dp.DateParam.Date)
	if err != nil {
		return err
	}

	return nil
}

func (dp *FilterParams) validateFilterParamsDateRange() error {
	if dp.RecentParam != nil {
		return fmt.Errorf("recent param should not be specified with FilterParamsDateRange. filter params: %v", dp)
	}
	if dp.DateParam.Date != "" {
		return fmt.Errorf("date should not be specified with FilterParamsDateRange. filter params: %v", dp)
	}
	if dp.DateParam.DateRange == nil {
		return fmt.Errorf("dateRange is required. filter params: %v", dp)
	}
	if err := dp.DateParam.DateRange.validate(); err != nil {
		return err
	}

	return nil
}

func (dp *FilterParams) validateFilterParamsYearToDate() error {
	if dp.RecentParam != nil {
		return fmt.Errorf("recent param should not be specified with FilterParamsYearToDate. filter params: %v", dp)
	}
	if dp.DateParam.DateRange != nil {
		return fmt.Errorf("date range should not be specified with FilterParamsYearToDate. filter params: %v", dp)
	}
	if dp.DateParam.Date == "" {
		return fmt.Errorf("date is required. filter params: %v", dp)
	}
	_, err := time.Parse("2006", dp.DateParam.Date)
	if err != nil {
		return err
	}

	return nil
}

func (dp *FilterParams) validateFilterParamsMRV() error {
	if dp.DateParam != nil {
		return fmt.Errorf("date param should not be specified with FilterParamsMRV. filter params: %v", dp)
	}
	if dp.RecentParam == nil {
		return fmt.Errorf("recent param is required. filter params: %v", dp)
	}

	switch dp.RecentParam.FrequencyType {
	case FrequencyMonthly, FrequencyQuarterly, FrequencyYearly:
	default:
		return fmt.Errorf("frequency type should be quarterly (Q), monthly (M) or yearly (Y). filter params: %v", dp)
	}

	if dp.RecentParam.MostRecentValues == 0 {
		return fmt.Errorf("most recent values should be larger then 0. filter params: %v", dp)
	}
	if dp.RecentParam.IsGapFill && dp.RecentParam.IsNotEmpty {
		return fmt.Errorf("IsGapFill cannot be true when IsNotEmpty is true. filter params: %v", dp)
	}

	return nil
}

func (dp *FilterParams) buildDateParams() string {
	switch dp.FilterParamsType {
	case FilterParamsDate:
		return dp.DateParam.Date
	case FilterParamsDateRange:
		return dp.DateParam.DateRange.join()
	case FilterParamsYearToDate:
		return "YTD:" + dp.DateParam.Date
	default:
		fmt.Printf("filter params type is invalid. filter params: %v", dp)
		return ""
	}
}

func (dp *DateRange) validate() error {
	startDate, err := parseDate(dp.Start)
	if err != nil {
		return err
	}
	endDate, err := parseDate(dp.End)
	if err != nil {
		return err
	}

	if err := dp.validateMonthlyAndQuarterly(); err != nil {
		return err
	}

	if startDate.After(endDate) {
		return fmt.Errorf("start should be before end, start: %v end: %v", dp.Start, dp.End)
	}

	return nil
}

func (dp *DateRange) validateMonthlyAndQuarterly() error {
	if dp.isYearly() || dp.isMonthly() || dp.isQuarterly() {
		return nil
	}

	return fmt.Errorf("yearly and quarterly and monthly cannot be used together, start: %v end: %v", dp.Start, dp.End)
}

func (dp *DateRange) isYearly() bool {
	return !strings.Contains(dp.Start, "M") && !strings.Contains(dp.End, "M") &&
		!strings.Contains(dp.End, "Q") && !strings.Contains(dp.End, "Q")
}

func (dp *DateRange) isMonthly() bool {
	return strings.Contains(dp.Start, "M") && strings.Contains(dp.End, "M")
}

func (dp *DateRange) isQuarterly() bool {
	return strings.Contains(dp.Start, "Q") && strings.Contains(dp.End, "Q")
}

func (dp *DateRange) join() string {
	return dp.Start + ":" + dp.End
}

func parseDate(dateStr string) (time.Time, error) {
	t, err := time.Parse("2006", dateStr)
	if err != nil {
		t, err = time.Parse("2006M01", dateStr)
		if err != nil {
			t, err = time.Parse("2006Q01", dateStr)
			if err != nil {
				return time.Time{}, err
			}
		}
	}

	return t, nil
}
