---
title: Development
description: Workflow, conventions, and practical guidance for contributing to StupidCloud.
---

## Project conventions

- **Module/imports:** The Go module is `StupidCloud`. The code lives under `src/` and `config/`, so imports look like `StupidCloud/src/internal/app` or `StupidCloud/config`.
- **Binaries:** Entrypoints live in `src/cmd/controller` and `src/cmd/node` (node is currently a stub).
- **Internal packages:** Implementation details belong under `src/internal/...`.

## Local workflow

Run checks before committing:

```bash
gofmt -w ./src ./config
go test ./...
go vet ./...
```

Build/run:

```bash
go run ./src/cmd/controller
go build ./src/cmd/controller
```

## Go guidelines

- Keep packages small and focused.
- Return errors explicitly; wrap with context using `%w` when helpful.
- Pass `context.Context` as the first parameter for cancellable/timeout-aware operations.
- Avoid unnecessary globals; inject dependencies via constructors.
- Start goroutines only with a clear lifecycle and cancellation path.
- Define interfaces at the point of use, not speculatively.

## Lifecycle and shutdown

Components that run continuously should respect context cancellation and have well-defined `Start`/shutdown behavior. A production-ready controller should:

- initialize components in a defined order,
- roll back partially started components on startup failure,
- shut down in reverse order with a timeout.

The current controller scaffold implements signal-based cancellation and a basic start/shutdown log path; full component orchestration is still in progress.

## Testing

- Prefer unit tests for domain/application logic.
- Use integration tests for infrastructure (Docker, PostgreSQL) when those features are implemented.
- Test error paths, cancellation, retries/idempotency (important for orchestration), and concurrency where relevant.

## Dependencies

Add dependencies only when needed for current work. After changes:

```bash
go mod tidy
```

Commit `go.mod` and `go.sum` together. Keep `go` directive/toolchain changes intentional.

## Documentation

- This site lives in `docs/` (Astro + Starlight) with English (`root`) and German (`de`) content.
- Keep docs aligned with the **actual** implementation state; label planned features clearly.
- Architectural decisions with long-term impact should be recorded as ADRs (the README references `docs/decisions/`; create that structure when formalizing decisions).