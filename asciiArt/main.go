package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func printAscii(str string, ascii map[rune][]string) {
	str = strings.ReplaceAll(str, `\n`, "\n")
	strs := strings.Split(str, "\n")
	for _, r := range strs {
		for i := 0; i < 8; i++ {
			for _, c := range r {
				if val, ok := ascii[c]; ok {
					fmt.Print(val[i])
				} else {
					fmt.Print("?") // for easy debugging
				}
			}
			fmt.Println() //   - for each row, print that row for every character in r
			//   - after completing a row, move to the next line
		}

	}
}
func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run . \"text\"")
	}
	input := os.Args[1]
	ascii, err := LoadBanner("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	printAscii(input, ascii)

}
