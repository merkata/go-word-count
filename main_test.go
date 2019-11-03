package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("one two three four\n")
	expected := 4
	actual, _ := count(b, false, false)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("one two three\nfour five six\nsix six six\n")
	expected := 3
	actual, _ := count(b, true, false)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, actual)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("one two three\n")
	expected := 14
	actual, _ := count(b, false, true)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, actual)
	}
}

func TestCountInvalidArgs(t *testing.T) {
	b := bytes.NewBufferString("failing test\n")
	_, err := count(b, true, true)
	if err == nil {
		t.Errorf("Expected an error, received nil\n")
	}
}
