package main

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	got := LetterCombinations("23")
	want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LetterCombinations(23) = %v, want %v", got, want)
	}

	if got := LetterCombinations(""); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("empty digits: %v", got)
	}
	if got := LetterCombinations("2"); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("single digit: %v", got)
	}
	// 7 and 9 carry four letters.
	if n := len(LetterCombinations("79")); n != 16 {
		t.Errorf("79 count = %d, want 16", n)
	}
	if got := LetterCombinations("1"); len(got) != 0 {
		t.Errorf("digit without letters: %v", got)
	}
}
