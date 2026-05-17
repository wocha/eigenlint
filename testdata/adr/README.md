# Test Fixtures

These ADR markdown files are used to verify EigenLint's check behavior. They are organized by expected outcome.

## Structure

- `valid/` — ADRs that should pass all checks
- `invalid/` — ADRs that should fail at least one check
- `warnings/` — ADRs that should produce warnings but no hard failures

## Expected Outcomes

| File | status_check | sections_check | changelog_check |
|------|--------------|----------------|-----------------|
| valid/ADR-001-use-postgres-for-state.md | PASS | PASS | PASS |
| invalid/ADR-002-missing-section.md | PASS | FAIL | PASS |
| invalid/ADR-003-broken-status.md | FAIL | PASS | PASS |
| warnings/ADR-004-no-changelog.md | PASS | PASS | WARN |
| invalid/ADR-005-invalid-status-value.md | FAIL | PASS | PASS |
| invalid/ADR-006-wrong-order.md | PASS | PASS | PASS (TBD) |
| invalid/ADR-007-empty.md | FAIL | FAIL | FAIL |
| invalid/ADR-008-typos.md | PASS | FAIL | PASS |

## Notes

- Fixture 6 (wrong section order) currently expected to pass all MVP checks. A future `section_order_check` would catch this case.
- Fixture 8 (typos) tests strictness of section name matching. Multiple typos in section headers ("Contxt" instead of "Context", "Cosidered" instead of "Considered", etc.) should be caught by `sections_check`.
