package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadNumbersFromFile(path string) []int {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		check(err)
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func GetResult(matchCount int) int {
	const MATCH = 2020
	entries := ReadNumbersFromFile("day1_input.txt")

	for x := 0; x < len(entries); x++ {
		for y := x; y < len(entries); y++ {
			if entries[x]+entries[y] == MATCH {
				return entries[x] * entries[y]
			}

			if matchCount == 3 {
				for z := y; z < len(entries); z++ {
					if entries[x]+entries[y]+entries[z] == MATCH {
						return entries[x] * entries[y] * entries[z]
					}
				}
			}
		}
	}

	return 0
}
