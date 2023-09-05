// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, published_at,title, description, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, updated_at, published_at, title, description, url, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
	Title       string
	Description sql.NullString
	Url         string
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.PublishedAt,
		arg.Title,
		arg.Description,
		arg.Url,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PublishedAt,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.FeedID,
	)
	return i, err
}

const getPostForUser = `-- name: GetPostForUser :many
SELECT posts.id, posts.created_at, posts.updated_at, posts.published_at, posts.title, posts.description, posts.url, posts.feed_id FROM posts JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
LIMIT $2
`

type GetPostForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

func (q *Queries) GetPostForUser(ctx context.Context, arg GetPostForUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PublishedAt,
			&i.Title,
			&i.Description,
			&i.Url,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}