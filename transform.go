package main

import "strings"

func splitText(text string) []string {
	return strings.Fields(text)
}

func applyUp(words []string) []string {
	result := []string{}

	for _, word := range words {
		if word == "(up)" {
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
		} else {
			result = append(result, word)
		}
	}

	return result
}

func transformText(text string) string {
	words := splitText(text)
	words = applyUp(words)

	return strings.Join(words, " ")
}
