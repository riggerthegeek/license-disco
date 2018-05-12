package scanner

type Language interface {
	Foo()
}

type Scanner struct {
	Languages []Language
}

func (obj *Scanner) RegisterLang(lang Language) []Language {
	obj.Languages = append(obj.Languages, lang)

	return obj.Languages
}

func GetFlags() []FlagElement {
	return []FlagElement{
		{
			Type: "hello",
			Value: 2,
		},
		{
			Type: "hello",
			Value: true,
		},
	}
}
