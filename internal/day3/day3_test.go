package day3

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	entries := ReadFromFile("day3_input.txt")

	t.Run("Should fetch correct result part 1", func(t *testing.T) {
		t.Run("The tree pattern repeats infinitely along the X axis", func(t *testing.T) {
			got := RideSlope(entries)
			assert.Equal(t, []int{94, 214, 99, 91, 46}, got)
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
