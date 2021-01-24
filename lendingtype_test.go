package wbdata

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestLendingTypesService_List(t *testing.T) {
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
	invalidPerPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestInvalidPerPage,
	}

	type args struct {
		pages *PageParams
	}
	tests := []struct {
		name                  string
		args                  args
		want                  *PageSummary
		wantLendingTypesCount int
		wantErr               bool
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
			wantLendingTypesCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:               false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: invalidPageParams,
			},
			want:                  nil,
			wantLendingTypesCount: 0,
			wantErr:               true,
		},
		{
			name: "failure because PerPage is less than 1",
			args: args{
				pages: invalidPerPageParams,
			},
			want:                  nil,
			wantLendingTypesCount: 0,
			wantErr:               true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lt := &LendingTypesService{
				client: client,
			}
			got, got1, err := lt.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("c.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("LendingTypesService.List() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantLendingTypesCount {
				t.Errorf("LendingTypesService.List() got1 = %v, want %v", got1, tt.wantLendingTypesCount)
			}
		})
	}
}

func TestLendingTypesService_Get(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		lendingTypeID string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *LendingType
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				lendingTypeID: testutils.TestDefaultLendingTypeID,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &LendingType{
				ID:       "IBD",
				Iso2Code: "XF",
				Value:    "IBRD",
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because lendingTypeID is invalid",
			args: args{
				lendingTypeID: testutils.TestInvalidLendingTypeID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/lendingTypes/%s?format=json",
					defaultBaseURL,
					apiVersion,
					testutils.TestInvalidLendingTypeID,
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
			lt := &LendingTypesService{
				client: client,
			}
			got, got1, err := lt.Get(tt.args.lendingTypeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LendingTypesService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LendingTypesService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LendingTypesService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
