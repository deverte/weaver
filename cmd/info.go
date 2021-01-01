/*
Package cmd implements base commands of the Weaver application.
*/
package cmd

import (
	"fmt"

	"github.com/deverte/weaver/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info [application]",
	Short: "Information about specified application.",
	Long:  "Information about specified application.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appManifest, app, tangle := utils.FindApp(args[0])

		fmt.Printf("Name: %v\n", appManifest.Name)
		fmt.Printf("Version: %v\n", appManifest.Version)
		fmt.Printf("Path: %v\n", app.Path)
		fmt.Printf("Tangle: %v\n", tangle.Path)
	},
}
