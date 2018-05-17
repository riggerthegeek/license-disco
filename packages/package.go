package packages

import (
	"github.com/spf13/cobra"
	"github.com/riggerthegeek/license-disco/npm"
	"github.com/riggerthegeek/license-disco/common"
)

type Package interface {
	Enabled(cmd *cobra.Command) bool
	Flags() []common.Flag
	Scan(path string) (error, *string)
}

func LoadPackages() []Package {
	return []Package{
		&npm.NPM{},
	}
}