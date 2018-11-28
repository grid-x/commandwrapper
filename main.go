package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type executionsArray []string

func (i *executionsArray) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *executionsArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var executions executionsArray
var stopOnFailure bool

func main() {
	flag.Var(&executions, "execute", "eg. -execute='/usr/local/bin/mybin test123 -o -a' -execute='ls -l'")
	flag.BoolVar(&stopOnFailure, "stop-on-failure", false, "If true, multiple execute steps get combined with && otherwise ||")
	flag.Parse()

	if len(executions) == 0 {
		fmt.Print("No commands found")
		os.Exit(1)
	}

	exitcode := runAllCommands(executions, stopOnFailure)

	os.Exit(exitcode)
}

func runAllCommands(executions executionsArray, stopOnFailure bool) int {
	exitcode := 0
	for _, element := range executions {
		s := strings.Split(element, " ")
		cmdName, cmdArgs := s[0], s[1:]

		exitcode = executeCommand(cmdName, cmdArgs)

		if stopOnFailure && exitcode != 0 {
			break
		}
	}
	return exitcode
}

func executeCommand(cmdName string, cmdArgs []string) int {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Command failed: ", err)

		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus()
		}

		fmt.Fprintln(os.Stderr, "Could not get exit code for failed program")
		return 1
	}
	return 0
}
