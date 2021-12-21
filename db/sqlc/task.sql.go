// Code generated by sqlc. DO NOT EDIT.
// source: task.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
	name,
	description,
	list_id
) VALUES (
	$1, $2, $3
) RETURNING id, name, description, list_id, completed, created_at
`

type CreateTaskParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	ListID      uuid.UUID      `json:"list_id"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.queryRow(ctx, q.createTaskStmt, createTask, arg.Name, arg.Description, arg.ListID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ListID,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteTaskStmt, deleteTask, id)
	return err
}

const getTask = `-- name: GetTask :one
SELECT id, name, description, list_id, completed, created_at FROM tasks
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id uuid.UUID) (Task, error) {
	row := q.queryRow(ctx, q.getTaskStmt, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ListID,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}

const getTasksByListId = `-- name: GetTasksByListId :many
SELECT id, name, description, list_id, completed, created_at FROM tasks
WHERE list_id = $1
ORDER BY created_at
`

func (q *Queries) GetTasksByListId(ctx context.Context, listID uuid.UUID) ([]Task, error) {
	rows, err := q.query(ctx, q.getTasksByListIdStmt, getTasksByListId, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ListID,
			&i.Completed,
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

const updateTask = `-- name: UpdateTask :one
UPDATE tasks
SET name = $2, description = $3
WHERE id = $1
RETURNING id, name, description, list_id, completed, created_at
`

type UpdateTaskParams struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error) {
	row := q.queryRow(ctx, q.updateTaskStmt, updateTask, arg.ID, arg.Name, arg.Description)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ListID,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}