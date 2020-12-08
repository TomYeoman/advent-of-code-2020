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
		}

		_, isFound := instructionsRan[currLine]
		isEndReached = isFound

		instructionsRan[currLine] = struct{}{}

	}

	return acc
}

func TestInputTerminates(input []string) (bool, int) {
	currExecutionLine := 0
	accTotal := 0
	runCount := 0
	isEndReached := false
	hasGracefullyTerminated := false

	for !isEndReached {

		currInstruction := strings.Split(input[currExecutionLine], " ")
		operation := currInstruction[0]
		arg, _ := strconv.Atoi(currInstruction[1])

		switch operation {
		case "nop":
			currExecutionLine += 1
		case "acc":
			accTotal += arg
			currExecutionLine += 1
		case "jmp":
			currExecutionLine += arg
		}

		// Deadlock detection, if we've ran over our instruction count
		if runCount > len(input) {
			isEndReached = true
		}
		// If we're trying to access an instruction, out of program bounds we've finished
		if currExecutionLine >= len(input) {
			isEndReached = true
			hasGracefullyTerminated = true
		}

		runCount++
	}

	return hasGracefullyTerminated, accTotal
}

func RunProgramTwo(input []string) int {

	// Get array entries that contain NOOP
	for i, inp := range input {

		// If a line is nop / jmp, we'll switch it to the other and test whether program terminates
		inputCopy := make([]string, len(input))
		copy(inputCopy, input)

		currInstruction := strings.Split(inp, " ")
		operation := currInstruction[0]
		arg := currInstruction[1]

		if operation == "nop" {
			inputCopy[i] = "jmp " + arg
		}
		if operation == "jmp" {
			inputCopy[i] = "nop " + arg
		}

		hasTerminated, acc := TestInputTerminates(inputCopy)

		if hasTerminated {
			return acc
		}
	}

	return 0
}

// If we haven't
