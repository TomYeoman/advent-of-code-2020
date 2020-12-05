package day4

import "fmt"

func BinaryCheck(data, upperCharacter, lowerCharacter string, baseRows, requiredIterations int) int {
	higher := baseRows
	lower := 0

	for i := 0; i < requiredIterations; i++ {
		mid := (higher + lower) / 2

		if string(data[i]) == lowerCharacter {
			higher = mid // Take lower half
		} else {
			lower = mid + 1 // Take higher half
		}
	}

	return lower
}

func FetchSeatInformation(binaryPartition string) (row int, col int, seatNum int) {

	rowNum := BinaryCheck(binaryPartition[0:7], "B", "F", 127, 7)
	columnNum := BinaryCheck(binaryPartition[7:10], "R", "L", 7, 3)

	return rowNum, columnNum, rowNum*8 + columnNum
}

// Part 1
func FindHighestSeat(entries []string) int {

	highest := 0
	for _, seat := range entries {
		_, _, seatNum := FetchSeatInformation(seat)
		if seatNum > highest {
			highest = seatNum
		}
	}

	return highest
}

// Part 2
func FindMissingSeat(entries []string) int {

	filledSeats := map[int]string{}

	// Populate map, with all current occupied seats
	for _, seat := range entries {
		_, _, seatNum := FetchSeatInformation(seat)
		filledSeats[seatNum] = "FOUND"
	}

	// Loop all seats availablle, checking whether there's a missing seat with neighbours both sides
	for i := 0; i < FindHighestSeat(entries); i++ {
		_, currSeatFound := filledSeats[i]
		_, neighbourOneFound := filledSeats[i-1]
		_, neighbourTwoFound := filledSeats[i+1]

		if (!currSeatFound) && neighbourOneFound && neighbourTwoFound {
			fmt.Printf("Was unable to find the seat %d \n", i)
			return i
		}
	}

	return 0
}
