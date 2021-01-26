package wbdata

import (
	"net/http"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestPageParams_addPageParams(t *testing.T) {
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
		fields  *fields
		wantErr bool
	}{
		{
			name:    "success",
			fields:  nil,
			wantErr: false,
		},
		{
			name: "success with pages",
			fields: &fields{
				Page:    testutils.TestDefaultPage,
				PerPage: testutils.TestDefaultPerPage,
			},
			wantErr: false,
		},
		{
			name: "failure because Page is less than 1",
			fields: &fields{
				Page:    testutils.TestInvalidPage,
				PerPage: testutils.TestDefaultPerPage,
			},
			wantErr: true,
		},
		{
			name: "failure because PerPage is less than 1",
			fields: &fields{
				Page:    testutils.TestDefaultPage,
				PerPage: testutils.TestInvalidPerPage,
			},
			wantErr: true,
		},
		{
			name: "failure because both Page and PerPage is less than 1",
			fields: &fields{
				Page:    testutils.TestInvalidPage,
				PerPage: testutils.TestInvalidPerPage,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pages *PageParams
			if tt.fields != nil {
				pages = &PageParams{
					Page:    tt.fields.Page,
					PerPage: tt.fields.PerPage,
				}
			} else {
				pages = nil
			}
			if err := pages.addPageParams(mockReq); (err != nil) != tt.wantErr {
				t.Errorf("PageParams.addPageParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
