package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(){
	var err error
	// api.db file will be created if not present. Technically when calling Open, we're not really creating a connection. Instead, we're just opening the database, and Go is then able to open connections thereafter.
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to DB")
	}

	// With this setMaxOpenCons method, we then control how many connections can be open simultaneously at most, let's say 10, and that will make sure that later when the application runs, we don't keep on opening new connections all the time, but instead we have a pool of ongoing connections that can be used whenever they're needed by different parts of the application. And at the same time, we make sure that this pool isn't too big. 
	DB.SetMaxOpenConns(10)

	// In addition to setting the max number, we can also set the max idle connection number, which simply means how many connections we want to keep open if no one's using these connections at the moment. And with that, I'm making sure that we have at least five connections open at all time, so that if we have an incoming request, it can immediately be handled, but more connections will be opened all the way up to a number of 10 connections if we have more requests coming in and if we need more connections. If we need more than 10, those other requests will have to wait until a connection is available again.
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables(){
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}