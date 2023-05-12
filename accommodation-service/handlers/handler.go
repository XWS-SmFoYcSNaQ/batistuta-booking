package handlers

import (
	"accommodation_service/model"
	"accommodation_service/proto/accommodation"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"log"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	DB *sql.DB
}

func (h AccommodationHandler) GetAll(ctx context.Context, request *accommodation.AM_GetAll_Request) (*accommodation.AM_GetAll_Response, error) {
	rows, err := h.DB.Query("SELECT * FROM Accommodation")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var accommodations []model.Accommodation
	for rows.Next() {
		var p model.Accommodation
		err := rows.Scan(&p.ID, &p.Name, &p.Benefits, &p.MinGuests, &p.MaxGuests)
		if err != nil {
			log.Fatal(err)
		}
		accommodations = append(accommodations, p)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	response, _ := json.Marshal(accommodations)

	return &accommodation.AM_GetAll_Response{
		Data: string(response),
	}, nil
}

func (h AccommodationHandler) Create(ctx context.Context, request *accommodation.AM_Create_Request) (*accommodation.AM_Create_Response, error) {
	stmt, err := h.DB.Prepare("INSERT INTO Accommodation VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	id := uuid.New()
	_, err = stmt.Exec(id, request.Name, request.Benefits, request.MinGuests, request.MaxGuests)
	if err != nil {
		log.Fatal(err)
	}

	return &accommodation.AM_Create_Response{
		Id: id.String(),
	}, nil
}
