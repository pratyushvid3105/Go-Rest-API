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
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
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