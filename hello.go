package main

import "fmt" // Importing a package that contains the Println funtion

const englishHelloPrefix = "Hello, " // Declaring a constant value

func Hello(name string) string { // string -> means the func returns string value
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Chris"))
}
