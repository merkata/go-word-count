package main

import (
	"testing"
	"bytes"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("one two three four\n")
	expected := 4
	actual := count(b)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, actual)
	}
}
