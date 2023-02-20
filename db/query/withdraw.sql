-- name: CreateWithdraw :one
INSERT INTO withdraw (
  from_account_id,
  amount_fiat,
  amount_coins,
  status_w
) VALUES (
  $1, $2, $3, 4
) RETURNING *;

-- name: UpdateWithdraw :one
UPDATE withdraw
SET amount_coins = $2, amount_fiat = $3
WHERE id = $1
RETURNING *;

-- name: UpdateWithdrawStatus :one
UPDATE withdraw
SET status_w = $2
WHERE id = $1
RETURNING *;

-- name: GetWithdraw :one
SELECT * FROM withdraw
WHERE id = $1 LIMIT 1;

-- name: GetAllWithdraw :many
SELECT * FROM withdraw
WHERE from_account_id = $1
ORDER BY created_at;