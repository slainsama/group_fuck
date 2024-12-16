package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunShellCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("command execution failed: %w", err)
	}
	return stdout.String(), nil
}
