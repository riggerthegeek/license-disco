package packages

type Package interface {
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