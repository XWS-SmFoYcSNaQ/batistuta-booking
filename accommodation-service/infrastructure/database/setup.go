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
		host_id uuid NOT NULL,
		name TEXT NOT NULL,
		benefits TEXT,
		min_guests integer,
		max_guests integer NOT NULL,
		base_price DOUBLE PRECISION NOT NULL,
		location TEXT NOT NULL,
		automatic_reservation INT)
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
    	guests integer NOT NULL,
    	FOREIGN KEY (accommodation_id) REFERENCES Accommodation (id)
	)`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Discount (
    	id uuid NOT NULL PRIMARY KEY,
    	d_start TIMESTAMPTZ NOT NULL,
    	d_end TIMESTAMPTZ NOT NULL,
    	accommodation_id uuid NOT NULL,
    	user_id uuid,
    	discount DOUBLE PRECISION NOT NULL,
    	FOREIGN KEY (accommodation_id) REFERENCES Accommodation (id)
	)`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`DROP TABLE IF EXISTS Rating;`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Rating (
    	id uuid NOT NULL PRIMARY KEY,
    	accommodation_id uuid NOT NULL,
    	user_id uuid NOT NULL,
    	value_ integer NOT NULL,
    	FOREIGN KEY (accommodation_id) REFERENCES Accommodation (id)
	)`)
	if err != nil {
		log.Fatalln(err)
	}
}
