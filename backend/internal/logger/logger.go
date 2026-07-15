package logger

import (
	"log/slog"
	"os"
)

func Init() {
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)

	slog.SetDefault(logger)
}
