package day7

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

			data := []string{"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			}

			res := CheckHitCount("shinygold", data, 1)
			assert.Equal(t, 4, res, "Expected to find 151 bags")

		})
		t.Run("shiny gold bags", func(t *testing.T) {

			data := ReadFromFile("fixtures/day7_input.txt")

			res := CheckHitCount("shinygold", data, 1)
			assert.Equal(t, 153, res, "Expected to find 151 bags")

		})

	})
	t.Run("Part 2", func(t *testing.T) {
		t.Run("Should pass example input", func(t *testing.T) {

			data := []string{"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			}

			res := CheckHitCount("shinygold", data, 2)
			assert.Equal(t, 32, res, "Expected to find 32 bags")

		})

		t.Run("shiny gold bags", func(t *testing.T) {

			data := ReadFromFile("fixtures/day7_input.txt")

			res := CheckHitCount("shinygold", data, 2)
			assert.Equal(t, 153, res, "Expected to find 151 bags")

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
