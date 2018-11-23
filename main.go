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

var myFlags arrayFlags

func main() {
	flag.Var(&myFlags, "execute", "eg. -execute='/usr/local/bin/mybin test123 -o -a' -execute='ls -l'")
	flag.Parse()

	if len(myFlags) == 0 {
		fmt.Print("No commands found")
		os.Exit(1)
	}

	for _, element := range myFlags {
		s := strings.Split(element, " ")
		cmdName, cmdArgs := s[0], s[1:]

		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, "Command failed: ", err)
			os.Exit(1)
		}
	}
}
