package util

import "strings"

func Check(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}
