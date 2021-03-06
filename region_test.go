package wbdata

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestRegionsService_List(t *testing.T) {
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
		name    string
		args    args
		want    *PageSummary
		want1   []*Region
		wantErr bool
	}{
		{
			name: "success",
			args: args{},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   42,
			},
			want1:   nil,
			wantErr: false,
		},
		{
			name: "success with page params",
			args: args{
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   21,
				PerPage: 2,
				Total:   42,
			},
			want1: []*Region{
				{
					ID:       "",
					Code:     "AFR",
					Iso2Code: "A9",
					Name:     "Africa",
				},
				{
					ID:       "",
					Code:     "ARB",
					Iso2Code: "1A",
					Name:     "Arab World",
				},
			},
			wantErr: false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: invalidPageParams,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegionsService.List() got = %v, want %v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("RegionsService.List() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
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
				code: testutils.TestDefaultRegionCode,
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
				code: testutils.TestInvalidRegionCode,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/regions/%s?format=json",
					defaultBaseURL,
					apiVersion,
					testutils.TestInvalidRegionCode,
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
