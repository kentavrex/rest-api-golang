package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	app "rest-api-golang"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) Create(segment app.Segment) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (slug) values($1) RETURNING id", segmentsTable)

	row := r.db.QueryRow(query, segment.Slug)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SegmentPostgres) GetAll() ([]app.Segment, error) {
	var segments []app.Segment

	query := fmt.Sprintf("SELECT * FROM %s", segmentsTable)

	err := r.db.Select(&segments, query)

	return segments, err
}

func (r *SegmentPostgres) GetById(segmentId int) (app.Segment, error) {
	var segment app.Segment

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", segmentsTable)

	err := r.db.Get(&segment, query, segmentId)

	return segment, err
}

func (r *SegmentPostgres) Delete(segmentId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", segmentsTable)

	_, err := r.db.Exec(query, segmentId)

	return err
}

func (r *SegmentPostgres) Update(segmentId int, input app.UpdateSegmentInput) error {
	query := fmt.Sprintf("UPDATE %s SET slug = $1 WHERE id = $2", segmentsTable)

	_, err := r.db.Exec(query, input.Slug, segmentId)
	return err

}
