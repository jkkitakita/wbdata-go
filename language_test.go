package wbdata

import (
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestLanguagesService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		pages *PageParams
	}
	tests := []struct {
		name               string
		args               args
		want               *PageSummary
		wantLanguagesCount int
		wantErr            bool
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
			wantLanguagesCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:            false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: &PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:               nil,
			wantLanguagesCount: 0,
			wantErr:            true,
		},
		{
			name: "failure because PerPage is less than 1",
			args: args{
				pages: &PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: 0,
				},
			},
			want:               nil,
			wantLanguagesCount: 0,
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lang := &LanguagesService{
				client: client,
			}

			got, got1, err := lang.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("LanguagesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("LanguagesService.List() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantLanguagesCount {
				t.Errorf("LanguagesService.List() got1 = %v, want %v", got1, tt.wantLanguagesCount)
			}
		})
	}
}

func TestLanguagesService_Get(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		languageCode string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *Language
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				languageCode: "ja",
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &Language{
				Code:       "ja",
				Name:       "Japanese ",
				NativeForm: "日本語",
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because languageCode is invalid",
			args: args{
				languageCode: invalidID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL:  defaultBaseURL + apiVersion + "/languages/" + invalidID + "?format=json",
				Code: 200,
				Message: []ErrorMessage{
					{
						ID:    "150",
						Key:   "Language is not yet supported in the API",
						Value: "Response requested in an unsupported language.",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LanguagesService{
				client: client,
			}
			got, got1, err := c.Get(tt.args.languageCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("LanguagesService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("LanguagesService.Get() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LanguagesService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LanguagesService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
