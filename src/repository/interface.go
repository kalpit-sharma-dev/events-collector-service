package repository

import (
	"context"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
)

type DatabaseRepository interface {
	CreateEvents(ctx context.Context, req models.Request) (eventResp models.Response, err error)
	GetAllEvents(ctx context.Context) (eventResp []models.Response, err error)
}
