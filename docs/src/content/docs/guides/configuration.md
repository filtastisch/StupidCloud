---
title: Configuration
description: Current controller configuration behavior and the planned configuration model.
---

## Current behavior (implemented)

The controller entrypoint loads configuration at startup:

```go
cfg, err := config.Load("./data/config.toml")
```

Implementation details (`config/config.go`):

- If `./data/` (or the config directory) does not exist, it is created.
- If `./data/config.toml` does not exist, the embedded default file `defaults/config.toml` is written to that path.
- The file is parsed as **TOML** using `github.com/pelletier/go-toml/v2`.

### Current configuration structure

The current `Config` model includes these top-level sections:

| Section | Purpose (current model) |
|---|---|
| `application` | Name, version, environment, debug flag |
| `controller` | Controller host/port |
| `database` | Database type/path (model field; full DB integration is planned) |
| `logging` | Level/format |
| `network` | Bind/advertise addresses |
| `nodes` | Static node entries (name, address, port, enabled) |

> **Note:** Not all fields are actively used by running components yet, because much of the system is still scaffolded. Treat this as the current config shape, not a guarantee of full runtime behavior.

## Planned configuration model (target)

The target system separates concerns more explicitly (subject to change):

- **Controller config:** listeners (gRPC/HTTP/SSH), database DSN, storage, auth/cluster settings.
- **Node config:** node identity, controller endpoint, local Docker socket, cache settings.
- **Templates:** runtime/Minecraft defaults, files, health checks, proxy/network exposure.
- **Service/instance overrides:** concrete deviations per service/instance.
- **Secrets:** managed separately; never stored in Git or public templates.

Older planning documents showed YAML examples; the implementation currently uses TOML for controller config loading. Final formats and schemas will be decided as features are implemented.

## Practical guidance

- Validate configuration on startup and report clear errors (partially planned).
- Do not commit secrets. Prefer environment-injected values or a dedicated secrets mechanism.
- Keep template versions immutable once introduced.
- Separate runtime configuration from persistent player/world data.