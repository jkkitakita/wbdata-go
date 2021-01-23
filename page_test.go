package wbdata

import (
	"net/http"
	"testing"
)

func TestPageParams_addPageParams(t *testing.T) {
	invalidNumber := 0
	validNumber := 1

	mockReq, err := http.NewRequest("GET", "/hoge", nil)
	if err != nil {
		t.Error("failed to generate Mock http.NewRequest")
	}

	type fields struct {
		Page    int
		PerPage int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Page:    validNumber,
				PerPage: validNumber,
			},
			wantErr: false,
		},
		{
			name: "failure because Page is less than 1",
			fields: fields{
				Page:    invalidNumber,
				PerPage: validNumber,
			},
			wantErr: true,
		},
		{
			name: "failure because PerPage is less than 1",
			fields: fields{
				Page:    validNumber,
				PerPage: invalidNumber,
			},
			wantErr: true,
		},
		{
			name: "failure because both Page and PerPage is less than 1",
			fields: fields{
				Page:    invalidNumber,
				PerPage: invalidNumber,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pages := &PageParams{
				Page:    tt.fields.Page,
				PerPage: tt.fields.PerPage,
			}
			if err := pages.addPageParams(mockReq); (err != nil) != tt.wantErr {
				t.Errorf("PageParams.addPageParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
