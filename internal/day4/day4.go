package day4

import (
	"regexp"
	"strings"
)

func GetPasswordCount(entries []string) int {
	passportsFound := 0
	fieldsIdentified := 0

	validationRules := map[string]string{
		"byr": `byr:(19[2-9][0-9]|200[0-2])\s`,
		"iyr": `iyr:(201[0-9]|2020)\s`,
		"eyr": `eyr:(202[0-9]|2030)\s`,
		"hgt": `hgt:(?:1(?:[5-8][0-9]|9[0-3])cm|(?:59|6[0-9]|7[0-6])in)\s`,
		"hcl": `hcl:#[a-f0-9]{6}\s`,
		"ecl": `ecl:(amb|blu|brn|gry|grn|hzl|oth)\s`,
		"pid": `pid:([0-9]{9})\s`,
	}

	for _, entry := range entries {
		for _, d := range regexp.MustCompile(`(\w+):([^\s ]+)*`).FindAllString(entry, -1) {

			data := d + " " // Ugly fix, to make sure regex matches correctly
			passportAttrbibutes := strings.Split(data, ":")

			// For every passport attribute, other than CID go match against regex map
			if passportAttrbibutes[0] != "cid" {
				repeatPatternMatch := regexp.MustCompile(validationRules[passportAttrbibutes[0]]).FindStringSubmatch(data)

				if len(repeatPatternMatch) > 0 {
					fieldsIdentified++
				} else {
					print("test")
				}
			}
		}

		if entry == "" {
			if fieldsIdentified == 7 {
				passportsFound++
			}
			fieldsIdentified = 0
		}
	}

	return passportsFound
}
