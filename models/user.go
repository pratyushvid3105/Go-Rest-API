package models

import (
	"errors"

	"github.com/pratyushvid3105/Go-Rest-API/db"
	"github.com/pratyushvid3105/Go-Rest-API/utils"
)

type User struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error{
	// the password that is stored in the database should not be the plain password the user entered. Instead, it should be hashed, so converted to a different value in a way that can't be reversed, so that we can't get the original password from the hashed value
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	hashedPassword, err := utils.HashPassword(u.Password) 
	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, u.Email, hashedPassword)

	if err != nil {
		return err
	}

	var userId int64
	userId, err = result.LastInsertId()
	u.ID = userId
	return err
}

func GetAllUsers() ([]User, error) {
	// Now I'm using Query instead of Exec here, which we could use as well, because Query is typically used if you have a query where you wanna get back a bunch of rows, which you then wanna use, which is exactly what's the case here. Whereas Exec is used whenever you have a query that changes data in the database inserts data, updates data and so on. 
	// If you have a query that changes stuff, it's Exec. If you have a query that fetches data, it's Query.
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	// Next returns a boolean, which is true as long as there are rows left and false thereafter. And therefore this will keep the loop running for as long as there are rows to read and it will proceed through all those rows step by step with every iteration
	for rows.Next() {
		var user User
		// the Scan method, which now reads the content of the row we're currently processing and Scan then works a little bit like the FMT Package Scan method where you pass a pointer to Scan so that it's populated with the data from the row, though you don't just pass one pointer, but a bunch of pointers, one for every column that can be found in the row, in the order in which the columns were defined.
		// So here for this event, we wanna start by populating the ID. So, we pass that ID to Scan, however not the raw value event.ID, but instead a pointer to that. So a pointer to that ID field in that event Struct (event.ID). And we then do this for all those fields.
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (u *User) ValidateCredentials() error{
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}