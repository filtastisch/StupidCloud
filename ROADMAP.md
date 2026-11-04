# Roadmap

Die Roadmap beschreibt Entwicklungsphasen und Ergebnisse, keine festen
Veröffentlichungstermine. Eine Phase gilt erst dann als abgeschlossen,
wenn ihre Akzeptanzkriterien erfüllt und getestet sind.

## Phase 0 -- Projektfundament

**Ziel:** Ein sauber startendes und testbares Go-Projekt.

Ergebnisse: - Festgelegte Repository- und Package-Struktur. -
Controller-Einstiegspunkt. - Logger und Konfigurationsgrundlage. -
Start-/Shutdown-Lifecycle. - Erste Unit-Tests und CI.

Akzeptanz: - Controller lässt sich lokal starten. - SIGINT und SIGTERM
führen zu einem kontrollierten Shutdown. - Fehler werden mit Kontext
protokolliert. - `go test ./...` läuft erfolgreich.

## Phase 1 -- Controller-Kern

**Ziel:** Die zentralen Anwendungsdienste und Persistenz bereitstellen.

Ergebnisse: - PostgreSQL-Verbindung und Migrationen. - Benutzer- und
Berechtigungsmodell. - Node-, Service-, Template- und Instanzmodelle. -
Event- und Job-Grundlage. - Konfigurationsvalidierung.

Akzeptanz: - Datenbank kann initialisiert und migriert werden. -
Kernobjekte können über Application Services verwaltet werden. -
Fachlogik ist unabhängig von HTTP, SSH und TUI testbar.

## Phase 2 -- Ein einzelner Node

**Ziel:** Einen Node zuverlässig mit dem Controller verbinden und
Docker-Container steuern.

Ergebnisse: - Eigenständiger Node-Agent. - gRPC-Protokoll. - Sichere
Registrierung und Heartbeats. - Docker-Integration. - Lokaler Cache und
Statusmeldungen.

Akzeptanz: - Node registriert sich beim Controller. - Controller kann
einen Testcontainer starten und stoppen. - Ein Verbindungsabbruch wird
erkannt und überstanden. - Bereits laufende Container werden bei einem
temporären Controller-Ausfall nicht unkontrolliert beendet.

## Phase 3 -- Minecraft-Instanzen und Templates

**Ziel:** Minecraft-Server aus versionierten Templates erzeugen.

Ergebnisse: - Template-Manager. - Runtime- und
Minecraft-Konfiguration. - Instanz-Lifecycle. - Persistente Daten. -
Healthchecks und Logzugriff. - Portreservierung.

Akzeptanz: - Eine Instanz kann aus einem Template erstellt, gestartet,
überwacht und entfernt werden. - Template-Versionen sind
nachvollziehbar. - Persistente Daten bleiben bei einem
Container-Neustart erhalten. - Portkonflikte werden erkannt und korrekt
behandelt.

## Phase 4 -- Multi-Node und Scheduling

**Ziel:** Instanzen über mehrere Hosts verteilen.

Ergebnisse: - Node-Ressourcen und Kapazitäten. - Scheduling-Regeln. -
Node-Labels und Einschränkungen. - Zentrale Portbereiche. - Privates
Netzwerk für Backend-Kommunikation.

Akzeptanz: - Controller kann geeignete Nodes auswählen. -
Ressourcenlimits werden berücksichtigt. - Ports sind pro Node eindeutig
reserviert. - Nicht erreichbare Nodes werden korrekt als nicht verfügbar
behandelt.

## Phase 5 -- Proxy-Integration

**Ziel:** Minecraft-Instanzen automatisch in Proxy-Netzwerke
integrieren.

Ergebnisse: - Bridge für Velocity. - Bridge für BungeeCord. -
Proxy-Gruppen. - Registrierung nach Healthcheck. - Reconciliation und
Draining.

Akzeptanz: - Eine bereite Lobby wird automatisch registriert. - Eine
gestoppte Instanz wird kontrolliert entfernt. - Proxy-Neustarts führen
zur Wiederherstellung des gewünschten Zustands.

## Phase 6 -- Management

**Ziel:** Die Plattform administrierbar machen.

Ergebnisse: - CLI. - SSH-Server mit Public-Key-only-Authentifizierung. -
Mehrbenutzerfähige TUI. - Rollen- und Rechteverwaltung. - Audit-Events.

Akzeptanz: - Mehrere Nutzer können parallel eigene SSH-Sessions
verwenden. - Berechtigungen werden bei jeder Aktion geprüft. -
SSH-Passwortauthentifizierung ist nicht vorgesehen.

## Phase 7 -- Webinterface und Betrieb

**Ziel:** Externe Verwaltung und produktionsnahe Betriebsabläufe
ermöglichen.

Ergebnisse: - REST API und WebSocket. - React-/Next.js-Webinterface. -
TLS- und Reverse-Proxy-Konzept. - Monitoring und Metriken. -
Installations-, Backup- und Upgrade-Prozess.

Akzeptanz: - Webinterface kann Nodes und Instanzen anzeigen und
verwalten. - Live-Status wird aktualisiert. - Backups und
Wiederherstellung sind getestet. - Installations- und
Upgrade-Anleitungen sind nachvollziehbar.

## Spätere Erweiterungen

Erst nach einem stabilen Grundsystem evaluieren: - Hochverfügbarkeit des
Controllers. - Geplante und automatische Wartungsfenster. -
Instanzmigration zwischen Nodes. - Erweiterte Ressourcen- und
Kostenplanung. - Zusätzliche Game-Server-Runtimes. -
Plugin-/Extension-System.
