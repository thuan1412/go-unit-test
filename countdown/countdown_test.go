package countdown

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}

	CountDown(buffer)

	got := buffer.String()

	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
