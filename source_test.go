package wbdata

import (
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestSourcesService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

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
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
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
				pages: &PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:             nil,
			wantSourcesCount: 0,
			wantErr:          true,
		},
		{
			name: "failure because PerPage is less than 1",
			args: args{
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: 0,
				},
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
				sourceID: "11",
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &Source{
				ID:                   "11",
				LastUpdated:          "2013-02-22",
				Name:                 "Africa Development Indicators",
				Code:                 "ADI",
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
				sourceID: invalidID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL:  defaultBaseURL + apiVersion + "/sources/" + invalidID + "?format=json",
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
