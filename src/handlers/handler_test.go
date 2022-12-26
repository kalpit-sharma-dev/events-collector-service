package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/mocks"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
)

func TestHandleCreateUrl(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEventCollectorService(ctrl)
	urlHandler := NewHandler(mockRepo)

	var successReq = models.Event{
		ID:          0,
		Name:        "",
		Description: "",
		CreatedAt:   "",
	}

	tests := []struct {
		name           string
		req            models.Event
		setupMocks     func()
		expectedStatus int
	}{
		{
			name: "Success",
			req:  successReq,
			setupMocks: func() {
				mockRepo.EXPECT().CreateEvent(gomock.Any(), gomock.Any()).Return(models.Event{}, nil).AnyTimes()
			},
			expectedStatus: 201,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			b, _ := json.Marshal(tt.req)
			rr := httptest.NewRecorder()
			httpReq, _ := http.NewRequest("POST", "/url", bytes.NewBuffer(b))
			urlHandler.HandleCreateEvents(rr, httpReq)
			if tt.expectedStatus != rr.Code {
				t.Errorf("Error got = %v and want = %v", rr.Code, tt.expectedStatus)
			}
		})
	}
}
