package main

import (
	"testing"

	"bytes"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	// Countdown will write the data into an io.Writer
	// os.Stdout for printed to user
	// bytes.Buffer for capturing during tests
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper want 4 got %d", spySleeper.Calls)
	}
}
