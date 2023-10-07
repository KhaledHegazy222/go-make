package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

var targetsSet map[string]Target

func loadTargets(TargetsList []Target) {
	targetsSet = map[string]Target{}
	for _, target := range TargetsList {
		targetsSet[target.Name] = target
	}
}

func ContainsCycles(TargetsList []Target) bool {

	for _, target := range TargetsList {
		selectedTargets := map[string]bool{}
		if isCyclicTarget(target, selectedTargets) {
			return true
		}
	}
	return false
}

func isCyclicTarget(selectedTarget Target, scannedTargets map[string]bool) bool {
	targetExist := scannedTargets[selectedTarget.Name]
	if targetExist {
		return true
	}
	scannedTargets[selectedTarget.Name] = true
	if len(selectedTarget.Dependencies) == 0 {
		return false
	}
	for _, dep := range selectedTarget.Dependencies {
		if targetExist := scannedTargets[dep]; targetExist {
			return true
		}
		if isCyclicTarget(targetsSet[dep], scannedTargets) {
			return true
		}
	}
	return false
}

func Execute(Target Target) {
	for _, dep := range Target.Dependencies {
		Execute(targetsSet[dep])
	}
	for _, command := range Target.Commands {
		commandSegments := strings.Split(command, " ")
		silentCommand := false
		prog := commandSegments[0]
		if prog[0] == '@' {
			silentCommand = true
			prog = prog[1:]
		}
		args := ""
		if len(commandSegments) > 1 {
			args = strings.Join(commandSegments[1:], " ")
		}
		cmd := exec.Command(prog, args)
		out, err := cmd.Output()

		if err != nil {
			fmt.Println("could not run command: ", err)
		}
		if !silentCommand {
			fmt.Print(string(out))
		}
	}
}
