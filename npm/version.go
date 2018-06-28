package npm

type Version struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Version string `json:"version"`
	License string `json:"license"`
	Dependencies map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}