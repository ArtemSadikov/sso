package logger

import (
	"log/slog"
	"os"

	"cinematic.com/sso/internal/infrastructure/config"
)

func SetupLogger(env config.Environment) *slog.Logger {
	var res *slog.Logger

	switch env {
	case config.Local:
		res = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.Dev:
		res = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.Prod:
		res = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return res
}
