# ADR-001: ADR-Linter Tool Design

## Status

| Field         | Value                                    |
|---------------|------------------------------------------|
| Status        | Accepted                                 |
| Date          | 2026-05-15                               |
| Author        | Claude (AI-generated, see AI-USAGE.md)   |
| Supersedes    | -                                        |
| Superseded by | -                                        |

## Kontext und Problem

Mit wachsender ADR-Sammlung in EigenState-Core wird manuelle Konsistenz-Prüfung unzuverlässig. Pflicht-Sektionen können vergessen werden, Status-Werte können inkonsistent sein, Change-Logs können fehlen.

Zusätzlich dient dieser Linter als Pilotprojekt für die in ADR-000 dokumentierte MADR+AI-Generation-Annex-Methode. Das Tool wird benutzt, um die Methode aus Anwender-Sicht zu validieren.

## Decision Drivers

- Schnelle Ausführbarkeit für CI/CD-Integration (sub-Sekunde für 20 ADRs)
- Single-Binary-Deployment ohne Runtime-Dependencies
- Idiomatische Sprache für CLI-Tools
- Erweiterbar um neue Checks ohne Architektur-Änderung
- Lernprojekt für erste vollständige Go-Implementation des Autors
- Validierung der MADR+Annex-Methode am realen Use Case

## Considered Options

### Option A: Python-Script
Vertraut, schnell zu schreiben, aber Runtime-Dependencies und kein Lernwert.

### Option B: Bash-Script
Kein Build, aber skaliert nicht über 100 Zeilen sauber.

### Option C: Go-Binary
Lernkurve, aber idiomatisch und Single-Binary-Distribution.

## Decision Outcome

Gewählt wird Option C: Go-Binary. Lernwert und Deployment-Vorteile überwiegen initialen Setup-Aufwand.

## AI-Generation-Annex

### Constraints für Code-Generierung

- MUST: Verwende github.com/spf13/cobra für CLI-Struktur, nicht das flag-Package
- MUST: Strukturiere nach Go-Standard-Layout mit cmd/eigenlint/main.go und internal/
- MUST: Jeder Check ist eine eigene Funktion mit klarer Signatur, die ADR-Daten nimmt und Issues zurückgibt
- MUST: Output ist JSON-serialisierbar UND human-readable, Flag --format=json|text (Default: text)
- MUST: Exit-Code 0 bei Pass, 1 bei Linter-Fehlern, 2 bei Tool-internen Fehlern
- MUST: Konfigurationspfad über --path Flag, Default docs/architecture/adr
- SHOULD: Verwende einen Markdown-Parser nur wenn nötig, sonst Regex
- SHOULD: Logge auf stderr, Output auf stdout (Unix-Pipeline-tauglich)

### Verbotene Implementierungsmuster

- Keine init()-Funktionen mit Side-Effects
- Kein direktes os.Exit() außerhalb von main.go
- Keine String-Concatenation für JSON-Output
- Keine Java-Style Builder-Pattern für einfache Structs
- Keine globalen Variablen für Konfiguration
- Kein Panic auf erwartbaren Fehlern

### Machine-Readable References

- Requires: ADR-000 (ADR-Format-Konvention)
- Conflicts-With: -
- Supersedes: -
- Implementation-Hint: notes/go-lernpfad.md
- Validation-Script: scripts/validate-adr-linter.sh

## Detaillierte Festlegungen

### Welche Checks zuerst implementieren (MVP)

1. status_check: Status-Tabelle vorhanden mit Status, Date, Author
2. sections_check: Pflicht-Sektionen vorhanden (Kontext und Problem, Decision Drivers, Considered Options, Decision Outcome, Konsequenzen)
3. changelog_check: Change Log existiert oder Hinweis

### Output-Format Anforderungen

- Text-Mode: pro ADR eine Block-Ausgabe mit Check-Name, Pass/Fail-Status, ggf. Zeilennummer und Message
- JSON-Mode: Summary mit total/passed/failed, plus Array von ADR-Ergebnissen mit allen Checks und Issues
- Beide Modi müssen mit grep beziehungsweise jq weiterverarbeitbar sein

## Konsequenzen

### Positiv
- Automatisch durchsetzbare ADR-Konsistenz
- Tool als Bewerbungs-Asset
- Abgeschlossener Go-Lernkreislauf
- Methodik-Validierung am realen Use Case

### Negativ
- Lernkurve Go
- Wartungs-Last bei ADR-Format-Erweiterungen

### Neutral
- Tool ist auf eigene ADR-Konventionen zugeschnitten

## Validation Criteria

- Linter prüft alle ADRs des Repos in unter 1 Sekunde
- Mindestens drei MVP-Checks implementiert
- JSON-Output ist gültiges JSON, parsbar von jq
- Tool erkennt mindestens eine bewusst eingeführte Fehler-ADR

## Change Log

| Datum      | Author | Änderung                        |
|------------|--------|---------------------------------|
| 2026-05-15 | Claude | Initial Acceptance, AI-generiert |
