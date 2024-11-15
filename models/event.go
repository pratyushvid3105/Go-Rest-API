package models

import (
	"time"

	"github.com/pratyushvid3105/Go-Rest-API/db"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int64
}

func (e *Event) Save() error{
	query := `
	INSERT INTO events(name, description, location, dateTime, userId)
	VALUES(?, ?, ?, ?, ?)
	`

	result, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error{
	query := `
	DELETE FROM events WHERE id = ?`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error{
	query := "INSERT INTO registrations(eventId, userId) VALUES(?, ?)"

	_, err := db.DB.Exec(query, e.ID, userId)
	if err != nil {
		return err
	}
	return err
}

func (e Event) CancelRegistration(userId int64) error{
	query := "DELETE FROM registrations WHERE eventId = ? AND userId = ?"

	_, err := db.DB.Exec(query, e.ID, userId)
	if err != nil {
		return err
	}
	return err
}