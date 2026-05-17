# ADR-006: Use Cilium for Network Policies

## Status

| Field         | Value      |
|---------------|------------|
| Status        | Accepted   |
| Date          | 2026-06-01 |
| Author        | Timo Bigdon |

## Considered Options

### Option A: Cilium
eBPF-basiert, modern.

### Option B: Calico
Älter, bewährt.

## Kontext und Problem

Wir brauchen Network-Policies in Kubernetes.

## Decision Drivers

- eBPF-Vorteile
- Performance

## Decision Outcome

Cilium.

## Konsequenzen

### Positiv
- Performance

### Negativ
- Lernkurve

## Change Log

| Datum      | Author      | Änderung           |
|------------|-------------|--------------------|
| 2026-06-01 | Timo Bigdon | Initial Acceptance |
