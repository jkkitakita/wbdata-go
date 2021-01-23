package wbdata

import (
	"fmt"
	"net/http"
	"time"
)

type (
	// DateParams is a struct for API's query params about date
	DateParams struct {
		Start string
		End   string
	}
)

func (dp *DateParams) validate() error {
	startYear, err := validateYear(dp.Start)
	if err != nil {
		return err
	}
	endYear, err := validateYear(dp.End)
	if err != nil {
		return err
	}

	if startYear.After(endYear) {
		return fmt.Errorf("start should be before end, start: %v end: %v", startYear, endYear)
	}

	return nil
}

func validateYear(yearStr string) (time.Time, error) {
	t, err := time.Parse("2006", yearStr)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func (dp *DateParams) join() string {
	return dp.Start + ":" + dp.End
}

func (dp *DateParams) addDateParams(req *http.Request) error {
	if dp == nil {
		return nil
	}

	if err := dp.validate(); err != nil {
		return err
	}

	params := req.URL.Query()
	params.Add(`date`, dp.join())
	req.URL.RawQuery = params.Encode()

	return nil
}
