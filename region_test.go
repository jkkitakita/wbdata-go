package wbdata

import (
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestRegionsService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		pages PageParams
	}
	tests := []struct {
		name             string
		args             args
		want             *PageSummary
		wantRegionsCount int
		wantErr          bool
	}{
		{
			name: "success",
			args: args{
				pages: PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantRegionsCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:          false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:             nil,
			wantRegionsCount: 0,
			wantErr:          true,
		},
		{
			name: "failure because PerPage is less than 1",
			args: args{
				pages: PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: 0,
				},
			},
			want:             nil,
			wantRegionsCount: 0,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RegionsService{
				client: client,
			}
			got, got1, err := r.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegionsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("RegionsService.List() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantRegionsCount {
				t.Errorf("RegionsService.List() got1 = %v, want %v", got1, tt.wantRegionsCount)
			}
		})
	}
}

func TestRegionsService_Get(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		code string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *Region
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				code: "afr",
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &Region{
				ID:       "", // empty
				Code:     "AFR",
				Iso2Code: "A9",
				Name:     "Africa",
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because code is invalid",
			args: args{
				code: invalidID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL:  defaultBaseURL + apiVersion + "/regions/" + invalidID + "?format=json",
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
			r := &RegionsService{
				client: client,
			}
			got, got1, err := r.Get(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegionsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionsService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RegionsService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
