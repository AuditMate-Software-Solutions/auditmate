# AuditMate

Offline System Auditing Tool

> Capture system state. Detect drift. No cloud. No telemetry. No installation.

---

## ⚡ Overview

AuditMate is a portable, offline system auditing tool that records deterministic snapshots of a machine’s state.

It is designed for:
- system change tracking
- forensic support
- infrastructure drift detection
- local evidence collection

It does NOT:
- send data anywhere
- run in the background
- perform scanning or exploitation
- require installation

---

## 🧠 Core Features

- Fully offline execution
- Open-source free version
- Windows & Linux binaries
- Real-time drift detection
- Deterministic system snapshots
- Human-readable + JSON output

---

## 📦 Download

- Windows: `auditmate.exe`
- Linux: `auditmate`

No installer required. Fully portable.

---

## 🔐 Privacy

- No telemetry
- No network communication
- No cloud dependency
- No background services
- All processing happens locally

---

## 💰 Pricing

### Free
- System information (OS, hostname, uptime)
- Users & services
- Installed packages
- Open ports
- Drift detection
- Open source

### Pro
- $25 / month
- $250 / year

Includes:
- Historical tracking
- Export reports (JSON / CSV / HTML)
- Audit-ready reporting formats
- Multi-system comparison
- Signed outputs

### Founding Member
- $200 / year (first 100 users, normally $250/year)

---

## 🖥 Example Output (Windows)

---

### Clean System

```text
C:\> auditmate.exe

AuditMate
-----------------------------------
Hostname : WORKSTATION-01
OS       : Windows 11 Pro
Uptime   : 1 day, 2 hours

Users    : 3
Services : 128
Packages : 94
Ports    : 6

Status   : Clean
-----------------------------------
Warnings : 0
Errors   : false
Reset    : false
Report   : auditmate-output\audit.json
Duration : 52ms
