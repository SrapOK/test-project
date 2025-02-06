// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: song.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createSong = `-- name: CreateSong :one
INSERT INTO song (
    song_id, 
    song, 
    author,
    release_date,
    text,
    link
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING song_id, song, author, release_date, text, link
`

type CreateSongParams struct {
	SongID      uuid.UUID   `json:"song_id"`
	Song        string      `json:"song"`
	Author      string      `json:"author"`
	ReleaseDate pgtype.Date `json:"release_date"`
	Text        string      `json:"text"`
	Link        string      `json:"link"`
}

func (q *Queries) CreateSong(ctx context.Context, arg CreateSongParams) (Song, error) {
	row := q.db.QueryRow(ctx, createSong,
		arg.SongID,
		arg.Song,
		arg.Author,
		arg.ReleaseDate,
		arg.Text,
		arg.Link,
	)
	var i Song
	err := row.Scan(
		&i.SongID,
		&i.Song,
		&i.Author,
		&i.ReleaseDate,
		&i.Text,
		&i.Link,
	)
	return i, err
}

const deleteSong = `-- name: DeleteSong :exec
DELETE FROM song
WHERE song_id = $1
`

func (q *Queries) DeleteSong(ctx context.Context, songID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteSong, songID)
	return err
}

const findSong = `-- name: FindSong :many
SELECT song_id, song, author, release_date, text, link
FROM song
WHERE song LIKE $1
`

func (q *Queries) FindSong(ctx context.Context, song string) ([]Song, error) {
	rows, err := q.db.Query(ctx, findSong, song)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Song
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.SongID,
			&i.Song,
			&i.Author,
			&i.ReleaseDate,
			&i.Text,
			&i.Link,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthorsBySong = `-- name: GetAuthorsBySong :many
SELECT author
FROM song
WHERE song = $1
`

func (q *Queries) GetAuthorsBySong(ctx context.Context, song string) ([]string, error) {
	rows, err := q.db.Query(ctx, getAuthorsBySong, song)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var author string
		if err := rows.Scan(&author); err != nil {
			return nil, err
		}
		items = append(items, author)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSongById = `-- name: GetSongById :one
SELECT song_id, song, author, release_date, text, link
FROM song
WHERE song_id = $1
`

func (q *Queries) GetSongById(ctx context.Context, songID uuid.UUID) (Song, error) {
	row := q.db.QueryRow(ctx, getSongById, songID)
	var i Song
	err := row.Scan(
		&i.SongID,
		&i.Song,
		&i.Author,
		&i.ReleaseDate,
		&i.Text,
		&i.Link,
	)
	return i, err
}

const getSongs = `-- name: GetSongs :many
SELECT song_id, song, author, release_date, text, link
FROM song
LIMIT $1
OFFSET $2
`

type GetSongsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetSongs(ctx context.Context, arg GetSongsParams) ([]Song, error) {
	rows, err := q.db.Query(ctx, getSongs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Song
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.SongID,
			&i.Song,
			&i.Author,
			&i.ReleaseDate,
			&i.Text,
			&i.Link,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSongsByField = `-- name: GetSongsByField :many
SELECT song_id, song, author, release_date, text, link
FROM song
WHERE 
    song = COALESCE($3, song)
    AND author = COALESCE($4, author)
    AND release_date = COALESCE($5, release_date)
    AND text = COALESCE($6, text)
    AND link = COALESCE($7, link)
LIMIT $1
OFFSET $2
`

type GetSongsByFieldParams struct {
	Limit       int32       `json:"limit"`
	Offset      int32       `json:"offset"`
	Song        pgtype.Text `json:"song"`
	Author      pgtype.Text `json:"author"`
	ReleaseDate pgtype.Date `json:"release_date"`
	Text        pgtype.Text `json:"text"`
	Link        pgtype.Text `json:"link"`
}

func (q *Queries) GetSongsByField(ctx context.Context, arg GetSongsByFieldParams) ([]Song, error) {
	rows, err := q.db.Query(ctx, getSongsByField,
		arg.Limit,
		arg.Offset,
		arg.Song,
		arg.Author,
		arg.ReleaseDate,
		arg.Text,
		arg.Link,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Song
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.SongID,
			&i.Song,
			&i.Author,
			&i.ReleaseDate,
			&i.Text,
			&i.Link,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSong = `-- name: UpdateSong :one
UPDATE song
SET 
    song = $2, 
    author = $3, 
    release_date = $5,
    text = $4,
    link = $6
WHERE song_id = $1
RETURNING song_id, song, author, release_date, text, link
`

type UpdateSongParams struct {
	SongID      uuid.UUID   `json:"song_id"`
	Song        string      `json:"song"`
	Author      string      `json:"author"`
	Text        string      `json:"text"`
	ReleaseDate pgtype.Date `json:"release_date"`
	Link        string      `json:"link"`
}

func (q *Queries) UpdateSong(ctx context.Context, arg UpdateSongParams) (Song, error) {
	row := q.db.QueryRow(ctx, updateSong,
		arg.SongID,
		arg.Song,
		arg.Author,
		arg.Text,
		arg.ReleaseDate,
		arg.Link,
	)
	var i Song
	err := row.Scan(
		&i.SongID,
		&i.Song,
		&i.Author,
		&i.ReleaseDate,
		&i.Text,
		&i.Link,
	)
	return i, err
}
