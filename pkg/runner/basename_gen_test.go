package runner

import (
	"fmt"
	"testing"
)

type fixedNumberGenerator struct{}

func (fng *fixedNumberGenerator) Intn(int) int { return 999 }

func TestGonelBasenameGenerator(t *testing.T) {
	__rnd = &fixedNumberGenerator{}
	__dtp = func() string {
		return "xxx"
	}
	expected := "gonel_xxx_0000999.go"
	out := generateBasename()

	if out != expected {
		fmt.Printf("Result %s isn't expected %s\n", out, expected)
		t.Fail()
	}
}
