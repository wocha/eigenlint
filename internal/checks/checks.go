// Package checks defines the Check interface and registers individual
// check implementations.
//
// Each check is a small, focused function that takes a parsed ADR and
// returns a CheckResult describing whether the ADR passed and which
// issues were found.
//
// Checks should be independent of each other and side-effect-free.
package checks
