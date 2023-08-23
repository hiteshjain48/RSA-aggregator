-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, published_at,title, description, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostForUser :many
SELECT posts.* FROM posts JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
LIMIT $2;