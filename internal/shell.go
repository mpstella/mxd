package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func RunShellCommand(name string, args ...string) (bytes.Buffer, bytes.Buffer, error) {

	cmd := exec.Command(name, args...)
	var stdOut, stdErr bytes.Buffer

	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	if err := cmd.Run(); err != nil {
		fmt.Printf("An error has occurred with %s\n", err)

		for _, s := range strings.Split(stdErr.String(), "\n") {
			fmt.Println(s)
		}
		return stdOut, stdErr, err
	}

	for _, s := range strings.Split(stdOut.String(), "\n") {
		fmt.Println(s)
	}
	return stdOut, stdErr, nil
}
