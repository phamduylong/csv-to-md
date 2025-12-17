package main

import (
	"testing"
	"time"
)

func TestDurationToString(t *testing.T) {
	d, _ := time.ParseDuration("3.0337h")
	expected := "3 hours 2 minutes 1 second 320 ms"
	res := durationToReadableString(d)
	if res != expected {
		t.Errorf("Failed when formatting duration. Expected: %s, received %s", expected, res)
	}
}
