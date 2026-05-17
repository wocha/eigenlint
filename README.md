# EigenLint

A linter for Architecture Decision Records (ADRs).

EigenLint validates ADR markdown files against a strict MADR-based format. It checks for structural completeness, status validity, and changelog hygiene.

## Status

Early development. MVP in progress.

This is the first complete Go project of the author and serves as a pilot application of the MADR-with-AI-Generation-Annex methodology developed in [EigenState-Core](https://github.com/wocha/EigenState-Core).

## Installation

```bash
go install github.com/wocha/eigenlint/cmd/eigenlint@latest
```

Or build from source:

```bash
git clone https://github.com/wocha/eigenlint
cd eigenlint
make build
```

## Usage

```bash
eigenlint --path docs/architecture/adr
eigenlint --path docs/architecture/adr --format json
```

### Exit Codes

- `0` — All ADRs pass all checks
- `1` — One or more ADRs failed at least one check
- `2` — Tool-internal error (e.g., path not found)

## Checks (MVP)

| Check | Description |
|-------|-------------|
| `status_check` | Status table is present and contains required fields (Status, Date, Author) |
| `sections_check` | All mandatory MADR sections are present |
| `changelog_check` | Change Log section exists |

## Design

The tool design is documented in [ADR-001](docs/architecture/adr/ADR-001-adr-linter-tool-design.md). The ADR serves as both specification and test case for the author's MADR-with-AI-Generation-Annex methodology.

## License

MIT
