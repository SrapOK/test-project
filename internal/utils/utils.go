package utils

import (
	"strings"

	flog "github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
)

func SetText(src string, dest *pgtype.Text) {
	dest.String = src
	dest.Valid = true
}

func StringEmpty(stringToCheck string) bool {
	return len(stringToCheck) == 0
}

func GetFlogLevel(logLevel string) flog.Level {
	var level flog.Level

	switch strings.ToLower(logLevel) {
	case "debug":
		level = flog.LevelDebug
	case "info":
		level = flog.LevelInfo
	case "warn":
		level = flog.LevelWarn
	case "error":
		level = flog.LevelError
	case "fatal":
		level = flog.LevelFatal
	default:
		level = flog.LevelDebug
	}

	return level
}
