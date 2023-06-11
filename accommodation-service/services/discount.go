package services

import (
	"database/sql"
	"errors"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/model"
	"github.com/google/uuid"
	"log"
	"time"
)

type DiscountService struct {
	DB *sql.DB
}

func (s DiscountService) GetAllByAccommodation(id uuid.UUID) ([]*model.Discount, error) {
	errorMessage := "error while fetching prices"
	stmt, err := s.DB.Prepare(`
		SELECT * FROM Discount
		WHERE accommodation_id = $1 AND user_id IS NULL
	`)
	if err != nil {
		log.Println(err)
		return nil, errors.New(errorMessage)
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New(errorMessage)
	}
	defer rows.Close()

	var data []*model.Discount
	for rows.Next() {
		var d model.Discount
		err := rows.Scan(&d.ID, &d.Start, &d.End, &d.AccommodationId, &d.UserId, &d.Discount)
		if err != nil {
			log.Println(err)
			return nil, errors.New(errorMessage)
		}
		data = append(data, &d)
	}
	return data, nil
}

func (s DiscountService) GetAllByAccommodationAndInterval(id uuid.UUID, start, end time.Time, userId uuid.UUID) ([]*model.Discount, error) {
	errorMessage := "error while fetching prices"
	stmt, err := s.DB.Prepare(`
		SELECT * FROM Discount 
		WHERE accommodation_id = $1 AND (d_start, d_end) OVERLAPS ($2, $3) AND user_id IS NULL
	`)
	if userId != uuid.Nil {
		stmt, err = s.DB.Prepare(`
			SELECT * FROM Discount 
			WHERE accommodation_id = $1 AND (d_start, d_end) OVERLAPS ($2, $3) AND user_id = $4
		`)
	}
	if err != nil {
		log.Println(err)
		return nil, errors.New(errorMessage)
	}
	rows, err := stmt.Query(id, start, end, userId)
	if err != nil {
		log.Println(err)
		return nil, errors.New(errorMessage)
	}
	defer rows.Close()

	var data []*model.Discount
	for rows.Next() {
		var d model.Discount
		err := rows.Scan(&d.ID, &d.Start, &d.End, &d.AccommodationId, &d.UserId, &d.Discount)
		if err != nil {
			log.Println(err)
			return nil, errors.New(errorMessage)
		}
		data = append(data, &d)
	}
	return data, nil
}

func (s DiscountService) Create(d *model.Discount) (uuid.UUID, error) {
	errorMessage := "error while creating discount"
	if d.Discount <= 0 || d.Discount > 100 {
		return uuid.Nil, errors.New("discount can't be zero, negative or higher than 100")
	}
	stmt, err := s.DB.Prepare(`INSERT INTO Discount VALUES ($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		log.Println(err)
		return uuid.Nil, errors.New(errorMessage)
	}
	defer stmt.Close()
	id := uuid.New()
	if d.UserId == uuid.Nil {
		_, err = stmt.Exec(id, d.Start, d.End, d.AccommodationId, nil, d.Discount)
	} else {
		_, err = stmt.Exec(id, d.Start, d.End, d.AccommodationId, d.UserId, d.Discount)
	}
	if err != nil {
		log.Println(err)
		return uuid.Nil, errors.New(errorMessage)
	}

	return id, nil
}
