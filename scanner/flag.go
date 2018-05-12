package scanner

type FlagElement struct {
	Name string
	Type string `json:"type",xml:"type"`
	Value interface{}
	Usage string
}