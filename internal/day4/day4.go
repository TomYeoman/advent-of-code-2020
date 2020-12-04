package day4

import (
	"regexp"
	"strings"
)

func GetPasswordCount(entries []string) int {
	passportsFound := 0
	fieldsIdentified := 0

	for _, entry := range entries {

		repeatPatternMatch := regexp.MustCompile(`(\w+):([^\s ]+)*`).FindAllString(entry, -1)

		for _, data := range repeatPatternMatch {
			if strings.Split(data, ":")[0] != "cid" {
				fieldsIdentified++
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
