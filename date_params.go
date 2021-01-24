package wbdata

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	// DateParamsUnknown is type of date params for unknown
	DateParamsUnknown DateParamsType = iota
	// DateParamsDate is type of date params for date
	DateParamsDate
	// DateParamsRange is type of date params for date range
	DateParamsRange
	// DateParamsYearToDate is type of date params for year-to-date
	DateParamsYearToDate
)

type (
	// DateParamsType is type of date params
	DateParamsType uint

	// DateParams is struct for date params
	DateParams struct {
		DateParamsType DateParamsType
		Date           string
		DateRange      *DateRange
	}

	// DateRange is a struct for API's query params about date
	DateRange struct {
		Start string
		End   string
	}
)

func (dp *DateParams) addDateParams(req *http.Request) error {
	if dp == nil {
		return nil
	}

	if err := dp.validate(); err != nil {
		return err
	}

	params := req.URL.Query()
	params.Add(`date`, dp.buildDateParams())
	req.URL.RawQuery = params.Encode()

	return nil
}

func (dp *DateParams) validate() error {
	switch dp.DateParamsType {
	case DateParamsDate:
		if dp.DateRange != nil {
			return fmt.Errorf("date range should not be specified with DateParamsDate. date params: %v", dp)
		}
		if dp.Date == "" {
			return fmt.Errorf("date is required. date param: %v", dp)
		}
		_, err := parseDate(dp.Date)
		if err != nil {
			return err
		}
	case DateParamsRange:
		if dp.Date != "" {
			return fmt.Errorf("date should not be specified with DateParamsRange. date params: %v", dp)
		}
		if dp.DateRange == nil {
			return fmt.Errorf("dateRange is required. date params: %v", dp)
		}
		if err := dp.DateRange.validate(); err != nil {
			return err
		}
	case DateParamsYearToDate:
		if dp.DateRange != nil {
			return fmt.Errorf("date range should not be specified with DateParamsYearToDate. date params: %v", dp)
		}
		if dp.Date == "" {
			return fmt.Errorf("date is required. date params: %v", dp)
		}
		_, err := time.Parse("2006", dp.Date)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("date params type is unknown. date params: %v", dp)
	}

	return nil
}

func (dp *DateParams) buildDateParams() string {
	switch dp.DateParamsType {
	case DateParamsDate:
		return dp.Date
	case DateParamsRange:
		return dp.DateRange.join()
	case DateParamsYearToDate:
		return "YTD:" + dp.Date
	default:
		fmt.Printf("date params type is invalid. date params: %v", dp)
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
	if dp.isYearlyRange() || dp.isMonthlyRange() || dp.isQuarterlyRange() {
		return nil
	}

	return fmt.Errorf("yearly and quarterly and monthly cannot be used together, start: %v end: %v", dp.Start, dp.End)
}

func (dp *DateRange) isYearlyRange() bool {
	return !strings.Contains(dp.Start, "M") && !strings.Contains(dp.End, "M") &&
		!strings.Contains(dp.End, "Q") && !strings.Contains(dp.End, "Q")
}

func (dp *DateRange) isMonthlyRange() bool {
	return strings.Contains(dp.Start, "M") && strings.Contains(dp.End, "M")
}

func (dp *DateRange) isQuarterlyRange() bool {
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
