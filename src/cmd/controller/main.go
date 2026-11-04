package main

import (
	"StupidCloud/src/internal/app"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	controller := app.NewController(logger)
	if err := controller.Start(ctx); err != nil {
		slog.Error("Controller stopped with error: ", "error", err)
		os.Exit(1)
	}

}
