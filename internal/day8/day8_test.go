package day8

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		t.Run("Should pass example input", func(t *testing.T) {

			data := []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			}

			res := RunProgram(data)
			assert.Equal(t, 5, res)

		})
		t.Run("should fetch correct result for input", func(t *testing.T) {

			data := ReadFromFile("fixtures/day8_input.txt")

			res := RunProgram(data)
			assert.Equal(t, 1675, res)

		})

	})
	t.Run("Part 2", func(t *testing.T) {
		t.Run("Should pass example input", func(t *testing.T) {

			data := []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			}

			res := RunProgramTwo(data)
			assert.Equal(t, 8, res)

		})
		t.Run("should fetch correct result for input", func(t *testing.T) {

			data := ReadFromFile("fixtures/day8_input.txt")

			res := RunProgramTwo(data)
			assert.Equal(t, 1532, res)

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
