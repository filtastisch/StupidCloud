---
title: Architektur
description: Der geplante Controller/Node-Entwurf und was heute im Code vorhanden ist.
---

StupidCloud ist um einen **zentralen Controller** und **Node-Agents** auf verwalteten Hosts herum aufgebaut. Der Controller ist die maßgebliche Quelle für den *Sollzustand*; Nodes setzen Arbeit lokal um (im Zielentwurf primär über Docker) und melden den *Istzustand*.

> **Umsetzungsstand:** Das Repository enthält Controller-Lifecycle-Grundgerüste und Manager-Platzhalter. Vollständige Controller–Node-Kommunikation, Docker-Orchestrierung, Persistenz, Proxy-Integration und Verwaltungsoberflächen sind noch nicht fertig. Die folgenden Abschnitte beschreiben, sofern nicht anders markiert, die **Zielarchitektur**.

## Komponenten im Überblick

### Controller

Geplante Aufgaben:

- Arbeit für Nodes/Services/Instanzen annehmen und koordinieren.
- Instanzen auf passende Nodes schedulen (Ressourcen, Labels, Kapazität).
- Templates, Konfiguration, Benutzer und Berechtigungen verwalten.
- API-Oberfläche bereitstellen (intern gRPC, Web REST/WebSocket, SSH/TUI/CLI-Integrationen).
- Autoritativen Zustand persistieren (geplant: PostgreSQL).

**Aktueller Code:** `src/cmd/controller/main.go` startet den Prozess, lädt die Konfiguration und führt `app.Controller.Start(ctx)` aus. `src/internal/app/controller.go` initialisiert Manager (`service`, `node`, `proxy`) und hält Felder für SSH/Web/gRPC-Server; `Start` wartet derzeit auf den Shutdown.

### Node-Agent

Geplante Aufgaben:

- Auf jedem verwalteten Host mit Zugriff auf die lokale Docker-Engine laufen.
- Eine **ausgehende, authentifizierte** Verbindung zum Controller aufbauen (geplant: gRPC/Protobuf).
- Containeroperationen ausführen (Erstellen/Starten/Stoppen/Entfernen, Logs, Health/Status).
- Heartbeats, Fähigkeiten/Ressourcen und Instanzstatus melden.
- Artefakte/Templates lokal cachen, wo sinnvoll.

**Aktueller Code:** `src/cmd/node/main.go` ist ein Platzhalter. Das Node-Verhalten ist ein Roadmap-Punkt.

### Verwaltungsoberflächen

Geplant:

- **SSH**-Verwaltung (Public-Key-orientiert) für CLI/TUI.
- **TUI** (Bubble Tea) und **CLI** über gemeinsame Application Services.
- **Web-UI** (geplant React/Next.js) über REST/WebSocket-APIs des Controllers.

Alle Oberflächen sollen dieselbe Autorisierung und dieselbe Anwendungslogik nutzen.

## Kommunikation und Daten (geplant)

| Grenze | Geplante Technologie | Hinweise |
|---|---|---|
| Controller ↔ Node | gRPC + Protocol Buffers | Authentifiziert/verschlüsselt; idempotente Operationen; Reconnect/Backoff. |
| Webclient ↔ Controller | REST-API | Versioniert (z. B. `/api/v1`). |
| Live-Updates | WebSocket | Authentifiziert; nach Berechtigungen gefiltert. |
| Persistenz | PostgreSQL | Migrationen, Constraints (z. B. Portreservierungen). |
| Konfiguration | Heute TOML / Modell in Entwicklung | Aktueller Loader nutzt TOML; ein breiteres Modell ist geplant. |

## Minecraft-/Proxy-Integration (geplant)

StupidCloud zielt auf Minecraft-Workloads mit **Velocity**- und **BungeeCord**-Integration. Eine Bridge-Komponente soll den gewünschten Proxy-Zustand vom Controller erhalten und über Proxy-APIs abgleichen, Backends nach Healthchecks registrieren und beim Shutdown drainen.

## Sicherheitsgrenzen

- Nodes sollten Verbindungen zum Controller initiieren (reduziert eingehende Exponierung auf Nodes).
- Interner Verkehr sollte TLS-geschützt sein, wo passend mit gegenseitiger Authentifizierung.
- Zugriff auf den Docker-Socket ist hochprivilegiert und sollte auf den Node-Agent beschränkt sein.
- Backend-Minecraft-Ports sollten nicht öffentlich exponiert werden, außer dies ist ausdrücklich gewollt.

Siehe [Sicherheit](/de/guides/security/) für Details und Hinweise.

## Weiterführende Seiten

- [Projektstatus](/de/reference/project-status/)
- [Roadmap](/de/reference/roadmap/)
- [Glossar](/de/reference/glossary/)