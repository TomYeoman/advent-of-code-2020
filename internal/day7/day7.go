package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	children []*Node
	parents  []*Node
	quantity int
}

var uniqueValues = map[string]string{}
var childrenQuantity int

func GetParentNames(node *Node) {
	for _, parent := range node.parents {
		name := parent.name
		// Lazy - Get recursive working
		uniqueValues[name] = name
		GetParentNames(parent)
	}
}

// TODO - Get working
func GetChildrenQuantity(node *Node) {
	for _, child := range node.children {
		GetChildrenQuantity(child)
	}
}

func CheckHitCount(searchterm string, data []string, part int) int {
	// Build tree

	bagEntries := make(map[string]*Node)

	for _, line := range data {

		// Sanitize line a little, to make it easier to regex
		// Before = "light red bags contain 1 bright white bag, 2 muted yellow bags.""
		// After = "lightredbags,1brightwhitebag,2mutedyellowbags,"
		line = strings.NewReplacer("contain", ",", ".", ",", " ", "", "bags", "", "bag", "").Replace(line)

		// Split on commas
		result := regexp.MustCompile(`(.*?,)`).FindAllString(line, -1)

		outerBagName := ""
		for _, d := range result {
			// Extract the count, and color name
			res := regexp.MustCompile(`([^,\d]+)|([\d]+)`).FindAllString(d, -1)

			// Outer bag
			if len(res) == 1 {
				outerBagName = res[0]
				// Create outer bag, only if it's a root
				if _, isOk := bagEntries[outerBagName]; !isOk {
					bagEntries[outerBagName] = &Node{
						name:     outerBagName,
						quantity: 1,
					}
				}

			} else {
				innerBagName := res[1]
				innerBagQuantity, _ := strconv.Atoi(res[0])

				if _, isOk := bagEntries[innerBagName]; !isOk {
					bagEntries[innerBagName] = &Node{
						name:     innerBagName,
						quantity: 1,
					}
				}

				bagEntries[innerBagName].quantity = innerBagQuantity
				bagEntries[innerBagName].parents = append(bagEntries[innerBagName].parents, bagEntries[outerBagName])
				bagEntries[outerBagName].children = append(bagEntries[outerBagName].children, bagEntries[innerBagName])

				fmt.Printf("bag %q, quantity %d \n", innerBagName, innerBagQuantity)
			}
		}

	}

	if part == 1 {
		GetParentNames(bagEntries[searchterm])
		return len(uniqueValues)
	} else {
		GetChildrenQuantity(bagEntries[searchterm])
		return childrenQuantity
	}

	return 0
}
