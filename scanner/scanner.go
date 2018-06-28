package scanner

import (
	"github.com/spf13/cobra"

	"github.com/riggerthegeek/license-disco/packages"
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

	return nil
}