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
	DateTime time.Time
	UserID int `binding:"required"`
}

var events = []Event{}

func (e Event) Save() error{
	// we want to use the actual data that we're getting from the request instead of hard coding something here and to inject that received data in a safe way into this query which is not vulnerable to SQL injection attacks, we should add a couple of question marks here. One for every column into which a value should be inserted. So five question marks in total here because that's a special syntax that is supported by these SQL Packages that gives us a SQL injection safe way of inserting values into this query
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)
	`

	// To prepare we parse our query and as a result we'll get the prepared SQL statement or an error if something goes wrong.
	statement, err1 := db.DB.Prepare(query)
	if err1 != nil {
		return err1
	}
	// One last thing we also should do with that statement is close it after we executed it and a good way of doing that, is with help of the defer keyword, which allows us to call close here without executing it here. Instead, it will be executed by go now whenever this safe function is done, no matter if it's done because we had an error anywhere or because it completed successfully, that's how we should close this statement.
	defer statement.Close()
	// We can use this prepared statement to execute it with this exec method which exists on this statement value. And to exec we now can parse as many arguments as we need, one for every placeholder we have here. So one for every question mark in the order in which those values should be assigned to those columns.
	result, err2 := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err2 != nil {
		return err2
	}
	// We can use this result to call LastInsertId to get the id of the event that was inserted because remember that we actually configured events table such that the ID is set automatically and we can get this automatically generated ID with help of this LastInsertId function here. So as a result we get back the id or an error if this somehow fails or if no id was found and I want to use that id to set it on my event. So I'll set the event ID to id. 
	id, err3 := result.LastInsertId()
	e.ID = id
	return err3
}

func GetAllEvents() []Event {
	return events
}