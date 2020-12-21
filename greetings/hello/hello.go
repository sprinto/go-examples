package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("MM consultants")
	fmt.Println(message)
}
