package main

import (
	"fmt"
)

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
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
