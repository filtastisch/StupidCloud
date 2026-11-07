---
title: Erste Schritte
description: Den aktuellen Controller-Grundstand von StupidCloud starten und die Repository-Struktur verstehen.
---

Dieser Leitfaden beschreibt den **aktuellen** Stand des Projekts. StupidCloud befindet sich in einer frühen Phase; heute ist nur grundlegendes Controller-Verhalten umgesetzt.

## Voraussetzungen

- **Go** in einer zu `go.mod` passenden Version (das Modul deklariert derzeit `go 1.27`).
- **Git**.
- Für spätere Arbeiten (nicht für den aktuellen Grundstand nötig): **Docker Engine** für Container-Integration und **PostgreSQL** für geplante Persistenz.

## Controller klonen und starten

Aus dem Repository-Root:

```bash
go run ./src/cmd/controller
```

Aktuelles Startverhalten:

1. Es wird eine Single-Instance-Lockdatei unter `<System-Temp>/stupidcloud-controller.lock` erstellt. Läuft bereits ein Controller, beendet sich der Prozess.
2. Der Standard-`slog`-Textlogger wird initialisiert.
3. Basisinformationen zu Betriebssystem/Architektur werden geloggt.
4. Die Konfiguration wird aus `./data/config.toml` geladen. Existieren Datei oder Verzeichnis nicht, erstellt `config.Load` `./data/` und schreibt die eingebettete Standardkonfiguration, anschließend wird sie geparst.
5. Der Controller startet und wartet auf `SIGINT` oder `SIGTERM`, danach wird der Shutdown geloggt.

> Der Controller bietet derzeit Lifecycle-/Grundgerüst-Verhalten. gRPC/SSH/Webserver, Node-Orchestrierung, Docker-Operationen und datenbankgestützte Funktionen sind **noch nicht vollständig verdrahtet**.

## Nützliche Befehle

```bash
# Controller starten
go run ./src/cmd/controller

# Alle Tests ausführen
go test ./...

# Statische Prüfungen
go vet ./...

# Code formatieren
gofmt -w ./src ./config

# Abhängigkeiten aufräumen
go mod tidy

# Controller bauen
go build ./src/cmd/controller
```

## Repository-Struktur

```text
StupidCloud/
├── config/              # Konfigurationsladen und eingebettete Defaults (TOML)
├── data/                # Lokale Laufzeitdaten (z. B. config.toml); bei Bedarf erstellt
├── docs/                # Astro-Starlight-Dokumentationsseite
├── src/
│   ├── cmd/
│   │   ├── controller/  # Controller-Einstiegspunkt
│   │   └── node/        # Node-Einstiegspunkt (derzeit Platzhalter)
│   └── internal/
│       ├── app/         # Controller-Lifecycle/Grundgerüst
│       ├── node/        # Node-Manager (Grundgerüst)
│       ├── proxy/       # Proxy-Manager (Grundgerüst)
│       ├── service/     # Service-Manager (Grundgerüst)
│       └── web/         # Webserver-Paket (Grundgerüst)
├── go.mod
├── README.md
├── TODO.md
└── ROADMAP.md
```

Das Verzeichnis `src/` ist ein normales Verzeichnis (keine besondere Go-Konvention). Imports nutzen den Modulpfad, z. B. `StupidCloud/src/internal/app`.

## Node-Agent

`src/cmd/node` ist derzeit nur ein Platzhalter (`package node`). Node-Registrierung, gRPC-Protokoll und Docker-Ausführung sind geplant; siehe [Roadmap](/de/reference/roadmap/).

## Nächste Schritte

- [Architektur](/de/guides/architecture/) — Zielentwurf Controller/Node.
- [Konfiguration](/de/guides/configuration/) — aktuelle TOML-Konfiguration und geplantes Modell.
- [Entwicklung](/de/guides/development/) — Konventionen und Arbeitsablauf.