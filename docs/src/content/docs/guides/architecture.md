---
title: Architecture
description: The planned controller, node agent, and management interfaces.
---

StupidCloud is designed around a central controller and a set of node agents. The controller holds the desired configuration and assigns work; nodes apply that work on their hosts and report their current state.

## Controller

The controller is the coordination point. Its planned responsibilities include scheduling, user and template management, the API, and management interfaces. It should be the source of truth for the desired state of the cluster.

## Node agent

A node agent runs on each managed host. It communicates with the controller, creates and manages server containers through Docker, and reports results and state back to the controller.

## Management interfaces

The project plans SSH based CLI and TUI access through the controller, plus a separate web interface. These interfaces should use the same controller APIs and operations.

## Planned communication and storage

The current target design uses gRPC and Protocol Buffers for internal controller to node communication, a REST API for web clients, WebSockets for live updates, and PostgreSQL for persistent data. These are plans and may change as implementation progresses.

## Related pages

- [Project status and planned technologies](/reference/project-status/)
