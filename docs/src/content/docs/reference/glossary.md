---
title: Glossary
description: Key terms used in StupidCloud documentation and design.
---

| Term | Meaning |
|---|---|
| **Controller** | Central control plane and source of desired cluster state. |
| **Node** | Managed host running a node agent and Docker. |
| **Node agent** | Local process that executes controller tasks on a node. |
| **Cluster** | Controller plus registered nodes. |
| **Service** | Logical grouping of one or more server instances. |
| **Instance** | A single runnable server unit, typically a Docker container. |
| **Template** | Versioned blueprint for runtime, files, and server configuration. |
| **Artifact** | Downloadable file/package required by an instance. |
| **Desired state** | The state the controller wants a resource to have. |
| **Observed state** | The state actually reported by a node or bridge. |
| **Reconciliation** | Aligning observed state with desired state. |
| **Scheduling** | Selecting a node for a new or moved instance. |
| **Proxy bridge** | Component/plugin (e.g., for Velocity/BungeeCord) that applies proxy changes and reports state. |
| **Port reservation** | Central allocation of a host port to an instance. |
| **Control plane** | Components that manage state and orchestration. |
| **Data plane** | Components that carry actual game/player traffic. |