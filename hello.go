package main

import "fmt" // Importing a package that contains the Println funtion

func Hello(name string) string { // string -> means the func returns string value
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello("Chris"))
}
