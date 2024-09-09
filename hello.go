package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name
	}
	if language == "French" {
		
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}

// func main() {
//     // Set properties of the predefined Logger, including
//     // the log entry prefix and a flag to disable printing
//     // the time, source file, and line number.
//     log.SetPrefix("greetings: ")
//     log.SetFlags(0)

//     // A slice of names.
//     names := []string{"Gladys", "Samantha", "Darrin"}

//     // Request greeting messages for the names.
//     messages, err := greetings.Hellos(names)
//     if err != nil {
//         log.Fatal(err)
//     }
//     // If no error was returned, print the returned map of
//     // messages to the console.
//     fmt.Println(messages)
// }
