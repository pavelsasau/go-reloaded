package main

import (
	"strconv"
	"strings"
	"unicode"
)

// capitalizeWord делает первую букву большой, а остальные маленькими.
// Пример: hELLO -> Hello
func capitalizeWord(word string) string {
	runes := []rune(strings.ToLower(word)) // делаем слово маленьким и превращаем в список букв
	if len(runes) == 0 {                   // если слово пустое
		return word
	}
	runes[0] = unicode.ToUpper(runes[0]) // первую букву делаем большой
	return string(runes)                 // превращаем список букв обратно в string
}

// applyMarkers обрабатывает:
// (up), (low), (cap)
// (up, n), (low, n), (cap, n)
// (hex), (bin)
func applyMarkers(words []string) []string {
	result := []string{} // сюда складываем уже обработанные слова

	for i := 0; i < len(words); i++ { // проходим по всем словам
		word := words[i] // текущее слово

		if word == "(up)" || word == "(low)" || word == "(cap)" { // проверяем простой тег
			if len(result) > 0 { // проверяем, что есть предыдущее слово
				lastIndex := len(result) - 1 // индекс последнего слова
				switch word {                // смотрим, какой именно тег встретился
				case "(up)":
					result[lastIndex] = strings.ToUpper(result[lastIndex]) // делаем слово большими буквами
				case "(low)":
					result[lastIndex] = strings.ToLower(result[lastIndex]) // делаем слово маленькими буквами
				case "(cap)":
					result[lastIndex] = capitalizeWord(result[lastIndex]) // делаем первую букву большой
				}
			}
			continue // сам тег в result не добавляем
		}

		if (word == "(up," || word == "(low," || word == "(cap,") && i+1 < len(words) { // проверяем тег с числом
			numberText := words[i+1]                         // например "3)"
			numberText = strings.TrimSuffix(numberText, ")") // "3)" превращаем в "3"
			number, err := strconv.Atoi(numberText)          // текст "3" превращаем в число 3

			if err == nil && number > 0 { // проверяем, что число прочиталось правильно
				start := len(result) - number // индекс, с которого начинаем изменять слова
				if start < 0 {                // если слов меньше, чем указано в теге
					start = 0 // начинаем с первого слова
				}

				for j := start; j < len(result); j++ { // проходим по N предыдущим словам
					switch word { // смотрим, какой тег встретился
					case "(up,":
						result[j] = strings.ToUpper(result[j]) // делаем слово большими буквами
					case "(low,":
						result[j] = strings.ToLower(result[j]) // делаем слово маленькими буквами
					case "(cap,":
						result[j] = capitalizeWord(result[j]) // делаем первую букву большой
					}
				}

				i++      // пропускаем следующий элемент, например "3)"
				continue // переходим к следующему слову
			}
		}

		if word == "(hex)" && len(result) > 0 { // если встретили тег hex
			lastIndex := len(result) - 1                               // индекс предыдущего слова
			number, err := strconv.ParseInt(result[lastIndex], 16, 64) // переводим hex в число
			if err == nil {                                            // если преобразование прошло нормально
				result[lastIndex] = strconv.FormatInt(number, 10) // число превращаем обратно в обычный текст
			}
			continue // тег (hex) не добавляем
		}

		if word == "(bin)" && len(result) > 0 { // если встретили тег bin
			lastIndex := len(result) - 1                              // индекс предыдущего слова
			number, err := strconv.ParseInt(result[lastIndex], 2, 64) // переводим binary в число
			if err == nil {                                           // если преобразование прошло нормально
				result[lastIndex] = strconv.FormatInt(number, 10) // число превращаем обратно в текст
			}
			continue // тег (bin) не добавляем
		}

		result = append(result, word) // если это обычное слово, добавляем его в result
	}

	return result // возвращаем готовый список слов
}

