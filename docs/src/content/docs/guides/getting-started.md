---
title: Getting started
description: Run the current StupidCloud controller scaffold and understand the repository layout.
---

This guide covers the **current** state of the project. StupidCloud is early-stage software; only foundational controller behavior is implemented today.

## Prerequisites

- **Go** version compatible with `go.mod` (the module currently declares `go 1.27`).
- **Git**.
- For future work (not required to run the current scaffold): **Docker Engine** for container integration and **PostgreSQL** for planned persistence.

## Clone and run the controller

From the repository root:

```bash
go run ./src/cmd/controller
```

What happens on startup (current implementation):

1. A single-instance lock file is created at `<system temp>/stupidcloud-controller.lock`. If another controller is already running, the process exits.
2. The default `slog` text logger is initialized.
3. Basic OS/architecture information is logged.
4. Configuration is loaded from `./data/config.toml`. If the file or directory does not exist, `config.Load` creates `./data/` and writes the embedded default configuration, then parses it.
5. The controller starts and waits for `SIGINT` or `SIGTERM`, then logs shutdown.

> The controller currently provides lifecycle/scaffolding behavior. gRPC/SSH/web servers, node orchestration, Docker operations, and database-backed features are **not** fully wired yet.

## Useful commands

```bash
# Run controller
go run ./src/cmd/controller

# Run all tests
go test ./...

# Static checks
go vet ./...

# Format code
gofmt -w ./src ./config

# Tidy dependencies
go mod tidy

# Build controller binary
go build ./src/cmd/controller
```

## Repository layout

```text
StupidCloud/
├── config/              # Config loading and embedded defaults (TOML)
├── data/                # Local runtime data (e.g., config.toml); created as needed
├── docs/                # Astro Starlight documentation site
├── src/
│   ├── cmd/
│   │   ├── controller/  # Controller entrypoint
│   │   └── node/        # Node entrypoint (currently a stub)
│   └── internal/
│       ├── app/         # Controller application lifecycle/scaffold
│       ├── node/        # Node manager (scaffold)
│       ├── proxy/       # Proxy manager (scaffold)
│       ├── service/     # Service manager (scaffold)
│       └── web/         # Web server package (scaffold)
├── go.mod
├── README.md
├── TODO.md
└── ROADMAP.md
```

The `src/` directory is a regular directory (not a special Go convention). Imports use the module path, e.g. `StupidCloud/src/internal/app`.

## Node agent

`src/cmd/node` is currently a placeholder (`package node` only). Node registration, gRPC protocol, and Docker execution are planned; see the [roadmap](/reference/roadmap/).

## Next steps

- [Architecture](/guides/architecture/) — target controller/node design.
- [Configuration](/guides/configuration/) — current TOML config and planned model.
- [Development](/guides/development/) — conventions and workflow.