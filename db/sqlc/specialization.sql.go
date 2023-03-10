// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: specialization.sql

package db

import (
	"context"
	"database/sql"
)

const createSpecialization = `-- name: CreateSpecialization :one
INSERT INTO specialization (
  branch,
  spec_name,
  descr
) VALUES (
  $1, $2, $3
) RETURNING id, branch, user_id, spec_name, descr, is_online
`

type CreateSpecializationParams struct {
	Branch   int64  `json:"branch"`
	SpecName string `json:"spec_name"`
	Descr    string `json:"descr"`
}

func (q *Queries) CreateSpecialization(ctx context.Context, arg CreateSpecializationParams) (Specialization, error) {
	row := q.db.QueryRowContext(ctx, createSpecialization, arg.Branch, arg.SpecName, arg.Descr)
	var i Specialization
	err := row.Scan(
		&i.ID,
		&i.Branch,
		&i.UserID,
		&i.SpecName,
		&i.Descr,
		&i.IsOnline,
	)
	return i, err
}

const deleteSpecialization = `-- name: DeleteSpecialization :exec
DELETE FROM specialization
WHERE id = $1
`

func (q *Queries) DeleteSpecialization(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSpecialization, id)
	return err
}

const getSpecialization = `-- name: GetSpecialization :one
SELECT id, branch, user_id, spec_name, descr, is_online FROM specialization
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSpecialization(ctx context.Context, id int64) (Specialization, error) {
	row := q.db.QueryRowContext(ctx, getSpecialization, id)
	var i Specialization
	err := row.Scan(
		&i.ID,
		&i.Branch,
		&i.UserID,
		&i.SpecName,
		&i.Descr,
		&i.IsOnline,
	)
	return i, err
}

const getSpecializationFrom = `-- name: GetSpecializationFrom :many
SELECT id, branch, user_id, spec_name, descr, is_online FROM specialization
WHERE user_id = $1
ORDER BY id
`

func (q *Queries) GetSpecializationFrom(ctx context.Context, userID int64) ([]Specialization, error) {
	rows, err := q.db.QueryContext(ctx, getSpecializationFrom, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Specialization{}
	for rows.Next() {
		var i Specialization
		if err := rows.Scan(
			&i.ID,
			&i.Branch,
			&i.UserID,
			&i.SpecName,
			&i.Descr,
			&i.IsOnline,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSpecialization = `-- name: UpdateSpecialization :one
UPDATE specialization
SET is_online = $2
WHERE id = $1
RETURNING id, branch, user_id, spec_name, descr, is_online
`

type UpdateSpecializationParams struct {
	ID       int64        `json:"id"`
	IsOnline sql.NullBool `json:"is_online"`
}

func (q *Queries) UpdateSpecialization(ctx context.Context, arg UpdateSpecializationParams) (Specialization, error) {
	row := q.db.QueryRowContext(ctx, updateSpecialization, arg.ID, arg.IsOnline)
	var i Specialization
	err := row.Scan(
		&i.ID,
		&i.Branch,
		&i.UserID,
		&i.SpecName,
		&i.Descr,
		&i.IsOnline,
	)
	return i, err
}
