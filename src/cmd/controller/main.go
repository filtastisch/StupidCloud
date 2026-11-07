package main

import (
	"StupidCloud/config"
	"StupidCloud/src/internal/app"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/allan-simon/go-singleinstance"
)

func main() {
	lockPath := filepath.Join(os.TempDir(), "stupidcloud-controller.lock")

	lockfile, err := singleinstance.CreateLockFile(lockPath)
	if err != nil {
		for i := 10; i > 0; i-- {
			fmt.Printf("\rStupidCloud already running! Exiting... (%d) ", i)
			time.Sleep(time.Second)
		}
		os.Exit(1)
	}

	defer func(lockfile *os.File) {
		err := lockfile.Close()
		if err != nil {
			fmt.Printf("Failed to close lockfile: %s\n", err)
		}
	}(lockfile)

	opsys := runtime.GOOS
	arch := runtime.GOARCH

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("System Information:", opsys, arch)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	cfg, err := config.Load("./data/config.toml")
	if err != nil {
		slog.Error("Failed to load config file:", err)
	}

	controller := app.NewController(logger, cfg)

	if err := controller.Start(ctx); err != nil {
		slog.Error("Controller stopped with error: ", "error", err)
		os.Exit(1)
	}

}
