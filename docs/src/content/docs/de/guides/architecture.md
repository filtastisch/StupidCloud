---
title: Architektur
description: Der geplante Controller, Node-Agent und die Verwaltungsoberflächen.
---

StupidCloud basiert auf einem zentralen Controller und mehreren Node-Agents. Der Controller hält die gewünschte Konfiguration vor und verteilt Aufgaben. Die Nodes setzen diese auf ihren Hosts um und melden ihren aktuellen Zustand zurück.

## Controller

Der Controller ist die Koordinationsstelle. Zu seinen geplanten Aufgaben gehören Scheduling, Benutzer- und Templateverwaltung, API sowie Verwaltungsoberflächen. Er soll die maßgebliche Quelle für den gewünschten Zustand des Clusters sein.

## Node-Agent

Auf jedem verwalteten Host läuft ein Node-Agent. Er kommuniziert mit dem Controller, erstellt und verwaltet Servercontainer über Docker und meldet Ergebnisse sowie Zustände an den Controller zurück.

## Verwaltungsoberflächen

Geplant sind SSH-basierte CLI- und TUI-Zugänge über den Controller sowie eine separate Weboberfläche. Alle Oberflächen sollen dieselben Controller-APIs und Abläufe verwenden.

## Geplante Kommunikation und Speicherung

Der aktuelle Entwurf sieht gRPC und Protocol Buffers für die interne Kommunikation zwischen Controller und Nodes vor, eine REST-API für Webclients, WebSockets für Live-Updates und PostgreSQL für dauerhafte Daten. Diese Technologien sind geplant und können sich im Entwicklungsverlauf ändern.

## Weiterführende Seite

- [Projektstatus und geplante Technologien](/de/reference/project-status/)
