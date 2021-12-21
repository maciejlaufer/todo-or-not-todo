// Code generated by sqlc. DO NOT EDIT.
// source: list.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createList = `-- name: CreateList :one
INSERT INTO lists (
	name,
	creator_id
) VALUES (
	$1, $2
) RETURNING id, name, creator_id, created_at
`

type CreateListParams struct {
	Name      string    `json:"name"`
	CreatorID uuid.UUID `json:"creator_id"`
}

func (q *Queries) CreateList(ctx context.Context, arg CreateListParams) (List, error) {
	row := q.queryRow(ctx, q.createListStmt, createList, arg.Name, arg.CreatorID)
	var i List
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatorID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteList = `-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1
`

func (q *Queries) DeleteList(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteListStmt, deleteList, id)
	return err
}

const getList = `-- name: GetList :one
SELECT id, name, creator_id, created_at FROM lists
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetList(ctx context.Context, id uuid.UUID) (List, error) {
	row := q.queryRow(ctx, q.getListStmt, getList, id)
	var i List
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatorID,
		&i.CreatedAt,
	)
	return i, err
}

const getListsByCreatorId = `-- name: GetListsByCreatorId :many
SELECT id, name, creator_id, created_at FROM lists
WHERE creator_id = $1
ORDER BY created_at
`

func (q *Queries) GetListsByCreatorId(ctx context.Context, creatorID uuid.UUID) ([]List, error) {
	rows, err := q.query(ctx, q.getListsByCreatorIdStmt, getListsByCreatorId, creatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []List{}
	for rows.Next() {
		var i List
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatorID,
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

const getListsForUser = `-- name: GetListsForUser :many
SELECT l.id, l.name, l.creator_id, l.created_at FROM lists l
INNER JOIN users_lists ul ON ul.list_id = l.id
WHERE ul.user_id = $1
ORDER BY l.created_at
`

func (q *Queries) GetListsForUser(ctx context.Context, userID uuid.UUID) ([]List, error) {
	rows, err := q.query(ctx, q.getListsForUserStmt, getListsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []List{}
	for rows.Next() {
		var i List
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatorID,
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

const updateList = `-- name: UpdateList :one
UPDATE lists
SET name = $2
WHERE id = $1
RETURNING id, name, creator_id, created_at
`

type UpdateListParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (q *Queries) UpdateList(ctx context.Context, arg UpdateListParams) (List, error) {
	row := q.queryRow(ctx, q.updateListStmt, updateList, arg.ID, arg.Name)
	var i List
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatorID,
		&i.CreatedAt,
	)
	return i, err
}