package day1

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	entries := ReadFromFile("day2_input.txt")

	t.Run("Should fetch correct result part 1", func(t *testing.T) {
		got := GetResult(entries, 1)
		assert.Equal(t, 439, got)
	})

	t.Run("Should fetch correct result part 2", func(t *testing.T) {
		got := GetResult(entries, 2)
		assert.Equal(t, 584, got)
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
