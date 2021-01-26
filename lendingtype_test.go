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

	type args struct {
		pages *PageParams
	}
	tests := []struct {
		name    string
		args    args
		want    *PageSummary
		want1   []*LendingType
		wantErr bool
	}{
		{
			name: "success",
			args: args{},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   4,
			},
			want1: []*LendingType{
				{
					ID:       "IBD",
					Iso2Code: "XF",
					Value:    "IBRD",
				},
				{
					ID:       "IDB",
					Iso2Code: "XH",
					Value:    "Blend",
				},
				{
					ID:       "IDX",
					Iso2Code: "XI",
					Value:    "IDA",
				},
				{
					ID:       "LNX",
					Iso2Code: "XX",
					Value:    "Not classified",
				},
			},
			wantErr: false,
		},
		{
			name: "success with page params",
			args: args{
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   2,
				PerPage: 2,
				Total:   4,
			},
			want1: []*LendingType{
				{
					ID:       "IBD",
					Iso2Code: "XF",
					Value:    "IBRD",
				},
				{
					ID:       "IDB",
					Iso2Code: "XH",
					Value:    "Blend",
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
			lt := &LendingTypesService{
				client: client,
			}
			got, got1, err := lt.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("LendingTypesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LendingTypesService.List() got = %v, want %v", got, tt.want)
			}
			for i := range got1 {
				if !reflect.DeepEqual(got1[i], tt.want1[i]) {
					t.Errorf("LendingTypesService.List() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
				}
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
