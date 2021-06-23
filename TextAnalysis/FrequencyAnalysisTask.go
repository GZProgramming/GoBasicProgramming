package main

import (
	"strings"
)

func getMostFrequentNextWords(text [][]string) map[string]string {
	ngrams := make(map[string]map[string]int)
	for _, sentence := range text {
		for i := 0; i < len(sentence)-1; i++ {
			addNgrams(ngrams, sentence[i], sentence[i+1])
			if i+2 < len(sentence) {
				twoWords := sentence[i] + " " + sentence[i+1]
				addNgrams(ngrams, twoWords, sentence[i+2])
			}
		}
	}
	return mostFrequentWords(ngrams)
}

func addNgrams(ngrams map[string]map[string]int, word string, nextWord string) {
	if _, ok := ngrams[word]; !ok {
		ngrams[word] = make(map[string]int)
	}
	if _, ok := ngrams[word][nextWord]; !ok {
		ngrams[word][nextWord] = 0
	}
	ngrams[word][nextWord]++
}

func mostFrequentWords(ngrams map[string]map[string]int) map[string]string {
	result := make(map[string]string)
	for mapKey, mapLine := range ngrams {
		maxCount := 0
		minWord := "Harry Potter and the Sorcerer's Stone"
		for word, wordFrequency := range mapLine {
			if wordFrequency > maxCount ||
				wordFrequency == maxCount &&
					strings.Compare(minWord, word) > 0 {
				maxCount = wordFrequency
				minWord = word
			}
		}
		result[mapKey] = minWord
	}
	return result
}
