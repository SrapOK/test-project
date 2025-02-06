package app

import (
	"test-project/internal/contract"
	"test-project/internal/repository/postgres"
	"test-project/internal/repository/song_details"
	"test-project/internal/service"

	_ "test-project/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func (a *App) newRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: a.config.HttpServer.IdleTimeout,
	})

	queries := postgres.New(a.db)
	songDetailsRepo := song_details.New(a.config.SongDetails.Url)

	songService := service.NewSongService(queries, songDetailsRepo)
	songContract := contract.NewSongContract(songService)

	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api/v1")

	api.Get("/swagger/*", swagger.HandlerDefault)

	song := api.Group("songs")

	song.Get("/", songContract.GetSongs)
	song.Post("/", songContract.PostSong)
	song.Put("/:id", songContract.PutSong)
	song.Delete("/:id", songContract.DeleteSong)
	song.Get("/:id/:row", songContract.GetSongRow)

	return app
}
