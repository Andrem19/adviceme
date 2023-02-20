-- name: CreateBranch :one
INSERT INTO branches (
  branch_name
) VALUES (
  $1
) RETURNING *;

-- name: UpdateBranch :one
UPDATE branches
SET branch_name = $2
WHERE id = $1
RETURNING *;

-- name: GetBranch :one
SELECT * FROM branches
WHERE id = $1 LIMIT 1;

-- name: GetBranchFrom :many
SELECT * FROM branches
ORDER BY id;

-- name: DeleteBranch :exec
DELETE FROM branches
WHERE id = $1;