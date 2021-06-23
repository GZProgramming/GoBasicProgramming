package main

import (
	"strings"
	"testing"
)

func TestReturnEmptyDictionary_OnEmptyText(t *testing.T) {
	text := ""
	parsedText := parseSentences(text)
	expected := make(map[string]string)
	actual := getMostFrequentNextWords(parsedText)

	AssertResult(t, expected, actual, text)
}

func TestReturnEmptyDictionary_OnTextWithOneSentenceWithOneWord(t *testing.T) {
	text := "abc"
	parsedText := parseSentences(text)
	expected := make(map[string]string)
	actual := getMostFrequentNextWords(parsedText)

	AssertResult(t, expected, actual, text)
}

func TestReturnCorrectResult_OnTextWithOneSentenceWithTwoWords(t *testing.T) {
	text := "x y"
	parsedText := parseSentences(text)
	expected := map[string]string{"x": "y"}
	actual := getMostFrequentNextWords(parsedText)

	AssertResult(t, expected, actual, text)
}

func TestReturnCorrectResult_OnTextWithOneSentenceWithMultipleWords(t *testing.T) {
	text := "x y z"
	parsedText := parseSentences(text)
	expected := map[string]string{"x": "y", "y": "z", "x y": "z"}
	actual := getMostFrequentNextWords(parsedText)

	AssertResult(t, expected, actual, text)
}

func TestReturnCorrectResult_OnTextWithTwoSentencesWithOneWord(t *testing.T) {
	text := "x.y"
	parsedText := parseSentences(text)
	expected := make(map[string]string)
	actual := getMostFrequentNextWords(parsedText)

	AssertResult(t, expected, actual, text)
}

func TestReturnResult_WithMostFrequentBigrams(t *testing.T) {
	texts := []string{"x y. x z. x y.", "x z. x y. x y", "x y. x y.", "x y"}
	for _, text := range texts {
		parsedText := parseSentences(text)
		expected := map[string]string{"x": "y"}
		actual := getMostFrequentNextWords(parsedText)

		AssertResult(t, expected, actual, text)
	}
}

func TestReturnResult_WithLexicographicallyFirstBigram_IfBigramsHaveSameFrequency(t *testing.T) {
	texts := []string{"x y. x z.", "x z. x y.", "x y. x yy.", "x yy. x y"}
	for _, text := range texts {
		parsedText := parseSentences(text)
		expected := map[string]string{"x": "y"}
		actual := getMostFrequentNextWords(parsedText)

		AssertResult(t, expected, actual, text)
	}
}

func AssertResult(t *testing.T, expected map[string]string, actual map[string]string, text string) {
	for key := range expected {
		if _, ok := actual[key]; !ok {
			t.Errorf("Input text: %s\nMissing expected key %s in dictionary\n", text, key)
			t.Fail()
		}
		if strings.Compare(expected[key], actual[key]) != 0 {
			t.Errorf("Input text: %s\nWrong value for key %s\n", text, key)
			t.Fail()
		}
	}

	for key := range actual {
		if _, ok := expected[key]; !ok {
			t.Errorf("Input text: %s\nKey %s should not be in dictionary\n", text, key)
			t.Fail()
		}
	}
}
