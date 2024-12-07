package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input/day-3.txt
var input string

func main() {

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches := re.FindAllString(input, -1)

	fmt.Println(solve(matches))
}

func mustMultiply(instruction string) int {
	trimmed := instruction[4 : len(instruction)-1]
	operands := strings.Split(trimmed, ",")

	op1, err := strconv.Atoi(operands[0])
	if err != nil {
		panic(err)
	}
	op2, err := strconv.Atoi(operands[1])
	if err != nil {
		panic(err)
	}

	return op1 * op2
}

func solve(instructions []string) int {
	sum := 0
	for _, inst := range instructions {
		sum += mustMultiply(inst)
	}
	return sum
}
