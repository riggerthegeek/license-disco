package packages

type NPM struct {
	
}

func (obj *NPM) Flags() []Flag {
	return []Flag{
		{
			Type: "Bool",
			Name: "detect.npm",
			Value: false,
			Usage: "Search for npm dependencies",
		},
	}
}