package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined log
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	// Get greetings message and print it
	message, err := greetings.Hello("Yosua")

	// If error was returned
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
