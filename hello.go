package main

import "fmt" // Importing a package that contains the Println funtion

// Spanish
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, " // spanish const variable
const spanishEndingSign = ":)"

// English/Empty string
const englishHelloPrefix = "Hello, " // Declaring a constant value -> english
const englishEndingSign = "!"

// French
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const frenchEndingSign = "^-^"

func Hello(name string, language string) string { // string -> means the func returns string value
	if name == "" {
		name = "World"
	}

	// Spanish
	if language == spanish {
		return spanishHelloPrefix + name + spanishEndingSign
	}

	//French
	if language == french {
		return frenchHelloPrefix + name + frenchEndingSign
	}

	return englishHelloPrefix + name + englishEndingSign
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
