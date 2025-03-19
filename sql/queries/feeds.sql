-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * from feeds
order by last_fetched_at asc nulls first
limit $1;

-- name: MarkFeedsAsFetched :one
update feeds
set last_fetched_at = now(),
updated_at = now()
WHERE id = $1
RETURNING *;