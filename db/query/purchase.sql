-- name: CreatePurchase :one
INSERT INTO purchase (
  from_account_id,
  amount_fiat,
  amount_coins,
  status_p
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdatePurchase :one
UPDATE purchase
SET amount_coins = $2, amount_fiat = $3
WHERE id = $1
RETURNING *;

-- name: UpdatePurchaseStatus :one
UPDATE purchase
SET status_p = $2
WHERE id = $1
RETURNING *;

-- name: GetPurchase :one
SELECT * FROM purchase
WHERE id = $1 LIMIT 1;

-- name: GetPurchases :many
SELECT * FROM purchase
WHERE from_account_id = $1
ORDER BY created_at;