---
title: Glossar
description: Wichtige Begriffe in der StupidCloud-Dokumentation und -Konzeption.
---

| Begriff | Bedeutung |
|---|---|
| **Controller** | Zentrale Control Plane und Quelle des gewünschten Clusterzustands. |
| **Node** | Verwalteter Host, auf dem ein Node-Agent und Docker laufen. |
| **Node-Agent** | Lokaler Prozess, der Controller-Aufträge auf einem Node ausführt. |
| **Cluster** | Gesamtheit aus Controller und registrierten Nodes. |
| **Service** | Logische Gruppierung einer oder mehrerer Serverinstanzen. |
| **Instanz** | Einzelne ausführbare Serverinstanz, typischerweise ein Docker-Container. |
| **Template** | Versionierte Vorlage für Runtime, Dateien und Serverkonfiguration. |
| **Artefakt** | Herunterladbare Datei/Paket, das eine Instanz benötigt. |
| **Sollzustand (Desired State)** | Vom Controller gewünschter Zustand einer Ressource. |
| **Istzustand (Observed State)** | Von Node oder Bridge tatsächlich gemeldeter Zustand. |
| **Reconciliation** | Abgleich zwischen Ist- und Sollzustand. |
| **Scheduling** | Auswahl eines Nodes für eine neue oder zu verschiebende Instanz. |
| **Proxy-Bridge** | Komponente/Plugin (z. B. für Velocity/BungeeCord), das Proxy-Änderungen umsetzt und Zustand meldet. |
| **Portreservierung** | Zentrale Reservierung eines Hostports für eine Instanz. |
| **Control Plane** | Komponenten, die Zustand und Orchestrierung verwalten. |
| **Data Plane** | Komponenten, die den eigentlichen Spielverkehr transportieren. |