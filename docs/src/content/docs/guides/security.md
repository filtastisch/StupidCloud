---
title: Security
description: Security goals, planned controls, and operational cautions for StupidCloud.
---

This document describes **security goals and planned measures**. It is **not** a claim that all controls are implemented. Verify behavior against the code and configuration before production use.

## Principles

- Expose the minimum network surface by default.
- Authenticate and authorize every administrative action.
- Encrypt and authenticate internal connections.
- Keep secrets out of Git, logs, and distributable templates.
- Audit administrative actions.
- Fail safely.

## SSH management (planned design)

The target design includes an integrated SSH server (commonly bound to localhost, e.g. `127.0.0.1:2222`) accessed through a restricted tunnel user on the system SSH daemon. Key expectations:

- **Public-key authentication only** (no password/keyboard-interactive fallbacks).
- Tunnel user without interactive shell or elevated privileges.
- Separate management of StupidCloud SSH keys from system accounts.
- Each SSH session isolated with its own context/TUI instance.
- Authorization checked per action.

> **Operational caution:** Changing system SSH configuration can lock you out. Test with `sshd -t`, keep a recovery session open, and validate tunnel restrictions before relying on them.

## Controller–node communication (planned)

- TLS-protected gRPC.
- Mutual authentication (e.g., mTLS) or an equivalent node identity mechanism.
- Revocable node identities; protection against registration/reconnect abuse.
- No tokens/certificates in logs.

## Docker and host security

Access to the Docker daemon is effectively high-privilege host access.

- Limit Docker access to the node agent.
- Never expose the Docker socket over unprotected TCP.
- Avoid privileged containers; restrict capabilities and mounts.
- Pin/select images deliberately.
- Validate and control persistent host paths.

## Network exposure

- Keep backend Minecraft servers on private networks where possible.
- Make public bind addresses explicit.
- Avoid exposing controller management ports unnecessarily.
- Serve web interfaces over TLS (typically behind a reviewed reverse-proxy setup).

## Users, permissions, and secrets

- Centralize RBAC/authorization checks in application services.
- Separate web authentication from SSH public-key authentication.
- Store passwords only with appropriate password hashing (when implemented).
- Manage secrets separately from templates; verify artifact hashes where applicable.
- Restrict permissions on config/data directories; encrypt sensitive backups.

## Pre-production review checklist

- Test authentication/authorization paths.
- Test node identity rotation/revocation.
- Validate SSH tunnel restrictions.
- Review container isolation and Docker permissions.
- Validate/sanitize API inputs.
- Scan logs/errors for secret leakage.
- Test backup and restore.