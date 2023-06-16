package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/infrastructure/database"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/utility"
	"github.com/google/uuid"
	"log"
)

type AccommodationService struct {
	DB *sql.DB
}

func (s AccommodationService) GetAll(filters *utility.Filter) ([]*model.Accommodation, error) {
	stmt, err := database.GetAllAccommodations(s.DB, filters)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	accommodations, err := parseAccommodations(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return accommodations, nil
}

func (s AccommodationService) GetAllByHostId(hostId *uuid.UUID) ([]*model.Accommodation, error) {
	stmt, err := database.GetAllAccommodationsByHostId(s.DB)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := stmt.Query(hostId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	accommodations, err := parseAccommodations(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return accommodations, nil
}

func (s AccommodationService) Create(a *model.Accommodation) (uuid.UUID, error) {
	if a.MinGuests < 0 || a.MaxGuests < 0 {
		return uuid.Nil, errors.New("minimum guests and maximum guests can't be negative")
	}
	if a.MinGuests > a.MaxGuests {
		return uuid.Nil, errors.New("maximum guests value has to be larger than minimum one")
	}
	if a.BasePrice < 0 {
		return uuid.Nil, errors.New("price can't be negative")
	}

	stmt, err := s.DB.Prepare("INSERT INTO Accommodation VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Println(err)
		return uuid.Nil, errors.New("error while creating accommodation")
	}
	defer stmt.Close()
	id := uuid.New()
	_, err = stmt.Exec(id, a.HostId, a.Name, a.Benefits, a.MinGuests, a.MaxGuests, a.BasePrice)
	if err != nil {
		log.Println(err)
		return uuid.Nil, errors.New("error while creating accommodation")
	}

	return id, nil
}

func (s AccommodationService) GetById(id uuid.UUID) (*model.Accommodation, error) {
	stmt, err := database.GetAccommodationById(s.DB)
	if err != nil {
		return nil, errors.New("accommodation not found")
	}
	defer stmt.Close()
	var a model.Accommodation
	var periodsJSON []byte
	var discountsJSON []byte
	err = stmt.QueryRow(id).Scan(&a.ID, &a.HostId, &a.Name, &a.Benefits, &a.MinGuests, &a.MaxGuests, &a.BasePrice, &periodsJSON, &discountsJSON)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error while fetching accommodation")
	}

	err = json.Unmarshal(periodsJSON, &a.Periods)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error while fetching accommodation's periods")
	}
	err = json.Unmarshal(discountsJSON, &a.Discounts)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error while fetching accommodation's discounts")
	}
	if len(a.Discounts) > 0 && a.Discounts[0] == nil {
		a.Discounts = append(a.Discounts[:0], a.Discounts[1:]...)
	}
	if len(a.Periods) > 0 && a.Periods[0] == nil {
		a.Periods = append(a.Periods[:0], a.Periods[1:]...)
	}
	return &a, nil
}

//private

func parseAccommodations(rows *sql.Rows) ([]*model.Accommodation, error) {
	var accommodations []*model.Accommodation

	for rows.Next() {
		var p model.Accommodation
		var ratingsJSON []byte

		err := rows.Scan(&p.ID, &p.HostId, &p.Name, &p.Benefits, &p.MinGuests, &p.MaxGuests, &p.BasePrice, &ratingsJSON)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(ratingsJSON, &p.Ratings)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if len(p.Ratings) > 0 && p.Ratings[0] == nil {
			p.Ratings = p.Ratings[1:]
		}

		accommodations = append(accommodations, &p)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return accommodations, nil
}
