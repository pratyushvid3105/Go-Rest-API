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

func (e Event) Save() error{
	// we want to use the actual data that we're getting from the request instead of hard coding something here and to inject that received data in a safe way into this query which is not vulnerable to SQL injection attacks, we should add a couple of question marks here. One for every column into which a value should be inserted. So five question marks in total here because that's a special syntax that is supported by these SQL Packages that gives us a SQL injection safe way of inserting values into this query
	query := `
	INSERT INTO events(name, description, location, dateTime, userId)
	VALUES(?, ?, ?, ?, ?)
	`

	// Using Prepare() is 100% optional! We could send all your commands directly via Exec() or Query().
	result, err1 := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err1 != nil {
		return err1
	}
	// We can use this result to call LastInsertId to get the id of the event that was inserted because remember that we actually configured events table such that the ID is set automatically and we can get this automatically generated ID with help of this LastInsertId function here. So as a result we get back the id or an error if this somehow fails or if no id was found and I want to use that id to set it on my event. So I'll set the event ID to id. 
	id, err2 := result.LastInsertId()
	e.ID = id
	return err2
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	// Use QueryRow method cause we know result will consist of only 1 row
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) 
	if err != nil {
		// Now we actually can't use nil as a zero value for event because the zero value for event is essentially an empty Struct (Event{}). But here, in order to be able to use nil, we will instead simply return a pointer to event because the null value for a pointer is nil. So if there is no address available because we have no event in this case, and therefore we can then return nil and the error.
		return nil, err
	}
	// here we will return a pointer to the created event and nil as a value for the error.
	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	// Now I'm using Query instead of Exec here, which we could use as well, because Query is typically used if you have a query where you wanna get back a bunch of rows, which you then wanna use, which is exactly what's the case here. Whereas Exec is used whenever you have a query that changes data in the database inserts data, updates data and so on. 
	// If you have a query that changes stuff, it's Exec. If you have a query that fetches data, it's Query.
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	// Next returns a boolean, which is true as long as there are rows left and false thereafter. And therefore this will keep the loop running for as long as there are rows to read and it will proceed through all those rows step by step with every iteration
	for rows.Next() {
		var event Event
		// the Scan method, which now reads the content of the row we're currently processing and Scan then works a little bit like the FMT Package Scan method where you pass a pointer to Scan so that it's populated with the data from the row, though you don't just pass one pointer, but a bunch of pointers, one for every column that can be found in the row, in the order in which the columns were defined.
		// So here for this event, we wanna start by populating the ID. So, we pass that ID to Scan, however not the raw value event.ID, but instead a pointer to that. So a pointer to that ID field in that event Struct (event.ID). And we then do this for all those fields.
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
/* 
Preparing Statements vs Directly Executing Queries (Prepare() vs Exec()/Query())
We started sending SQL commands to the SQLite database.

And we did this by following different approaches:

DB.Exec() (when we created the tables)

Prepare() + stmt.Exec() (when we inserted data into the database)

DB.Query() (when we fetched data from the database)

Using Prepare() is 100% optional! You could send all your commands directly via Exec() or Query().

The difference between those two methods then just is whether you're fetching data from the database (=> use Query()) or your manipulating the database / data in the database (=> use Exec()).

But what's the advantage of using Prepare()?

Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).

This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.

And, indeed, in this application, we are calling stmt.Close() directly after calling stmt.Exec(). So here, it really wouldn't matter which approach you're using.

But in order to show you the different ways of using the sql package, I decided to also include this preparation approach in this course.
*/