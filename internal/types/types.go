// Package types defines the shared data structures used across all
// EigenLint packages: parsed ADR representations, lint issues, and
// check results.
//
// These types form the contract between parser, checks, and reporters:
// the parser produces ADR values, checks consume ADR values and produce
// CheckResult values, reporters consume CheckResult values.
package types
