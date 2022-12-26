package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/mocks"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
)

func TestCreateEvents(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(mockCtrl)
	eventService := NewEventService(mockRepo)

	type args struct {
		ctx      context.Context
		eventReq models.Event
	}
	tests := []struct {
		name        string
		args        args
		setupMocks  func()
		wantUrlResp models.Event
		wantErr     bool
	}{

		{
			name: "Failure",
			args: args{
				ctx: context.Background(),
				eventReq: models.Event{
					Name:        "",
					Description: "",
				},
			},
			setupMocks: func() {
				mockRepo.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(models.Event{}, errors.New("Some error")).Times(1)

			},
			wantUrlResp: models.Event{},
			wantErr:     true,
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				eventReq: models.Event{
					Name:        "google",
					Description: "web",
				},
			},
			setupMocks: func() {
				mockRepo.EXPECT().CreateEvents(gomock.Any(), gomock.Any()).Return(models.Event{
					ID:          0,
					Name:        "google",
					Description: "web",
					CreatedAt:   "",
				}, nil).Times(1)

			},
			wantUrlResp: models.Event{
				ID:          0,
				Name:        "google",
				Description: "web",
				CreatedAt:   "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			gotUrlResp, err := eventService.CreateEvent(tt.args.ctx, tt.args.eventReq)
			if (err != nil) != tt.wantErr {
				if err != errors.New("Some error") {
					t.Errorf("CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotUrlResp, tt.wantUrlResp) {
				t.Errorf("CreateEvent() got = %v, want %v", gotUrlResp, tt.wantUrlResp)
			}

		})
	}
}

func TestGetAllEvents(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(mockCtrl)
	eventService := NewEventService(mockRepo)

	type args struct {
		ctx      context.Context
		eventReq models.Event
	}
	tests := []struct {
		name        string
		args        args
		setupMocks  func()
		wantUrlResp models.EventList
		wantErr     bool
	}{

		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				eventReq: models.Event{
					Name:        "google",
					Description: "web",
				},
			},
			setupMocks: func() {
				mockRepo.EXPECT().GetAllEvents(gomock.Any()).Return(models.EventList{
					Events: []models.Event{},
				}, nil).Times(1)

			},
			wantUrlResp: models.EventList{
				Events: []models.Event{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			gotUrlResp, err := eventService.GetAllEvents(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				if err != errors.New("Some error") {
					t.Errorf("GetAllEvents() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotUrlResp, tt.wantUrlResp) {
				t.Errorf("GetAllEvents() got = %v, want %v", gotUrlResp, tt.wantUrlResp)
			}

		})
	}
}
