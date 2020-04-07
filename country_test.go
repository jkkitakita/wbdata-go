package wbdata

import (
	"flag"
	"os"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

var (
	update = flag.Bool("update", false, "update fixtures")
)

func TestMain(m *testing.M) {
	flag.Parse()
	testutils.UpdateFixture(update)
	os.Exit(m.Run())
}

func TestCountriesService_List(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		pages PageParams
	}
	tests := []struct {
		name               string
		args               args
		want               *PageSummary
		wantCountriesCount int
		wantErr            bool
	}{
		{
			name: "success",
			args: args{
				pages: PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantCountriesCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:            false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:               nil,
			wantCountriesCount: 0,
			wantErr:            true,
		},
		{
			name: "failure because PerPage is less than 1",
			args: args{
				pages: PageParams{
					Page:    testutils.TestDefaultPage,
					PerPage: 0,
				},
			},
			want:               nil,
			wantCountriesCount: 0,
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CountriesService{
				client: client,
			}

			got, got1, err := c.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountriesService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("CountriesService.List() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantCountriesCount {
				t.Errorf("CountriesService.List() got1 = %v, want %v", got1, tt.wantCountriesCount)
			}
		})
	}
}

func TestCountriesService_Get(t *testing.T) {
	invalidCountryID := "ABCDEFG"

	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		countryID string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *Country
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				countryID: "jpn",
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &Country{
				ID:          "JPN",
				Name:        "Japan",
				CapitalCity: "Tokyo",
				Iso2Code:    "JP",
				Longitude:   "139.77",
				Latitude:    "35.67",
				Region: CountryRegion{
					ID:       "EAS",
					Iso2Code: "Z4",
					Value:    "East Asia & Pacific",
				},
				IncomeLevel: IncomeLevel{
					ID:       "HIC",
					Iso2Code: "XD",
					Value:    "High income",
				},
				LendingType: LendingType{
					ID:       "LNX",
					Iso2Code: "XX",
					Value:    "Not classified",
				},
				AdminRegion: CountryRegion{
					ID:       "",
					Iso2Code: "",
					Value:    "",
				},
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because countryID is invalid",
			args: args{
				countryID: invalidCountryID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL:  defaultBaseURL + apiVersion + "/countries/" + invalidCountryID + "?format=json",
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
			c := &CountriesService{
				client: client,
			}
			got, got1, err := c.Get(tt.args.countryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountriesService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				if !reflect.DeepEqual(err, tt.wantErrRes) {
					t.Errorf("CountriesService.Get() err = %v, wantErrRes %v", err, tt.wantErrRes)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountriesService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CountriesService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
