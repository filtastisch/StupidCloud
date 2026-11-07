---
title: Entwicklung
description: Arbeitsablauf, Konventionen und praktische Hinweise für Beiträge zu StupidCloud.
---

## Projektkonventionen

- **Modul/Imports:** Das Go-Modul heißt `StupidCloud`. Der Code liegt unter `src/` und `config/`, daher sehen Imports z. B. so aus: `StupidCloud/src/internal/app` oder `StupidCloud/config`.
- **Binaries:** Einstiegspunkte liegen in `src/cmd/controller` und `src/cmd/node` (Node ist derzeit ein Platzhalter).
- **Interne Pakete:** Implementierungsdetails gehören unter `src/internal/...`.

## Lokaler Arbeitsablauf

Vor dem Commit prüfen:

```bash
gofmt -w ./src ./config
go test ./...
go vet ./...
```

Bauen/Starten:

```bash
go run ./src/cmd/controller
go build ./src/cmd/controller
```

## Go-Richtlinien

- Pakete klein und fokussiert halten.
- Fehler explizit zurückgeben; bei Bedarf mit `%w` Kontext ergänzen.
- `context.Context` als ersten Parameter für abbrechbare/timeout-fähige Operationen übergeben.
- Unnötige Globals vermeiden; Abhängigkeiten über Konstruktoren injizieren.
- Goroutinen nur mit klarem Lebenszyklus und Abbruchpfad starten.
- Interfaces dort definieren, wo sie gebraucht werden, nicht vorsorglich.

## Lifecycle und Shutdown

Dauerhaft laufende Komponenten sollten Context-Abbruch respektieren und klar definiertes Start-/Shutdown-Verhalten haben. Ein produktionsreifer Controller sollte:

- Komponenten in definierter Reihenfolge initialisieren,
- bei Startfehlern teilweise gestartete Komponenten zurückrollen,
- in umgekehrter Reihenfolge mit Timeout herunterfahren.

Der aktuelle Controller-Grundstand implementiert signalbasierten Abbruch und einen einfachen Start-/Shutdown-Pfad; die vollständige Komponenten-Orchestrierung folgt.

## Tests

- Unit-Tests für Domänen-/Anwendungslogik bevorzugen.
- Integrationstests für Infrastruktur (Docker, PostgreSQL), sobald diese Funktionen umgesetzt sind.
- Fehlerpfade, Abbruch, Wiederholungen/Idempotenz (wichtig für Orchestrierung) und Nebenläufigkeit testen, wo relevant.

## Abhängigkeiten

Abhängigkeiten nur hinzufügen, wenn sie für die aktuelle Arbeit benötigt werden. Nach Änderungen:

```bash
go mod tidy
```

`go.mod` und `go.sum` gemeinsam committen. Änderungen an `go`-Direktive/Toolchain bewusst vornehmen.

## Dokumentation

- Diese Seite liegt in `docs/` (Astro + Starlight) mit englischen (`root`) und deutschen (`de`) Inhalten.
- Dokumentation am **tatsächlichen** Umsetzungsstand ausrichten; geplante Funktionen klar kennzeichnen.
- Architekturentscheidungen mit langfristiger Wirkung sollten als ADRs festgehalten werden (die README verweist auf `docs/decisions/`; Struktur anlegen, wenn Entscheidungen formalisiert werden).