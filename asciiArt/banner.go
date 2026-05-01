package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	m := make(map[rune][]string)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	text := string(data)
	lines := strings.Split(strings.ReplaceAll(text, "\r", ""), "\n")

	for i := 32; i <= 126; i++ { //trying to take i back to a new indexing system starting from 0
		index := i - 32
		start := index * 9

		if start+8 > len(lines) {
			return nil, fmt.Errorf("invalid banner format")
		}
		charlines := lines[start : start+8]
		m[rune(i)] = charlines // why i is used not index is because when want the cahracter, the key does not deal with position or indexing just the values
	}
	return m, nil

}
