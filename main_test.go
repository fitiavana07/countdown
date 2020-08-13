package main

import (
	"testing"

	"bytes"
	"reflect"
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

	// printed correct
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	// sleep 4 times
	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper want 4 got %d", spySleeper.Calls)
	}

	// write - sleep - write - sleep
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

	})
}
