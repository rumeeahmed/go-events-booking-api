package models

import (
	"database/sql"
	"go-events-booking-api/db"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	sqlStmt := `INSERT INTO users (email, password) VALUES (? , ?)`
	database := db.GetDb()
	stmt, err := database.Prepare(sqlStmt)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	result, err := stmt.Exec(u.Email, u.Password)

	id, err := result.LastInsertId() // gets the last inserted id of the item
	if err != nil {
		return err
	}

	u.ID = id
	return err
}
