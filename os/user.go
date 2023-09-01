package osinfo

import (
	"os/exec"
	"strings"
)

// Username returns the username of the current user
func Username() string {
	// get the username from the whoami command
	output, err := exec.Command("whoami").Output()
	if err != nil {
		return "Unknown"
	}
	// remove the trailing newline
	username := strings.TrimSuffix(string(output), "\n")
	return username
}
