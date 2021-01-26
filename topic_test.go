package wbdata

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jkkitakita/wbdata-go/testutils"
)

func TestTopicsService_List(t *testing.T) {
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
		pages *PageParams
	}
	tests := []struct {
		name    string
		args    args
		want    *PageSummary
		want1   []*Topic
		wantErr bool
	}{
		{
			name: "success",
			args: args{},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   21,
			},
			want1:   nil,
			wantErr: false,
		},
		{
			name: "success with page params",
			args: args{
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   11,
				PerPage: 2,
				Total:   21,
			},
			want1: []*Topic{
				{
					ID:    "1",
					Value: "Agriculture & Rural Development",
					SourceNote: "For the 70 percent of the world's poor who " +
						"live in rural areas, agriculture is the main source of income and employment. " +
						"But depletion and degradation of land and water pose serious challenges to producing " +
						"enough food and other agricultural products to sustain livelihoods here " +
						"and meet the needs of urban populations. Data presented here include measures of " +
						"agricultural inputs, outputs, and productivity compiled by the UN's Food and Agriculture Organization.",
				},
				{
					ID:    "2",
					Value: "Aid Effectiveness",
					SourceNote: "Aid effectiveness is the impact that aid has " +
						"in reducing poverty and inequality, increasing growth, building capacity, and " +
						"accelerating achievement of the Millennium Development Goals set by the international community. " +
						"Indicators here cover aid received as well as progress in reducing poverty and " +
						"improving education, health, and other measures of human welfare.",
				},
			},
			wantErr: false,
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
			topicService := &TopicsService{
				client: client,
			}
			got, got1, err := topicService.List(tt.args.pages)
			if (err != nil) != tt.wantErr {
				t.Errorf("TopicsService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopicsService.List() got = %v, want %v", got, tt.want)
			}
			if tt.want1 != nil {
				for i := range got1 {
					if !reflect.DeepEqual(got1[i], tt.want1[i]) {
						t.Errorf("TopicsService.List() got1[i] = %v, want[i] %v", got1[i], tt.want1[i])
					}
				}
			}
		})
	}
}

func TestTopicsService_Get(t *testing.T) {
	client, save := NewTestClient(t, *update)
	defer save()

	type args struct {
		topicID string
	}
	tests := []struct {
		name       string
		args       args
		want       *PageSummary
		want1      *Topic
		wantErr    bool
		wantErrRes *ErrorResponse
	}{
		{
			name: "success",
			args: args{
				topicID: testutils.TestDefaultTopicID,
			},
			want: &PageSummary{
				Page:    1,
				Pages:   1,
				PerPage: 50,
				Total:   1,
			},
			want1: &Topic{
				ID:    "1",
				Value: "Agriculture & Rural Development",
				SourceNote: "For the 70 percent of the world's poor who " +
					"live in rural areas, agriculture is the main source of income and employment. " +
					"But depletion and degradation of land and water pose serious challenges to producing " +
					"enough food and other agricultural products to sustain livelihoods here " +
					"and meet the needs of urban populations. Data presented here include measures of " +
					"agricultural inputs, outputs, and productivity compiled by the UN's Food and Agriculture Organization.",
			},
			wantErr:    false,
			wantErrRes: nil,
		},
		{
			name: "failure because topicID is invalid",
			args: args{
				topicID: testutils.TestInvalidTopicID,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
			wantErrRes: &ErrorResponse{
				URL: fmt.Sprintf(
					"%s%s/topics/%s?format=json",
					defaultBaseURL,
					apiVersion,
					testutils.TestInvalidTopicID,
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
			topicService := &TopicsService{
				client: client,
			}
			got, got1, err := topicService.Get(tt.args.topicID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TopicsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopicsService.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TopicsService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
