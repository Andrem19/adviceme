// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: branches.sql

package db

import (
	"context"
)

const createBranch = `-- name: CreateBranch :one
INSERT INTO branches (
  branch_name
) VALUES (
  $1
) RETURNING id, branch_name
`

func (q *Queries) CreateBranch(ctx context.Context, branchName string) (Branch, error) {
	row := q.db.QueryRowContext(ctx, createBranch, branchName)
	var i Branch
	err := row.Scan(&i.ID, &i.BranchName)
	return i, err
}

const deleteBranch = `-- name: DeleteBranch :exec
DELETE FROM branches
WHERE id = $1
`

func (q *Queries) DeleteBranch(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBranch, id)
	return err
}

const getBranch = `-- name: GetBranch :one
SELECT id, branch_name FROM branches
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBranch(ctx context.Context, id int64) (Branch, error) {
	row := q.db.QueryRowContext(ctx, getBranch, id)
	var i Branch
	err := row.Scan(&i.ID, &i.BranchName)
	return i, err
}

const getBranchFrom = `-- name: GetBranchFrom :many
SELECT id, branch_name FROM branches
ORDER BY id
`

func (q *Queries) GetBranchFrom(ctx context.Context) ([]Branch, error) {
	rows, err := q.db.QueryContext(ctx, getBranchFrom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Branch{}
	for rows.Next() {
		var i Branch
		if err := rows.Scan(&i.ID, &i.BranchName); err != nil {
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

const updateBranch = `-- name: UpdateBranch :one
UPDATE branches
SET branch_name = $2
WHERE id = $1
RETURNING id, branch_name
`

type UpdateBranchParams struct {
	ID         int64  `json:"id"`
	BranchName string `json:"branch_name"`
}

func (q *Queries) UpdateBranch(ctx context.Context, arg UpdateBranchParams) (Branch, error) {
	row := q.db.QueryRowContext(ctx, updateBranch, arg.ID, arg.BranchName)
	var i Branch
	err := row.Scan(&i.ID, &i.BranchName)
	return i, err
}