package repository

import (
	app "avito-dynamic-segment-back"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user app.CreateUserInput) (int, error)
	GetAll() ([]app.User, error)
	GetById(userId int) (app.User, error)
	Delete(userId int) error
	Update(userId int, input app.UpdateUserInput) error

	AddSegments(userId int, input app.AddUserSegmentInput) error
	GetSegments(userId int) ([]app.Segment, error)
}

type Segment interface {
	Create(segment app.Segment) (int, error)
	GetAll() ([]app.Segment, error)
	GetById(segmentId int) (app.Segment, error)
	Delete(segmentId int) error
	Update(segmentId int, input app.UpdateSegmentInput) error
}

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Segment: NewSegmentPostgres(db),
	}
}
