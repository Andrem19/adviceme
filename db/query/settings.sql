-- name: CreateSettings :one
INSERT INTO settings (
  rate
) VALUES (
  $1
) RETURNING *;

-- name: UpdateSettingsRate :one
UPDATE settings
SET rate = $2
WHERE id = $1
RETURNING *;

-- name: GetSettings :one
SELECT * FROM settings
WHERE id = $1 LIMIT 1;

-- name: GetAllSettings :many
SELECT * FROM settings
ORDER BY created_at;