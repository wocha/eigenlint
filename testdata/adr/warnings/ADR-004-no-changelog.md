# ADR-004: Use OpenTelemetry for Observability

## Status

| Field         | Value         |
|---------------|---------------|
| Status        | Proposed      |
| Date          | 2026-04-01    |
| Author        | Timo Bigdon   |

## Context and Problem

The system needs standardized observability.

## Decision Drivers

- Vendor neutrality
- Broad tooling support
- Standardization

## Considered Options

### Option A: OpenTelemetry
Vendor-neutral standard.

### Option B: Vendor-specific solution
Faster start, but lock-in.

## Decision Outcome

Selected: Option A, OpenTelemetry.

## Consequences

### Positive
- Vendor neutrality

### Negative
- Slightly higher setup overhead
