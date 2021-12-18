-- name: AddUserToList :one
INSERT INTO users_lists (
	user_id,
	list_id
) VALUES (
	$1, $2
) RETURNING *;

-- name: DeleteUserFromList :exec
DELETE FROM users_lists
WHERE user_id = $1 AND list_id = $2;