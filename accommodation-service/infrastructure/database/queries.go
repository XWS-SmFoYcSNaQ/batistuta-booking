package database

import (
	"database/sql"
)

func GetAllAccommodations(DB *sql.DB) (*sql.Stmt, error) {
	return DB.Prepare(`
		SELECT a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, json_agg(DISTINCT r) as ratings 
		FROM Accommodation a
		LEFT JOIN Rating r ON a.id = r.accommodation_id 
		GROUP BY a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price
	`)
}

func GetAllAccommodationsByHostId(DB *sql.DB) (*sql.Stmt, error) {
	return DB.Prepare(`
		SELECT a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, json_agg(DISTINCT r) as ratings 
		FROM Accommodation a
		LEFT JOIN Rating r ON a.id = r.accommodation_id 
		WHERE host_id = $1 
		GROUP BY a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price
	`)
}

func GetAccommodationById(DB *sql.DB) (*sql.Stmt, error) {
	return DB.Prepare(`
		SELECT a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, json_agg(DISTINCT p) as periods, json_agg(DISTINCT d) as discounts
		FROM accommodation a 
		LEFT JOIN period p ON a.id = p.accommodation_id
		LEFT JOIN discount d on a.id = d.accommodation_id
		WHERE a.id = $1 GROUP BY a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price
	`)
}
