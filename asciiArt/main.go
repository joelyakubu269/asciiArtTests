package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run . \"text\"")
	}
	input := os.Args[1]
	text := strings.ReplaceAll(input, `\n`, "\n")
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := generateArt(text, banner)
	fmt.Print(result)

}
