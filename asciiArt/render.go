package main

import (
	"strings"
)

func render(text string, banner map[rune][]string) []string {
	result := make([]string, 8)
	for i := 0; i < 8; i++ {
		var row strings.Builder
		for _, c := range text {
			if val, ok := banner[c]; ok {
				row.WriteString(val[i])
			}
		}
		result[i] = row.String() // builds value row by row and stores them
	}
	return result

}
