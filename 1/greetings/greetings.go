package greetings

import "fmt"

//? := is the shortcut to declaring and initializing a variable in go.
//? Go uses the value on its right to determine the variable's type. But instead, we could have used:
//! var message string
//! message = fmt.Sprintf("Hi, %v. Welcome!", name)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
