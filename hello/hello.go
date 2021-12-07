package hello

const (
	Spanish = "Spanish"
	French  = "French"

	EnglishHelloPrefix = "Hello"
	SpanishHelloPrefix = "Hola"
	FrenchHelloPrefix  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return Greeting(language) + ", " + name
}

func Greeting(language string) (prefix string) {

	switch language {
	case French:
		prefix = FrenchHelloPrefix
	case Spanish:
		prefix = SpanishHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return prefix
}
