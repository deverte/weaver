/*
Package tangle implements commands connected with tangle management.
*/
package tangle

import (
	"fmt"

	"github.com/deverte/weaver/internal/fs"
	"github.com/spf13/cobra"
)

// TangleListCmd represents "tangle list" command.
var TangleListCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all added tangles.",
	Long:  "List of all added tangles.",
	Run: func(cmd *cobra.Command, args []string) {
		weaverfs := fs.NewWeaverFS()
		weaverfs.Fill()

		for _, tangle := range weaverfs.Home.Tangle {
			fmt.Printf("%s\t|\t%s\n", tangle.Name, tangle.Path)
		}
	},
}
