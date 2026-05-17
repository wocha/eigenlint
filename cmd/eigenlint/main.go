package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eigenlint",
	Short: "Linter for Architecture Decision Records",
	Long:  "EigenLint checks ADR files for structural completeness, status validity, and changelog hygiene.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("EigenLint v0.1.0 — placeholder, no checks implemented yet")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
