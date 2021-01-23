package wbdata

import (
	"errors"
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
		return errors.New(`Start should be before End`)
	}

	return nil
}

func validateYear(year string) (time.Time, error) {
	t, err := time.Parse("2006", year)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func (dp *DateParams) join(separator string) string {
	return dp.Start + ":" + dp.End
}

func (dp *DateParams) addDateParams(req *http.Request) error {
	if err := dp.validate(); err != nil {
		return err
	}

	params := req.URL.Query()
	params.Add(`date`, dp.join(":"))
	req.URL.RawQuery = params.Encode()

	return nil
}
