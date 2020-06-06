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
	"path/filepath"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

const (
	defaultProtocol = "http"
	defaultHost     = "api.worldbank.org"
	defaultBaseURL  = defaultProtocol + "://" + defaultHost + "/"
	apiVersion      = "v2"
	userAgent       = "wbdata-go"
	defaultFormat   = "json"
)

// A Client manages communication with the World Bank Open Data API
type Client struct {
	client *http.Client

	// Base URL for API requests. Defaults to the World Bank Open Data API
	BaseURL *url.URL

	// Local Language
	LocalLanguage string

	// Logger
	Logger *log.Logger

	// User agent used when communicating with the World Bank Open Data API
	UserAgent string

	// Services to talk to different APIs
	Countries    *CountriesService
	Indicators   *IndicatorsService
	IncomeLevels *IncomeLevelsService
	LendingTypes *LendingTypesService
	Regions      *RegionsService
	Sources      *SourcesService
	Topics       *TopicsService
}

type service struct {
	client *Client //nolint:structcheck
}

// LocalLanguage sets local language
func LocalLanguage(lang string) func(*Client) {
	return func(s *Client) {
		s.LocalLanguage = lang
	}
}

// NewClient returns a new World Bank Open Data API client.
func NewClient(httpClient *http.Client, options ...func(*Client)) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL + apiVersion + "/")
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	for _, option := range options {
		option(c)
	}
	c.Countries = &CountriesService{client: c}
	c.Sources = &SourcesService{client: c}
	c.Topics = &TopicsService{client: c}
	c.Indicators = &IndicatorsService{client: c}
	c.IncomeLevels = &IncomeLevelsService{client: c}
	c.LendingTypes = &LendingTypesService{client: c}
	c.Regions = &RegionsService{client: c}
	return c
}

// NewRequest returns a new World Bank Open Data API http request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	v := url.Values{}
	v.Set("format", defaultFormat)
	// Set Local language
	if c.LocalLanguage != `` {
		urlStr = fmt.Sprintf("%s/%s", c.LocalLanguage, urlStr)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse from %s: %v", urlStr, err)
	}
	url := fmt.Sprintf("%s?%s", u, v.Encode())

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, fmt.Errorf("failed to encode from %s: %v", url, err)
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

func (c *Client) do(req *http.Request, v *[]interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := checkStatusCode(resp); err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read all from %q: %v", req.URL, err)
	}

	var errReses []ErrorResponse
	if err := json.Unmarshal(data, &errReses); err != io.EOF && err != nil {
		if len(errReses) != 0 {
			errReses = []ErrorResponse{}
		}

		if err := json.Unmarshal(data, v); err != io.EOF && err != nil {
			return fmt.Errorf("failed to unmarshal from %q", req.URL)
		}
	}

	if len(errReses) != 0 {
		errReses[0].URL = req.URL.String()
		errReses[0].Code = resp.StatusCode
		return &errReses[0]
	}

	return nil
}

func checkStatusCode(resp *http.Response) error {
	// NOTE: StatusCode is 'always' 200 Eeven if ErrorMessage exists.
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	if c := resp.StatusCode; 500 <= c && c <= 599 {
		return NewAPIError(resp.Request.URL.String(), resp.StatusCode, ErrInvalidServer)
	}

	return nil
}

// NewTestClient returns a new World Bank Open Data API client for Test using go-vcr.
func NewTestClient(t testing.TB, update bool) (*Client, func()) {
	fixtureDir := filepath.Join("testdata", "fixtures")
	cassette := filepath.Join(fixtureDir, t.Name())

	r, err := recorder.New(cassette)
	if err != nil {
		t.Fatal(err)
	}
	customHTTPClient := &http.Client{
		Transport: r,
	}

	return NewClient(customHTTPClient), func() {
		if err := r.Stop(); err != nil {
			t.Errorf("failed to update fixtures: %s", err)
		}
	}
}
