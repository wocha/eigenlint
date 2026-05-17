# ADR-001: Use PostgreSQL for State Storage

## Status

| Field         | Value      |
|---------------|------------|
| Status        | Accepted   |
| Date          | 2026-01-15 |
| Author        | Timo Bigdon |
| Supersedes    | -          |
| Superseded by | -          |

## Kontext und Problem

Das System benötigt eine persistente Datenbank für State-Storage. Die Anforderungen umfassen ACID-Garantien, JSON-Unterstützung und gute Operability.

## Decision Drivers

- ACID-Compliance für Transaktions-Sicherheit
- JSON-Datentyp für flexible Schema-Evolution
- Verbreitete Tool-Unterstützung im Cloud-Native-Ökosystem
- Operability: Backup, Replication, Monitoring
- Kosten-Effizienz im Self-Hosting

## Considered Options

### Option A: PostgreSQL
Bewährte relationale Datenbank mit starker JSON-Unterstützung.

### Option B: MySQL
Verbreitet, aber schwächere JSON-Features und weniger strikte Defaults.

### Option C: SQLite
Einfach, aber nicht geeignet für Multi-Node-Deployments.

## Decision Outcome

Gewählt wird Option A: PostgreSQL. Die Kombination aus ACID-Garantien und JSON-Support trifft die Anforderungen am besten.

## Konsequenzen

### Positiv
- Robuste Transaktions-Sicherheit
- Flexibles Schema durch JSONB
- Breites Tool-Ökosystem

### Negativ
- Komplexere Operations als SQLite
- Höhere Ressourcen-Anforderungen

### Neutral
- Migration zu anderer Datenbank wäre aufwändig

## Validation Criteria

- PostgreSQL 16 deployed in der Test-Umgebung
- Backup-Strategie mit RPO < 1h validiert
- Monitoring mit Prometheus-Exporter aktiv

## Change Log

| Datum      | Author      | Änderung                          |
|------------|-------------|-----------------------------------|
| 2026-01-15 | Timo Bigdon | Initial Acceptance                |
