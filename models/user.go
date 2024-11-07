package models

import "github.com/pratyushvid3105/Go-Rest-API/db"

type User struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error{
	// the password that is stored in the database should not be the plain password the user entered. Instead, it should be hashed, so converted to a different value in a way that can't be reversed, so that we can't get the original password from the hashed value
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	result, err := db.DB.Exec(query, u.Email, u.Password)

	if err != nil {
		return err
	}

	var userId int64
	userId, err = result.LastInsertId()
	u.ID = userId
	return err
}