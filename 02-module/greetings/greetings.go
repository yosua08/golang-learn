package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello function returns greeting for the name person
func Hello(name string) (string, error) {
	// If no name was given
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greetings
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// func Hellos(name []string) (map[string]string, error) {
// 	message := make(map[string]string)
// 	for
// }

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
