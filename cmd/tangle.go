package cmd

import (
	"github.com/deverte/weaver/cmd/tangle"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tangleCmd)
	tangleCmd.AddCommand(tangle.TangleListCmd)
}

var tangleCmd = &cobra.Command{
	Use:   "tangle",
	Short: "Tangle manipulation.",
	Long:  "Tangle manipulation.",
}
