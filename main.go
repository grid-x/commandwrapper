package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var executions arrayFlags
var stopOnFailure bool

func main() {
	flag.Var(&executions, "execute", "eg. -execute='/usr/local/bin/mybin test123 -o -a' -execute='ls -l'")
	flag.BoolVar(&stopOnFailure, "stop-on-failure", false, "Should multiple execute steps combined with && or ||")
	flag.Parse()

	if len(executions) == 0 {
		fmt.Print("No commands found")
		os.Exit(1)
	}

	exitcode := 0
	for _, element := range executions {
		s := strings.Split(element, " ")
		cmdName, cmdArgs := s[0], s[1:]

		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, "Command failed: ", err)
			exitcode = 1
			if stopOnFailure {
				break
			}
		}
	}
	os.Exit(exitcode)
}
