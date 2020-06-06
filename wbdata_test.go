package wbdata

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestNewClient(t *testing.T) {
	baseURL, _ := url.Parse(defaultBaseURL + apiVersion + "/")
	defaultClient := &Client{
		client:        &http.Client{},
		BaseURL:       baseURL,
		LocalLanguage: "",
		Logger:        nil,
		UserAgent:     userAgent,
	}
	jaClient := &Client{
		client:        &http.Client{},
		BaseURL:       baseURL,
		LocalLanguage: testutils.JaLocalLanguage,
		Logger:        nil,
		UserAgent:     userAgent,
	}
	optIgnoreUnexported := cmpopts.IgnoreUnexported(Client{})
	optIgnoreFields := cmpopts.IgnoreFields(Client{},
		"Countries",
		"Indicators",
		"IncomeLevels",
		"LendingTypes",
		"Regions",
		"Sources",
		"Topics",
	)

	type args struct {
		httpClient *http.Client
		options    []func(*Client)
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "default",
			args: args{
				httpClient: &http.Client{},
				options:    nil,
			},
			want: defaultClient,
		},
		{
			name: "ja",
			args: args{
				httpClient: &http.Client{},
				options: []func(*Client){
					LocalLanguage(testutils.JaLocalLanguage),
				},
			},
			want: jaClient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.httpClient, tt.args.options...)
			if !cmp.Equal(got, tt.want, optIgnoreUnexported, optIgnoreFields) {
				t.Errorf("NewClient() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	baseURL, _ := url.Parse(fmt.Sprintf("%s%s/", defaultBaseURL, apiVersion))
	urlStr := "countries"
	defaultRequestURL, _ := url.Parse(fmt.Sprintf("%s%s?format=%s", baseURL, urlStr, defaultFormat))
	jaRequestURL, _ := url.Parse(fmt.Sprintf("%s%s/%s?format=%s", baseURL, testutils.JaLocalLanguage, urlStr, defaultFormat))

	defaultClient := &Client{
		client:        &http.Client{},
		BaseURL:       baseURL,
		LocalLanguage: "",
		Logger:        nil,
		UserAgent:     userAgent,
	}
	jaClient := &Client{
		client:        &http.Client{},
		BaseURL:       baseURL,
		LocalLanguage: testutils.JaLocalLanguage,
		Logger:        nil,
		UserAgent:     userAgent,
	}
	defaultHTTPRequest := &http.Request{
		Method:     "GET",
		URL:        defaultRequestURL,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: map[string][]string{
			"User-Agent": {
				userAgent,
			},
		},
		Host: defaultHost,
	}
	jaHTTPRequest := &http.Request{
		Method:     "GET",
		URL:        jaRequestURL,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: map[string][]string{
			"User-Agent": {
				userAgent,
			},
		},
		Host: defaultHost,
	}

	type args struct {
		method string
		urlStr string
		body   interface{}
	}
	tests := []struct {
		name    string
		client  *Client
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name:   "default",
			client: defaultClient,
			args: args{
				method: "GET",
				urlStr: urlStr,
				body:   nil,
			},
			want:    defaultHTTPRequest,
			wantErr: false,
		},
		{
			name:   "ja",
			client: jaClient,
			args: args{
				method: "GET",
				urlStr: urlStr,
				body:   nil,
			},
			want:    jaHTTPRequest,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.client
			got, err := c.NewRequest(tt.args.method, tt.args.urlStr, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want, cmpopts.IgnoreUnexported(http.Request{})) {
				t.Errorf("Client.NewRequest() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
