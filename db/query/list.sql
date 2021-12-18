-- name: CreateList :one
INSERT INTO lists (
	name,
	creator_id
) VALUES (
	$1, $2
) RETURNING *;

-- name: GetList :one
SELECT * FROM lists
WHERE id = $1 LIMIT 1;

-- name: GetListsByCreatorId :many
SELECT * FROM lists
WHERE creator_id = $1
ORDER BY created_at;

-- name: GetListsForUser :many
SELECT l.* FROM lists l
INNER JOIN users_lists ul ON ul.list_id = l.id
WHERE ul.user_id = $1
ORDER BY l.created_at;

-- name: UpdateList :one
UPDATE lists
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1;
