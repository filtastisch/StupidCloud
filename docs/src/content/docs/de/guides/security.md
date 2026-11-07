---
title: Sicherheit
description: Sicherheitsziele, geplante Maßnahmen und betriebliche Hinweise für StupidCloud.
---

Dieses Dokument beschreibt **Sicherheitsziele und geplante Maßnahmen**. Es ist **keine** Bestätigung, dass alle Kontrollen bereits implementiert sind. Verhalten vor dem Produktiveinsatz am Code und an der Konfiguration prüfen.

## Grundprinzipien

- Standardmäßig minimale Netzwerkangriffsfläche.
- Jede administrative Aktion authentifizieren und autorisieren.
- Interne Verbindungen verschlüsseln und authentifizieren.
- Secrets aus Git, Logs und verteilbaren Templates heraushalten.
- Administrative Aktionen auditieren.
- Fehlerzustände sicher behandeln.

## SSH-Verwaltung (geplantes Design)

Das Zielsystem enthält einen integrierten SSH-Server (typisch lokal gebunden, z. B. `127.0.0.1:2222`), erreichbar über einen eingeschränkten Tunnelbenutzer auf dem System-SSH-Daemon. Wichtige Erwartungen:

- **Nur Public-Key-Authentifizierung** (keine Passwort-/Keyboard-Interactive-Fallbacks).
- Tunnelbenutzer ohne interaktive Shell und ohne erhöhte Rechte.
- StupidCloud-SSH-Schlüssel getrennt von Systemkonten verwalten.
- Jede SSH-Session isoliert mit eigenem Kontext/TUI.
- Berechtigungen pro Aktion prüfen.

> **Betrieblicher Hinweis:** Änderungen an der System-SSH-Konfiguration können den Zugang sperren. Mit `sshd -t` testen, eine Rettungssitzung offen halten und Tunnelbeschränkungen validieren.

## Controller–Node-Kommunikation (geplant)

- TLS-geschütztes gRPC.
- Gegenseitige Authentifizierung (z. B. mTLS) oder gleichwertige Node-Identität.
- Widerrufbare Node-Identitäten; Schutz gegen Registrierungs-/Reconnect-Missbrauch.
- Keine Tokens/Zertifikate in Logs.

## Docker- und Host-Sicherheit

Zugriff auf den Docker-Daemon bedeutet faktisch weitreichende Host-Rechte.

- Docker-Zugriff auf den Node-Agent beschränken.
- Docker-Socket niemals ungeschützt über TCP exponieren.
- Privileged-Container vermeiden; Capabilities und Mounts einschränken.
- Images bewusst auswählen/pinnen.
- Persistente Hostpfade validieren und kontrollieren.

## Netzwerk-Exponierung

- Backend-Minecraft-Server möglichst in privaten Netzwerken halten.
- Öffentliche Bind-Adressen explizit machen.
- Controller-Management-Ports nicht unnötig exponieren.
- Weboberflächen über TLS bereitstellen (typisch hinter geprüftem Reverse Proxy).

## Benutzer, Berechtigungen und Secrets

- RBAC/Autorisierung zentral in Application Services prüfen.
- Web-Authentifizierung von SSH-Public-Key-Authentifizierung trennen.
- Passwörter nur mit geeignetem Hashing speichern (bei Umsetzung).
- Secrets getrennt von Templates verwalten; Artefakt-Hashes prüfen, wo sinnvoll.
- Rechte auf Konfig-/Datenverzeichnisse einschränken; sensible Backups verschlüsseln.

## Checkliste vor Produktion

- Authentifizierung/Autorisierung testen.
- Node-Identitätsrotation/-widerruf testen.
- SSH-Tunnelbeschränkungen validieren.
- Container-Isolation und Docker-Rechte prüfen.
- API-Eingaben validieren/filtern.
- Logs/Fehler auf Secret-Leaks prüfen.
- Backup und Restore testen.