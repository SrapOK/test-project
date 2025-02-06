-- name: GetSongs :many
SELECT *
FROM song
LIMIT $1
OFFSET $2;

-- name: GetSongsByField :many
SELECT *
FROM song
WHERE 
    song = COALESCE(sqlc.narg('song'), song)
    AND author = COALESCE(sqlc.narg('author'), author)
    AND release_date = COALESCE(sqlc.narg('release_date'), release_date)
    AND text = COALESCE(sqlc.narg('text'), text)
    AND link = COALESCE(sqlc.narg('link'), link)
LIMIT $1
OFFSET $2;


-- name: GetSongById :one
SELECT *
FROM song
WHERE song_id = $1;


-- name: GetAuthorsBySong :many
SELECT author
FROM song
WHERE song = $1;

-- name: FindSong :many
SELECT *
FROM song
WHERE song LIKE $1;

-- name: CreateSong :one
INSERT INTO song (
    song_id, 
    song, 
    author,
    release_date,
    text,
    link
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateSong :one
UPDATE song
SET 
    song = $2, 
    author = $3, 
    release_date = $5,
    text = $4,
    link = $6
WHERE song_id = $1
RETURNING *;

-- name: DeleteSong :exec
DELETE FROM song
WHERE song_id = $1;
