package main

import (
	"testing"

	"bytes"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	// Countdown will write the data into an io.Writer
	// os.Stdout for printed to user
	// bytes.Buffer for capturing during tests
	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
