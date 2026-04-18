# SOC 2 Control Mapping (CC6 & CC7)

This document describes how AuditMate supports **technical evidence collection** commonly requested for SOC 2 assessments.

AuditMate **does not certify compliance** and **does not replace audits**.

---

# CC6 — Logical Access Controls

## CC6.1 – User Access Management

Collected Evidence:

* Local users
* Administrator group membership

Audit Use:

* Identify privileged accounts
* Detect changes over time

---

## CC6.2 – Privileged Access

Collected Evidence:

* Administrator group changes

Audit Use:

* Track privilege modifications

---

# CC7 — System Operations

## CC7.1 – System Visibility

Collected Evidence:

* Running services
* Open ports
* Firewall status

Audit Use:

* Identify unexpected system changes

---

## CC7.2 – Change Detection

Mechanism:

* Baselines
* Deterministic diffs

Audit Use:

* Identify system modifications

---

## CC7.3 – Patch Awareness

Collected Evidence:

* Pending updates (where supported)

Audit Use:

* Patch awareness at execution time

---

# Output Characteristics

Professional tier supports:

* Timestamped output
* Evidence bundles
* Signed artifacts
* Historical snapshots

---

# Scope Statement

AuditMate provides **technical system evidence**.

Compliance decisions remain the responsibility of:

* Organizations
* Auditors
* Policies
