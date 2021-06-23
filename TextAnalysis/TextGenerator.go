package main

import "strings"

func continuePhrase(ngrams map[string]string, phraseBeginning string, wordsCount int) string {
	words := strings.Split(phraseBeginning, " ")

	for i := 0; i < wordsCount; i++ {
		twoLastWords := ""
		if len(words) >= 2 {
			twoLastWords = words[len(words)-2] + " " + words[len(words)-1]
		}
		words = nextWords(words, ngrams, twoLastWords)
	}

	return strings.Join(words, " ")
}

func nextWords(words []string, ngrams map[string]string, twoLastWords string) []string {
	if _, ok := ngrams[twoLastWords]; ok && len(words) >= 2 {
		words = append(words, ngrams[twoLastWords])
	} else if _, ok := ngrams[words[len(words)-1]]; ok {
		words = append(words, ngrams[words[len(words)-1]])
	}

	return words
}
