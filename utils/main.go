package utils

import (
	"os"
	"strings"
)

func FetchInput(fn string) (string, error) {

	content, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}
	var input string = string(content)
	input = strings.Trim(input, "\n")
	return input, nil
}
