package fortnight

import (
	"fortnight"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func TestAdd(t *testing.T) {
	x := fortnight.AddNumber(2, 3)
	if x != 5 {
		t.Fail()
	}
}

func TestFailedAdd(t *testing.T) {
	x := fortnight.AddNumber(3, 3)
	if x != 5 {
		t.Fail()
	}
}
