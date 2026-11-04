# TODO

Diese Liste enthält konkrete Arbeitspakete. Der Status bezieht sich auf
den bekannten Projektstand und muss bei Änderungen aktualisiert werden.

Legende:

-   `[ ]` Offen
-   `[x]` Erledigt
-   `[ ]` **BLOCKED** -- abhängig von einer noch offenen Entscheidung
    oder Voraussetzung

## Aktueller Fokus: Controller-Grundlage

-   [x] Projektstruktur und Go-Modulpfad vereinheitlichen.
-   [x] `cmd/controller/main.go` als Programmeinstieg einrichten.
-   [x] `internal/app/controller.go` implementieren.
-   [x] Logger über `log/slog` initialisieren und injizieren.
-   [x] `Controller.Start(context.Context)` implementieren.
-   [x] Graceful Shutdown bei SIGINT und SIGTERM umsetzen.
-   [ ] Doppelte Starts verhindern und einen passenden Fehler
    zurückgeben.
-   [ ] Erste Tests für Controller-Lifecycle schreiben.
-   [ ] `go test ./...` und `go vet ./...` erfolgreich ausführen.

## Projekt- und Konfigurationsgrundlage

-   [ ] Konfigurationsmodell für Controller und Node entwerfen.
-   [ ] YAML-Konfiguration laden.
-   [ ] Standardwerte definieren.
-   [ ] Konfiguration validieren und verständliche Fehlermeldungen
    liefern.
-   [ ] Pfade für Daten, Logs, Cache und Konfiguration festlegen.
-   [ ] Erststart-/Setup-Wizard konzipieren.
-   [ ] Version und Build-Informationen bereitstellen.

## Controller-Lifecycle und Komponenten

-   [ ] Komponenten mit klarer Startreihenfolge initialisieren.
-   [ ] Fehler bei der Initialisierung sauber zurückrollen.
-   [ ] Komponenten in umgekehrter Reihenfolge herunterfahren.
-   [ ] Shutdown-Timeout definieren.
-   [ ] Gesundheitsstatus des Controllers bereitstellen.
-   [ ] Signal- und Fehlerbehandlung zentralisieren.

## Datenbank

-   [ ] PostgreSQL-Anbindung mit `pgx` einrichten.
-   [ ] Verbindungs-Pool konfigurieren.
-   [ ] Migrationssystem auswählen und integrieren.
-   [ ] Tabellen für Benutzer und SSH-Schlüssel entwerfen.
-   [ ] Tabellen für Nodes und Node-Registrierungen entwerfen.
-   [ ] Tabellen für Services, Templates und Instanzen entwerfen.
-   [ ] Portreservierungen mit Datenbank-Constraints absichern.
-   [ ] Transaktionsgrenzen dokumentieren.
-   [ ] Backup- und Wiederherstellungsstrategie definieren.

## Node-Agent und Kommunikation

-   [ ] Eigenständigen Node-Einstiegspunkt erstellen.
-   [ ] Node-Konfiguration und lokale Identität implementieren.
-   [ ] gRPC-Protokoll und Protobuf-Nachrichten definieren.
-   [ ] Sichere Node-Registrierung umsetzen.
-   [ ] mTLS oder gleichwertige gegenseitige Authentifizierung
    implementieren.
-   [ ] Heartbeats und Statusmeldungen einführen.
-   [ ] Reconnect- und Backoff-Strategie implementieren.
-   [ ] Verhalten bei Controller-Ausfall definieren.
-   [ ] Node-Fähigkeiten und Ressourcen erfassen.
-   [ ] Systemdienst-Installationen (systemd) bereitstellen.

## Docker und Instanzen

