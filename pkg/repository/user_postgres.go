package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	app "rest-api-golang"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user app.CreateUserInput) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, segments) VALUES ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, pq.Array(*user.Segments))
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetAll() ([]app.UserGet, error) {
	var users []app.UserGet

	query := fmt.Sprintf("SELECT * FROM %s", usersTable)

	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserPostgres) GetById(userId int) (app.UserGet, error) {
	var user app.UserGet

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)

	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *UserPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)

	_, err := r.db.Exec(query, userId)

	return err
}

func (r *UserPostgres) Update(userId int, input app.UpdateUserInput) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, segments = $2 WHERE id = $3", usersTable)

	_, err := r.db.Exec(query, input.Name, pq.Array(*input.Segments), userId)

	return err

}

func (r *UserPostgres) AddSegments(userId int, input app.AddUserSegmentInput) error {
	query := fmt.Sprintf("UPDATE %s SET segments = ARRAY_CAT(segments, $1) WHERE id = $2", usersTable)

	_, err := r.db.Exec(query, pq.Array(*input.Segments), userId)

	return err

}

func (r *UserPostgres) GetSegments(userId int) ([]app.Segment, error) {
	var userSegmentsIds pq.Int64Array
	query := fmt.Sprintf("SELECT segments FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&userSegmentsIds, query, userId)

	var userSegments []app.Segment
	query = fmt.Sprintf("SELECT * FROM %s WHERE id = ANY($1)", segmentsTable)
	err = r.db.Select(&userSegments, query, userSegmentsIds)

	return userSegments, err
}

func (r *UserPostgres) DeleteSegments(userId int) error {
	query := fmt.Sprintf("UPDATE %s SET segments = [] WHERE id = $1", usersTable)

	_, err := r.db.Exec(query, userId)

	return err
}
