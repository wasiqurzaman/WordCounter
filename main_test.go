package main

import (
	"bytes"
	"testing"
)

// TestCountLines tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("hello world how are you?\n")

	exp := 5
	res := count(b, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("This is line1\n This is line2\n And this is line3")

	exp := 3
	res := count(b, true)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
