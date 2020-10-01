// Code generated by sqlc. DO NOT EDIT.
// source: token.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (
    user_id,
    blacklisted,
    token,
    type,
    expires
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, user_id, blacklisted, token, type, expires, created_at, updated_at
`

type CreateTokenParams struct {
	UserID      int64          `json:"user_id"`
	Blacklisted sql.NullBool   `json:"blacklisted"`
	Token       sql.NullString `json:"token"`
	Type        sql.NullInt32  `json:"type"`
	Expires     time.Time      `json:"expires"`
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, createToken,
		arg.UserID,
		arg.Blacklisted,
		arg.Token,
		arg.Type,
		arg.Expires,
	)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Blacklisted,
		&i.Token,
		&i.Type,
		&i.Expires,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteToken = `-- name: DeleteToken :exec
DELETE FROM tokens
WHERE id = $1
`

func (q *Queries) DeleteToken(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteToken, id)
	return err
}

const getToken = `-- name: GetToken :one
SELECT id, user_id, blacklisted, token, type, expires, created_at, updated_at FROM tokens
WHERE token = $1 LIMIT 1
`

func (q *Queries) GetToken(ctx context.Context, token sql.NullString) (Token, error) {
	row := q.db.QueryRowContext(ctx, getToken, token)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Blacklisted,
		&i.Token,
		&i.Type,
		&i.Expires,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listToken = `-- name: ListToken :many
SELECT id, user_id, blacklisted, token, type, expires, created_at, updated_at FROM tokens
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTokenParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListToken(ctx context.Context, arg ListTokenParams) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listToken, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Blacklisted,
			&i.Token,
			&i.Type,
			&i.Expires,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateToken = `-- name: UpdateToken :one
UPDATE tokens
SET blacklisted = $2, expires = $3
WHERE id = $1
RETURNING id, user_id, blacklisted, token, type, expires, created_at, updated_at
`

type UpdateTokenParams struct {
	ID          int64        `json:"id"`
	Blacklisted sql.NullBool `json:"blacklisted"`
	Expires     time.Time    `json:"expires"`
}

func (q *Queries) UpdateToken(ctx context.Context, arg UpdateTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, updateToken, arg.ID, arg.Blacklisted, arg.Expires)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Blacklisted,
		&i.Token,
		&i.Type,
		&i.Expires,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}