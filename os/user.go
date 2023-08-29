package osinfo

import (
	"os/exec"
	"strings"
)

func Username() string {
	output, err := exec.Command("whoami").Output()
	if err != nil {
		return "Unknown"
	}
	username := strings.TrimSuffix(string(output), "\n")
	return username
}
