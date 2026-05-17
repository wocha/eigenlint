# ADR-005: Use mTLS for Service-to-Service Communication

## Status

| Field         | Value         |
|---------------|---------------|
| Status        | Maybe-Someday |
| Date          | 2026-05-01    |
| Author        | Timo Bigdon   |

## Context and Problem

Service-to-service communication requires authentication.

## Decision Drivers

- Security
- Standards compliance
- Low performance overhead

## Considered Options

### Option A: mTLS
Standard, widely adopted.

### Option B: JWT tokens
More flexible, but more complex.

## Decision Outcome

Selected: Option A, mTLS.

## Consequences

### Positive
- High security

### Negative
- Certificate management overhead

## Change Log

| Date       | Author      | Change             |
|------------|-------------|--------------------|
| 2026-05-01 | Timo Bigdon | Initial acceptance |
