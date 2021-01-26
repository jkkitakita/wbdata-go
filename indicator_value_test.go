package wbdata

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestIndicatorValuesService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	optIgnoreFields := cmpopts.IgnoreFields(
		PageSummaryWithSourceID{},
		"Pages",
		"Total",
		"LastUpdated",
	)

	defaultFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestDefaultDateStart,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultFilterRecentParams := &FilterParams{
		FilterParamsType: FilterParamsMRV,
		RecentParam: &RecentParam{
			FrequencyType:    FrequencyYearly,
			MostRecentValues: 1,
			IsNotEmpty:       false,
			IsGapFill:        false,
		},
	}
	invalidFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestInvalidDate,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		indicatorID  string
		filterParams *FilterParams
		pages        *PageParams
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummaryWithSourceID
		want1      []*IndicatorValue
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				indicatorID: testutils.TestDefaultIndicatorID,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  50,
				SourceID: "2",
			},
			want1:      nil,
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with filter date params",
			args: args{
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: defaultFilterDateParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  2,
				SourceID: "2",
			},
			want1: []*IndicatorValue{
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "1A",
						Value: "Arab World",
					},
					Countryiso3code: "ARB",
					Date:            "2019",
					Value:           2.81741458466511e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "1A",
						Value: "Arab World",
					},
					Countryiso3code: "ARB",
					Date:            "2018",
					Value:           2.77138409790453e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with filter recent params",
			args: args{
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: defaultFilterRecentParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  2,
				SourceID: "2",
			},
			want1: []*IndicatorValue{
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "1A",
						Value: "Arab World",
					},
					Countryiso3code: "ARB",
					Date:            "2019",
					Value:           2.81741458466511e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "S3",
						Value: "Caribbean small states",
					},
					Countryiso3code: "CSS",
					Date:            "2019",
					Value:           7.77217149178506e+10,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid indicator id",
			args: args{
				indicatorID: testutils.TestInvalidIndicatorID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/all/indicators/%s?format=json",
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
		{
			name: "failure because invalid filter params",
			args: args{
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: invalidFilterDateParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				indicatorID: testutils.TestDefaultIndicatorID,
				pages:       invalidPageParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.List(tt.args.indicatorID, tt.args.filterParams, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.List() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil && !cmp.Equal(got, tt.want, nil, optIgnoreFields) {
				t.Errorf("NewClient() = %+v, want %+v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("IndicatorValuesService.List() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListWithFootnote(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	optIgnoreFields := cmpopts.IgnoreFields(
		PageSummaryWithSourceID{},
		"Pages",
		"Total",
		"LastUpdated",
	)

	defaultFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestDefaultDateStart,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	invalidFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestInvalidDate,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		indicatorID  string
		filterParams *FilterParams
		pages        *PageParams
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummaryWithSourceID
		want1      []*IndicatorValueWithFootnote
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				indicatorID: testutils.TestDefaultIndicatorID,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  50,
				SourceID: "2",
			},
			want1:      nil,
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with filter params",
			args: args{
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: defaultFilterDateParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  2,
				SourceID: "2",
			},
			want1: []*IndicatorValueWithFootnote{
				{
					IndicatorValue: IndicatorValue{
						Indicator: IDAndValue{
							ID:    "NY.GDP.MKTP.CD",
							Value: "GDP (current US$)",
						},
						Country: IDAndValue{
							ID:    "1A",
							Value: "Arab World",
						},
						Countryiso3code: "ARB",
						Date:            "2019",
						Value:           2.81741458466511e+12,
						Unit:            "",
						ObsStatus:       "",
						Decimal:         0,
					},
					Footnote: "",
				},
				{
					IndicatorValue: IndicatorValue{
						Indicator: IDAndValue{
							ID:    "NY.GDP.MKTP.CD",
							Value: "GDP (current US$)",
						},
						Country: IDAndValue{
							ID:    "1A",
							Value: "Arab World",
						},
						Countryiso3code: "ARB",
						Date:            "2018",
						Value:           2.77138409790453e+12,
						Unit:            "",
						ObsStatus:       "",
						Decimal:         0,
					},
					Footnote: "",
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid indicator id",
			args: args{
				indicatorID: testutils.TestInvalidIndicatorID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/all/indicators/%s?footnote=y&format=json",
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
		{
			name: "failure because invalid filter params",
			args: args{
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: invalidFilterDateParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				indicatorID: testutils.TestDefaultIndicatorID,
				pages:       invalidPageParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.ListWithFootnote(tt.args.indicatorID, tt.args.filterParams, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListWithFootnote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListWithFootnote() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil && !cmp.Equal(got, tt.want, nil, optIgnoreFields) {
				t.Errorf("NewClient() = %+v, want %+v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("IndicatorValuesService.ListWithFootnote() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListByCountryIDs(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	optIgnoreFields := cmpopts.IgnoreFields(
		PageSummaryWithSourceID{},
		"Pages",
		"Total",
		"LastUpdated",
	)

	defaultFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestDefaultDateStart,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	invalidFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestInvalidDate,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		countryIDs   []string
		indicatorID  string
		filterParams *FilterParams
		pages        *PageParams
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummaryWithSourceID
		want1      []*IndicatorValue
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				countryIDs:  testutils.TestDefaultCountryIDs,
				indicatorID: testutils.TestDefaultIndicatorID,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  50,
				SourceID: "2",
			},
			want1:      nil,
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: defaultFilterDateParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  2,
				SourceID: "2",
			},
			want1: []*IndicatorValue{
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "JP",
						Value: "Japan",
					},
					Countryiso3code: "JPN",
					Date:            "2019",
					Value:           5.08176954237977e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "JP",
						Value: "Japan",
					},
					Countryiso3code: "JPN",
					Date:            "2018",
					Value:           4.95480661999519e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid CountryIDs",
			args: args{
				countryIDs:  testutils.TestInvalidCountryIDs,
				indicatorID: testutils.TestDefaultIndicatorID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?format=json",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestInvalidCountryIDs, ";"),
					testutils.TestDefaultIndicatorID,
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
		{
			name: "failure because invalid indicatorID",
			args: args{
				countryIDs:  testutils.TestDefaultCountryIDs,
				indicatorID: testutils.TestInvalidIndicatorID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?format=json",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestDefaultCountryIDs, ";"),
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
		{
			name: "failure because invalid filter params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: invalidFilterDateParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				countryIDs:  testutils.TestDefaultCountryIDs,
				indicatorID: testutils.TestDefaultIndicatorID,
				pages:       invalidPageParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.ListByCountryIDs(tt.args.countryIDs, tt.args.indicatorID, tt.args.filterParams, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListByCountryIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListByCountryIDs() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil && !cmp.Equal(got, tt.want, nil, optIgnoreFields) {
				t.Errorf("IndicatorValuesService.ListByCountryIDs() got = %+v, want %+v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("IndicatorValuesService.ListByCountryIDs() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListByCountryIDsWithFootnote(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	optIgnoreFields := cmpopts.IgnoreFields(
		PageSummaryWithSourceID{},
		"Pages",
		"Total",
		"LastUpdated",
	)

	defaultFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestDefaultDateStart,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	invalidFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestInvalidDate,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		countryIDs   []string
		indicatorID  string
		filterParams *FilterParams
		pages        *PageParams
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummaryWithSourceID
		want1      []*IndicatorValueWithFootnote
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				countryIDs:  testutils.TestDefaultCountryIDs,
				indicatorID: testutils.TestDefaultIndicatorID,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  50,
				SourceID: "2",
			},
			want1:      nil,
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: defaultFilterDateParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithSourceID{
				Page:     1,
				PerPage:  2,
				SourceID: "2",
			},
			want1: []*IndicatorValueWithFootnote{
				{
					IndicatorValue: IndicatorValue{
						Indicator: IDAndValue{
							ID:    "NY.GDP.MKTP.CD",
							Value: "GDP (current US$)",
						},
						Country: IDAndValue{
							ID:    "JP",
							Value: "Japan",
						},
						Countryiso3code: "JPN",
						Date:            "2019",
						Value:           5.08176954237977e+12,
						Unit:            "",
						ObsStatus:       "",
						Decimal:         0,
					},
					Footnote: "",
				},
				{
					IndicatorValue: IndicatorValue{
						Indicator: IDAndValue{
							ID:    "NY.GDP.MKTP.CD",
							Value: "GDP (current US$)",
						},
						Country: IDAndValue{
							ID:    "JP",
							Value: "Japan",
						},
						Countryiso3code: "JPN",
						Date:            "2018",
						Value:           4.95480661999519e+12,
						Unit:            "",
						ObsStatus:       "",
						Decimal:         0,
					},
					Footnote: "",
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid CountryIDs",
			args: args{
				countryIDs:  testutils.TestInvalidCountryIDs,
				indicatorID: testutils.TestDefaultIndicatorID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?footnote=y&format=json",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestInvalidCountryIDs, ";"),
					testutils.TestDefaultIndicatorID,
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
		{
			name: "failure because invalid indicatorID",
			args: args{
				countryIDs:  testutils.TestDefaultCountryIDs,
				indicatorID: testutils.TestInvalidIndicatorID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?footnote=y&format=json",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestDefaultCountryIDs, ";"),
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
		{
			name: "failure because invalid filter params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorID:  testutils.TestDefaultIndicatorID,
				filterParams: invalidFilterDateParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				countryIDs:  testutils.TestDefaultCountryIDs,
				indicatorID: testutils.TestDefaultIndicatorID,
				pages:       invalidPageParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.ListByCountryIDsWithFootnote(tt.args.countryIDs, tt.args.indicatorID, tt.args.filterParams, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListByCountryIDsWithFootnote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListByCountryIDsWithFootnote() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil && !cmp.Equal(got, tt.want, nil, optIgnoreFields) {
				t.Errorf("IndicatorValuesService.ListByCountryIDsWithFootnote() got = %+v, want %+v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("IndicatorValuesService.ListByCountryIDsWithFootnote() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListBySourceID(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	optIgnoreFields := cmpopts.IgnoreFields(
		PageSummaryWithLastUpdated{},
		"Pages",
		"Total",
		"LastUpdated",
	)

	defaultFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestDefaultDateStart,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	invalidFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestInvalidDate,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		indicatorIDs []string
		sourceID     string
		filterParams *FilterParams
		pages        *PageParams
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummaryWithLastUpdated
		want1      []*IndicatorValue
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want: &PageSummaryWithLastUpdated{
				Page:    1,
				PerPage: 50,
			},
			want1:      nil,
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with params",
			args: args{
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				filterParams: defaultFilterDateParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithLastUpdated{
				Page:    1,
				PerPage: 2,
			},
			want1: []*IndicatorValue{
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "1A",
						Value: "Arab World",
					},
					Countryiso3code: "ARB",
					Date:            "2019",
					Value:           2.81741458466511e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "1A",
						Value: "Arab World",
					},
					Countryiso3code: "ARB",
					Date:            "2018",
					Value:           2.77138409790453e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid indicator id",
			args: args{
				indicatorIDs: testutils.TestInvalidIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/all/indicators/%s?format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestInvalidIndicatorIDs, ";"),
					testutils.TestDefaultSourceID,
				),
				Code: 200,
				// http://api.worldbank.org/v2/countries/all/indicators/INVALID.INDICATOR.ID;SP.POP.TOTL?format=json&source=2
				Message: []ErrorMessage{
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
				},
			},
		},
		{
			name: "failure because invalid source id",
			args: args{
				indicatorIDs: testutils.TestInvalidIndicatorIDs,
				sourceID:     testutils.TestInvalidSourceID,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid filter params",
			args: args{
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				filterParams: invalidFilterDateParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				pages:        invalidPageParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.ListBySourceID(tt.args.indicatorIDs, tt.args.sourceID, tt.args.filterParams, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListBySourceID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListBySourceID() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil && !cmp.Equal(got, tt.want, nil, optIgnoreFields) {
				t.Errorf("IndicatorValuesService.ListBySourceID() got = %+v, want %+v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("IndicatorValuesService.ListBySourceID() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListBySourceIDWithFootnote(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	optIgnoreFields := cmpopts.IgnoreFields(
		PageSummaryWithLastUpdated{},
		"Pages",
		"Total",
		"LastUpdated",
	)

	defaultFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestDefaultDateStart,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	invalidFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestInvalidDate,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		indicatorIDs []string
		sourceID     string
		filterParams *FilterParams
		pages        *PageParams
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummaryWithLastUpdated
		want1      []*IndicatorValueWithFootnote
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want: &PageSummaryWithLastUpdated{
				Page:    1,
				PerPage: 50,
			},
			want1:      nil,
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with params",
			args: args{
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				filterParams: defaultFilterDateParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithLastUpdated{
				Page:    1,
				PerPage: 2,
			},
			want1: []*IndicatorValueWithFootnote{
				{
					IndicatorValue: IndicatorValue{
						Indicator: IDAndValue{
							ID:    "NY.GDP.MKTP.CD",
							Value: "GDP (current US$)",
						},
						Country: IDAndValue{
							ID:    "1A",
							Value: "Arab World",
						},
						Countryiso3code: "ARB",
						Date:            "2019",
						Value:           2.81741458466511e+12,
						Unit:            "",
						ObsStatus:       "",
						Decimal:         0,
					},
					Footnote: "",
				},
				{
					IndicatorValue: IndicatorValue{
						Indicator: IDAndValue{
							ID:    "NY.GDP.MKTP.CD",
							Value: "GDP (current US$)",
						},
						Country: IDAndValue{
							ID:    "1A",
							Value: "Arab World",
						},
						Countryiso3code: "ARB",
						Date:            "2018",
						Value:           2.77138409790453e+12,
						Unit:            "",
						ObsStatus:       "",
						Decimal:         0,
					},
					Footnote: "",
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid indicator id",
			args: args{
				indicatorIDs: testutils.TestInvalidIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/all/indicators/%s?footnote=y&format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestInvalidIndicatorIDs, ";"),
					testutils.TestDefaultSourceID,
				),
				Code: 200,
				// http://api.worldbank.org/v2/countries/all/indicators/INVALID.INDICATOR.ID;SP.POP.TOTL?format=json&source=2
				Message: []ErrorMessage{
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
				},
			},
		},
		{
			name: "failure because invalid source id",
			args: args{
				indicatorIDs: testutils.TestInvalidIndicatorIDs,
				sourceID:     testutils.TestInvalidSourceID,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid filter params",
			args: args{
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				filterParams: invalidFilterDateParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				pages:        invalidPageParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.ListBySourceIDWithFootnote(
				tt.args.indicatorIDs,
				tt.args.sourceID,
				tt.args.filterParams,
				tt.args.pages,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListBySourceIDWithFootnote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListBySourceIDWithFootnote() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil && !cmp.Equal(got, tt.want, nil, optIgnoreFields) {
				t.Errorf("IndicatorValuesService.ListBySourceIDWithFootnote() got = %+v, want %+v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("IndicatorValuesService.ListBySourceIDWithFootnote() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListByCountryIDsAndSourceID(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	optIgnoreFields := cmpopts.IgnoreFields(
		PageSummaryWithLastUpdated{},
		"Pages",
		"Total",
		"LastUpdated",
	)

	defaultFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestDefaultDateStart,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	invalidFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestInvalidDate,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		countryIDs   []string
		indicatorIDs []string
		sourceID     string
		filterParams *FilterParams
		pages        *PageParams
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummaryWithLastUpdated
		want1      []*IndicatorValue
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want:       nil,
			want1:      nil,
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				filterParams: defaultFilterDateParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithLastUpdated{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			want1: []*IndicatorValue{
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "JP",
						Value: "Japan",
					},
					Countryiso3code: "JPN",
					Date:            "2019",
					Value:           5.08176954237977e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
				{
					Indicator: IDAndValue{
						ID:    "NY.GDP.MKTP.CD",
						Value: "GDP (current US$)",
					},
					Country: IDAndValue{
						ID:    "JP",
						Value: "Japan",
					},
					Countryiso3code: "JPN",
					Date:            "2018",
					Value:           4.95480661999519e+12,
					Unit:            "",
					ObsStatus:       "",
					Decimal:         0,
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid country ids",
			args: args{
				countryIDs:   testutils.TestInvalidCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestInvalidCountryIDs, ";"),
					strings.Join(testutils.TestDefaultIndicatorIDs, ";"),
					testutils.TestDefaultSourceID,
				),
				Code: 200,
				// http://api.worldbank.org/v2/countries/all/indicators/INVALID.INDICATOR.ID;SP.POP.TOTL?format=json&source=2
				Message: []ErrorMessage{
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
				},
			},
		},
		{
			name: "failure because invalid indicator id",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestInvalidIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestDefaultCountryIDs, ";"),
					strings.Join(testutils.TestInvalidIndicatorIDs, ";"),
					testutils.TestDefaultSourceID,
				),
				Code: 200,
				// http://api.worldbank.org/v2/countries/all/indicators/INVALID.INDICATOR.ID;SP.POP.TOTL?format=json&source=2
				Message: []ErrorMessage{
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
				},
			},
		},
		{
			name: "failure because invalid source id",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestInvalidSourceID,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid filter params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				filterParams: invalidFilterDateParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				pages:        invalidPageParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.ListByCountryIDsAndSourceID(
				tt.args.countryIDs,
				tt.args.indicatorIDs,
				tt.args.sourceID,
				tt.args.filterParams,
				tt.args.pages,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListByCountryIDsAndSourceID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListBySourceID() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil && !cmp.Equal(got, tt.want, nil, optIgnoreFields) {
				t.Errorf("IndicatorValuesService.ListBySourceID() got = %+v, want %+v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("IndicatorValuesService.ListBySourceID() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListByCountryIDsAndSourceIDWithFootnote(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	optIgnoreFields := cmpopts.IgnoreFields(
		PageSummaryWithLastUpdated{},
		"Pages",
		"Total",
		"LastUpdated",
	)

	defaultFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestDefaultDateStart,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	invalidFilterDateParams := &FilterParams{
		FilterParamsType: FilterParamsDateRange,
		DateParam: &DateParam{
			DateRange: &DateRange{
				Start: testutils.TestInvalidDate,
				End:   testutils.TestDefaultDateEnd,
			},
		},
	}
	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		countryIDs   []string
		indicatorIDs []string
		sourceID     string
		filterParams *FilterParams
		pages        *PageParams
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummaryWithLastUpdated
		want1      []*IndicatorValueWithFootnote
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want:       nil,
			want1:      nil,
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "success with params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				filterParams: defaultFilterDateParams,
				pages:        defaultPageParams,
			},
			want: &PageSummaryWithLastUpdated{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			want1: []*IndicatorValueWithFootnote{
				{
					IndicatorValue: IndicatorValue{
						Indicator: IDAndValue{
							ID:    "NY.GDP.MKTP.CD",
							Value: "GDP (current US$)",
						},
						Country: IDAndValue{
							ID:    "JP",
							Value: "Japan",
						},
						Countryiso3code: "JPN",
						Date:            "2019",
						Value:           5.08176954237977e+12,
						Unit:            "",
						ObsStatus:       "",
						Decimal:         0,
					},
					Footnote: "",
				},
				{
					IndicatorValue: IndicatorValue{
						Indicator: IDAndValue{
							ID:    "NY.GDP.MKTP.CD",
							Value: "GDP (current US$)",
						},
						Country: IDAndValue{
							ID:    "JP",
							Value: "Japan",
						},
						Countryiso3code: "JPN",
						Date:            "2018",
						Value:           4.95480661999519e+12,
						Unit:            "",
						ObsStatus:       "",
						Decimal:         0,
					},
					Footnote: "",
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid country ids",
			args: args{
				countryIDs:   testutils.TestInvalidCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?footnote=y&format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestInvalidCountryIDs, ";"),
					strings.Join(testutils.TestDefaultIndicatorIDs, ";"),
					testutils.TestDefaultSourceID,
				),
				Code: 200,
				// http://api.worldbank.org/v2/countries/all/indicators/INVALID.INDICATOR.ID;SP.POP.TOTL?format=json&source=2
				Message: []ErrorMessage{
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
				},
			},
		},
		{
			name: "failure because invalid indicator id",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestInvalidIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?footnote=y&format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(testutils.TestDefaultCountryIDs, ";"),
					strings.Join(testutils.TestInvalidIndicatorIDs, ";"),
					testutils.TestDefaultSourceID,
				),
				Code: 200,
				// http://api.worldbank.org/v2/countries/all/indicators/INVALID.INDICATOR.ID;SP.POP.TOTL?footnote=y&format=json&source=2
				Message: []ErrorMessage{
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
					{
						ID:    "120",
						Key:   "Invalid value",
						Value: "The provided parameter value is not valid",
					},
				},
			},
		},
		{
			name: "failure because invalid source id",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestInvalidSourceID,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid filter params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				filterParams: invalidFilterDateParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				countryIDs:   testutils.TestDefaultCountryIDs,
				indicatorIDs: testutils.TestDefaultIndicatorIDs,
				sourceID:     testutils.TestDefaultSourceID,
				pages:        invalidPageParams,
			},
			want:       nil,
			want1:      nil,
			wantErr:    true,
			wantErrRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.ListByCountryIDsAndSourceIDWithFootnote(
				tt.args.countryIDs,
				tt.args.indicatorIDs,
				tt.args.sourceID,
				tt.args.filterParams,
				tt.args.pages,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"IndicatorValuesService.ListByCountryIDsAndSourceIDWithFootnote() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf(
						"IndicatorValuesService.ListByCountryIDsAndSourceIDWithFootnote() err = %v, wantErrRes %v",
						err,
						tt.wantErrRes,
					)
				}
			}
			if tt.want != nil && !cmp.Equal(got, tt.want, nil, optIgnoreFields) {
				t.Errorf(
					"IndicatorValuesService.ListByCountryIDsAndSourceIDWithFootnote() got = %+v, want %+v",
					got,
					tt.want,
				)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf(
							"IndicatorValuesService.ListByCountryIDsAndSourceIDWithFootnote() got1[i] = %v, want[i] %v",
							got1[i],
							tt.want1[i],
						)
					}
				}
			}
		})
	}
}
