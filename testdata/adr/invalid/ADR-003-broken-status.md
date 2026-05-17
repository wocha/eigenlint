# ADR-003: Use Redis for Caching

## Status

Status: Accepted
Date: 2026-03-01

## Kontext und Problem

Wir brauchen eine schnelle Cache-Schicht.

## Decision Drivers

- Niedrige Latenz
- Einfache Operability

## Considered Options

### Option A: Redis
Verbreiteter In-Memory-Store.

### Option B: Memcached
Älter, weniger Features.

## Decision Outcome

Gewählt wird Option A: Redis.

## Konsequenzen

### Positiv
- Niedrige Latenz

### Negativ
- Memory-Bound

## Change Log

| Datum      | Author      | Änderung           |
|------------|-------------|--------------------|
| 2026-03-01 | Timo Bigdon | Initial Acceptance |
