package internal

import (
	"errors"
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
		if isCyclicTarget(targetsSet[dep], scannedTargets) {
			return true
		}
	}
	return false
}

var errExecutableNotFound = errors.New("executable file not found")

func Execute(Target Target) (output []byte, err error) {
	output = []byte{}
	for _, dep := range Target.Dependencies {
		execOutput, execError := Execute(targetsSet[dep])
		if execError != nil {
			fmt.Println("Error: ", errExecutableNotFound)
			return nil, errExecutableNotFound
		}
		output = append(output, execOutput...)
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

		cmdOutput, cmdError := cmd.Output()
		if cmdError != nil {
			fmt.Println("Error: ", cmdError)
			return nil, cmdError
		}
		if !silentCommand {
			output = append(output, cmdOutput...)
		}
	}
	return output, nil
}
