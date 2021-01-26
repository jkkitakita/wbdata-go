package wbdata

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestIndicatorsService_List(t *testing.T) {
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
		name                string
		args                args
		want                *PageSummary
		wantIndicatorsCount int
		wantErr             bool
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
			wantIndicatorsCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:             false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: invalidPageParams,
			},
			want:                nil,
			wantIndicatorsCount: 0,
			wantErr:             true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorsService{
				client: client,
			}
			got, got1, err := i.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("IndicatorsService.List() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantIndicatorsCount {
				t.Errorf("IndicatorsService.List() got1 = %v, want %v", got1, tt.wantIndicatorsCount)
			}
		})
	}
}

func TestIndicatorsService_Get(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		indicatorID string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *Indicator
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				indicatorID: "1.0.hcount.1.90usd",
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &Indicator{
				ID:   "1.0.HCount.1.90usd",
				Name: "Poverty Headcount ($1.90 a day)",
				Unit: "", // NOTE: always empty?
				Source: &IDAndValue{
					ID:    "37",
					Value: "LAC Equity Lab",
				},
				SourceNote: "The poverty headcount index measures the proportion of " +
					"the population with daily per capita income (in 2011 PPP) below the poverty line.",
				SourceOrganization: "LAC Equity Lab tabulations of SEDLAC (CEDLAS and the World Bank).",
				Topics: []*IDAndValue{
					{
						ID:    "11",
						Value: "Poverty ",
					},
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because indicatorID is invalid",
			args: args{
				indicatorID: testutils.TestInvalidIndicatorID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/indicators/%s?format=json",
					defaultBaseURL,
					apiVersion,
					testutils.TestInvalidIndicatorID,
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
			i := &IndicatorsService{
				client: client,
			}
			got, got1, err := i.Get(tt.args.indicatorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndicatorsService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IndicatorsService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIndicatorsService_ListByTopicID(t *testing.T) {
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
		topicID string
		pages   *PageParams
	}
	tests := []struct {
		name                string
		args                args
		want                *PageSummary
		wantIndicatorsCount int
		wantErr             bool
	}{
		{
			name: "success",
			args: args{
				topicID: testutils.TestDefaultTopicID,
				pages:   defaultPageParams,
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantIndicatorsCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:             false,
		},
		{
			name: "failure because invalid topic id",
			args: args{
				topicID: testutils.TestInvalidTopicID,
				pages:   defaultPageParams,
			},
			want:                nil,
			wantIndicatorsCount: 0,
			wantErr:             true,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				topicID: testutils.TestDefaultTopicID,
				pages:   invalidPageParams,
			},
			want:                nil,
			wantIndicatorsCount: 0,
			wantErr:             true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorsService{
				client: client,
			}
			got, got1, err := i.ListByTopicID(tt.args.topicID, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorsService.ListByTopicID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("IndicatorsService.ListByTopicID() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantIndicatorsCount {
				t.Errorf("IndicatorsService.ListByTopicID() got1 = %v, want %v", got1, tt.wantIndicatorsCount)
			} else {
				var topics []*IDAndValue
				var cnt int
				for i := range got1 {
					topics = got1[i].Topics
					cnt = 0
					for j := range topics {
						if topics[j].ID == tt.args.topicID {
							cnt++
						}
					}
					if cnt == 0 {
						t.Errorf("invalid topic id. got1 = %+v", got1)
					}
				}
			}
		})
	}
}
