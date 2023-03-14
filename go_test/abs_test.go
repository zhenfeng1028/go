package test

import (
	"testing"
)

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
	t.Logf("The abs of -1 is: %d", got)
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// go test abs_test.go -v
