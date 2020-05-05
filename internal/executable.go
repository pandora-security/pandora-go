/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

import (
	"fmt"
	"os/exec"
)

type Executable struct {
	Name string `json:"name"`
}

func (executable *Executable) LookPath() string {
	executablePath, err := exec.LookPath(executable.Name)
	if err != nil {
		return ""
	}
	return executablePath
}

func (executable *Executable) Found() bool {
	return executable.LookPath() != ""
}

func (executable *Executable) Run(args ...string) (string, error) {
	command := exec.Command(executable.Name, args...)
	output, err := command.Output()
	if err != nil {
		return "", PandoraRuntimeError
	}
	if output == nil {
		return "", nil
	}
	return fmt.Sprintf("%s", output), nil
}

func NewExecutable(executableName string) *Executable {
	return &Executable{ Name: executableName }
}