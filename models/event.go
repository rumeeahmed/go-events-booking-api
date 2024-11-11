package models

import (
	"database/sql"
	"go-events-booking-api/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	sqlStmt := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (? , ?, ?, ?, ?)`
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

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	id, err := result.LastInsertId() // gets the last inserted id of the item
	if err != nil {
		return err
	}

	e.ID = id
	return err
}

func (e *Event) Update() (*Event, error) {
	database := db.GetDb()
	sqlStmt := `UPDATE events SET name=?, description=?, location=?, date_time=? WHERE id =? RETURNING *`
	stmt, err := database.Prepare(sqlStmt)

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	result := database.QueryRow(sqlStmt, e.Name, e.Description, e.Location, e.DateTime, e.ID)

	var event Event
	err = result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, err
}

func (e *Event) Delete() error {
	database := db.GetDb()
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := database.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents() ([]Event, error) {
	database := db.GetDb()
	sqlStmt := `SELECT * FROM events`
	rows, err := database.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	var events []Event

	// Iterate through all the rows, until there are no more and create an event object and append to list for all rows.
	for rows.Next() {
		var event Event
		// Pass a pointer per the required fields to the scan method which will populate the event object with values.
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	database := db.GetDb()
	sqlStmt := `SELECT * FROM events WHERE id = ?`

	result := database.QueryRow(sqlStmt, id)

	var event Event
	err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, err
}
