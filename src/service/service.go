package service

import (
	"context"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/repository"
)

type EventCollectorServiceImpl struct {
	repository repository.DatabaseRepository
}

// CreateEvent implements EventService
func (s *EventCollectorServiceImpl) CreateEvent(ctx context.Context, eventReq models.Event) (eventResponse models.Event, err error) {
	eventResponse, err = s.repository.CreateEvents(ctx, eventReq)
	return
}

// GetEvent implements EventService
func (s *EventCollectorServiceImpl) GetAllEvents(ctx context.Context) (eventResponse models.EventList, err error) {
	eventResponse, err = s.repository.GetAllEvents(ctx)
	return
}

func NewEventService(repository repository.DatabaseRepository) EventCollectorService {
	return &EventCollectorServiceImpl{repository: repository}
}
