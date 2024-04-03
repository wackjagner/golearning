package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	//Return error if no name given
	if name == "" {
		return "", errors.New("You forgot to provide a name.")
	}

	//Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
