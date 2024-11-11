package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // Global sqlite db to be used outside the package.

func InitDb() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	if DB == nil {
		panic("DB is nil after opening connection.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    email TEXT NOT NULL UNIQUE, 
	    password TEXT NOT NULL
  	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL, 
	    description TEXT NOT NULL, 
	    location TEXT NOT NULL, 
	    date_time DATETIME NOT NULL,
	    user_id INTEGER,
	    FOREIGN KEY (user_id) REFERENCES users(id)
  	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table: " + err.Error())
	}
}

func GetDb() *sql.DB {
	return DB
}
