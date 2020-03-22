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
	defaultBaseURL = "http://api.worldbank.org/"
	apiVersion     = "v2"
	userAgent      = "wbdata"
	defaultFormat  = "json"
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
	baseURL, _ := url.Parse(defaultBaseURL + apiVersion + "/")
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
	v := url.Values{}
	v.Set("format", defaultFormat)
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?%s", u, v.Encode())

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
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

func (c *Client) do(req *http.Request, v *[]interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// log.Printf(`req: %+v`, req)

	if err := checkStatusCode(resp); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var errReses []ErrorResponse
	if err := json.Unmarshal(data, &errReses); err != io.EOF && err != nil {
		// intialize []ErrorResponse
		if len(errReses) != 0 {
			errReses = []ErrorResponse{}
		}

		if err := json.Unmarshal(data, v); err != io.EOF && err != nil {
			return nil, err
		}
	}

	if len(errReses) != 0 {
		return nil, &errReses[0]
	}

	return resp, nil
}

func checkStatusCode(resp *http.Response) error {
	// NOTE: StatusCode is 'always' 200 Eeven if ErrorMessage exists.
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	if c := resp.StatusCode; 500 <= c && c <= 599 {
		return &APIError{
			Status:       resp.StatusCode,
			ErrorMessage: ErrInvalidServer,
		}
	}

	return nil
}
