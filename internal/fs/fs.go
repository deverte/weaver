/*
Package fs (folder structure) implements navigation inside Weaver directories
and installed packages.
*/
package fs

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// App represents file structure of an application.
type App struct {
	Name string
	Path string
}

// Tangle represent file structure of a tangle.
type Tangle struct {
	Name     string
	Path     string
	Manifest string
}

// WeaverFS represents Weaver folder structure.
type WeaverFS struct {
	Home struct {
		Path   string
		Apps   []App
		Tangle []Tangle
	}
	External struct {
		Targets struct {
			Name string
			Path string
		}
		Sources struct {
			Name string
			Path string
		}
	}
}

// WeaverFSInterface is a WeaverFS interface.
type WeaverFSInterface interface {
	Fill()
}

// NewWeaverFS is a WeaverFS constructor.
func NewWeaverFS() WeaverFS {
	return WeaverFS{}
}

// Fill fills WeaverFS with directories information.
func (weaverfs *WeaverFS) Fill() {
	weaverfs.Home.Path = os.Getenv(`WEAVER_HOME`)
	if len(weaverfs.Home.Path) == 0 {
		log.Fatal("Error: WEAVER_HOME environment variable does not specified.")
	}

	// Fill Tangle
	weaverTanglePath := filepath.Join(weaverfs.Home.Path, "tangle")
	files, err := ioutil.ReadDir(weaverTanglePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, repo := range files {
		weaverfs.Home.Tangle = append(
			weaverfs.Home.Tangle,
			struct {
				Name     string
				Path     string
				Manifest string
			}{
				Name:     repo.Name(),
				Path:     filepath.Join(weaverTanglePath, repo.Name()),
				Manifest: filepath.Join(weaverTanglePath, repo.Name()),
			},
		)
	}

	// Fill Apps
	weaverAppsPath := filepath.Join(weaverfs.Home.Path, "apps")
	files, err = ioutil.ReadDir(weaverAppsPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, repo := range files {
		weaverfs.Home.Apps = append(
			weaverfs.Home.Apps,
			struct {
				Name string
				Path string
			}{
				Name: repo.Name(),
				Path: filepath.Join(weaverAppsPath, repo.Name()),
			},
		)
	}
}
