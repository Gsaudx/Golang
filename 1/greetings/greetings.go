package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Empty name")
	}

	// If a name was received, return a value that embeds the name
	// In a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
	//!nil in this case means no errors
}

//? := is the shortcut to declaring and initializing a variable in go.
//? Go uses the value on its right to determine the variable's type. But instead, we could have used:
//! var message string
//! message = fmt.Sprintf("Hi, %v. Welcome!", name)
