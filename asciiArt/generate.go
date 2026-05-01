package main

import (
	"strings"
)

func generateArt(text string, banner map[rune][]string) string {
	if text == "" {
		return ""
	}
	
	texts := Split(text)
	var result strings.Builder
	for i, r := range texts {
		if r == "" {
			if i < len(texts)-1 {
				result.WriteString("\n")
			}
		} else {
			rows := Render(r, banner)
			for _, row := range rows {
				result.WriteString(row + "\n") // helps to move to the next line
			}
		}

	}
	return result.String()
}
