package scanner

import (
	"github.com/spf13/cobra"

	"github.com/riggerthegeek/license-disco/packages"
<<<<<<< HEAD
	"errors"
	"path/filepath"
	"os"
	"fmt"
)

func Scan(pkgs []packages.Package, cmd *cobra.Command, path string) error {
	rootDir, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	stat, err := os.Stat(rootDir)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return errors.New(path + " is not a directory")
	}

	var hasActive = false

	for _, pkg := range pkgs {
		if pkg.Enabled(cmd) {
			if hasActive == false {
				hasActive = true
			}

			err, packages := pkg.Scan(rootDir)
			if err != nil {
				return err
			}

			fmt.Println(packages)
		}
	}

	if hasActive == false {
		return errors.New("NO_SCANNERS_ENABLED")
	}
=======
	"fmt"
	"errors"
)

func Scan(pkgs []packages.Package, cmd *cobra.Command, path string) error {
	var activePkgs = []packages.Package{}

	for _, pkg := range pkgs {
		if pkg.Enabled(cmd) {
			activePkgs = append(activePkgs, pkg)
		}
	}

	if len(activePkgs) == 0 {
		return errors.New("no scanners enabled")
	}

	fmt.Println(activePkgs)
>>>>>>> parent of 8fa3794... Done a few more things

	return nil
}