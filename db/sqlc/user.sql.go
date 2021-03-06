// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	email,
	password,
	first_name,
	last_name
) VALUES (
	$1, $2, $3, $4
) RETURNING id, email, password, first_name, last_name, created_at
`

type CreateUserParams struct {
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Email,
		arg.Password,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, first_name, last_name, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, password, first_name, last_name, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, email, password, first_name, last_name, created_at FROM users
ORDER BY email
LIMIT $1
OFFSET $2
`

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.getUsersStmt, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.FirstName,
			&i.LastName,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET first_name = $2, last_name = $3
WHERE id = $1
RETURNING id, email, password, first_name, last_name, created_at
`

type UpdateUserParams struct {
	ID        uuid.UUID      `json:"id"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserStmt, updateUser, arg.ID, arg.FirstName, arg.LastName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}
