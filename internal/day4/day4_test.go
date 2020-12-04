package day4

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	t.Run("Should process all invalid passwords correctly", func(t *testing.T) {
		invalidPassports := ReadFromFile("fixtures/invalid_passports.txt")

		got := GetPasswordCount(invalidPassports)
		assert.Equal(t, 5, got)
	})

	t.Run("Should process all valid passports correctly", func(t *testing.T) {
		validPassports := ReadFromFile("fixtures/valid_passports.txt")

		got := GetPasswordCount(validPassports)
		assert.Equal(t, 5, got)
	})

	t.Run("Should process full input correctly invalid password", func(t *testing.T) {
		fullPassportSet := ReadFromFile("fixtures/fixtures/day4_input.txt")

		got := GetPasswordCount(fullPassportSet)
		assert.Equal(t, 5, got)
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
