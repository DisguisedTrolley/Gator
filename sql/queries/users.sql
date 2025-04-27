-- name: CreateUser :one
INSERT INTO users (
	id, created_at, updated_at, name
) VALUES ( 
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1
)
RETURNING *;

-- name: GetUser :one
SELECT * from users
WHERE name = $1;

-- name: DeleteUsers :exec
delete from users;
