package song_details

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"test-project/internal/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

type SongDetailsRepository struct {
	url string
}

func New(url string) *SongDetailsRepository {
	return &SongDetailsRepository{url: url}
}

func (s *SongDetailsRepository) GetSongDetails(song, group string) (postgres.Song, error) {
	var songDetails postgres.Song
	agent := fiber.Get(s.url)

	agent.QueryString(fmt.Sprintf("song=%s", song))
	agent.QueryString(fmt.Sprintf("group=%s", group))

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return songDetails, errors.Join(errs...)
	}

	if statusCode != http.StatusOK {
		return songDetails, errors.New("некорректный код ответа")
	}

	if err := json.Unmarshal(body, &songDetails); err != nil {
		return songDetails, err
	}

	return songDetails, nil
}
