package day6

func GetGroupTotals(entries []string) []int {

	answerMap := make(map[string]int)
	groupCounts := []int{}
	groupCount := 0
	usersInGroup := 0

	// For each user
	for _, e := range entries {
		// For each answer
		for _, answer := range e {
			answerMap[string(answer)]++
		}

		if e == "" {
			groupCounts = append(groupCounts, 0)
			for _, count := range answerMap {
				if count == usersInGroup {
					groupCounts[groupCount]++
				}
			}
			groupCount++
			usersInGroup = 0
			// Reset map between runs
			answerMap = make(map[string]int)

		} else {
			usersInGroup++
		}
	}

	return groupCounts
}

func GetSumOfTotalsAcrossGroups(entries []string) int {
	results := GetGroupTotals(entries)
	total := 0

	for _, res := range results {
		total += res
	}
	return total

}
