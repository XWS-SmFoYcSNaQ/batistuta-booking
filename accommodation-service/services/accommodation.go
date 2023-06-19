package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/infrastructure/database"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/utility"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	"github.com/google/uuid"
	"log"
	"strconv"
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
	if a.AutomaticReservation != 0 && a.AutomaticReservation != 1 {
		return uuid.Nil, errors.New("wrong data for automatic reservation (must be 0 or 1)")
	}

	stmt, err := s.DB.Prepare("INSERT INTO Accommodation VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		log.Println(err)
		return uuid.Nil, errors.New("error while creating accommodation")
	}
	defer stmt.Close()
	id := uuid.New()
	_, err = stmt.Exec(id, a.HostId, a.Name, a.Benefits, a.MinGuests, a.MaxGuests, a.BasePrice, a.Location, a.AutomaticReservation)
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
	err = stmt.QueryRow(id).Scan(&a.ID, &a.HostId, &a.Name, &a.Benefits, &a.MinGuests, &a.MaxGuests, &a.BasePrice, &a.Location, &a.AutomaticReservation, &periodsJSON, &discountsJSON)
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

func (s AccommodationService) GetAccommodationSearchResults(a *accommodation.AM_SearchAccommodations_Request) ([]*model.Accommodation, error) {
	query := "SELECT * FROM Accommodation WHERE min_guests <= " + strconv.Itoa(int(a.NumberOfGuests)) +
		" AND max_guests >= " + strconv.Itoa(int(a.NumberOfGuests))
	if a.Location != "" {
		query += " AND location = '" + a.Location + "'"
	}
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, errors.New("Number of guests : " + strconv.Itoa(int(a.NumberOfGuests)))
	}
	defer rows.Close()

	var accommodations []*model.Accommodation
	for rows.Next() {
		var p model.Accommodation
		err := rows.Scan(&p.ID, &p.HostId, &p.Name, &p.Benefits, &p.MinGuests, &p.MaxGuests, &p.BasePrice, &p.Location, &p.AutomaticReservation)
		if err == sql.ErrNoRows {
			return accommodations, nil
		} else {
			if err != nil {
				return nil, errors.New(err.Error())
			}
		}
		accommodations = append(accommodations, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("bad rows")
	}

	return accommodations, nil
}

//private

func parseAccommodations(rows *sql.Rows) ([]*model.Accommodation, error) {
	var accommodations []*model.Accommodation

	for rows.Next() {
		var p model.Accommodation
		var ratingsJSON []byte

		err := rows.Scan(&p.ID, &p.HostId, &p.Name, &p.Benefits, &p.MinGuests, &p.MaxGuests, &p.BasePrice, &p.Location, &p.AutomaticReservation, &ratingsJSON)
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

func (s AccommodationService) GetAutomaticReservationValue(id string) (int32, error) {
	var automaticReservation int32
	stmt, err := s.DB.Prepare("SELECT automatic_reservation FROM Accommodation WHERE id = $1")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&automaticReservation)
	if err != nil {
		return -1, err
	}

	return automaticReservation, nil
}

func (s AccommodationService) GetHostIdByAccommodationId(id string) (string, error) {
	var hostId string
	stmt, err := s.DB.Prepare("SELECT host_id FROM Accommodation WHERE id = $1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&hostId)
	if err != nil {
		return "", err
	}

	return hostId, nil
}
