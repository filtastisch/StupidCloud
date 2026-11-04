---
title: Project status and technology
description: Current maturity, intended capabilities, and planned technology choices.
---

## Maturity

StupidCloud is in an early development stage. The repository and plans are evolving, so the capabilities below describe the intended direction and should not be read as a guarantee that a feature is available today.

## Intended capabilities

- Manage multiple remote hosts from a central controller.
- Run individual Minecraft server instances in Docker containers.
- Manage templates, configuration, scheduling, and ports centrally.
- Connect proxy networks through Velocity and BungeeCord integrations.
- Provide SSH based CLI and TUI tools and a separate web interface.

## Planned technology

| Area | Planned choice |
| --- | --- |
| Backend | Go |
| Containers | Docker Engine / Moby |
| Internal communication | gRPC and Protocol Buffers |
| Web API | REST |
| Live updates | WebSocket |
| Database | PostgreSQL |
| Configuration | YAML |
| Terminal management | Charmbracelet SSH and Bubble Tea |
| HTTP routing | Chi |
| Web frontend | React / Next.js |

Technology choices can change during development. Architectural decisions intended to be binding should be recorded in the repository's `docs/decisions/` directory.
