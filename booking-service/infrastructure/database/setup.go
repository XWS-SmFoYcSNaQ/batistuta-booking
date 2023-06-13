package database

import (
	"database/sql"
	"log"
)

func SetupDatabase(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DATABASE CONNECTED")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS BookingRequest (
		id uuid NOT NULL PRIMARY KEY,
		accommodation_id uuid NOT NULL,
		start_date TEXT NOT NULL,
		end_date TEXT NOT NULL,
		number_of_guests integer NOT NULL,
    	user_id uuid NOT NULL)
    `)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Reservation (
		id uuid NOT NULL PRIMARY KEY,
		accommodation_id uuid NOT NULL,
		start_date TEXT NOT NULL,
		end_date TEXT NOT NULL,
		number_of_guests integer NOT NULL,
    	user_id uuid NOT NULL)
    `)
	if err != nil {
		log.Fatalln(err)
	}
}
