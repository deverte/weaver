package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Weaver's version.",
	Long:  "Print the version number of the Weaver.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Weaver v0.1")
	},
}
