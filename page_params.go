package wbdata

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

type intOrString int

type PageParams struct {
	Page    int
	PerPage int
}

type PageSummary struct {
	Page    intOrString `json:"page"`
	Pages   intOrString `json:"pages"`
	PerPage intOrString `json:"per_page"`
	Total   intOrString `json:"total"`
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
