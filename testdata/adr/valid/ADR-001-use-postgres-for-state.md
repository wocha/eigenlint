# ADR-001: Use PostgreSQL for State Storage

## Status

| Field         | Value         |
|---------------|---------------|
| Status        | Accepted      |
| Date          | 2026-01-15    |
| Author        | Timo Bigdon   |
| Supersedes    | -             |
| Superseded by | -             |

## Context and Problem

The system requires a persistent database for state storage. Requirements include ACID guarantees, JSON support, and good operability.

## Decision Drivers

- ACID compliance for transactional safety
- JSON data type for flexible schema evolution
- Broad tooling support in cloud-native ecosystem
- Operability: backup, replication, monitoring
- Cost efficiency in self-hosting

## Considered Options

### Option A: PostgreSQL
Proven relational database with strong JSON support.

### Option B: MySQL
Widely used, but weaker JSON features and less strict defaults.

### Option C: SQLite
Simple, but not suitable for multi-node deployments.

## Decision Outcome

Selected: Option A, PostgreSQL. The combination of ACID guarantees and JSON support best meets the requirements.

## Consequences

### Positive
- Robust transactional safety
- Flexible schema through JSONB
- Wide tooling ecosystem

### Negative
- More complex operations than SQLite
- Higher resource requirements

### Neutral
- Migration to a different database would be expensive

## Validation Criteria

- PostgreSQL 16 deployed in the test environment
- Backup strategy with RPO < 1h validated
- Monitoring with Prometheus exporter active

## Change Log

| Date       | Author      | Change                       |
|------------|-------------|------------------------------|
| 2026-01-15 | Timo Bigdon | Initial acceptance           |
