package app

import (
	"context"
	"fmt"
	"test-project/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
)

type App struct {
	config config.Config
	router *fiber.App
	db     *pgx.Conn
}

func NewApp(config config.Config) (*App, error) {

	db, err := ConnectPostgres(config.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("не удаётся инициализировать приложение: %s", err.Error())
	}

	app := &App{
		config: config,
		db:     db,
	}

	router := app.newRouter()
	app.router = router

	return app, nil
}

func (a *App) Run() {
	if err := a.router.Listen(a.config.HttpServer.Address); err != nil {
		log.Info(fmt.Errorf("прекращение работы сервера: %s", err.Error()))
	}
}

func (a *App) Stop(ctx context.Context) {
	if err := a.router.ShutdownWithContext(ctx); err != nil {
		log.Info(fmt.Errorf("не удалось завершить работу сервера: %s", err.Error()))
	}

	if err := a.db.Close(ctx); err != nil {
		log.Info(fmt.Errorf("не удалось закрыть подключение к бд: %s", err.Error()))
	}

	log.Info("приложение успешно завершило свою работу")
}
