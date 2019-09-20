package main

import (
	"fmt"
	"os"

	"github.com/patrick-huber-pivotal/ymldiff/formatters"

	"github.com/patrick-huber-pivotal/ymldiff/diff"
)

func main() {
	if len(os.Args) != 3 {
		printHelpAndExit()
	}

	fromFile := os.Args[1]
	toFile := os.Args[2]

	_, err := os.Stat(fromFile)
	if err != nil {
		printErrorAndExit(err.Error())
	}

	_, err = os.Stat(toFile)
	if err != nil {
		printErrorAndExit(err.Error())
	}

	changeLog, err := diff.NewChangeLogFromFiles(fromFile, toFile)
	if err != nil {
		printErrorAndExit(err.Error())
	}

	formatter := formatters.NewBOSH(changeLog)
	formatter.Write(os.Stdout)
}

func printErrorAndExit(message string) {
	os.Stderr.WriteString(message)
	os.Exit(1)
}

func printHelpAndExit() {
	printHelp()
	os.Exit(1)
}

func printHelp() {
	fmt.Printf("usage:\n%s <from> <to>", os.Args[0])
	fmt.Println()
}
