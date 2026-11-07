---
title: Roadmap
description: Geplante Entwicklungsphasen und Ergebnisse (ohne feste Veröffentlichungstermine).
---

Diese Seite fasst die Projekt-Roadmap zusammen. Die maßgebliche, fortlaufend gepflegte Version steht in `ROADMAP.md` im Repository-Root. Phasen gelten erst als abgeschlossen, wenn Akzeptanzkriterien erfüllt und getestet sind.

## Phase 0 — Projektfundament

**Ziel:** Ein sauber startendes und testbares Go-Projekt.
- Repository-/Paketstruktur, Controller-Einstieg, Logging/Konfig-Grundlagen.
- Start-/Shutdown-Lifecycle, erste Tests und CI.
- **Akzeptanz:** Controller startet lokal; SIGINT/SIGTERM führen zu kontrolliertem Shutdown; Fehler mit Kontext geloggt; `go test ./...` läuft.

## Phase 1 — Controller-Kern

**Ziel:** Zentrale Anwendungsdienste und Persistenz.
- PostgreSQL-Anbindung/Migrationen; Benutzer-/Berechtigungsmodell.
- Node-/Service-/Template-/Instanzmodelle; Event-/Job-Grundlage; Konfig-Validierung.
- **Akzeptanz:** DB initialisiert/migriert; Kernobjekte über Application Services verwaltbar; Fachlogik unabhängig von HTTP/SSH/TUI testbar.

## Phase 2 — Ein einzelner Node

**Ziel:** Zuverlässige Controller–Node-Verbindung und Docker-Steuerung.
- Eigenständiger Node-Agent, gRPC-Protokoll, sichere Registrierung/Heartbeats.
- Docker-Integration, lokaler Cache, Statusmeldungen.
- **Akzeptanz:** Node registriert sich; Controller startet/stoppt Testcontainer; Verbindungsabbrüche werden behandelt; Container werden bei temporärem Controller-Ausfall nicht unerwartet beendet.

## Phase 3 — Minecraft-Instanzen und Templates

**Ziel:** Minecraft-Server aus versionierten Templates erzeugen.
- Template-Manager, Runtime-/Minecraft-Konfiguration, Instanz-Lifecycle.
- Persistente Daten, Healthchecks, Logzugriff, Portreservierung.
- **Akzeptanz:** Erstellen/Starten/Überwachen/Entfernen aus Template; versionierte Templates; persistente Daten überleben Neustarts; Portkonflikte korrekt behandelt.

## Phase 4 — Multi-Node und Scheduling

**Ziel:** Instanzen über Hosts verteilen.
- Node-Ressourcen/Kapazitäten, Scheduling-Regeln, Labels/Einschränkungen.
- Zentrale Portbereiche, privates Backend-Netzwerk.
- **Akzeptanz:** passende Node-Auswahl; Ressourcenlimits beachtet; Ports pro Node eindeutig reserviert; nicht erreichbare Nodes korrekt behandelt.

## Phase 5 — Proxy-Integration

**Ziel:** Instanzen automatisch in Proxy-Netzwerke integrieren.
- Velocity/BungeeCord-Bridges, Proxy-Gruppen.
- Registrierung nach Healthcheck, Reconciliation, Draining.
- **Akzeptanz:** bereite Lobby registriert; gestoppte Instanz sauber entfernt; Proxy-Neustarts werden abgeglichen.

## Phase 6 — Verwaltung

**Ziel:** Die Plattform administrierbar machen.
- CLI, SSH-Server mit nur Public-Key-Auth, Multi-User-TUI.
- Rollen/Berechtigungen, Audit-Events.
- **Akzeptanz:** parallele SSH-Sessions; Autorisierung bei jeder Aktion; keine SSH-Passwortauth.

## Phase 7 — Weboberfläche und Betrieb

**Ziel:** Externe Verwaltung und produktionsnahe Abläufe.
- REST-API + WebSocket, React/Next.js-UI.
- TLS/Reverse-Proxy-Konzept, Monitoring/Metriken.
- Installations-/Backup-/Upgrade-Prozess.
- **Akzeptanz:** UI verwaltet Nodes/Instanzen; Live-Status; Backup/Restore getestet; Install-/Upgrade-Doku.

## Spätere Erweiterungen

Nach stabilem Grundsystem evaluieren: Controller-HA, Wartungsfenster, Instanzmigration, erweiterte Ressourcen-/Kostenplanung, zusätzliche Runtimes, Plugin-/Extension-System.