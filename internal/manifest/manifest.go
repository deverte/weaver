/*
Package manifest implements parsing and processing of apllications manifests,
including substitution of environment variables.
*/
package manifest

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Symlink represents symlink structure.
type Symlink struct {
	Source string
	Target string
}

// AppManifest represents manifest of the symlink application.
// !!! Add reboot flag
type AppManifest struct {
	Name    string
	Version string
	Install struct {
		Symlinks []Symlink
		Copy     []struct {
			Source string
			Target string
		}
		Reg []struct {
			Path string
		}
		Scripts []struct {
			Path string
		}
	}
	Uninstall struct {
		Symlinks []struct {
			Target string
		}
		Delete []struct {
			Path string
		}
		Reg []struct {
			Path string
		}
		Scripts []struct {
			Path string
		}
	}
}

// AppManifestInterface is an AppManifest interface.
type AppManifestInterface interface {
	Read(path string) AppManifest
	Parse() AppManifest
}

// NewAppManifest is a constructor for AppManifest.
func NewAppManifest() AppManifest {
	return AppManifest{}
}

// Read reads application manifest in specified file path.
func (app *AppManifest) Read(path string) AppManifest {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(file), &app)
	if err != nil {
		log.Fatal(err)
	}

	return *app // ability for method chaining
}

// Parse substitutes environment variables inside manifest's strings.
func (app AppManifest) Parse() AppManifest {
	for idx, symlink := range app.Install.Symlinks {
		app.Install.Symlinks[idx].Source = substituteEnv(symlink.Source)
		app.Install.Symlinks[idx].Target = substituteEnv(symlink.Target)
	}
	for idx, symlink := range app.Install.Copy {
		app.Install.Copy[idx].Source = substituteEnv(symlink.Source)
		app.Install.Copy[idx].Target = substituteEnv(symlink.Target)
	}
	for idx, symlink := range app.Install.Reg {
		app.Install.Reg[idx].Path = substituteEnv(symlink.Path)
	}
	for idx, symlink := range app.Install.Scripts {
		app.Install.Scripts[idx].Path = substituteEnv(symlink.Path)
	}
	for idx, symlink := range app.Uninstall.Symlinks {
		app.Uninstall.Symlinks[idx].Target = substituteEnv(symlink.Target)
	}
	for idx, symlink := range app.Uninstall.Delete {
		app.Uninstall.Delete[idx].Path = substituteEnv(symlink.Path)
	}
	for idx, symlink := range app.Uninstall.Reg {
		app.Uninstall.Reg[idx].Path = substituteEnv(symlink.Path)
	}
	for idx, symlink := range app.Uninstall.Scripts {
		app.Uninstall.Scripts[idx].Path = substituteEnv(symlink.Path)
	}
	return app // ability for method chaining
}

// substituteEnv is a helper function for ParseManifest. It substitutes
// environment variables values into secified string.
func substituteEnv(path string) string {
	variables := findVariables(path)
	newPath := path
	for _, variable := range variables {
		if os.Getenv(variable) != "" {
			newPath = strings.Replace(
				newPath, "${"+variable+"}", os.Getenv(variable), -1,
			)
		}
	}
	return newPath
}

// findVariables is a helper function for substituteEnv. It searches variable
// pattern `${variable}` (where `variable` can be custom) in specified string
// and returns list of variables in format [`${variable1}`, ...].
func findVariables(path string) []string {
	var variables []string
	var variable []rune
	isAmp := false
	isVariable := false
	for _, char := range path {
		if char == '$' {
			isAmp = true
		}

		if isVariable && char == '}' {
			isVariable = false
			variables = append(variables, string(variable))
			variable = nil
		}

		if isVariable {
			variable = append(variable, char)
		}

		if isAmp && char == '{' {
			isVariable = true
		}

		if char != '$' {
			isAmp = false
		}
	}
	return variables
}
