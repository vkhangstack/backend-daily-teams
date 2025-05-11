package utils

import "strings"

func SplitName(name string) (string, string) {
	strings.TrimSpace(name)
	parts := strings.Fields(name)

	firstName := ""
	lastName := parts[0]

	for i := 1; i < len(parts); i++ {
		firstName = firstName + " " + parts[i]
	}

	return strings.TrimSpace(firstName), strings.TrimSpace(lastName)
}
