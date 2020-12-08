package day8

import (
	"fmt"
	"strconv"
	"strings"
)

func RunProgram(input []string) int {
	instructionsRan := make(map[int]struct{})
	currLine := 0
	acc := 0
	isEndReached := false

	for !isEndReached {

		currInstruction := strings.Split(input[currLine], " ")
		operation := currInstruction[0]
		arg, _ := strconv.Atoi(currInstruction[1])

		fmt.Printf("Running op %q, arg %d, line %d \n", operation, arg, currLine)

		switch operation {
		case "nop":
			currLine += 1
		case "acc":
			acc += arg
			currLine += 1
		case "jmp":
			currLine += arg

		_, isFound := instructionsRan[currLine]
		isEndReached = isFound

		instructionsRan[currLine] = struct{}{}

	}

	return acc
}
