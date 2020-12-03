package day3

import "fmt"

type tobogganer struct {
	x             int
	y             int
	right         int
	down          int
	treesCollided int
}

type slope struct {
	grid         [][]string
	tobogganists []tobogganer
}

// rideDown returns the path taken
func (s *slope) rideDown() []int {

	var result []int

	// Run every tobogannism down the tracks
	for i, _ := range s.tobogganists {
		// Whilst there's still a path to travel down
		for s.tobogganists[i].y < len(s.grid)-1 {
			newX := s.tobogganists[i].x + s.tobogganists[i].right
			newY := s.tobogganists[i].y + s.tobogganists[i].down

			// CHeck whether our move, would exceed row length, if so dynamically re-size
			if newX >= len(s.grid[newY]) {
				s.resizeRowPath(newY)
			}

			s.tobogganists[i].x = newX
			s.tobogganists[i].y = newY

			// Get value from grid
			if s.grid[newY][newX] == "#" {
				s.tobogganists[i].treesCollided++
			}
		}

		fmt.Print("Completed Check")
		result = append(result, s.tobogganists[i].treesCollided)
	}

	return result

}

func (s *slope) resizeRowPath(row int) {

	resizeCount := row / 5
	dataToDupe := s.grid[row]

	// Re-size by just enough, probably a little overkill here but works :)
	for x := 0; x < resizeCount; x++ {
		s.grid[row] = append(s.grid[row], dataToDupe...)
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

func RideSlope(entries []string) []int {
	slope := &slope{
		tobogganists: []tobogganer{
			{
				x: 0, y: 0, right: 1, down: 1, treesCollided: 0,
			},
			{
				x: 0, y: 0, right: 3, down: 1, treesCollided: 0,
			},
			{
				x: 0, y: 0, right: 5, down: 1, treesCollided: 0,
			},
			{
				x: 0, y: 0, right: 7, down: 1, treesCollided: 0,
			},
			{
				x: 0, y: 0, right: 1, down: 2, treesCollided: 0,
			},
		},
		grid: [][]string{},
	}

	slope.generateTerrain(entries)

	results := slope.rideDown()

	return results
}
