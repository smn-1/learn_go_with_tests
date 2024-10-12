package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	name := "Bob"
	err := Greet(&buffer, name)

	got := buffer.String()
	want := "Hello, Bob"

	if err != nil {
		t.Errorf("Greet failed: %v", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
