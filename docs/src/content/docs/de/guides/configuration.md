---
title: Konfiguration
description: Aktuelles Verhalten der Controller-Konfiguration und geplantes Konfigurationsmodell.
---

## Aktuelles Verhalten (implementiert)

Der Controller-Einstiegspunkt lädt beim Start die Konfiguration:

```go
cfg, err := config.Load("./data/config.toml")
```

Implementierungsdetails (`config/config.go`):

- Wenn `./data/` (bzw. das Konfigurationsverzeichnis) nicht existiert, wird es erstellt.
- Wenn `./data/config.toml` nicht existiert, wird die eingebettete Standarddatei `defaults/config.toml` an diesen Pfad geschrieben.
- Die Datei wird als **TOML** mit `github.com/pelletier/go-toml/v2` geparst.

### Aktuelle Konfigurationsstruktur

Das aktuelle `Config`-Modell enthält diese Top-Level-Abschnitte:

| Abschnitt | Zweck (aktuelles Modell) |
|---|---|
| `application` | Name, Version, Umgebung, Debug-Flag |
| `controller` | Host/Port des Controllers |
| `database` | Datenbanktyp/-pfad (Modellfeld; vollständige DB-Integration geplant) |
| `logging` | Level/Format |
| `network` | Bind-/Advertise-Adressen |
| `nodes` | Statische Node-Einträge (Name, Adresse, Port, Enabled) |

> **Hinweis:** Nicht alle Felder werden bereits von laufenden Komponenten aktiv genutzt, da große Teile des Systems noch als Grundgerüst vorliegen. Dies beschreibt die aktuelle Config-Form, nicht garantiertes vollständiges Laufzeitverhalten.

## Geplantes Konfigurationsmodell (Ziel)

Das Zielsystem trennt Zuständigkeiten deutlicher (kann sich ändern):

- **Controller-Konfiguration:** Listener (gRPC/HTTP/SSH), Datenbank-DSN, Storage, Auth/Cluster.
- **Node-Konfiguration:** Node-Identität, Controller-Endpunkt, lokaler Docker-Socket, Cache.
- **Templates:** Runtime/Minecraft-Defaults, Dateien, Healthchecks, Proxy/Netzwerk-Exponierung.
- **Service-/Instanz-Overrides:** konkrete Abweichungen pro Service/Instanz.
- **Secrets:** separat verwaltet; niemals in Git oder öffentliche Templates.

Ältere Planungsdokumente zeigten YAML-Beispiele; die aktuelle Implementierung lädt Controller-Konfiguration als TOML. Endgültige Formate und Schemata werden mit der Umsetzung festgelegt.

## Praktische Hinweise

- Konfiguration beim Start validieren und verständliche Fehler liefern (teilweise geplant).
- Keine Secrets committen. Werte bevorzugt per Umgebung oder dediziertem Secrets-Mechanismus bereitstellen.
- Template-Versionen nach Einführung unveränderlich behandeln.
- Laufzeitkonfiguration von persistenten Spieler-/Weltdaten trennen.