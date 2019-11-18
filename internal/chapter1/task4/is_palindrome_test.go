package task4_test

import (
	"github.com/ivkalita/cti/internal/chapter1/task4"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	RunTests(t, task4.IsPalindrome)
}
