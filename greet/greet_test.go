package greet_test

import (
	"testing"

	"github.com/koketani/playground-ci/greet"
)

func TestGreet(t *testing.T) {
	actual := greet.Greet("koketani")
	expected := "hello koketani"
	if actual != expected {
		t.Fatalf("expected %s, but %s", expected, actual)
	}
}