// fixArticles меняет a на an, если следующее слово начинается на a, e, i, o, u или h.
// Пример: a apple -> an apple
func fixArticles(words []string) []string {
	for i := 0; i < len(words); i++ { // проходим по всем словам
		if !strings.EqualFold(words[i], "a") { // если текущее слово не a и не A
			continue
		}

		for j := i + 1; j < len(words); j++ { // ищем следующее настоящее слово
			next := strings.TrimLeft(words[j], "'\"") // убираем кавычки слева
			nextRunes := []rune(next)                 // превращаем слово в список букв

			if len(nextRunes) == 0 { // если после удаления кавычек ничего не осталось
				continue
			}

			firstLetter := unicode.ToLower(nextRunes[0])     // берём первую букву следующего слова
			if strings.ContainsRune("aeiouh", firstLetter) { // проверяем гласную или h
				if words[i] == "A" { // если было A с большой буквы
					words[i] = "An" // A -> An
				} else {
					words[i] = "an" // a -> an
				}
			}
			break // следующее настоящее слово уже нашли
		}
	}

	return words // возвращаем изменённый список
}

// isPunctuation проверяет, является ли символ знаком пунктуации.
func isPunctuation(r rune) bool {
	return strings.ContainsRune(".,!?:;", r) // возвращает true для нужных знаков
}

// formatPunctuation исправляет пробелы возле пунктуации.
// Пример: Hello , world ! -> Hello, world!
func formatPunctuation(text string) string {
	runes := []rune(text) // превращаем текст в список символов
	result := []rune{}    // сюда собираем готовый текст

	for i := 0; i < len(runes); i++ { // проходим по всем символам
		current := runes[i] // текущий символ

		if isPunctuation(current) { // если текущий символ — пунктуация
			for len(result) > 0 && result[len(result)-1] == ' ' { // пока перед знаком есть пробел
				result = result[:len(result)-1] // удаляем пробел перед пунктуацией
			}

			result = append(result, current) // добавляем сам знак пунктуации

			for i+1 < len(runes) && runes[i+1] == ' ' { // если после знака есть пробелы
				i++ // пропускаем лишние пробелы
			}

			if i+1 < len(runes) && !isPunctuation(runes[i+1]) && runes[i+1] != '\'' { // если дальше обычный символ
				result = append(result, ' ') // добавляем один правильный пробел
			}

			continue // переходим к следующему символу
		}

		result = append(result, current) // обычный символ просто добавляем
	}

	return strings.TrimSpace(string(result)) // убираем пробелы в начале и конце
}

// formatQuotes исправляет пробелы внутри одинарных кавычек.
// Пример: ' hello world ' -> 'hello world'
func formatQuotes(text string) string {
	runes := []rune(text) // превращаем текст в список символов
	result := []rune{}    // сюда собираем готовый текст
	insideQuote := false  // false = сейчас мы не внутри кавычек

	for i := 0; i < len(runes); i++ { // проходим по всем символам
		current := runes[i] // текущий символ

		if current != '\'' { // если это не кавычка
			result = append(result, current) // просто добавляем символ
			continue
		}

		if !insideQuote { // если это открывающая кавычка
			result = append(result, '\'') // добавляем открывающую кавычку
			insideQuote = true            // теперь мы внутри кавычек

			for i+1 < len(runes) && runes[i+1] == ' ' { // если после кавычки есть пробелы
				i++ // убираем пробел после открывающей кавычки
			}
		} else { // если это закрывающая кавычка
			for len(result) > 0 && result[len(result)-1] == ' ' { // если перед кавычкой есть пробел
				result = result[:len(result)-1] // удаляем пробел перед закрывающей кавычкой
			}

			result = append(result, '\'') // добавляем закрывающую кавычку
			insideQuote = false           // теперь мы вышли из кавычек
		}
	}

	return string(result) // возвращаем готовый текст
}

// transformText запускает все правила проекта по порядку.
func transformText(text string) string {
	words := strings.Fields(text)   // разбиваем текст на отдельные слова
	words = applyMarkers(words)     // обрабатываем up, low, cap, hex, bin и варианты с n
	words = fixArticles(words)      // исправляем a -> an
	text = strings.Join(words, " ") // соединяем слова обратно в один текст
	text = formatPunctuation(text)  // исправляем пробелы возле пунктуации
	text = formatQuotes(text)       // исправляем пробелы возле кавычек
	return text                     // возвращаем полностью готовый текст
}
