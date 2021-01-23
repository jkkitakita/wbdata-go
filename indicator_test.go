package wbdata

import (
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestIndicatorsService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

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
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
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
				pages: &PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:                nil,
			wantIndicatorsCount: 0,
			wantErr:             true,
		},
		{
			name: "failure because PerPage is less than 1",
			args: args{
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: 0,
				},
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
				indicatorID: invalidID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL:  defaultBaseURL + apiVersion + "/indicators/" + invalidID + "?format=json",
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

	defaultTopicID := "1"
	invalidTopicID := "1000"

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
				topicID: defaultTopicID,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
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
				topicID: invalidTopicID,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:                nil,
			wantIndicatorsCount: 0,
			wantErr:             true,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				topicID: defaultTopicID,
				pages: &PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:                nil,
			wantIndicatorsCount: 0,
			wantErr:             true,
		},
		{
			name: "failure because PerPage is less than 1",
			args: args{
				topicID: defaultTopicID,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: 0,
				},
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
						if topics[j].ID == defaultTopicID {
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
