package contract

import (
	"context"
	"fmt"
	"strconv"
	"test-project/internal/contract/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type SongContract struct {
	songService SongService
}

func NewSongContract(songService SongService) *SongContract {

	return &SongContract{
		songService: songService,
	}
}

// GetSongs Get songs
//
//	@Summary		Get songs
//	@Description	Get songs
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int		false	"Page"						example(1)
//	@Param			pageSize	query		int		false	"Page Size"					example(10)
//	@Param			song		query		string	fales	"Filter for Song"			example(The Weeping)
//	@Param			group		query		string	fales	"Filter for Group"			example(DIM)
//	@Param			text		query		string	fales	"Filter for Text"			example(lorem ipsum...)
//	@Param			link		query		string	fales	"Filter for Link"			example(https://www.youtube.com/watch?v=_UCo04xk2Ik)
//	@Param			releaseDate	query		string	fales	"Filter for Release Date"	example(2025-01-03)
//	@Success		200			{array}		dto.GetSongsResponseDto
//	@Failure		400			{string}	string	"некорректные параметры"
//	@Failure		500			{string}	string	"не удалось получить песни"
//	@Router			/songs [get]
func (s *SongContract) GetSongs(c *fiber.Ctx) error {
	ctx := context.Background()

	params := new(dto.GetSongsDto)

	if err := c.QueryParser(params); err != nil {
		errMsg := fmt.Sprintf("некорректные параметры: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusBadRequest).SendString(errMsg)
	}

	filter := params.ToGetSongsByFieldParams()

	songs, err := s.songService.GetSongs(ctx, params.Page, params.PageSize, filter)
	if err != nil {
		errMsg := fmt.Sprintf("не удалось получить песни: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusInternalServerError).SendString(errMsg)
	}

	items := make([]dto.SongDto, len(songs))
	for i := range items {
		items[i].FromSong(songs[i])
	}

	response := dto.GetSongsResponseDto{
		Page:     params.Page,
		PageSize: params.PageSize,
		Items:    items,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// PostSong Post song
//
//	@Summary		Post song
//	@Description	Post song
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			song	body		string	true	"Song Title"	example(The Weeping)
//	@Param			group	body		string	true	"Group"			example(DIM)
//	@Success		200		{object}	dto.SongDto
//	@Failure		400		{string}	string	"некорректные параметры"
//	@Failure		500		{string}	string	"не удалось создать песню"
//	@Router			/songs [post]
func (s *SongContract) PostSong(c *fiber.Ctx) error {
	ctx := context.Background()

	songReq := new(dto.PostSongDto)

	if err := c.BodyParser(songReq); err != nil {
		errMsg := fmt.Sprintf("некорректные параметры: %s", err.Error())
		log.Info(errMsg)
		return c.Status(fiber.StatusBadRequest).SendString(errMsg)
	}

	song, err := s.songService.CreateSong(ctx, songReq.Song, songReq.Group)
	if err != nil {
		errMsg := fmt.Sprintf("не удалось создать песню: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusInternalServerError).SendString(errMsg)
	}

	var res dto.SongDto
	res.FromSong(song)

	return c.Status(fiber.StatusOK).JSON(res)
}

// PutSong Put song
//
//	@Summary		Put song
//	@Description	Put song
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			song		body		string	true	"Song Title"	example(The Weeping)
//	@Param			group		body		string	true	"Group"			example(DIM)
//	@Param			releaseDate	body		string	true	"Release Date"	example(2024-04-17)
//	@Param			text		body		string	true	"Text"			example(lorem ipsum...)
//	@Param			link		body		string	true	"Link"			example(https://www.youtube.com/watch?v=_UCo04xk2Ik)
//	@Success		200			{object}	dto.SongDto
//	@Failure		400			{string}	string	"некорректные параметры"
//	@Failure		500			{string}	string
//	@Router			/songs [put]
func (s *SongContract) PutSong(c *fiber.Ctx) error {
	ctx := context.Background()
	songReq := new(dto.PutSongDto)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		errMsg := fmt.Sprintf("не удалось декодировать uuid: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusBadRequest).SendString(errMsg)
	}

	if err := c.BodyParser(songReq); err != nil {
		errMsg := fmt.Sprintf("некорректные параметры: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusBadRequest).SendString(errMsg)
	}

	params, err := songReq.ToUpdateSongParams(id)
	if err != nil {
		errMsg := fmt.Sprintf("параметры не прошли валидацию: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusBadRequest).SendString(errMsg)
	}

	song, err := s.songService.UpdateSong(ctx, params)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var res dto.SongDto
	res.FromSong(song)

	return c.Status(fiber.StatusOK).JSON(res)
}

// DeleteSong Delete song
//
//	@Summary		Delete song
//	@Description	Delete song
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Song UUID"	example(048fec20-b500-4d44-8698-db46f7d86ae8)
//	@Success		200	{string}	string	"песня удалена"
//	@Failure		400	{string}	string	"не удалось декодировать uuid"
//	@Failure		500	{string}	string	"не удалось удалить песню"
//	@Router			/songs/{id} [delete]
func (s *SongContract) DeleteSong(c *fiber.Ctx) error {
	ctx := context.Background()

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		errMsg := fmt.Sprintf("не удалось декодировать uuid: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusBadRequest).SendString(errMsg)
	}

	if err := s.songService.DeleteSong(ctx, id); err != nil {
		errMsg := fmt.Sprintf("не удалось удалить песню: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusInternalServerError).SendString(errMsg)
	}

	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("песня удалена: %s", id.String()))
}

// GetSongRow Get song's row
//
//	@Summary		Get song's row
//	@Description	Get song's row
//	@Tags			songs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Song UUID"		example(048fec20-b500-4d44-8698-db46f7d86ae8)
//	@Param			row	path		string	true	"Song's row"	example(1)
//	@Success		200	{string}	string	"куплет"
//	@Failure		400	{string}	string	"не удалось декодировать uuid"
//	@Failure		500	{string}	string	"не удалось получить куплет"
//	@Router			/songs/{id}/{row} [get]
func (s *SongContract) GetSongRow(c *fiber.Ctx) error {
	ctx := context.Background()

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		errMsg := fmt.Sprintf("не удалось декодировать uuid: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusBadRequest).SendString(errMsg)
	}

	row, err := strconv.Atoi(c.Params("row", "1"))
	if err != nil {
		errMsg := fmt.Sprintf("не удалось декодировать row: %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusBadRequest).SendString(errMsg)
	}

	result, err := s.songService.GetSongRow(ctx, id, row)
	if err != nil {
		errMsg := fmt.Sprintf(": %s", err.Error())
		log.Error(errMsg)
		return c.Status(fiber.StatusInternalServerError).SendString(errMsg)
	}

	return c.Status(fiber.StatusOK).SendString(result)
}
