package wbdata

import (
	"net/http"
	"net/url"
	"testing"
)

func TestDateParams_addDateParams(t *testing.T) {
	baseURL, _ := url.Parse(defaultBaseURL + apiVersion)

	type fields struct {
		Start string
		End   string
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Start: "2018",
				End:   "2019",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "success with same time",
			fields: fields{
				Start: "2018",
				End:   "2018",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: false,
		},
		{
			name: "failure because invalid start",
			fields: fields{
				Start: "hoge",
				End:   "2018",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure because invalid end",
			fields: fields{
				Start: "2018",
				End:   "hoge",
			},
			args: args{
				req: &http.Request{
					URL: baseURL,
				},
			},
			wantErr: true,
		},
		{
			name: "failure because start should be before end",
			fields: fields{
				Start: "2018",
				End:   "2017",
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
			dp := &DateParams{
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			if err := dp.addDateParams(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("DateParams.addDateParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
