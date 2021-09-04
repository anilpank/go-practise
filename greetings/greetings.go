package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil

}

// returns a greeting for named person
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	//returns the greeting that embeds name in a message
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprintf(randomFormat())
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A Slice of message formats
	formats := []string{
		"Hi %v Welcome",
		"Great to see you %v",
		"Hail %v, Well Met",
	}

	return formats[rand.Intn(len(formats))]
}
