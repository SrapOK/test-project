package service

import (
	"context"
	"test-project/internal/repository/postgres"

	"github.com/google/uuid"
)

type SongRepository interface {
	GetSongs(ctx context.Context, params postgres.GetSongsParams) ([]postgres.Song, error)
	GetSongsByField(ctx context.Context, params postgres.GetSongsByFieldParams) ([]postgres.Song, error)
	GetSongById(ctx context.Context, id uuid.UUID) (postgres.Song, error)

	CreateSong(ctx context.Context, params postgres.CreateSongParams) (postgres.Song, error)
	UpdateSong(ctx context.Context, params postgres.UpdateSongParams) (postgres.Song, error)
	DeleteSong(ctx context.Context, id uuid.UUID) error
}

type SongDetailsRepository interface {
	GetSongDetails(song, author string) (postgres.Song, error)
}
