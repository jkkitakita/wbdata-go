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
		name    string
		args    args
		want    *PageSummary
		want1   []*Indicator
		wantErr bool
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
			want1: []*Indicator{
				{
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
				{
					ID:   "1.0.HCount.2.5usd",
					Name: "Poverty Headcount ($2.50 a day)",
					Unit: "", // NOTE: always empty?
					Source: &IDAndValue{
						ID:    "37",
						Value: "LAC Equity Lab",
					},
					SourceNote: "The poverty headcount index measures the proportion of " +
						"the population with daily per capita income (in 2005 PPP) below the poverty line.",
					SourceOrganization: "LAC Equity Lab tabulations of SEDLAC (CEDLAS and the World Bank).",
					Topics: []*IDAndValue{
						{
							ID:    "11",
							Value: "Poverty ",
						},
					},
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
			i := &IndicatorsService{
				client: client,
			}
			got, got1, err := i.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && (got.Page != tt.want.Page || got.PerPage != tt.want.PerPage) {
				t.Errorf("IndicatorsService.List() got = %v, want %v", got, tt.want)
			}
			for i := range got1 {
				if !reflect.DeepEqual(got1[i], tt.want1[i]) {
					t.Errorf("IndicatorsService.List() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
				}
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
		name    string
		args    args
		want    *PageSummary
		want1   []*Indicator
		wantErr bool
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
			want1: []*Indicator{
				{
					ID:   "AG.AGR.TRAC.NO",
					Name: "Agricultural machinery, tractors",
					Unit: "", // NOTE: always empty?
					Source: &IDAndValue{
						ID:    "2",
						Value: "World Development Indicators",
					},
					SourceNote: "Agricultural machinery refers to the number of wheel and crawler tractors (excluding garden tractors) " +
						"in use in agriculture at the end of the calendar year specified or during the first quarter of the following year.",
					SourceOrganization: "Food and Agriculture Organization, electronic files and web site.",
					Topics: []*IDAndValue{
						{
							ID:    "1",
							Value: "Agriculture & Rural Development  ",
						},
					},
				},
				{
					ID:   "AG.CON.FERT.PT.ZS",
					Name: "Fertilizer consumption (% of fertilizer production)",
					Unit: "", // NOTE: always empty?
					Source: &IDAndValue{
						ID:    "2",
						Value: "World Development Indicators",
					},
					SourceNote: "Fertilizer consumption measures the quantity of plant nutrients " +
						"used per unit of arable land. Fertilizer products cover nitrogenous, " +
						"potash, and phosphate fertilizers (including ground rock phosphate). " +
						"Traditional nutrients--animal and plant manures--are not included. " +
						"For the purpose of data dissemination, " +
						"FAO has adopted the concept of a calendar year (January to December). " +
						"Some countries compile fertilizer data on a calendar year basis, " +
						"while others are on a split-year basis.",
					SourceOrganization: "Food and Agriculture Organization, electronic files and web site.",
					Topics: []*IDAndValue{
						{
							ID:    "1",
							Value: "Agriculture & Rural Development  ",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "failure because invalid topic id",
			args: args{
				topicID: testutils.TestInvalidTopicID,
				pages:   defaultPageParams,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				topicID: testutils.TestDefaultTopicID,
				pages:   invalidPageParams,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
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
			if tt.want != nil && (got.Page != tt.want.Page || got.PerPage != tt.want.PerPage) {
				t.Errorf("IndicatorsService.ListByTopicID() got = %v, want %v", got, tt.want)
			}
			for i := range got1 {
				if !reflect.DeepEqual(got1[i], tt.want1[i]) {
					t.Errorf("IndicatorsService.ListByTopicID() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
				}
			}
		})
	}
}
