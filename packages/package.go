package packages

import (
	"github.com/spf13/cobra"
)

type Package interface {
	Enabled(cmd *cobra.Command) bool
<<<<<<< HEAD
	Flags() []common.Flag
	Name() string
	Scan(path string) (error, []common.Dependency)
=======
	Flags() []Flag
}

type Flag struct {
	Type string
	Name string
	Value interface{}
	Usage string
>>>>>>> parent of 8fa3794... Done a few more things
}

func LoadPackages() []Package {
	return []Package{
		&NPM{},
	}
}

func (t *Flag) Twat()  {

}