# AuditMate — Portable System Audit Evidence Tool

AuditMate is a portable, offline system audit tool that produces **deterministic, verifiable snapshots of system state**.

AuditMate records **factual system evidence**.  
It does **not judge**, **score**, or **interpret** results.

This makes AuditMate useful for:

* Audits  
* Incident response  
* Change tracking  
* System trust verification  
* Forensics  

---

# Why AuditMate?

* **Deterministic Evidence** — Record system state consistently over time  
* **Change Visibility** — Compare snapshots and see exactly what changed  
* **Audit Friendly** — Human-readable and verifiable output  
* **Offline First** — No network access required  
* **Lightweight** — No agents, no background services  
* **Transparent** — Free version fully auditable  

---

# Product Tiers

## Free Tier — $0

Designed for admins, engineers, and individuals who want **system visibility and change tracking**.

Includes:

* ✅ System info (hostname, OS, kernel, uptime)  
* ✅ User accounts and admin membership  
* ✅ Running services  
* ✅ Installed packages  
* ✅ Open ports  
* ✅ Firewall state  
* ✅ Baseline snapshots  
* ✅ Deterministic diffs  
* ✅ Terminal summaries  
* ✅ Exit codes for automation  
* ✅ Fully offline  
* ✅ Open source and auditable  

Ideal for:

* Personal machines  
* Servers  
* Small teams  
* Labs  

---

## Professional Tier — $25 / month or $250 / year

Everything in Free, plus **audit-ready evidence features**.

Includes:

* ✅ Evidence signing  
* ✅ Evidence manifest generation  
* ✅ Bundle export for auditors  
* ✅ Historical snapshot tracking  
* ✅ License-based deployment  
* ✅ Offline verification  

Ideal for:

* Companies preparing for audits  
* Compliance teams  
* Incident response  
* Forensics  

---

# How It Works

1. **Collect** — AuditMate gathers system state  
2. **Baseline** — Save snapshot  
3. **Diff** — Compare changes  
4. **Report** — Human-readable output  
5. **Evidence (Pro)** — Signed audit artifacts  

---

# Example Output

```
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
```

---

# Pricing

| Plan            | Price | Notes                                                 |
| --------------- | ----- | ----------------------------------------------------- |
| Free            | $0    | Fully offline, open source                            |
| Monthly         | $25   | Professional features                                 |
| Yearly          | $250  | Professional features                                 |
| Founding Member | $200  | Lifetime single-server license (limited to first 100) |

Contact: **[audit-mate@proton.me](mailto:audit-mate@proton.me)**

---

# Trust & Verification

AuditMate is designed for transparency.

* Free version is open source
* No telemetry
* No network access
* No auto updates
* Deterministic output
* Verifiable binaries

Both tiers share the **same collectors** for consistent results.

---

# Philosophy

AuditMate does not:

* Score security
* Judge configuration
* Detect vulnerabilities

AuditMate **records facts**.

You decide what matters.
