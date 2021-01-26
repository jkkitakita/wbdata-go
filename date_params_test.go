package wbdata

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestFilterParams_addFilterParams(t *testing.T) {
	defaultFormatStr := "format=" + defaultFormat.String()
	baseURLStr := defaultBaseURL + apiVersion
	baseURL, _ := url.Parse(baseURLStr + "?" + defaultFormatStr)

	invalidYearStr := "2018-01"
	defaultYearStr2018 := "2018"
	defaultYearStr2019 := "2019"
	defaultMonthStr2018 := "2018M01"
	defaultMonthStr2019 := "2019M01"
	defaultQuarterStr2018 := "2018Q01"
	defaultQuarterStr2019 := "2019Q01"

	expectedDefaultURL, _ := url.Parse(baseURLStr + "?" + defaultFormatStr)
	expectedYearURL, _ := url.Parse(baseURLStr + "?date=" + defaultYearStr2018 + "&" + defaultFormatStr)
	expectedYTDURL, _ := url.Parse(baseURLStr + "?date=YTD%3A" + defaultYearStr2018 + "&" + defaultFormatStr)
	expectedMonthURL, _ := url.Parse(baseURLStr + "?date=" + defaultMonthStr2018 + "&" + defaultFormatStr)
	expectedMonth2018To2019URL, _ := url.Parse(
		baseURLStr + "?date=" + defaultMonthStr2018 + "%3A" + defaultMonthStr2019 + "&" + defaultFormatStr,
	)
	expectedQuarterURL, _ := url.Parse(baseURLStr + "?date=" + defaultQuarterStr2018 + "&" + defaultFormatStr)
	expectedQuarter2018To2019URL, _ := url.Parse(
		baseURLStr + "?date=" + defaultQuarterStr2018 + "%3A" + defaultQuarterStr2019 + "&" + defaultFormatStr,
	)
	expectedFrequencyYearMRVURL, _ := url.Parse(baseURLStr + "?" + defaultFormatStr + "&frequency=Y&mrv=1")
	expectedFrequencyYearMRNEVURL, _ := url.Parse(baseURLStr + "?" + defaultFormatStr + "&frequency=Y&mrnev=1")
	expectedFrequencyYearGapFillURL, _ := url.Parse(baseURLStr + "?" + defaultFormatStr + "&frequency=Y&gapfill=Y&mrv=1")
	expectedFrequencyMonthURL, _ := url.Parse(baseURLStr + "?" + defaultFormatStr + "&frequency=M&mrv=1")
	expectedFrequencyQuarterURL, _ := url.Parse(baseURLStr + "?" + defaultFormatStr + "&frequency=Q&mrv=1")

	type fields struct {
		FilterParamsType FilterParamsType
		DateParam        *DateParam
		RecentParam      *RecentParam
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  *fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name:   "success with nil fields",
			fields: nil,
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedDefaultURL,
			},
			wantErr: false,
		},
		{
			name: "success with FilterParamsDate and date of yearly",
			fields: &fields{
				FilterParamsType: FilterParamsDate,
				DateParam: &DateParam{
					Date: defaultYearStr2018,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedYearURL,
			},
			wantErr: false,
		},
		{
			name: "success with FilterParamsDate and date of monthly",
			fields: &fields{
				FilterParamsType: FilterParamsDate,
				DateParam: &DateParam{
					Date: defaultMonthStr2018,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedMonthURL,
			},
			wantErr: false,
		},
		{
			name: "success with FilterParamsDate and date of quarterly",
			fields: &fields{
				FilterParamsType: FilterParamsDate,
				DateParam: &DateParam{
					Date: defaultQuarterStr2018,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedQuarterURL,
			},
			wantErr: false,
		},
		{
			name: "failure with FilterParamsDate because DateRange is specified",
			fields: &fields{
				FilterParamsType: FilterParamsDate,
				DateParam: &DateParam{
					Date: defaultYearStr2018,
					DateRange: &DateRange{
						Start: defaultYearStr2018,
						End:   defaultYearStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDate because Date is not found",
			fields: &fields{
				FilterParamsType: FilterParamsDate,
				DateParam: &DateParam{
					Date: "",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDate because Date is invalid",
			fields: &fields{
				FilterParamsType: FilterParamsDate,
				DateParam: &DateParam{
					Date: invalidYearStr,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDate because RecentParam is specified",
			fields: &fields{
				FilterParamsType: FilterParamsDate,
				DateParam: &DateParam{
					Date: defaultYearStr2018,
				},
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success with FilterParamsDateRange",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: defaultYearStr2018,
						End:   defaultYearStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "success with FilterParamsDateRange M",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: defaultMonthStr2018,
						End:   defaultMonthStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedMonth2018To2019URL,
			},
			wantErr: false,
		},
		{
			name: "success with FilterParamsDateRange Q",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: defaultQuarterStr2018,
						End:   defaultQuarterStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedQuarter2018To2019URL,
			},
			wantErr: false,
		},
		{
			name: "failure with FilterParamsDateRange because Date is specified",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					Date: defaultYearStr2018,
					DateRange: &DateRange{
						Start: defaultYearStr2018,
						End:   defaultYearStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDateRange because DateRange is nil",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: nil,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDateRange because DateRange.Start is invalid",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: invalidYearStr,
						End:   defaultYearStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDateRange because DateRange.End is invalid",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: defaultYearStr2018,
						End:   invalidYearStr,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDateRange because start should be before end",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: defaultYearStr2019,
						End:   defaultYearStr2018,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDateRange because both monthly and quarterly is used",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: defaultMonthStr2018,
						End:   defaultQuarterStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDateRange because both yearly and quarterly is used",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: defaultYearStr2018,
						End:   defaultQuarterStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsDateRange because RecentParam is specified",
			fields: &fields{
				FilterParamsType: FilterParamsDateRange,
				DateParam: &DateParam{
					DateRange: &DateRange{
						Start: defaultYearStr2018,
						End:   defaultYearStr2019,
					},
				},
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success with FilterParamsYearToDate",
			fields: &fields{
				FilterParamsType: FilterParamsYearToDate,
				DateParam: &DateParam{
					Date: defaultYearStr2018,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedYTDURL,
			},
			wantErr: false,
		},
		{
			name: "failure with FilterParamsYearToDate because DateRange is specified",
			fields: &fields{
				FilterParamsType: FilterParamsYearToDate,
				DateParam: &DateParam{
					Date: defaultYearStr2018,
					DateRange: &DateRange{
						Start: defaultYearStr2018,
						End:   defaultYearStr2019,
					},
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsYearToDate because Date is not found",
			fields: &fields{
				FilterParamsType: FilterParamsYearToDate,
				DateParam: &DateParam{
					Date: "",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsYearToDate because Date is invalid",
			fields: &fields{
				FilterParamsType: FilterParamsYearToDate,
				DateParam: &DateParam{
					Date: invalidYearStr,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsYearToDate because Date is monthly",
			fields: &fields{
				FilterParamsType: FilterParamsYearToDate,
				DateParam: &DateParam{
					Date: defaultMonthStr2018,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsYearToDate because Date is quarterly",
			fields: &fields{
				FilterParamsType: FilterParamsYearToDate,
				DateParam: &DateParam{
					Date: defaultQuarterStr2018,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FilterParamsYearToDate because RecentParam is specified",
			fields: &fields{
				FilterParamsType: FilterParamsYearToDate,
				DateParam: &DateParam{
					Date: defaultYearStr2018,
				},
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success with FilterParamsMRV with Yearly",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedFrequencyYearMRVURL,
			},
			wantErr: false,
		},
		{
			name: "success with FilterParamsMRV with Monthly",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyMonthly,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedFrequencyMonthURL,
			},
			wantErr: false,
		},
		{
			name: "success with FilterParamsMRV with Quarterly",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyQuarterly,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedFrequencyQuarterURL,
			},
			wantErr: false,
		},
		{
			name: "success with FilterParamsMRV with IsNotEmpty",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 1,
					IsNotEmpty:       true,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedFrequencyYearMRNEVURL,
			},
			wantErr: false,
		},
		{
			name: "success with FilterParamsMRV with IsGapFill",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        true,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want: &http.Request{
				URL: expectedFrequencyYearGapFillURL,
			},
			wantErr: false,
		},
		{
			name: "failure with FilterParamsUnknown",
			fields: &fields{
				FilterParamsType: FilterParamsUnknown,
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with DateParam",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				DateParam: &DateParam{
					Date: defaultYearStr2018,
				},
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure without RecentParam",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with FrequencyUnknown",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyUnknown,
					MostRecentValues: 1,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with zero MostRecentValues",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 0,
					IsNotEmpty:       false,
					IsGapFill:        false,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failure with both IsNotEmpty and IsGapFill is true",
			fields: &fields{
				FilterParamsType: FilterParamsMRV,
				RecentParam: &RecentParam{
					FrequencyType:    FrequencyYearly,
					MostRecentValues: 1,
					IsNotEmpty:       true,
					IsGapFill:        true,
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.args.req.URL, _ = url.Parse(baseURLStr + "?" + defaultFormatStr)
		t.Run(tt.name, func(t *testing.T) {
			var dp *FilterParams
			if tt.fields != nil {
				dp = &FilterParams{
					FilterParamsType: tt.fields.FilterParamsType,
					DateParam:        tt.fields.DateParam,
					RecentParam:      tt.fields.RecentParam,
				}
			} else {
				dp = nil
			}
			if err := dp.addFilterParams(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("FilterParams.addFilterParams() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != nil && !reflect.DeepEqual(tt.args.req, tt.want) {
				t.Errorf("FilterParams.addFilterParams() tt.args.req = %v, want %v", tt.args.req, tt.want)
			}
		})
	}
}
