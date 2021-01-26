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
		name            string
		args            args
		want            *PageSummary
		wantTopicsCount int
		wantErr         bool
	}{
		{
			name: "success",
			args: args{
				pages: defaultPageParams,
			},
			want: &PageSummary{
				Page:    intOrString(testutils.TestDefaultPage),
				PerPage: intOrString(testutils.TestDefaultPerPage),
			},
			wantTopicsCount: testutils.TestDefaultPage * testutils.TestDefaultPerPage,
			wantErr:         false,
		},
		{
			name: "failure because Page is less than 1",
			args: args{
				pages: invalidPageParams,
			},
			want:            nil,
			wantTopicsCount: 0,
			wantErr:         true,
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
			if tt.want != nil {
				if got.Page != tt.want.Page || got.PerPage != tt.want.PerPage {
					t.Errorf("TopicsService.List() got = %v, want %v", got, tt.want)
				}
			}
			if len(got1) != tt.wantTopicsCount {
				t.Errorf("TopicsService.List() got1 = %v, want %v", got1, tt.wantTopicsCount)
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
