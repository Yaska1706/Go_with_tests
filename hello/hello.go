package hello

const EnglishHelloPrefix = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return EnglishHelloPrefix + ", " + name
}
