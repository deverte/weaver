package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Weaver",
	Short: "Weaver package manager.",
	Long: "The Weaver program is a package manager for symlink programs " +
		"(variation of portable applications).",
	Args: cobra.MinimumNArgs(1),
}

// Execute executes specified command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
