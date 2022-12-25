package service

import (
	"context"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
)

type EventCollectorService interface {
	CreateEvent(ctx context.Context, eventReq models.Request) (eventResponse models.Response, err error)
	GetAllEvents(ctx context.Context) (eventResponse models.Response, err error)
}
