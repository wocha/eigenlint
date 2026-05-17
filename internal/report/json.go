// json.go implements a Reporter that produces machine-readable JSON
// output suitable for further processing with tools like jq.
//
// The JSON structure contains a summary block and a per-ADR results
// array, where each ADR has nested check results and issues.
package report
