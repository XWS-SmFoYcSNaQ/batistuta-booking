package services

import (
	"accommodation_service/model"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type AccommodationService struct {
	DB *sql.DB
}

func (s AccommodationService) GetAll(hostId uuid.UUID) ([]*model.Accommodation, error) {
	var rows *sql.Rows
	if hostId != uuid.Nil {
		stmt, err := s.DB.Prepare("SELECT * FROM Accommodation WHERE host_id = $1")
		if err != nil {
			return nil, err
		}
		rows, err = stmt.Query(hostId)
		if err != nil {
			return nil, err
		}
	} else {
		stmt, err := s.DB.Prepare("SELECT * FROM Accommodation")
		if err != nil {
			return nil, err
		}
		rows, err = stmt.Query()
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	var accommodations []*model.Accommodation
	for rows.Next() {
		var p model.Accommodation
		err := rows.Scan(&p.ID, &p.HostId, &p.Name, &p.Benefits, &p.MinGuests, &p.MaxGuests, &p.BasePrice)
		if err != nil {
			return nil, err
		}
		accommodations = append(accommodations, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accommodations, nil
}

func (s AccommodationService) Create(a *model.Accommodation) (uuid.UUID, error) {
	stmt, err := s.DB.Prepare("INSERT INTO Accommodation VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return uuid.Nil, err
	}
	defer stmt.Close()
	id := uuid.New()
	_, err = stmt.Exec(id, a.HostId, a.Name, a.Benefits, a.MinGuests, a.MaxGuests, a.BasePrice)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s AccommodationService) GetById(id uuid.UUID) (*model.Accommodation, error) {
	stmt, err := s.DB.Prepare(`
		SELECT a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price, array_agg(p) as periods, array_agg(d) as discounts
		FROM accommodation a 
		JOIN period p ON a.id = p.accommodation_id
		JOIN discount d on a.id = d.accommodation_id
		WHERE a.id = $1 GROUP BY a.id, a.host_id, a.name, a.benefits, a.min_guests, a.max_guests, a.base_price
	`)
	if err != nil {
		return nil, errors.New("accommodation not found")
	}
	defer stmt.Close()
	var a model.Accommodation
	var periodsJSON []byte
	var discountsJSON []byte
	err = stmt.QueryRow(id).Scan(&a.ID, &a.HostId, &a.Name, &a.Benefits, &a.MinGuests, &a.MaxGuests, &a.BasePrice, &periodsJSON, &discountsJSON)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(periodsJSON, &a.Periods)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(discountsJSON, &a.Discounts)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
