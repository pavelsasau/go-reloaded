package main

import (
	"strconv"
	"strings"
)

// splitText разбивает обычный текст на отдельные слова.
func splitText(text string) []string {
	return strings.Fields(text)
}

// applyUp делает предыдущее слово большими буквами.
// Пример: hello (up) -> HELLO
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

// applyLow делает предыдущее слово маленькими буквами.
// Пример: HELLO (low) -> hello
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

// applyCap делает первую букву предыдущего слова большой.
// Пример: hELLO (cap) -> Hello
func applyCap(words []string) []string {
	result := []string{}

	for _, word := range words {
		if word == "(cap)" {
			if len(result) > 0 {
				lastIndex := len(result) - 1
				lastWord := result[lastIndex]

				lowerWord := strings.ToLower(lastWord)

				if len(lowerWord) > 0 {
					capWord := strings.ToUpper(lowerWord[:1]) + lowerWord[1:]
					result[lastIndex] = capWord
				}
			}
		} else {
			result = append(result, word)
		}
	}

	return result
}

// applyUpN делает N предыдущих слов большими буквами.
// Пример: hello my friend (up, 3) -> HELLO MY FRIEND
func applyUpN(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(up," && i+1 < len(words) {
			numberText := words[i+1]

			// "3)" -> "3"
			numberText = strings.TrimSuffix(numberText, ")")

			// "3" -> число 3
			number, err := strconv.Atoi(numberText)

			if err == nil && number > 0 {
				start := len(result) - number

				if start < 0 {
					start = 0
				}

				for j := start; j < len(result); j++ {
					result[j] = strings.ToUpper(result[j])
				}

				// Пропускаем элемент "3)"
				i++

				continue
			}
		}

		result = append(result, word)
	}

	return result
}

// applyLowN делает N предыдущих слов маленькими буквами.
// Пример: HELLO MY FRIEND (low, 3) -> hello my friend
func applyLowN(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(low," && i+1 < len(words) {
			numberText := words[i+1]

			// "3)" -> "3"
			numberText = strings.TrimSuffix(numberText, ")")

			// "3" -> число 3
			number, err := strconv.Atoi(numberText)

			if err == nil && number > 0 {
				start := len(result) - number

				if start < 0 {
					start = 0
				}

				for j := start; j < len(result); j++ {
					result[j] = strings.ToLower(result[j])
				}

				// Пропускаем элемент "3)"
				i++

				continue
			}
		}

		result = append(result, word)
	}

	return result
}

// applyCapN делает первую букву большой у N предыдущих слов.
// Пример: hello MY FRIEND (cap, 3) -> Hello My Friend
func applyCapN(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(cap," && i+1 < len(words) {
			numberText := words[i+1]

			// "3)" -> "3"
			numberText = strings.TrimSuffix(numberText, ")")

			// "3" -> число 3
			number, err := strconv.Atoi(numberText)

			if err == nil && number > 0 {
				start := len(result) - number

				if start < 0 {
					start = 0
				}

				for j := start; j < len(result); j++ {
					lowerWord := strings.ToLower(result[j])

					if len(lowerWord) > 0 {
						capWord := strings.ToUpper(lowerWord[:1]) + lowerWord[1:]
						result[j] = capWord
					}
				}

				// Пропускаем элемент "3)"
				i++

				continue
			}
		}

		result = append(result, word)
	}

	return result
}

// transformText запускает все правила преобразования текста.
func transformText(text string) string {
	words := splitText(text)

	words = applyUp(words)
	words = applyLow(words)
	words = applyCap(words)

	words = applyUpN(words)
	words = applyLowN(words)
	words = applyCapN(words)

	return strings.Join(words, " ")
}
