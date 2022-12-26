package service

import (
	"context"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
)

type EventCollectorService interface {
	CreateEvent(ctx context.Context, eventReq models.Event) (eventResponse models.Event, err error)
	GetAllEvents(ctx context.Context) (eventResponse models.EventList, err error)
}
