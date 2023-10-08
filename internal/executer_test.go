package internal

import (
	"os"
	"testing"
)

func TestContainsCycles(t *testing.T) {
	t.Run("Test No Cycles Validation", func(t *testing.T) {
		parsedTargets := parseFile(t, "../Makefile_No_Cycle")
		result := ContainsCycles(parsedTargets)
		expected := false
		if result != expected {
			t.Errorf("Unmatched Result Expected:\n %t\nfound:\n %t\n", expected, result)
		}

	})

	t.Run("Test Cycles Validation", func(t *testing.T) {
		parsedTargets := parseFile(t, "../Makefile_Cycles")
		result := ContainsCycles(parsedTargets)
		expected := true
		if result != expected {
			t.Errorf("Unmatched Result Expected:\n %t\nfound:\n %t\n", expected, result)
		}

	})

}

func parseFile(t testing.TB, filePath string) []Target {
	t.Helper()
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("%q", err.Error())
		return nil
	}
	return ParseContent(data)
}

func TestExecute(t *testing.T) {
	t.Run("Test Execute", func(t *testing.T) {
		parsedTargets := parseFile(t, "../Makefile_Execute_Dep")
		out, err := Execute(parsedTargets[0])
		if err != nil {
			t.Errorf("Unexpected Error : %q\n", err.Error())
			return
		}
		result := string(out)
		expected := `First
Third
Second
all
`
		if result != expected {
			t.Errorf("Unmatched Result Expected:\n %q\nfound:\n %q\n", expected, result)
		}
	})
	t.Run("Test Execute Error", func(t *testing.T) {
		parsedTargets := parseFile(t, "../Makefile_Execute_Error")
		_, err := Execute(parsedTargets[0])
		if err != errExecutableNotFound {
			t.Errorf("Unmatched Error Expected:\n %q\nfound:\n %q\n", errExecutableNotFound.Error(), err.Error())
			return
		}
	})
}
