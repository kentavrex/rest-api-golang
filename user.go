package app

import (
	"errors"
	"github.com/lib/pq"
)

type User struct {
	Id       int           `json:"id" db:"id"`
	Name     string        `json:"name" db:"name" binding:"required"`
	Segments pq.Int64Array `json:"segments" db:"segments"`
}

type CreateUserInput struct {
	Name     *string        `json:"name" db:"name"`
	Segments *pq.Int64Array `json:"segments" db:"segments"`
}

type UpdateUserInput struct {
	Name     *string        `json:"name" db:"name"`
	Segments *pq.Int64Array `json:"segments" db:"segments"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type AddUserSegmentInput struct {
	Segments *pq.Int64Array `json:"segments" db:"segments" binding:"required"`
}

func (i AddUserSegmentInput) Validate() error {
	if i.Segments == nil {
		return errors.New("add segments structure has no values")
	}

	return nil
}
