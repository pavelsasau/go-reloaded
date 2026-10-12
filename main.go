package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Erreur : il faut un fichier d'entrée et un fichier de sortie.")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)

	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		os.Exit(1)
	}

	text := string(data)

	fmt.Println("Input :", inputFile)
	fmt.Println("Output :", outputFile)
	fmt.Println(text)
}
