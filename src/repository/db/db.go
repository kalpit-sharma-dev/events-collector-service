package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/repository"
)

type DatabaseImpl struct {
	Conn *sql.DB
}

const (
	getAllEventsQuery = `SELECT * FROM events ORDER BY ID DESC`
	createEventsQuery = `INSERT INTO events (name, description) VALUES ($1, $2) RETURNING id, description, name, created_at`
)

// CreateEvents implements repository.DatabaseRepository
func (db *DatabaseImpl) CreateEvents(ctx context.Context, req models.Event) (eventResp models.Event, err error) {
	var id int
	var createdAt string
	var description string
	var name string

	err = db.Conn.QueryRow(createEventsQuery, req.Name, req.Description).Scan(&id, &description, &name, &createdAt)
	if err != nil {
		log.Println(err)
		return
	}

	eventResp.ID = id
	eventResp.CreatedAt = createdAt
	eventResp.Name = name
	eventResp.Description = description
	return
}

// GetAllEvents implements repository.DatabaseRepository
func (db *DatabaseImpl) GetAllEvents(ctx context.Context) (eventList models.EventList, err error) {
	rows, err := db.Conn.Query(getAllEventsQuery)
	if err != nil {
		return
	}
	for rows.Next() {
		var event models.Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.CreatedAt)
		if err != nil {
			return
		}
		eventList.Events = append(eventList.Events, event)
	}
	return eventList, nil
}

// NewRepository returns instance of Database repository
func NewDatabaseRepository(db *sql.DB) repository.DatabaseRepository {

	return &DatabaseImpl{Conn: db}
}
