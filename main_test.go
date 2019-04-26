package main

import (
	"os/exec"
	"testing"
	"time"
)

func TestSuccessfulExecute(t *testing.T) {
	cmdName := "true"
	cmdArgs := []string{}
	returncode, _ := executeCommand(cmdName, cmdArgs)

	if returncode != 0 {
		t.Errorf("Returncode was incorrect, got: %d, want: %d.", returncode, 0)
	}
}

func TestFailedExecute(t *testing.T) {
	cmdName := "./fail_with_56.sh"
	cmdArgs := []string{}
	returncode, _ := executeCommand(cmdName, cmdArgs)

	if returncode != 56 {
		t.Errorf("Returncode was incorrect, got: %d, want: %d.", returncode, 56)
	}
}

func TestMultipleExecute(t *testing.T) {
	executions := []string{"true", "true"}
	returncode := runAllCommands(executions, true)

	if returncode != 0 {
		t.Errorf("Returncode was incorrect, got: %d, want: %d.", returncode, 0)
	}
}

func TestMultipleExecuteAndDontStopOnFailure(t *testing.T) {
	executions := []string{"./fail_with_56.sh", "true"}
	returncode := runAllCommands(executions, false)

	if returncode != 0 {
		t.Errorf("Returncode was incorrect, got: %d, want: %d.", returncode, 0)
	}
}

func TestMultipleExecuteAndStopOnFailure(t *testing.T) {
	executions := []string{"./fail_with_56.sh", "true"}
	returncode := runAllCommands(executions, true)
	if returncode != 56 {
		t.Errorf("Returncode was incorrect, got: %d, want: %d.", returncode, 56)
	}
}

func TestSIGTERMHandling(t *testing.T) {
	go func() {
		time.Sleep(2 * time.Second)

		cmd := exec.Command("/bin/bash", "./sigterm_commandwrapper.sh")
		cmd.Run()
	}()

	cmdName := "./endless.sh"
	cmdArgs := []string{}
	returncode, sigterm := executeCommand(cmdName, cmdArgs)

	if sigterm != true {
		t.Errorf("Sigterm not set, got: %t, want: %t.", sigterm, true)
	}
	if returncode != -1 {
		t.Errorf("Returncode was incorrect, got: %d, want: %d.", returncode, -1)
	}
}

func TestSignalHandlingWithMultilple(t *testing.T) {
	go func() {
		time.Sleep(2 * time.Second)

		cmd := exec.Command("/bin/bash", "./sigterm_commandwrapper.sh")
		cmd.Run()
	}()

	executions := []string{"./endless.sh", "true"}
	returncode := runAllCommands(executions, false)

	if returncode != -1 {
		t.Errorf("Returncode was incorrect, got: %d, want: %d.", returncode, -1)
	}
}
