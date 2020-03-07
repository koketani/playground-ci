package greet_test

import (
	"testing"

	"github.com/koketani/playground-ci/greet"
)

func TestGreet(t *testing.T) {
	actual := greet.Greet("Koketani")
	expected := "Koketani, hello, world."
	if actual != expected {
		t.Fatalf("expected %q, but %q", expected, actual)
	}
}
