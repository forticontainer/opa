package util

import (
	"bytes"
	"os/exec"
	"strings"
)

// ExecuteShell executes the shell command and return result back
func ExecuteShell(command string) (string, error) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("/bin/sh")
	cmd.Stdin = strings.NewReader(command)
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	// run command and return result
	if err := cmd.Run(); err != nil {
		return errOut.String(), err
	}

	return out.String(), nil
}
