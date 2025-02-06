package dto

import (
	"fmt"
	"test-project/internal/repository/postgres"
	"test-project/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *SongDto) FromSong(song postgres.Song) {
	s.SongId = song.SongID.String()
	s.Group = song.Author
	s.Song = song.Song
	s.ReleaseDate = ReleaseDateToString(song.ReleaseDate)
	s.Text = song.Text
	s.Link = song.Link
}

func (s *SongDto) ToSong() (postgres.Song, error) {
	var song postgres.Song
	var err error

	song.Author = s.Group
	song.Song = s.Song
	song.Link = s.Link
	song.Text = s.Text

	song.SongID, err = uuid.Parse(s.SongId)
	if err != nil {
		return song, err
	}

	song.ReleaseDate, err = StringToReleaseDate(s.ReleaseDate)
	if err != nil {
		return song, err
	}

	return song, nil
}

func (p *PutSongDto) ToUpdateSongParams(songId uuid.UUID) (postgres.UpdateSongParams, error) {
	params := postgres.UpdateSongParams{
		SongID: songId,
		Song:   p.Song,
		Author: p.Group,
		Text:   p.Text,
		Link:   p.Link,
	}

	releaseDate, err := StringToReleaseDate(p.ReleaseDate)
	if err != nil {
		return params, err
	}

	params.ReleaseDate = releaseDate

	return params, nil
}

func (g *GetSongsDto) ToGetSongsByFieldParams() postgres.GetSongsByFieldParams {
	var params postgres.GetSongsByFieldParams
	p := &params

	if !utils.StringEmpty(g.Text) {
		utils.SetText(g.Text, &p.Text)
	}

	if !utils.StringEmpty(g.Group) {
		utils.SetText(g.Group, &p.Author)
	}

	if !utils.StringEmpty(g.Song) {
		utils.SetText(g.Song, &p.Song)
	}

	if !utils.StringEmpty(g.Link) {
		utils.SetText(g.Link, &p.Link)
	}

	if !utils.StringEmpty(g.ReleaseDate) {
		date, err := StringToReleaseDate(g.ReleaseDate)
		if err != nil {
			log.Debugf("не удалось замапить GetSongsDto в GetSongsByFieldParams: %s", err.Error())
		} else {
			p.ReleaseDate = date
		}
	}

	return params
}

func ReleaseDateToString(releaseDate pgtype.Date) string {
	return releaseDate.Time.Format(time.DateOnly)
}

func StringToReleaseDate(releaseDate string) (pgtype.Date, error) {
	var result pgtype.Date

	err := result.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", releaseDate)))
	if err != nil {
		return result, fmt.Errorf("не удалось декодировать releaseDate: %s", err.Error())
	}

	return result, nil
}
