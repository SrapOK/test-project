package service

import (
	"context"
	"fmt"
	"math"
	"strings"
	"test-project/internal/repository/postgres"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type SongService struct {
	songRepo        SongRepository
	songDetailsRepo SongDetailsRepository
}

func NewSongService(songRepo SongRepository, songDetailsRepo SongDetailsRepository) *SongService {
	return &SongService{
		songRepo:        songRepo,
		songDetailsRepo: songDetailsRepo,
	}
}

func (s *SongService) GetSongs(ctx context.Context, page, pageSize int, filter postgres.GetSongsByFieldParams) ([]postgres.Song, error) {
	limit := int32(math.Abs(float64(pageSize)))
	offset := int32(math.Abs(float64(page))-1) * limit
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	log.Debug(fmt.Sprintf("Get Songs params limit: %d offset: %d filter %v", limit, offset, filter))

	filter.Limit = limit
	filter.Offset = offset

	songs, err := s.songRepo.GetSongsByField(ctx, filter)

	log.Debug(fmt.Sprintf("Get Songs result: %v", songs))
	return songs, err
}

func (s *SongService) GetSong(ctx context.Context, id uuid.UUID) (postgres.Song, error) {
	log.Debug(fmt.Sprintf("Get Song params id: %s", id.String()))

	song, err := s.songRepo.GetSongById(ctx, id)
	if err != nil {
		return song, fmt.Errorf("не удалось получить песню: %s", err.Error())
	}

	log.Debug(fmt.Sprintf("Get Song result: %v", song))
	return song, nil
}

func (c *SongService) CreateSong(ctx context.Context, song, author string) (postgres.Song, error) {
	songDetails, err := c.songDetailsRepo.GetSongDetails(song, author)
	if err != nil {
		return songDetails, fmt.Errorf("не удалось получить подробную информацию о песне: %s", err.Error())
	}

	params := postgres.CreateSongParams{
		SongID:      uuid.New(),
		Song:        song,
		Author:      author,
		ReleaseDate: songDetails.ReleaseDate,
		Text:        songDetails.Text,
		Link:        songDetails.Link,
	}

	log.Debug(fmt.Sprintf("Create Song params: %v", params))

	result, err := c.songRepo.CreateSong(ctx, params)
	if err != nil {
		return songDetails, fmt.Errorf("не удалось создать запись о песне: %s", err.Error())
	}

	log.Debug(fmt.Sprintf("Create Song result: %v", result))
	return result, nil
}

func (c *SongService) DeleteSong(ctx context.Context, id uuid.UUID) error {
	log.Debug(fmt.Sprintf("Delete Song params id: %s", id.String()))

	if _, err := c.songRepo.GetSongById(ctx, id); err != nil {
		return fmt.Errorf("не найдена песня для удаления: %s", err.Error())
	}

	if err := c.songRepo.DeleteSong(ctx, id); err != nil {
		return fmt.Errorf("не удалось удалить песню: %s", err.Error())
	}

	return nil
}

func (c *SongService) UpdateSong(ctx context.Context, params postgres.UpdateSongParams) (postgres.Song, error) {
	log.Debug(fmt.Sprintf("Update Song params: %v", params))

	res, err := c.songRepo.UpdateSong(ctx, params)
	if err != nil {
		return res, fmt.Errorf("не удалось обновить песню: %s", err.Error())
	}

	log.Debug(fmt.Sprintf("Update Song result: %v", res))
	return res, nil
}

func (s *SongService) GetSongRow(ctx context.Context, id uuid.UUID, row int) (string, error) {
	var result string

	log.Debug(fmt.Sprintf("Get Song Row params id: %s row: %d", id.String(), row))

	song, err := s.songRepo.GetSongById(ctx, id)
	if err != nil {
		return result, fmt.Errorf("песня не найдена: %s", err.Error())
	}

	rows := strings.Split(song.Text, "\n")
	if row <= 0 || row > len(rows) {
		return result, fmt.Errorf("некорректный индекс строки для текста песни")
	}

	result = rows[row-1]

	log.Debug(fmt.Sprintf("Get Song Row result: %s", result))
	return result, nil
}
