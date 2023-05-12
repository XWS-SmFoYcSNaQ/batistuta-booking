package database

import (
	"database/sql"
	"log"
)

func SetupDatabase(db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS accommodation")
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DATABASE WORKS!")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Accommodation (
		id uuid NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		benefits TEXT,
		min_guests integer,
		max_guests integer)
    `)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Period (
    	id uuid NOT NULL PRIMARY KEY,
    	p_start TIMESTAMPTZ NOT NULL,
    	p_end TIMESTAMPTZ NOT NULL,
    	period_type integer NOT NULL,
    	accommodation_id uuid,
    	user_id uuid,
    	FOREIGN KEY (accommodation_id) REFERENCES Accommodation (id)
	)`)
	if err != nil {
		log.Fatalln(err)
	}
}
