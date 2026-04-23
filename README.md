# AuditMate — Portable System Audit Evidence Tool

AuditMate is a portable, offline system audit tool that produces deterministic, verifiable snapshots of system state.

AuditMate records factual system evidence. It does not judge, score, or interpret results.

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
