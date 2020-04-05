package wbdata

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type (
	intOrString int

	// PageParams is a struct for API's query params about pages
	PageParams struct {
		Page    int
		PerPage int
	}

	// PageSummary is a struct for a Summary about pages
	PageSummary struct {
		Page    intOrString `json:"page"`
		Pages   intOrString `json:"pages"`
		PerPage intOrString `json:"per_page"`
		Total   intOrString `json:"total"`
	}
)

func (pages *PageParams) pageParams(req *http.Request) error {
	params := req.URL.Query()

	if pages.Page > 0 {
		params.Add(`page`, strconv.Itoa(pages.Page))
	} else {
		return errors.New("page of params should be larger than 0")
	}

	if pages.PerPage > 0 {
		params.Add(`per_page`, strconv.Itoa(pages.PerPage))
	} else {
		return errors.New("per_page of params should be larger than 0")
	}

	req.URL.RawQuery = params.Encode()

	return nil
}

func (ios *intOrString) UnmarshalJSON(data []byte) error {
	var intRegex = regexp.MustCompile(`\d+`)
	trimData := strings.Trim(string(data), "\"")
	if intRegex.MatchString(trimData) {
		if ios != nil {
			intIos, err := strconv.Atoi(trimData)
			if err != nil {
				return err
			}
			*ios = intOrString(intIos)
		}
		return nil
	}

	var i int
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}
	p := (*int)(ios)
	*p = i
	return nil
}
