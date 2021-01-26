package wbdata

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestIncomeLevelsService_List(t *testing.T) {
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
		want1   []*IncomeLevel
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   4,
				PerPage: 2,
				Total:   7,
			},
			want1: []*IncomeLevel{
				{
					ID:       "HIC",
					Iso2Code: "XD",
					Value:    "High income",
				},
				{
					ID:       "INX",
					Iso2Code: "XY",
					Value:    "Not classified",
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
			il := &IncomeLevelsService{
				client: client,
			}
			got, got1, err := il.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IncomeLevelsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncomeLevelsService.List() got = %v, want %v", got, tt.want)
			}
			for i := range got1 {
				if !reflect.DeepEqual(got1[i], tt.want1[i]) {
					t.Errorf("IncomeLevelsService.List() got1 = %v, want %v", got1[i], tt.want1[i])
				}
			}
		})
	}
}

func TestIncomeLevelsService_Get(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		incomeLevelID string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *IncomeLevel
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				incomeLevelID: testutils.TestDefaultIncomeLevelID,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &IncomeLevel{
				ID:       "HIC",
				Iso2Code: "XD",
				Value:    "High income",
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because incomeLevelID is invalid",
			args: args{
				incomeLevelID: testutils.TestInvalidIncomeLevelID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/incomeLevels/%s?format=json",
					defaultBaseURL,
					apiVersion,
					testutils.TestInvalidIncomeLevelID,
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
			il := &IncomeLevelsService{
				client: client,
			}
			got, got1, err := il.Get(tt.args.incomeLevelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IncomeLevelsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncomeLevelsService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IncomeLevelsService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
