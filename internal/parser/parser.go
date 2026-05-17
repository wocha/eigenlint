// Package parser reads ADR markdown files from disk and extracts
// their structural elements into typed ADR values.
//
// The parser is intentionally permissive: it does not validate
// correctness, only extracts what it finds. Validation is the
// responsibility of the checks package.
package parser
