-- name: CreateFeedFollow :one
WITH inserted AS (
  INSERT INTO feed_follows (
    id, created_at, updated_at, user_id, feed_id
  ) VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
  )
  RETURNING id, user_id, feed_id
)
SELECT
  inserted.id,
  users.name AS user_name,
  feeds.name AS feed_name
FROM
  inserted
JOIN users ON inserted.user_id = users.id
JOIN feeds ON inserted.feed_id = feeds.id;


-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM 
    feed_follows
JOIN users ON feed_follows.user_id = users.id
JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;
