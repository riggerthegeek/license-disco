package scanner

import (
	"github.com/spf13/cobra"

	"github.com/riggerthegeek/license-disco/packages"
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

	return nil
}