-- name: CreateList :one
INSERT INTO lists (
	name,
	user_id
) VALUES (
	$1, $2
) RETURNING *;

-- name: GetList :one
SELECT * FROM lists
WHERE id = $1 LIMIT 1;

-- name: GetLists :many
SELECT * FROM lists
ORDER BY created_at;

-- name: UpdateList :one
UPDATE lists
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1;
