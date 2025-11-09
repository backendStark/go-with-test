package fuzzing

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	username string
	age      int
}

func ParseUserInput(input string) (*User, error) {
	res := strings.Split(input, ":")

	if len(res) != 2 {
		return nil, fmt.Errorf("error parsing user input, expected two parts")
	}

	age, err := strconv.Atoi(res[1])

	if err != nil {
		return nil, fmt.Errorf("error converting age to int")
	}

	if age <= 0 {
		return nil, fmt.Errorf("error parsing user input, age should be more then 0")
	}

	trimmedUsername := strings.TrimSpace(res[0])

	if trimmedUsername == "" {
		return nil, fmt.Errorf("error parsing user input, username is empty")
	}

	return &User{
		username: trimmedUsername,
		age:      age,
	}, nil
}
