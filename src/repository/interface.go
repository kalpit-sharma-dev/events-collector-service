package repository

import (
	"context"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
)

type DatabaseRepository interface {
	CreateEvents(ctx context.Context, req models.Event) (eventResp models.Event, err error)
	GetAllEvents(ctx context.Context) (eventList models.EventList, err error)
}
