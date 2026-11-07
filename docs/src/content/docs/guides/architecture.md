---
title: Architecture
description: The intended controller/node design and what exists in the codebase today.
---

StupidCloud is designed around a **central controller** and **node agents** running on managed hosts. The controller is the source of truth for *desired state*; nodes apply work locally (primarily via Docker in the target design) and report *observed state*.

> **Implementation status:** The repository contains controller lifecycle scaffolding and manager placeholders. Full controller–node communication, Docker orchestration, persistence, proxy integration, and management interfaces are not complete. Treat sections below as the **target architecture** unless marked as implemented.

## High-level components

### Controller

Planned responsibilities:

- Accept and coordinate work for nodes/services/instances.
- Schedule instances onto suitable nodes (resources, labels, capacity).
- Manage templates, configuration, users, and permissions.
- Own the API surface (internal gRPC, web REST/WebSocket, SSH/TUI/CLI integrations).
- Persist authoritative state (planned: PostgreSQL).

**Current code:** `src/cmd/controller/main.go` starts the process, loads config, and runs `app.Controller.Start(ctx)`. `src/internal/app/controller.go` initializes managers (`service`, `node`, `proxy`) and holds fields for SSH/web/gRPC servers, but `Start` currently waits for shutdown.

### Node agent

Planned responsibilities:

- Run on each managed host with access to the local Docker engine.
- Establish an **outbound, authenticated** connection to the controller (planned: gRPC/Protobuf).
- Execute container operations (create/start/stop/remove, logs, health/state).
- Report heartbeats, capabilities/resources, and instance status.
- Cache artifacts/templates locally where appropriate.

**Current code:** `src/cmd/node/main.go` is a stub. Node behavior is a roadmap item.

### Management interfaces

Planned:

- **SSH** management (public-key oriented), hosting CLI/TUI experiences.
- **TUI** (Bubble Tea) and **CLI** over shared application services.
- **Web UI** (planned React/Next.js) backed by controller REST/WebSocket APIs.

These interfaces should share the same authorization and application-service logic.

## Communication and data (planned)

| Boundary | Planned technology | Notes |
|---|---|---|
| Controller ↔ Node | gRPC + Protocol Buffers | Authenticated/encrypted; idempotent operations; reconnection/backoff. |
| Web client ↔ Controller | REST API | Versioned (e.g., `/api/v1`). |
| Live updates | WebSocket | Authenticated; scoped by permissions. |
| Persistence | PostgreSQL | Migrations, constraints (e.g., port reservations). |
| Configuration | TOML today / model evolving | Current loader uses TOML; broader config model is planned. |

## Minecraft/proxy integration (planned)

StupidCloud targets Minecraft server workloads with **Velocity** and **BungeeCord** integration. A bridge component is expected to receive desired proxy state from the controller and reconcile it via proxy APIs, registering backends after health checks and draining on shutdown.

## Security boundaries

- Nodes should initiate connections to the controller (reduces inbound exposure on nodes).
- Internal traffic should be TLS-protected with mutual authentication where appropriate.
- Docker socket access is high-privilege and should be limited to the node agent.
- Backend Minecraft ports should not be exposed publicly unless explicitly intended.

See [Security](/guides/security/) for details and caveats.

## Related pages

- [Project status](/reference/project-status/)
- [Roadmap](/reference/roadmap/)
- [Glossary](/reference/glossary/)