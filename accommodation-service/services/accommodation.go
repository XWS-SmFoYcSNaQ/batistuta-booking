package services

import (
	"accommodation_service/model"
	"accommodation_service/proto/accommodation"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"strconv"
)

type AccommodationService struct {
	DB *sql.DB
	PeriodService
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

func (s AccommodationService) GetSearchedAccommodations(a *accommodation.AM_SearchAccommodations_Request) ([]*model.Accommodation, error) {
	//errorMessage := "error while fetching accommodations"
	rows, err := s.DB.Query("SELECT * FROM Accommodation WHERE min_guests <= " + strconv.Itoa(int(a.NumberOfGuests)) +
		" AND max_guests >= " + strconv.Itoa(int(a.NumberOfGuests)))
	if err != nil {
		return nil, errors.New("Number of guests : " + strconv.Itoa(int(a.NumberOfGuests)))
	}
	defer rows.Close()

	var accommodations []*model.Accommodation
	for rows.Next() {
		var p model.Accommodation
		err := rows.Scan(&p.ID, &p.Name, &p.Benefits, &p.MinGuests, &p.MaxGuests, &p.BasePrice)
		if err != nil {
			return nil, errors.New("lose appendovao")
		}
		accommodations = append(accommodations, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("losi redovi")
	}

	return accommodations, nil
}

func (s AccommodationService) GetById(id uuid.UUID) (*model.Accommodation, error) {
	stmt, err := s.DB.Prepare(`
		SELECT * FROM Accommodation WHERE id = $1
	`)
	if err != nil {
		return nil, errors.New("accommodation not found")
	}
	defer stmt.Close()
	var a model.Accommodation
	err = stmt.QueryRow(id).Scan(&a.ID, &a.HostId, &a.Name, &a.Benefits, &a.MinGuests, &a.MaxGuests, &a.BasePrice)
	if err != nil {
		return nil, errors.New("error while fetching accommodation")
	}
	return &a, nil
}
