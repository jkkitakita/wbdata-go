package wbdata

import (
	"reflect"
	"strings"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestIndicatorValuesService_ListByCountryIDs(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	defaultCountryIDs := []string{"JPN", "USA"}
	invalidCountryIDs := []string{"ABCDEFG", "HIJKLM"}
	defaultIndicatorID := "NY.GDP.MKTP.CD"
	invalidIndicatorID := "INVALID.INDICATOR.ID"
	defaultDateStart := "2018"
	defaultDateEnd := "2019"
	invalidDate := "hoge"
	defaultDateParams := &DateParams{
		DateParamsType: DateParamsRange,
		DateRange: &DateRange{
			Start: defaultDateStart,
			End:   defaultDateEnd,
		},
	}
	invalidDateParams := &DateParams{
		DateParamsType: DateParamsRange,
		DateRange: &DateRange{
			Start: invalidDate,
			End:   defaultDateEnd,
		},
	}

	type args struct {
		countryIDs  []string
		indicatorID string
		datePatams  *DateParams
		pages       *PageParams
	}
	tests := []struct {
		name                     string
		args                     args
		want                     *PageSummaryWithSource
		wantIndicatorValuesCount int
		wantErr                  bool
		wantErrRes               *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				countryIDs:  defaultCountryIDs,
				indicatorID: defaultIndicatorID,
				datePatams:  defaultDateParams,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want: &PageSummaryWithSource{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantIndicatorValuesCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:                  false,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid CountryIDs",
			args: args{
				countryIDs:  invalidCountryIDs,
				indicatorID: defaultIndicatorID,
				datePatams:  defaultDateParams,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes: &ErrorResponse{
				URL: defaultBaseURL + apiVersion +
					"/countries/" + strings.Join(invalidCountryIDs, ";") +
					"/indicators/" + defaultIndicatorID +
					"?date=" + defaultDateStart + "%3A" + defaultDateEnd +
					"&format=json",
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
				countryIDs:  defaultCountryIDs,
				indicatorID: invalidIndicatorID,
				datePatams:  defaultDateParams,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes: &ErrorResponse{
				URL: defaultBaseURL + apiVersion +
					"/countries/" + strings.Join(defaultCountryIDs, ";") +
					"/indicators/" + invalidIndicatorID +
					"?date=" + defaultDateStart + "%3A" + defaultDateEnd +
					"&format=json",
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
			name: "failure because invalid date params",
			args: args{
				countryIDs:  defaultCountryIDs,
				indicatorID: defaultIndicatorID,
				datePatams:  invalidDateParams,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				countryIDs:  defaultCountryIDs,
				indicatorID: defaultIndicatorID,
				datePatams:  defaultDateParams,
				pages: &PageParams{
					Page:    testutils.TestInvalidPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes:               nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndicatorValuesService{
				client: client,
			}
			got, got1, err := i.ListByCountryIDs(tt.args.countryIDs, tt.args.indicatorID, tt.args.datePatams, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListByCountryIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListByCountryIDs() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("IndicatorValuesService.ListByCountryIDs() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantIndicatorValuesCount {
				t.Errorf("invalid length of IndicatorValuesService.ListByCountryIDs() got1 = %v, want %v", got1, tt.wantIndicatorValuesCount)
			}
			if !tt.wantErr {
				for i := range got1 {
					if got1[i].Date != defaultDateStart && got1[i].Date != defaultDateEnd {
						t.Errorf(
							"invalid date of IndicatorValuesService.ListByCountryIDs() got1 = %v, want %v or %v",
							got1[i].Date,
							defaultDateStart,
							defaultDateEnd,
						)
					}
					if got1[i].Indicator.ID != defaultIndicatorID {
						t.Errorf(
							"invalid indicator ID of IndicatorValuesService.ListByCountryIDs() got1[i].Indicator.ID = %v, want %v",
							got1[i].Indicator.ID,
							defaultIndicatorID,
						)
					}
				}
			}
		})
	}
}
