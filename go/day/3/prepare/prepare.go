package main

import (
	_ "embed"
	"os"
	"regexp"
)

//go:embed input.txt
var input string

func main() {
	const output = "prepared.txt"

	file, err := os.OpenFile(output, os.O_CREATE, 644)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	matches := re.FindAllString(input, -1)

	for _, m := range matches {
		file.WriteString(m + "\n")
	}
}
