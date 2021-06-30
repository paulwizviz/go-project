package mylib

import (
	"strings"
	"testing"
)

func TestForSentence(t *testing.T) {
	expected := "Hello World and Earth."
	actual := FormSentence("Hello", "World", "and", "Earth")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("Expected: %s Actual: %s", expected, actual)
	}
}
