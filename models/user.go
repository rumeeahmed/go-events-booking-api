package models

import (
	"database/sql"
	"errors"
	"go-events-booking-api/db"
	"go-events-booking-api/utils"
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

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	id, err := result.LastInsertId() // gets the last inserted id of the item
	if err != nil {
		return err
	}

	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	database := db.GetDb()
	sqlStmt := `SELECT password FROM users WHERE email = ?`

	result := database.QueryRow(sqlStmt, u.Email)

	var retrievedPassword string
	err := result.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	isValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isValid {
		return errors.New("invalid credentials")
	}

	return nil
}