-   [ ] Docker-Client kapseln.
-   [ ] Docker-Verfügbarkeit und API-Kompatibilität prüfen.
-   [ ] Container erstellen, starten, stoppen und entfernen.
-   [ ] Container-Logs abrufen.
-   [ ] Containerzustand überwachen.
-   [ ] Ressourcenlimits abbilden.
-   [ ] Volumes und persistente Daten unterstützen.
-   [ ] Container-Labels für StupidCloud-Metadaten definieren.
-   [ ] Fehler bei Docker-Operationen klassifizieren.
-   [ ] Idempotente Operationen und Wiederholungen implementieren.

## Templates und Storage

-   [ ] Template-Datenmodell definieren.
-   [ ] Versionierte, unveränderliche Template-Versionen unterstützen.
-   [ ] Template-Dateien und Artefakte zentral speichern.
-   [ ] Hash-Prüfung für heruntergeladene Artefakte implementieren.
-   [ ] Node-seitigen Cache einführen.
-   [ ] Konfigurationsvererbung definieren: global, Template, Service,
    Instanz.
-   [ ] Secrets getrennt von Template-Dateien verwalten.
-   [ ] Persistente Welt- und Spielerdaten von Template-Dateien trennen.
-   [ ] Rollout- und Aktualisierungsstrategie definieren.

## Scheduling und Ports

-   [ ] Node-Auswahl anhand von Ressourcen und Labels implementieren.
-   [ ] CPU-, RAM- und Kapazitätsinformationen erfassen.
-   [ ] Portbereiche pro Node konfigurieren.
-   [ ] Ports transaktionssicher reservieren.
-   [ ] Belegungsprüfung beim Containerstart durchführen.
-   [ ] Freigabe und Wiederverwendung von Ports implementieren.
-   [ ] Expositionstypen (öffentlich, intern, container-only)
    unterstützen.
-   [ ] Fehler- und Wiederholungsstrategie für Portkonflikte definieren.

## Proxy-Integration

-   [ ] Gemeinsames Bridge-Protokoll definieren.
-   [ ] Velocity-Bridge entwickeln.
-   [ ] BungeeCord-Bridge entwickeln.
-   [ ] Server erst nach erfolgreichem Healthcheck registrieren.
-   [ ] Registrierung idempotent machen.
-   [ ] Proxy-Reconnect und Reconciliation implementieren.
-   [ ] Graceful Draining vor dem Stoppen einer Instanz umsetzen.
-   [ ] Backend-Verbindungen über privates Netzwerk absichern.

## Management und API

-   [ ] Gemeinsame Application Services für CLI, TUI und Web definieren.
-   [ ] CLI-Kommandostruktur entwerfen.
-   [ ] SSH-Server mit Public-Key-Authentifizierung integrieren.
-   [ ] SSH-Session-Isolation für mehrere gleichzeitige Nutzer
    sicherstellen.
-   [ ] Bubble Tea TUI implementieren.
-   [ ] REST API versionieren.
-   [ ] WebSocket-Live-Events implementieren.
-   [ ] RBAC und Berechtigungsprüfung zentralisieren.
-   [ ] Webfrontend entwickeln.
-   [ ] Audit-Logging für administrative Aktionen einführen.

## Qualität und Betrieb

-   [ ] Unit-Tests für fachliche Logik schreiben.
-   [ ] Integrationstests für PostgreSQL und Docker ergänzen.
-   [ ] Testumgebung mit Docker Compose bereitstellen.
-   [ ] CI für Formatierung, Tests, Vet und Builds einrichten.
-   [ ] Strukturierte Logs und Request-/Correlation-IDs einführen.
-   [ ] Metriken und Health-Endpunkte bereitstellen.
-   [ ] Release- und Versionsstrategie definieren.
-   [ ] Installations- und Upgrade-Prozess dokumentieren.
-   [ ] Sicherheitsreview vor produktivem Einsatz durchführen.

## Bewusst noch nicht priorisiert

-   Multi-Controller-HA.
-   Automatische Migration laufender Instanzen zwischen Nodes.
-   Kubernetes-Unterstützung.
-   Mandantenfähiges Hosting als Produkt.
-   Komplexe Abrechnung und Quotas.
