# ADR-003: Use Redis for Caching

## Status

Status: Accepted
Date: 2026-03-01

## Context and Problem

We need a fast caching layer.

## Decision Drivers

- Low latency
- Simple operability

## Considered Options

### Option A: Redis
Widely used in-memory store.

### Option B: Memcached
Older, fewer features.

## Decision Outcome

Selected: Option A, Redis.

## Consequences

### Positive
- Low latency

### Negative
- Memory-bound

## Change Log

| Date       | Author      | Change             |
|------------|-------------|--------------------|
| 2026-03-01 | Timo Bigdon | Initial acceptance |
