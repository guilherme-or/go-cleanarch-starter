package util

import (
	"fmt"
	"strings"
)

func Say(msg string, args ...any) {
	if _, err := fmt.Printf(msg, args...); err != nil {
		panic(err)
	}
}

func Ask(msg string, args ...any) string {
	var input string = ""

	msg += " "
	Say(msg, args...)
	if _, err := fmt.Scanln(&input); err != nil {
		if err.Error() != "unexpected newline" {
			panic(err)
		} else {
			input = ""
		}
	}

	return strings.TrimSpace(input)
}

func AskBool(msg string, args ...any) bool {
	msg += " (Y/n) "
	input := strings.ToLower(Ask(msg, args...))

	if input == "" || input == "y" {
		return true
	} else if input == "n" {
		return false
	} else {
		panic("Invalid input. Please type 'y' or 'n'")
	}
}
