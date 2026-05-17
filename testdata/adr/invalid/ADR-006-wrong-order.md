# ADR-006: Use Cilium for Network Policies

## Status

| Field         | Value         |
|---------------|---------------|
| Status        | Accepted      |
| Date          | 2026-06-01    |
| Author        | Timo Bigdon   |

## Considered Options

### Option A: Cilium
eBPF-based, modern.

### Option B: Calico
Older, proven.

## Context and Problem

We need network policies in Kubernetes.

## Decision Drivers

- eBPF advantages
- Performance

## Decision Outcome

Cilium.

## Consequences

### Positive
- Performance

### Negative
- Learning curve

## Change Log

| Date       | Author      | Change             |
|------------|-------------|--------------------|
| 2026-06-01 | Timo Bigdon | Initial acceptance |
