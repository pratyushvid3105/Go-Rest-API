package models

import "time"

type Event struct {
	ID int `binding:"required"`
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time
	UserID int `binding:"required"`
}

var events = []Event{}

func (e Event) Save(){
	// later: add it to a DB
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}