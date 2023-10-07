package internal

import (
	"strings"
)

type Target struct {
	Name         string
	Dependencies []string
	Commands     []string
}

func splitTargets(file_lines []string) [][]string {
	separatedTargets := make([][]string, 0)
	for _, line := range file_lines {
		if line == "" {
			continue
		}
		if line[0] != '\t' {
			// Append New Target
			newTargetLines := make([]string, 1)
			newTargetLines[0] = line
			separatedTargets = append(separatedTargets, newTargetLines)
		} else {
			// Append Commands to the last Target
			separatedTargets[len(separatedTargets)-1] = append(separatedTargets[len(separatedTargets)-1], line)
		}

	}
	return separatedTargets
}

func ParseContent(Data []byte) (TargetsList []Target) {
	TargetsList = make([]Target, 0)
	charData := string(Data)
	lines := strings.Split(charData, "\n")
	separatedTargets := splitTargets(lines)
	for _, target := range separatedTargets {
		addedTarget := Target{}

		line := target[0]
		lineSegments := strings.Split(line, ":")
		addedTarget.Name = lineSegments[0]
		dependenciesNames := strings.Trim(lineSegments[1], " ")

		addedTarget.Dependencies = make([]string, 0)
		if len(dependenciesNames) >= 1 {
			for _, dep := range strings.Split(dependenciesNames, " ") {
				addedTarget.Dependencies = append(addedTarget.Dependencies, strings.Trim(dep, " \t"))
			}
		}
		for _, cmd := range target[1:] {
			trimmedCommand := strings.Trim(cmd, " \t")
			if len(trimmedCommand) > 0 {
				addedTarget.Commands = append(addedTarget.Commands, trimmedCommand)
			}
		}

		TargetsList = append(TargetsList, addedTarget)
	}
	loadTargets(TargetsList)
	return TargetsList
}
