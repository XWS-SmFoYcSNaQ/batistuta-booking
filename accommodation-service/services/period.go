package services

import (
	"accommodation_service/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"
)

type PeriodService struct {
	DB *sql.DB
}

func (s PeriodService) GetAllByAccommodation(id uuid.UUID) ([]*model.Period, error) {
	errorMessage := "error while fetching periods"
	stmt, err := s.DB.Prepare("SELECT * FROM Period WHERE accommodation_id = $1")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	defer rows.Close()

	var data []*model.Period
	for rows.Next() {
		var d model.Period
		err := rows.Scan(&d.ID, &d.Start, &d.End, &d.AccommodationId, &d.UserId)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
		data = append(data, &d)
	}
	return data, nil
}

func (s PeriodService) Create(p *model.Period) (uuid.UUID, error) {
	errorMessage := "error while creating period"

	free, err := s.isAvailableForGivenInterval(p.Start, p.End)
	if err != nil {
		return uuid.Nil, errors.New("please choose valid start and end dates")
	}
	if !free {
		return uuid.Nil, errors.New("accommodation is not available for the given time interval")
	}
	hasSpace, err := s.hasEnoughSpaceForGivenInterval(p.Start, p.End)
	if err != nil {
		return uuid.Nil, errors.New("please choose valid start and end dates")
	}
	if !hasSpace {
		return uuid.Nil, errors.New("there is not enough space for the given time interval")
	}

	stmt, err := s.DB.Prepare(`INSERT INTO Period VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}
	defer stmt.Close()
	id := uuid.New()
	if p.UserId == uuid.Nil {
		_, err = stmt.Exec(id, p.Start, p.End, p.AccommodationId, nil)
	} else {
		_, err = stmt.Exec(id, p.Start, p.End, p.AccommodationId, p.UserId)
	}
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}

	return id, nil
}

func (s PeriodService) isAvailableForGivenInterval(start, end time.Time) (bool, error) {
	stmt, err := s.DB.Prepare(`
		SELECT COUNT(*) AS accommodation_count FROM Period p
		WHERE (p.p_start, p.p_end) OVERLAPS ($1, $2) AND user_id IS NULL
	`)
	defer stmt.Close()
	if err != nil {
		return false, err
	}
	var count int
	err = stmt.QueryRow(start, end).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (s PeriodService) hasEnoughSpaceForGivenInterval(start, end time.Time) (bool, error) {
	stmt, err := s.DB.Prepare(`
		SELECT p.accommodation_id, COUNT(*) AS accommodation_count, a.max_guests
		FROM Period p
		INNER JOIN Accommodation a ON a.id = p.accommodation_id
		WHERE (p.p_start, p.p_end) OVERLAPS ($1, $2) AND p.user_id IS NOT NULL
		GROUP BY p.accommodation_id, a.max_guests
	`)
	defer stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	var count int
	var id string
	var maxGuests int
	err = stmt.QueryRow(start, end).Scan(&id, &count, &maxGuests)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		log.Println(err.Error())
		return false, err
	}
	return maxGuests > count, nil
}
