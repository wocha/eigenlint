# ADR-005: Use mTLS for Service-to-Service Communication

## Status

| Field         | Value         |
|---------------|---------------|
| Status        | Maybe-Someday |
| Date          | 2026-05-01    |
| Author        | Timo Bigdon   |

## Kontext und Problem

Service-zu-Service-Kommunikation braucht Authentifizierung.

## Decision Drivers

- Sicherheit
- Standardkonformität
- Geringer Performance-Overhead

## Considered Options

### Option A: mTLS
Standard, weit verbreitet.

### Option B: JWT-Token
Flexibler, aber komplexer.

## Decision Outcome

Gewählt wird Option A: mTLS.

## Konsequenzen

### Positiv
- Hohe Sicherheit

### Negativ
- Zertifikats-Management

## Change Log

| Datum      | Author      | Änderung           |
|------------|-------------|--------------------|
| 2026-05-01 | Timo Bigdon | Initial Acceptance |
