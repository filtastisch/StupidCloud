---
title: Projektstatus
description: Aktueller Reifegrad, vorhandene Grundlagen und geplante Technologieentscheidungen.
---

## Reifegrad

StupidCloud befindet sich in einer **frühen Entwicklungsphase**. Repository und Entwurf entwickeln sich weiter. Als geplant oder als Ziel beschriebenen Funktionen sollte man nicht vertrauen, solange sie nicht im Code und in Tests bestätigt sind.

## Was heute vorhanden ist (Code-Stand)

- **Controller-Einstiegspunkt** (`src/cmd/controller/main.go`): Single-Instance-Lock, `slog`-Logging, Signalbehandlung (`SIGINT`/`SIGTERM`), Laden der Konfiguration aus `./data/config.toml`.
- **Controller-Grundgerüst** (`src/internal/app/controller.go`): Initialisierung der Manager (`service`, `node`, `proxy`), Platzhalter für SSH/Web/gRPC-Server, einfaches Starten/Warten auf Shutdown.
- **Konfigurationsladen** (`config/`): TOML-Parsing mit eingebetteten Defaults.
- **Node-Einstiegspunkt**: nur Platzhalter.

Viele Teilbereiche (Datenbankpersistenz, Node-gRPC-Protokoll, Docker-Orchestrierung, Scheduling, Proxy-Bridges, SSH/TUI/Web-Verwaltung) sind **geplant** und in `TODO.md` und `ROADMAP.md` festgehalten.

## Geplante Fähigkeiten (Ziel)

- Mehrere entfernte Hosts über einen zentralen Controller verwalten.
- Einzelne Minecraft-Serverinstanzen in Docker-Containern ausführen.
- Templates, Konfiguration, Scheduling und Ports zentral verwalten.
- Proxy-Netzwerke über Velocity und BungeeCord integrieren.
- SSH-basierte CLI/TUI und eine separate Weboberfläche bereitstellen.

## Geplante/genutzte Technologien

| Bereich | Geplant oder im Einsatz | Hinweise |
|---|---|---|
| Backend | Go | Modul `StupidCloud`. |
| Container | Docker Engine / Moby | Integration geplant. |
| Interne Kommunikation | gRPC + Protocol Buffers | Controller–Node-Protokoll geplant. |
| Web-API | REST | Geplant. |
| Live-Updates | WebSocket | Geplant. |
| Datenbank | PostgreSQL | Geplant (`pgx` im TODO genannt). |
| Konfiguration | TOML (aktueller Loader); Modell in Entwicklung | Aktuelle Controller-Config ist TOML via `config.Load`. |
| Terminalverwaltung | Charmbracelet SSH + Bubble Tea | Geplant/Abhängigkeiten vorhanden. |
| HTTP-Routing | Chi | Abhängigkeit vorhanden; vollständige Verdrahtung geplant. |
| Webfrontend | React / Next.js | Geplant; `docs/` nutzt Astro/Starlight. |

Technologieentscheidungen können sich ändern. Verbindliche Entscheidungen bei Formalisierung als ADRs unter `docs/decisions/` festhalten.