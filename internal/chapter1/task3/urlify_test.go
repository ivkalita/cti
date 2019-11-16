package task3_test

import (
	"github.com/ivkalita/cti/internal/chapter1/task3"
	"testing"
)

func TestUrlify(t *testing.T) {
	RunTests(t, task3.Urlify)
}
