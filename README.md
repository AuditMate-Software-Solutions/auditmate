# AuditMate — Portable System Audit Evidence Tool

AuditMate is a portable, offline system audit tool that produces **deterministic, verifiable snapshots of system state**.

AuditMate records **factual system evidence**. It does not judge, score, or interpret results.

It is used for:
- audits
- incident response
- change tracking
- system verification
- forensics

---

# Why AuditMate

- Deterministic system snapshots
- Exact change tracking (diffs)
- Offline-first (no network required)
- Lightweight CLI tool
- Human-readable + JSON output
- Fully transparent free version

---

# Product Tiers

## Free Tier — $0

- System information (hostname, OS, kernel, uptime)
- User accounts and admin membership
- Running services
- Installed packages
- Open ports
- Firewall state
- Baseline snapshot creation
- Deterministic diff engine
- Fully offline operation
- Open source

---

## Professional Tier — $25 / month or $250 / year

- Evidence signing
- Evidence manifest generation
- Exportable audit bundles
- Historical snapshot tracking
- Multi-system comparison
- Offline verification
- License-based deployment

Founding Member:
- $200 / year (first 100 users, normally $250/year)

---

# How It Works

1. Collect system state
2. Create baseline snapshot
3. Compare snapshots (diff engine)
4. Generate report
5. (Pro) Sign and export evidence bundle

---

# Example Output (Windows)

```text
C:\> auditmate-pro.exe

============== AuditMate Report ==============
--------------------------------------------------
SYSTEM INFORMATION
--------------------------------------------------
Hostname          : MY-LAPTOP
Host ID           : b577b2d675174a94b735d3b2e3e47683
Baseline Created  : 2026-04-09T01:25:04Z
Current Snapshot  : 2026-04-09T03:55:42Z
OS                : Windows 11 Pro
Kernel            : 10.0.22631
Uptime            : 1 day, 3 hours, 12 minutes

USER ACCOUNTS
--------------------------------------------------
Administrators:
  - Administrator
  - hacker
  - john

Regular Users:
  - JOHN
  - Guest

SYSTEM SUMMARY
--------------------------------------------------
Services Running   : 130
Installed Programs : 95
Open Ports         : 8
Firewall State     : enabled

BASELINE COMPARISON
--------------------------------------------------
Changes detected:

  Administrators:
  + hacker (NEW ADMINISTRATOR)

  Services:
  + sshd (NEW SERVICE)
  + backdoor-svc (NEW SERVICE)

  Open Ports:
  + 22/tcp (SSH enabled)
  + 3389/tcp (RDP enabled)

================================================

JSON report saved to: auditmate-output\audit_report.json

---

# Pricing

Free:
$0 — fully offline open-source version

Professional:
$25 / month
$250 / year

Includes:
- historical tracking
- export reports (JSON / CSV / HTML)
- audit-ready bundles
- multi-system comparison
- signed outputs

Founding Member:
$200 / year (first 100 users, normally $250/year)

---

# Trust & Verification

- No telemetry
- No network communication
- No background services
- No cloud dependency
- Deterministic output
- Verifiable builds

Free version is fully open source.

---

# Philosophy

AuditMate does not:
- score systems
- assign risk
- detect vulnerabilities

AuditMate records facts.

You decide what matters.
