package main

import (
	"fmt"
	"os"
)

// main читает входной файл, обрабатывает текст
// и записывает результат в выходной файл.
func main() {
	if len(os.Args) < 3 { // проверяем, что указаны 2 файла
		fmt.Println("Erreur : il faut un fichier d'entrée et un fichier de sortie.")
		os.Exit(1) // останавливаем программу
	}

	inputFile := os.Args[1]  // имя входного файла
	outputFile := os.Args[2] // имя выходного файла

	data, err := os.ReadFile(inputFile) // читаем входной файл

	if err != nil { // если файл не удалось прочитать
		fmt.Println("Erreur de lecture :", err)
		os.Exit(1)
	}

	text := string(data) // превращаем содержимое файла в string

	text = transformText(text) // запускаем все правила обработки текста

	err = os.WriteFile(outputFile, []byte(text), 0644) // записываем результат в файл

	if err != nil { // если не получилось записать файл
		fmt.Println("Erreur d'écriture :", err)
		os.Exit(1)
	}

	fmt.Println("Terminé :", outputFile) // сообщаем, что программа закончила работу
}
