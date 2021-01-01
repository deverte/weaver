package cmd

import (
	"fmt"

	"github.com/deverte/weaver/internal/fs"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all added applications.",
	Long:  "List of all added applications with installation status.",
	Run: func(cmd *cobra.Command, args []string) {
		weaverfs := fs.NewWeaverFS()
		weaverfs.Fill()

		for _, app := range weaverfs.Home.Apps {
			fmt.Printf("%s\t|\t%s\n", app.Name, app.Path)
		}
	},
}
