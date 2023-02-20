-- name: CreateMessages :one
INSERT INTO messages (
  who_answer_id,
  who_ask_id,
  specialization,
  message_text
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetMessages :one
SELECT * FROM messages
WHERE id = $1 LIMIT 1;

-- name: GetMessagesFrom :many
SELECT * FROM messages
WHERE who_ask_id = $1
ORDER BY created_at;

-- name: GetMessagesTo :many
SELECT * FROM messages
WHERE who_answer_id = $1
ORDER BY created_at;

-- name: DeleteMessages :exec
DELETE FROM messages
WHERE id = $1;