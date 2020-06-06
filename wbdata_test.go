package wbdata

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewClient(t *testing.T) {
	baseURL, _ := url.Parse(defaultBaseURL + apiVersion + "/")
	jaLang := "ja"
	defaultClient := &Client{
		client:        &http.Client{},
		BaseURL:       baseURL,
		LocalLanguage: defaultLocalLanguage,
		Logger:        nil,
		UserAgent:     userAgent,
	}
	jaClient := &Client{
		client:        &http.Client{},
		BaseURL:       baseURL,
		LocalLanguage: jaLang,
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
					LocalLanguage(jaLang),
				},
			},
			want: jaClient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.httpClient, tt.args.options...)
			if !cmp.Equal(got, tt.want, optIgnoreUnexported, optIgnoreFields) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
