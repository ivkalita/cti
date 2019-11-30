package task5_test

import (
	"github.com/ivkalita/cti/internal/chapter2/task5"
	"testing"
)

func TestSumReverse(t *testing.T) {
	RunTests(t, true, task5.SumReverse)
}

func TestSum(t *testing.T) {
	RunTests(t, false, task5.Sum)
}
