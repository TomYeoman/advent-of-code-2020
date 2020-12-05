package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	t.Run("Should fetch the correct seat information, given a binary space partition", func(t *testing.T) {

		ttData := []struct {
			data    string
			row     int
			col     int
			seatNum int
		}{
			{data: "BFFFBBFRRR", row: 70, col: 7, seatNum: 567},
			{data: "FFFBBBFRRR", row: 14, col: 7, seatNum: 119},
			{data: "BBFFBBFRLL", row: 102, col: 4, seatNum: 820},
		}

		for i, d := range ttData {
			row, col, seatNum := FetchSeatInformation(d.data)

			assert.Equal(t, d.row, row, fmt.Sprintf("Row was incorrect for run %v", i))
			assert.Equal(t, d.col, col, fmt.Sprintf("Col was incorrect for run %v", i))
			assert.Equal(t, d.seatNum, seatNum, fmt.Sprintf("seatNum was incorrect for run %v", i))
		}
	})

	t.Run("Should find the highest seat, across all of the input data", func(t *testing.T) {
		seatData := ReadFromFile("fixtures/day5_input.txt")

		got := FindHighestSeat(seatData)
		assert.Equal(t, 965, got)
	})

	t.Run("Should find the correct missing seat", func(t *testing.T) {
		seatData := ReadFromFile("fixtures/day5_input.txt")

		got := FindMissingSeat(seatData)
		assert.Equal(t, 524, got)
	})

}

func ReadFromFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}
