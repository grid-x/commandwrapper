package main

import (
	"testing"
)

func TestSuccessfulExecute(t *testing.T) {
	cmdName := "true"
	cmdArgs := []string{}
	returncode := executeCommand(cmdName, cmdArgs)

	if returncode != 0 {
		t.Errorf("Returncode was incorrect, got: %d, want: %d.", returncode, 0)
	}
}

func TestSFailedExecute(t *testing.T) {
	cmdName := "./fail_with_56.sh"
	cmdArgs := []string{}
	returncode := executeCommand(cmdName, cmdArgs)

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
