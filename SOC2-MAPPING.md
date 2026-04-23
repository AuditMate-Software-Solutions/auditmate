# System Evidence Coverage Mapping

This document maps what AuditMate collects to common categories used in enterprise audit reviews.

This is NOT a compliance claim.

AuditMate does not certify, enforce, or validate compliance.

It only records system state.

---

## Access-related Evidence

Collected:
- local users
- administrator membership

Purpose:
- visibility into privilege changes

---

## System Activity Evidence

Collected:
- running services
- installed packages
- open ports
- firewall state

Purpose:
- system state snapshot comparison
- change detection over time

---

## Change Detection Model

AuditMate uses:
- baseline snapshots
- deterministic diffing

Purpose:
- identify differences between two system states

---

## Important Clarification

AuditMate does NOT:
- enforce controls
- validate policies
- determine compliance status

It only provides **raw system evidence**