---
title: StupidCloud
description: Dokumentation für die geplante, in Go entwickelte Orchestrierungsplattform für Minecraft-Server.
template: splash
hero:
  tagline: Verteilte Minecraft-Server-Infrastruktur planen, betreiben und verwalten.
  actions:
    - text: Erste Schritte
      link: /de/guides/getting-started/
      icon: right-arrow
    - text: Architektur
      link: /de/guides/architecture/
      icon: open-book
    - text: Projektstatus
      link: /de/reference/project-status/
      variant: minimal
---

StupidCloud ist eine **frühe, in Go entwickelte Orchestrierungsplattform** zur Verwaltung verteilter **Minecraft**-Server-Infrastrukturen. Im Zielsystem laufen einzelne Minecraft-Server in separaten **Docker-Containern**, koordiniert durch einen zentralen **Controller** und ausgeführt von **Node-Agents** auf verwalteten Hosts.

> **Status:** Diese Dokumentation beschreibt das Zielsystem und den aktuellen Stand des Codes. Viele Komponenten sind geplant, aber noch nicht umgesetzt. Prüfe [Projektstatus](/de/reference/project-status/) und [Roadmap](/de/reference/roadmap/), bevor du dich auf eine Funktion verlässt.

## Ziele von StupidCloud

- Zentrale Verwaltung mehrerer Hosts (**Nodes**).
- Minecraft-Serverinstanzen in Docker-Containern.
- Wiederverwendbare Templates und konsistente Konfiguration.
- Scheduling und zentrale Portverwaltung/-reservierung.
- Integration von **Velocity**- und **BungeeCord**-Proxies.
- Verwaltung über **SSH** (Public-Key-orientiertes Design), CLI, TUI und Weboberfläche.
- Sichere, authentifizierte Kommunikation zwischen Controller und Nodes.

## Einstieg

- **Aktuellen Controller-Grundstand starten:** [Erste Schritte](/de/guides/getting-started/)
- **Zielarchitektur verstehen:** [Architektur](/de/guides/architecture/)
- **Sicher beitragen:** [Entwicklung](/de/guides/development/) und [Sicherheit](/de/guides/security/)
- **Fortschritt verfolgen:** [Projektstatus](/de/reference/project-status/) und [Roadmap](/de/reference/roadmap/)