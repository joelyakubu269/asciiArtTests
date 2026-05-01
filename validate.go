package main

import (
	"fmt"
)

func validate(text string) (rune, error) {
	for _, r := range text {
		if r < 32 || r > 126 {
			return r, fmt.Errorf("invalid character") // check for alphanumeric characters and space
		}
	}
	return 0, nil
}
