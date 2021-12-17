-- name: CreateTask :one
INSERT INTO tasks (
	name,
	description,
	list_id
) VALUES (
	$1, $2, $3
) RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: GetTasksByListId :many
SELECT * FROM tasks
WHERE list_id = $1
ORDER BY created_at;

-- name: UpdateTask :one
UPDATE tasks
SET name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;
