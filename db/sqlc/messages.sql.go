// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: messages.sql

package db

import (
	"context"
)

const createMessages = `-- name: CreateMessages :one
INSERT INTO messages (
  who_answer_id,
  who_ask_id,
  specialization,
  message_text
) VALUES (
  $1, $2, $3, $4
) RETURNING id, who_ask_id, who_answer_id, specialization, message_text, created_at
`

type CreateMessagesParams struct {
	WhoAnswerID    int64  `json:"who_answer_id"`
	WhoAskID       int64  `json:"who_ask_id"`
	Specialization int64  `json:"specialization"`
	MessageText    string `json:"message_text"`
}

func (q *Queries) CreateMessages(ctx context.Context, arg CreateMessagesParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, createMessages,
		arg.WhoAnswerID,
		arg.WhoAskID,
		arg.Specialization,
		arg.MessageText,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.WhoAskID,
		&i.WhoAnswerID,
		&i.Specialization,
		&i.MessageText,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMessages = `-- name: DeleteMessages :exec
DELETE FROM messages
WHERE id = $1
`

func (q *Queries) DeleteMessages(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMessages, id)
	return err
}

const getMessages = `-- name: GetMessages :one
SELECT id, who_ask_id, who_answer_id, specialization, message_text, created_at FROM messages
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMessages(ctx context.Context, id int64) (Message, error) {
	row := q.db.QueryRowContext(ctx, getMessages, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.WhoAskID,
		&i.WhoAnswerID,
		&i.Specialization,
		&i.MessageText,
		&i.CreatedAt,
	)
	return i, err
}

const getMessagesFrom = `-- name: GetMessagesFrom :many
SELECT id, who_ask_id, who_answer_id, specialization, message_text, created_at FROM messages
WHERE who_ask_id = $1
ORDER BY created_at
`

func (q *Queries) GetMessagesFrom(ctx context.Context, whoAskID int64) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getMessagesFrom, whoAskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.WhoAskID,
			&i.WhoAnswerID,
			&i.Specialization,
			&i.MessageText,
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

const getMessagesTo = `-- name: GetMessagesTo :many
SELECT id, who_ask_id, who_answer_id, specialization, message_text, created_at FROM messages
WHERE who_answer_id = $1
ORDER BY created_at
`

func (q *Queries) GetMessagesTo(ctx context.Context, whoAnswerID int64) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getMessagesTo, whoAnswerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.WhoAskID,
			&i.WhoAnswerID,
			&i.Specialization,
			&i.MessageText,
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
