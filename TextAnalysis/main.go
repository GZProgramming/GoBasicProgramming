package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//TODO
	parsedSentences := parseSentences("abc")
	fmt.Println(parsedSentences)

}

func parseSentences(text string) [][]string {
	var parsedSentences = make([][]string, 0)

	splitFunc := func(r rune) bool {
		return strings.ContainsRune(".!?;:()", r)
	}

	sentences := strings.FieldsFunc(text, splitFunc)
	for _, sentence := range sentences {
		words := parseWords(sentence)
		if len(words) > 0 {
			parsedSentences = append(parsedSentences, words)
		}
	}
	return parsedSentences
}

func parseWords(sentence string) []string {
	words := make([]string, 0)
	var builder strings.Builder
	builder.Grow(len(sentence))
	for _, symbol := range sentence {
		if isTrueSymbol(symbol) {
			builder.WriteRune(symbol)
		} else if builder.Len() > 0 {
			words = append(words, strings.ToLower(builder.String()))
			builder.Reset()
		}
	}
	if builder.Len() > 0 {
		words = append(words, strings.ToLower(builder.String()))
	}
	return words
}

func isTrueSymbol(s rune) bool {
	return unicode.IsLetter(s) || s == '\''
}
