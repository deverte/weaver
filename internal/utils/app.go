/*
Package utils implements uncategorized helper functions.
*/
package utils

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/deverte/weaver/internal/fs"
	"github.com/deverte/weaver/internal/manifest"
)

// FindApp retrieves information about added application.
func FindApp(name string) (manifest.AppManifest, fs.App, fs.Tangle) {
	appManifest := manifest.NewAppManifest()
	app := fs.App{}
	tangle := fs.Tangle{}

	weaverfs := fs.NewWeaverFS()
	weaverfs.Fill()

	for _, tangleTmp := range weaverfs.Home.Tangle {
		for _, appTmp := range weaverfs.Home.Apps {
			if name == appTmp.Name {
				manifestPath := filepath.Join(
					tangleTmp.Path, "fiber", appTmp.Name+".yaml",
				)
				if _, err := os.Stat(manifestPath); !os.IsNotExist(err) {
					appManifest.Read(manifestPath)
					app = appTmp
					tangle = tangleTmp
				}
				break
			}
		}
	}
	return appManifest, app, tangle
}

// ExpandPath makes absolute path from rootPath and relativePath.
// rootPath must be without "/" sign at the end.
// relativePath must be in format "./some/relative/path".
func ExpandPath(rootPath string, relativePath string) string {
	// !!! Add error checking
	return filepath.Join(rootPath, relativePath[1:])
}

// RunScript ...
func RunScript(scriptPath string) {
	if _, err := os.Stat(scriptPath); !os.IsNotExist(err) {
		if filepath.Ext(scriptPath) == ".ps1" {
			powershellCmd := exec.Command(
				"powershell", scriptPath,
			)

			err := powershellCmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else if filepath.Ext(scriptPath) == ".cmd" || filepath.Ext(scriptPath) == ".bat" {
			cmdCmd := exec.Command(
				"cmd", "/c", scriptPath,
			)

			err := cmdCmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(
				"\"" + filepath.Ext(scriptPath) + "\" script is not supported.",
			)
		}
	}
}
