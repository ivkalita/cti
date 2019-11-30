package task6_test

import (
	"github.com/ivkalita/cti/internal/chapter2/task6"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	RunTests(t, task6.IsPalindrome)
}
