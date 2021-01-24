package wbdata

import (
	"net/http"
	"net/url"
	"testing"
)

func TestDateParams_addDateParams(t *testing.T) {
	baseURL, _ := url.Parse(defaultBaseURL + apiVersion)

	type fields struct {
		DateParamsType DateParamsType
		Date           string
		DateRange      *DateRange
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  *fields
		args    args
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
			wantErr: false,
		},
		{
			name: "success with DateParamsDate and date of yearly",
			fields: &fields{
				DateParamsType: DateParamsDate,
				Date:           "2018",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "success with DateParamsDate and date of monthly",
			fields: &fields{
				DateParamsType: DateParamsDate,
				Date:           "2018M01",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "success with DateParamsDate and date of quarterly",
			fields: &fields{
				DateParamsType: DateParamsDate,
				Date:           "2018Q01",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "failure with DateParamsDate because DateRange is specified",
			fields: &fields{
				DateParamsType: DateParamsDate,
				Date:           "2018",
				DateRange: &DateRange{
					Start: "2018",
					End:   "2019",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsDate because Date is not found",
			fields: &fields{
				DateParamsType: DateParamsDate,
				Date:           "",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsDate because Date is invalid",
			fields: &fields{
				DateParamsType: DateParamsDate,
				Date:           "hoge",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "success with DateParamsRange",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange: &DateRange{
					Start: "2018",
					End:   "2019",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "success with DateParamsRange M",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange: &DateRange{
					Start: "2018M01",
					End:   "2019M01",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "success with DateParamsRange Q",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange: &DateRange{
					Start: "2018Q01",
					End:   "2019Q01",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "failure with DateParamsRange because Date is specified",
			fields: &fields{
				DateParamsType: DateParamsRange,
				Date:           "2018",
				DateRange: &DateRange{
					Start: "2018",
					End:   "2019",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsRange because DateRange is nil",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange:      nil,
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsRange because DateRange.Start is invalid",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange: &DateRange{
					Start: "hoge",
					End:   "2019",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsRange because DateRange.End is invalid",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange: &DateRange{
					Start: "2018",
					End:   "hoge",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsRange because start should be before end",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange: &DateRange{
					Start: "2019",
					End:   "2018",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsRange because both monthly and quarterly is used",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange: &DateRange{
					Start: "2018M01",
					End:   "2019Q01",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsRange because both yearly and quarterly is used",
			fields: &fields{
				DateParamsType: DateParamsRange,
				DateRange: &DateRange{
					Start: "2018",
					End:   "2019Q01",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "success with DateParamsYearToDate",
			fields: &fields{
				DateParamsType: DateParamsYearToDate,
				Date:           "2018",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "failure with DateParamsYearToDate because DateRange is specified",
			fields: &fields{
				DateParamsType: DateParamsYearToDate,
				Date:           "2018",
				DateRange: &DateRange{
					Start: "2018",
					End:   "2019",
				},
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsYearToDate because Date is not found",
			fields: &fields{
				DateParamsType: DateParamsYearToDate,
				Date:           "",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsYearToDate because Date is invalid",
			fields: &fields{
				DateParamsType: DateParamsYearToDate,
				Date:           "hoge",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsYearToDate because Date is monthly",
			fields: &fields{
				DateParamsType: DateParamsYearToDate,
				Date:           "2018M01",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsYearToDate because Date is quarterly",
			fields: &fields{
				DateParamsType: DateParamsYearToDate,
				Date:           "2018Q01",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure with DateParamsUnknown",
			fields: &fields{
				DateParamsType: DateParamsUnknown,
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp *DateParams
			if tt.fields != nil {
				dp = &DateParams{
					DateParamsType: tt.fields.DateParamsType,
					Date:           tt.fields.Date,
					DateRange:      tt.fields.DateRange,
				}
			} else {
				dp = nil
			}
			if err := dp.addDateParams(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("DateParams.addDateParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
