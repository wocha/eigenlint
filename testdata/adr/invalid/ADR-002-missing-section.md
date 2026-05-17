# ADR-002: Use Kafka for Event Streaming

## Status

| Field         | Value         |
|---------------|---------------|
| Status        | Accepted      |
| Date          | 2026-02-01    |
| Author        | Timo Bigdon   |

## Context and Problem

The system requires an event streaming platform for asynchronous communication between services.

## Considered Options

### Option A: Apache Kafka
Industry standard with high throughput.

### Option B: RabbitMQ
Easier to operate, but different guarantees.

## Decision Outcome

Selected: Option A, Apache Kafka.

## Consequences

### Positive
- High throughput
- Strongly established tooling

### Negative
- High operational complexity

## Change Log

| Date       | Author      | Change             |
|------------|-------------|--------------------|
| 2026-02-01 | Timo Bigdon | Initial acceptance |
