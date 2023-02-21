-- name: CreateUser :one
INSERT INTO user_account (
  nickname,
  email,
  balance,
  hashed_password,
  resp
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM user_account
WHERE id = $1 LIMIT 1;

-- name: GetUserAccount :one
SELECT * FROM user_account
WHERE nickname = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM user_account
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM user_account
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: SubUserBalance :one
UPDATE user_account
SET balance = balance - $2
WHERE id = $1
RETURNING *;

-- name: AddUserBalance :one
UPDATE user_account
SET balance = balance + $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM user_account
WHERE id = $1;