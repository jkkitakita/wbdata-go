package wbdata

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

const (
	defaultIndicatorID = "NY.GDP.MKTP.CD"
	invalidIndicatorID = "INVALID.INDICATOR.ID"
	defaultSourceID    = "2"
	invalidSourceID    = "hoge"
	defaultDateStart   = "2018"
	defaultDateEnd     = "2019"
	invalidDate        = "hoge"
)

func TestIndicatorValuesService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

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
		indicatorID string
		datePatams  *DateParams
		pages       *PageParams
	}
	tests := []struct {
		name                     string
		args                     args
		want                     *PageSummaryWithSourceID
		wantIndicatorValuesCount int
		wantErr                  bool
		wantErrRes               *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				indicatorID: defaultIndicatorID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  false,
			wantErrRes:               nil,
		},
		{
			name: "success with params",
			args: args{
				indicatorID: defaultIndicatorID,
				datePatams:  defaultDateParams,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want: &PageSummaryWithSourceID{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantIndicatorValuesCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:                  false,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid indicator id",
			args: args{
				indicatorID: invalidIndicatorID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/all/indicators/%s?format=json",
					defaultBaseURL,
					apiVersion,
					invalidIndicatorID,
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
			name: "failure because invalid date params",
			args: args{
				indicatorID: defaultIndicatorID,
				datePatams:  invalidDateParams,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				indicatorID: defaultIndicatorID,
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
			got, got1, err := i.List(tt.args.indicatorID, tt.args.datePatams, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.List() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("IndicatorValuesService.List() got = %v, want %v", got, tt.want)
				}
				if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
					t.Errorf(
						"IndicatorValuesService.List() reflect.TypeOf(got) = %v, reflect.TypeOf(tt.want) %v",
						reflect.TypeOf(got),
						reflect.TypeOf(tt.want),
					)
				}
				if len(got1) != tt.wantIndicatorValuesCount {
					t.Errorf("invalid length of IndicatorValuesService.List() got1 = %v, want %v", got1, tt.wantIndicatorValuesCount)
				}
				if !tt.wantErr {
					for i := range got1 {
						if got1[i].Date != defaultDateStart && got1[i].Date != defaultDateEnd {
							t.Errorf(
								"invalid date of IndicatorValuesService.List() got1 = %v, want %v or %v",
								got1[i].Date,
								defaultDateStart,
								defaultDateEnd,
							)
						}
						if got1[i].Indicator.ID != defaultIndicatorID {
							t.Errorf(
								"invalid indicator ID of IndicatorValuesService.List() got1[i].Indicator.ID = %v, want %v",
								got1[i].Indicator.ID,
								defaultIndicatorID,
							)
						}
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListByCountryIDs(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	defaultCountryIDs := []string{"JPN", "USA"}
	invalidCountryIDs := []string{"ABCDEFG", "HIJKLM"}
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
		want                     *PageSummaryWithSourceID
		wantIndicatorValuesCount int
		wantErr                  bool
		wantErrRes               *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				countryIDs:  defaultCountryIDs,
				indicatorID: defaultIndicatorID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  false,
			wantErrRes:               nil,
		},
		{
			name: "success with params",
			args: args{
				countryIDs:  defaultCountryIDs,
				indicatorID: defaultIndicatorID,
				datePatams:  defaultDateParams,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want: &PageSummaryWithSourceID{
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
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?format=json",
					defaultBaseURL,
					apiVersion,
					strings.Join(invalidCountryIDs, ";"),
					defaultIndicatorID,
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
				countryIDs:  defaultCountryIDs,
				indicatorID: invalidIndicatorID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?format=json",
					defaultBaseURL,
					apiVersion,
					strings.Join(defaultCountryIDs, ";"),
					invalidIndicatorID,
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
			name: "failure because invalid date params",
			args: args{
				countryIDs:  defaultCountryIDs,
				indicatorID: defaultIndicatorID,
				datePatams:  invalidDateParams,
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
				if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
					t.Errorf(
						"IndicatorValuesService.ListByCountryIDs() reflect.TypeOf(got) = %v, reflect.TypeOf(tt.want) %v",
						reflect.TypeOf(got),
						reflect.TypeOf(tt.want),
					)
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
			}
		})
	}
}

func TestIndicatorValuesService_ListBySourceID(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	defaultIndicatorIDs := []string{"NY.GDP.MKTP.CD", "SP.POP.TOTL"}
	invalidIndicatorIDs := []string{"INVALID.INDICATOR.ID", "SP.POP.TOTL"}
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
		indicatorIDs []string
		sourceID     string
		datePatams   *DateParams
		pages        *PageParams
	}
	tests := []struct {
		name                     string
		args                     args
		want                     *PageSummaryWithLastUpdated
		wantIndicatorValuesCount int
		wantErr                  bool
		wantErrRes               *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     defaultSourceID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  false,
			wantErrRes:               nil,
		},
		{
			name: "success with params",
			args: args{
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     defaultSourceID,
				datePatams:   defaultDateParams,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want: &PageSummaryWithLastUpdated{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantIndicatorValuesCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:                  false,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid indicator id",
			args: args{
				indicatorIDs: invalidIndicatorIDs,
				sourceID:     defaultSourceID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/all/indicators/%s?format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(invalidIndicatorIDs, ";"),
					defaultSourceID,
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
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     invalidSourceID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid date params",
			args: args{
				indicatorIDs: invalidIndicatorIDs,
				sourceID:     defaultSourceID,
				datePatams:   invalidDateParams,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     defaultSourceID,
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
			got, got1, err := i.ListBySourceID(tt.args.indicatorIDs, tt.args.sourceID, tt.args.datePatams, tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListBySourceID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListBySourceID() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("IndicatorValuesService.ListBySourceID() got = %v, want %v", got, tt.want)
				}
				if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
					t.Errorf(
						"IndicatorValuesService.ListBySourceID() reflect.TypeOf(got) = %v, reflect.TypeOf(tt.want) %v",
						reflect.TypeOf(got),
						reflect.TypeOf(tt.want),
					)
				}
				if len(got1) != tt.wantIndicatorValuesCount {
					t.Errorf("invalid length of IndicatorValuesService.ListBySourceID() got1 = %v, want %v", got1, tt.wantIndicatorValuesCount)
				}
				if !tt.wantErr {
					for i := range got1 {
						if got1[i].Date != defaultDateStart && got1[i].Date != defaultDateEnd {
							t.Errorf(
								"invalid date of IndicatorValuesService.ListBySourceID() got1 = %v, want %v or %v",
								got1[i].Date,
								defaultDateStart,
								defaultDateEnd,
							)
						}
					}
				}
			}
		})
	}
}

func TestIndicatorValuesService_ListByCountryIDsAndSourceID(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	defaultCountryIDs := []string{"JPN", "USA"}
	invalidCountryIDs := []string{"ABCDEFG", "HIJKLM"}
	defaultIndicatorIDs := []string{"NY.GDP.MKTP.CD", "SP.POP.TOTL"}
	invalidIndicatorIDs := []string{"INVALID.INDICATOR.ID", "SP.POP.TOTL"}
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
		countryIDs   []string
		indicatorIDs []string
		sourceID     string
		datePatams   *DateParams
		pages        *PageParams
	}
	tests := []struct {
		name                     string
		args                     args
		want                     *PageSummaryWithLastUpdated
		wantIndicatorValuesCount int
		wantErr                  bool
		wantErrRes               *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				countryIDs:   defaultCountryIDs,
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     defaultSourceID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  false,
			wantErrRes:               nil,
		},
		{
			name: "success with params",
			args: args{
				countryIDs:   defaultCountryIDs,
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     defaultSourceID,
				datePatams:   defaultDateParams,
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want: &PageSummaryWithLastUpdated{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantIndicatorValuesCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:                  false,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid country ids",
			args: args{
				countryIDs:   invalidCountryIDs,
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     defaultSourceID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(invalidCountryIDs, ";"),
					strings.Join(defaultIndicatorIDs, ";"),
					defaultSourceID,
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
				countryIDs:   defaultCountryIDs,
				indicatorIDs: invalidIndicatorIDs,
				sourceID:     defaultSourceID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s/indicators/%s?format=json&source=%s",
					defaultBaseURL,
					apiVersion,
					strings.Join(defaultCountryIDs, ";"),
					strings.Join(invalidIndicatorIDs, ";"),
					defaultSourceID,
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
				countryIDs:   defaultCountryIDs,
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     invalidSourceID,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid date params",
			args: args{
				countryIDs:   defaultCountryIDs,
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     defaultSourceID,
				datePatams:   invalidDateParams,
			},
			want:                     nil,
			wantIndicatorValuesCount: 0,
			wantErr:                  true,
			wantErrRes:               nil,
		},
		{
			name: "failure because invalid page params",
			args: args{
				countryIDs:   defaultCountryIDs,
				indicatorIDs: defaultIndicatorIDs,
				sourceID:     defaultSourceID,
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
			got, got1, err := i.ListByCountryIDsAndSourceID(
				tt.args.countryIDs,
				tt.args.indicatorIDs,
				tt.args.sourceID,
				tt.args.datePatams,
				tt.args.pages,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndicatorValuesService.ListByCountryIDsAndSourceID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErrRes != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("IndicatorValuesService.ListByCountryIDsAndSourceID() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("IndicatorValuesService.ListByCountryIDsAndSourceID() got = %v, want %v", got, tt.want)
				}
				if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
					t.Errorf(
						"IndicatorValuesService.ListByCountryIDsAndSourceID() reflect.TypeOf(got) = %v, reflect.TypeOf(tt.want) %v",
						reflect.TypeOf(got),
						reflect.TypeOf(tt.want),
					)
				}
				if len(got1) != tt.wantIndicatorValuesCount {
					t.Errorf(
						"invalid length of IndicatorValuesService.ListByCountryIDsAndSourceID() got1 = %v, want %v",
						got1,
						tt.wantIndicatorValuesCount,
					)
				}
				if !tt.wantErr {
					for i := range got1 {
						if got1[i].Date != defaultDateStart && got1[i].Date != defaultDateEnd {
							t.Errorf(
								"invalid date of IndicatorValuesService.ListByCountryIDsAndSourceID() got1 = %v, want %v or %v",
								got1[i].Date,
								defaultDateStart,
								defaultDateEnd,
							)
						}
					}
				}
			}
		})
	}
}
