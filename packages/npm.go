package packages

import (
	"github.com/spf13/cobra"
)

type NPM struct {
	
}

func (obj *NPM) Enabled(cmd *cobra.Command) bool  {
	return cmd.Flags().Lookup("detect.npm").Value.String() == "true"
}

func (obj *NPM) Flags() []Flag {
	return []Flag{
		{
			Type: "Bool",
			Name: "detect.npm",
			Value: false,
			Usage: "Search for npm dependencies",
		},
		{
			Type: "Bool",
			Name: "detect.npm2",
			Value: false,
			Usage: "Search for npm dependencies",
		},
	}
}