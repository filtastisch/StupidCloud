# StupidCloud

**StupidCloud** ist eine geplante, in Go entwickelte
Server-Orchestrierungsplattform. Das Projekt soll die Verwaltung
verteilter Minecraft-Server-Infrastrukturen ermöglichen. Einzelne
Minecraft-Server werden dabei in separaten Docker-Containern betrieben
und zentral über einen Controller verwaltet.

> **Projektstatus:** Frühe Entwicklungsphase. Die Dokumentation
> beschreibt das angestrebte Zielsystem. Nicht alle beschriebenen
> Komponenten sind bereits implementiert.

## Ziele

-   Zentrale Verwaltung mehrerer Hosts (Nodes).
-   Ausführung einzelner Serverinstanzen in Docker-Containern.
-   Einheitliche Verwaltung von Templates und Konfigurationen.
-   Automatisches Scheduling und zentrale Portverwaltung.
-   Integration von Velocity und BungeeCord.
-   Verwaltung über SSH, CLI, TUI und Webinterface.
-   Sichere Kommunikation zwischen Controller und Nodes.
-   Erweiterbarkeit und klare Trennung der Verantwortlichkeiten.

## Architektur in Kürze

StupidCloud besteht konzeptionell aus drei Bereichen:

1.  **Controller:** Zentrale Steuerung, Scheduling, Benutzerverwaltung,
    Templates, API und Management-Oberflächen.
2.  **Node-Agent:** Läuft auf jedem verwalteten Host, kommuniziert mit
    dem Controller und setzt Container-Aufträge lokal über Docker um.
3.  **Management-Oberflächen:** SSH-basierte CLI/TUI auf dem Controller
    sowie ein separates Webinterface.

Der Controller ist die zentrale Quelle der gewünschten Konfiguration.
Nodes führen die ihnen zugewiesenen Aufgaben aus und melden ihren
Zustand zurück.

## Geplante Technologien

  Bereich                 Technologie
  ----------------------- ---------------------------------
  Backend                 Go
  Container               Docker Engine / Moby
  Interne Kommunikation   gRPC und Protocol Buffers
  Web API                 REST
  Live-Updates            WebSocket
  Datenbank               PostgreSQL
  Konfiguration           YAML
  SSH-Management          Charmbracelet SSH
  TUI                     Bubble Tea
  HTTP-Routing            Chi
  Proxy-Anbindung         Velocity- und BungeeCord-Bridge
  Webfrontend             React / Next.js

Technologien können sich während der Entwicklung ändern. Verbindliche
Entscheidungen werden in `docs/decisions/` dokumentiert.

## Repository-Struktur

Die aktuelle Projektstruktur verwendet ein `src`-Verzeichnis:

``` text
StupidCloud/
├── src/
│   ├── cmd/
│   │   ├── controller/
│   │   └── node/
│   ├── internal/
│   │   ├── app/
│   │   ├── auth/
│   │   ├── config/
│   │   ├── database/
│   │   ├── docker/
│   │   ├── event/
│   │   ├── instance/
│   │   ├── job/
│   │   ├── logging/
│   │   ├── management/
│   │   ├── node/
│   │   ├── port/
│   │   ├── proxy/
│   │   ├── scheduling/
│   │   ├── service/
│   │   ├── storage/
│   │   ├── template/
│   │   └── web/
│   └── pkg/
│       └── protocol/
├── web/
├── deployments/
├── templates/
├── docs/
├── go.mod
└── go.sum
```

Die Struktur ist ein langfristiges Zielbild. Verzeichnisse sollen erst
angelegt werden, wenn sie tatsächlich benötigt werden. Go behandelt
`src` als normales Unterverzeichnis; es ist keine besondere
Go-Konvention.

## Entwicklungsstand

Der aktuelle Schwerpunkt liegt auf den Grundlagen des Go-Projekts und
dem Lifecycle des Controllers. Die vollständige Zielarchitektur ist noch
nicht umgesetzt.

Siehe: - [TODO](TODO.md) - [Roadmap](ROADMAP.md) -
[Architektur](docs/ARCHITECTURE.md) -
[Entwicklungsleitfaden](docs/DEVELOPMENT.md)

## Lokale Entwicklung

Voraussetzungen:

-   Eine aktuelle Go-Version passend zur `go.mod`.
-   Git.
-   Docker Engine für spätere Docker-Integrationstests.
-   PostgreSQL für spätere Datenbankfunktionen.

Controller starten (aus dem Repository-Root):

``` powershell
go run ./src/cmd/controller
```

Node starten (sobald der Einstiegspunkt implementiert ist):

``` powershell
go run ./src/cmd/node
```

Abhängigkeiten aktualisieren und aufräumen:

``` powershell
go mod tidy
```

Tests ausführen:

``` powershell
go test ./...
```

## Sicherheitsgrundsätze

-   Node-Agenten bauen ausgehende, authentifizierte Verbindungen zum
    Controller auf.
-   Interne Kommunikation soll TLS-geschützt und gegenseitig
    authentifiziert werden.
-   SSH-Management verwendet ausschließlich
    Public-Key-Authentifizierung.
-   Web-Authentifizierung und SSH-Schlüssel sind getrennte
    Authentifizierungswege.
-   Backend-Minecraft-Ports werden nicht unnötig öffentlich freigegeben.
-   Secrets gehören nicht in Templates, Git oder Logs.

Details: [Security](docs/SECURITY.md).

## Lizenz

Noch nicht festgelegt. Bis zur ausdrücklichen Festlegung gilt: Es wird
keine Open-Source-Lizenz angenommen. Nutzung und Weiterverteilung
richten sich nach den jeweils geltenden Rechten des Urhebers.
