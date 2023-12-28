-- name: CreatedFeed :one
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;



-- name: GetNextFeedToFetch :one
SELECT * from feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;

-- name: MarkFeedAsFetched :one
UPDATE feeds 
Set last_fetched_at =NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;