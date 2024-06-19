package main

import (
	"fmt"

	"productionwebsite.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Guilherme")
	fmt.Println(message)
}
