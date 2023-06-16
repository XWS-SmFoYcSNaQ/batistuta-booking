package domain

import "github.com/google/uuid"

type RatingRepository interface {
	Insert(rating *Rating) error
	Update(rating *Rating) error
	GetAll() (*[]Rating, error)
	DeleteAll()
	Delete(rating *Rating) error
	GetById(id *uuid.UUID) (*Rating, error)
	GetByUserAndTarget(userId *uuid.UUID, targetId *uuid.UUID, targetType uint32) (*Rating, error)
	GetTargetAverage(targetId *uuid.UUID, targetType uint32) (float64, error)
	GetByTargetType(targetType uint32) (*[]Rating, error)
	GetByTargetId(targetId *uuid.UUID) (*[]Rating, error)
}
