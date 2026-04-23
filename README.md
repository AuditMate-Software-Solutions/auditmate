# AuditMate — Portable System Audit Evidence Tool

AuditMate is a portable, offline system audit tool that produces deterministic, verifiable snapshots of system state.

AuditMate records factual system evidence. It does not judge, score, or interpret results.

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

- System info (hostname, OS, kernel, uptime)  
- Users & admin membership  
- Running services  
- Installed packages  
- Open ports  
- Firewall state  
- Baseline snapshots  
- Deterministic diffs  
- Fully offline  
- Open source  

---

## Professional Tier — $25 / month or $250 / year

- Evidence signing  
- Evidence manifest generation  
- Exportable audit bundles  
- Historical tracking  
- Multi-system comparison  
- Offline verification  

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

===================================

```
```
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
```
