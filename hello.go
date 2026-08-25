package main

// import "fmt" // Importing a package that contains the Println funtion

const (
	spanish = "Spanish"
	french  = "French"
	bengali = "Bengali"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	bengaliHelloprefix = "Nomoskar, "

	spanishEndingSign = ":)"
	englishEndingSign = "!"
	frenchEndingSign  = "^-^"
	bengaliEndingSign = "^<>^"
)

// Base function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(name, language)
}

// Helper function
func greetingPrefix(name string, language string) string {

	prefix := englishHelloPrefix
	endingSign := englishEndingSign

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
		endingSign = spanishEndingSign
	case french:
		prefix = frenchHelloPrefix
		endingSign = frenchEndingSign
	case bengali:
		prefix = bengaliHelloprefix
		endingSign = bengaliEndingSign
	}

	return prefix + name + endingSign
}
