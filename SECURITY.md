# Security Policy

## Overview

AuditMate is designed to be a **read-only, offline system auditing tool**.

AuditMate records system state and produces deterministic output.

Security and transparency are core design principles.

---

## What AuditMate Does

* Executes OS-provided commands
* Reads system state
* Writes output locally
* Exits immediately after execution

---

## What AuditMate Does NOT Do

AuditMate intentionally avoids:

❌ Network connections
❌ Telemetry or analytics
❌ Automatic updates
❌ Background services
❌ Privilege escalation
❌ Configuration changes
❌ Vulnerability scanning
❌ Exploitation
❌ Remote execution

---

## Security Model

AuditMate:

* Runs with current user privileges
* Uses OS-native commands
* Produces deterministic output
* Stores files locally

---

## Supply Chain Trust

Users can:

* Build from source
* Verify checksums
* Review collectors
* Audit output

Professional tier:

* Signed evidence artifacts
* Offline verification
* License validation

---

## Reporting a Vulnerability

Email: **[audit-mate@proton.me](mailto:audit-mate@proton.me)**

Please include:

* OS version
* AuditMate version
* Steps to reproduce

Response target: **72 hours**

---

## Supported Versions

* Latest Free release
* Supported Professional releases

Thank you for helping keep AuditMate secure.
