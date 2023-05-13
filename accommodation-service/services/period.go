package services

import (
	"accommodation_service/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
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
	overlap, err := s.checkOverlap(p.Start, p.End)
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}
	if overlap {
		return uuid.Nil, errors.New("timestamps are overlapping")
	}

	stmt, err := s.DB.Prepare(`INSERT INTO Period VALUES ($1, $2, $4, $5, $6)`)
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}
	id := uuid.New()
	_, err = stmt.Exec(id, p.Start, p.End, p.AccommodationId, p.UserId)
	if err != nil {
		return uuid.Nil, errors.New(errorMessage)
	}

	return id, nil
}

func (s PeriodService) checkOverlap(start, end time.Time) (bool, error) {
	stmt, err := s.DB.Prepare(`
		SELECT COUNT(*) as c FROM Period p INNER JOIN Accommodation a ON a.id = p.accommodation_id
		WHERE (p_start, p_end) OVERLAPS ($1, $2) AND (a.max_guests >= c OR p.user_id = NULL)
	`)
	if err != nil {
		return false, err
	}
	var count int
	err = stmt.QueryRow(start, end).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
