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
				lastIndex := len(result) - 1
				lastWord := result[lastIndex]
				upperWord := strings.ToUpper(lastWord)
				result[lastIndex] = upperWord
			}
		} else {
			result = append(result, word)
		}
	}

	return result
}

func applyLow(words []string) []string {
	result := []string{}

	for _, word := range words {
		if word == "(low)" {
			if len(result) > 0 {
				lastIndex := len(result) - 1
				lastWord := result[lastIndex]
				lowerWord := strings.ToLower(lastWord)
				result[lastIndex] = lowerWord
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
	words = applyLow(words)
	words = applyCap(words)

	return strings.Join(words, " ")
}

func applyCap(words []string) []string {
	result := []string{}

	for _, word := range words {
		if word == "(cap)" {
			if len(result) > 0 {
				lastIndex := len(result) - 1
				lastWord := result[lastIndex]

				lowerWord := strings.ToLower(lastWord)
				capWord := strings.ToUpper(lowerWord[:1]) + lowerWord[1:]

				result[lastIndex] = capWord
			}
		} else {
			result = append(result, word)
		}
	}

	return result
}
