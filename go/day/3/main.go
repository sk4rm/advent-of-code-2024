package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed prepared.txt
var input string

func main() {
	instructions := strings.Split(input, "\n")
	ans := solve(instructions)
	fmt.Println(ans)
}

func solve(instructions []string) int {
	shouldMul := true
	sum := 0

	for _, inst := range instructions {
		switch inst {

		case "do()":
			shouldMul = true

		case "don't()":
			shouldMul = false

		default:
			if shouldMul {
				sum += mustMultiply(inst)
			}
		}
	}

	return sum
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
