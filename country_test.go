package wbdata

import (
	"flag"
	"os"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

const (
	invalidID = "ABCDEFG"
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
		params *ListCountryParams
		pages  *PageParams
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
				params: nil,
				pages: &PageParams{
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
			name: "success with region id",
			args: args{
				params: &ListCountryParams{
					RegionID: "EAS",
				},
				pages: &PageParams{
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
			name: "success with income level id",
			args: args{
				params: &ListCountryParams{
					IncomeLevelID: "HIC",
				},
				pages: &PageParams{
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
			name: "success with lending type id",
			args: args{
				params: &ListCountryParams{
					LendingTypeID: "LNX",
				},
				pages: &PageParams{
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
			name: "success with params",
			args: args{
				params: &ListCountryParams{
					RegionID:      "EAS",
					IncomeLevelID: "HIC",
					LendingTypeID: "LNX",
				},
				pages: &PageParams{
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
			name: "failure because invalid region id",
			args: args{
				params: &ListCountryParams{
					RegionID:      invalidID,
					IncomeLevelID: "",
					LendingTypeID: "",
				},
				pages: &PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:               nil,
			wantCountriesCount: 0,
			wantErr:            true,
		},
		{
			name: "failure because invalid income level id",
			args: args{
				params: &ListCountryParams{
					RegionID:      "",
					IncomeLevelID: invalidID,
					LendingTypeID: "",
				},
				pages: &PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:               nil,
			wantCountriesCount: 0,
			wantErr:            true,
		},
		{
			name: "failure because invalid lending type id",
			args: args{
				params: &ListCountryParams{
					RegionID:      "",
					IncomeLevelID: "",
					LendingTypeID: invalidID,
				},
				pages: &PageParams{
					Page:    0,
					PerPage: testutils.TestDefaultPerPage,
				},
			},
			want:               nil,
			wantCountriesCount: 0,
			wantErr:            true,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				params: &ListCountryParams{
					RegionID:      "",
					IncomeLevelID: "",
					LendingTypeID: "",
				},
				pages: &PageParams{
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
				params: &ListCountryParams{
					RegionID:      "",
					IncomeLevelID: "",
					LendingTypeID: "",
				},
				pages: &PageParams{
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

			got, got1, err := c.List(tt.args.params, tt.args.pages)
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

			for i := range got1 {
				if tt.args.params != nil {
					if tt.args.params.RegionID != "" && got1[i].Region.ID != tt.args.params.RegionID {
						t.Errorf("invalid region id. got1[i].Region.ID = %v, want %v", got1[i].Region.ID, tt.args.params.RegionID)
					}
					if tt.args.params.IncomeLevelID != "" && got1[i].IncomeLevel.ID != tt.args.params.IncomeLevelID {
						t.Errorf("invalid region id. got1[i].IncomeLevel.ID = %v, want %v", got1[i].IncomeLevel.ID, tt.args.params.IncomeLevelID)
					}
					if tt.args.params.LendingTypeID != "" && got1[i].LendingType.ID != tt.args.params.LendingTypeID {
						t.Errorf("invalid region id. got1[i].LendingType.ID = %v, want %v", got1[i].LendingType.ID, tt.args.params.LendingTypeID)
					}
				}
			}
		})
	}
}

func TestCountriesService_Get(t *testing.T) {
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
				countryID: invalidID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL:  defaultBaseURL + apiVersion + "/countries/" + invalidID + "?format=json",
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

func TestListCountryParams_toQueryParams(t *testing.T) {
	type fields struct {
		RegionID      string
		IncomeLevelID string
		LendingTypeID string
	}
	tests := []struct {
		name   string
		fields *fields
		want   map[string]string
	}{
		{
			name: "success",
			fields: &fields{
				RegionID:      "EAS",
				IncomeLevelID: "HIC",
				LendingTypeID: "LNX",
			},
			want: map[string]string{
				"region":      "EAS",
				"incomelevel": "HIC",
				"lendingtype": "LNX",
			},
		},
		{
			name:   "success with nil",
			fields: nil,
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var params *ListCountryParams
			if tt.fields != nil {
				params = &ListCountryParams{
					RegionID:      tt.fields.RegionID,
					IncomeLevelID: tt.fields.IncomeLevelID,
					LendingTypeID: tt.fields.LendingTypeID,
				}
			} else {
				params = nil
			}
			if got := params.toQueryParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCountryParams.toQueryParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
