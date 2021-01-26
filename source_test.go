package wbdata

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestSourcesService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		pages *PageParams
	}
	tests := []struct {
		name             string
		args             args
		want             *PageSummary
		wantSourcesCount int
		wantErr          bool
	}{
		{
			name: "success",
			args: args{
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantSourcesCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:          false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: invalidPageParams,
			},
			want:             nil,
			wantSourcesCount: 0,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SourcesService{
				client: client,
			}
			got, got1, err := s.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("SourcesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("SourcesService.List() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantSourcesCount {
				t.Errorf("SourcesService.List() got1 = %v, want %v", got1, tt.wantSourcesCount)
			}
		})
	}
}

func TestSourcesService_Get(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		sourceID string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *Source
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				sourceID: testutils.TestDefaultSourceID,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &Source{
				ID:                   "2",
				LastUpdated:          "2020-12-16",
				Name:                 "World Development Indicators",
				Code:                 "WDI",
				Description:          "", // NOTE: always empty?
				URL:                  "", // NOTE: always empty?
				DataAvailability:     "Y",
				MetadataAvailability: "Y",
				Concepts:             "3",
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because sourceID is invalid",
			args: args{
				sourceID: testutils.TestInvalidSourceID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/sources/%s?format=json",
					defaultBaseURL,
					apiVersion,
					testutils.TestInvalidSourceID,
				),
				Code: 200,
				Message: []ErrorMessage{
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SourcesService{
				client: client,
			}
			got, got1, err := s.Get(tt.args.sourceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SourcesService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SourcesService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SourcesService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
