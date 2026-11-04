---
title: Projektstatus und Technologien
description: Entwicklungsstand, geplante Funktionen und vorgesehene Technologien.
---

## Entwicklungsstand

StupidCloud befindet sich in einer frühen Entwicklungsphase. Repository und Planung entwickeln sich weiter. Die folgenden Funktionen beschreiben die Zielrichtung und sind keine Zusicherung, dass sie bereits verfügbar sind.

## Geplante Funktionen

- Mehrere entfernte Hosts über einen zentralen Controller verwalten.
- Einzelne Minecraft-Serverinstanzen in Docker-Containern ausführen.
- Templates, Konfiguration, Scheduling und Ports zentral verwalten.
- Proxy-Netzwerke über Velocity- und BungeeCord-Integrationen anbinden.
- SSH-basierte CLI- und TUI-Werkzeuge sowie eine separate Weboberfläche bereitstellen.

## Vorgesehene Technologien

| Bereich | Geplante Technologie |
| --- | --- |
| Backend | Go |
| Container | Docker Engine / Moby |
| Interne Kommunikation | gRPC und Protocol Buffers |
| Web-API | REST |
| Live-Updates | WebSocket |
| Datenbank | PostgreSQL |
| Konfiguration | YAML |
| Terminalverwaltung | Charmbracelet SSH und Bubble Tea |
| HTTP-Routing | Chi |
| Webfrontend | React / Next.js |

Technologieentscheidungen können sich während der Entwicklung ändern. Verbindliche Architekturentscheidungen sollten im Verzeichnis `docs/decisions/` des Repositorys dokumentiert werden.
