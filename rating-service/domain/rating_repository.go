package domain

type RatingRepository interface {
	Insert(rating *Rating) error
	GetAll() (*[]Rating, error)
	DeleteAll()
	Delete(rating *Rating) error
}
