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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Accommodation (
		id uuid NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		benefits TEXT,
		min_guests integer,
		max_guests integer NOT NULL,
		base_price DOUBLE PRECISION NOT NULL)
    `)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Period (
    	id uuid NOT NULL PRIMARY KEY,
    	p_start TIMESTAMPTZ NOT NULL,
    	p_end TIMESTAMPTZ NOT NULL,
    	accommodation_id uuid NOT NULL,
    	user_id uuid,
    	FOREIGN KEY (accommodation_id) REFERENCES Accommodation (id)
	)`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Price (
    	id uuid NOT NULL PRIMARY KEY,
    	p_start TIMESTAMPTZ NOT NULL,
    	p_end TIMESTAMPTZ NOT NULL,
    	accommodation_id uuid NOT NULL,
    	user_id uuid,
    	price DOUBLE PRECISION NOT NULL,
    	FOREIGN KEY (accommodation_id) REFERENCES Accommodation (id)
	)`)
	if err != nil {
		log.Fatalln(err)
	}
}
