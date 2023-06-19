package database

import (
	"database/sql"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/utility"
	"strconv"
	"strings"
)

func GetAllAccommodations(DB *sql.DB, filter *utility.Filter) (*sql.Stmt, error) {
	query := `SELECT a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, a.location, a.automatic_reservation, json_agg(DISTINCT r) as ratings `
	query += ` FROM Accommodation a`
	query += ` LEFT JOIN Rating r ON a.id = r.accommodation_id`
	query += FormFilterWhereClause(filter)
	query += ` GROUP BY a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, a.location, a.automatic_reservation`
	return DB.Prepare(query)
}

func GetAllAccommodationsByHostId(DB *sql.DB) (*sql.Stmt, error) {
	return DB.Prepare(`
		SELECT a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, a.location, a.automatic_reservation, json_agg(DISTINCT r) as ratings 
		FROM Accommodation a
		LEFT JOIN Rating r ON a.id = r.accommodation_id 
		WHERE host_id = $1 
		GROUP BY a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, a.location, a.automatic_reservation
	`)
}

func GetAccommodationById(DB *sql.DB) (*sql.Stmt, error) {
	return DB.Prepare(`
		SELECT a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, a.location, a.automatic_reservation, json_agg(DISTINCT p) as periods, json_agg(DISTINCT d) as discounts
		FROM accommodation a 
		LEFT JOIN period p ON a.id = p.accommodation_id
		LEFT JOIN discount d on a.id = d.accommodation_id
		WHERE a.id = $1 GROUP BY a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, a.location, a.automatic_reservation
	`)
}

func FormFilterWhereClause(filter *utility.Filter) string {
	var strs []string
	if filter.Range != nil {
		r := *filter.Range
		if r.Min != -1 {
			strs = append(strs, "base_price >= "+strconv.FormatFloat(r.Min, 'g', 5, 64))
		}
		if r.Max != -1 {
			strs = append(strs, "base_price <= "+strconv.FormatFloat(r.Max, 'g', 5, 64))
		}
	}

	for _, f := range filter.Benefits {
		strs = append(strs, "benefits LIKE '%"+f+"%'")
	}

	if len(strs) > 0 {
		return " WHERE " + strings.Join(strs, " AND ")
	}
	return ""
}
