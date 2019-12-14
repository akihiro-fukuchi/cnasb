package stringutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	in := "Hello, world"
	want := "dlrow ,olleH"

	got := Reverse(in)
	if got != want {
		t.Errorf("Reverse(%q) == %q, want %q", in, got, want)
	}
}
