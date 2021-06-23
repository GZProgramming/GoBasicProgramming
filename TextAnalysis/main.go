package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("HarryPotterText.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}

	parsedSentences := parseSentences(string(data))
	ngrams := getMostFrequentNextWords(parsedSentences)

	for {
		fmt.Println("Введите первое слово (например, harry): ")
		beginning := ""
		fmt.Fscan(os.Stdin, &beginning)
		if beginning == "" {
			break
		}
		phrase := continuePhrase(ngrams, strings.ToLower(beginning), 10)
		fmt.Println(phrase)
	}
}
