package packages

import (
	"github.com/spf13/cobra"
)

type Package interface {
	Enabled(cmd *cobra.Command) bool
	Flags() []Flag
}

type Flag struct {
	Type string
	Name string
	Value interface{}
	Usage string
}

func LoadPackages() []Package {
	return []Package{
		&NPM{},
	}
}

func (t *Flag) Twat()  {

}