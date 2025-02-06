package contract

import (
	"context"
	"test-project/internal/repository/postgres"

	"github.com/google/uuid"
)

type SongService interface {
	GetSongs(ctx context.Context, page, pageSize int, filter postgres.GetSongsByFieldParams) ([]postgres.Song, error)
	GetSongRow(ctx context.Context, id uuid.UUID, row int) (string, error)
	CreateSong(ctx context.Context, song, group string) (postgres.Song, error)
	UpdateSong(ctx context.Context, params postgres.UpdateSongParams) (postgres.Song, error)
	DeleteSong(ctx context.Context, id uuid.UUID) error
}
