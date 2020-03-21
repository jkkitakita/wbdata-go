package wbdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.worldbank.org/"
	apiVersion     = "v2"
	userAgent      = "wbdata"
)

// A Client manages communication with the World Bank Open Data API.
type Client struct {
	client *http.Client

	// Base URL for API requests. Defaults to the World Bank Open Data API
	BaseURL *url.URL

	// Logger
	Logger *log.Logger

	// User agent used when communicating with the GitHub API.
	UserAgent string

	//Services to talk to different APIs
	Countries    *CountriesService
	Indicators   *IndicatorsService
	IncomeLevels *IncomeLevelsService
	LendingTypes *LendingTypesService
	Regions      *RegionsService
	Sources      *SourcesService
	Topics       *TopicsService
}

type service struct {
	client *Client
}

// NewClient returns a new World Bank Open Data API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.Countries = &CountriesService{client: c}
	c.Sources = &SourcesService{client: c}
	c.Topics = &TopicsService{client: c}
	c.Indicators = &IndicatorsService{client: c}
	c.IncomeLevels = &IncomeLevelsService{client: c}
	c.LendingTypes = &LendingTypesService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String()+apiVersion, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}


func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := checkRes(resp); err != nil {
		return resp, err
	}

	return resp, err
}

type ErrorResponse struct {
	Message []struct {
		Id    string
		Key   string
		Value string
	}
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%+v", r.Message)
}

func checkRes(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errRes := &ErrorResponse{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil && data != nil {
		json.Unmarshal(data, errRes)
	}
	return errRes
}