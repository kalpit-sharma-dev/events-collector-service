package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/repository"
)

type DatabaseImpl struct {
	Conn *sql.DB
}

// CreateEvents implements repository.DatabaseRepository
func (db *DatabaseImpl) CreateEvents(ctx context.Context, req models.Event) (eventResp models.Event, err error) {
	var id int
	var createdAt string
	query := `INSERT INTO events (name, description) VALUES ($1, $2) RETURNING id, created_at`
	fmt.Println("$$$$$$$$$$$$QUERY$$$$$$$$$$$$$$$$$", query)
	err = db.Conn.QueryRow(query, req.Name, req.Description).Scan(&id, &createdAt)
	if err != nil {
		return
	}
	eventResp.ID = id
	eventResp.CreatedAt = createdAt
	return
}

// GetAllEvents implements repository.DatabaseRepository
func (db *DatabaseImpl) GetAllEvents(ctx context.Context) (eventList models.EventList, err error) {
	rows, err := db.Conn.Query("SELECT * FROM events ORDER BY ID DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		var event models.Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.CreatedAt)
		if err != nil {
			return
		}
		eventList.Items = append(eventList.Items, event)
	}
	return eventList, nil
}

// NewRepository returns instance of Database repository
func NewDatabaseRepository(db *sql.DB) repository.DatabaseRepository {

	return &DatabaseImpl{Conn: db}
}
