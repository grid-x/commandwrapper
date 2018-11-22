package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
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
	var wg sync.WaitGroup
	flag.Var(&myFlags, "execute", "eg. -execute='/usr/local/bin/mybin test123 -o -a' -execute='ls -l'")
	flag.Parse()

	if len(myFlags) == 0 {
		fmt.Print("No commands found")
		os.Exit(1)
	}

	for _, element := range myFlags {
		s := strings.Split(element, " ")
		cmdName, cmdArgs := s[0], s[1:]

		var stdoutBuf, stderrBuf bytes.Buffer
		cmd := exec.Command(cmdName, cmdArgs...)

		stdoutIn, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create StdoutPipe: ", err)
			os.Exit(1)
		}
		stderrIn, err := cmd.StderrPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create StderrPipe: ", err)
			os.Exit(1)
		}

		var errStdout, errStderr error
		stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
		stderr := io.MultiWriter(os.Stderr, &stderrBuf)

		if err := cmd.Start(); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to start the command: ", err)
			os.Exit(1)
		}

		wg.Add(2)
		go func() {
			_, errStdout = io.Copy(stdout, stdoutIn)
			defer wg.Done()
		}()

		go func() {
			_, errStderr = io.Copy(stderr, stderrIn)
			defer wg.Done()
		}()

		if err := cmd.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, "Command failed: ", err)
			os.Exit(1)
		}
		if errStdout != nil || errStderr != nil {
			fmt.Fprintln(os.Stderr, "Failed to capture stdout or stderr", err)
			os.Exit(1)
		}
	}
	wg.Wait()
}
