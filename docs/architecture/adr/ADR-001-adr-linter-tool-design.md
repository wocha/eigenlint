# ADR-001: ADR-Linter Tool Design

## Status

| Field         | Value                                    |
|---------------|------------------------------------------|
| Status        | Accepted                                 |
| Date          | 2026-05-15                               |
| Author        | Claude (AI-generated, see AI-USAGE.md)   |
| Supersedes    | -                                        |
| Superseded by | -                                        |

## Context and Problem

As the ADR collection in EigenState-Core grows, manual consistency checking becomes unreliable. Mandatory sections may be forgotten, status values may be inconsistent, change logs may be missing.

Additionally, this linter serves as a pilot project for the MADR-with-AI-Generation-Annex methodology documented in ADR-000. The tool is used to validate the methodology from an applier's perspective.

## Decision Drivers

- Fast execution for CI/CD integration (sub-second for 20 ADRs)
- Single-binary deployment without runtime dependencies
- Idiomatic language for CLI tools
- Extensible with new checks without architectural changes
- Learning project for the author's first complete Go implementation
- Validation of the MADR+Annex methodology against a real use case

## Considered Options

### Option A: Python Script
Familiar, fast to write, but introduces runtime dependencies and offers no learning value.

### Option B: Bash Script
Zero build overhead, but does not scale cleanly beyond 100 lines.

### Option C: Go Binary
Learning curve required, but idiomatic and produces single-binary distributables.

## Decision Outcome

Selected: Option C, Go binary. Learning value and deployment advantages outweigh the initial setup effort.

## AI-Generation-Annex

### Constraints for Code Generation

- MUST: Use github.com/spf13/cobra for CLI structure, not the flag package
- MUST: Follow Go standard layout with cmd/eigenlint/main.go and internal/
- MUST: Each check is its own function with a clear signature taking ADR data and returning Issues
- MUST: Output is both JSON-serializable AND human-readable, controlled via --format=json|text (default: text)
- MUST: Exit code 0 on pass, 1 on lint failures, 2 on tool-internal errors
- MUST: Configuration path via --path flag, default docs/architecture/adr
- SHOULD: Use a markdown parser only when necessary, otherwise regex
- SHOULD: Log to stderr, output to stdout (Unix pipeline compatible)

### Forbidden Implementation Patterns

- No init() functions with side effects
- No direct os.Exit() outside of main.go
- No string concatenation for JSON output
- No Java-style Builder pattern for simple structs
- No global variables for configuration
- No panic on expected errors

### Machine-Readable References

- Requires: ADR-000 (ADR format convention)
- Conflicts-With: -
- Supersedes: -
- Implementation-Hint: notes/go-lernpfad.md
- Validation-Script: scripts/validate-adr-linter.sh

## Detailed Specifications

### MVP Checks to Implement First

1. status_check: status table present with Status, Date, Author fields
2. sections_check: mandatory sections present (Context and Problem, Decision Drivers, Considered Options, Decision Outcome, Consequences)
3. changelog_check: Change Log section exists or warning is raised

### Output Format Requirements

- Text mode: one block per ADR with check name, pass/fail status, optional line number and message
- JSON mode: summary with total/passed/failed, plus per-ADR results array
- Both modes must be processable with grep or jq respectively

## Consequences

### Positive
- Automatically enforceable ADR consistency
- Tool as portfolio asset
- Complete Go learning cycle
- Methodology validation against real use case

### Negative
- Go learning curve
- Maintenance burden as ADR format evolves

### Neutral
- Tool is tailored to the author's specific ADR conventions

## Validation Criteria

- Linter checks all ADRs in the repo within one second
- At least three MVP checks implemented
- JSON output is valid JSON, parseable by jq
- Tool correctly identifies at least one deliberately introduced faulty ADR

## Change Log

| Date       | Author | Change                            |
|------------|--------|-----------------------------------|
| 2026-05-15 | Claude | Initial acceptance, AI-generated  |
