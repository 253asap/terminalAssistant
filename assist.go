package main

import (
	commandfuncs "assist/commandFuncs"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		os.Exit(1)
	}
	command := flag.Arg(0)
	runCommand(command, flag.Args()[1:])
}

func runCommand(command string, args []string) error {
	switch command {
	case "hello":
		fmt.Println("Hello", args[0])
	case "GPTepal":
		commandfuncs.ChatGpt()
	default:
		fmt.Println("Command not valid")
	}
	return nil
}
