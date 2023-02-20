-- name: CreateEntries :one
INSERT INTO entries (
  user_id,
  amount,
  messages
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetAllUserEntries :many
SELECT * FROM entries
WHERE user_id = $1
ORDER BY id;