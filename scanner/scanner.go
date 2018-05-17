package scanner

import (
	"github.com/spf13/cobra"

	"github.com/riggerthegeek/license-disco/packages"
	"path/filepath"
)

func Scan(pkgs []packages.Package, cmd *cobra.Command, path string) error {
	//var activePkgs = []packages.Package{}

	for _, pkg := range pkgs {
		if pkg.Enabled(cmd) {
			// @todo check rootDir exists and is a directory
			rootDir, err := filepath.Abs(path)

			if err != nil {
				return err
			}

			pkg.Scan(rootDir)
		}
	}

	//if len(activePkgs) == 0 {
	//	return errors.New("no scanners enabled")
	//}

	return nil
}