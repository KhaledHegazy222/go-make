package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func ContainsCycles(TargetsList []Target) bool {
	targetsSet := map[string]Target{}
	for _, target := range TargetsList {
		targetsSet[target.Name] = target
	}
	for _, target := range TargetsList {
		selectedTargets := map[string]bool{}
		if isCyclicTarget(target, selectedTargets, targetsSet) {
			return true
		}
	}
	return false
}

func isCyclicTarget(selectedTarget Target, scannedTargets map[string]bool, targetsSet map[string]Target) bool {
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
		if isCyclicTarget(targetsSet[dep], scannedTargets, targetsSet) {
			return true
		}
	}
	return false
}

func Execute(Target Target) {
	for _, command := range Target.Commands {
		commandSegments := strings.Split(command, " ")
		prog := commandSegments[0]
		args := ""
		if len(commandSegments) > 1 {
			args = strings.Join(commandSegments[1:], " ")
		}
		cmd := exec.Command(prog, args)
		out, err := cmd.Output()

		if err != nil {
			fmt.Println("could not run command: ", err)
		}
		fmt.Println(string(out))
	}
}
