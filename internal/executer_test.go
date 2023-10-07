package internal

import (
	"os"
	"testing"
)

func TestContainsCycles(t *testing.T) {
	t.Run("Test No Cycles Validation", func(t *testing.T) {
		data, err := os.ReadFile("../Makefile_No_Cycle")
		if err != nil {
			t.Errorf("%q", err.Error())
			return
		}
		parsedTargets := ParseContent(data)
		result := ContainsCycles(parsedTargets)
		expected := false
		if result != expected {
			t.Errorf("Unmatched Result Expected:\n %t\nfound:\n %t\n", expected, result)
		}

	})
	t.Run("Test Cycles Validation", func(t *testing.T) {
		data, err := os.ReadFile("../Makefile_Cycles")
		if err != nil {
			t.Errorf("%q", err.Error())
			return
		}
		parsedTargets := ParseContent(data)
		result := ContainsCycles(parsedTargets)
		expected := true
		if result != expected {
			t.Errorf("Unmatched Result Expected:\n %t\nfound:\n %t\n", expected, result)
		}

	})
}
