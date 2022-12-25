package db

import (
	"context"
	"database/sql"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/repository"
)

type DatabaseImpl struct {
	databaseProvider *sql.DB
}

// CreateEvents implements repository.DatabaseRepository
func (*DatabaseImpl) CreateEvents(ctx context.Context, req models.Request) (eventResp models.Response, err error) {

	return eventResp, nil
	//panic("unimplemented")
}

// GetAllEvents implements repository.DatabaseRepository
func (*DatabaseImpl) GetAllEvents(ctx context.Context) (eventResp []models.Response, err error) {
	panic("unimplemented")
}

// NewRepository returns instance of Database repository
func NewDatabaseRepository(db *sql.DB) repository.DatabaseRepository {

	return &DatabaseImpl{databaseProvider: db}
}
