package day1

import (
	"regexp"
	"strconv"
)

func GetResult(entries []string, part int) int {

	validPasswords := 0

	for _, entry := range entries {

		repeatPatternMatch := regexp.MustCompile(`(\d+)-(\d+)`).FindStringSubmatch(entry)
		lowerCount, _ := strconv.Atoi(repeatPatternMatch[1])
		upperCount, _ := strconv.Atoi(repeatPatternMatch[2])

		requiredCharacterMatch := regexp.MustCompile(`(\w+):`).FindStringSubmatch(entry)
		requiredCharacter := requiredCharacterMatch[1]

		passwordMatch := regexp.MustCompile(`\s(\w+)$`).FindStringSubmatch(entry)
		password := passwordMatch[1]

		if part == 1 {
			requiredCharactersHit := 0

			for _, c := range password {
				if string(c) == requiredCharacter {
					requiredCharactersHit++
				}
			}

			if requiredCharactersHit >= lowerCount && requiredCharactersHit <= upperCount {
				validPasswords++
			}
		}

		if part == 2 {
			if (string(password[lowerCount-1]) == requiredCharacter) != (string(password[upperCount-1]) == requiredCharacter) {
				validPasswords++
			}
		}

	}
	return validPasswords
}
