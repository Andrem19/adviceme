-- name: CreateSpecialization :one
INSERT INTO specialization (
  branch,
  spec_name,
  descr
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateSpecialization :one
UPDATE specialization
SET is_online = $2
WHERE id = $1
RETURNING *;

-- name: GetSpecialization :one
SELECT * FROM specialization
WHERE id = $1 LIMIT 1;

-- name: GetSpecializationFrom :many
SELECT * FROM specialization
WHERE user_id = $1
ORDER BY id;

-- name: DeleteSpecialization :exec
DELETE FROM specialization
WHERE id = $1;