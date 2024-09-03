// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package persistence

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  full_name, email, password
) VALUES (
  $1, $2, $3
) RETURNING id, full_name, email, password, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	FullName string         `db:"full_name" json:"full_name"`
	Email    string         `db:"email" json:"email"`
	Password sql.NullString `db:"password" json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.FullName, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const existsAnUserUsingTheSameEmail = `-- name: ExistsAnUserUsingTheSameEmail :one
SELECT count(1) > 0 as exists FROM users WHERE email = $1 and deleted_at is null
`

func (q *Queries) ExistsAnUserUsingTheSameEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsAnUserUsingTheSameEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getLastUser = `-- name: GetLastUser :one
SELECT id, full_name, email, password, created_at, updated_at, deleted_at FROM users where deleted_at is null ORDER BY id desc LIMIT 1
`

func (q *Queries) GetLastUser(ctx context.Context) (*User, error) {
	row := q.db.QueryRowContext(ctx, getLastUser)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getUser = `-- name: GetUser :one
SELECT id, full_name, email, password, created_at, updated_at, deleted_at FROM users WHERE id = $1 and deleted_at is null
`

func (q *Queries) GetUser(ctx context.Context, id int32) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, full_name, email, password, created_at, updated_at, deleted_at FROM users WHERE email = $1 and deleted_at is null
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
