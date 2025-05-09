-- name: CreatePost :one
INSERT INTO posts (
id, created_at, updated_at, title, url, description, published_at, feed_id 
) VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts
ORDER BY created_at DESC
LIMIT $1;
