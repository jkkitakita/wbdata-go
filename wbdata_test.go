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
	jaLanguage := &Language{
		Code: "ja",
	}
	optIgnoreUnexported := cmpopts.IgnoreUnexported(Client{})
	optIgnoreFields := cmpopts.IgnoreFields(Client{},
		"Countries",
		"Indicators",
		"IndicatorValues",
		"IncomeLevels",
		"LendingTypes",
		"Regions",
		"Sources",
		"Topics",
		"Languages",
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
			want: &Client{
				client:       &http.Client{},
				BaseURL:      baseURL,
				OutputFormat: OutputFormatJSON,
				UserAgent:    userAgent,
			},
		},
		{
			name: "Language_ja",
			args: args{
				httpClient: &http.Client{},
				options: []func(*Client){
					SetLanguage(jaLanguage),
				},
			},
			want: &Client{
				client:       &http.Client{},
				BaseURL:      baseURL,
				Language:     testutils.TestDefaultLanguageCode,
				OutputFormat: OutputFormatJSON,
				UserAgent:    userAgent,
			},
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
	defaultRequestURL, _ := url.Parse(fmt.Sprintf("%s%s?format=%s", baseURL, urlStr, OutputFormatJSON))
	jaRequestURL, _ := url.Parse(fmt.Sprintf("%s%s/%s?format=%s", baseURL, testutils.TestDefaultLanguageCode, urlStr, OutputFormatJSON))

	type args struct {
		method      string
		urlStr      string
		queryParams map[string]string
		body        interface{}
	}
	tests := []struct {
		name    string
		client  *Client
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "default",
			client: &Client{
				client:       &http.Client{},
				BaseURL:      baseURL,
				OutputFormat: OutputFormatJSON,
				UserAgent:    userAgent,
			},
			args: args{
				method:      "GET",
				urlStr:      urlStr,
				queryParams: nil,
				body:        nil,
			},
			want: &http.Request{
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
			},
			wantErr: false,
		},
		{
			name: "Language_ja",
			client: &Client{
				client:       &http.Client{},
				BaseURL:      baseURL,
				Language:     testutils.TestDefaultLanguageCode,
				OutputFormat: OutputFormatJSON,
				UserAgent:    userAgent,
			},
			args: args{
				method:      "GET",
				urlStr:      urlStr,
				queryParams: nil,
				body:        nil,
			},
			want: &http.Request{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.client
			got, err := c.NewRequest(tt.args.method, tt.args.urlStr, tt.args.queryParams, tt.args.body)
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
