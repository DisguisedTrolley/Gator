-- name: CreateFeed :one
INSERT INTO feeds (
	id, created_at, updated_at, name, url, user_id
) VALUES (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2,
	$3
) RETURNING *;

-- name: GetFeeds :many
SELECT
    feeds.id,
    feeds.name AS feed_name,
    feeds.url,
    feeds.user_id,
    users.name AS user_name
FROM
    feeds
JOIN
    users ON feeds.user_id = users.id;

