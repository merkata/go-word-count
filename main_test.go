package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("one two three four\n")
	expected := 4
	actual := count(b, false)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("one two three\nfour five six\nsix six six\n")
	expected := 3
	actual := count(b, true)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, actual)
	}
}
