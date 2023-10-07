package main

import (
	"flag"
	"github.com/khaledhegazy222/go-make/internal"
	"os"
)

func main() {
	target := flag.String("t", "", "Target name")
	filePath := flag.String("f", "./Makefile", "Makefile Path")
	flag.Parse()

	args := os.Args

	makefilePath := *filePath
	if len(args) > 1 {
		makefilePath = args[1]
	}

	fileData, _ := os.ReadFile(makefilePath)
	parsedTargets := internal.ParseContent(fileData)
	var selectedTarget internal.Target
	if len(*target) > 0 {

		for _, targetItem := range parsedTargets {
			if targetItem.Name == *target {
				selectedTarget = targetItem
			}
		}
	} else {
		selectedTarget = parsedTargets[0]

		if !internal.ContainsCycles(parsedTargets) {
			internal.Execute(selectedTarget)
		}
	}
}
