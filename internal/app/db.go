package app

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ConnectPostgres(dsn string) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	connetion, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключитсья к бд: %s", err.Error())
	}

	if err := connetion.Ping(ctx); err != nil {
		return nil, fmt.Errorf("не удалось проверить состояние подключения к бд: %s", err.Error())
	}

	m, err := migrate.New(
		"file://internal/repository/migrations",
		dsn,
	)
	if err != nil {
		return connetion, fmt.Errorf("не удолось начать миграцию: %s", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return connetion, fmt.Errorf("не удалось совершить миграцию: %s", err.Error())
	}

	return connetion, nil
}
