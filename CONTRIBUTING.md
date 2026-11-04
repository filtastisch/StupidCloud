# Contributing

Dieses Dokument beschreibt die erwarteten Arbeitsweisen bei Änderungen
an StupidCloud.

## Vor einer Änderung

-   Prüfen, ob das Thema bereits in `TODO.md` oder `ROADMAP.md` steht.
-   Bestehende Architekturgrenzen beachten.
-   Bei größeren Architekturänderungen ein ADR anlegen.
-   Änderungen möglichst klein und klar abgrenzen.

## Codequalität

Vor dem Einreichen:

``` powershell
gofmt -w ./src
go test ./...
go vet ./...
```

Neue fachliche Logik soll Tests erhalten. Fehler dürfen nicht
kommentarlos ignoriert werden.

## Architektur

-   UI- und Transportcode ruft Application Services auf.
-   Fachlogik soll nicht direkt von Docker, PostgreSQL oder HTTP
    abhängen.
-   Infrastrukturzugriffe werden hinter passenden Grenzen gekapselt.
-   Interfaces nur dort einführen, wo sie einen konkreten Entkopplungs-
    oder Testnutzen haben.
-   Keine zyklischen Package-Abhängigkeiten.

## Dokumentation

-   Öffentliche Konfigurationsänderungen in `docs/CONFIGURATION.md`
    aktualisieren.
-   Änderungen an Kommunikationsverträgen in `docs/PROTOCOLS.md`
    dokumentieren.
-   Neue Architekturentscheidungen als ADR festhalten.
-   Erledigte Aufgaben in `TODO.md` aktualisieren.
-   Roadmap nur ändern, wenn sich Prioritäten oder Phasenziele ändern.
