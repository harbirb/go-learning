package main

import (
	"fmt"
)

// can group constants in a block, declare "const" just once
const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	spanish            = "Spanish"
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
)

// Functions starting with capital letter are public
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// (prefix string) makes a named return value, creates variable prefix for this function
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// dont have to explicitly say "return prefix" since named return is defined above
	return
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
