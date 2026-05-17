# ADR-004: Use OpenTelemetry for Observability

## Status

| Field         | Value      |
|---------------|------------|
| Status        | Proposed   |
| Date          | 2026-04-01 |
| Author        | Timo Bigdon |

## Kontext und Problem

Das System braucht standardisierte Observability.

## Decision Drivers

- Vendor-Neutralität
- Breite Tool-Unterstützung
- Standardisierung

## Considered Options

### Option A: OpenTelemetry
Vendor-neutraler Standard.

### Option B: Vendor-spezifische Lösung
Schneller Start, aber Lock-In.

## Decision Outcome

Gewählt wird Option A: OpenTelemetry.

## Konsequenzen

### Positiv
- Vendor-Neutralität

### Negativ
- Etwas mehr Setup-Aufwand
