---
title: Project status
description: Current maturity, implemented foundations, and planned technology choices.
---

## Maturity

StupidCloud is in an **early development stage**. The repository and design are evolving. Features described as planned or targeted should not be treated as available unless confirmed in code and tests.

## What exists today (codebase snapshot)

- **Controller entrypoint** (`src/cmd/controller/main.go`): single-instance lock, `slog` logging, OS signal handling (`SIGINT`/`SIGTERM`), config load from `./data/config.toml`.
- **Controller scaffold** (`src/internal/app/controller.go`): manager initialization (`service`, `node`, `proxy`), placeholders for SSH/web/gRPC servers, basic start/wait-for-shutdown behavior.
- **Config loading** (`config/`): TOML parsing with embedded defaults.
- **Node entrypoint**: placeholder only.

Many subsystems (database persistence, node gRPC protocol, Docker orchestration, scheduling, proxy bridges, SSH/TUI/web management) are **planned** and tracked in `TODO.md` and `ROADMAP.md`.

## Intended capabilities (target)

- Manage multiple remote hosts from a central controller.
- Run individual Minecraft server instances in Docker containers.
- Manage templates, configuration, scheduling, and ports centrally.
- Integrate proxy networks via Velocity and BungeeCord.
- Provide SSH-based CLI/TUI and a separate web interface.

## Planned/selected technology

| Area | Planned or in use | Notes |
|---|---|---|
| Backend | Go | Module `StupidCloud`. |
| Containers | Docker Engine / Moby | Planned integration. |
| Internal communication | gRPC + Protocol Buffers | Planned controller–node protocol. |
| Web API | REST | Planned. |
| Live updates | WebSocket | Planned. |
| Database | PostgreSQL | Planned (`pgx` noted in TODO). |
| Configuration | TOML (current loader); model evolving | Current controller config is TOML via `config.Load`. |
| Terminal management | Charmbracelet SSH + Bubble Tea | Planned/deps present. |
| HTTP routing | Chi | Dependency present; full wiring planned. |
| Web frontend | React / Next.js | Planned; `docs/` site uses Astro/Starlight. |

Technology choices can change during development. Record binding decisions as ADRs under `docs/decisions/` when that process is formalized.