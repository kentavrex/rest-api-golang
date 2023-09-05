package repository

import (
	"github.com/jmoiron/sqlx"
	app "rest-api-golang"
)

type User interface {
	Create(user app.CreateUserInput) (int, error)
	GetAll() ([]app.UserGet, error)
	GetById(userId int) (app.UserGet, error)
	Delete(userId int) error
	Update(userId int, input app.UpdateUserInput) error

	AddSegments(userId int, input app.AddUserSegmentInput) error
	GetSegments(userId int) ([]app.Segment, error)
	DeleteSegments(userId int) error
}

type Segment interface {
	Create(segment app.CreateSegmentInput) (int, error)
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
