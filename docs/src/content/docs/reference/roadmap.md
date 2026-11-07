---
title: Roadmap
description: Planned development phases and outcomes (no fixed release dates).
---

This page summarizes the project roadmap. The authoritative, evolving version lives in `ROADMAP.md` at the repository root. Phases are considered complete only when acceptance criteria are met and tested.

## Phase 0 — Project foundation

**Goal:** A clean, startable, testable Go project.
- Repository/package structure, controller entrypoint, logging/config basics.
- Start/shutdown lifecycle, initial tests and CI.
- **Acceptance:** controller starts locally; SIGINT/SIGTERM cause controlled shutdown; errors logged with context; `go test ./...` passes.

## Phase 1 — Controller core

**Goal:** Core application services and persistence.
- PostgreSQL connection/migrations; user/permission model.
- Node/service/template/instance models; event/job foundation; config validation.
- **Acceptance:** DB initialized/migrated; core objects manageable via application services; domain logic testable independently of HTTP/SSH/TUI.

## Phase 2 — Single node

**Goal:** Reliable controller–node connection and Docker control.
- Standalone node agent, gRPC protocol, secure registration/heartbeats.
- Docker integration, local cache, status reporting.
- **Acceptance:** node registers; controller can start/stop a test container; disconnects handled; containers not killed unexpectedly during temporary controller outage.

## Phase 3 — Minecraft instances and templates

**Goal:** Create Minecraft servers from versioned templates.
- Template manager, runtime/Minecraft config, instance lifecycle.
- Persistent data, health checks, log access, port reservation.
- **Acceptance:** create/start/monitor/remove from template; versioned templates; persistent data survives restarts; port conflicts handled.

## Phase 4 — Multi-node and scheduling

**Goal:** Distribute instances across hosts.
- Node resources/capacity, scheduling rules, labels/constraints.
- Central port ranges, private backend networking.
- **Acceptance:** suitable node selection; resource limits respected; ports uniquely reserved per node; unreachable nodes handled correctly.

## Phase 5 — Proxy integration

**Goal:** Automatically integrate instances into proxy networks.
- Velocity/BungeeCord bridges, proxy groups.
- Register after health check, reconciliation, draining.
- **Acceptance:** ready lobby registered; stopped instance removed cleanly; proxy restarts reconciled.

## Phase 6 — Management

**Goal:** Make the platform administrable.
- CLI, SSH server with public-key-only auth, multi-user TUI.
- Roles/permissions, audit events.
- **Acceptance:** parallel SSH sessions; authorization on every action; no SSH password auth.

## Phase 7 — Web interface and operations

**Goal:** External management and production-oriented operations.
- REST API + WebSocket, React/Next.js UI.
- TLS/reverse-proxy concept, monitoring/metrics.
- Install/backup/upgrade process.
- **Acceptance:** UI manages nodes/instances; live status updates; backup/restore tested; docs for install/upgrade.

## Later extensions

Evaluate after a stable base: controller HA, maintenance windows, instance migration, advanced resource/cost planning, additional runtimes, plugin/extension system.