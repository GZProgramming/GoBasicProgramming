package main

import (
	"strings"
	"testing"
)

func TestContinuePhrase_StopWhenNoBigrammsAndTrigramms(t *testing.T) {

	ngrams := map[string]string{"x": "y", "y": "q", "x y": "z"}
	actual := continuePhrase(ngrams, "x", 4)

	if strings.Compare(actual, "x y z") != 0 {
		t.Error()
	}
}
