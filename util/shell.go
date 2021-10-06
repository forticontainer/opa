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

	// if errOut buffer is not empty, then it means something is wrong
	//if errOut.Len() != 0 {
	//	return "", errors.New(errOut.String())
	//}
	return out.String(), nil
}
