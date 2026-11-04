---
title: StupidCloud
description: Dokumentation der geplanten Go-Plattform zur Orchestrierung von Minecraft-Servern.
template: splash
hero:
  tagline: Verteilte Minecraft-Server-Infrastrukturen planen, betreiben und verwalten.
  actions:
    - text: Architektur kennenlernen
      link: /de/guides/architecture/
      icon: right-arrow
    - text: Projektstatus
      link: /de/reference/project-status/
      icon: open-book
---

StupidCloud ist ein frühes Projekt zur zentralen Verwaltung von Minecraft-Servern auf mehreren Hosts. Die Server sollen in Docker-Containern laufen. Ein Controller koordiniert Konfiguration und Aufgaben der Node-Agents.

> **Projektstatus:** StupidCloud befindet sich in einer frühen Entwicklungsphase. Diese Dokumentation beschreibt das geplante System; noch sind nicht alle Komponenten implementiert.

## Ziele

- Mehrere Hosts (Nodes) zentral verwalten.
- Minecraft-Serverinstanzen in Docker-Containern ausführen.
- Templates und Konfigurationen wiederverwendbar verwalten.
- Scheduling und Portverwaltung zentral organisieren.
- Velocity und BungeeCord integrieren.
- Verwaltung über SSH, CLI, TUI und Weboberfläche ermöglichen.

## Einstieg

- Der [Architekturleitfaden](/de/guides/architecture/) erklärt die geplanten Komponenten und ihre Aufgaben.
- Der [Projektstatus und die Technologien](/de/reference/project-status/) zeigen, was geplant und was noch im Aufbau ist.
