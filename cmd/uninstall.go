package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/deverte/weaver/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [application name]",
	Short: "Uninstall application.",
	Long: "Uninstall application, with ability " +
		"to restore it with all user data.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appManifest, app, _ := utils.FindApp(args[0])

		if app.Name != "" {
			appManifest.Parse()

			// Pre-uninstall scripts
			for _, script := range appManifest.Uninstall.Scripts {
				if script.Type == "pre" {
					scriptPath := utils.ExpandPath(app.Path, script.Path)
					utils.RunScript(scriptPath)
				}
			}
			// Symlink
			for _, symlink := range appManifest.Uninstall.Symlinks {
				os.Remove(symlink.Target)
			}
			// Registry
			for _, reg := range appManifest.Uninstall.Reg {
				regCmd := exec.Command(
					"reg", "import", utils.ExpandPath(app.Path, reg.Path),
				)

				err := regCmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			}
			// Delete
			for _, file := range appManifest.Uninstall.Delete {
				err := os.RemoveAll(file.Path)
				if err != nil {
					log.Fatal(err)
				}
			}
			// Post-uninstall scripts
			for _, script := range appManifest.Uninstall.Scripts {
				if script.Type == "post" || script.Type == "" {
					scriptPath := utils.ExpandPath(app.Path, script.Path)
					utils.RunScript(scriptPath)
				}
			}

			fmt.Println(
				"Application \"" + app.Name + "\" successfully uninstalled.",
			)
		} else {
			log.Fatal("App \"" + args[0] + "\" is not in any Tangle.")
		}
	},
}
