package main

import (
	"fmt"
	"testing"
)

func TestSimpleSentence(t *testing.T) {
	parsedSentences := parseSentences("abc")
	expected := make([][]string, 0)
	expected = append(expected, []string{"abc"})

	if len(parsedSentences) != len(expected) {
		t.Errorf("Expected length %d, but was %d", len(expected), len(parsedSentences))
	}
	if !AssertAllSentencesEqual(parsedSentences, expected) {
		t.Error("Wrong result!")
	}
}

func TestSimpleSentenceTwoWords(t *testing.T) {
	parsedSentences := parseSentences("a, bc")
	expected := make([][]string, 0)
	expected = append(expected, []string{"a", "bc"})

	if len(parsedSentences) != len(expected) {
		t.Errorf("Expected length %d, but was %d", len(expected), len(parsedSentences))
	}
	if !AssertAllSentencesEqual(parsedSentences, expected) {
		t.Error("Wrong result!")
	}
}

func TestReturnCorrectResult_OnTextWithOneSentence_WithWordContainingApostrophe(t *testing.T) {
	parsedSentences := parseSentences("it's")
	expected := make([][]string, 0)
	expected = append(expected, []string{"it's"})

	if len(parsedSentences) != len(expected) {
		t.Errorf("Expected length %d, but was %d", len(expected), len(parsedSentences))
	}
	if !AssertAllSentencesEqual(parsedSentences, expected) {
		t.Error("Wrong result!")
	}
}

func TestCorrectlyParse_SentenceDelimiters(t *testing.T) {
	parsedSentences := parseSentences("A.b!c?d : E;f(g)h;i")
	expected := make([][]string, 0)
	for _, sent := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"} {
		expected = append(expected, []string{sent})
	}

	if len(parsedSentences) != len(expected) {
		t.Errorf("Expected length %d, but was %d", len(expected), len(parsedSentences))
	}
	if !AssertAllSentencesEqual(parsedSentences, expected) {
		t.Error("Wrong result!")
	}
}

func AssertAllSentencesEqual(parsedSentences, expected [][]string) bool {
	for i := 0; i < len(parsedSentences); i++ {
		for j := 0; j < len(parsedSentences[i]); j++ {
			if parsedSentences[i][j] != expected[i][j] {
				fmt.Printf("Expected %s, but was %s!\n", string(expected[i][j]), parsedSentences[i][j])
				return false
			}
		}
	}
	return true
}
