# AuditMate — Portable System Audit Evidence Tool

AuditMate is a portable, offline system audit tool that produces deterministic, verifiable snapshots of system state.

AuditMate records factual system evidence. It does not judge, score, or interpret results.

It is used for audits, incident response, change tracking, system verification, and forensics.

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

---

C:\> auditmate.exe

AuditMate
-----------------------------------
Hostname : WORKSTATION-01
OS       : Windows 11 Pro
Uptime   : 1 day, 3 hours

Users    : 4
Services : 130
Packages : 96
Ports    : 8

Status   : Drift detected

Changes:
 + Users:
   + hacker (admin)

 + Services:
   + sshd
   + remote-control

 + Open ports:
   + 22/tcp
   + 3389/tcp

-----------------------------------
Warnings : 1
Errors   : false
Reset    : false
Report   : auditmate-output\audit.json
Duration : 61ms

---

# Pricing

Free:
$0 — fully offline open-source version

Professional:
$25 / month
$250 / year
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
