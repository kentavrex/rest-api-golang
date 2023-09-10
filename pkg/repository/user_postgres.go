package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	app "rest-api-golang"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user app.CreateUserInput) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var userId int
	createUserQuery := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id;", usersTable)
	row := tx.QueryRow(createUserQuery, user.Name)
	if err := row.Scan(&userId); err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, segmentId := range *user.Segments {
		createUsersSegmentsQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2);", usersSegmentsTable)
		_, err := tx.Exec(createUsersSegmentsQuery, userId, segmentId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return userId, tx.Commit()
}

func (r *UserPostgres) GetAll() ([]app.UserGet, error) {
	var users []app.UserGet

	query := fmt.Sprintf("SELECT * FROM %s", usersTable)

	err := r.db.Select(&users, query)

	for idx, user_x := range users {
		var userSegments []app.UsersSegments
		getUserSegmentsQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", usersSegmentsTable)
		err = r.db.Select(&userSegments, getUserSegmentsQuery, user_x.Id)

		var segmentsIds []int
		for _, userSegment := range userSegments {
			segmentsIds = append(segmentsIds, userSegment.SegmentId)
		}

		users[idx].Segments = segmentsIds
	}

	return users, err
}

func (r *UserPostgres) GetById(userId int) (app.UserGet, error) {
	var user app.UserGet
	getUserQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, getUserQuery, userId)

	var userSegments []app.UsersSegments
	getUserSegmentsQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", usersSegmentsTable)
	err = r.db.Select(&userSegments, getUserSegmentsQuery, userId)

	var segmentsIds []int
	for _, userSegment := range userSegments {
		segmentsIds = append(segmentsIds, userSegment.SegmentId)
	}

	user.Segments = segmentsIds

	return user, err
}

func (r *UserPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)

	_, err := r.db.Exec(query, userId)

	return err
}

func (r *UserPostgres) Update(userId int, input app.UpdateUserInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	updateUserQuery := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", usersTable)
	_, err = tx.Exec(updateUserQuery, input.Name, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	removeOldSegmentsQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersSegmentsTable)
	_, err = tx.Exec(removeOldSegmentsQuery, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, segmentId := range *input.Segments {
		createUsersSegmentsQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2);", usersSegmentsTable)
		_, err = tx.Exec(createUsersSegmentsQuery, userId, segmentId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *UserPostgres) AddSegments(userId int, input app.AddUserSegmentInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, segmentId := range *input.Segments {
		query := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2);", usersSegmentsTable)
		_, err = tx.Exec(query, userId, segmentId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *UserPostgres) GetSegments(userId int) ([]app.Segment, error) {
	var userSegments []app.Segment
	query := fmt.Sprintf("SELECT * FROM %s WHERE id IN (SELECT segment_id FROM %s WHERE user_id = $1)", segmentsTable, usersSegmentsTable)
	err := r.db.Select(&userSegments, query, userId)

	return userSegments, err
}

func (r *UserPostgres) DeleteSegments(userId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	removeOldSegmentsQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersSegmentsTable)
	_, err = tx.Exec(removeOldSegmentsQuery, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *UserPostgres) DeleteSegment(userId int, segmentId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_id = $2 ", usersSegmentsTable)

	_, err := r.db.Exec(query, userId, segmentId)

	return err
}
