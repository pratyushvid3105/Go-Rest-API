package models

import (
	"github.com/pratyushvid3105/Go-Rest-API/db"
)

type Registration struct {
	ID int64
	EventID int64 `binding:"required"`
	UserID int64 `binding:"required"`
}

func GetAllRegistrations() ([]Registration, error) {
	query := "SELECT * FROM registrations"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registrations []Registration

	for rows.Next() {
		var registration Registration
		err := rows.Scan(&registration.ID, &registration.EventID, &registration.UserID)

		if err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}
	return registrations, nil
}