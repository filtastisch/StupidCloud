---
title: StupidCloud
description: Documentation for the planned Go-based Minecraft server orchestration platform.
template: splash
hero:
  tagline: Plan, run, and manage distributed Minecraft server infrastructure.
  actions:
    - text: Get started
      link: /guides/getting-started/
      icon: right-arrow
    - text: Architecture
      link: /guides/architecture/
      icon: open-book
    - text: Project status
      link: /reference/project-status/
      variant: minimal
---

StupidCloud is an **early-stage, Go-based server orchestration platform** for managing distributed **Minecraft** server infrastructure. The target design runs individual Minecraft servers in separate **Docker containers**, coordinated by a central **controller** and executed by **node agents** on managed hosts.

> **Status:** The documentation describes the intended system and the current state of the codebase. Many components are planned but not implemented yet. Check the [project status](/reference/project-status/) and [roadmap](/reference/roadmap/) before relying on a feature.

## What StupidCloud aims to provide

- Central management of multiple hosts (**nodes**).
- Minecraft server instances running in Docker containers.
- Reusable templates and consistent configuration.
- Scheduling and centralized port management/reservation.
- Integration with **Velocity** and **BungeeCord** proxies.
- Management via **SSH** (public-key oriented design), CLI, TUI, and a web interface.
- Secure, authenticated controller–node communication.

## Start here

- **Run the current controller scaffold:** [Getting started](/guides/getting-started/)
- **Understand the target design:** [Architecture](/guides/architecture/)
- **Contribute safely:** [Development](/guides/development/) and [Security](/guides/security/)
- **Track progress:** [Project status](/reference/project-status/) and [Roadmap](/reference/roadmap/)