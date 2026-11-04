---
title: StupidCloud
description: Documentation for the planned Go based Minecraft server orchestration platform.
template: splash
hero:
  tagline: Plan, run, and manage distributed Minecraft server infrastructure.
  actions:
    - text: Explore the architecture
      link: /guides/architecture/
      icon: right-arrow
    - text: Project status
      link: /reference/project-status/
      icon: open-book
---

StupidCloud is an early stage project for centrally managing Minecraft servers across multiple hosts. Servers are intended to run in Docker containers, with a controller coordinating configuration and work across node agents.

> **Project status:** StupidCloud is under early development. This documentation describes the intended system; not every component is implemented yet.

## What StupidCloud aims to provide

- Central management of multiple hosts (nodes).
- Minecraft server instances running in Docker containers.
- Reusable templates and consistent configuration.
- Scheduling and centralized port management.
- Integration with Velocity and BungeeCord.
- SSH, CLI, TUI, and web management interfaces.

## Start here

- Read the [architecture guide](/guides/architecture/) for the planned components and their responsibilities.
- Check the [project status reference](/reference/project-status/) before relying on a feature.
