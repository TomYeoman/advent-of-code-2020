package day3

import (
	"fmt"
)

type toboggan struct {
	xPos int
	yPos int
}

type slope struct {
	grid [][]string
	toboggan
}

// rideDown returns the path taken
func (s *slope) rideDown() []string {

	var result []string
	// Get current toboggan position
	for s.toboggan.yPos < len(s.grid)-1 {

		newXPosition := s.toboggan.xPos + 3
		newYPosition := s.toboggan.yPos + 1

		fmt.Printf("Toboggan at x: %v, y: %v \n", newXPosition, newYPosition)

		if newYPosition == 322 {
			fmt.Print("Test")
		}
		// CHeck whether our move, would exceed row length, if so dynamically re-size
		if newXPosition > len(s.grid[newYPosition]) {
			fmt.Printf("Resizing row %v \n", newYPosition)
			s.resizeRowPath(newYPosition)
		}

		s.toboggan.xPos = newXPosition
		s.toboggan.yPos = newYPosition

		// Get value from grid
		result = append(result, s.grid[newYPosition][newXPosition])

	}

	return result

}

func (s *slope) resizeRowPath(row int) {

	increaseBy := row / 5
	currRowData := s.grid[row]

	for x := 0; x < increaseBy; x++ {
		s.grid[row] = append(s.grid[row], currRowData...)
	}

}

// Generate terrain
func (s *slope) generateTerrain(entries []string) {

	dataGrid := make([][]string, len(entries))
	for x, row := range entries {
		dataGrid[x] = make([]string, len(row))
		for y, rowEntry := range row {
			dataGrid[x][y] = string(rowEntry)
		}
	}

	s.grid = dataGrid
}

func RideSlope(entries []string) int {

	slope := &slope{
		toboggan: toboggan{

			0, 0,
		},
		grid: [][]string{},
	}

	slope.generateTerrain(entries)
	result := slope.rideDown()

	treesHit := 0
	for _, res := range result {
		if res == "#" {
			treesHit++
		}
	}
	return treesHit
}
