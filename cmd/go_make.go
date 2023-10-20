package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/khaledhegazy222/go-make/internal"
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

	fileData, err := os.ReadFile(makefilePath)
	if err != nil {
		fmt.Printf("Error : %q\n", err.Error())
		return
	}
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
