package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/1set/gut/yos"
	"github.com/deverte/weaver/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install [application name]",
	Short: "Install application.",
	Long:  "Install application.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appManifest, app, _ := utils.FindApp(args[0])
		appManifest.Parse()

		// Symlink
		for _, symlink := range appManifest.Install.Symlinks {
			// !!! Add recursive directory creation if dir is symlinking
			// into unexisting directory
			os.Symlink(utils.ExpandPath(app.Path, symlink.Source), symlink.Target)
		}
		// Registry
		for _, reg := range appManifest.Install.Reg {
			// !!! Make one function for install and uninstall
			regCmd := exec.Command(
				"reg", "import", utils.ExpandPath(app.Path, reg.Path),
			)

			err := regCmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}
		// Copy
		for _, file := range appManifest.Install.Copy {
			// !!! Add check if target exists.
			fileStat, err := os.Stat(utils.ExpandPath(app.Path, file.Source))
			if err != nil {
				log.Fatal(err)
			}

			// !!! Add recursive directory creation if file/dir is copying into
			// unexisting directory
			if fileStat.IsDir() {
				err = yos.CopyDir(
					utils.ExpandPath(app.Path, file.Source), file.Target,
				)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				err = yos.CopyFile(
					utils.ExpandPath(app.Path, file.Source), file.Target,
				)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		// Scripts
		for _, script := range appManifest.Install.Scripts {
			// !!! Make one function for install and uninstall
			// !!! Add support for another scripts (.cmd, .bat and etc.)
			// !!! Add pre- and post-install scripts
			scriptPath := utils.ExpandPath(app.Path, script.Path)
			if _, err := os.Stat(scriptPath); !os.IsNotExist(err) {
				if filepath.Ext(scriptPath) == ".ps1" {
					powershellCmd := exec.Command(
						"powershell", scriptPath,
					)

					err := powershellCmd.Run()
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
		// !!! Add installation complete.
	},
}
