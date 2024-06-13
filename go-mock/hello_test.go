package main

import (
	"testing"
)

// Testing is writen into the language; no external packages necessary
// Its best to be writing the tests first in many circumstances
func TestHello (t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}