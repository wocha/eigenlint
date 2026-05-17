# ADR-002: Use Kafka for Event Streaming

## Status

| Field         | Value      |
|---------------|------------|
| Status        | Accepted   |
| Date          | 2026-02-01 |
| Author        | Timo Bigdon |

## Kontext und Problem

Das System benötigt eine Event-Streaming-Plattform für asynchrone Kommunikation zwischen Services.

## Considered Options

### Option A: Apache Kafka
Industry-Standard mit hoher Throughput.

### Option B: RabbitMQ
Einfacher zu betreiben, aber andere Garantien.

## Decision Outcome

Gewählt wird Option A: Apache Kafka.

## Konsequenzen

### Positiv
- Hohe Throughput
- Stark etabliertes Tooling

### Negativ
- Hohe Operations-Komplexität

## Change Log

| Datum      | Author      | Änderung           |
|------------|-------------|--------------------|
| 2026-02-01 | Timo Bigdon | Initial Acceptance |
