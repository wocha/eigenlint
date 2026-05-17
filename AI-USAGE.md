# AI Usage Disclosure

This repository is developed in an AI-augmented engineering practice. This file documents transparently which artifacts were created in which way.

## Disclosure Levels

Every artifact in this repository falls into one of the following categories:

### Level 1: Human-Authored
Fully written by the author without AI assistance. AI may have been used at most for spell-check or grammar correction, but no substantive contributions.

### Level 2: Human-Directed AI Generation
The author provided detailed specification and direction; AI handled the wording. The author read, reviewed, and consciously accepted the result. Substantive content is from the author, linguistic form from AI.

### Level 3: AI-Drafted, Human-Edited
AI generated a first draft based on rough direction. The author substantially revised the draft, adjusted content, and corrected inaccuracies.

### Level 4: AI-Translated
Author wrote in one language; AI translated to another. Content fully from author, language from AI.

### Level 5: AI-Refined
Author wrote a complete draft; AI proposed improvements to style, clarity, or structure. The author individually accepted or rejected each suggestion.

## Per-Artifact Assignment

### Architecture Decision Records

| File | Level | Notes |
|------|-------|-------|
| ADR-001-adr-linter-tool-design.md | Level 2 (Human-Directed AI Generation) | AI-generated based on author specification. Deliberately not written by author so the author can test the ADR from an applier's perspective. |

### Code

| File | Level | Notes |
|------|-------|-------|
| cmd/eigenlint/main.go (initial skeleton) | Level 2 | Minimal placeholder generated to validate toolchain |
| All other Go code | Level 1 (Human-Authored) | Author writes the actual implementation as a Go learning project |

## Methodological Background

The separation into levels follows disclosure frameworks from academia (Weaver 2024, Suchikova 2026), adapted for engineering repositories.

The goal is not to justify or hide AI usage, but to make transparent what was created how. This matters for audit traceability and for application contexts where the question "how do you work with AI" arises.

## Author Reflection

The author works in a Lead Architect role with AI augmentation. The substance of architectural decisions comes from the author, based on 13 years of experience in regulated IT environments. AI is used as an implementation layer, comparable to a team of implementers receiving architectural direction.

The discipline of "when AI, when self" follows a simple principle: substantial methodological decisions are written by the author, so they can be defended under pressure (e.g., in interviews). Implementation artifacts and code may be AI-augmented as long as their conformance against author-formulated constraints is verifiable.

In this specific project, ADR-001 is AI-generated as a test of the methodology — the author tests whether an ADR in this format suffices to guide implementation without the author present. The implementation itself (Go code) is written by the author.
