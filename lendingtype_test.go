package wbdata

import (
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestLendingTypesService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		pages PageParams
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
				pages: PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
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
				pages: PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:                  nil,
			wantLendingTypesCount: 0,
			wantErr:               true,
		},
		{
			name: "failure because PerPage is less than 1",
			args: args{
				pages: PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: 0,
				},
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
				lendingTypeID: "ibd",
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
				lendingTypeID: invalidID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL:  defaultBaseURL + apiVersion + "/lendingTypes/" + invalidID + "?format=json",
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
