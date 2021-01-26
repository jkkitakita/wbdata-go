package wbdata

import (
	"flag"
	"fmt"
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

	defaultPageParams := &PageParams{
		Page:    testutils.TestDefaultPage,
		PerPage: testutils.TestDefaultPerPage,
	}
	invalidPageParams := &PageParams{
		Page:    testutils.TestInvalidPage,
		PerPage: testutils.TestDefaultPerPage,
	}

	type args struct {
		params *ListCountryParams
		pages  *PageParams
	}
	tests := []struct {
		name    string
		args    args
		want    *PageSummary
		want1   []*Country
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				params: nil,
				pages:  defaultPageParams,
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			want1: []*Country{
				{
					ID:          "ABW",
					Name:        "Aruba",
					CapitalCity: "Oranjestad",
					Iso2Code:    "AW",
					Longitude:   "-70.0167",
					Latitude:    "12.5167",
					Region: CountryRegion{
						ID:       "LCN",
						Iso2Code: "ZJ",
						Value:    "Latin America & Caribbean ",
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
				{
					ID:          "AFG",
					Name:        "Afghanistan",
					CapitalCity: "Kabul",
					Iso2Code:    "AF",
					Longitude:   "69.1761",
					Latitude:    "34.5228",
					Region: CountryRegion{
						ID:       "SAS",
						Iso2Code: "8S",
						Value:    "South Asia",
					},
					IncomeLevel: IncomeLevel{
						ID:       "LIC",
						Iso2Code: "XM",
						Value:    "Low income",
					},
					LendingType: LendingType{
						ID:       "IDX",
						Iso2Code: "XI",
						Value:    "IDA",
					},
					AdminRegion: CountryRegion{
						ID:       "SAS",
						Iso2Code: "8S",
						Value:    "South Asia",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success with region id",
			args: args{
				params: &ListCountryParams{
					RegionID: "EAS",
				},
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			want1: []*Country{
				{
					ID:          "ASM",
					Name:        "American Samoa",
					CapitalCity: "Pago Pago",
					Iso2Code:    "AS",
					Longitude:   "-170.691",
					Latitude:    "-14.2846",
					Region: CountryRegion{
						ID:       "EAS",
						Iso2Code: "Z4",
						Value:    "East Asia & Pacific",
					},
					IncomeLevel: IncomeLevel{
						ID:       "UMC",
						Iso2Code: "XT",
						Value:    "Upper middle income",
					},
					LendingType: LendingType{
						ID:       "LNX",
						Iso2Code: "XX",
						Value:    "Not classified",
					},
					AdminRegion: CountryRegion{
						ID:       "EAP",
						Iso2Code: "4E",
						Value:    "East Asia & Pacific (excluding high income)",
					},
				},
				{
					ID:          "AUS",
					Name:        "Australia",
					CapitalCity: "Canberra",
					Iso2Code:    "AU",
					Longitude:   "149.129",
					Latitude:    "-35.282",
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
			},
			wantErr: false,
		},
		{
			name: "success with income level id",
			args: args{
				params: &ListCountryParams{
					IncomeLevelID: "HIC",
				},
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			want1: []*Country{
				{
					ID:          "ABW",
					Name:        "Aruba",
					CapitalCity: "Oranjestad",
					Iso2Code:    "AW",
					Longitude:   "-70.0167",
					Latitude:    "12.5167",
					Region: CountryRegion{
						ID:       "LCN",
						Iso2Code: "ZJ",
						Value:    "Latin America & Caribbean ",
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
				{
					ID:          "AND",
					Name:        "Andorra",
					CapitalCity: "Andorra la Vella",
					Iso2Code:    "AD",
					Longitude:   "1.5218",
					Latitude:    "42.5075",
					Region: CountryRegion{
						ID:       "ECS",
						Iso2Code: "Z7",
						Value:    "Europe & Central Asia",
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
			},
			wantErr: false,
		},
		{
			name: "success with lending type id",
			args: args{
				params: &ListCountryParams{
					LendingTypeID: "LNX",
				},
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			want1: []*Country{
				{
					ID:          "ABW",
					Name:        "Aruba",
					CapitalCity: "Oranjestad",
					Iso2Code:    "AW",
					Longitude:   "-70.0167",
					Latitude:    "12.5167",
					Region: CountryRegion{
						ID:       "LCN",
						Iso2Code: "ZJ",
						Value:    "Latin America & Caribbean ",
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
				{
					ID:          "AND",
					Name:        "Andorra",
					CapitalCity: "Andorra la Vella",
					Iso2Code:    "AD",
					Longitude:   "1.5218",
					Latitude:    "42.5075",
					Region: CountryRegion{
						ID:       "ECS",
						Iso2Code: "Z7",
						Value:    "Europe & Central Asia",
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
			},
			wantErr: false,
		},
		{
			name: "success with params",
			args: args{
				params: &ListCountryParams{
					RegionID:      "EAS",
					IncomeLevelID: "HIC",
					LendingTypeID: "LNX",
				},
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			want1: []*Country{
				{
					ID:          "AUS",
					Name:        "Australia",
					CapitalCity: "Canberra",
					Iso2Code:    "AU",
					Longitude:   "149.129",
					Latitude:    "-35.282",
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
				{
					ID:          "BRN",
					Name:        "Brunei Darussalam",
					CapitalCity: "Bandar Seri Begawan",
					Iso2Code:    "BN",
					Longitude:   "114.946",
					Latitude:    "4.94199",
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
			},
			wantErr: false,
		},
		{
			name: "failure because invalid region id",
			args: args{
				params: &ListCountryParams{
					RegionID: testutils.TestInvalidRegionCode,
				},
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "failure because invalid income level id",
			args: args{
				params: &ListCountryParams{
					IncomeLevelID: testutils.TestInvalidIncomeLevelID,
				},
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "failure because invalid lending type id",
			args: args{
				params: &ListCountryParams{
					LendingTypeID: testutils.TestInvalidLendingTypeID,
				},
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: invalidPageParams,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
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

			for i := range got1 {
				if !reflect.DeepEqual(got1[i], tt.want1[i]) {
					t.Errorf("CountriesService.List() got1 = %v, want %v", got1[i], tt.want1[i])
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
				countryID: testutils.TestDefaultCountryID,
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
				countryID: testutils.TestInvalidCountryID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/countries/%s?format=json",
					defaultBaseURL,
					apiVersion,
					testutils.TestInvalidCountryID,
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
