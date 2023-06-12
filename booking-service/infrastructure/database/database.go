package database

import (
	"booking_service/config"
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

	SetupDatabase(db)
	return db
}
