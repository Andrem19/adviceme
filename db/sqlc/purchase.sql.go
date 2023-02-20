// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: purchase.sql

package db

import (
	"context"
)

const createPurchase = `-- name: CreatePurchase :one
INSERT INTO purchase (
  from_account_id,
  amount_fiat,
  amount_coins,
  status_p
) VALUES (
  $1, $2, $3, $4
) RETURNING id, from_account_id, amount_fiat, amount_coins, status_p, created_at
`

type CreatePurchaseParams struct {
	FromAccountID int64   `json:"from_account_id"`
	AmountFiat    float64 `json:"amount_fiat"`
	AmountCoins   int64   `json:"amount_coins"`
	StatusP       Status  `json:"status_p"`
}

func (q *Queries) CreatePurchase(ctx context.Context, arg CreatePurchaseParams) (Purchase, error) {
	row := q.db.QueryRowContext(ctx, createPurchase,
		arg.FromAccountID,
		arg.AmountFiat,
		arg.AmountCoins,
		arg.StatusP,
	)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.AmountFiat,
		&i.AmountCoins,
		&i.StatusP,
		&i.CreatedAt,
	)
	return i, err
}

const getPurchase = `-- name: GetPurchase :one
SELECT id, from_account_id, amount_fiat, amount_coins, status_p, created_at FROM purchase
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPurchase(ctx context.Context, id int64) (Purchase, error) {
	row := q.db.QueryRowContext(ctx, getPurchase, id)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.AmountFiat,
		&i.AmountCoins,
		&i.StatusP,
		&i.CreatedAt,
	)
	return i, err
}

const getPurchases = `-- name: GetPurchases :many
SELECT id, from_account_id, amount_fiat, amount_coins, status_p, created_at FROM purchase
WHERE from_account_id = $1
ORDER BY created_at
`

func (q *Queries) GetPurchases(ctx context.Context, fromAccountID int64) ([]Purchase, error) {
	rows, err := q.db.QueryContext(ctx, getPurchases, fromAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Purchase{}
	for rows.Next() {
		var i Purchase
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.AmountFiat,
			&i.AmountCoins,
			&i.StatusP,
			&i.CreatedAt,
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

const updatePurchase = `-- name: UpdatePurchase :one
UPDATE purchase
SET amount_coins = $2, amount_fiat = $3
WHERE id = $1
RETURNING id, from_account_id, amount_fiat, amount_coins, status_p, created_at
`

type UpdatePurchaseParams struct {
	ID          int64   `json:"id"`
	AmountCoins int64   `json:"amount_coins"`
	AmountFiat  float64 `json:"amount_fiat"`
}

func (q *Queries) UpdatePurchase(ctx context.Context, arg UpdatePurchaseParams) (Purchase, error) {
	row := q.db.QueryRowContext(ctx, updatePurchase, arg.ID, arg.AmountCoins, arg.AmountFiat)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.AmountFiat,
		&i.AmountCoins,
		&i.StatusP,
		&i.CreatedAt,
	)
	return i, err
}

const updatePurchaseStatus = `-- name: UpdatePurchaseStatus :one
UPDATE purchase
SET status_p = $2
WHERE id = $1
RETURNING id, from_account_id, amount_fiat, amount_coins, status_p, created_at
`

type UpdatePurchaseStatusParams struct {
	ID      int64  `json:"id"`
	StatusP Status `json:"status_p"`
}

func (q *Queries) UpdatePurchaseStatus(ctx context.Context, arg UpdatePurchaseStatusParams) (Purchase, error) {
	row := q.db.QueryRowContext(ctx, updatePurchaseStatus, arg.ID, arg.StatusP)
	var i Purchase
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.AmountFiat,
		&i.AmountCoins,
		&i.StatusP,
		&i.CreatedAt,
	)
	return i, err
}
