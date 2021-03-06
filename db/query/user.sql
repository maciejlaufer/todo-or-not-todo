-- name: CreateUser :one
INSERT INTO users (
	email,
	password,
	first_name,
	last_name
) VALUES (
	$1, $2, $3, $4
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY email
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2, last_name = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
