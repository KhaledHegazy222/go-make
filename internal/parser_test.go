package internal

import (
	"os"
	"reflect"
	"testing"
)

// We need to Remove Test for private functions
func TestSplitTargets(t *testing.T) {
	t.Run("Test Normal Behavior no empty lines", func(t *testing.T) {
		result := splitTargets([]string{
			"go: all1",
			"\t@echo all1",
			"\techo Done",
			"all1: all2",
			"\techo all2",
			"\techo Done",
			"all2: all3",
			"\techo all3",
			"\techo Done",
			"all3:",
			"\techo all5",
			"\techo Done",
		})
		expected := [][]string{
			{
				"go: all1",
				"\t@echo all1",
				"\techo Done",
			}, {
				"all1: all2",
				"\techo all2",
				"\techo Done",
			}, {
				"all2: all3",
				"\techo all3",
				"\techo Done",
			}, {
				"all3:",
				"\techo all5",
				"\techo Done",
			},
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unmatched Result Expected:\n %q\nfound:\n %q\n", expected, result)
		}
	})

	t.Run("Test Empty Lines filtration", func(t *testing.T) {
		result := splitTargets([]string{
			"go: all1",
			"\t@echo all1",
			"",
			"",
			"\techo Done",
			"all1: all2",
			"",
			"",
			"\techo all2",
			"",
			"\techo Done",
		})
		expected := [][]string{
			{
				"go: all1",
				"\t@echo all1",
				"\techo Done",
			}, {
				"all1: all2",
				"\techo all2",
				"\techo Done",
			},
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unmatched Result Expected:\n %q\nfound:\n %q\n", expected, result)
		}
	})
}

func TestTargetParsing(t *testing.T) {
	t.Run("Parsing Single Target", func(t *testing.T) {
		data, err := os.ReadFile("../Makefile_Single_Target")
		if err != nil {
			t.Errorf("%q", err.Error())
			return
		}
		result := ParseContent(data)
		expected := []Target{
			{
				Name:         "go",
				Dependencies: []string{"all1"},
				Commands:     []string{"@echo \"all1\"", "echo Done"},
			},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unmatched Result Expected:\n %q\nfound:\n %q\n", expected, result)
		}

	})
	t.Run("Parsing Multiple Target per file", func(t *testing.T) {
		data, err := os.ReadFile("../Makefile_Multiple_Targets")
		if err != nil {
			t.Errorf("%q", err.Error())
			return
		}
		result := ParseContent(data)
		expected := []Target{
			{
				Name:         "go",
				Dependencies: []string{"all1", "all2", "all3"},
				Commands:     []string{"@echo all1", "echo Done"},
			}, {
				Name:         "all1",
				Dependencies: []string{"all2"},
				Commands:     []string{"echo all2", "echo Done"},
			}, {
				Name:         "all2",
				Dependencies: []string{"all3"},
				Commands:     []string{"echo all3", "echo Done"},
			}, {
				Name:         "all3",
				Dependencies: []string{},
				Commands:     []string{"echo all5", "echo Done"},
			},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unmatched Result Expected:\n %q\nfound:\n %q\n", expected, result)
		}

	})
}
