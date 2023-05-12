package database

import (
	"accommodation_service/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

func Connect(cfg config.Config) *sql.DB {
	split := strings.Split(cfg.DBAddress, ":")
	host := "localhost"
	port := split[1]
	user := cfg.DBUsername
	password := cfg.DBPassword
	if split[0] != "" {
		host = split[0]
	}
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", user, password, host, port)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE DATABASE accommodation")
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DATABASE WORKS!")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Accommodation (" +
		"id uuid NOT NULL PRIMARY KEY, " +
		"name TEXT NOT NULL," +
		"benefits TEXT," +
		"min_guests integer," +
		"max_guests integer)",
	)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
