package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		t.Run("Should find the correct amount of 'Yes' questions per group", func(t *testing.T) {

			ttData := []struct {
				data        []string
				expectedYes int
			}{
				{data: []string{"abc", ""}, expectedYes: 3},
				{data: []string{"a", "b", "c", ""}, expectedYes: 0},
				{data: []string{"ab", "ac", ""}, expectedYes: 1},
				{data: []string{"a", "a", "a", "a", ""}, expectedYes: 1},
				{data: []string{"b", ""}, expectedYes: 1},
			}

			for i, d := range ttData {
				res := GetSumOfTotalsAcrossGroups(d.data)

				assert.Equal(t, d.expectedYes, res, fmt.Sprintf("Incorrect yes answer found for run %v", i))
			}
		})
		t.Run("Should return correct result for test input", func(t *testing.T) {

			invalidPassports := ReadFromFile("fixtures/day6_input.txt")

			got := GetSumOfTotalsAcrossGroups(invalidPassports)
			assert.Equal(t, 3640, got)

		})
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
