package gotest

import "testing"

func TestAdd(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected %d", a, b, actual, expected)
	}
}
